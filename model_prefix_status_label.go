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

// PrefixStatusLabel the model 'PrefixStatusLabel'
type PrefixStatusLabel string

// List of Prefix_status_label
const (
	PREFIXSTATUSLABEL_CONTAINER PrefixStatusLabel = "Container"
	PREFIXSTATUSLABEL_ACTIVE PrefixStatusLabel = "Active"
	PREFIXSTATUSLABEL_RESERVED PrefixStatusLabel = "Reserved"
	PREFIXSTATUSLABEL_DEPRECATED PrefixStatusLabel = "Deprecated"
)

// All allowed values of PrefixStatusLabel enum
var AllowedPrefixStatusLabelEnumValues = []PrefixStatusLabel{
	"Container",
	"Active",
	"Reserved",
	"Deprecated",
}

func (v *PrefixStatusLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrefixStatusLabel(value)
	for _, existing := range AllowedPrefixStatusLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrefixStatusLabel", value)
}

// NewPrefixStatusLabelFromValue returns a pointer to a valid PrefixStatusLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrefixStatusLabelFromValue(v string) (*PrefixStatusLabel, error) {
	ev := PrefixStatusLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrefixStatusLabel: valid values are %v", v, AllowedPrefixStatusLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrefixStatusLabel) IsValid() bool {
	for _, existing := range AllowedPrefixStatusLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Prefix_status_label value
func (v PrefixStatusLabel) Ptr() *PrefixStatusLabel {
	return &v
}

type NullablePrefixStatusLabel struct {
	value *PrefixStatusLabel
	isSet bool
}

func (v NullablePrefixStatusLabel) Get() *PrefixStatusLabel {
	return v.value
}

func (v *NullablePrefixStatusLabel) Set(val *PrefixStatusLabel) {
	v.value = val
	v.isSet = true
}

func (v NullablePrefixStatusLabel) IsSet() bool {
	return v.isSet
}

func (v *NullablePrefixStatusLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrefixStatusLabel(val *PrefixStatusLabel) *NullablePrefixStatusLabel {
	return &NullablePrefixStatusLabel{value: val, isSet: true}
}

func (v NullablePrefixStatusLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrefixStatusLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

