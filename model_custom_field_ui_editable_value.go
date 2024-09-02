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

// CustomFieldUiEditableValue * `yes` - Yes * `no` - No * `hidden` - Hidden
type CustomFieldUiEditableValue string

// List of CustomField_ui_editable_value
const (
	CUSTOMFIELDUIEDITABLEVALUE_YES CustomFieldUiEditableValue = "yes"
	CUSTOMFIELDUIEDITABLEVALUE_NO CustomFieldUiEditableValue = "no"
	CUSTOMFIELDUIEDITABLEVALUE_HIDDEN CustomFieldUiEditableValue = "hidden"
)

// All allowed values of CustomFieldUiEditableValue enum
var AllowedCustomFieldUiEditableValueEnumValues = []CustomFieldUiEditableValue{
	"yes",
	"no",
	"hidden",
}

func (v *CustomFieldUiEditableValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomFieldUiEditableValue(value)
	for _, existing := range AllowedCustomFieldUiEditableValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomFieldUiEditableValue", value)
}

// NewCustomFieldUiEditableValueFromValue returns a pointer to a valid CustomFieldUiEditableValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomFieldUiEditableValueFromValue(v string) (*CustomFieldUiEditableValue, error) {
	ev := CustomFieldUiEditableValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomFieldUiEditableValue: valid values are %v", v, AllowedCustomFieldUiEditableValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomFieldUiEditableValue) IsValid() bool {
	for _, existing := range AllowedCustomFieldUiEditableValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomField_ui_editable_value value
func (v CustomFieldUiEditableValue) Ptr() *CustomFieldUiEditableValue {
	return &v
}

type NullableCustomFieldUiEditableValue struct {
	value *CustomFieldUiEditableValue
	isSet bool
}

func (v NullableCustomFieldUiEditableValue) Get() *CustomFieldUiEditableValue {
	return v.value
}

func (v *NullableCustomFieldUiEditableValue) Set(val *CustomFieldUiEditableValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldUiEditableValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldUiEditableValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldUiEditableValue(val *CustomFieldUiEditableValue) *NullableCustomFieldUiEditableValue {
	return &NullableCustomFieldUiEditableValue{value: val, isSet: true}
}

func (v NullableCustomFieldUiEditableValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldUiEditableValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

