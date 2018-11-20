// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateInstancePoolPlacementConfigurationDetails The location for where an instance pool will place instances.
type UpdateInstancePoolPlacementConfigurationDetails struct {

	// The availability domain to place instances.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the primary subnet to place instances.
	PrimarySubnetId *string `mandatory:"true" json:"primarySubnetId"`

	// The set of subnet OCIDs for secondary VNICs for instances in the pool.
	SecondaryVnicSubnets []InstancePoolPlacementSecondaryVnicSubnet `mandatory:"false" json:"secondaryVnicSubnets"`
}

func (m UpdateInstancePoolPlacementConfigurationDetails) String() string {
	return common.PointerString(m)
}
