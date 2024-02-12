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

// checks if the AppsUpdateAppRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppsUpdateAppRequest{}

// AppsUpdateAppRequest struct for AppsUpdateAppRequest
type AppsUpdateAppRequest struct {
	Title NullableString `json:"title,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

// NewAppsUpdateAppRequest instantiates a new AppsUpdateAppRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsUpdateAppRequest() *AppsUpdateAppRequest {
	this := AppsUpdateAppRequest{}
	return &this
}

// NewAppsUpdateAppRequestWithDefaults instantiates a new AppsUpdateAppRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsUpdateAppRequestWithDefaults() *AppsUpdateAppRequest {
	this := AppsUpdateAppRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppsUpdateAppRequest) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppsUpdateAppRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *AppsUpdateAppRequest) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *AppsUpdateAppRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *AppsUpdateAppRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *AppsUpdateAppRequest) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppsUpdateAppRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppsUpdateAppRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AppsUpdateAppRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AppsUpdateAppRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AppsUpdateAppRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AppsUpdateAppRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppsUpdateAppRequest) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppsUpdateAppRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AppsUpdateAppRequest) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AppsUpdateAppRequest) SetTags(v []string) {
	o.Tags = v
}

func (o AppsUpdateAppRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppsUpdateAppRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableAppsUpdateAppRequest struct {
	value *AppsUpdateAppRequest
	isSet bool
}

func (v NullableAppsUpdateAppRequest) Get() *AppsUpdateAppRequest {
	return v.value
}

func (v *NullableAppsUpdateAppRequest) Set(val *AppsUpdateAppRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsUpdateAppRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsUpdateAppRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsUpdateAppRequest(val *AppsUpdateAppRequest) *NullableAppsUpdateAppRequest {
	return &NullableAppsUpdateAppRequest{value: val, isSet: true}
}

func (v NullableAppsUpdateAppRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsUpdateAppRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


