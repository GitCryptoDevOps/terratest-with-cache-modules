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

// ShippingAddress The representation of ShippingAddress
type ShippingAddress struct {
	Addressee *string `mandatory:"false" json:"addressee"`

	CareOf *string `mandatory:"false" json:"careOf"`

	Address1 *string `mandatory:"false" json:"address1"`

	Address2 *string `mandatory:"false" json:"address2"`

	Address3 *string `mandatory:"false" json:"address3"`

	Address4 *string `mandatory:"false" json:"address4"`

	CityOrLocality *string `mandatory:"false" json:"cityOrLocality"`

	StateOrRegion *string `mandatory:"false" json:"stateOrRegion"`

	Zipcode *string `mandatory:"false" json:"zipcode"`

	Country *string `mandatory:"false" json:"country"`

	PhoneNumber *string `mandatory:"false" json:"phoneNumber"`

	Email *string `mandatory:"false" json:"email"`
}

func (m ShippingAddress) String() string {
	return common.PointerString(m)
}
