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

// checks if the PatchedInventoryItemTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedInventoryItemTemplateRequest{}

// PatchedInventoryItemTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedInventoryItemTemplateRequest struct {
	DeviceType *BriefDeviceTypeRequest `json:"device_type,omitempty"`
	Parent NullableInt32 `json:"parent,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Role NullableBriefInventoryItemRoleRequest `json:"role,omitempty"`
	Manufacturer NullableBriefManufacturerRequest `json:"manufacturer,omitempty"`
	// Manufacturer-assigned part identifier
	PartId *string `json:"part_id,omitempty"`
	Description *string `json:"description,omitempty"`
	ComponentType NullableString `json:"component_type,omitempty"`
	ComponentId NullableInt64 `json:"component_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedInventoryItemTemplateRequest PatchedInventoryItemTemplateRequest

// NewPatchedInventoryItemTemplateRequest instantiates a new PatchedInventoryItemTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedInventoryItemTemplateRequest() *PatchedInventoryItemTemplateRequest {
	this := PatchedInventoryItemTemplateRequest{}
	return &this
}

// NewPatchedInventoryItemTemplateRequestWithDefaults instantiates a new PatchedInventoryItemTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedInventoryItemTemplateRequestWithDefaults() *PatchedInventoryItemTemplateRequest {
	this := PatchedInventoryItemTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *PatchedInventoryItemTemplateRequest) GetDeviceType() BriefDeviceTypeRequest {
	if o == nil || IsNil(o.DeviceType) {
		var ret BriefDeviceTypeRequest
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PatchedInventoryItemTemplateRequest) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given BriefDeviceTypeRequest and assigns it to the DeviceType field.
func (o *PatchedInventoryItemTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest) {
	o.DeviceType = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedInventoryItemTemplateRequest) GetParent() int32 {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret int32
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedInventoryItemTemplateRequest) GetParentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *PatchedInventoryItemTemplateRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableInt32 and assigns it to the Parent field.
func (o *PatchedInventoryItemTemplateRequest) SetParent(v int32) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *PatchedInventoryItemTemplateRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *PatchedInventoryItemTemplateRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedInventoryItemTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedInventoryItemTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedInventoryItemTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedInventoryItemTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedInventoryItemTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedInventoryItemTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedInventoryItemTemplateRequest) GetRole() BriefInventoryItemRoleRequest {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefInventoryItemRoleRequest
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedInventoryItemTemplateRequest) GetRoleOk() (*BriefInventoryItemRoleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedInventoryItemTemplateRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefInventoryItemRoleRequest and assigns it to the Role field.
func (o *PatchedInventoryItemTemplateRequest) SetRole(v BriefInventoryItemRoleRequest) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *PatchedInventoryItemTemplateRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *PatchedInventoryItemTemplateRequest) UnsetRole() {
	o.Role.Unset()
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedInventoryItemTemplateRequest) GetManufacturer() BriefManufacturerRequest {
	if o == nil || IsNil(o.Manufacturer.Get()) {
		var ret BriefManufacturerRequest
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedInventoryItemTemplateRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *PatchedInventoryItemTemplateRequest) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableBriefManufacturerRequest and assigns it to the Manufacturer field.
func (o *PatchedInventoryItemTemplateRequest) SetManufacturer(v BriefManufacturerRequest) {
	o.Manufacturer.Set(&v)
}
// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *PatchedInventoryItemTemplateRequest) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *PatchedInventoryItemTemplateRequest) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *PatchedInventoryItemTemplateRequest) GetPartId() string {
	if o == nil || IsNil(o.PartId) {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemTemplateRequest) GetPartIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartId) {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *PatchedInventoryItemTemplateRequest) HasPartId() bool {
	if o != nil && !IsNil(o.PartId) {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *PatchedInventoryItemTemplateRequest) SetPartId(v string) {
	o.PartId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedInventoryItemTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedInventoryItemTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedInventoryItemTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedInventoryItemTemplateRequest) GetComponentType() string {
	if o == nil || IsNil(o.ComponentType.Get()) {
		var ret string
		return ret
	}
	return *o.ComponentType.Get()
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedInventoryItemTemplateRequest) GetComponentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComponentType.Get(), o.ComponentType.IsSet()
}

