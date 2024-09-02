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

// checks if the NestedInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedInterface{}

// NestedInterface Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedInterface struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Device NestedDevice `json:"device"`
	Name string `json:"name"`
	Cable NullableInt32 `json:"cable,omitempty"`
	Occupied bool `json:"_occupied"`
	AdditionalProperties map[string]interface{}
}

type _NestedInterface NestedInterface

// NewNestedInterface instantiates a new NestedInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedInterface(id int32, url string, display string, device NestedDevice, name string, occupied bool) *NestedInterface {
	this := NestedInterface{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Device = device
	this.Name = name
	this.Occupied = occupied
	return &this
}

// NewNestedInterfaceWithDefaults instantiates a new NestedInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedInterfaceWithDefaults() *NestedInterface {
	this := NestedInterface{}
	return &this
}

// GetId returns the Id field value
func (o *NestedInterface) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedInterface) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedInterface) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedInterface) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedInterface) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedInterface) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedInterface) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedInterface) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedInterface) SetDisplay(v string) {
	o.Display = v
}

// GetDevice returns the Device field value
func (o *NestedInterface) GetDevice() NestedDevice {
	if o == nil {
		var ret NestedDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *NestedInterface) GetDeviceOk() (*NestedDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *NestedInterface) SetDevice(v NestedDevice) {
	o.Device = v
}

// GetName returns the Name field value
func (o *NestedInterface) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedInterface) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedInterface) SetName(v string) {
	o.Name = v
}

// GetCable returns the Cable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedInterface) GetCable() int32 {
	if o == nil || IsNil(o.Cable.Get()) {
		var ret int32
		return ret
	}
	return *o.Cable.Get()
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedInterface) GetCableOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cable.Get(), o.Cable.IsSet()
}

// HasCable returns a boolean if a field has been set.
func (o *NestedInterface) HasCable() bool {
	if o != nil && o.Cable.IsSet() {
		return true
	}

	return false
}

// SetCable gets a reference to the given NullableInt32 and assigns it to the Cable field.
func (o *NestedInterface) SetCable(v int32) {
	o.Cable.Set(&v)
}
// SetCableNil sets the value for Cable to be an explicit nil
func (o *NestedInterface) SetCableNil() {
	o.Cable.Set(nil)
}

// UnsetCable ensures that no value is present for Cable, not even an explicit nil
func (o *NestedInterface) UnsetCable() {
	o.Cable.Unset()
}

// GetOccupied returns the Occupied field value
func (o *NestedInterface) GetOccupied() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Occupied
}

// GetOccupiedOk returns a tuple with the Occupied field value
// and a boolean to check if the value has been set.
func (o *NestedInterface) GetOccupiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Occupied, true
}

// SetOccupied sets field value
func (o *NestedInterface) SetOccupied(v bool) {
	o.Occupied = v
}

func (o NestedInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["device"] = o.Device
	toSerialize["name"] = o.Name
	if o.Cable.IsSet() {
		toSerialize["cable"] = o.Cable.Get()
	}
	toSerialize["_occupied"] = o.Occupied

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"device",
		"name",
		"_occupied",
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

	varNestedInterface := _NestedInterface{}

	err = json.Unmarshal(data, &varNestedInterface)

	if err != nil {
		return err
	}

	*o = NestedInterface(varNestedInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "device")
		delete(additionalProperties, "name")
		delete(additionalProperties, "cable")
		delete(additionalProperties, "_occupied")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedInterface struct {
	value *NestedInterface
	isSet bool
}

func (v NullableNestedInterface) Get() *NestedInterface {
	return v.value
}

func (v *NullableNestedInterface) Set(val *NestedInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedInterface(val *NestedInterface) *NullableNestedInterface {
	return &NullableNestedInterface{value: val, isSet: true}
}

func (v NullableNestedInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


