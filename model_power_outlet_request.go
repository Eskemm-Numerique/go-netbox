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

// checks if the PowerOutletRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerOutletRequest{}

// PowerOutletRequest Adds support for custom fields and tags.
type PowerOutletRequest struct {
	Device BriefDeviceRequest `json:"device"`
	Module NullableBriefModuleRequest `json:"module,omitempty"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	// * `iec-60320-c5` - C5 * `iec-60320-c7` - C7 * `iec-60320-c13` - C13 * `iec-60320-c15` - C15 * `iec-60320-c19` - C19 * `iec-60320-c21` - C21 * `iec-60309-p-n-e-4h` - P+N+E 4H * `iec-60309-p-n-e-6h` - P+N+E 6H * `iec-60309-p-n-e-9h` - P+N+E 9H * `iec-60309-2p-e-4h` - 2P+E 4H * `iec-60309-2p-e-6h` - 2P+E 6H * `iec-60309-2p-e-9h` - 2P+E 9H * `iec-60309-3p-e-4h` - 3P+E 4H * `iec-60309-3p-e-6h` - 3P+E 6H * `iec-60309-3p-e-9h` - 3P+E 9H * `iec-60309-3p-n-e-4h` - 3P+N+E 4H * `iec-60309-3p-n-e-6h` - 3P+N+E 6H * `iec-60309-3p-n-e-9h` - 3P+N+E 9H * `iec-60906-1` - IEC 60906-1 * `nbr-14136-10a` - 2P+T 10A (NBR 14136) * `nbr-14136-20a` - 2P+T 20A (NBR 14136) * `nema-1-15r` - NEMA 1-15R * `nema-5-15r` - NEMA 5-15R * `nema-5-20r` - NEMA 5-20R * `nema-5-30r` - NEMA 5-30R * `nema-5-50r` - NEMA 5-50R * `nema-6-15r` - NEMA 6-15R * `nema-6-20r` - NEMA 6-20R * `nema-6-30r` - NEMA 6-30R * `nema-6-50r` - NEMA 6-50R * `nema-10-30r` - NEMA 10-30R * `nema-10-50r` - NEMA 10-50R * `nema-14-20r` - NEMA 14-20R * `nema-14-30r` - NEMA 14-30R * `nema-14-50r` - NEMA 14-50R * `nema-14-60r` - NEMA 14-60R * `nema-15-15r` - NEMA 15-15R * `nema-15-20r` - NEMA 15-20R * `nema-15-30r` - NEMA 15-30R * `nema-15-50r` - NEMA 15-50R * `nema-15-60r` - NEMA 15-60R * `nema-l1-15r` - NEMA L1-15R * `nema-l5-15r` - NEMA L5-15R * `nema-l5-20r` - NEMA L5-20R * `nema-l5-30r` - NEMA L5-30R * `nema-l5-50r` - NEMA L5-50R * `nema-l6-15r` - NEMA L6-15R * `nema-l6-20r` - NEMA L6-20R * `nema-l6-30r` - NEMA L6-30R * `nema-l6-50r` - NEMA L6-50R * `nema-l10-30r` - NEMA L10-30R * `nema-l14-20r` - NEMA L14-20R * `nema-l14-30r` - NEMA L14-30R * `nema-l14-50r` - NEMA L14-50R * `nema-l14-60r` - NEMA L14-60R * `nema-l15-20r` - NEMA L15-20R * `nema-l15-30r` - NEMA L15-30R * `nema-l15-50r` - NEMA L15-50R * `nema-l15-60r` - NEMA L15-60R * `nema-l21-20r` - NEMA L21-20R * `nema-l21-30r` - NEMA L21-30R * `nema-l22-30r` - NEMA L22-30R * `CS6360C` - CS6360C * `CS6364C` - CS6364C * `CS8164C` - CS8164C * `CS8264C` - CS8264C * `CS8364C` - CS8364C * `CS8464C` - CS8464C * `ita-e` - ITA Type E (CEE 7/5) * `ita-f` - ITA Type F (CEE 7/3) * `ita-g` - ITA Type G (BS 1363) * `ita-h` - ITA Type H * `ita-i` - ITA Type I * `ita-j` - ITA Type J * `ita-k` - ITA Type K * `ita-l` - ITA Type L (CEI 23-50) * `ita-m` - ITA Type M (BS 546) * `ita-n` - ITA Type N * `ita-o` - ITA Type O * `ita-multistandard` - ITA Multistandard * `usb-a` - USB Type A * `usb-micro-b` - USB Micro B * `usb-c` - USB Type C * `molex-micro-fit-1x2` - Molex Micro-Fit 1x2 * `molex-micro-fit-2x2` - Molex Micro-Fit 2x2 * `molex-micro-fit-2x4` - Molex Micro-Fit 2x4 * `dc-terminal` - DC Terminal * `hdot-cx` - HDOT Cx * `saf-d-grid` - Saf-D-Grid * `neutrik-powercon-20a` - Neutrik powerCON (20A) * `neutrik-powercon-32a` - Neutrik powerCON (32A) * `neutrik-powercon-true1` - Neutrik powerCON TRUE1 * `neutrik-powercon-true1-top` - Neutrik powerCON TRUE1 TOP * `ubiquiti-smartpower` - Ubiquiti SmartPower * `hardwired` - Hardwired * `other` - Other
	Type NullableString `json:"type,omitempty"`
	PowerPort NullableBriefPowerPortRequest `json:"power_port,omitempty"`
	// * `A` - A * `B` - B * `C` - C
	FeedLeg NullableString `json:"feed_leg,omitempty"`
	Description *string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected *bool `json:"mark_connected,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PowerOutletRequest PowerOutletRequest

// NewPowerOutletRequest instantiates a new PowerOutletRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerOutletRequest(device BriefDeviceRequest, name string) *PowerOutletRequest {
	this := PowerOutletRequest{}
	this.Device = device
	this.Name = name
	return &this
}

