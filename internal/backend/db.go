package backend

import (
	"fmt"

	"github.com/hashicorp/go-memdb"
)

// User represents a user in the database.
type User struct {
	// Email must be unique, and is treated as the User's UUID.
	Email string
	Name  string
	Age   int
}

// Region represents an availability region in the database.
type Region struct {
	Name string
}

// A Client manages communication with the memdb.
type Client struct {
	db *memdb.MemDB
}

var db *memdb.MemDB

func init() {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"users": {
				Name: "users",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
					"age": {
						Name:    "age",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "Age"},
					},
				},
			},
			"regions": {
				Name: "regions",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
		},
	}

	var err error
	db, err = memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	// populate regions for use in data source
	txn := db.Txn(true)
	defer txn.Abort()

	regions := []*Region{{Name: "UK"}, {Name: "EU"}, {Name: "USA"}}
	for _, r := range regions {
		err := txn.Insert("regions", r)
		if err != nil {
			panic(err)
		}
	}

	txn.Commit()
}

// NewClient returns a new memdb client.
func NewClient() (*Client, error) {
	c := &Client{}

	c.db = db

	return c, nil
}

func (c *Client) CreateUser(user *User) error {
	txn := c.db.Txn(true)
	defer txn.Abort()

	// uniqueness: error if email already exists in db
	existingUser, err := c.ReadUser(user.Email)
	if err != nil {
		return fmt.Errorf("Error determining if user with email %s already exists in db: %s", user.Email, err)
	}

	if existingUser != nil {
		return fmt.Errorf("Cannot create user: user already exists with email %s", user.Email)
	}

	if err := txn.Insert("users", user); err != nil {
		return err
	}

	txn.Commit()

	return nil
}

func (c *Client) ReadUser(email string) (*User, error) {
	// Create read-only transaction
	txn := c.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("users", "id", email)
	if err != nil {
		return nil, err
	}

	if raw != nil {
		p := raw.(*User)
		return p, nil
	}

	return nil, nil
}

func (c *Client) UpdateUser(user *User) error {
	txn := c.db.Txn(true)
	defer txn.Abort()

	// email is unique (enforced on write)
	raw, err := txn.First("users", "id", user.Email)
	if err != nil {
		return err
	}

	p := raw.(*User)

	p.Name = user.Name
	p.Age = user.Age

	err = txn.Insert("users", p)
	if err != nil {
		return err
	}

	txn.Commit()

	return nil
}

func (c *Client) DeleteUser(user *User) error {
	txn := c.db.Txn(true)
	defer txn.Abort()

	existingUser, err := c.ReadUser(user.Email)
	if err != nil {
		return fmt.Errorf("Error determining if user with email %s exists in db: %s", user.Email, err)
	}

	if existingUser == nil {
		return fmt.Errorf("Cannot delete user with email %s: email not in db", user.Email)
	}

	err = txn.Delete("users", user)
	if err != nil {
		return err
	}

	txn.Commit()

	return nil
}

func (c *Client) ReadRegions() ([]*Region, error) {
	regions := []*Region{}

	txn := c.db.Txn(false)

	it, err := txn.Get("regions", "id")
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		r := obj.(*Region)
		regions = append(regions, r)
	}

	return regions, nil
}
