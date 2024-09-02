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

// CableRequestLengthUnit * `km` - Kilometers * `m` - Meters * `cm` - Centimeters * `mi` - Miles * `ft` - Feet * `in` - Inches
type CableRequestLengthUnit string

// List of CableRequest_length_unit
const (
	CABLEREQUESTLENGTHUNIT_KM CableRequestLengthUnit = "km"
	CABLEREQUESTLENGTHUNIT_M CableRequestLengthUnit = "m"
	CABLEREQUESTLENGTHUNIT_CM CableRequestLengthUnit = "cm"
	CABLEREQUESTLENGTHUNIT_MI CableRequestLengthUnit = "mi"
	CABLEREQUESTLENGTHUNIT_FT CableRequestLengthUnit = "ft"
	CABLEREQUESTLENGTHUNIT_IN CableRequestLengthUnit = "in"
	CABLEREQUESTLENGTHUNIT_EMPTY CableRequestLengthUnit = ""
)

// All allowed values of CableRequestLengthUnit enum
var AllowedCableRequestLengthUnitEnumValues = []CableRequestLengthUnit{
	"km",
	"m",
	"cm",
	"mi",
	"ft",
	"in",
	"",
}

func (v *CableRequestLengthUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CableRequestLengthUnit(value)
	for _, existing := range AllowedCableRequestLengthUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CableRequestLengthUnit", value)
}

// NewCableRequestLengthUnitFromValue returns a pointer to a valid CableRequestLengthUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCableRequestLengthUnitFromValue(v string) (*CableRequestLengthUnit, error) {
	ev := CableRequestLengthUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CableRequestLengthUnit: valid values are %v", v, AllowedCableRequestLengthUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CableRequestLengthUnit) IsValid() bool {
	for _, existing := range AllowedCableRequestLengthUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CableRequest_length_unit value
func (v CableRequestLengthUnit) Ptr() *CableRequestLengthUnit {
	return &v
}

type NullableCableRequestLengthUnit struct {
	value *CableRequestLengthUnit
	isSet bool
}

func (v NullableCableRequestLengthUnit) Get() *CableRequestLengthUnit {
	return v.value
}

func (v *NullableCableRequestLengthUnit) Set(val *CableRequestLengthUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableCableRequestLengthUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableCableRequestLengthUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableRequestLengthUnit(val *CableRequestLengthUnit) *NullableCableRequestLengthUnit {
	return &NullableCableRequestLengthUnit{value: val, isSet: true}
}

func (v NullableCableRequestLengthUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableRequestLengthUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

