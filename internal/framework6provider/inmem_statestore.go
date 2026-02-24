// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package framework

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
	"sync"
	"testing/fstest"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/statestore"
	"github.com/hashicorp/terraform-plugin-framework/statestore/schema"
)

const (
	defaultLockFileName  = "terraform.tfstate.tflock"
	defaultStateFileName = "terraform.tfstate"
)

var _ statestore.StateStore = &InMemStateStore{}

// InMemStateStore implements a Terraform state store that keeps all state
// and lock files in memory, which is only useful for testing purposes.
type InMemStateStore struct {
	memFS fstest.MapFS
	mu    sync.RWMutex
}

func NewInMemStateStore(memFS fstest.MapFS) statestore.StateStore {
	return &InMemStateStore{
		memFS: memFS,
	}
}

func (s *InMemStateStore) Metadata(ctx context.Context, req statestore.MetadataRequest, resp *statestore.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inmem"
}

func (s *InMemStateStore) Schema(ctx context.Context, req statestore.SchemaRequest, resp *statestore.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "An in-memory state store for testing purposes. All state and lock files are stored in memory " +
			"and will be lost when the provider process ends.",
		Attributes: map[string]schema.Attribute{
			"region": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("us-east-1", "us-west-2"),
				},
			},
		},
	}
}

func (s *InMemStateStore) Initialize(ctx context.Context, req statestore.InitializeRequest, resp *statestore.InitializeResponse) {
	// No initialization needed for in-memory store
}

func (s *InMemStateStore) GetStates(ctx context.Context, req statestore.GetStatesRequest, resp *statestore.GetStatesResponse) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	directories, err := s.memFS.ReadDir(".")
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading In-Memory Filesystem",
			fmt.Sprintf("Could not read state store directory: %s", err),
		)
		return
	}

	workspaces := make([]string, 0, len(directories))
	for _, dir := range directories {
		workspaces = append(workspaces, filepath.Base(dir.Name()))
	}

	resp.StateIDs = workspaces
}

func (s *InMemStateStore) DeleteState(ctx context.Context, req statestore.DeleteStateRequest, resp *statestore.DeleteStateResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for filePath := range s.memFS {
		if strings.HasPrefix(filePath, req.StateID) {
			delete(s.memFS, filePath)
			return
		}
	}

	resp.Diagnostics.AddError(
		"Error Deleting State",
		fmt.Sprintf("Provider was asked to delete state with ID %q, which doesn't exist", req.StateID),
	)
}

func (s *InMemStateStore) Lock(ctx context.Context, req statestore.LockRequest, resp *statestore.LockResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()

	lockFilePath := filepath.Join(req.StateID, defaultLockFileName)

	if lockFile, lockExists := s.memFS[lockFilePath]; lockExists {
		var existingLock statestore.LockInfo
		if err := json.Unmarshal(lockFile.Data, &existingLock); err != nil {
			resp.Diagnostics.AddError(
				"Error Reading Existing Lock",
				fmt.Sprintf("Could not unmarshal lock data: %s", err),
			)
			return
		}

		resp.Diagnostics.Append(statestore.WorkspaceAlreadyLockedDiagnostic(req, existingLock))
		return
	}

	lockInfo := statestore.NewLockInfo(req)
	lockBytes, err := json.Marshal(lockInfo)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Lock",
			fmt.Sprintf("Could not marshal lock data: %s", err),
		)
		return
	}

	s.memFS[lockFilePath] = &fstest.MapFile{
		Data:    lockBytes,
		Mode:    fs.ModePerm,
		ModTime: time.Now(),
	}

	resp.LockID = lockInfo.ID
}

func (s *InMemStateStore) Unlock(ctx context.Context, req statestore.UnlockRequest, resp *statestore.UnlockResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()

	lockFilePath := filepath.Join(req.StateID, defaultLockFileName)

	lockFile, lockExists := s.memFS[lockFilePath]
	if !lockExists {
		resp.Diagnostics.AddError(
			"Lock Does Not Exist",
			fmt.Sprintf("Workspace %q has already been unlocked or was never locked.", req.StateID),
		)
		return
	}

	var existingLock statestore.LockInfo
	if err := json.Unmarshal(lockFile.Data, &existingLock); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Lock",
			fmt.Sprintf("Could not unmarshal lock data: %s", err),
		)
		return
	}

	if existingLock.ID != req.LockID {
		resp.Diagnostics.AddError(
			"Lock ID Mismatch",
			fmt.Sprintf(
				"Workspace %q is locked with ID %q, but unlock was attempted with ID %q. "+
					"This likely indicates a bug in the Lock method or concurrent access to the same workspace.",
				req.StateID, existingLock.ID, req.LockID,
			),
		)
		return
	}

	delete(s.memFS, lockFilePath)
}

func (s *InMemStateStore) Read(ctx context.Context, req statestore.ReadRequest, resp *statestore.ReadResponse) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	stateFilePath := filepath.Join(req.StateID, defaultStateFileName)
	stateFile, err := s.memFS.Open(stateFilePath)

	if err != nil && errors.Is(err, fs.ErrNotExist) {
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading State File",
			fmt.Sprintf("Could not open state file for workspace %q: %s", req.StateID, err),
		)
		return
	}

	stateBytes, err := io.ReadAll(stateFile)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading State Data",
			fmt.Sprintf("Could not read state data for workspace %q: %s", req.StateID, err),
		)
		return
	}

	resp.StateBytes = stateBytes
}

func (s *InMemStateStore) Write(ctx context.Context, req statestore.WriteRequest, resp *statestore.WriteResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()

	stateFilePath := filepath.Join(req.StateID, defaultStateFileName)
	s.memFS[stateFilePath] = &fstest.MapFile{
		Data:    req.StateBytes,
		Mode:    fs.ModePerm,
		ModTime: time.Now(),
	}
}
