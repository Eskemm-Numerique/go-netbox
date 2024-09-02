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

// checks if the BriefL2VPNRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefL2VPNRequest{}

// BriefL2VPNRequest Adds support for custom fields and tags.
type BriefL2VPNRequest struct {
	Identifier NullableInt64 `json:"identifier,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	// * `vpws` - VPWS * `vpls` - VPLS * `vxlan` - VXLAN * `vxlan-evpn` - VXLAN-EVPN * `mpls-evpn` - MPLS EVPN * `pbb-evpn` - PBB EVPN * `epl` - EPL * `evpl` - EVPL * `ep-lan` - Ethernet Private LAN * `evp-lan` - Ethernet Virtual Private LAN * `ep-tree` - Ethernet Private Tree * `evp-tree` - Ethernet Virtual Private Tree
	Type *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefL2VPNRequest BriefL2VPNRequest

// NewBriefL2VPNRequest instantiates a new BriefL2VPNRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefL2VPNRequest(name string, slug string) *BriefL2VPNRequest {
	this := BriefL2VPNRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewBriefL2VPNRequestWithDefaults instantiates a new BriefL2VPNRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefL2VPNRequestWithDefaults() *BriefL2VPNRequest {
	this := BriefL2VPNRequest{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BriefL2VPNRequest) GetIdentifier() int64 {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret int64
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BriefL2VPNRequest) GetIdentifierOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *BriefL2VPNRequest) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableInt64 and assigns it to the Identifier field.
func (o *BriefL2VPNRequest) SetIdentifier(v int64) {
	o.Identifier.Set(&v)
}
// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *BriefL2VPNRequest) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *BriefL2VPNRequest) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetName returns the Name field value
func (o *BriefL2VPNRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefL2VPNRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefL2VPNRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *BriefL2VPNRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefL2VPNRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefL2VPNRequest) SetSlug(v string) {
	o.Slug = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BriefL2VPNRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefL2VPNRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BriefL2VPNRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BriefL2VPNRequest) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefL2VPNRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefL2VPNRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefL2VPNRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefL2VPNRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefL2VPNRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefL2VPNRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefL2VPNRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBriefL2VPNRequest := _BriefL2VPNRequest{}

	err = json.Unmarshal(data, &varBriefL2VPNRequest)

	if err != nil {
		return err
	}

	*o = BriefL2VPNRequest(varBriefL2VPNRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefL2VPNRequest struct {
	value *BriefL2VPNRequest
	isSet bool
}

func (v NullableBriefL2VPNRequest) Get() *BriefL2VPNRequest {
	return v.value
}

func (v *NullableBriefL2VPNRequest) Set(val *BriefL2VPNRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefL2VPNRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefL2VPNRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefL2VPNRequest(val *BriefL2VPNRequest) *NullableBriefL2VPNRequest {
	return &NullableBriefL2VPNRequest{value: val, isSet: true}
}

func (v NullableBriefL2VPNRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefL2VPNRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


