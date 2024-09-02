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

// WirelessLANStatusValue * `active` - Active * `reserved` - Reserved * `disabled` - Disabled * `deprecated` - Deprecated
type WirelessLANStatusValue string

// List of WirelessLAN_status_value
const (
	WIRELESSLANSTATUSVALUE_ACTIVE WirelessLANStatusValue = "active"
	WIRELESSLANSTATUSVALUE_RESERVED WirelessLANStatusValue = "reserved"
	WIRELESSLANSTATUSVALUE_DISABLED WirelessLANStatusValue = "disabled"
	WIRELESSLANSTATUSVALUE_DEPRECATED WirelessLANStatusValue = "deprecated"
	WIRELESSLANSTATUSVALUE_EMPTY WirelessLANStatusValue = ""
)

// All allowed values of WirelessLANStatusValue enum
var AllowedWirelessLANStatusValueEnumValues = []WirelessLANStatusValue{
	"active",
	"reserved",
	"disabled",
	"deprecated",
	"",
}

func (v *WirelessLANStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WirelessLANStatusValue(value)
	for _, existing := range AllowedWirelessLANStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WirelessLANStatusValue", value)
}

// NewWirelessLANStatusValueFromValue returns a pointer to a valid WirelessLANStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWirelessLANStatusValueFromValue(v string) (*WirelessLANStatusValue, error) {
	ev := WirelessLANStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WirelessLANStatusValue: valid values are %v", v, AllowedWirelessLANStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WirelessLANStatusValue) IsValid() bool {
	for _, existing := range AllowedWirelessLANStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WirelessLAN_status_value value
func (v WirelessLANStatusValue) Ptr() *WirelessLANStatusValue {
	return &v
}

type NullableWirelessLANStatusValue struct {
	value *WirelessLANStatusValue
	isSet bool
}

func (v NullableWirelessLANStatusValue) Get() *WirelessLANStatusValue {
	return v.value
}

func (v *NullableWirelessLANStatusValue) Set(val *WirelessLANStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessLANStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessLANStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessLANStatusValue(val *WirelessLANStatusValue) *NullableWirelessLANStatusValue {
	return &NullableWirelessLANStatusValue{value: val, isSet: true}
}

func (v NullableWirelessLANStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessLANStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

