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

// checks if the AppsUpdateReleaseChannelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppsUpdateReleaseChannelRequest{}

// AppsUpdateReleaseChannelRequest struct for AppsUpdateReleaseChannelRequest
type AppsUpdateReleaseChannelRequest struct {
	VersionId string `json:"versionId"`
}

type _AppsUpdateReleaseChannelRequest AppsUpdateReleaseChannelRequest

// NewAppsUpdateReleaseChannelRequest instantiates a new AppsUpdateReleaseChannelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsUpdateReleaseChannelRequest(versionId string) *AppsUpdateReleaseChannelRequest {
	this := AppsUpdateReleaseChannelRequest{}
	this.VersionId = versionId
	return &this
}

// NewAppsUpdateReleaseChannelRequestWithDefaults instantiates a new AppsUpdateReleaseChannelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsUpdateReleaseChannelRequestWithDefaults() *AppsUpdateReleaseChannelRequest {
	this := AppsUpdateReleaseChannelRequest{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *AppsUpdateReleaseChannelRequest) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *AppsUpdateReleaseChannelRequest) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *AppsUpdateReleaseChannelRequest) SetVersionId(v string) {
	o.VersionId = v
}

func (o AppsUpdateReleaseChannelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppsUpdateReleaseChannelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	return toSerialize, nil
}

func (o *AppsUpdateReleaseChannelRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"versionId",
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

	varAppsUpdateReleaseChannelRequest := _AppsUpdateReleaseChannelRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppsUpdateReleaseChannelRequest)

	if err != nil {
		return err
	}

	*o = AppsUpdateReleaseChannelRequest(varAppsUpdateReleaseChannelRequest)

	return err
}

type NullableAppsUpdateReleaseChannelRequest struct {
	value *AppsUpdateReleaseChannelRequest
	isSet bool
}

func (v NullableAppsUpdateReleaseChannelRequest) Get() *AppsUpdateReleaseChannelRequest {
	return v.value
}

func (v *NullableAppsUpdateReleaseChannelRequest) Set(val *AppsUpdateReleaseChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsUpdateReleaseChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsUpdateReleaseChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsUpdateReleaseChannelRequest(val *AppsUpdateReleaseChannelRequest) *NullableAppsUpdateReleaseChannelRequest {
	return &NullableAppsUpdateReleaseChannelRequest{value: val, isSet: true}
}

func (v NullableAppsUpdateReleaseChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsUpdateReleaseChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

