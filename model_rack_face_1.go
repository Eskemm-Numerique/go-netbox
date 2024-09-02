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

// RackFace1 * `front` - Front * `rear` - Rear
type RackFace1 string

// List of Rack_face_1
const (
	RACKFACE1_FRONT RackFace1 = "front"
	RACKFACE1_REAR RackFace1 = "rear"
	RACKFACE1_EMPTY RackFace1 = ""
)

// All allowed values of RackFace1 enum
var AllowedRackFace1EnumValues = []RackFace1{
	"front",
	"rear",
	"",
}

func (v *RackFace1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackFace1(value)
	for _, existing := range AllowedRackFace1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackFace1", value)
}

// NewRackFace1FromValue returns a pointer to a valid RackFace1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackFace1FromValue(v string) (*RackFace1, error) {
	ev := RackFace1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackFace1: valid values are %v", v, AllowedRackFace1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackFace1) IsValid() bool {
	for _, existing := range AllowedRackFace1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Rack_face_1 value
func (v RackFace1) Ptr() *RackFace1 {
	return &v
}

type NullableRackFace1 struct {
	value *RackFace1
	isSet bool
}

func (v NullableRackFace1) Get() *RackFace1 {
	return v.value
}

func (v *NullableRackFace1) Set(val *RackFace1) {
	v.value = val
	v.isSet = true
}

func (v NullableRackFace1) IsSet() bool {
	return v.isSet
}

func (v *NullableRackFace1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackFace1(val *RackFace1) *NullableRackFace1 {
	return &NullableRackFace1{value: val, isSet: true}
}

func (v NullableRackFace1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackFace1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

