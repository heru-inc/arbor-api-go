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

// checks if the PaginatorLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatorLink{}

// PaginatorLink A link to a page of paginated results.
type PaginatorLink struct {
	Url NullableString `json:"url,omitempty"`
	Label string `json:"label"`
	Active bool `json:"active"`
	AdditionalProperties map[string]interface{}
}

type _PaginatorLink PaginatorLink

// NewPaginatorLink instantiates a new PaginatorLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatorLink(label string, active bool) *PaginatorLink {
	this := PaginatorLink{}
	this.Label = label
	this.Active = active
	return &this
}

// NewPaginatorLinkWithDefaults instantiates a new PaginatorLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatorLinkWithDefaults() *PaginatorLink {
	this := PaginatorLink{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatorLink) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatorLink) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *PaginatorLink) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *PaginatorLink) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *PaginatorLink) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *PaginatorLink) UnsetUrl() {
	o.Url.Unset()
}

// GetLabel returns the Label field value
func (o *PaginatorLink) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PaginatorLink) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PaginatorLink) SetLabel(v string) {
	o.Label = v
}

// GetActive returns the Active field value
func (o *PaginatorLink) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *PaginatorLink) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *PaginatorLink) SetActive(v bool) {
	o.Active = v
}

func (o PaginatorLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatorLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	toSerialize["label"] = o.Label
	toSerialize["active"] = o.Active

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatorLink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"active",
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

	varPaginatorLink := _PaginatorLink{}

	err = json.Unmarshal(data, &varPaginatorLink)

	if err != nil {
		return err
	}

	*o = PaginatorLink(varPaginatorLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "label")
		delete(additionalProperties, "active")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatorLink struct {
	value *PaginatorLink
	isSet bool
}

func (v NullablePaginatorLink) Get() *PaginatorLink {
	return v.value
}

func (v *NullablePaginatorLink) Set(val *PaginatorLink) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatorLink) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatorLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatorLink(val *PaginatorLink) *NullablePaginatorLink {
	return &NullablePaginatorLink{value: val, isSet: true}
}

func (v NullablePaginatorLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatorLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


