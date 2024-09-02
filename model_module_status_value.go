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

// ModuleStatusValue * `offline` - Offline * `active` - Active * `planned` - Planned * `staged` - Staged * `failed` - Failed * `decommissioning` - Decommissioning
type ModuleStatusValue string

// List of Module_status_value
const (
	MODULESTATUSVALUE_OFFLINE ModuleStatusValue = "offline"
	MODULESTATUSVALUE_ACTIVE ModuleStatusValue = "active"
	MODULESTATUSVALUE_PLANNED ModuleStatusValue = "planned"
	MODULESTATUSVALUE_STAGED ModuleStatusValue = "staged"
	MODULESTATUSVALUE_FAILED ModuleStatusValue = "failed"
	MODULESTATUSVALUE_DECOMMISSIONING ModuleStatusValue = "decommissioning"
)

// All allowed values of ModuleStatusValue enum
var AllowedModuleStatusValueEnumValues = []ModuleStatusValue{
	"offline",
	"active",
	"planned",
	"staged",
	"failed",
	"decommissioning",
}

func (v *ModuleStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModuleStatusValue(value)
	for _, existing := range AllowedModuleStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModuleStatusValue", value)
}

// NewModuleStatusValueFromValue returns a pointer to a valid ModuleStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModuleStatusValueFromValue(v string) (*ModuleStatusValue, error) {
	ev := ModuleStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModuleStatusValue: valid values are %v", v, AllowedModuleStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModuleStatusValue) IsValid() bool {
	for _, existing := range AllowedModuleStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Module_status_value value
func (v ModuleStatusValue) Ptr() *ModuleStatusValue {
	return &v
}

type NullableModuleStatusValue struct {
	value *ModuleStatusValue
	isSet bool
}

func (v NullableModuleStatusValue) Get() *ModuleStatusValue {
	return v.value
}

func (v *NullableModuleStatusValue) Set(val *ModuleStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleStatusValue(val *ModuleStatusValue) *NullableModuleStatusValue {
	return &NullableModuleStatusValue{value: val, isSet: true}
}

func (v NullableModuleStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

