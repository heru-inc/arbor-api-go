/*
ArborXR Public API

This API provides a RESTful interface to interact with your organization's data.

API version: v2
Contact: support@arborxr.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arborapi

import (
	"encoding/json"
)

// checks if the UnshareReleaseChannel200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnshareReleaseChannel200Response{}

// UnshareReleaseChannel200Response struct for UnshareReleaseChannel200Response
type UnshareReleaseChannel200Response struct {
	Message *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UnshareReleaseChannel200Response UnshareReleaseChannel200Response

// NewUnshareReleaseChannel200Response instantiates a new UnshareReleaseChannel200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnshareReleaseChannel200Response() *UnshareReleaseChannel200Response {
	this := UnshareReleaseChannel200Response{}
	return &this
}

// NewUnshareReleaseChannel200ResponseWithDefaults instantiates a new UnshareReleaseChannel200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnshareReleaseChannel200ResponseWithDefaults() *UnshareReleaseChannel200Response {
	this := UnshareReleaseChannel200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UnshareReleaseChannel200Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnshareReleaseChannel200Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UnshareReleaseChannel200Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UnshareReleaseChannel200Response) SetMessage(v string) {
	o.Message = &v
}

func (o UnshareReleaseChannel200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnshareReleaseChannel200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UnshareReleaseChannel200Response) UnmarshalJSON(data []byte) (err error) {
	varUnshareReleaseChannel200Response := _UnshareReleaseChannel200Response{}

	err = json.Unmarshal(data, &varUnshareReleaseChannel200Response)

	if err != nil {
		return err
	}

	*o = UnshareReleaseChannel200Response(varUnshareReleaseChannel200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUnshareReleaseChannel200Response struct {
	value *UnshareReleaseChannel200Response
	isSet bool
}

func (v NullableUnshareReleaseChannel200Response) Get() *UnshareReleaseChannel200Response {
	return v.value
}

func (v *NullableUnshareReleaseChannel200Response) Set(val *UnshareReleaseChannel200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUnshareReleaseChannel200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUnshareReleaseChannel200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnshareReleaseChannel200Response(val *UnshareReleaseChannel200Response) *NullableUnshareReleaseChannel200Response {
	return &NullableUnshareReleaseChannel200Response{value: val, isSet: true}
}

func (v NullableUnshareReleaseChannel200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnshareReleaseChannel200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


