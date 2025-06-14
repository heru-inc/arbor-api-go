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

// checks if the GetAppReleaseChannels200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAppReleaseChannels200ResponseDataInner{}

// GetAppReleaseChannels200ResponseDataInner struct for GetAppReleaseChannels200ResponseDataInner
type GetAppReleaseChannels200ResponseDataInner struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Version *GetAppVersions200ResponseDataInner `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAppReleaseChannels200ResponseDataInner GetAppReleaseChannels200ResponseDataInner

// NewGetAppReleaseChannels200ResponseDataInner instantiates a new GetAppReleaseChannels200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAppReleaseChannels200ResponseDataInner() *GetAppReleaseChannels200ResponseDataInner {
	this := GetAppReleaseChannels200ResponseDataInner{}
	return &this
}

// NewGetAppReleaseChannels200ResponseDataInnerWithDefaults instantiates a new GetAppReleaseChannels200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAppReleaseChannels200ResponseDataInnerWithDefaults() *GetAppReleaseChannels200ResponseDataInner {
	this := GetAppReleaseChannels200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetAppReleaseChannels200ResponseDataInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppReleaseChannels200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetAppReleaseChannels200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetAppReleaseChannels200ResponseDataInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetAppReleaseChannels200ResponseDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppReleaseChannels200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetAppReleaseChannels200ResponseDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetAppReleaseChannels200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetAppReleaseChannels200ResponseDataInner) GetVersion() GetAppVersions200ResponseDataInner {
	if o == nil || IsNil(o.Version) {
		var ret GetAppVersions200ResponseDataInner
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppReleaseChannels200ResponseDataInner) GetVersionOk() (*GetAppVersions200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetAppReleaseChannels200ResponseDataInner) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given GetAppVersions200ResponseDataInner and assigns it to the Version field.
func (o *GetAppReleaseChannels200ResponseDataInner) SetVersion(v GetAppVersions200ResponseDataInner) {
	o.Version = &v
}

func (o GetAppReleaseChannels200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAppReleaseChannels200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAppReleaseChannels200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	varGetAppReleaseChannels200ResponseDataInner := _GetAppReleaseChannels200ResponseDataInner{}

	err = json.Unmarshal(data, &varGetAppReleaseChannels200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = GetAppReleaseChannels200ResponseDataInner(varGetAppReleaseChannels200ResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAppReleaseChannels200ResponseDataInner struct {
	value *GetAppReleaseChannels200ResponseDataInner
	isSet bool
}

func (v NullableGetAppReleaseChannels200ResponseDataInner) Get() *GetAppReleaseChannels200ResponseDataInner {
	return v.value
}

func (v *NullableGetAppReleaseChannels200ResponseDataInner) Set(val *GetAppReleaseChannels200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAppReleaseChannels200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAppReleaseChannels200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAppReleaseChannels200ResponseDataInner(val *GetAppReleaseChannels200ResponseDataInner) *NullableGetAppReleaseChannels200ResponseDataInner {
	return &NullableGetAppReleaseChannels200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetAppReleaseChannels200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAppReleaseChannels200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


