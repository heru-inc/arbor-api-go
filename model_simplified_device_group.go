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

// checks if the SimplifiedDeviceGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplifiedDeviceGroup{}

// SimplifiedDeviceGroup Represents a device group that is managed by your organization (id and title only).
type SimplifiedDeviceGroup struct {
	Id string `json:"id"`
	Title string `json:"title"`
	AdditionalProperties map[string]interface{}
}

type _SimplifiedDeviceGroup SimplifiedDeviceGroup

// NewSimplifiedDeviceGroup instantiates a new SimplifiedDeviceGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedDeviceGroup(id string, title string) *SimplifiedDeviceGroup {
	this := SimplifiedDeviceGroup{}
	this.Id = id
	this.Title = title
	return &this
}

// NewSimplifiedDeviceGroupWithDefaults instantiates a new SimplifiedDeviceGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedDeviceGroupWithDefaults() *SimplifiedDeviceGroup {
	this := SimplifiedDeviceGroup{}
	return &this
}

// GetId returns the Id field value
func (o *SimplifiedDeviceGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimplifiedDeviceGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimplifiedDeviceGroup) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *SimplifiedDeviceGroup) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SimplifiedDeviceGroup) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *SimplifiedDeviceGroup) SetTitle(v string) {
	o.Title = v
}

func (o SimplifiedDeviceGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplifiedDeviceGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimplifiedDeviceGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"title",
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

	varSimplifiedDeviceGroup := _SimplifiedDeviceGroup{}

	err = json.Unmarshal(data, &varSimplifiedDeviceGroup)

	if err != nil {
		return err
	}

	*o = SimplifiedDeviceGroup(varSimplifiedDeviceGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimplifiedDeviceGroup struct {
	value *SimplifiedDeviceGroup
	isSet bool
}

func (v NullableSimplifiedDeviceGroup) Get() *SimplifiedDeviceGroup {
	return v.value
}

func (v *NullableSimplifiedDeviceGroup) Set(val *SimplifiedDeviceGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedDeviceGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedDeviceGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedDeviceGroup(val *SimplifiedDeviceGroup) *NullableSimplifiedDeviceGroup {
	return &NullableSimplifiedDeviceGroup{value: val, isSet: true}
}

func (v NullableSimplifiedDeviceGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedDeviceGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


