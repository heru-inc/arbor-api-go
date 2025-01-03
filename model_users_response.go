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

// checks if the UsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersResponse{}

// UsersResponse A paginated list of users.
type UsersResponse struct {
	Data []User `json:"data"`
	Links Links `json:"links"`
	Meta Meta `json:"meta"`
	AdditionalProperties map[string]interface{}
}

type _UsersResponse UsersResponse

// NewUsersResponse instantiates a new UsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersResponse(data []User, links Links, meta Meta) *UsersResponse {
	this := UsersResponse{}
	this.Data = data
	this.Links = links
	this.Meta = meta
	return &this
}

// NewUsersResponseWithDefaults instantiates a new UsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersResponseWithDefaults() *UsersResponse {
	this := UsersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UsersResponse) GetData() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UsersResponse) GetDataOk() ([]User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *UsersResponse) SetData(v []User) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *UsersResponse) GetLinks() Links {
	if o == nil {
		var ret Links
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *UsersResponse) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *UsersResponse) SetLinks(v Links) {
	o.Links = v
}

// GetMeta returns the Meta field value
func (o *UsersResponse) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *UsersResponse) GetMetaOk() (*Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *UsersResponse) SetMeta(v Meta) {
	o.Meta = v
}

func (o UsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UsersResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUsersResponse := _UsersResponse{}

	err = json.Unmarshal(data, &varUsersResponse)

	if err != nil {
		return err
	}

	*o = UsersResponse(varUsersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "links")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsersResponse struct {
	value *UsersResponse
	isSet bool
}

func (v NullableUsersResponse) Get() *UsersResponse {
	return v.value
}

func (v *NullableUsersResponse) Set(val *UsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersResponse(val *UsersResponse) *NullableUsersResponse {
	return &NullableUsersResponse{value: val, isSet: true}
}

func (v NullableUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


