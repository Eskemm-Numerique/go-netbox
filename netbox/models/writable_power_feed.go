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

// WritablePowerFeed writable power feed
// swagger:model WritablePowerFeed
type WritablePowerFeed struct {

	// Amperage
	// Maximum: 32767
	// Minimum: 1
	Amperage int64 `json:"amperage,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Max utilization
	//
	// Maximum permissible draw (percentage)
	// Maximum: 100
	// Minimum: 1
	MaxUtilization int64 `json:"max_utilization,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// Phase
	// Enum: [single-phase three-phase]
	Phase string `json:"phase,omitempty"`

	// Power panel
	// Required: true
	PowerPanel *int64 `json:"power_panel"`

	// Rack
	Rack *int64 `json:"rack,omitempty"`

	// Status
	// Enum: [offline active planned failed]
	Status string `json:"status,omitempty"`

	// Supply
	// Enum: [ac dc]
	Supply string `json:"supply,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// Type
	// Enum: [primary redundant]
	Type string `json:"type,omitempty"`

	// Voltage
	// Maximum: 32767
	// Minimum: 1
	Voltage int64 `json:"voltage,omitempty"`
}

// Validate validates this writable power feed
func (m *WritablePowerFeed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmperage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxUtilization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerPanel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupply(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoltage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritablePowerFeed) validateAmperage(formats strfmt.Registry) error {

	if swag.IsZero(m.Amperage) { // not required
		return nil
	}

	if err := validate.MinimumInt("amperage", "body", int64(m.Amperage), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("amperage", "body", int64(m.Amperage), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateMaxUtilization(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxUtilization) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_utilization", "body", int64(m.MaxUtilization), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_utilization", "body", int64(m.MaxUtilization), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

var writablePowerFeedTypePhasePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["single-phase","three-phase"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypePhasePropEnum = append(writablePowerFeedTypePhasePropEnum, v)
	}
}

const (

	// WritablePowerFeedPhaseSinglePhase captures enum value "single-phase"
	WritablePowerFeedPhaseSinglePhase string = "single-phase"

	// WritablePowerFeedPhaseThreePhase captures enum value "three-phase"
	WritablePowerFeedPhaseThreePhase string = "three-phase"
)

// prop value enum
func (m *WritablePowerFeed) validatePhaseEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writablePowerFeedTypePhasePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validatePhase(formats strfmt.Registry) error {

	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	// value enum
	if err := m.validatePhaseEnum("phase", "body", m.Phase); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validatePowerPanel(formats strfmt.Registry) error {

	if err := validate.Required("power_panel", "body", m.PowerPanel); err != nil {
		return err
	}

	return nil
}

var writablePowerFeedTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["offline","active","planned","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypeStatusPropEnum = append(writablePowerFeedTypeStatusPropEnum, v)
	}
}

const (

	// WritablePowerFeedStatusOffline captures enum value "offline"
	WritablePowerFeedStatusOffline string = "offline"

	// WritablePowerFeedStatusActive captures enum value "active"
	WritablePowerFeedStatusActive string = "active"

	// WritablePowerFeedStatusPlanned captures enum value "planned"
	WritablePowerFeedStatusPlanned string = "planned"

	// WritablePowerFeedStatusFailed captures enum value "failed"
	WritablePowerFeedStatusFailed string = "failed"
)

// prop value enum
func (m *WritablePowerFeed) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writablePowerFeedTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var writablePowerFeedTypeSupplyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ac","dc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypeSupplyPropEnum = append(writablePowerFeedTypeSupplyPropEnum, v)
	}
}

const (

	// WritablePowerFeedSupplyAc captures enum value "ac"
	WritablePowerFeedSupplyAc string = "ac"

	// WritablePowerFeedSupplyDc captures enum value "dc"
	WritablePowerFeedSupplyDc string = "dc"
)

// prop value enum
func (m *WritablePowerFeed) validateSupplyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writablePowerFeedTypeSupplyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validateSupply(formats strfmt.Registry) error {

	if swag.IsZero(m.Supply) { // not required
		return nil
	}

	// value enum
	if err := m.validateSupplyEnum("supply", "body", m.Supply); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateTags(formats strfmt.Registry) error {

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

var writablePowerFeedTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","redundant"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypeTypePropEnum = append(writablePowerFeedTypeTypePropEnum, v)
	}
}

const (

	// WritablePowerFeedTypePrimary captures enum value "primary"
	WritablePowerFeedTypePrimary string = "primary"

	// WritablePowerFeedTypeRedundant captures enum value "redundant"
	WritablePowerFeedTypeRedundant string = "redundant"
)

// prop value enum
func (m *WritablePowerFeed) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writablePowerFeedTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateVoltage(formats strfmt.Registry) error {

	if swag.IsZero(m.Voltage) { // not required
		return nil
	}

	if err := validate.MinimumInt("voltage", "body", int64(m.Voltage), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("voltage", "body", int64(m.Voltage), 32767, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePowerFeed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePowerFeed) UnmarshalBinary(b []byte) error {
	var res WritablePowerFeed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
