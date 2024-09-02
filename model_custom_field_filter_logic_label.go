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

// CustomFieldFilterLogicLabel the model 'CustomFieldFilterLogicLabel'
type CustomFieldFilterLogicLabel string

// List of CustomField_filter_logic_label
const (
	CUSTOMFIELDFILTERLOGICLABEL_DISABLED CustomFieldFilterLogicLabel = "Disabled"
	CUSTOMFIELDFILTERLOGICLABEL_LOOSE CustomFieldFilterLogicLabel = "Loose"
	CUSTOMFIELDFILTERLOGICLABEL_EXACT CustomFieldFilterLogicLabel = "Exact"
)

// All allowed values of CustomFieldFilterLogicLabel enum
var AllowedCustomFieldFilterLogicLabelEnumValues = []CustomFieldFilterLogicLabel{
	"Disabled",
	"Loose",
	"Exact",
}

func (v *CustomFieldFilterLogicLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomFieldFilterLogicLabel(value)
	for _, existing := range AllowedCustomFieldFilterLogicLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomFieldFilterLogicLabel", value)
}

// NewCustomFieldFilterLogicLabelFromValue returns a pointer to a valid CustomFieldFilterLogicLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomFieldFilterLogicLabelFromValue(v string) (*CustomFieldFilterLogicLabel, error) {
	ev := CustomFieldFilterLogicLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomFieldFilterLogicLabel: valid values are %v", v, AllowedCustomFieldFilterLogicLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomFieldFilterLogicLabel) IsValid() bool {
	for _, existing := range AllowedCustomFieldFilterLogicLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomField_filter_logic_label value
func (v CustomFieldFilterLogicLabel) Ptr() *CustomFieldFilterLogicLabel {
	return &v
}

type NullableCustomFieldFilterLogicLabel struct {
	value *CustomFieldFilterLogicLabel
	isSet bool
}

func (v NullableCustomFieldFilterLogicLabel) Get() *CustomFieldFilterLogicLabel {
	return v.value
}

func (v *NullableCustomFieldFilterLogicLabel) Set(val *CustomFieldFilterLogicLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldFilterLogicLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldFilterLogicLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldFilterLogicLabel(val *CustomFieldFilterLogicLabel) *NullableCustomFieldFilterLogicLabel {
	return &NullableCustomFieldFilterLogicLabel{value: val, isSet: true}
}

func (v NullableCustomFieldFilterLogicLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldFilterLogicLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

