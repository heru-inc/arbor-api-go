/*
arborxr-api-v2

This API provides a RESTful interface to interact with your organization's data.

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AppsStartVersionUploadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppsStartVersionUploadRequest{}

// AppsStartVersionUploadRequest struct for AppsStartVersionUploadRequest
type AppsStartVersionUploadRequest struct {
	Filename string `json:"filename"`
	ReleaseChannelId NullableString `json:"releaseChannelId,omitempty"`
}

type _AppsStartVersionUploadRequest AppsStartVersionUploadRequest

// NewAppsStartVersionUploadRequest instantiates a new AppsStartVersionUploadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsStartVersionUploadRequest(filename string) *AppsStartVersionUploadRequest {
	this := AppsStartVersionUploadRequest{}
	this.Filename = filename
	return &this
}

// NewAppsStartVersionUploadRequestWithDefaults instantiates a new AppsStartVersionUploadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsStartVersionUploadRequestWithDefaults() *AppsStartVersionUploadRequest {
	this := AppsStartVersionUploadRequest{}
	return &this
}

// GetFilename returns the Filename field value
func (o *AppsStartVersionUploadRequest) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *AppsStartVersionUploadRequest) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *AppsStartVersionUploadRequest) SetFilename(v string) {
	o.Filename = v
}

// GetReleaseChannelId returns the ReleaseChannelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppsStartVersionUploadRequest) GetReleaseChannelId() string {
	if o == nil || IsNil(o.ReleaseChannelId.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseChannelId.Get()
}

// GetReleaseChannelIdOk returns a tuple with the ReleaseChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppsStartVersionUploadRequest) GetReleaseChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseChannelId.Get(), o.ReleaseChannelId.IsSet()
}

// HasReleaseChannelId returns a boolean if a field has been set.
func (o *AppsStartVersionUploadRequest) HasReleaseChannelId() bool {
	if o != nil && o.ReleaseChannelId.IsSet() {
		return true
	}

	return false
}

// SetReleaseChannelId gets a reference to the given NullableString and assigns it to the ReleaseChannelId field.
func (o *AppsStartVersionUploadRequest) SetReleaseChannelId(v string) {
	o.ReleaseChannelId.Set(&v)
}
// SetReleaseChannelIdNil sets the value for ReleaseChannelId to be an explicit nil
func (o *AppsStartVersionUploadRequest) SetReleaseChannelIdNil() {
	o.ReleaseChannelId.Set(nil)
}

// UnsetReleaseChannelId ensures that no value is present for ReleaseChannelId, not even an explicit nil
func (o *AppsStartVersionUploadRequest) UnsetReleaseChannelId() {
	o.ReleaseChannelId.Unset()
}

func (o AppsStartVersionUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppsStartVersionUploadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filename"] = o.Filename
	if o.ReleaseChannelId.IsSet() {
		toSerialize["releaseChannelId"] = o.ReleaseChannelId.Get()
	}
	return toSerialize, nil
}

func (o *AppsStartVersionUploadRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filename",
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

	varAppsStartVersionUploadRequest := _AppsStartVersionUploadRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppsStartVersionUploadRequest)

	if err != nil {
		return err
	}

	*o = AppsStartVersionUploadRequest(varAppsStartVersionUploadRequest)

	return err
}

type NullableAppsStartVersionUploadRequest struct {
	value *AppsStartVersionUploadRequest
	isSet bool
}

func (v NullableAppsStartVersionUploadRequest) Get() *AppsStartVersionUploadRequest {
	return v.value
}

func (v *NullableAppsStartVersionUploadRequest) Set(val *AppsStartVersionUploadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsStartVersionUploadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsStartVersionUploadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsStartVersionUploadRequest(val *AppsStartVersionUploadRequest) *NullableAppsStartVersionUploadRequest {
	return &NullableAppsStartVersionUploadRequest{value: val, isSet: true}
}

func (v NullableAppsStartVersionUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsStartVersionUploadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

