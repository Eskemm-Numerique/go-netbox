// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableIPAddress writable IP address
// swagger:model WritableIPAddress
type WritableIPAddress struct {

	// Address
	//
	// IPv4 or IPv6 address (with mask)
	// Required: true
	Address *string `json:"address"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// DNS Name
	//
	// Hostname or FQDN (not case-sensitive)
	// Max Length: 255
	// Pattern: ^[0-9A-Za-z._-]+$
	DNSName string `json:"dns_name,omitempty"`

	// Family
	// Read Only: true
	Family int64 `json:"family,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Interface
	Interface *int64 `json:"interface,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// NAT (Inside)
	//
	// The IP for which this address is the "outside" IP
	NatInside *int64 `json:"nat_inside,omitempty"`

	// Nat outside
	// Required: true
	NatOutside *int64 `json:"nat_outside"`

	// Role
	//
	// The functional role of this IP
	// Enum: [loopback secondary anycast vip vrrp hsrp glbp carp]
	Role string `json:"role,omitempty"`

	// Status
	//
	// The operational status of this IP
	// Enum: [active reserved deprecated dhcp]
	Status string `json:"status,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// VRF
	Vrf *int64 `json:"vrf,omitempty"`
}

// Validate validates this writable IP address
func (m *WritableIPAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNatOutside(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableIPAddress) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateDNSName(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSName) { // not required
		return nil
	}

	if err := validate.MaxLength("dns_name", "body", string(m.DNSName), 255); err != nil {
		return err
	}

	if err := validate.Pattern("dns_name", "body", string(m.DNSName), `^[0-9A-Za-z._-]+$`); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateNatOutside(formats strfmt.Registry) error {

	if err := validate.Required("nat_outside", "body", m.NatOutside); err != nil {
		return err
	}

	return nil
}

var writableIpAddressTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["loopback","secondary","anycast","vip","vrrp","hsrp","glbp","carp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableIpAddressTypeRolePropEnum = append(writableIpAddressTypeRolePropEnum, v)
	}
}

const (

	// WritableIPAddressRoleLoopback captures enum value "loopback"
	WritableIPAddressRoleLoopback string = "loopback"

	// WritableIPAddressRoleSecondary captures enum value "secondary"
	WritableIPAddressRoleSecondary string = "secondary"

	// WritableIPAddressRoleAnycast captures enum value "anycast"
	WritableIPAddressRoleAnycast string = "anycast"

	// WritableIPAddressRoleVip captures enum value "vip"
	WritableIPAddressRoleVip string = "vip"

	// WritableIPAddressRoleVrrp captures enum value "vrrp"
	WritableIPAddressRoleVrrp string = "vrrp"

	// WritableIPAddressRoleHsrp captures enum value "hsrp"
	WritableIPAddressRoleHsrp string = "hsrp"

	// WritableIPAddressRoleGlbp captures enum value "glbp"
	WritableIPAddressRoleGlbp string = "glbp"

	// WritableIPAddressRoleCarp captures enum value "carp"
	WritableIPAddressRoleCarp string = "carp"
)

// prop value enum
func (m *WritableIPAddress) validateRoleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableIpAddressTypeRolePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableIPAddress) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

var writableIpAddressTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","reserved","deprecated","dhcp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableIpAddressTypeStatusPropEnum = append(writableIpAddressTypeStatusPropEnum, v)
	}
}

const (

	// WritableIPAddressStatusActive captures enum value "active"
	WritableIPAddressStatusActive string = "active"

	// WritableIPAddressStatusReserved captures enum value "reserved"
	WritableIPAddressStatusReserved string = "reserved"

	// WritableIPAddressStatusDeprecated captures enum value "deprecated"
	WritableIPAddressStatusDeprecated string = "deprecated"

	// WritableIPAddressStatusDhcp captures enum value "dhcp"
	WritableIPAddressStatusDhcp string = "dhcp"
)

// prop value enum
func (m *WritableIPAddress) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableIpAddressTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableIPAddress) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableIPAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableIPAddress) UnmarshalBinary(b []byte) error {
	var res WritableIPAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
