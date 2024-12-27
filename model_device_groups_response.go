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

// checks if the DeviceGroupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceGroupsResponse{}

// DeviceGroupsResponse A paginated list of device groups.
type DeviceGroupsResponse struct {
	Data []DeviceGroup `json:"data"`
	Links Links `json:"links"`
	Meta Meta `json:"meta"`
	AdditionalProperties map[string]interface{}
}

type _DeviceGroupsResponse DeviceGroupsResponse

// NewDeviceGroupsResponse instantiates a new DeviceGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceGroupsResponse(data []DeviceGroup, links Links, meta Meta) *DeviceGroupsResponse {
	this := DeviceGroupsResponse{}
	this.Data = data
	this.Links = links
	this.Meta = meta
	return &this
}

// NewDeviceGroupsResponseWithDefaults instantiates a new DeviceGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceGroupsResponseWithDefaults() *DeviceGroupsResponse {
	this := DeviceGroupsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *DeviceGroupsResponse) GetData() []DeviceGroup {
	if o == nil {
		var ret []DeviceGroup
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupsResponse) GetDataOk() ([]DeviceGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *DeviceGroupsResponse) SetData(v []DeviceGroup) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *DeviceGroupsResponse) GetLinks() Links {
	if o == nil {
		var ret Links
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupsResponse) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *DeviceGroupsResponse) SetLinks(v Links) {
	o.Links = v
}

// GetMeta returns the Meta field value
func (o *DeviceGroupsResponse) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupsResponse) GetMetaOk() (*Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *DeviceGroupsResponse) SetMeta(v Meta) {
	o.Meta = v
}

func (o DeviceGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceGroupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"links",
		"meta",
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

	varDeviceGroupsResponse := _DeviceGroupsResponse{}

	err = json.Unmarshal(data, &varDeviceGroupsResponse)

	if err != nil {
		return err
	}

	*o = DeviceGroupsResponse(varDeviceGroupsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "links")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceGroupsResponse struct {
	value *DeviceGroupsResponse
	isSet bool
}

func (v NullableDeviceGroupsResponse) Get() *DeviceGroupsResponse {
	return v.value
}

func (v *NullableDeviceGroupsResponse) Set(val *DeviceGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceGroupsResponse(val *DeviceGroupsResponse) *NullableDeviceGroupsResponse {
	return &NullableDeviceGroupsResponse{value: val, isSet: true}
}

func (v NullableDeviceGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


