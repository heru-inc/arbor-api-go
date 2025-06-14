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

// checks if the UpdateUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserRequest{}

// UpdateUserRequest struct for UpdateUserRequest
type UpdateUserRequest struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	OrganizationRoleId NullableString `json:"organizationRoleId,omitempty"`
	// Required when `groupIds` is provided.
	GroupRoleId NullableString `json:"groupRoleId,omitempty"`
	// Required when `groupRoleId` is provided.
	GroupIds []string `json:"groupIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateUserRequest UpdateUserRequest

// NewUpdateUserRequest instantiates a new UpdateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserRequest() *UpdateUserRequest {
	this := UpdateUserRequest{}
	return &this
}

// NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserRequestWithDefaults() *UpdateUserRequest {
	this := UpdateUserRequest{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UpdateUserRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UpdateUserRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetOrganizationRoleId returns the OrganizationRoleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserRequest) GetOrganizationRoleId() string {
	if o == nil || IsNil(o.OrganizationRoleId.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationRoleId.Get()
}

// GetOrganizationRoleIdOk returns a tuple with the OrganizationRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserRequest) GetOrganizationRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationRoleId.Get(), o.OrganizationRoleId.IsSet()
}

// HasOrganizationRoleId returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasOrganizationRoleId() bool {
	if o != nil && o.OrganizationRoleId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationRoleId gets a reference to the given NullableString and assigns it to the OrganizationRoleId field.
func (o *UpdateUserRequest) SetOrganizationRoleId(v string) {
	o.OrganizationRoleId.Set(&v)
}
// SetOrganizationRoleIdNil sets the value for OrganizationRoleId to be an explicit nil
func (o *UpdateUserRequest) SetOrganizationRoleIdNil() {
	o.OrganizationRoleId.Set(nil)
}

// UnsetOrganizationRoleId ensures that no value is present for OrganizationRoleId, not even an explicit nil
func (o *UpdateUserRequest) UnsetOrganizationRoleId() {
	o.OrganizationRoleId.Unset()
}

// GetGroupRoleId returns the GroupRoleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserRequest) GetGroupRoleId() string {
	if o == nil || IsNil(o.GroupRoleId.Get()) {
		var ret string
		return ret
	}
	return *o.GroupRoleId.Get()
}

// GetGroupRoleIdOk returns a tuple with the GroupRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserRequest) GetGroupRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupRoleId.Get(), o.GroupRoleId.IsSet()
}

// HasGroupRoleId returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasGroupRoleId() bool {
	if o != nil && o.GroupRoleId.IsSet() {
		return true
	}

	return false
}

// SetGroupRoleId gets a reference to the given NullableString and assigns it to the GroupRoleId field.
func (o *UpdateUserRequest) SetGroupRoleId(v string) {
	o.GroupRoleId.Set(&v)
}
// SetGroupRoleIdNil sets the value for GroupRoleId to be an explicit nil
func (o *UpdateUserRequest) SetGroupRoleIdNil() {
	o.GroupRoleId.Set(nil)
}

// UnsetGroupRoleId ensures that no value is present for GroupRoleId, not even an explicit nil
func (o *UpdateUserRequest) UnsetGroupRoleId() {
	o.GroupRoleId.Unset()
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetGroupIds() []string {
	if o == nil || IsNil(o.GroupIds) {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupIds) {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasGroupIds() bool {
	if o != nil && !IsNil(o.GroupIds) {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *UpdateUserRequest) SetGroupIds(v []string) {
	o.GroupIds = v
}

func (o UpdateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if o.OrganizationRoleId.IsSet() {
		toSerialize["organizationRoleId"] = o.OrganizationRoleId.Get()
	}
	if o.GroupRoleId.IsSet() {
		toSerialize["groupRoleId"] = o.GroupRoleId.Get()
	}
	if !IsNil(o.GroupIds) {
		toSerialize["groupIds"] = o.GroupIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateUserRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateUserRequest := _UpdateUserRequest{}

	err = json.Unmarshal(data, &varUpdateUserRequest)

	if err != nil {
		return err
	}

	*o = UpdateUserRequest(varUpdateUserRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "organizationRoleId")
		delete(additionalProperties, "groupRoleId")
		delete(additionalProperties, "groupIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateUserRequest struct {
	value *UpdateUserRequest
	isSet bool
}

func (v NullableUpdateUserRequest) Get() *UpdateUserRequest {
	return v.value
}

func (v *NullableUpdateUserRequest) Set(val *UpdateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserRequest(val *UpdateUserRequest) *NullableUpdateUserRequest {
	return &NullableUpdateUserRequest{value: val, isSet: true}
}

func (v NullableUpdateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


