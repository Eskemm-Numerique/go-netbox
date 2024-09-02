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

// checks if the RearPortTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RearPortTemplateRequest{}

// RearPortTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type RearPortTemplateRequest struct {
	DeviceType NullableBriefDeviceTypeRequest `json:"device_type,omitempty"`
	ModuleType NullableBriefModuleTypeRequest `json:"module_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	// * `8p8c` - 8P8C * `8p6c` - 8P6C * `8p4c` - 8P4C * `8p2c` - 8P2C * `6p6c` - 6P6C * `6p4c` - 6P4C * `6p2c` - 6P2C * `4p4c` - 4P4C * `4p2c` - 4P2C * `gg45` - GG45 * `tera-4p` - TERA 4P * `tera-2p` - TERA 2P * `tera-1p` - TERA 1P * `110-punch` - 110 Punch * `bnc` - BNC * `f` - F Connector * `n` - N Connector * `mrj21` - MRJ21 * `fc` - FC * `lc` - LC * `lc-pc` - LC/PC * `lc-upc` - LC/UPC * `lc-apc` - LC/APC * `lsh` - LSH * `lsh-pc` - LSH/PC * `lsh-upc` - LSH/UPC * `lsh-apc` - LSH/APC * `lx5` - LX.5 * `lx5-pc` - LX.5/PC * `lx5-upc` - LX.5/UPC * `lx5-apc` - LX.5/APC * `mpo` - MPO * `mtrj` - MTRJ * `sc` - SC * `sc-pc` - SC/PC * `sc-upc` - SC/UPC * `sc-apc` - SC/APC * `st` - ST * `cs` - CS * `sn` - SN * `sma-905` - SMA 905 * `sma-906` - SMA 906 * `urm-p2` - URM-P2 * `urm-p4` - URM-P4 * `urm-p8` - URM-P8 * `splice` - Splice * `other` - Other
	Type string `json:"type"`
	Color *string `json:"color,omitempty" validate:"regexp=^[0-9a-f]{6}$"`
	Positions *int32 `json:"positions,omitempty"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RearPortTemplateRequest RearPortTemplateRequest

// NewRearPortTemplateRequest instantiates a new RearPortTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRearPortTemplateRequest(name string, type_ string) *RearPortTemplateRequest {
	this := RearPortTemplateRequest{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewRearPortTemplateRequestWithDefaults instantiates a new RearPortTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRearPortTemplateRequestWithDefaults() *RearPortTemplateRequest {
	this := RearPortTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RearPortTemplateRequest) GetDeviceType() BriefDeviceTypeRequest {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret BriefDeviceTypeRequest
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RearPortTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *RearPortTemplateRequest) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableBriefDeviceTypeRequest and assigns it to the DeviceType field.
func (o *RearPortTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest) {
	o.DeviceType.Set(&v)
}
// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *RearPortTemplateRequest) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *RearPortTemplateRequest) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RearPortTemplateRequest) GetModuleType() BriefModuleTypeRequest {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret BriefModuleTypeRequest
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RearPortTemplateRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *RearPortTemplateRequest) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableBriefModuleTypeRequest and assigns it to the ModuleType field.
func (o *RearPortTemplateRequest) SetModuleType(v BriefModuleTypeRequest) {
	o.ModuleType.Set(&v)
}
// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *RearPortTemplateRequest) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *RearPortTemplateRequest) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value
func (o *RearPortTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RearPortTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RearPortTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RearPortTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RearPortTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RearPortTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *RearPortTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value
func (o *RearPortTemplateRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RearPortTemplateRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RearPortTemplateRequest) SetType(v string) {
	o.Type = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *RearPortTemplateRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RearPortTemplateRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *RearPortTemplateRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *RearPortTemplateRequest) SetColor(v string) {
	o.Color = &v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *RearPortTemplateRequest) GetPositions() int32 {
	if o == nil || IsNil(o.Positions) {
		var ret int32
		return ret
	}
	return *o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RearPortTemplateRequest) GetPositionsOk() (*int32, bool) {
	if o == nil || IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *RearPortTemplateRequest) HasPositions() bool {
	if o != nil && !IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given int32 and assigns it to the Positions field.
func (o *RearPortTemplateRequest) SetPositions(v int32) {
	o.Positions = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RearPortTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RearPortTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RearPortTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RearPortTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o RearPortTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RearPortTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RearPortTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
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

	varRearPortTemplateRequest := _RearPortTemplateRequest{}

	err = json.Unmarshal(data, &varRearPortTemplateRequest)

	if err != nil {
		return err
	}

	*o = RearPortTemplateRequest(varRearPortTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "color")
		delete(additionalProperties, "positions")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRearPortTemplateRequest struct {
	value *RearPortTemplateRequest
	isSet bool
}

func (v NullableRearPortTemplateRequest) Get() *RearPortTemplateRequest {
	return v.value
}

func (v *NullableRearPortTemplateRequest) Set(val *RearPortTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRearPortTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRearPortTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRearPortTemplateRequest(val *RearPortTemplateRequest) *NullableRearPortTemplateRequest {
	return &NullableRearPortTemplateRequest{value: val, isSet: true}
}

func (v NullableRearPortTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRearPortTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


