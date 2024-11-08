// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package framework

// EphemeralResourceSpyClient is used in tests to verify that an ephemeral resource lifecycle handler has been executed.
type EphemeralResourceSpyClient struct {
	renewInvocations int
	closeInvocations int
}

// Renew will increment the number of invocations for this instance, which can be retrieved with the `RenewInvocations` method
func (e *EphemeralResourceSpyClient) Renew() {
	e.renewInvocations++
}

// RenewInvocations returns the number of times the `Renew` method has been called on this instance.
func (e *EphemeralResourceSpyClient) RenewInvocations() int {
	return e.renewInvocations
}

// Close will increment the number of invocations for this instance, which can be retrieved with the `CloseInvocations` method
func (e *EphemeralResourceSpyClient) Close() {
	e.closeInvocations++
}

// CloseInvocations returns the number of times the `Close` method has been called on this instance.
func (e *EphemeralResourceSpyClient) CloseInvocations() int {
	return e.closeInvocations
}
