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

// checks if the WritableIPSecPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableIPSecPolicyRequest{}

// WritableIPSecPolicyRequest Adds support for custom fields and tags.
type WritableIPSecPolicyRequest struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Proposals []int32 `json:"proposals,omitempty"`
	PfsGroup NullablePatchedWritableIPSecPolicyRequestPfsGroup `json:"pfs_group,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableIPSecPolicyRequest WritableIPSecPolicyRequest

// NewWritableIPSecPolicyRequest instantiates a new WritableIPSecPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableIPSecPolicyRequest(name string) *WritableIPSecPolicyRequest {
	this := WritableIPSecPolicyRequest{}
	this.Name = name
	return &this
}

// NewWritableIPSecPolicyRequestWithDefaults instantiates a new WritableIPSecPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableIPSecPolicyRequestWithDefaults() *WritableIPSecPolicyRequest {
	this := WritableIPSecPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableIPSecPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableIPSecPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableIPSecPolicyRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableIPSecPolicyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPSecPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableIPSecPolicyRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableIPSecPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetProposals returns the Proposals field value if set, zero value otherwise.
func (o *WritableIPSecPolicyRequest) GetProposals() []int32 {
	if o == nil || IsNil(o.Proposals) {
		var ret []int32
		return ret
	}
	return o.Proposals
}

// GetProposalsOk returns a tuple with the Proposals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPSecPolicyRequest) GetProposalsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Proposals) {
		return nil, false
	}
	return o.Proposals, true
}

// HasProposals returns a boolean if a field has been set.
func (o *WritableIPSecPolicyRequest) HasProposals() bool {
	if o != nil && !IsNil(o.Proposals) {
		return true
	}

	return false
}

// SetProposals gets a reference to the given []int32 and assigns it to the Proposals field.
func (o *WritableIPSecPolicyRequest) SetProposals(v []int32) {
	o.Proposals = v
}

// GetPfsGroup returns the PfsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIPSecPolicyRequest) GetPfsGroup() PatchedWritableIPSecPolicyRequestPfsGroup {
	if o == nil || IsNil(o.PfsGroup.Get()) {
		var ret PatchedWritableIPSecPolicyRequestPfsGroup
		return ret
	}
	return *o.PfsGroup.Get()
}

// GetPfsGroupOk returns a tuple with the PfsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIPSecPolicyRequest) GetPfsGroupOk() (*PatchedWritableIPSecPolicyRequestPfsGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.PfsGroup.Get(), o.PfsGroup.IsSet()
}

// HasPfsGroup returns a boolean if a field has been set.
func (o *WritableIPSecPolicyRequest) HasPfsGroup() bool {
	if o != nil && o.PfsGroup.IsSet() {
		return true
	}

	return false
}

// SetPfsGroup gets a reference to the given NullablePatchedWritableIPSecPolicyRequestPfsGroup and assigns it to the PfsGroup field.
func (o *WritableIPSecPolicyRequest) SetPfsGroup(v PatchedWritableIPSecPolicyRequestPfsGroup) {
	o.PfsGroup.Set(&v)
}
// SetPfsGroupNil sets the value for PfsGroup to be an explicit nil
func (o *WritableIPSecPolicyRequest) SetPfsGroupNil() {
	o.PfsGroup.Set(nil)
}

// UnsetPfsGroup ensures that no value is present for PfsGroup, not even an explicit nil
func (o *WritableIPSecPolicyRequest) UnsetPfsGroup() {
	o.PfsGroup.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableIPSecPolicyRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPSecPolicyRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableIPSecPolicyRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableIPSecPolicyRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableIPSecPolicyRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPSecPolicyRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableIPSecPolicyRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableIPSecPolicyRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableIPSecPolicyRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPSecPolicyRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableIPSecPolicyRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableIPSecPolicyRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableIPSecPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableIPSecPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Proposals) {
		toSerialize["proposals"] = o.Proposals
	}
	if o.PfsGroup.IsSet() {
		toSerialize["pfs_group"] = o.PfsGroup.Get()
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableIPSecPolicyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varWritableIPSecPolicyRequest := _WritableIPSecPolicyRequest{}

	err = json.Unmarshal(data, &varWritableIPSecPolicyRequest)

	if err != nil {
		return err
	}

	*o = WritableIPSecPolicyRequest(varWritableIPSecPolicyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "proposals")
		delete(additionalProperties, "pfs_group")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableIPSecPolicyRequest struct {
	value *WritableIPSecPolicyRequest
	isSet bool
}

func (v NullableWritableIPSecPolicyRequest) Get() *WritableIPSecPolicyRequest {
	return v.value
}

func (v *NullableWritableIPSecPolicyRequest) Set(val *WritableIPSecPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableIPSecPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableIPSecPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableIPSecPolicyRequest(val *WritableIPSecPolicyRequest) *NullableWritableIPSecPolicyRequest {
	return &NullableWritableIPSecPolicyRequest{value: val, isSet: true}
}

func (v NullableWritableIPSecPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableIPSecPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


