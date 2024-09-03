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

// checks if the CableRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CableRequest{}

// CableRequest Adds support for custom fields and tags.
type CableRequest struct {
	Type *CableType `json:"type,omitempty"`
	ATerminations []GenericObjectRequest `json:"a_terminations,omitempty"`
	BTerminations []GenericObjectRequest `json:"b_terminations,omitempty"`
	Status *CableStatusValue `json:"status,omitempty"`
	Tenant NullableBriefTenantRequest `json:"tenant,omitempty"`
	Label *string `json:"label,omitempty"`
	Color *string `json:"color,omitempty" validate:"regexp=^[0-9a-f]{6}$"`
	Length NullableFloat64 `json:"length,omitempty"`
	LengthUnit NullableCableRequestLengthUnit `json:"length_unit,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CableRequest CableRequest

// NewCableRequest instantiates a new CableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCableRequest() *CableRequest {
	this := CableRequest{}
	return &this
}

// NewCableRequestWithDefaults instantiates a new CableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCableRequestWithDefaults() *CableRequest {
	this := CableRequest{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CableRequest) GetType() CableType {
	if o == nil || IsNil(o.Type) {
		var ret CableType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableRequest) GetTypeOk() (*CableType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CableRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CableType and assigns it to the Type field.
func (o *CableRequest) SetType(v CableType) {
	o.Type = &v
}

// GetATerminations returns the ATerminations field value if set, zero value otherwise.
func (o *CableRequest) GetATerminations() []GenericObjectRequest {
	if o == nil || IsNil(o.ATerminations) {
		var ret []GenericObjectRequest
		return ret
	}
	return o.ATerminations
}

// GetATerminationsOk returns a tuple with the ATerminations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableRequest) GetATerminationsOk() ([]GenericObjectRequest, bool) {
	if o == nil || IsNil(o.ATerminations) {
		return nil, false
	}
	return o.ATerminations, true
}

// HasATerminations returns a boolean if a field has been set.
func (o *CableRequest) HasATerminations() bool {
	if o != nil && !IsNil(o.ATerminations) {
		return true
	}

	return false
}

// SetATerminations gets a reference to the given []GenericObjectRequest and assigns it to the ATerminations field.
func (o *CableRequest) SetATerminations(v []GenericObjectRequest) {
	o.ATerminations = v
}

// GetBTerminations returns the BTerminations field value if set, zero value otherwise.
func (o *CableRequest) GetBTerminations() []GenericObjectRequest {
	if o == nil || IsNil(o.BTerminations) {
		var ret []GenericObjectRequest
		return ret
	}
	return o.BTerminations
}

// GetBTerminationsOk returns a tuple with the BTerminations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableRequest) GetBTerminationsOk() ([]GenericObjectRequest, bool) {
	if o == nil || IsNil(o.BTerminations) {
		return nil, false
	}
	return o.BTerminations, true
}

// HasBTerminations returns a boolean if a field has been set.
func (o *CableRequest) HasBTerminations() bool {
	if o != nil && !IsNil(o.BTerminations) {
		return true
	}

	return false
}

// SetBTerminations gets a reference to the given []GenericObjectRequest and assigns it to the BTerminations field.
func (o *CableRequest) SetBTerminations(v []GenericObjectRequest) {
	o.BTerminations = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CableRequest) GetStatus() CableStatusValue {
	if o == nil || IsNil(o.Status) {
		var ret CableStatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableRequest) GetStatusOk() (*CableStatusValue, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CableRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CableStatusValue and assigns it to the Status field.
func (o *CableRequest) SetStatus(v CableStatusValue) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CableRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CableRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *CableRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *CableRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *CableRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *CableRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CableRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CableRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CableRequest) SetLabel(v string) {
	o.Label = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *CableRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *CableRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *CableRequest) SetColor(v string) {
	o.Color = &v
}

// GetLength returns the Length field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CableRequest) GetLength() float64 {
	if o == nil || IsNil(o.Length.Get()) {
		var ret float64
		return ret
	}
	return *o.Length.Get()
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CableRequest) GetLengthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Length.Get(), o.Length.IsSet()
}

// HasLength returns a boolean if a field has been set.
func (o *CableRequest) HasLength() bool {
	if o != nil && o.Length.IsSet() {
		return true
	}

	return false
}

// SetLength gets a reference to the given NullableFloat64 and assigns it to the Length field.
func (o *CableRequest) SetLength(v float64) {
	o.Length.Set(&v)
}
// SetLengthNil sets the value for Length to be an explicit nil
func (o *CableRequest) SetLengthNil() {
	o.Length.Set(nil)
}

// UnsetLength ensures that no value is present for Length, not even an explicit nil
func (o *CableRequest) UnsetLength() {
	o.Length.Unset()
}

// GetLengthUnit returns the LengthUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CableRequest) GetLengthUnit() CableRequestLengthUnit {
	if o == nil || IsNil(o.LengthUnit.Get()) {
		var ret CableRequestLengthUnit
		return ret
	}
	return *o.LengthUnit.Get()
}

// GetLengthUnitOk returns a tuple with the LengthUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CableRequest) GetLengthUnitOk() (*CableRequestLengthUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.LengthUnit.Get(), o.LengthUnit.IsSet()
}

// HasLengthUnit returns a boolean if a field has been set.
func (o *CableRequest) HasLengthUnit() bool {
	if o != nil && o.LengthUnit.IsSet() {
		return true
	}

	return false
}

// SetLengthUnit gets a reference to the given NullableCableRequestLengthUnit and assigns it to the LengthUnit field.
func (o *CableRequest) SetLengthUnit(v CableRequestLengthUnit) {
	o.LengthUnit.Set(&v)
}
// SetLengthUnitNil sets the value for LengthUnit to be an explicit nil
func (o *CableRequest) SetLengthUnitNil() {
	o.LengthUnit.Set(nil)
}

// UnsetLengthUnit ensures that no value is present for LengthUnit, not even an explicit nil
func (o *CableRequest) UnsetLengthUnit() {
	o.LengthUnit.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CableRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CableRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CableRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *CableRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *CableRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *CableRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CableRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CableRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *CableRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CableRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CableRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CableRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CableRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o CableRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CableRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ATerminations) {
		toSerialize["a_terminations"] = o.ATerminations
	}
	if !IsNil(o.BTerminations) {
		toSerialize["b_terminations"] = o.BTerminations
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if o.Length.IsSet() {
		toSerialize["length"] = o.Length.Get()
	}
	if o.LengthUnit.IsSet() {
		toSerialize["length_unit"] = o.LengthUnit.Get()
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

func (o *CableRequest) UnmarshalJSON(data []byte) (err error) {
	varCableRequest := _CableRequest{}

	err = json.Unmarshal(data, &varCableRequest)

	if err != nil {
		return err
	}

	*o = CableRequest(varCableRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "a_terminations")
		delete(additionalProperties, "b_terminations")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "label")
		delete(additionalProperties, "color")
		delete(additionalProperties, "length")
		delete(additionalProperties, "length_unit")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCableRequest struct {
	value *CableRequest
	isSet bool
}

func (v NullableCableRequest) Get() *CableRequest {
	return v.value
}

func (v *NullableCableRequest) Set(val *CableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableRequest(val *CableRequest) *NullableCableRequest {
	return &NullableCableRequest{value: val, isSet: true}
}

func (v NullableCableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


