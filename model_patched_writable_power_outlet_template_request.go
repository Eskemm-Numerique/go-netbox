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

// checks if the PatchedWritablePowerOutletTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritablePowerOutletTemplateRequest{}

// PatchedWritablePowerOutletTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritablePowerOutletTemplateRequest struct {
	DeviceType NullableBriefDeviceTypeRequest `json:"device_type,omitempty"`
	ModuleType NullableBriefModuleTypeRequest `json:"module_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	// * `iec-60320-c5` - C5 * `iec-60320-c7` - C7 * `iec-60320-c13` - C13 * `iec-60320-c15` - C15 * `iec-60320-c19` - C19 * `iec-60320-c21` - C21 * `iec-60309-p-n-e-4h` - P+N+E 4H * `iec-60309-p-n-e-6h` - P+N+E 6H * `iec-60309-p-n-e-9h` - P+N+E 9H * `iec-60309-2p-e-4h` - 2P+E 4H * `iec-60309-2p-e-6h` - 2P+E 6H * `iec-60309-2p-e-9h` - 2P+E 9H * `iec-60309-3p-e-4h` - 3P+E 4H * `iec-60309-3p-e-6h` - 3P+E 6H * `iec-60309-3p-e-9h` - 3P+E 9H * `iec-60309-3p-n-e-4h` - 3P+N+E 4H * `iec-60309-3p-n-e-6h` - 3P+N+E 6H * `iec-60309-3p-n-e-9h` - 3P+N+E 9H * `iec-60906-1` - IEC 60906-1 * `nbr-14136-10a` - 2P+T 10A (NBR 14136) * `nbr-14136-20a` - 2P+T 20A (NBR 14136) * `nema-1-15r` - NEMA 1-15R * `nema-5-15r` - NEMA 5-15R * `nema-5-20r` - NEMA 5-20R * `nema-5-30r` - NEMA 5-30R * `nema-5-50r` - NEMA 5-50R * `nema-6-15r` - NEMA 6-15R * `nema-6-20r` - NEMA 6-20R * `nema-6-30r` - NEMA 6-30R * `nema-6-50r` - NEMA 6-50R * `nema-10-30r` - NEMA 10-30R * `nema-10-50r` - NEMA 10-50R * `nema-14-20r` - NEMA 14-20R * `nema-14-30r` - NEMA 14-30R * `nema-14-50r` - NEMA 14-50R * `nema-14-60r` - NEMA 14-60R * `nema-15-15r` - NEMA 15-15R * `nema-15-20r` - NEMA 15-20R * `nema-15-30r` - NEMA 15-30R * `nema-15-50r` - NEMA 15-50R * `nema-15-60r` - NEMA 15-60R * `nema-l1-15r` - NEMA L1-15R * `nema-l5-15r` - NEMA L5-15R * `nema-l5-20r` - NEMA L5-20R * `nema-l5-30r` - NEMA L5-30R * `nema-l5-50r` - NEMA L5-50R * `nema-l6-15r` - NEMA L6-15R * `nema-l6-20r` - NEMA L6-20R * `nema-l6-30r` - NEMA L6-30R * `nema-l6-50r` - NEMA L6-50R * `nema-l10-30r` - NEMA L10-30R * `nema-l14-20r` - NEMA L14-20R * `nema-l14-30r` - NEMA L14-30R * `nema-l14-50r` - NEMA L14-50R * `nema-l14-60r` - NEMA L14-60R * `nema-l15-20r` - NEMA L15-20R * `nema-l15-30r` - NEMA L15-30R * `nema-l15-50r` - NEMA L15-50R * `nema-l15-60r` - NEMA L15-60R * `nema-l21-20r` - NEMA L21-20R * `nema-l21-30r` - NEMA L21-30R * `nema-l22-30r` - NEMA L22-30R * `CS6360C` - CS6360C * `CS6364C` - CS6364C * `CS8164C` - CS8164C * `CS8264C` - CS8264C * `CS8364C` - CS8364C * `CS8464C` - CS8464C * `ita-e` - ITA Type E (CEE 7/5) * `ita-f` - ITA Type F (CEE 7/3) * `ita-g` - ITA Type G (BS 1363) * `ita-h` - ITA Type H * `ita-i` - ITA Type I * `ita-j` - ITA Type J * `ita-k` - ITA Type K * `ita-l` - ITA Type L (CEI 23-50) * `ita-m` - ITA Type M (BS 546) * `ita-n` - ITA Type N * `ita-o` - ITA Type O * `ita-multistandard` - ITA Multistandard * `usb-a` - USB Type A * `usb-micro-b` - USB Micro B * `usb-c` - USB Type C * `molex-micro-fit-1x2` - Molex Micro-Fit 1x2 * `molex-micro-fit-2x2` - Molex Micro-Fit 2x2 * `molex-micro-fit-2x4` - Molex Micro-Fit 2x4 * `dc-terminal` - DC Terminal * `hdot-cx` - HDOT Cx * `saf-d-grid` - Saf-D-Grid * `neutrik-powercon-20a` - Neutrik powerCON (20A) * `neutrik-powercon-32a` - Neutrik powerCON (32A) * `neutrik-powercon-true1` - Neutrik powerCON TRUE1 * `neutrik-powercon-true1-top` - Neutrik powerCON TRUE1 TOP * `ubiquiti-smartpower` - Ubiquiti SmartPower * `hardwired` - Hardwired * `other` - Other
	Type *string `json:"type,omitempty"`
	PowerPort NullableBriefPowerPortTemplateRequest `json:"power_port,omitempty"`
	// Phase (for three-phase feeds)  * `A` - A * `B` - B * `C` - C
	FeedLeg *string `json:"feed_leg,omitempty"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritablePowerOutletTemplateRequest PatchedWritablePowerOutletTemplateRequest

