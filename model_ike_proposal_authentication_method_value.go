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

// IKEProposalAuthenticationMethodValue * `preshared-keys` - Pre-shared keys * `certificates` - Certificates * `rsa-signatures` - RSA signatures * `dsa-signatures` - DSA signatures
type IKEProposalAuthenticationMethodValue string

// List of IKEProposal_authentication_method_value
const (
	IKEPROPOSALAUTHENTICATIONMETHODVALUE_PRESHARED_KEYS IKEProposalAuthenticationMethodValue = "preshared-keys"
	IKEPROPOSALAUTHENTICATIONMETHODVALUE_CERTIFICATES IKEProposalAuthenticationMethodValue = "certificates"
	IKEPROPOSALAUTHENTICATIONMETHODVALUE_RSA_SIGNATURES IKEProposalAuthenticationMethodValue = "rsa-signatures"
	IKEPROPOSALAUTHENTICATIONMETHODVALUE_DSA_SIGNATURES IKEProposalAuthenticationMethodValue = "dsa-signatures"
)

// All allowed values of IKEProposalAuthenticationMethodValue enum
var AllowedIKEProposalAuthenticationMethodValueEnumValues = []IKEProposalAuthenticationMethodValue{
	"preshared-keys",
	"certificates",
	"rsa-signatures",
	"dsa-signatures",
}

func (v *IKEProposalAuthenticationMethodValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IKEProposalAuthenticationMethodValue(value)
	for _, existing := range AllowedIKEProposalAuthenticationMethodValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IKEProposalAuthenticationMethodValue", value)
}

// NewIKEProposalAuthenticationMethodValueFromValue returns a pointer to a valid IKEProposalAuthenticationMethodValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIKEProposalAuthenticationMethodValueFromValue(v string) (*IKEProposalAuthenticationMethodValue, error) {
	ev := IKEProposalAuthenticationMethodValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IKEProposalAuthenticationMethodValue: valid values are %v", v, AllowedIKEProposalAuthenticationMethodValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IKEProposalAuthenticationMethodValue) IsValid() bool {
	for _, existing := range AllowedIKEProposalAuthenticationMethodValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IKEProposal_authentication_method_value value
func (v IKEProposalAuthenticationMethodValue) Ptr() *IKEProposalAuthenticationMethodValue {
	return &v
}

type NullableIKEProposalAuthenticationMethodValue struct {
	value *IKEProposalAuthenticationMethodValue
	isSet bool
}

func (v NullableIKEProposalAuthenticationMethodValue) Get() *IKEProposalAuthenticationMethodValue {
	return v.value
}

func (v *NullableIKEProposalAuthenticationMethodValue) Set(val *IKEProposalAuthenticationMethodValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEProposalAuthenticationMethodValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEProposalAuthenticationMethodValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEProposalAuthenticationMethodValue(val *IKEProposalAuthenticationMethodValue) *NullableIKEProposalAuthenticationMethodValue {
	return &NullableIKEProposalAuthenticationMethodValue{value: val, isSet: true}
}

func (v NullableIKEProposalAuthenticationMethodValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEProposalAuthenticationMethodValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

