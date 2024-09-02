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

// checks if the CustomFieldUiVisible type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldUiVisible{}

// CustomFieldUiVisible struct for CustomFieldUiVisible
type CustomFieldUiVisible struct {
	// * `always` - Always * `if-set` - If set * `hidden` - Hidden
	Value *string `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomFieldUiVisible CustomFieldUiVisible

// NewCustomFieldUiVisible instantiates a new CustomFieldUiVisible object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldUiVisible() *CustomFieldUiVisible {
	this := CustomFieldUiVisible{}
	return &this
}

// NewCustomFieldUiVisibleWithDefaults instantiates a new CustomFieldUiVisible object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldUiVisibleWithDefaults() *CustomFieldUiVisible {
	this := CustomFieldUiVisible{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CustomFieldUiVisible) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldUiVisible) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CustomFieldUiVisible) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CustomFieldUiVisible) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CustomFieldUiVisible) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldUiVisible) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CustomFieldUiVisible) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CustomFieldUiVisible) SetLabel(v string) {
	o.Label = &v
}

func (o CustomFieldUiVisible) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldUiVisible) ToMap() (map[string]interface{}, error) {
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

func (o *CustomFieldUiVisible) UnmarshalJSON(data []byte) (err error) {
	varCustomFieldUiVisible := _CustomFieldUiVisible{}

	err = json.Unmarshal(data, &varCustomFieldUiVisible)

	if err != nil {
		return err
	}

	*o = CustomFieldUiVisible(varCustomFieldUiVisible)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomFieldUiVisible struct {
	value *CustomFieldUiVisible
	isSet bool
}

func (v NullableCustomFieldUiVisible) Get() *CustomFieldUiVisible {
	return v.value
}

func (v *NullableCustomFieldUiVisible) Set(val *CustomFieldUiVisible) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldUiVisible) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldUiVisible) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldUiVisible(val *CustomFieldUiVisible) *NullableCustomFieldUiVisible {
	return &NullableCustomFieldUiVisible{value: val, isSet: true}
}

func (v NullableCustomFieldUiVisible) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldUiVisible) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


