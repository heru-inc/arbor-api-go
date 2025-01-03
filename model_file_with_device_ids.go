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

// checks if the FileWithDeviceIds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileWithDeviceIds{}

// FileWithDeviceIds Represents a file that is stored in your organization's content library with a list of device ids that it is installed on.
type FileWithDeviceIds struct {
	DeviceIds []string `json:"deviceIds"`
	Id string `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
	Size string `json:"size"`
	Tags []string `json:"tags"`
	Added Time `json:"added"`
	AdditionalProperties map[string]interface{}
}

type _FileWithDeviceIds FileWithDeviceIds

// NewFileWithDeviceIds instantiates a new FileWithDeviceIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileWithDeviceIds(deviceIds []string, id string, name string, location string, size string, tags []string, added Time) *FileWithDeviceIds {
	this := FileWithDeviceIds{}
	this.Id = id
	this.Name = name
	this.Location = location
	this.Size = size
	this.Tags = tags
	this.Added = added
	return &this
}

// NewFileWithDeviceIdsWithDefaults instantiates a new FileWithDeviceIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileWithDeviceIdsWithDefaults() *FileWithDeviceIds {
	this := FileWithDeviceIds{}
	return &this
}

// GetDeviceIds returns the DeviceIds field value
func (o *FileWithDeviceIds) GetDeviceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value
// and a boolean to check if the value has been set.
func (o *FileWithDeviceIds) GetDeviceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceIds, true
}

// SetDeviceIds sets field value
func (o *FileWithDeviceIds) SetDeviceIds(v []string) {
	o.DeviceIds = v
}

// GetId returns the Id field value
func (o *FileWithDeviceIds) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileWithDeviceIds) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileWithDeviceIds) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FileWithDeviceIds) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FileWithDeviceIds) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FileWithDeviceIds) SetName(v string) {
	o.Name = v
}

// GetLocation returns the Location field value
func (o *FileWithDeviceIds) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *FileWithDeviceIds) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *FileWithDeviceIds) SetLocation(v string) {
	o.Location = v
}

// GetSize returns the Size field value
func (o *FileWithDeviceIds) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *FileWithDeviceIds) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *FileWithDeviceIds) SetSize(v string) {
	o.Size = v
}

// GetTags returns the Tags field value
func (o *FileWithDeviceIds) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *FileWithDeviceIds) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *FileWithDeviceIds) SetTags(v []string) {
	o.Tags = v
}

// GetAdded returns the Added field value
func (o *FileWithDeviceIds) GetAdded() Time {
	if o == nil {
		var ret Time
		return ret
	}

	return o.Added
}

// GetAddedOk returns a tuple with the Added field value
// and a boolean to check if the value has been set.
func (o *FileWithDeviceIds) GetAddedOk() (*Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Added, true
}

// SetAdded sets field value
func (o *FileWithDeviceIds) SetAdded(v Time) {
	o.Added = v
}

func (o FileWithDeviceIds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileWithDeviceIds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceIds"] = o.DeviceIds
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["location"] = o.Location
	toSerialize["size"] = o.Size
	toSerialize["tags"] = o.Tags
	toSerialize["added"] = o.Added

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FileWithDeviceIds) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deviceIds",
		"id",
		"name",
		"location",
		"size",
		"tags",
		"added",
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

	varFileWithDeviceIds := _FileWithDeviceIds{}

	err = json.Unmarshal(data, &varFileWithDeviceIds)

	if err != nil {
		return err
	}

	*o = FileWithDeviceIds(varFileWithDeviceIds)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deviceIds")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "location")
		delete(additionalProperties, "size")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "added")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileWithDeviceIds struct {
	value *FileWithDeviceIds
	isSet bool
}

func (v NullableFileWithDeviceIds) Get() *FileWithDeviceIds {
	return v.value
}

func (v *NullableFileWithDeviceIds) Set(val *FileWithDeviceIds) {
	v.value = val
	v.isSet = true
}

func (v NullableFileWithDeviceIds) IsSet() bool {
	return v.isSet
}

func (v *NullableFileWithDeviceIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileWithDeviceIds(val *FileWithDeviceIds) *NullableFileWithDeviceIds {
	return &NullableFileWithDeviceIds{value: val, isSet: true}
}

func (v NullableFileWithDeviceIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileWithDeviceIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


