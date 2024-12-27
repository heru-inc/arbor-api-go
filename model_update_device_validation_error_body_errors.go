/*
arborxr-api-v2

This API provides a RESTful interface to interact with your organization's data.

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arborapi

import (
	"encoding/json"
)

// checks if the UpdateDeviceValidationErrorBodyErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceValidationErrorBodyErrors{}

// UpdateDeviceValidationErrorBodyErrors struct for UpdateDeviceValidationErrorBodyErrors
type UpdateDeviceValidationErrorBodyErrors struct {
	Title []string `json:"title,omitempty"`
	DeviceGroupId []string `json:"deviceGroupId,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Tags0 []string `json:"tags.0,omitempty"`
	Tags1 []string `json:"tags.1,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateDeviceValidationErrorBodyErrors UpdateDeviceValidationErrorBodyErrors

// NewUpdateDeviceValidationErrorBodyErrors instantiates a new UpdateDeviceValidationErrorBodyErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceValidationErrorBodyErrors() *UpdateDeviceValidationErrorBodyErrors {
	this := UpdateDeviceValidationErrorBodyErrors{}
	return &this
}

// NewUpdateDeviceValidationErrorBodyErrorsWithDefaults instantiates a new UpdateDeviceValidationErrorBodyErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceValidationErrorBodyErrorsWithDefaults() *UpdateDeviceValidationErrorBodyErrors {
	this := UpdateDeviceValidationErrorBodyErrors{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDeviceValidationErrorBodyErrors) GetTitle() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDeviceValidationErrorBodyErrors) GetTitleOk() ([]string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateDeviceValidationErrorBodyErrors) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given []string and assigns it to the Title field.
func (o *UpdateDeviceValidationErrorBodyErrors) SetTitle(v []string) {
	o.Title = v
}

// GetDeviceGroupId returns the DeviceGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDeviceValidationErrorBodyErrors) GetDeviceGroupId() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DeviceGroupId
}

// GetDeviceGroupIdOk returns a tuple with the DeviceGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDeviceValidationErrorBodyErrors) GetDeviceGroupIdOk() ([]string, bool) {
	if o == nil || IsNil(o.DeviceGroupId) {
		return nil, false
	}
	return o.DeviceGroupId, true
}

// HasDeviceGroupId returns a boolean if a field has been set.
func (o *UpdateDeviceValidationErrorBodyErrors) HasDeviceGroupId() bool {
	if o != nil && !IsNil(o.DeviceGroupId) {
		return true
	}

	return false
}

// SetDeviceGroupId gets a reference to the given []string and assigns it to the DeviceGroupId field.
func (o *UpdateDeviceValidationErrorBodyErrors) SetDeviceGroupId(v []string) {
	o.DeviceGroupId = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDeviceValidationErrorBodyErrors) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDeviceValidationErrorBodyErrors) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateDeviceValidationErrorBodyErrors) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateDeviceValidationErrorBodyErrors) SetTags(v []string) {
	o.Tags = v
}

// GetTags0 returns the Tags0 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDeviceValidationErrorBodyErrors) GetTags0() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags0
}

// GetTags0Ok returns a tuple with the Tags0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDeviceValidationErrorBodyErrors) GetTags0Ok() ([]string, bool) {
	if o == nil || IsNil(o.Tags0) {
		return nil, false
	}
	return o.Tags0, true
}

// HasTags0 returns a boolean if a field has been set.
func (o *UpdateDeviceValidationErrorBodyErrors) HasTags0() bool {
	if o != nil && !IsNil(o.Tags0) {
		return true
	}

	return false
}

// SetTags0 gets a reference to the given []string and assigns it to the Tags0 field.
func (o *UpdateDeviceValidationErrorBodyErrors) SetTags0(v []string) {
	o.Tags0 = v
}

// GetTags1 returns the Tags1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDeviceValidationErrorBodyErrors) GetTags1() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags1
}

// GetTags1Ok returns a tuple with the Tags1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDeviceValidationErrorBodyErrors) GetTags1Ok() ([]string, bool) {
	if o == nil || IsNil(o.Tags1) {
		return nil, false
	}
	return o.Tags1, true
}

// HasTags1 returns a boolean if a field has been set.
func (o *UpdateDeviceValidationErrorBodyErrors) HasTags1() bool {
	if o != nil && !IsNil(o.Tags1) {
		return true
	}

	return false
}

// SetTags1 gets a reference to the given []string and assigns it to the Tags1 field.
func (o *UpdateDeviceValidationErrorBodyErrors) SetTags1(v []string) {
	o.Tags1 = v
}

func (o UpdateDeviceValidationErrorBodyErrors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceValidationErrorBodyErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.DeviceGroupId != nil {
		toSerialize["deviceGroupId"] = o.DeviceGroupId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Tags0 != nil {
		toSerialize["tags.0"] = o.Tags0
	}
	if o.Tags1 != nil {
		toSerialize["tags.1"] = o.Tags1
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateDeviceValidationErrorBodyErrors) UnmarshalJSON(data []byte) (err error) {
	varUpdateDeviceValidationErrorBodyErrors := _UpdateDeviceValidationErrorBodyErrors{}

	err = json.Unmarshal(data, &varUpdateDeviceValidationErrorBodyErrors)

	if err != nil {
		return err
	}

	*o = UpdateDeviceValidationErrorBodyErrors(varUpdateDeviceValidationErrorBodyErrors)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "deviceGroupId")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "tags.0")
		delete(additionalProperties, "tags.1")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateDeviceValidationErrorBodyErrors struct {
	value *UpdateDeviceValidationErrorBodyErrors
	isSet bool
}

func (v NullableUpdateDeviceValidationErrorBodyErrors) Get() *UpdateDeviceValidationErrorBodyErrors {
	return v.value
}

func (v *NullableUpdateDeviceValidationErrorBodyErrors) Set(val *UpdateDeviceValidationErrorBodyErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceValidationErrorBodyErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceValidationErrorBodyErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceValidationErrorBodyErrors(val *UpdateDeviceValidationErrorBodyErrors) *NullableUpdateDeviceValidationErrorBodyErrors {
	return &NullableUpdateDeviceValidationErrorBodyErrors{value: val, isSet: true}
}

func (v NullableUpdateDeviceValidationErrorBodyErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceValidationErrorBodyErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


