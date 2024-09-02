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

// checks if the IPAddressRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPAddressRole{}

// IPAddressRole struct for IPAddressRole
type IPAddressRole struct {
	// * `loopback` - Loopback * `secondary` - Secondary * `anycast` - Anycast * `vip` - VIP * `vrrp` - VRRP * `hsrp` - HSRP * `glbp` - GLBP * `carp` - CARP
	Value *string `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IPAddressRole IPAddressRole

// NewIPAddressRole instantiates a new IPAddressRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPAddressRole() *IPAddressRole {
	this := IPAddressRole{}
	return &this
}

// NewIPAddressRoleWithDefaults instantiates a new IPAddressRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPAddressRoleWithDefaults() *IPAddressRole {
	this := IPAddressRole{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IPAddressRole) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddressRole) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IPAddressRole) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IPAddressRole) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IPAddressRole) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddressRole) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IPAddressRole) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *IPAddressRole) SetLabel(v string) {
	o.Label = &v
}

func (o IPAddressRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPAddressRole) ToMap() (map[string]interface{}, error) {
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

func (o *IPAddressRole) UnmarshalJSON(data []byte) (err error) {
	varIPAddressRole := _IPAddressRole{}

	err = json.Unmarshal(data, &varIPAddressRole)

	if err != nil {
		return err
	}

	*o = IPAddressRole(varIPAddressRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPAddressRole struct {
	value *IPAddressRole
	isSet bool
}

func (v NullableIPAddressRole) Get() *IPAddressRole {
	return v.value
}

func (v *NullableIPAddressRole) Set(val *IPAddressRole) {
	v.value = val
	v.isSet = true
}

func (v NullableIPAddressRole) IsSet() bool {
	return v.isSet
}

func (v *NullableIPAddressRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPAddressRole(val *IPAddressRole) *NullableIPAddressRole {
	return &NullableIPAddressRole{value: val, isSet: true}
}

func (v NullableIPAddressRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPAddressRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


