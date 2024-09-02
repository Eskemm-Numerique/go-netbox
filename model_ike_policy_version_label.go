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

// IKEPolicyVersionLabel the model 'IKEPolicyVersionLabel'
type IKEPolicyVersionLabel string

// List of IKEPolicy_version_label
const (
	IKEPOLICYVERSIONLABEL_IKEV1 IKEPolicyVersionLabel = "IKEv1"
	IKEPOLICYVERSIONLABEL_IKEV2 IKEPolicyVersionLabel = "IKEv2"
)

// All allowed values of IKEPolicyVersionLabel enum
var AllowedIKEPolicyVersionLabelEnumValues = []IKEPolicyVersionLabel{
	"IKEv1",
	"IKEv2",
}

func (v *IKEPolicyVersionLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IKEPolicyVersionLabel(value)
	for _, existing := range AllowedIKEPolicyVersionLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IKEPolicyVersionLabel", value)
}

// NewIKEPolicyVersionLabelFromValue returns a pointer to a valid IKEPolicyVersionLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIKEPolicyVersionLabelFromValue(v string) (*IKEPolicyVersionLabel, error) {
	ev := IKEPolicyVersionLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IKEPolicyVersionLabel: valid values are %v", v, AllowedIKEPolicyVersionLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IKEPolicyVersionLabel) IsValid() bool {
	for _, existing := range AllowedIKEPolicyVersionLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IKEPolicy_version_label value
func (v IKEPolicyVersionLabel) Ptr() *IKEPolicyVersionLabel {
	return &v
}

type NullableIKEPolicyVersionLabel struct {
	value *IKEPolicyVersionLabel
	isSet bool
}

func (v NullableIKEPolicyVersionLabel) Get() *IKEPolicyVersionLabel {
	return v.value
}

func (v *NullableIKEPolicyVersionLabel) Set(val *IKEPolicyVersionLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEPolicyVersionLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEPolicyVersionLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEPolicyVersionLabel(val *IKEPolicyVersionLabel) *NullableIKEPolicyVersionLabel {
	return &NullableIKEPolicyVersionLabel{value: val, isSet: true}
}

func (v NullableIKEPolicyVersionLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEPolicyVersionLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

