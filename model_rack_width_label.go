/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.10 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// RackWidthLabel the model 'RackWidthLabel'
type RackWidthLabel string

// List of Rack_width_label
const (
	RACKWIDTHLABEL__10_INCHES RackWidthLabel = "10 inches"
	RACKWIDTHLABEL__19_INCHES RackWidthLabel = "19 inches"
	RACKWIDTHLABEL__21_INCHES RackWidthLabel = "21 inches"
	RACKWIDTHLABEL__23_INCHES RackWidthLabel = "23 inches"
)

// All allowed values of RackWidthLabel enum
var AllowedRackWidthLabelEnumValues = []RackWidthLabel{
	"10 inches",
	"19 inches",
	"21 inches",
	"23 inches",
}

func (v *RackWidthLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackWidthLabel(value)
	for _, existing := range AllowedRackWidthLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackWidthLabel", value)
}

// NewRackWidthLabelFromValue returns a pointer to a valid RackWidthLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackWidthLabelFromValue(v string) (*RackWidthLabel, error) {
	ev := RackWidthLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackWidthLabel: valid values are %v", v, AllowedRackWidthLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackWidthLabel) IsValid() bool {
	for _, existing := range AllowedRackWidthLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Rack_width_label value
func (v RackWidthLabel) Ptr() *RackWidthLabel {
	return &v
}

type NullableRackWidthLabel struct {
	value *RackWidthLabel
	isSet bool
}

func (v NullableRackWidthLabel) Get() *RackWidthLabel {
	return v.value
}

func (v *NullableRackWidthLabel) Set(val *RackWidthLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableRackWidthLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableRackWidthLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackWidthLabel(val *RackWidthLabel) *NullableRackWidthLabel {
	return &NullableRackWidthLabel{value: val, isSet: true}
}

func (v NullableRackWidthLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackWidthLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

