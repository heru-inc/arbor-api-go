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

// checks if the AppVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppVersion{}

// AppVersion Represents an uploaded version of an app in your content library.
type AppVersion struct {
	Id string `json:"id"`
	Version string `json:"version"`
	Code NullableInt32 `json:"code,omitempty"`
	Size NullableString `json:"size,omitempty"`
	Added Time `json:"added"`
	Status AppVersionStatus `json:"status"`
	DownloadUrl NullableString `json:"downloadUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppVersion AppVersion

// NewAppVersion instantiates a new AppVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppVersion(id string, version string, added Time, status AppVersionStatus) *AppVersion {
	this := AppVersion{}
	this.Id = id
	this.Version = version
	this.Added = added
	this.Status = status
	return &this
}

// NewAppVersionWithDefaults instantiates a new AppVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppVersionWithDefaults() *AppVersion {
	this := AppVersion{}
	return &this
}

// GetId returns the Id field value
func (o *AppVersion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppVersion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppVersion) SetId(v string) {
	o.Id = v
}

// GetVersion returns the Version field value
func (o *AppVersion) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AppVersion) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AppVersion) SetVersion(v string) {
	o.Version = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppVersion) GetCode() int32 {
	if o == nil || IsNil(o.Code.Get()) {
		var ret int32
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppVersion) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *AppVersion) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableInt32 and assigns it to the Code field.
func (o *AppVersion) SetCode(v int32) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *AppVersion) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *AppVersion) UnsetCode() {
	o.Code.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppVersion) GetSize() string {
	if o == nil || IsNil(o.Size.Get()) {
		var ret string
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppVersion) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *AppVersion) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableString and assigns it to the Size field.
func (o *AppVersion) SetSize(v string) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *AppVersion) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *AppVersion) UnsetSize() {
	o.Size.Unset()
}

// GetAdded returns the Added field value
func (o *AppVersion) GetAdded() Time {
	if o == nil {
		var ret Time
		return ret
	}

	return o.Added
}

// GetAddedOk returns a tuple with the Added field value
// and a boolean to check if the value has been set.
func (o *AppVersion) GetAddedOk() (*Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Added, true
}

// SetAdded sets field value
func (o *AppVersion) SetAdded(v Time) {
	o.Added = v
}

// GetStatus returns the Status field value
func (o *AppVersion) GetStatus() AppVersionStatus {
	if o == nil {
		var ret AppVersionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AppVersion) GetStatusOk() (*AppVersionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AppVersion) SetStatus(v AppVersionStatus) {
	o.Status = v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppVersion) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadUrl.Get()
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppVersion) GetDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadUrl.Get(), o.DownloadUrl.IsSet()
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *AppVersion) HasDownloadUrl() bool {
	if o != nil && o.DownloadUrl.IsSet() {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given NullableString and assigns it to the DownloadUrl field.
func (o *AppVersion) SetDownloadUrl(v string) {
	o.DownloadUrl.Set(&v)
}
// SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil
func (o *AppVersion) SetDownloadUrlNil() {
	o.DownloadUrl.Set(nil)
}

// UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
func (o *AppVersion) UnsetDownloadUrl() {
	o.DownloadUrl.Unset()
}

func (o AppVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["version"] = o.Version
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	toSerialize["added"] = o.Added
	toSerialize["status"] = o.Status
	if o.DownloadUrl.IsSet() {
		toSerialize["downloadUrl"] = o.DownloadUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppVersion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"version",
		"added",
		"status",
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

	varAppVersion := _AppVersion{}

	err = json.Unmarshal(data, &varAppVersion)

	if err != nil {
		return err
	}

	*o = AppVersion(varAppVersion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "code")
		delete(additionalProperties, "size")
		delete(additionalProperties, "added")
		delete(additionalProperties, "status")
		delete(additionalProperties, "downloadUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppVersion struct {
	value *AppVersion
	isSet bool
}

func (v NullableAppVersion) Get() *AppVersion {
	return v.value
}

func (v *NullableAppVersion) Set(val *AppVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableAppVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableAppVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppVersion(val *AppVersion) *NullableAppVersion {
	return &NullableAppVersion{value: val, isSet: true}
}

func (v NullableAppVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


