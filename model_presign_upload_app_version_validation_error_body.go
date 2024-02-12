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

// checks if the PresignUploadAppVersionValidationErrorBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PresignUploadAppVersionValidationErrorBody{}

// PresignUploadAppVersionValidationErrorBody struct for PresignUploadAppVersionValidationErrorBody
type PresignUploadAppVersionValidationErrorBody struct {
	Message string `json:"message"`
	Errors PresignUploadAppVersionValidationErrorBodyErrors `json:"errors"`
}

type _PresignUploadAppVersionValidationErrorBody PresignUploadAppVersionValidationErrorBody

// NewPresignUploadAppVersionValidationErrorBody instantiates a new PresignUploadAppVersionValidationErrorBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresignUploadAppVersionValidationErrorBody(message string, errors PresignUploadAppVersionValidationErrorBodyErrors) *PresignUploadAppVersionValidationErrorBody {
	this := PresignUploadAppVersionValidationErrorBody{}
	this.Message = message
	this.Errors = errors
	return &this
}

// NewPresignUploadAppVersionValidationErrorBodyWithDefaults instantiates a new PresignUploadAppVersionValidationErrorBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresignUploadAppVersionValidationErrorBodyWithDefaults() *PresignUploadAppVersionValidationErrorBody {
	this := PresignUploadAppVersionValidationErrorBody{}
	return &this
}

// GetMessage returns the Message field value
func (o *PresignUploadAppVersionValidationErrorBody) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PresignUploadAppVersionValidationErrorBody) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PresignUploadAppVersionValidationErrorBody) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value
func (o *PresignUploadAppVersionValidationErrorBody) GetErrors() PresignUploadAppVersionValidationErrorBodyErrors {
	if o == nil {
		var ret PresignUploadAppVersionValidationErrorBodyErrors
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *PresignUploadAppVersionValidationErrorBody) GetErrorsOk() (*PresignUploadAppVersionValidationErrorBodyErrors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *PresignUploadAppVersionValidationErrorBody) SetErrors(v PresignUploadAppVersionValidationErrorBodyErrors) {
	o.Errors = v
}

func (o PresignUploadAppVersionValidationErrorBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PresignUploadAppVersionValidationErrorBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

func (o *PresignUploadAppVersionValidationErrorBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"errors",
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

	varPresignUploadAppVersionValidationErrorBody := _PresignUploadAppVersionValidationErrorBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPresignUploadAppVersionValidationErrorBody)

	if err != nil {
		return err
	}

	*o = PresignUploadAppVersionValidationErrorBody(varPresignUploadAppVersionValidationErrorBody)

	return err
}

type NullablePresignUploadAppVersionValidationErrorBody struct {
	value *PresignUploadAppVersionValidationErrorBody
	isSet bool
}

func (v NullablePresignUploadAppVersionValidationErrorBody) Get() *PresignUploadAppVersionValidationErrorBody {
	return v.value
}

func (v *NullablePresignUploadAppVersionValidationErrorBody) Set(val *PresignUploadAppVersionValidationErrorBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePresignUploadAppVersionValidationErrorBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePresignUploadAppVersionValidationErrorBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresignUploadAppVersionValidationErrorBody(val *PresignUploadAppVersionValidationErrorBody) *NullablePresignUploadAppVersionValidationErrorBody {
	return &NullablePresignUploadAppVersionValidationErrorBody{value: val, isSet: true}
}

func (v NullablePresignUploadAppVersionValidationErrorBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresignUploadAppVersionValidationErrorBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

