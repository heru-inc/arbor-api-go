/*
arborxr-api-v2

This API provides a RESTful interface to interact with your organization's data.

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arbor-api-go

import (
	"encoding/json"
	"fmt"
)

// checks if the DevicesCheckFingerprintRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicesCheckFingerprintRequest{}

// DevicesCheckFingerprintRequest struct for DevicesCheckFingerprintRequest
type DevicesCheckFingerprintRequest struct {
	Fingerprint string `json:"fingerprint"`
	AdditionalProperties map[string]interface{}
}

type _DevicesCheckFingerprintRequest DevicesCheckFingerprintRequest

// NewDevicesCheckFingerprintRequest instantiates a new DevicesCheckFingerprintRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesCheckFingerprintRequest(fingerprint string) *DevicesCheckFingerprintRequest {
	this := DevicesCheckFingerprintRequest{}
	this.Fingerprint = fingerprint
	return &this
}

// NewDevicesCheckFingerprintRequestWithDefaults instantiates a new DevicesCheckFingerprintRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesCheckFingerprintRequestWithDefaults() *DevicesCheckFingerprintRequest {
	this := DevicesCheckFingerprintRequest{}
	return &this
}

// GetFingerprint returns the Fingerprint field value
func (o *DevicesCheckFingerprintRequest) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *DevicesCheckFingerprintRequest) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value
func (o *DevicesCheckFingerprintRequest) SetFingerprint(v string) {
	o.Fingerprint = v
}

func (o DevicesCheckFingerprintRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicesCheckFingerprintRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fingerprint"] = o.Fingerprint

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DevicesCheckFingerprintRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fingerprint",
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

	varDevicesCheckFingerprintRequest := _DevicesCheckFingerprintRequest{}

	err = json.Unmarshal(data, &varDevicesCheckFingerprintRequest)

	if err != nil {
		return err
	}

	*o = DevicesCheckFingerprintRequest(varDevicesCheckFingerprintRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fingerprint")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevicesCheckFingerprintRequest struct {
	value *DevicesCheckFingerprintRequest
	isSet bool
}

func (v NullableDevicesCheckFingerprintRequest) Get() *DevicesCheckFingerprintRequest {
	return v.value
}

func (v *NullableDevicesCheckFingerprintRequest) Set(val *DevicesCheckFingerprintRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesCheckFingerprintRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesCheckFingerprintRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesCheckFingerprintRequest(val *DevicesCheckFingerprintRequest) *NullableDevicesCheckFingerprintRequest {
	return &NullableDevicesCheckFingerprintRequest{value: val, isSet: true}
}

func (v NullableDevicesCheckFingerprintRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesCheckFingerprintRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


