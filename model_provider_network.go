/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.10 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the ProviderNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderNetwork{}

// ProviderNetwork Adds support for custom fields and tags.
type ProviderNetwork struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Provider BriefProvider `json:"provider"`
	Name string `json:"name"`
	ServiceId *string `json:"service_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _ProviderNetwork ProviderNetwork

// NewProviderNetwork instantiates a new ProviderNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderNetwork(id int32, url string, display string, provider BriefProvider, name string, created NullableTime, lastUpdated NullableTime) *ProviderNetwork {
	this := ProviderNetwork{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Provider = provider
	this.Name = name
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewProviderNetworkWithDefaults instantiates a new ProviderNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderNetworkWithDefaults() *ProviderNetwork {
	this := ProviderNetwork{}
	return &this
}

// GetId returns the Id field value
func (o *ProviderNetwork) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProviderNetwork) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProviderNetwork) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ProviderNetwork) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ProviderNetwork) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ProviderNetwork) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *ProviderNetwork) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ProviderNetwork) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ProviderNetwork) SetDisplay(v string) {
	o.Display = v
}

// GetProvider returns the Provider field value
func (o *ProviderNetwork) GetProvider() BriefProvider {
	if o == nil {
		var ret BriefProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *ProviderNetwork) GetProviderOk() (*BriefProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *ProviderNetwork) SetProvider(v BriefProvider) {
	o.Provider = v
}

// GetName returns the Name field value
func (o *ProviderNetwork) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProviderNetwork) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProviderNetwork) SetName(v string) {
	o.Name = v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ProviderNetwork) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderNetwork) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ProviderNetwork) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *ProviderNetwork) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProviderNetwork) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderNetwork) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProviderNetwork) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProviderNetwork) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ProviderNetwork) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderNetwork) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ProviderNetwork) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ProviderNetwork) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProviderNetwork) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderNetwork) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProviderNetwork) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *ProviderNetwork) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ProviderNetwork) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderNetwork) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ProviderNetwork) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ProviderNetwork) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ProviderNetwork) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderNetwork) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *ProviderNetwork) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ProviderNetwork) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderNetwork) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *ProviderNetwork) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o ProviderNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["provider"] = o.Provider
	toSerialize["name"] = o.Name
	if !IsNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProviderNetwork) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"provider",
		"name",
		"created",
		"last_updated",
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

	varProviderNetwork := _ProviderNetwork{}

	err = json.Unmarshal(data, &varProviderNetwork)

	if err != nil {
		return err
	}

	*o = ProviderNetwork(varProviderNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "name")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProviderNetwork struct {
	value *ProviderNetwork
	isSet bool
}

func (v NullableProviderNetwork) Get() *ProviderNetwork {
	return v.value
}

func (v *NullableProviderNetwork) Set(val *ProviderNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderNetwork(val *ProviderNetwork) *NullableProviderNetwork {
	return &NullableProviderNetwork{value: val, isSet: true}
}

func (v NullableProviderNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


