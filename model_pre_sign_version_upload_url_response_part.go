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

// checks if the PreSignVersionUploadUrlResponsePart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreSignVersionUploadUrlResponsePart{}

// PreSignVersionUploadUrlResponsePart A presigned-URL/part-number pair for uploading the file.
type PreSignVersionUploadUrlResponsePart struct {
	PartNumber int32 `json:"partNumber"`
	PresignedUrl string `json:"presignedUrl"`
}

type _PreSignVersionUploadUrlResponsePart PreSignVersionUploadUrlResponsePart

// NewPreSignVersionUploadUrlResponsePart instantiates a new PreSignVersionUploadUrlResponsePart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreSignVersionUploadUrlResponsePart(partNumber int32, presignedUrl string) *PreSignVersionUploadUrlResponsePart {
	this := PreSignVersionUploadUrlResponsePart{}
	this.PartNumber = partNumber
	this.PresignedUrl = presignedUrl
	return &this
}

// NewPreSignVersionUploadUrlResponsePartWithDefaults instantiates a new PreSignVersionUploadUrlResponsePart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreSignVersionUploadUrlResponsePartWithDefaults() *PreSignVersionUploadUrlResponsePart {
	this := PreSignVersionUploadUrlResponsePart{}
	return &this
}

// GetPartNumber returns the PartNumber field value
func (o *PreSignVersionUploadUrlResponsePart) GetPartNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value
// and a boolean to check if the value has been set.
func (o *PreSignVersionUploadUrlResponsePart) GetPartNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartNumber, true
}

// SetPartNumber sets field value
func (o *PreSignVersionUploadUrlResponsePart) SetPartNumber(v int32) {
	o.PartNumber = v
}

// GetPresignedUrl returns the PresignedUrl field value
func (o *PreSignVersionUploadUrlResponsePart) GetPresignedUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PresignedUrl
}

// GetPresignedUrlOk returns a tuple with the PresignedUrl field value
// and a boolean to check if the value has been set.
func (o *PreSignVersionUploadUrlResponsePart) GetPresignedUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresignedUrl, true
}

// SetPresignedUrl sets field value
func (o *PreSignVersionUploadUrlResponsePart) SetPresignedUrl(v string) {
	o.PresignedUrl = v
}

func (o PreSignVersionUploadUrlResponsePart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreSignVersionUploadUrlResponsePart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["partNumber"] = o.PartNumber
	toSerialize["presignedUrl"] = o.PresignedUrl
	return toSerialize, nil
}

func (o *PreSignVersionUploadUrlResponsePart) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"partNumber",
		"presignedUrl",
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

	varPreSignVersionUploadUrlResponsePart := _PreSignVersionUploadUrlResponsePart{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPreSignVersionUploadUrlResponsePart)

	if err != nil {
		return err
	}

	*o = PreSignVersionUploadUrlResponsePart(varPreSignVersionUploadUrlResponsePart)

	return err
}

type NullablePreSignVersionUploadUrlResponsePart struct {
	value *PreSignVersionUploadUrlResponsePart
	isSet bool
}

func (v NullablePreSignVersionUploadUrlResponsePart) Get() *PreSignVersionUploadUrlResponsePart {
	return v.value
}

func (v *NullablePreSignVersionUploadUrlResponsePart) Set(val *PreSignVersionUploadUrlResponsePart) {
	v.value = val
	v.isSet = true
}

func (v NullablePreSignVersionUploadUrlResponsePart) IsSet() bool {
	return v.isSet
}

func (v *NullablePreSignVersionUploadUrlResponsePart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreSignVersionUploadUrlResponsePart(val *PreSignVersionUploadUrlResponsePart) *NullablePreSignVersionUploadUrlResponsePart {
	return &NullablePreSignVersionUploadUrlResponsePart{value: val, isSet: true}
}

func (v NullablePreSignVersionUploadUrlResponsePart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreSignVersionUploadUrlResponsePart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


