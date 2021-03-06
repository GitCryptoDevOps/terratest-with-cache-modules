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

// ChangeTransferJobCompartmentDetails The representation of ChangeTransferJobCompartmentDetails
type ChangeTransferJobCompartmentDetails struct {

	// The OCID  (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment into which the resources should be moved.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeTransferJobCompartmentDetails) String() string {
	return common.PointerString(m)
}
