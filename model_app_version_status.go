/*
arborxr-api-v2

This API provides a RESTful interface to interact with your organization's data.

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AppVersionStatus the model 'AppVersionStatus'
type AppVersionStatus string

// List of AppVersionStatus
const (
	ERROR AppVersionStatus = "ERROR"
	PROCESSING AppVersionStatus = "PROCESSING"
	UPLOADED AppVersionStatus = "UPLOADED"
	PENDING AppVersionStatus = "PENDING"
	AVAILABLE AppVersionStatus = "AVAILABLE"
)

// All allowed values of AppVersionStatus enum
var AllowedAppVersionStatusEnumValues = []AppVersionStatus{
	"ERROR",
	"PROCESSING",
	"UPLOADED",
	"PENDING",
	"AVAILABLE",
}

func (v *AppVersionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppVersionStatus(value)
	for _, existing := range AllowedAppVersionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppVersionStatus", value)
}

// NewAppVersionStatusFromValue returns a pointer to a valid AppVersionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppVersionStatusFromValue(v string) (*AppVersionStatus, error) {
	ev := AppVersionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppVersionStatus: valid values are %v", v, AllowedAppVersionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppVersionStatus) IsValid() bool {
	for _, existing := range AllowedAppVersionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppVersionStatus value
func (v AppVersionStatus) Ptr() *AppVersionStatus {
	return &v
}

type NullableAppVersionStatus struct {
	value *AppVersionStatus
	isSet bool
}

func (v NullableAppVersionStatus) Get() *AppVersionStatus {
	return v.value
}

func (v *NullableAppVersionStatus) Set(val *AppVersionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAppVersionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAppVersionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppVersionStatus(val *AppVersionStatus) *NullableAppVersionStatus {
	return &NullableAppVersionStatus{value: val, isSet: true}
}

func (v NullableAppVersionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppVersionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

