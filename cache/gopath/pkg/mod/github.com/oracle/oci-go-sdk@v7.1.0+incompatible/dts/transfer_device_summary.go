// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DTS API
//
// A description of the DTS API
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TransferDeviceSummary The representation of TransferDeviceSummary
type TransferDeviceSummary struct {
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	IscsiIQN *string `mandatory:"false" json:"iscsiIQN"`

	Label *string `mandatory:"false" json:"label"`

	LifecycleState TransferDeviceSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	AttachedTransferPackageLabel *string `mandatory:"false" json:"attachedTransferPackageLabel"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`

	UploadStatusLogUri *string `mandatory:"false" json:"uploadStatusLogUri"`
}

func (m TransferDeviceSummary) String() string {
	return common.PointerString(m)
}

// TransferDeviceSummaryLifecycleStateEnum Enum with underlying type: string
type TransferDeviceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for TransferDeviceSummaryLifecycleStateEnum
const (
	TransferDeviceSummaryLifecycleStatePreparing  TransferDeviceSummaryLifecycleStateEnum = "PREPARING"
	TransferDeviceSummaryLifecycleStateReady      TransferDeviceSummaryLifecycleStateEnum = "READY"
	TransferDeviceSummaryLifecycleStatePackaged   TransferDeviceSummaryLifecycleStateEnum = "PACKAGED"
	TransferDeviceSummaryLifecycleStateActive     TransferDeviceSummaryLifecycleStateEnum = "ACTIVE"
	TransferDeviceSummaryLifecycleStateProcessing TransferDeviceSummaryLifecycleStateEnum = "PROCESSING"
	TransferDeviceSummaryLifecycleStateComplete   TransferDeviceSummaryLifecycleStateEnum = "COMPLETE"
	TransferDeviceSummaryLifecycleStateMissing    TransferDeviceSummaryLifecycleStateEnum = "MISSING"
	TransferDeviceSummaryLifecycleStateError      TransferDeviceSummaryLifecycleStateEnum = "ERROR"
	TransferDeviceSummaryLifecycleStateDeleted    TransferDeviceSummaryLifecycleStateEnum = "DELETED"
	TransferDeviceSummaryLifecycleStateCancelled  TransferDeviceSummaryLifecycleStateEnum = "CANCELLED"
)

var mappingTransferDeviceSummaryLifecycleState = map[string]TransferDeviceSummaryLifecycleStateEnum{
	"PREPARING":  TransferDeviceSummaryLifecycleStatePreparing,
	"READY":      TransferDeviceSummaryLifecycleStateReady,
	"PACKAGED":   TransferDeviceSummaryLifecycleStatePackaged,
	"ACTIVE":     TransferDeviceSummaryLifecycleStateActive,
	"PROCESSING": TransferDeviceSummaryLifecycleStateProcessing,
	"COMPLETE":   TransferDeviceSummaryLifecycleStateComplete,
	"MISSING":    TransferDeviceSummaryLifecycleStateMissing,
	"ERROR":      TransferDeviceSummaryLifecycleStateError,
	"DELETED":    TransferDeviceSummaryLifecycleStateDeleted,
	"CANCELLED":  TransferDeviceSummaryLifecycleStateCancelled,
}

// GetTransferDeviceSummaryLifecycleStateEnumValues Enumerates the set of values for TransferDeviceSummaryLifecycleStateEnum
func GetTransferDeviceSummaryLifecycleStateEnumValues() []TransferDeviceSummaryLifecycleStateEnum {
	values := make([]TransferDeviceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingTransferDeviceSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
