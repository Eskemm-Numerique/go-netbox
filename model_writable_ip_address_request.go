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

// checks if the WritableIPAddressRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableIPAddressRequest{}

// WritableIPAddressRequest Adds support for custom fields and tags.
type WritableIPAddressRequest struct {
	Address string `json:"address"`
	Vrf NullableBriefVRFRequest `json:"vrf,omitempty"`
	Tenant NullableBriefTenantRequest `json:"tenant,omitempty"`
	Status *PatchedWritableIPAddressRequestStatus `json:"status,omitempty"`
	Role *PatchedWritableIPAddressRequestRole `json:"role,omitempty"`
	AssignedObjectType NullableString `json:"assigned_object_type,omitempty"`
	AssignedObjectId NullableInt64 `json:"assigned_object_id,omitempty"`
	// The IP for which this address is the \"outside\" IP
	NatInside NullableInt32 `json:"nat_inside,omitempty"`
	// Hostname or FQDN (not case-sensitive)
	DnsName *string `json:"dns_name,omitempty" validate:"regexp=^([0-9A-Za-z_-]+|\\\\*)(\\\\.[0-9A-Za-z_-]+)*\\\\.?$"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableIPAddressRequest WritableIPAddressRequest

// NewWritableIPAddressRequest instantiates a new WritableIPAddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableIPAddressRequest(address string) *WritableIPAddressRequest {
	this := WritableIPAddressRequest{}
	this.Address = address
	return &this
}

// NewWritableIPAddressRequestWithDefaults instantiates a new WritableIPAddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableIPAddressRequestWithDefaults() *WritableIPAddressRequest {
	this := WritableIPAddressRequest{}
	return &this
}

// GetAddress returns the Address field value
func (o *WritableIPAddressRequest) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *WritableIPAddressRequest) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *WritableIPAddressRequest) SetAddress(v string) {
	o.Address = v
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIPAddressRequest) GetVrf() BriefVRFRequest {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret BriefVRFRequest
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIPAddressRequest) GetVrfOk() (*BriefVRFRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *WritableIPAddressRequest) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableBriefVRFRequest and assigns it to the Vrf field.
func (o *WritableIPAddressRequest) SetVrf(v BriefVRFRequest) {
	o.Vrf.Set(&v)
}
// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *WritableIPAddressRequest) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *WritableIPAddressRequest) UnsetVrf() {
	o.Vrf.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIPAddressRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIPAddressRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableIPAddressRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *WritableIPAddressRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableIPAddressRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableIPAddressRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WritableIPAddressRequest) GetStatus() PatchedWritableIPAddressRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret PatchedWritableIPAddressRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddressRequest) GetStatusOk() (*PatchedWritableIPAddressRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WritableIPAddressRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PatchedWritableIPAddressRequestStatus and assigns it to the Status field.
func (o *WritableIPAddressRequest) SetStatus(v PatchedWritableIPAddressRequestStatus) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *WritableIPAddressRequest) GetRole() PatchedWritableIPAddressRequestRole {
	if o == nil || IsNil(o.Role) {
		var ret PatchedWritableIPAddressRequestRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddressRequest) GetRoleOk() (*PatchedWritableIPAddressRequestRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *WritableIPAddressRequest) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given PatchedWritableIPAddressRequestRole and assigns it to the Role field.
func (o *WritableIPAddressRequest) SetRole(v PatchedWritableIPAddressRequestRole) {
	o.Role = &v
}

// GetAssignedObjectType returns the AssignedObjectType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIPAddressRequest) GetAssignedObjectType() string {
	if o == nil || IsNil(o.AssignedObjectType.Get()) {
		var ret string
		return ret
	}
	return *o.AssignedObjectType.Get()
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIPAddressRequest) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedObjectType.Get(), o.AssignedObjectType.IsSet()
}

// HasAssignedObjectType returns a boolean if a field has been set.
func (o *WritableIPAddressRequest) HasAssignedObjectType() bool {
	if o != nil && o.AssignedObjectType.IsSet() {
		return true
	}

	return false
}

// SetAssignedObjectType gets a reference to the given NullableString and assigns it to the AssignedObjectType field.
func (o *WritableIPAddressRequest) SetAssignedObjectType(v string) {
	o.AssignedObjectType.Set(&v)
}
// SetAssignedObjectTypeNil sets the value for AssignedObjectType to be an explicit nil
func (o *WritableIPAddressRequest) SetAssignedObjectTypeNil() {
	o.AssignedObjectType.Set(nil)
}

// UnsetAssignedObjectType ensures that no value is present for AssignedObjectType, not even an explicit nil
func (o *WritableIPAddressRequest) UnsetAssignedObjectType() {
	o.AssignedObjectType.Unset()
}

// GetAssignedObjectId returns the AssignedObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIPAddressRequest) GetAssignedObjectId() int64 {
	if o == nil || IsNil(o.AssignedObjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.AssignedObjectId.Get()
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIPAddressRequest) GetAssignedObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedObjectId.Get(), o.AssignedObjectId.IsSet()
}

// HasAssignedObjectId returns a boolean if a field has been set.
func (o *WritableIPAddressRequest) HasAssignedObjectId() bool {
	if o != nil && o.AssignedObjectId.IsSet() {
		return true
	}

	return false
}

// SetAssignedObjectId gets a reference to the given NullableInt64 and assigns it to the AssignedObjectId field.
func (o *WritableIPAddressRequest) SetAssignedObjectId(v int64) {
	o.AssignedObjectId.Set(&v)
}
// SetAssignedObjectIdNil sets the value for AssignedObjectId to be an explicit nil
func (o *WritableIPAddressRequest) SetAssignedObjectIdNil() {
	o.AssignedObjectId.Set(nil)
}

// UnsetAssignedObjectId ensures that no value is present for AssignedObjectId, not even an explicit nil
func (o *WritableIPAddressRequest) UnsetAssignedObjectId() {
	o.AssignedObjectId.Unset()
}

// GetNatInside returns the NatInside field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableIPAddressRequest) GetNatInside() int32 {
	if o == nil || IsNil(o.NatInside.Get()) {
		var ret int32
		return ret
	}
	return *o.NatInside.Get()
}

// GetNatInsideOk returns a tuple with the NatInside field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableIPAddressRequest) GetNatInsideOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NatInside.Get(), o.NatInside.IsSet()
}

// HasNatInside returns a boolean if a field has been set.
func (o *WritableIPAddressRequest) HasNatInside() bool {
	if o != nil && o.NatInside.IsSet() {
		return true
	}

	return false
}

// SetNatInside gets a reference to the given NullableInt32 and assigns it to the NatInside field.
func (o *WritableIPAddressRequest) SetNatInside(v int32) {
	o.NatInside.Set(&v)
}
// SetNatInsideNil sets the value for NatInside to be an explicit nil
func (o *WritableIPAddressRequest) SetNatInsideNil() {
	o.NatInside.Set(nil)
}

// UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
func (o *WritableIPAddressRequest) UnsetNatInside() {
	o.NatInside.Unset()
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *WritableIPAddressRequest) GetDnsName() string {
	if o == nil || IsNil(o.DnsName) {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddressRequest) GetDnsNameOk() (*string, bool) {
	if o == nil || IsNil(o.DnsName) {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *WritableIPAddressRequest) HasDnsName() bool {
	if o != nil && !IsNil(o.DnsName) {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *WritableIPAddressRequest) SetDnsName(v string) {
	o.DnsName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableIPAddressRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddressRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableIPAddressRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableIPAddressRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableIPAddressRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddressRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableIPAddressRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableIPAddressRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableIPAddressRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddressRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableIPAddressRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableIPAddressRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableIPAddressRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableIPAddressRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableIPAddressRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableIPAddressRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableIPAddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableIPAddressRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if o.Vrf.IsSet() {
		toSerialize["vrf"] = o.Vrf.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if o.AssignedObjectType.IsSet() {
		toSerialize["assigned_object_type"] = o.AssignedObjectType.Get()
	}
	if o.AssignedObjectId.IsSet() {
		toSerialize["assigned_object_id"] = o.AssignedObjectId.Get()
	}
	if o.NatInside.IsSet() {
		toSerialize["nat_inside"] = o.NatInside.Get()
	}
	if !IsNil(o.DnsName) {
		toSerialize["dns_name"] = o.DnsName
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

func (o *WritableIPAddressRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
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

	varWritableIPAddressRequest := _WritableIPAddressRequest{}

	err = json.Unmarshal(data, &varWritableIPAddressRequest)

	if err != nil {
		return err
	}

	*o = WritableIPAddressRequest(varWritableIPAddressRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "assigned_object_type")
		delete(additionalProperties, "assigned_object_id")
		delete(additionalProperties, "nat_inside")
		delete(additionalProperties, "dns_name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableIPAddressRequest struct {
	value *WritableIPAddressRequest
	isSet bool
}

func (v NullableWritableIPAddressRequest) Get() *WritableIPAddressRequest {
	return v.value
}

func (v *NullableWritableIPAddressRequest) Set(val *WritableIPAddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableIPAddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableIPAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableIPAddressRequest(val *WritableIPAddressRequest) *NullableWritableIPAddressRequest {
	return &NullableWritableIPAddressRequest{value: val, isSet: true}
}

func (v NullableWritableIPAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableIPAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


