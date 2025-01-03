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

// checks if the AppsRevokeReleaseChannelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppsRevokeReleaseChannelRequest{}

// AppsRevokeReleaseChannelRequest struct for AppsRevokeReleaseChannelRequest
type AppsRevokeReleaseChannelRequest struct {
	OrganizationSlug string `json:"organizationSlug"`
	AdditionalProperties map[string]interface{}
}

type _AppsRevokeReleaseChannelRequest AppsRevokeReleaseChannelRequest

// NewAppsRevokeReleaseChannelRequest instantiates a new AppsRevokeReleaseChannelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsRevokeReleaseChannelRequest(organizationSlug string) *AppsRevokeReleaseChannelRequest {
	this := AppsRevokeReleaseChannelRequest{}
	this.OrganizationSlug = organizationSlug
	return &this
}

// NewAppsRevokeReleaseChannelRequestWithDefaults instantiates a new AppsRevokeReleaseChannelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsRevokeReleaseChannelRequestWithDefaults() *AppsRevokeReleaseChannelRequest {
	this := AppsRevokeReleaseChannelRequest{}
	return &this
}

// GetOrganizationSlug returns the OrganizationSlug field value
func (o *AppsRevokeReleaseChannelRequest) GetOrganizationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationSlug
}

// GetOrganizationSlugOk returns a tuple with the OrganizationSlug field value
// and a boolean to check if the value has been set.
func (o *AppsRevokeReleaseChannelRequest) GetOrganizationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationSlug, true
}

// SetOrganizationSlug sets field value
func (o *AppsRevokeReleaseChannelRequest) SetOrganizationSlug(v string) {
	o.OrganizationSlug = v
}

func (o AppsRevokeReleaseChannelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppsRevokeReleaseChannelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["organizationSlug"] = o.OrganizationSlug

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppsRevokeReleaseChannelRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"organizationSlug",
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

	varAppsRevokeReleaseChannelRequest := _AppsRevokeReleaseChannelRequest{}

	err = json.Unmarshal(data, &varAppsRevokeReleaseChannelRequest)

	if err != nil {
		return err
	}

	*o = AppsRevokeReleaseChannelRequest(varAppsRevokeReleaseChannelRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "organizationSlug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppsRevokeReleaseChannelRequest struct {
	value *AppsRevokeReleaseChannelRequest
	isSet bool
}

func (v NullableAppsRevokeReleaseChannelRequest) Get() *AppsRevokeReleaseChannelRequest {
	return v.value
}

func (v *NullableAppsRevokeReleaseChannelRequest) Set(val *AppsRevokeReleaseChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsRevokeReleaseChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsRevokeReleaseChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsRevokeReleaseChannelRequest(val *AppsRevokeReleaseChannelRequest) *NullableAppsRevokeReleaseChannelRequest {
	return &NullableAppsRevokeReleaseChannelRequest{value: val, isSet: true}
}

func (v NullableAppsRevokeReleaseChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsRevokeReleaseChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


