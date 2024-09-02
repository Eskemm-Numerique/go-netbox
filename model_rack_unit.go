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

// checks if the RackUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackUnit{}

// RackUnit A rack unit is an abstraction formed by the set (rack, position, face); it does not exist as a row in the database.
type RackUnit struct {
	Id float64 `json:"id"`
	Name string `json:"name"`
	Face RackUnitFace `json:"face"`
	Device BriefDevice `json:"device"`
	Occupied bool `json:"occupied"`
	Display string `json:"display"`
	AdditionalProperties map[string]interface{}
}

type _RackUnit RackUnit

// NewRackUnit instantiates a new RackUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackUnit(id float64, name string, face RackUnitFace, device BriefDevice, occupied bool, display string) *RackUnit {
	this := RackUnit{}
	this.Id = id
	this.Name = name
	this.Face = face
	this.Device = device
	this.Occupied = occupied
	this.Display = display
	return &this
}

// NewRackUnitWithDefaults instantiates a new RackUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackUnitWithDefaults() *RackUnit {
	this := RackUnit{}
	return &this
}

// GetId returns the Id field value
func (o *RackUnit) GetId() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RackUnit) GetIdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RackUnit) SetId(v float64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RackUnit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RackUnit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RackUnit) SetName(v string) {
	o.Name = v
}

// GetFace returns the Face field value
func (o *RackUnit) GetFace() RackUnitFace {
	if o == nil {
		var ret RackUnitFace
		return ret
	}

	return o.Face
}

// GetFaceOk returns a tuple with the Face field value
// and a boolean to check if the value has been set.
func (o *RackUnit) GetFaceOk() (*RackUnitFace, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Face, true
}

// SetFace sets field value
func (o *RackUnit) SetFace(v RackUnitFace) {
	o.Face = v
}

// GetDevice returns the Device field value
func (o *RackUnit) GetDevice() BriefDevice {
	if o == nil {
		var ret BriefDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *RackUnit) GetDeviceOk() (*BriefDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *RackUnit) SetDevice(v BriefDevice) {
	o.Device = v
}

// GetOccupied returns the Occupied field value
func (o *RackUnit) GetOccupied() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Occupied
}

// GetOccupiedOk returns a tuple with the Occupied field value
// and a boolean to check if the value has been set.
func (o *RackUnit) GetOccupiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Occupied, true
}

// SetOccupied sets field value
func (o *RackUnit) SetOccupied(v bool) {
	o.Occupied = v
}

// GetDisplay returns the Display field value
func (o *RackUnit) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *RackUnit) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *RackUnit) SetDisplay(v string) {
	o.Display = v
}

func (o RackUnit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackUnit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["face"] = o.Face
	toSerialize["device"] = o.Device
	toSerialize["occupied"] = o.Occupied
	toSerialize["display"] = o.Display

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RackUnit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"face",
		"device",
		"occupied",
		"display",
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

	varRackUnit := _RackUnit{}

	err = json.Unmarshal(data, &varRackUnit)

	if err != nil {
		return err
	}

	*o = RackUnit(varRackUnit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "face")
		delete(additionalProperties, "device")
		delete(additionalProperties, "occupied")
		delete(additionalProperties, "display")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRackUnit struct {
	value *RackUnit
	isSet bool
}

func (v NullableRackUnit) Get() *RackUnit {
	return v.value
}

func (v *NullableRackUnit) Set(val *RackUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableRackUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableRackUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackUnit(val *RackUnit) *NullableRackUnit {
	return &NullableRackUnit{value: val, isSet: true}
}

func (v NullableRackUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


