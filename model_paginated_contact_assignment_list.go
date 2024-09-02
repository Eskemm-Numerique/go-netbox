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

// checks if the PaginatedContactAssignmentList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedContactAssignmentList{}

// PaginatedContactAssignmentList struct for PaginatedContactAssignmentList
type PaginatedContactAssignmentList struct {
	Count int32 `json:"count"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []ContactAssignment `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedContactAssignmentList PaginatedContactAssignmentList

// NewPaginatedContactAssignmentList instantiates a new PaginatedContactAssignmentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedContactAssignmentList(count int32, results []ContactAssignment) *PaginatedContactAssignmentList {
	this := PaginatedContactAssignmentList{}
	this.Count = count
	this.Results = results
	return &this
}

// NewPaginatedContactAssignmentListWithDefaults instantiates a new PaginatedContactAssignmentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedContactAssignmentListWithDefaults() *PaginatedContactAssignmentList {
	this := PaginatedContactAssignmentList{}
	return &this
}

// GetCount returns the Count field value
func (o *PaginatedContactAssignmentList) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PaginatedContactAssignmentList) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PaginatedContactAssignmentList) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedContactAssignmentList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedContactAssignmentList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedContactAssignmentList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedContactAssignmentList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedContactAssignmentList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedContactAssignmentList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedContactAssignmentList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedContactAssignmentList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedContactAssignmentList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedContactAssignmentList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedContactAssignmentList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedContactAssignmentList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value
func (o *PaginatedContactAssignmentList) GetResults() []ContactAssignment {
	if o == nil {
		var ret []ContactAssignment
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedContactAssignmentList) GetResultsOk() ([]ContactAssignment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedContactAssignmentList) SetResults(v []ContactAssignment) {
	o.Results = v
}

func (o PaginatedContactAssignmentList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedContactAssignmentList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedContactAssignmentList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"results",
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

	varPaginatedContactAssignmentList := _PaginatedContactAssignmentList{}

	err = json.Unmarshal(data, &varPaginatedContactAssignmentList)

	if err != nil {
		return err
	}

	*o = PaginatedContactAssignmentList(varPaginatedContactAssignmentList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "next")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedContactAssignmentList struct {
	value *PaginatedContactAssignmentList
	isSet bool
}

func (v NullablePaginatedContactAssignmentList) Get() *PaginatedContactAssignmentList {
	return v.value
}

func (v *NullablePaginatedContactAssignmentList) Set(val *PaginatedContactAssignmentList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedContactAssignmentList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedContactAssignmentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedContactAssignmentList(val *PaginatedContactAssignmentList) *NullablePaginatedContactAssignmentList {
	return &NullablePaginatedContactAssignmentList{value: val, isSet: true}
}

func (v NullablePaginatedContactAssignmentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedContactAssignmentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


