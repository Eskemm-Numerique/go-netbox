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

// JournalEntryKindValue * `info` - Info * `success` - Success * `warning` - Warning * `danger` - Danger
type JournalEntryKindValue string

// List of JournalEntry_kind_value
const (
	JOURNALENTRYKINDVALUE_INFO JournalEntryKindValue = "info"
	JOURNALENTRYKINDVALUE_SUCCESS JournalEntryKindValue = "success"
	JOURNALENTRYKINDVALUE_WARNING JournalEntryKindValue = "warning"
	JOURNALENTRYKINDVALUE_DANGER JournalEntryKindValue = "danger"
)

// All allowed values of JournalEntryKindValue enum
var AllowedJournalEntryKindValueEnumValues = []JournalEntryKindValue{
	"info",
	"success",
	"warning",
	"danger",
}

func (v *JournalEntryKindValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JournalEntryKindValue(value)
	for _, existing := range AllowedJournalEntryKindValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JournalEntryKindValue", value)
}

// NewJournalEntryKindValueFromValue returns a pointer to a valid JournalEntryKindValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJournalEntryKindValueFromValue(v string) (*JournalEntryKindValue, error) {
	ev := JournalEntryKindValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JournalEntryKindValue: valid values are %v", v, AllowedJournalEntryKindValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JournalEntryKindValue) IsValid() bool {
	for _, existing := range AllowedJournalEntryKindValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JournalEntry_kind_value value
func (v JournalEntryKindValue) Ptr() *JournalEntryKindValue {
	return &v
}

type NullableJournalEntryKindValue struct {
	value *JournalEntryKindValue
	isSet bool
}

func (v NullableJournalEntryKindValue) Get() *JournalEntryKindValue {
	return v.value
}

func (v *NullableJournalEntryKindValue) Set(val *JournalEntryKindValue) {
	v.value = val
	v.isSet = true
}

func (v NullableJournalEntryKindValue) IsSet() bool {
	return v.isSet
}

func (v *NullableJournalEntryKindValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJournalEntryKindValue(val *JournalEntryKindValue) *NullableJournalEntryKindValue {
	return &NullableJournalEntryKindValue{value: val, isSet: true}
}

func (v NullableJournalEntryKindValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJournalEntryKindValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

