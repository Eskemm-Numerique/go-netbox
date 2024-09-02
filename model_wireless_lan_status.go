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

// checks if the WirelessLANStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WirelessLANStatus{}

// WirelessLANStatus struct for WirelessLANStatus
type WirelessLANStatus struct {
	// * `active` - Active * `reserved` - Reserved * `disabled` - Disabled * `deprecated` - Deprecated
	Value *string `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WirelessLANStatus WirelessLANStatus

// NewWirelessLANStatus instantiates a new WirelessLANStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWirelessLANStatus() *WirelessLANStatus {
	this := WirelessLANStatus{}
	return &this
}

// NewWirelessLANStatusWithDefaults instantiates a new WirelessLANStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWirelessLANStatusWithDefaults() *WirelessLANStatus {
	this := WirelessLANStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WirelessLANStatus) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLANStatus) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WirelessLANStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WirelessLANStatus) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WirelessLANStatus) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLANStatus) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WirelessLANStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WirelessLANStatus) SetLabel(v string) {
	o.Label = &v
}

func (o WirelessLANStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WirelessLANStatus) ToMap() (map[string]interface{}, error) {
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

func (o *WirelessLANStatus) UnmarshalJSON(data []byte) (err error) {
	varWirelessLANStatus := _WirelessLANStatus{}

	err = json.Unmarshal(data, &varWirelessLANStatus)

	if err != nil {
		return err
	}

	*o = WirelessLANStatus(varWirelessLANStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWirelessLANStatus struct {
	value *WirelessLANStatus
	isSet bool
}

func (v NullableWirelessLANStatus) Get() *WirelessLANStatus {
	return v.value
}

func (v *NullableWirelessLANStatus) Set(val *WirelessLANStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessLANStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessLANStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessLANStatus(val *WirelessLANStatus) *NullableWirelessLANStatus {
	return &NullableWirelessLANStatus{value: val, isSet: true}
}

func (v NullableWirelessLANStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessLANStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


