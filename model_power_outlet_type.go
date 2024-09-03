/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.10 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PowerOutletType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerOutletType{}

// PowerOutletType struct for PowerOutletType
type PowerOutletType struct {
	Value *PatchedWritablePowerOutletTemplateRequestType `json:"value,omitempty"`
	Label *PowerOutletTypeLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PowerOutletType PowerOutletType

// NewPowerOutletType instantiates a new PowerOutletType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerOutletType() *PowerOutletType {
	this := PowerOutletType{}
	return &this
}

// NewPowerOutletTypeWithDefaults instantiates a new PowerOutletType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerOutletTypeWithDefaults() *PowerOutletType {
	this := PowerOutletType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PowerOutletType) GetValue() PatchedWritablePowerOutletTemplateRequestType {
	if o == nil || IsNil(o.Value) {
		var ret PatchedWritablePowerOutletTemplateRequestType
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletType) GetValueOk() (*PatchedWritablePowerOutletTemplateRequestType, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PowerOutletType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given PatchedWritablePowerOutletTemplateRequestType and assigns it to the Value field.
func (o *PowerOutletType) SetValue(v PatchedWritablePowerOutletTemplateRequestType) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PowerOutletType) GetLabel() PowerOutletTypeLabel {
	if o == nil || IsNil(o.Label) {
		var ret PowerOutletTypeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletType) GetLabelOk() (*PowerOutletTypeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PowerOutletType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given PowerOutletTypeLabel and assigns it to the Label field.
func (o *PowerOutletType) SetLabel(v PowerOutletTypeLabel) {
	o.Label = &v
}

func (o PowerOutletType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerOutletType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PowerOutletType) UnmarshalJSON(data []byte) (err error) {
	varPowerOutletType := _PowerOutletType{}

	err = json.Unmarshal(data, &varPowerOutletType)

	if err != nil {
		return err
	}

	*o = PowerOutletType(varPowerOutletType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerOutletType struct {
	value *PowerOutletType
	isSet bool
}

func (v NullablePowerOutletType) Get() *PowerOutletType {
	return v.value
}

func (v *NullablePowerOutletType) Set(val *PowerOutletType) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerOutletType) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerOutletType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerOutletType(val *PowerOutletType) *NullablePowerOutletType {
	return &NullablePowerOutletType{value: val, isSet: true}
}

func (v NullablePowerOutletType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerOutletType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


