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

// checks if the ModuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleRequest{}

// ModuleRequest Adds support for custom fields and tags.
type ModuleRequest struct {
	Device BriefDeviceRequest `json:"device"`
	ModuleBay NestedModuleBayRequest `json:"module_bay"`
	ModuleType BriefModuleTypeRequest `json:"module_type"`
	Status *ModuleStatusValue `json:"status,omitempty"`
	Serial *string `json:"serial,omitempty"`
	// A unique tag used to identify this device
	AssetTag NullableString `json:"asset_tag,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModuleRequest ModuleRequest

// NewModuleRequest instantiates a new ModuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleRequest(device BriefDeviceRequest, moduleBay NestedModuleBayRequest, moduleType BriefModuleTypeRequest) *ModuleRequest {
	this := ModuleRequest{}
	this.Device = device
	this.ModuleBay = moduleBay
	this.ModuleType = moduleType
	return &this
}

// NewModuleRequestWithDefaults instantiates a new ModuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleRequestWithDefaults() *ModuleRequest {
	this := ModuleRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *ModuleRequest) GetDevice() BriefDeviceRequest {
	if o == nil {
		var ret BriefDeviceRequest
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *ModuleRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *ModuleRequest) SetDevice(v BriefDeviceRequest) {
	o.Device = v
}

// GetModuleBay returns the ModuleBay field value
func (o *ModuleRequest) GetModuleBay() NestedModuleBayRequest {
	if o == nil {
		var ret NestedModuleBayRequest
		return ret
	}

	return o.ModuleBay
}

// GetModuleBayOk returns a tuple with the ModuleBay field value
// and a boolean to check if the value has been set.
func (o *ModuleRequest) GetModuleBayOk() (*NestedModuleBayRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleBay, true
}

// SetModuleBay sets field value
func (o *ModuleRequest) SetModuleBay(v NestedModuleBayRequest) {
	o.ModuleBay = v
}

// GetModuleType returns the ModuleType field value
func (o *ModuleRequest) GetModuleType() BriefModuleTypeRequest {
	if o == nil {
		var ret BriefModuleTypeRequest
		return ret
	}

	return o.ModuleType
}

// GetModuleTypeOk returns a tuple with the ModuleType field value
// and a boolean to check if the value has been set.
func (o *ModuleRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleType, true
}

// SetModuleType sets field value
func (o *ModuleRequest) SetModuleType(v BriefModuleTypeRequest) {
	o.ModuleType = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModuleRequest) GetStatus() ModuleStatusValue {
	if o == nil || IsNil(o.Status) {
		var ret ModuleStatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleRequest) GetStatusOk() (*ModuleStatusValue, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModuleRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ModuleStatusValue and assigns it to the Status field.
func (o *ModuleRequest) SetStatus(v ModuleStatusValue) {
	o.Status = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *ModuleRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *ModuleRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *ModuleRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModuleRequest) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModuleRequest) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *ModuleRequest) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *ModuleRequest) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}
// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *ModuleRequest) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *ModuleRequest) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModuleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModuleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModuleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ModuleRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ModuleRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ModuleRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ModuleRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ModuleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *ModuleRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ModuleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ModuleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ModuleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ModuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	toSerialize["module_bay"] = o.ModuleBay
	toSerialize["module_type"] = o.ModuleType
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device",
		"module_bay",
		"module_type",
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

	varModuleRequest := _ModuleRequest{}

	err = json.Unmarshal(data, &varModuleRequest)

	if err != nil {
		return err
	}

	*o = ModuleRequest(varModuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module_bay")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "asset_tag")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleRequest struct {
	value *ModuleRequest
	isSet bool
}

func (v NullableModuleRequest) Get() *ModuleRequest {
	return v.value
}

func (v *NullableModuleRequest) Set(val *ModuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleRequest(val *ModuleRequest) *NullableModuleRequest {
	return &NullableModuleRequest{value: val, isSet: true}
}

func (v NullableModuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


