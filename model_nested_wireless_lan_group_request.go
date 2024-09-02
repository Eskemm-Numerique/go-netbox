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

// checks if the NestedWirelessLANGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedWirelessLANGroupRequest{}

// NestedWirelessLANGroupRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedWirelessLANGroupRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	AdditionalProperties map[string]interface{}
}

type _NestedWirelessLANGroupRequest NestedWirelessLANGroupRequest

// NewNestedWirelessLANGroupRequest instantiates a new NestedWirelessLANGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedWirelessLANGroupRequest(name string, slug string) *NestedWirelessLANGroupRequest {
	this := NestedWirelessLANGroupRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedWirelessLANGroupRequestWithDefaults instantiates a new NestedWirelessLANGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedWirelessLANGroupRequestWithDefaults() *NestedWirelessLANGroupRequest {
	this := NestedWirelessLANGroupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedWirelessLANGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedWirelessLANGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedWirelessLANGroupRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedWirelessLANGroupRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedWirelessLANGroupRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedWirelessLANGroupRequest) SetSlug(v string) {
	o.Slug = v
}

func (o NestedWirelessLANGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedWirelessLANGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedWirelessLANGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNestedWirelessLANGroupRequest := _NestedWirelessLANGroupRequest{}

	err = json.Unmarshal(data, &varNestedWirelessLANGroupRequest)

	if err != nil {
		return err
	}

	*o = NestedWirelessLANGroupRequest(varNestedWirelessLANGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedWirelessLANGroupRequest struct {
	value *NestedWirelessLANGroupRequest
	isSet bool
}

func (v NullableNestedWirelessLANGroupRequest) Get() *NestedWirelessLANGroupRequest {
	return v.value
}

func (v *NullableNestedWirelessLANGroupRequest) Set(val *NestedWirelessLANGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedWirelessLANGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedWirelessLANGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedWirelessLANGroupRequest(val *NestedWirelessLANGroupRequest) *NullableNestedWirelessLANGroupRequest {
	return &NullableNestedWirelessLANGroupRequest{value: val, isSet: true}
}

func (v NullableNestedWirelessLANGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedWirelessLANGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


