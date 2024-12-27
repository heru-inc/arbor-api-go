/*
arborxr-api-v2

This API provides a RESTful interface to interact with your organization's data.

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arbor-api-go

import (
	"encoding/json"
)

// checks if the UsersUpdateUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersUpdateUserRequest{}

// UsersUpdateUserRequest struct for UsersUpdateUserRequest
type UsersUpdateUserRequest struct {
	FirstName NullableString `json:"firstName,omitempty"`
	LastName NullableString `json:"lastName,omitempty"`
	Email NullableString `json:"email,omitempty"`
	OrganizationRoleId NullableString `json:"organizationRoleId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UsersUpdateUserRequest UsersUpdateUserRequest

// NewUsersUpdateUserRequest instantiates a new UsersUpdateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersUpdateUserRequest() *UsersUpdateUserRequest {
	this := UsersUpdateUserRequest{}
	return &this
}

// NewUsersUpdateUserRequestWithDefaults instantiates a new UsersUpdateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersUpdateUserRequestWithDefaults() *UsersUpdateUserRequest {
	this := UsersUpdateUserRequest{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsersUpdateUserRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersUpdateUserRequest) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *UsersUpdateUserRequest) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *UsersUpdateUserRequest) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *UsersUpdateUserRequest) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *UsersUpdateUserRequest) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsersUpdateUserRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersUpdateUserRequest) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *UsersUpdateUserRequest) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *UsersUpdateUserRequest) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *UsersUpdateUserRequest) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *UsersUpdateUserRequest) UnsetLastName() {
	o.LastName.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsersUpdateUserRequest) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersUpdateUserRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UsersUpdateUserRequest) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UsersUpdateUserRequest) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UsersUpdateUserRequest) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UsersUpdateUserRequest) UnsetEmail() {
	o.Email.Unset()
}

// GetOrganizationRoleId returns the OrganizationRoleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsersUpdateUserRequest) GetOrganizationRoleId() string {
	if o == nil || IsNil(o.OrganizationRoleId.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationRoleId.Get()
}

// GetOrganizationRoleIdOk returns a tuple with the OrganizationRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersUpdateUserRequest) GetOrganizationRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationRoleId.Get(), o.OrganizationRoleId.IsSet()
}

// HasOrganizationRoleId returns a boolean if a field has been set.
func (o *UsersUpdateUserRequest) HasOrganizationRoleId() bool {
	if o != nil && o.OrganizationRoleId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationRoleId gets a reference to the given NullableString and assigns it to the OrganizationRoleId field.
func (o *UsersUpdateUserRequest) SetOrganizationRoleId(v string) {
	o.OrganizationRoleId.Set(&v)
}
// SetOrganizationRoleIdNil sets the value for OrganizationRoleId to be an explicit nil
func (o *UsersUpdateUserRequest) SetOrganizationRoleIdNil() {
	o.OrganizationRoleId.Set(nil)
}

// UnsetOrganizationRoleId ensures that no value is present for OrganizationRoleId, not even an explicit nil
func (o *UsersUpdateUserRequest) UnsetOrganizationRoleId() {
	o.OrganizationRoleId.Unset()
}

func (o UsersUpdateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersUpdateUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName.IsSet() {
		toSerialize["firstName"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["lastName"] = o.LastName.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.OrganizationRoleId.IsSet() {
		toSerialize["organizationRoleId"] = o.OrganizationRoleId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UsersUpdateUserRequest) UnmarshalJSON(data []byte) (err error) {
	varUsersUpdateUserRequest := _UsersUpdateUserRequest{}

	err = json.Unmarshal(data, &varUsersUpdateUserRequest)

	if err != nil {
		return err
	}

	*o = UsersUpdateUserRequest(varUsersUpdateUserRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "email")
		delete(additionalProperties, "organizationRoleId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsersUpdateUserRequest struct {
	value *UsersUpdateUserRequest
	isSet bool
}

func (v NullableUsersUpdateUserRequest) Get() *UsersUpdateUserRequest {
	return v.value
}

func (v *NullableUsersUpdateUserRequest) Set(val *UsersUpdateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersUpdateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersUpdateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersUpdateUserRequest(val *UsersUpdateUserRequest) *NullableUsersUpdateUserRequest {
	return &NullableUsersUpdateUserRequest{value: val, isSet: true}
}

func (v NullableUsersUpdateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersUpdateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


