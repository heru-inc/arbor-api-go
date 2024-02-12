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

// checks if the ReleaseChannelsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseChannelsResponse{}

// ReleaseChannelsResponse A paginated list of release channels.
type ReleaseChannelsResponse struct {
	Data []ReleaseChannel `json:"data"`
	Links Links `json:"links"`
	Meta Meta `json:"meta"`
}

type _ReleaseChannelsResponse ReleaseChannelsResponse

// NewReleaseChannelsResponse instantiates a new ReleaseChannelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseChannelsResponse(data []ReleaseChannel, links Links, meta Meta) *ReleaseChannelsResponse {
	this := ReleaseChannelsResponse{}
	this.Data = data
	this.Links = links
	this.Meta = meta
	return &this
}

// NewReleaseChannelsResponseWithDefaults instantiates a new ReleaseChannelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseChannelsResponseWithDefaults() *ReleaseChannelsResponse {
	this := ReleaseChannelsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ReleaseChannelsResponse) GetData() []ReleaseChannel {
	if o == nil {
		var ret []ReleaseChannel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReleaseChannelsResponse) GetDataOk() ([]ReleaseChannel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ReleaseChannelsResponse) SetData(v []ReleaseChannel) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ReleaseChannelsResponse) GetLinks() Links {
	if o == nil {
		var ret Links
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ReleaseChannelsResponse) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ReleaseChannelsResponse) SetLinks(v Links) {
	o.Links = v
}

// GetMeta returns the Meta field value
func (o *ReleaseChannelsResponse) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ReleaseChannelsResponse) GetMetaOk() (*Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ReleaseChannelsResponse) SetMeta(v Meta) {
	o.Meta = v
}

func (o ReleaseChannelsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseChannelsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

func (o *ReleaseChannelsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varReleaseChannelsResponse := _ReleaseChannelsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReleaseChannelsResponse)

	if err != nil {
		return err
	}

	*o = ReleaseChannelsResponse(varReleaseChannelsResponse)

	return err
}

type NullableReleaseChannelsResponse struct {
	value *ReleaseChannelsResponse
	isSet bool
}

func (v NullableReleaseChannelsResponse) Get() *ReleaseChannelsResponse {
	return v.value
}

func (v *NullableReleaseChannelsResponse) Set(val *ReleaseChannelsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseChannelsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseChannelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseChannelsResponse(val *ReleaseChannelsResponse) *NullableReleaseChannelsResponse {
	return &NullableReleaseChannelsResponse{value: val, isSet: true}
}

func (v NullableReleaseChannelsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseChannelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

