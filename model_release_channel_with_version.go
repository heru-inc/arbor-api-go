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

// checks if the ReleaseChannelWithVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseChannelWithVersion{}

// ReleaseChannelWithVersion Represents a release channel that is designated in your organization with the current app version that is assigned to it.
type ReleaseChannelWithVersion struct {
	Version AppVersion `json:"version"`
	Id string `json:"id"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ReleaseChannelWithVersion ReleaseChannelWithVersion

// NewReleaseChannelWithVersion instantiates a new ReleaseChannelWithVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseChannelWithVersion(version AppVersion, id string, name string) *ReleaseChannelWithVersion {
	this := ReleaseChannelWithVersion{}
	this.Id = id
	this.Name = name
	return &this
}

// NewReleaseChannelWithVersionWithDefaults instantiates a new ReleaseChannelWithVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseChannelWithVersionWithDefaults() *ReleaseChannelWithVersion {
	this := ReleaseChannelWithVersion{}
	return &this
}

// GetVersion returns the Version field value
func (o *ReleaseChannelWithVersion) GetVersion() AppVersion {
	if o == nil {
		var ret AppVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ReleaseChannelWithVersion) GetVersionOk() (*AppVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ReleaseChannelWithVersion) SetVersion(v AppVersion) {
	o.Version = v
}

// GetId returns the Id field value
func (o *ReleaseChannelWithVersion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReleaseChannelWithVersion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReleaseChannelWithVersion) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ReleaseChannelWithVersion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReleaseChannelWithVersion) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReleaseChannelWithVersion) SetName(v string) {
	o.Name = v
}

func (o ReleaseChannelWithVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseChannelWithVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReleaseChannelWithVersion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"id",
		"name",
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

	varReleaseChannelWithVersion := _ReleaseChannelWithVersion{}

	err = json.Unmarshal(data, &varReleaseChannelWithVersion)

	if err != nil {
		return err
	}

	*o = ReleaseChannelWithVersion(varReleaseChannelWithVersion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "version")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReleaseChannelWithVersion struct {
	value *ReleaseChannelWithVersion
	isSet bool
}

func (v NullableReleaseChannelWithVersion) Get() *ReleaseChannelWithVersion {
	return v.value
}

func (v *NullableReleaseChannelWithVersion) Set(val *ReleaseChannelWithVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseChannelWithVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseChannelWithVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseChannelWithVersion(val *ReleaseChannelWithVersion) *NullableReleaseChannelWithVersion {
	return &NullableReleaseChannelWithVersion{value: val, isSet: true}
}

func (v NullableReleaseChannelWithVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseChannelWithVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


