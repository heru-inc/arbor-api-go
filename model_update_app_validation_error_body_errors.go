/*
arborxr-api-v2

This API provides a RESTful interface to interact with your organization's data.

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateAppValidationErrorBodyErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAppValidationErrorBodyErrors{}

// UpdateAppValidationErrorBodyErrors struct for UpdateAppValidationErrorBodyErrors
type UpdateAppValidationErrorBodyErrors struct {
	Title []string `json:"title,omitempty"`
	Description []string `json:"description,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Tags0 []string `json:"tags.0,omitempty"`
	Tags1 []string `json:"tags.1,omitempty"`
}

// NewUpdateAppValidationErrorBodyErrors instantiates a new UpdateAppValidationErrorBodyErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAppValidationErrorBodyErrors() *UpdateAppValidationErrorBodyErrors {
	this := UpdateAppValidationErrorBodyErrors{}
	return &this
}

// NewUpdateAppValidationErrorBodyErrorsWithDefaults instantiates a new UpdateAppValidationErrorBodyErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAppValidationErrorBodyErrorsWithDefaults() *UpdateAppValidationErrorBodyErrors {
	this := UpdateAppValidationErrorBodyErrors{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAppValidationErrorBodyErrors) GetTitle() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAppValidationErrorBodyErrors) GetTitleOk() ([]string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateAppValidationErrorBodyErrors) HasTitle() bool {
	if o != nil && IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given []string and assigns it to the Title field.
func (o *UpdateAppValidationErrorBodyErrors) SetTitle(v []string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAppValidationErrorBodyErrors) GetDescription() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAppValidationErrorBodyErrors) GetDescriptionOk() ([]string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateAppValidationErrorBodyErrors) HasDescription() bool {
	if o != nil && IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []string and assigns it to the Description field.
func (o *UpdateAppValidationErrorBodyErrors) SetDescription(v []string) {
	o.Description = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAppValidationErrorBodyErrors) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAppValidationErrorBodyErrors) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateAppValidationErrorBodyErrors) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateAppValidationErrorBodyErrors) SetTags(v []string) {
	o.Tags = v
}

// GetTags0 returns the Tags0 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAppValidationErrorBodyErrors) GetTags0() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags0
}

// GetTags0Ok returns a tuple with the Tags0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAppValidationErrorBodyErrors) GetTags0Ok() ([]string, bool) {
	if o == nil || IsNil(o.Tags0) {
		return nil, false
	}
	return o.Tags0, true
}

// HasTags0 returns a boolean if a field has been set.
func (o *UpdateAppValidationErrorBodyErrors) HasTags0() bool {
	if o != nil && IsNil(o.Tags0) {
		return true
	}

	return false
}

// SetTags0 gets a reference to the given []string and assigns it to the Tags0 field.
func (o *UpdateAppValidationErrorBodyErrors) SetTags0(v []string) {
	o.Tags0 = v
}

// GetTags1 returns the Tags1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAppValidationErrorBodyErrors) GetTags1() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags1
}

// GetTags1Ok returns a tuple with the Tags1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAppValidationErrorBodyErrors) GetTags1Ok() ([]string, bool) {
	if o == nil || IsNil(o.Tags1) {
		return nil, false
	}
	return o.Tags1, true
}

// HasTags1 returns a boolean if a field has been set.
func (o *UpdateAppValidationErrorBodyErrors) HasTags1() bool {
	if o != nil && IsNil(o.Tags1) {
		return true
	}

	return false
}

// SetTags1 gets a reference to the given []string and assigns it to the Tags1 field.
func (o *UpdateAppValidationErrorBodyErrors) SetTags1(v []string) {
	o.Tags1 = v
}

func (o UpdateAppValidationErrorBodyErrors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAppValidationErrorBodyErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
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
	return toSerialize, nil
}

type NullableUpdateAppValidationErrorBodyErrors struct {
	value *UpdateAppValidationErrorBodyErrors
	isSet bool
}

func (v NullableUpdateAppValidationErrorBodyErrors) Get() *UpdateAppValidationErrorBodyErrors {
	return v.value
}

func (v *NullableUpdateAppValidationErrorBodyErrors) Set(val *UpdateAppValidationErrorBodyErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAppValidationErrorBodyErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAppValidationErrorBodyErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAppValidationErrorBodyErrors(val *UpdateAppValidationErrorBodyErrors) *NullableUpdateAppValidationErrorBodyErrors {
	return &NullableUpdateAppValidationErrorBodyErrors{value: val, isSet: true}
}

func (v NullableUpdateAppValidationErrorBodyErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAppValidationErrorBodyErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