// HasComponentType returns a boolean if a field has been set.
func (o *PatchedInventoryItemTemplateRequest) HasComponentType() bool {
	if o != nil && o.ComponentType.IsSet() {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given NullableString and assigns it to the ComponentType field.
func (o *PatchedInventoryItemTemplateRequest) SetComponentType(v string) {
	o.ComponentType.Set(&v)
}
// SetComponentTypeNil sets the value for ComponentType to be an explicit nil
func (o *PatchedInventoryItemTemplateRequest) SetComponentTypeNil() {
	o.ComponentType.Set(nil)
}

// UnsetComponentType ensures that no value is present for ComponentType, not even an explicit nil
func (o *PatchedInventoryItemTemplateRequest) UnsetComponentType() {
	o.ComponentType.Unset()
}

// GetComponentId returns the ComponentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedInventoryItemTemplateRequest) GetComponentId() int64 {
	if o == nil || IsNil(o.ComponentId.Get()) {
		var ret int64
		return ret
	}
	return *o.ComponentId.Get()
}

// GetComponentIdOk returns a tuple with the ComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedInventoryItemTemplateRequest) GetComponentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComponentId.Get(), o.ComponentId.IsSet()
}

// HasComponentId returns a boolean if a field has been set.
func (o *PatchedInventoryItemTemplateRequest) HasComponentId() bool {
	if o != nil && o.ComponentId.IsSet() {
		return true
	}

	return false
}

// SetComponentId gets a reference to the given NullableInt64 and assigns it to the ComponentId field.
func (o *PatchedInventoryItemTemplateRequest) SetComponentId(v int64) {
	o.ComponentId.Set(&v)
}
// SetComponentIdNil sets the value for ComponentId to be an explicit nil
func (o *PatchedInventoryItemTemplateRequest) SetComponentIdNil() {
	o.ComponentId.Set(nil)
}

// UnsetComponentId ensures that no value is present for ComponentId, not even an explicit nil
func (o *PatchedInventoryItemTemplateRequest) UnsetComponentId() {
	o.ComponentId.Unset()
}

func (o PatchedInventoryItemTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedInventoryItemTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceType) {
		toSerialize["device_type"] = o.DeviceType
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Manufacturer.IsSet() {
		toSerialize["manufacturer"] = o.Manufacturer.Get()
	}
	if !IsNil(o.PartId) {
		toSerialize["part_id"] = o.PartId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.ComponentType.IsSet() {
		toSerialize["component_type"] = o.ComponentType.Get()
	}
	if o.ComponentId.IsSet() {
		toSerialize["component_id"] = o.ComponentId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedInventoryItemTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedInventoryItemTemplateRequest := _PatchedInventoryItemTemplateRequest{}

	err = json.Unmarshal(data, &varPatchedInventoryItemTemplateRequest)

	if err != nil {
		return err
	}

	*o = PatchedInventoryItemTemplateRequest(varPatchedInventoryItemTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "role")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "part_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "component_type")
		delete(additionalProperties, "component_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedInventoryItemTemplateRequest struct {
	value *PatchedInventoryItemTemplateRequest
	isSet bool
}

func (v NullablePatchedInventoryItemTemplateRequest) Get() *PatchedInventoryItemTemplateRequest {
	return v.value
}

func (v *NullablePatchedInventoryItemTemplateRequest) Set(val *PatchedInventoryItemTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedInventoryItemTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedInventoryItemTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedInventoryItemTemplateRequest(val *PatchedInventoryItemTemplateRequest) *NullablePatchedInventoryItemTemplateRequest {
	return &NullablePatchedInventoryItemTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedInventoryItemTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedInventoryItemTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


