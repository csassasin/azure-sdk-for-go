//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azdatalake

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/lease"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/sas"
)

// SharedKeyCredential contains an account's name and its primary or secondary key.
type SharedKeyCredential = exported.SharedKeyCredential

// NewSharedKeyCredential creates an immutable SharedKeyCredential containing the
// storage account's name and either its primary or secondary key.
func NewSharedKeyCredential(accountName, accountKey string) (*SharedKeyCredential, error) {
	return exported.NewSharedKeyCredential(accountName, accountKey)
}

// URLParts object represents the components that make up an Azure Storage Container/Blob URL.
// NOTE: Changing any SAS-related field requires computing a new SAS signature.
type URLParts = sas.URLParts

// ParseURL parses a URL initializing URLParts' fields including any SAS-related & snapshot query parameters. Any other
// query parameters remain in the UnparsedParams field. This method overwrites all fields in the URLParts object.
func ParseURL(u string) (URLParts, error) {
	return sas.ParseURL(u)
}

// HTTPRange defines a range of bytes within an HTTP resource, starting at offset and
// ending at offset+count. A zero-value HTTPRange indicates the entire resource. An HTTPRange
// which has an offset but no zero value count indicates from the offset to the resource's end.
type HTTPRange = exported.HTTPRange

// ===================================== LEASE CONSTANTS ============================================================

// StatusType defines values for StatusType
type StatusType = lease.StatusType

const (
	StatusTypeLocked   StatusType = lease.StatusTypeLocked
	StatusTypeUnlocked StatusType = lease.StatusTypeUnlocked
)

// PossibleStatusTypeValues returns the possible values for the StatusType const type.
func PossibleStatusTypeValues() []StatusType {
	return lease.PossibleStatusTypeValues()
}

// DurationType defines values for DurationType
type DurationType = lease.DurationType

const (
	DurationTypeInfinite DurationType = lease.DurationTypeInfinite
	DurationTypeFixed    DurationType = lease.DurationTypeFixed
)

// PossibleDurationTypeValues returns the possible values for the DurationType const type.
func PossibleDurationTypeValues() []DurationType {
	return lease.PossibleDurationTypeValues()
}

// StateType defines values for StateType
type StateType = lease.StateType

const (
	StateTypeAvailable StateType = lease.StateTypeAvailable
	StateTypeLeased    StateType = lease.StateTypeLeased
	StateTypeExpired   StateType = lease.StateTypeExpired
	StateTypeBreaking  StateType = lease.StateTypeBreaking
	StateTypeBroken    StateType = lease.StateTypeBroken
)

// PossibleStateTypeValues returns the possible values for the StateType const type.
func PossibleStateTypeValues() []StateType {
	return lease.PossibleStateTypeValues()
}
