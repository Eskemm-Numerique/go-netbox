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

// ClusterStatusValue * `planned` - Planned * `staging` - Staging * `active` - Active * `decommissioning` - Decommissioning * `offline` - Offline
type ClusterStatusValue string

// List of Cluster_status_value
const (
	CLUSTERSTATUSVALUE_PLANNED ClusterStatusValue = "planned"
	CLUSTERSTATUSVALUE_STAGING ClusterStatusValue = "staging"
	CLUSTERSTATUSVALUE_ACTIVE ClusterStatusValue = "active"
	CLUSTERSTATUSVALUE_DECOMMISSIONING ClusterStatusValue = "decommissioning"
	CLUSTERSTATUSVALUE_OFFLINE ClusterStatusValue = "offline"
)

// All allowed values of ClusterStatusValue enum
var AllowedClusterStatusValueEnumValues = []ClusterStatusValue{
	"planned",
	"staging",
	"active",
	"decommissioning",
	"offline",
}

func (v *ClusterStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterStatusValue(value)
	for _, existing := range AllowedClusterStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterStatusValue", value)
}

// NewClusterStatusValueFromValue returns a pointer to a valid ClusterStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterStatusValueFromValue(v string) (*ClusterStatusValue, error) {
	ev := ClusterStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterStatusValue: valid values are %v", v, AllowedClusterStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterStatusValue) IsValid() bool {
	for _, existing := range AllowedClusterStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Cluster_status_value value
func (v ClusterStatusValue) Ptr() *ClusterStatusValue {
	return &v
}

type NullableClusterStatusValue struct {
	value *ClusterStatusValue
	isSet bool
}

func (v NullableClusterStatusValue) Get() *ClusterStatusValue {
	return v.value
}

func (v *NullableClusterStatusValue) Set(val *ClusterStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStatusValue(val *ClusterStatusValue) *NullableClusterStatusValue {
	return &NullableClusterStatusValue{value: val, isSet: true}
}

func (v NullableClusterStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

