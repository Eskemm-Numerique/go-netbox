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

// checks if the PatchedWritableIPSecPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableIPSecPolicyRequest{}

// PatchedWritableIPSecPolicyRequest Adds support for custom fields and tags.
type PatchedWritableIPSecPolicyRequest struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Proposals []int32 `json:"proposals,omitempty"`
	// Diffie-Hellman group for Perfect Forward Secrecy  * `1` - Group 1 * `2` - Group 2 * `5` - Group 5 * `14` - Group 14 * `15` - Group 15 * `16` - Group 16 * `17` - Group 17 * `18` - Group 18 * `19` - Group 19 * `20` - Group 20 * `21` - Group 21 * `22` - Group 22 * `23` - Group 23 * `24` - Group 24 * `25` - Group 25 * `26` - Group 26 * `27` - Group 27 * `28` - Group 28 * `29` - Group 29 * `30` - Group 30 * `31` - Group 31 * `32` - Group 32 * `33` - Group 33 * `34` - Group 34
	PfsGroup NullableInt32 `json:"pfs_group,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableIPSecPolicyRequest PatchedWritableIPSecPolicyRequest

// NewPatchedWritableIPSecPolicyRequest instantiates a new PatchedWritableIPSecPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableIPSecPolicyRequest() *PatchedWritableIPSecPolicyRequest {
	this := PatchedWritableIPSecPolicyRequest{}
	return &this
}

// NewPatchedWritableIPSecPolicyRequestWithDefaults instantiates a new PatchedWritableIPSecPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableIPSecPolicyRequestWithDefaults() *PatchedWritableIPSecPolicyRequest {
	this := PatchedWritableIPSecPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableIPSecPolicyRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableIPSecPolicyRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableIPSecPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableIPSecPolicyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableIPSecPolicyRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableIPSecPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetProposals returns the Proposals field value if set, zero value otherwise.
func (o *PatchedWritableIPSecPolicyRequest) GetProposals() []int32 {
	if o == nil || IsNil(o.Proposals) {
		var ret []int32
		return ret
	}
	return o.Proposals
}

// GetProposalsOk returns a tuple with the Proposals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecPolicyRequest) GetProposalsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Proposals) {
		return nil, false
	}
	return o.Proposals, true
}

// HasProposals returns a boolean if a field has been set.
func (o *PatchedWritableIPSecPolicyRequest) HasProposals() bool {
	if o != nil && !IsNil(o.Proposals) {
		return true
	}

	return false
}

// SetProposals gets a reference to the given []int32 and assigns it to the Proposals field.
func (o *PatchedWritableIPSecPolicyRequest) SetProposals(v []int32) {
	o.Proposals = v
}

// GetPfsGroup returns the PfsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableIPSecPolicyRequest) GetPfsGroup() int32 {
	if o == nil || IsNil(o.PfsGroup.Get()) {
		var ret int32
		return ret
	}
	return *o.PfsGroup.Get()
}

// GetPfsGroupOk returns a tuple with the PfsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableIPSecPolicyRequest) GetPfsGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PfsGroup.Get(), o.PfsGroup.IsSet()
}

// HasPfsGroup returns a boolean if a field has been set.
func (o *PatchedWritableIPSecPolicyRequest) HasPfsGroup() bool {
	if o != nil && o.PfsGroup.IsSet() {
		return true
	}

	return false
}

// SetPfsGroup gets a reference to the given NullableInt32 and assigns it to the PfsGroup field.
func (o *PatchedWritableIPSecPolicyRequest) SetPfsGroup(v int32) {
	o.PfsGroup.Set(&v)
}
// SetPfsGroupNil sets the value for PfsGroup to be an explicit nil
func (o *PatchedWritableIPSecPolicyRequest) SetPfsGroupNil() {
	o.PfsGroup.Set(nil)
}

// UnsetPfsGroup ensures that no value is present for PfsGroup, not even an explicit nil
func (o *PatchedWritableIPSecPolicyRequest) UnsetPfsGroup() {
	o.PfsGroup.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableIPSecPolicyRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecPolicyRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableIPSecPolicyRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableIPSecPolicyRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableIPSecPolicyRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecPolicyRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableIPSecPolicyRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableIPSecPolicyRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableIPSecPolicyRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecPolicyRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableIPSecPolicyRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableIPSecPolicyRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableIPSecPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableIPSecPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
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

func (o *PatchedWritableIPSecPolicyRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableIPSecPolicyRequest := _PatchedWritableIPSecPolicyRequest{}

	err = json.Unmarshal(data, &varPatchedWritableIPSecPolicyRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableIPSecPolicyRequest(varPatchedWritableIPSecPolicyRequest)

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

type NullablePatchedWritableIPSecPolicyRequest struct {
	value *PatchedWritableIPSecPolicyRequest
	isSet bool
}

func (v NullablePatchedWritableIPSecPolicyRequest) Get() *PatchedWritableIPSecPolicyRequest {
	return v.value
}

func (v *NullablePatchedWritableIPSecPolicyRequest) Set(val *PatchedWritableIPSecPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableIPSecPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableIPSecPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableIPSecPolicyRequest(val *PatchedWritableIPSecPolicyRequest) *NullablePatchedWritableIPSecPolicyRequest {
	return &NullablePatchedWritableIPSecPolicyRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableIPSecPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableIPSecPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


