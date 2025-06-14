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

// checks if the GetAppReleaseChannels200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAppReleaseChannels200Response{}

// GetAppReleaseChannels200Response Get a paginated list of release channels.
type GetAppReleaseChannels200Response struct {
	Data []GetAppReleaseChannels200ResponseDataInner `json:"data,omitempty"`
	Links *GetAppReleaseChannels200ResponseLinks `json:"links,omitempty"`
	Meta *GetAppReleaseChannels200ResponseMeta `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAppReleaseChannels200Response GetAppReleaseChannels200Response

// NewGetAppReleaseChannels200Response instantiates a new GetAppReleaseChannels200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAppReleaseChannels200Response() *GetAppReleaseChannels200Response {
	this := GetAppReleaseChannels200Response{}
	return &this
}

// NewGetAppReleaseChannels200ResponseWithDefaults instantiates a new GetAppReleaseChannels200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAppReleaseChannels200ResponseWithDefaults() *GetAppReleaseChannels200Response {
	this := GetAppReleaseChannels200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAppReleaseChannels200Response) GetData() []GetAppReleaseChannels200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetAppReleaseChannels200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppReleaseChannels200Response) GetDataOk() ([]GetAppReleaseChannels200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAppReleaseChannels200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetAppReleaseChannels200ResponseDataInner and assigns it to the Data field.
func (o *GetAppReleaseChannels200Response) SetData(v []GetAppReleaseChannels200ResponseDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetAppReleaseChannels200Response) GetLinks() GetAppReleaseChannels200ResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret GetAppReleaseChannels200ResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppReleaseChannels200Response) GetLinksOk() (*GetAppReleaseChannels200ResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetAppReleaseChannels200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GetAppReleaseChannels200ResponseLinks and assigns it to the Links field.
func (o *GetAppReleaseChannels200Response) SetLinks(v GetAppReleaseChannels200ResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetAppReleaseChannels200Response) GetMeta() GetAppReleaseChannels200ResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret GetAppReleaseChannels200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppReleaseChannels200Response) GetMetaOk() (*GetAppReleaseChannels200ResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetAppReleaseChannels200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given GetAppReleaseChannels200ResponseMeta and assigns it to the Meta field.
func (o *GetAppReleaseChannels200Response) SetMeta(v GetAppReleaseChannels200ResponseMeta) {
	o.Meta = &v
}

func (o GetAppReleaseChannels200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAppReleaseChannels200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAppReleaseChannels200Response) UnmarshalJSON(data []byte) (err error) {
	varGetAppReleaseChannels200Response := _GetAppReleaseChannels200Response{}

	err = json.Unmarshal(data, &varGetAppReleaseChannels200Response)

	if err != nil {
		return err
	}

	*o = GetAppReleaseChannels200Response(varGetAppReleaseChannels200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "links")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAppReleaseChannels200Response struct {
	value *GetAppReleaseChannels200Response
	isSet bool
}

func (v NullableGetAppReleaseChannels200Response) Get() *GetAppReleaseChannels200Response {
	return v.value
}

func (v *NullableGetAppReleaseChannels200Response) Set(val *GetAppReleaseChannels200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAppReleaseChannels200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAppReleaseChannels200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAppReleaseChannels200Response(val *GetAppReleaseChannels200Response) *NullableGetAppReleaseChannels200Response {
	return &NullableGetAppReleaseChannels200Response{value: val, isSet: true}
}

func (v NullableGetAppReleaseChannels200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAppReleaseChannels200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


