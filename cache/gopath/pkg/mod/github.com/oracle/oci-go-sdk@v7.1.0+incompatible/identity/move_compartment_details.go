// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// MoveCompartmentDetails The representation of MoveCompartmentDetails
type MoveCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the parent compartment
	// into which the compartment should be moved.
	TargetCompartmentId *string `mandatory:"true" json:"targetCompartmentId"`
}

func (m MoveCompartmentDetails) String() string {
	return common.PointerString(m)
}
