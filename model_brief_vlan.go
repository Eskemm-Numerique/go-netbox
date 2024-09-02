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

// checks if the BriefVLAN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefVLAN{}

// BriefVLAN Adds support for custom fields and tags.
type BriefVLAN struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	// Numeric VLAN ID (1-4094)
	Vid int32 `json:"vid"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefVLAN BriefVLAN

// NewBriefVLAN instantiates a new BriefVLAN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefVLAN(id int32, url string, display string, vid int32, name string) *BriefVLAN {
	this := BriefVLAN{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Vid = vid
	this.Name = name
	return &this
}

// NewBriefVLANWithDefaults instantiates a new BriefVLAN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefVLANWithDefaults() *BriefVLAN {
	this := BriefVLAN{}
	return &this
}

// GetId returns the Id field value
func (o *BriefVLAN) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefVLAN) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefVLAN) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefVLAN) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefVLAN) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefVLAN) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefVLAN) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefVLAN) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefVLAN) SetDisplay(v string) {
	o.Display = v
}

// GetVid returns the Vid field value
func (o *BriefVLAN) GetVid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vid
}

// GetVidOk returns a tuple with the Vid field value
// and a boolean to check if the value has been set.
func (o *BriefVLAN) GetVidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vid, true
}

// SetVid sets field value
func (o *BriefVLAN) SetVid(v int32) {
	o.Vid = v
}

// GetName returns the Name field value
func (o *BriefVLAN) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefVLAN) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefVLAN) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefVLAN) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefVLAN) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefVLAN) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefVLAN) SetDescription(v string) {
	o.Description = &v
}

func (o BriefVLAN) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefVLAN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["vid"] = o.Vid
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefVLAN) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"vid",
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

	varBriefVLAN := _BriefVLAN{}

	err = json.Unmarshal(data, &varBriefVLAN)

	if err != nil {
		return err
	}

	*o = BriefVLAN(varBriefVLAN)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "vid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefVLAN struct {
	value *BriefVLAN
	isSet bool
}

func (v NullableBriefVLAN) Get() *BriefVLAN {
	return v.value
}

func (v *NullableBriefVLAN) Set(val *BriefVLAN) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefVLAN) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefVLAN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefVLAN(val *BriefVLAN) *NullableBriefVLAN {
	return &NullableBriefVLAN{value: val, isSet: true}
}

func (v NullableBriefVLAN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefVLAN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


