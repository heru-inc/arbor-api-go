/*
arborxr-api-v2

This API provides a RESTful interface to interact with your organization's data.

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arborapi

import (
	"encoding/json"
	"fmt"
)

// checks if the DeviceGroupsAddReleaseChannelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceGroupsAddReleaseChannelRequest{}

// DeviceGroupsAddReleaseChannelRequest struct for DeviceGroupsAddReleaseChannelRequest
type DeviceGroupsAddReleaseChannelRequest struct {
	ReleaseChannelId string `json:"releaseChannelId"`
	AdditionalProperties map[string]interface{}
}

type _DeviceGroupsAddReleaseChannelRequest DeviceGroupsAddReleaseChannelRequest

// NewDeviceGroupsAddReleaseChannelRequest instantiates a new DeviceGroupsAddReleaseChannelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceGroupsAddReleaseChannelRequest(releaseChannelId string) *DeviceGroupsAddReleaseChannelRequest {
	this := DeviceGroupsAddReleaseChannelRequest{}
	this.ReleaseChannelId = releaseChannelId
	return &this
}

// NewDeviceGroupsAddReleaseChannelRequestWithDefaults instantiates a new DeviceGroupsAddReleaseChannelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceGroupsAddReleaseChannelRequestWithDefaults() *DeviceGroupsAddReleaseChannelRequest {
	this := DeviceGroupsAddReleaseChannelRequest{}
	return &this
}

// GetReleaseChannelId returns the ReleaseChannelId field value
func (o *DeviceGroupsAddReleaseChannelRequest) GetReleaseChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseChannelId
}

// GetReleaseChannelIdOk returns a tuple with the ReleaseChannelId field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupsAddReleaseChannelRequest) GetReleaseChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseChannelId, true
}

// SetReleaseChannelId sets field value
func (o *DeviceGroupsAddReleaseChannelRequest) SetReleaseChannelId(v string) {
	o.ReleaseChannelId = v
}

func (o DeviceGroupsAddReleaseChannelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceGroupsAddReleaseChannelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["releaseChannelId"] = o.ReleaseChannelId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceGroupsAddReleaseChannelRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"releaseChannelId",
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

	varDeviceGroupsAddReleaseChannelRequest := _DeviceGroupsAddReleaseChannelRequest{}

	err = json.Unmarshal(data, &varDeviceGroupsAddReleaseChannelRequest)

	if err != nil {
		return err
	}

	*o = DeviceGroupsAddReleaseChannelRequest(varDeviceGroupsAddReleaseChannelRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "releaseChannelId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceGroupsAddReleaseChannelRequest struct {
	value *DeviceGroupsAddReleaseChannelRequest
	isSet bool
}

func (v NullableDeviceGroupsAddReleaseChannelRequest) Get() *DeviceGroupsAddReleaseChannelRequest {
	return v.value
}

func (v *NullableDeviceGroupsAddReleaseChannelRequest) Set(val *DeviceGroupsAddReleaseChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceGroupsAddReleaseChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceGroupsAddReleaseChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceGroupsAddReleaseChannelRequest(val *DeviceGroupsAddReleaseChannelRequest) *NullableDeviceGroupsAddReleaseChannelRequest {
	return &NullableDeviceGroupsAddReleaseChannelRequest{value: val, isSet: true}
}

func (v NullableDeviceGroupsAddReleaseChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceGroupsAddReleaseChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