// NewPatchedWritablePowerOutletTemplateRequest instantiates a new PatchedWritablePowerOutletTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritablePowerOutletTemplateRequest() *PatchedWritablePowerOutletTemplateRequest {
	this := PatchedWritablePowerOutletTemplateRequest{}
	return &this
}

// NewPatchedWritablePowerOutletTemplateRequestWithDefaults instantiates a new PatchedWritablePowerOutletTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritablePowerOutletTemplateRequestWithDefaults() *PatchedWritablePowerOutletTemplateRequest {
	this := PatchedWritablePowerOutletTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerOutletTemplateRequest) GetDeviceType() BriefDeviceTypeRequest {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret BriefDeviceTypeRequest
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerOutletTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletTemplateRequest) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableBriefDeviceTypeRequest and assigns it to the DeviceType field.
func (o *PatchedWritablePowerOutletTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest) {
	o.DeviceType.Set(&v)
}
// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *PatchedWritablePowerOutletTemplateRequest) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *PatchedWritablePowerOutletTemplateRequest) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerOutletTemplateRequest) GetModuleType() BriefModuleTypeRequest {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret BriefModuleTypeRequest
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerOutletTemplateRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletTemplateRequest) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableBriefModuleTypeRequest and assigns it to the ModuleType field.
func (o *PatchedWritablePowerOutletTemplateRequest) SetModuleType(v BriefModuleTypeRequest) {
	o.ModuleType.Set(&v)
}
// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *PatchedWritablePowerOutletTemplateRequest) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *PatchedWritablePowerOutletTemplateRequest) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritablePowerOutletTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritablePowerOutletTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletTemplateRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletTemplateRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletTemplateRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PatchedWritablePowerOutletTemplateRequest) SetType(v string) {
	o.Type = &v
}

// GetPowerPort returns the PowerPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerOutletTemplateRequest) GetPowerPort() BriefPowerPortTemplateRequest {
	if o == nil || IsNil(o.PowerPort.Get()) {
		var ret BriefPowerPortTemplateRequest
		return ret
	}
	return *o.PowerPort.Get()
}

// GetPowerPortOk returns a tuple with the PowerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerOutletTemplateRequest) GetPowerPortOk() (*BriefPowerPortTemplateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerPort.Get(), o.PowerPort.IsSet()
}

// HasPowerPort returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletTemplateRequest) HasPowerPort() bool {
	if o != nil && o.PowerPort.IsSet() {
		return true
	}

	return false
}

// SetPowerPort gets a reference to the given NullableBriefPowerPortTemplateRequest and assigns it to the PowerPort field.
func (o *PatchedWritablePowerOutletTemplateRequest) SetPowerPort(v BriefPowerPortTemplateRequest) {
	o.PowerPort.Set(&v)
}
// SetPowerPortNil sets the value for PowerPort to be an explicit nil
func (o *PatchedWritablePowerOutletTemplateRequest) SetPowerPortNil() {
	o.PowerPort.Set(nil)
}

// UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
func (o *PatchedWritablePowerOutletTemplateRequest) UnsetPowerPort() {
	o.PowerPort.Unset()
}

// GetFeedLeg returns the FeedLeg field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletTemplateRequest) GetFeedLeg() string {
	if o == nil || IsNil(o.FeedLeg) {
		var ret string
		return ret
	}
	return *o.FeedLeg
}

// GetFeedLegOk returns a tuple with the FeedLeg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletTemplateRequest) GetFeedLegOk() (*string, bool) {
	if o == nil || IsNil(o.FeedLeg) {
		return nil, false
	}
	return o.FeedLeg, true
}

// HasFeedLeg returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletTemplateRequest) HasFeedLeg() bool {
	if o != nil && !IsNil(o.FeedLeg) {
		return true
	}

	return false
}

// SetFeedLeg gets a reference to the given string and assigns it to the FeedLeg field.
func (o *PatchedWritablePowerOutletTemplateRequest) SetFeedLeg(v string) {
	o.FeedLeg = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritablePowerOutletTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerOutletTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritablePowerOutletTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritablePowerOutletTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o PatchedWritablePowerOutletTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritablePowerOutletTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.PowerPort.IsSet() {
		toSerialize["power_port"] = o.PowerPort.Get()
	}
	if !IsNil(o.FeedLeg) {
		toSerialize["feed_leg"] = o.FeedLeg
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritablePowerOutletTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritablePowerOutletTemplateRequest := _PatchedWritablePowerOutletTemplateRequest{}

	err = json.Unmarshal(data, &varPatchedWritablePowerOutletTemplateRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritablePowerOutletTemplateRequest(varPatchedWritablePowerOutletTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "power_port")
		delete(additionalProperties, "feed_leg")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritablePowerOutletTemplateRequest struct {
	value *PatchedWritablePowerOutletTemplateRequest
	isSet bool
}

func (v NullablePatchedWritablePowerOutletTemplateRequest) Get() *PatchedWritablePowerOutletTemplateRequest {
	return v.value
}

func (v *NullablePatchedWritablePowerOutletTemplateRequest) Set(val *PatchedWritablePowerOutletTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerOutletTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerOutletTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerOutletTemplateRequest(val *PatchedWritablePowerOutletTemplateRequest) *NullablePatchedWritablePowerOutletTemplateRequest {
	return &NullablePatchedWritablePowerOutletTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerOutletTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerOutletTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


