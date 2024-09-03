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

// checks if the DeviceTypeWeightUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceTypeWeightUnit{}

// DeviceTypeWeightUnit struct for DeviceTypeWeightUnit
type DeviceTypeWeightUnit struct {
	Value *DeviceTypeWeightUnitValue `json:"value,omitempty"`
	Label *DeviceTypeWeightUnitLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceTypeWeightUnit DeviceTypeWeightUnit

// NewDeviceTypeWeightUnit instantiates a new DeviceTypeWeightUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceTypeWeightUnit() *DeviceTypeWeightUnit {
	this := DeviceTypeWeightUnit{}
	return &this
}

// NewDeviceTypeWeightUnitWithDefaults instantiates a new DeviceTypeWeightUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTypeWeightUnitWithDefaults() *DeviceTypeWeightUnit {
	this := DeviceTypeWeightUnit{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DeviceTypeWeightUnit) GetValue() DeviceTypeWeightUnitValue {
	if o == nil || IsNil(o.Value) {
		var ret DeviceTypeWeightUnitValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeWeightUnit) GetValueOk() (*DeviceTypeWeightUnitValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DeviceTypeWeightUnit) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given DeviceTypeWeightUnitValue and assigns it to the Value field.
func (o *DeviceTypeWeightUnit) SetValue(v DeviceTypeWeightUnitValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DeviceTypeWeightUnit) GetLabel() DeviceTypeWeightUnitLabel {
	if o == nil || IsNil(o.Label) {
		var ret DeviceTypeWeightUnitLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeWeightUnit) GetLabelOk() (*DeviceTypeWeightUnitLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DeviceTypeWeightUnit) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given DeviceTypeWeightUnitLabel and assigns it to the Label field.
func (o *DeviceTypeWeightUnit) SetLabel(v DeviceTypeWeightUnitLabel) {
	o.Label = &v
}

func (o DeviceTypeWeightUnit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceTypeWeightUnit) ToMap() (map[string]interface{}, error) {
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

func (o *DeviceTypeWeightUnit) UnmarshalJSON(data []byte) (err error) {
	varDeviceTypeWeightUnit := _DeviceTypeWeightUnit{}

	err = json.Unmarshal(data, &varDeviceTypeWeightUnit)

	if err != nil {
		return err
	}

	*o = DeviceTypeWeightUnit(varDeviceTypeWeightUnit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceTypeWeightUnit struct {
	value *DeviceTypeWeightUnit
	isSet bool
}

func (v NullableDeviceTypeWeightUnit) Get() *DeviceTypeWeightUnit {
	return v.value
}

func (v *NullableDeviceTypeWeightUnit) Set(val *DeviceTypeWeightUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTypeWeightUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTypeWeightUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTypeWeightUnit(val *DeviceTypeWeightUnit) *NullableDeviceTypeWeightUnit {
	return &NullableDeviceTypeWeightUnit{value: val, isSet: true}
}

func (v NullableDeviceTypeWeightUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTypeWeightUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


