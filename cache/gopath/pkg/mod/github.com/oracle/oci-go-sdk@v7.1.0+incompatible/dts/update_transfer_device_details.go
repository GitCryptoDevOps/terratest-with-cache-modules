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

// UpdateTransferDeviceDetails The representation of UpdateTransferDeviceDetails
type UpdateTransferDeviceDetails struct {
	LifecycleState UpdateTransferDeviceDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m UpdateTransferDeviceDetails) String() string {
	return common.PointerString(m)
}

// UpdateTransferDeviceDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateTransferDeviceDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateTransferDeviceDetailsLifecycleStateEnum
const (
	UpdateTransferDeviceDetailsLifecycleStatePreparing UpdateTransferDeviceDetailsLifecycleStateEnum = "PREPARING"
	UpdateTransferDeviceDetailsLifecycleStateReady     UpdateTransferDeviceDetailsLifecycleStateEnum = "READY"
	UpdateTransferDeviceDetailsLifecycleStateCancelled UpdateTransferDeviceDetailsLifecycleStateEnum = "CANCELLED"
)

var mappingUpdateTransferDeviceDetailsLifecycleState = map[string]UpdateTransferDeviceDetailsLifecycleStateEnum{
	"PREPARING": UpdateTransferDeviceDetailsLifecycleStatePreparing,
	"READY":     UpdateTransferDeviceDetailsLifecycleStateReady,
	"CANCELLED": UpdateTransferDeviceDetailsLifecycleStateCancelled,
}

// GetUpdateTransferDeviceDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateTransferDeviceDetailsLifecycleStateEnum
func GetUpdateTransferDeviceDetailsLifecycleStateEnumValues() []UpdateTransferDeviceDetailsLifecycleStateEnum {
	values := make([]UpdateTransferDeviceDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateTransferDeviceDetailsLifecycleState {
		values = append(values, v)
	}
	return values
}