// NewPowerOutletRequestWithDefaults instantiates a new PowerOutletRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerOutletRequestWithDefaults() *PowerOutletRequest {
	this := PowerOutletRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *PowerOutletRequest) GetDevice() BriefDeviceRequest {
	if o == nil {
		var ret BriefDeviceRequest
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *PowerOutletRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *PowerOutletRequest) SetDevice(v BriefDeviceRequest) {
	o.Device = v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOutletRequest) GetModule() BriefModuleRequest {
	if o == nil || IsNil(o.Module.Get()) {
		var ret BriefModuleRequest
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutletRequest) GetModuleOk() (*BriefModuleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *PowerOutletRequest) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableBriefModuleRequest and assigns it to the Module field.
func (o *PowerOutletRequest) SetModule(v BriefModuleRequest) {
	o.Module.Set(&v)
}
// SetModuleNil sets the value for Module to be an explicit nil
func (o *PowerOutletRequest) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *PowerOutletRequest) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value
func (o *PowerOutletRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PowerOutletRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PowerOutletRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PowerOutletRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PowerOutletRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PowerOutletRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOutletRequest) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutletRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *PowerOutletRequest) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *PowerOutletRequest) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *PowerOutletRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *PowerOutletRequest) UnsetType() {
	o.Type.Unset()
}

// GetPowerPort returns the PowerPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOutletRequest) GetPowerPort() BriefPowerPortRequest {
	if o == nil || IsNil(o.PowerPort.Get()) {
		var ret BriefPowerPortRequest
		return ret
	}
	return *o.PowerPort.Get()
}

// GetPowerPortOk returns a tuple with the PowerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutletRequest) GetPowerPortOk() (*BriefPowerPortRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerPort.Get(), o.PowerPort.IsSet()
}

// HasPowerPort returns a boolean if a field has been set.
func (o *PowerOutletRequest) HasPowerPort() bool {
	if o != nil && o.PowerPort.IsSet() {
		return true
	}

	return false
}

// SetPowerPort gets a reference to the given NullableBriefPowerPortRequest and assigns it to the PowerPort field.
func (o *PowerOutletRequest) SetPowerPort(v BriefPowerPortRequest) {
	o.PowerPort.Set(&v)
}
// SetPowerPortNil sets the value for PowerPort to be an explicit nil
func (o *PowerOutletRequest) SetPowerPortNil() {
	o.PowerPort.Set(nil)
}

// UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
func (o *PowerOutletRequest) UnsetPowerPort() {
	o.PowerPort.Unset()
}

// GetFeedLeg returns the FeedLeg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerOutletRequest) GetFeedLeg() string {
	if o == nil || IsNil(o.FeedLeg.Get()) {
		var ret string
		return ret
	}
	return *o.FeedLeg.Get()
}

// GetFeedLegOk returns a tuple with the FeedLeg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerOutletRequest) GetFeedLegOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeedLeg.Get(), o.FeedLeg.IsSet()
}

// HasFeedLeg returns a boolean if a field has been set.
func (o *PowerOutletRequest) HasFeedLeg() bool {
	if o != nil && o.FeedLeg.IsSet() {
		return true
	}

	return false
}

// SetFeedLeg gets a reference to the given NullableString and assigns it to the FeedLeg field.
func (o *PowerOutletRequest) SetFeedLeg(v string) {
	o.FeedLeg.Set(&v)
}
// SetFeedLegNil sets the value for FeedLeg to be an explicit nil
func (o *PowerOutletRequest) SetFeedLegNil() {
	o.FeedLeg.Set(nil)
}

// UnsetFeedLeg ensures that no value is present for FeedLeg, not even an explicit nil
func (o *PowerOutletRequest) UnsetFeedLeg() {
	o.FeedLeg.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PowerOutletRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PowerOutletRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PowerOutletRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *PowerOutletRequest) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletRequest) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *PowerOutletRequest) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *PowerOutletRequest) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PowerOutletRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PowerOutletRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PowerOutletRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PowerOutletRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PowerOutletRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PowerOutletRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PowerOutletRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerOutletRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.PowerPort.IsSet() {
		toSerialize["power_port"] = o.PowerPort.Get()
	}
	if o.FeedLeg.IsSet() {
		toSerialize["feed_leg"] = o.FeedLeg.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
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

func (o *PowerOutletRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device",
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

	varPowerOutletRequest := _PowerOutletRequest{}

	err = json.Unmarshal(data, &varPowerOutletRequest)

	if err != nil {
		return err
	}

	*o = PowerOutletRequest(varPowerOutletRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "power_port")
		delete(additionalProperties, "feed_leg")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerOutletRequest struct {
	value *PowerOutletRequest
	isSet bool
}

func (v NullablePowerOutletRequest) Get() *PowerOutletRequest {
	return v.value
}

func (v *NullablePowerOutletRequest) Set(val *PowerOutletRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerOutletRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerOutletRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerOutletRequest(val *PowerOutletRequest) *NullablePowerOutletRequest {
	return &NullablePowerOutletRequest{value: val, isSet: true}
}

func (v NullablePowerOutletRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerOutletRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


