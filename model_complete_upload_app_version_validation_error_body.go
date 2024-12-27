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

// checks if the CompleteUploadAppVersionValidationErrorBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompleteUploadAppVersionValidationErrorBody{}

// CompleteUploadAppVersionValidationErrorBody struct for CompleteUploadAppVersionValidationErrorBody
type CompleteUploadAppVersionValidationErrorBody struct {
	Message string `json:"message"`
	Errors CompleteUploadAppVersionValidationErrorBodyErrors `json:"errors"`
	AdditionalProperties map[string]interface{}
}

type _CompleteUploadAppVersionValidationErrorBody CompleteUploadAppVersionValidationErrorBody

// NewCompleteUploadAppVersionValidationErrorBody instantiates a new CompleteUploadAppVersionValidationErrorBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteUploadAppVersionValidationErrorBody(message string, errors CompleteUploadAppVersionValidationErrorBodyErrors) *CompleteUploadAppVersionValidationErrorBody {
	this := CompleteUploadAppVersionValidationErrorBody{}
	this.Message = message
	this.Errors = errors
	return &this
}

// NewCompleteUploadAppVersionValidationErrorBodyWithDefaults instantiates a new CompleteUploadAppVersionValidationErrorBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteUploadAppVersionValidationErrorBodyWithDefaults() *CompleteUploadAppVersionValidationErrorBody {
	this := CompleteUploadAppVersionValidationErrorBody{}
	return &this
}

// GetMessage returns the Message field value
func (o *CompleteUploadAppVersionValidationErrorBody) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CompleteUploadAppVersionValidationErrorBody) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CompleteUploadAppVersionValidationErrorBody) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value
func (o *CompleteUploadAppVersionValidationErrorBody) GetErrors() CompleteUploadAppVersionValidationErrorBodyErrors {
	if o == nil {
		var ret CompleteUploadAppVersionValidationErrorBodyErrors
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *CompleteUploadAppVersionValidationErrorBody) GetErrorsOk() (*CompleteUploadAppVersionValidationErrorBodyErrors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *CompleteUploadAppVersionValidationErrorBody) SetErrors(v CompleteUploadAppVersionValidationErrorBodyErrors) {
	o.Errors = v
}

func (o CompleteUploadAppVersionValidationErrorBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompleteUploadAppVersionValidationErrorBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["errors"] = o.Errors

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompleteUploadAppVersionValidationErrorBody) UnmarshalJSON(data []byte) (err error) {
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

	varCompleteUploadAppVersionValidationErrorBody := _CompleteUploadAppVersionValidationErrorBody{}

	err = json.Unmarshal(data, &varCompleteUploadAppVersionValidationErrorBody)

	if err != nil {
		return err
	}

	*o = CompleteUploadAppVersionValidationErrorBody(varCompleteUploadAppVersionValidationErrorBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompleteUploadAppVersionValidationErrorBody struct {
	value *CompleteUploadAppVersionValidationErrorBody
	isSet bool
}

func (v NullableCompleteUploadAppVersionValidationErrorBody) Get() *CompleteUploadAppVersionValidationErrorBody {
	return v.value
}

func (v *NullableCompleteUploadAppVersionValidationErrorBody) Set(val *CompleteUploadAppVersionValidationErrorBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteUploadAppVersionValidationErrorBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteUploadAppVersionValidationErrorBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteUploadAppVersionValidationErrorBody(val *CompleteUploadAppVersionValidationErrorBody) *NullableCompleteUploadAppVersionValidationErrorBody {
	return &NullableCompleteUploadAppVersionValidationErrorBody{value: val, isSet: true}
}

func (v NullableCompleteUploadAppVersionValidationErrorBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteUploadAppVersionValidationErrorBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


