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

// checks if the Meta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Meta{}

// Meta Metadata about a paginated response.
type Meta struct {
	CurrentPage int32 `json:"current_page"`
	From int32 `json:"from"`
	LastPage int32 `json:"last_page"`
	Links []PaginatorLink `json:"links"`
	Path string `json:"path"`
	PerPage int32 `json:"per_page"`
	To int32 `json:"to"`
	Total int32 `json:"total"`
}

type _Meta Meta

// NewMeta instantiates a new Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeta(currentPage int32, from int32, lastPage int32, links []PaginatorLink, path string, perPage int32, to int32, total int32) *Meta {
	this := Meta{}
	this.CurrentPage = currentPage
	this.From = from
	this.LastPage = lastPage
	this.Links = links
	this.Path = path
	this.PerPage = perPage
	this.To = to
	this.Total = total
	return &this
}

// NewMetaWithDefaults instantiates a new Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaWithDefaults() *Meta {
	this := Meta{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value
func (o *Meta) GetCurrentPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value
// and a boolean to check if the value has been set.
func (o *Meta) GetCurrentPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPage, true
}

// SetCurrentPage sets field value
func (o *Meta) SetCurrentPage(v int32) {
	o.CurrentPage = v
}

// GetFrom returns the From field value
func (o *Meta) GetFrom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *Meta) GetFromOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *Meta) SetFrom(v int32) {
	o.From = v
}

// GetLastPage returns the LastPage field value
func (o *Meta) GetLastPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LastPage
}

// GetLastPageOk returns a tuple with the LastPage field value
// and a boolean to check if the value has been set.
func (o *Meta) GetLastPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastPage, true
}

// SetLastPage sets field value
func (o *Meta) SetLastPage(v int32) {
	o.LastPage = v
}

// GetLinks returns the Links field value
func (o *Meta) GetLinks() []PaginatorLink {
	if o == nil {
		var ret []PaginatorLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *Meta) GetLinksOk() ([]PaginatorLink, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *Meta) SetLinks(v []PaginatorLink) {
	o.Links = v
}

// GetPath returns the Path field value
func (o *Meta) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *Meta) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *Meta) SetPath(v string) {
	o.Path = v
}

// GetPerPage returns the PerPage field value
func (o *Meta) GetPerPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value
// and a boolean to check if the value has been set.
func (o *Meta) GetPerPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerPage, true
}

// SetPerPage sets field value
func (o *Meta) SetPerPage(v int32) {
	o.PerPage = v
}

// GetTo returns the To field value
func (o *Meta) GetTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *Meta) GetToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *Meta) SetTo(v int32) {
	o.To = v
}

// GetTotal returns the Total field value
func (o *Meta) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *Meta) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *Meta) SetTotal(v int32) {
	o.Total = v
}

func (o Meta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Meta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current_page"] = o.CurrentPage
	toSerialize["from"] = o.From
	toSerialize["last_page"] = o.LastPage
	toSerialize["links"] = o.Links
	toSerialize["path"] = o.Path
	toSerialize["per_page"] = o.PerPage
	toSerialize["to"] = o.To
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

func (o *Meta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"current_page",
		"from",
		"last_page",
		"links",
		"path",
		"per_page",
		"to",
		"total",
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

	varMeta := _Meta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMeta)

	if err != nil {
		return err
	}

	*o = Meta(varMeta)

	return err
}

type NullableMeta struct {
	value *Meta
	isSet bool
}

func (v NullableMeta) Get() *Meta {
	return v.value
}

func (v *NullableMeta) Set(val *Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeta(val *Meta) *NullableMeta {
	return &NullableMeta{value: val, isSet: true}
}

func (v NullableMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


