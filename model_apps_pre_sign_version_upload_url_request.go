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

// checks if the AppsPreSignVersionUploadUrlRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppsPreSignVersionUploadUrlRequest{}

// AppsPreSignVersionUploadUrlRequest struct for AppsPreSignVersionUploadUrlRequest
type AppsPreSignVersionUploadUrlRequest struct {
	Key string `json:"key"`
	UploadId string `json:"uploadId"`
	PartNumbers []int32 `json:"partNumbers"`
	AdditionalProperties map[string]interface{}
}

type _AppsPreSignVersionUploadUrlRequest AppsPreSignVersionUploadUrlRequest

// NewAppsPreSignVersionUploadUrlRequest instantiates a new AppsPreSignVersionUploadUrlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsPreSignVersionUploadUrlRequest(key string, uploadId string, partNumbers []int32) *AppsPreSignVersionUploadUrlRequest {
	this := AppsPreSignVersionUploadUrlRequest{}
	this.Key = key
	this.UploadId = uploadId
	this.PartNumbers = partNumbers
	return &this
}

// NewAppsPreSignVersionUploadUrlRequestWithDefaults instantiates a new AppsPreSignVersionUploadUrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsPreSignVersionUploadUrlRequestWithDefaults() *AppsPreSignVersionUploadUrlRequest {
	this := AppsPreSignVersionUploadUrlRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *AppsPreSignVersionUploadUrlRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AppsPreSignVersionUploadUrlRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AppsPreSignVersionUploadUrlRequest) SetKey(v string) {
	o.Key = v
}

// GetUploadId returns the UploadId field value
func (o *AppsPreSignVersionUploadUrlRequest) GetUploadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value
// and a boolean to check if the value has been set.
func (o *AppsPreSignVersionUploadUrlRequest) GetUploadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadId, true
}

// SetUploadId sets field value
func (o *AppsPreSignVersionUploadUrlRequest) SetUploadId(v string) {
	o.UploadId = v
}

// GetPartNumbers returns the PartNumbers field value
func (o *AppsPreSignVersionUploadUrlRequest) GetPartNumbers() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.PartNumbers
}

// GetPartNumbersOk returns a tuple with the PartNumbers field value
// and a boolean to check if the value has been set.
func (o *AppsPreSignVersionUploadUrlRequest) GetPartNumbersOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartNumbers, true
}

// SetPartNumbers sets field value
func (o *AppsPreSignVersionUploadUrlRequest) SetPartNumbers(v []int32) {
	o.PartNumbers = v
}

func (o AppsPreSignVersionUploadUrlRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppsPreSignVersionUploadUrlRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["uploadId"] = o.UploadId
	toSerialize["partNumbers"] = o.PartNumbers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppsPreSignVersionUploadUrlRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"uploadId",
		"partNumbers",
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

	varAppsPreSignVersionUploadUrlRequest := _AppsPreSignVersionUploadUrlRequest{}

	err = json.Unmarshal(data, &varAppsPreSignVersionUploadUrlRequest)

	if err != nil {
		return err
	}

	*o = AppsPreSignVersionUploadUrlRequest(varAppsPreSignVersionUploadUrlRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "uploadId")
		delete(additionalProperties, "partNumbers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppsPreSignVersionUploadUrlRequest struct {
	value *AppsPreSignVersionUploadUrlRequest
	isSet bool
}

func (v NullableAppsPreSignVersionUploadUrlRequest) Get() *AppsPreSignVersionUploadUrlRequest {
	return v.value
}

func (v *NullableAppsPreSignVersionUploadUrlRequest) Set(val *AppsPreSignVersionUploadUrlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsPreSignVersionUploadUrlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsPreSignVersionUploadUrlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsPreSignVersionUploadUrlRequest(val *AppsPreSignVersionUploadUrlRequest) *NullableAppsPreSignVersionUploadUrlRequest {
	return &NullableAppsPreSignVersionUploadUrlRequest{value: val, isSet: true}
}

func (v NullableAppsPreSignVersionUploadUrlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsPreSignVersionUploadUrlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


