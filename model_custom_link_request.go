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

// checks if the CustomLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomLinkRequest{}

// CustomLinkRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type CustomLinkRequest struct {
	ObjectTypes []string `json:"object_types"`
	Name string `json:"name"`
	Enabled *bool `json:"enabled,omitempty"`
	// Jinja2 template code for link text
	LinkText string `json:"link_text"`
	// Jinja2 template code for link URL
	LinkUrl string `json:"link_url"`
	Weight *int32 `json:"weight,omitempty"`
	// Links with the same group will appear as a dropdown menu
	GroupName *string `json:"group_name,omitempty"`
	ButtonClass *CustomLinkButtonClass `json:"button_class,omitempty"`
	// Force link to open in a new window
	NewWindow *bool `json:"new_window,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomLinkRequest CustomLinkRequest

// NewCustomLinkRequest instantiates a new CustomLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomLinkRequest(objectTypes []string, name string, linkText string, linkUrl string) *CustomLinkRequest {
	this := CustomLinkRequest{}
	this.ObjectTypes = objectTypes
	this.Name = name
	this.LinkText = linkText
	this.LinkUrl = linkUrl
	return &this
}

// NewCustomLinkRequestWithDefaults instantiates a new CustomLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomLinkRequestWithDefaults() *CustomLinkRequest {
	this := CustomLinkRequest{}
	return &this
}

// GetObjectTypes returns the ObjectTypes field value
func (o *CustomLinkRequest) GetObjectTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value
// and a boolean to check if the value has been set.
func (o *CustomLinkRequest) GetObjectTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectTypes, true
}

// SetObjectTypes sets field value
func (o *CustomLinkRequest) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetName returns the Name field value
func (o *CustomLinkRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomLinkRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomLinkRequest) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustomLinkRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLinkRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustomLinkRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustomLinkRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLinkText returns the LinkText field value
func (o *CustomLinkRequest) GetLinkText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkText
}

// GetLinkTextOk returns a tuple with the LinkText field value
// and a boolean to check if the value has been set.
func (o *CustomLinkRequest) GetLinkTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkText, true
}

// SetLinkText sets field value
func (o *CustomLinkRequest) SetLinkText(v string) {
	o.LinkText = v
}

// GetLinkUrl returns the LinkUrl field value
func (o *CustomLinkRequest) GetLinkUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkUrl
}

// GetLinkUrlOk returns a tuple with the LinkUrl field value
// and a boolean to check if the value has been set.
func (o *CustomLinkRequest) GetLinkUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkUrl, true
}

// SetLinkUrl sets field value
func (o *CustomLinkRequest) SetLinkUrl(v string) {
	o.LinkUrl = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *CustomLinkRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLinkRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *CustomLinkRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *CustomLinkRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *CustomLinkRequest) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLinkRequest) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *CustomLinkRequest) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *CustomLinkRequest) SetGroupName(v string) {
	o.GroupName = &v
}

// GetButtonClass returns the ButtonClass field value if set, zero value otherwise.
func (o *CustomLinkRequest) GetButtonClass() CustomLinkButtonClass {
	if o == nil || IsNil(o.ButtonClass) {
		var ret CustomLinkButtonClass
		return ret
	}
	return *o.ButtonClass
}

// GetButtonClassOk returns a tuple with the ButtonClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLinkRequest) GetButtonClassOk() (*CustomLinkButtonClass, bool) {
	if o == nil || IsNil(o.ButtonClass) {
		return nil, false
	}
	return o.ButtonClass, true
}

// HasButtonClass returns a boolean if a field has been set.
func (o *CustomLinkRequest) HasButtonClass() bool {
	if o != nil && !IsNil(o.ButtonClass) {
		return true
	}

	return false
}

// SetButtonClass gets a reference to the given CustomLinkButtonClass and assigns it to the ButtonClass field.
func (o *CustomLinkRequest) SetButtonClass(v CustomLinkButtonClass) {
	o.ButtonClass = &v
}

// GetNewWindow returns the NewWindow field value if set, zero value otherwise.
func (o *CustomLinkRequest) GetNewWindow() bool {
	if o == nil || IsNil(o.NewWindow) {
		var ret bool
		return ret
	}
	return *o.NewWindow
}

// GetNewWindowOk returns a tuple with the NewWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomLinkRequest) GetNewWindowOk() (*bool, bool) {
	if o == nil || IsNil(o.NewWindow) {
		return nil, false
	}
	return o.NewWindow, true
}

// HasNewWindow returns a boolean if a field has been set.
func (o *CustomLinkRequest) HasNewWindow() bool {
	if o != nil && !IsNil(o.NewWindow) {
		return true
	}

	return false
}

// SetNewWindow gets a reference to the given bool and assigns it to the NewWindow field.
func (o *CustomLinkRequest) SetNewWindow(v bool) {
	o.NewWindow = &v
}

func (o CustomLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_types"] = o.ObjectTypes
	toSerialize["name"] = o.Name
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["link_text"] = o.LinkText
	toSerialize["link_url"] = o.LinkUrl
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	if !IsNil(o.ButtonClass) {
		toSerialize["button_class"] = o.ButtonClass
	}
	if !IsNil(o.NewWindow) {
		toSerialize["new_window"] = o.NewWindow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomLinkRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_types",
		"name",
		"link_text",
		"link_url",
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

	varCustomLinkRequest := _CustomLinkRequest{}

	err = json.Unmarshal(data, &varCustomLinkRequest)

	if err != nil {
		return err
	}

	*o = CustomLinkRequest(varCustomLinkRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "link_text")
		delete(additionalProperties, "link_url")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "group_name")
		delete(additionalProperties, "button_class")
		delete(additionalProperties, "new_window")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomLinkRequest struct {
	value *CustomLinkRequest
	isSet bool
}

func (v NullableCustomLinkRequest) Get() *CustomLinkRequest {
	return v.value
}

func (v *NullableCustomLinkRequest) Set(val *CustomLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomLinkRequest(val *CustomLinkRequest) *NullableCustomLinkRequest {
	return &NullableCustomLinkRequest{value: val, isSet: true}
}

func (v NullableCustomLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


