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

// checks if the GetGroupRoles200ResponseMetaLinksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGroupRoles200ResponseMetaLinksInner{}

// GetGroupRoles200ResponseMetaLinksInner struct for GetGroupRoles200ResponseMetaLinksInner
type GetGroupRoles200ResponseMetaLinksInner struct {
	Url NullableString `json:"url,omitempty"`
	Label *string `json:"label,omitempty"`
	Active *bool `json:"active,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetGroupRoles200ResponseMetaLinksInner GetGroupRoles200ResponseMetaLinksInner

// NewGetGroupRoles200ResponseMetaLinksInner instantiates a new GetGroupRoles200ResponseMetaLinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGroupRoles200ResponseMetaLinksInner() *GetGroupRoles200ResponseMetaLinksInner {
	this := GetGroupRoles200ResponseMetaLinksInner{}
	return &this
}

// NewGetGroupRoles200ResponseMetaLinksInnerWithDefaults instantiates a new GetGroupRoles200ResponseMetaLinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGroupRoles200ResponseMetaLinksInnerWithDefaults() *GetGroupRoles200ResponseMetaLinksInner {
	this := GetGroupRoles200ResponseMetaLinksInner{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetGroupRoles200ResponseMetaLinksInner) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetGroupRoles200ResponseMetaLinksInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *GetGroupRoles200ResponseMetaLinksInner) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *GetGroupRoles200ResponseMetaLinksInner) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *GetGroupRoles200ResponseMetaLinksInner) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *GetGroupRoles200ResponseMetaLinksInner) UnsetUrl() {
	o.Url.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetGroupRoles200ResponseMetaLinksInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupRoles200ResponseMetaLinksInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetGroupRoles200ResponseMetaLinksInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetGroupRoles200ResponseMetaLinksInner) SetLabel(v string) {
	o.Label = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GetGroupRoles200ResponseMetaLinksInner) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupRoles200ResponseMetaLinksInner) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GetGroupRoles200ResponseMetaLinksInner) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GetGroupRoles200ResponseMetaLinksInner) SetActive(v bool) {
	o.Active = &v
}

func (o GetGroupRoles200ResponseMetaLinksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGroupRoles200ResponseMetaLinksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetGroupRoles200ResponseMetaLinksInner) UnmarshalJSON(data []byte) (err error) {
	varGetGroupRoles200ResponseMetaLinksInner := _GetGroupRoles200ResponseMetaLinksInner{}

	err = json.Unmarshal(data, &varGetGroupRoles200ResponseMetaLinksInner)

	if err != nil {
		return err
	}

	*o = GetGroupRoles200ResponseMetaLinksInner(varGetGroupRoles200ResponseMetaLinksInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "label")
		delete(additionalProperties, "active")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetGroupRoles200ResponseMetaLinksInner struct {
	value *GetGroupRoles200ResponseMetaLinksInner
	isSet bool
}

func (v NullableGetGroupRoles200ResponseMetaLinksInner) Get() *GetGroupRoles200ResponseMetaLinksInner {
	return v.value
}

func (v *NullableGetGroupRoles200ResponseMetaLinksInner) Set(val *GetGroupRoles200ResponseMetaLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGroupRoles200ResponseMetaLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGroupRoles200ResponseMetaLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGroupRoles200ResponseMetaLinksInner(val *GetGroupRoles200ResponseMetaLinksInner) *NullableGetGroupRoles200ResponseMetaLinksInner {
	return &NullableGetGroupRoles200ResponseMetaLinksInner{value: val, isSet: true}
}

func (v NullableGetGroupRoles200ResponseMetaLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGroupRoles200ResponseMetaLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


