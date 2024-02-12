# PaginatorLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** |  | [optional] 
**Label** | **string** |  | 
**Active** | **bool** |  | 

## Methods

### NewPaginatorLink

`func NewPaginatorLink(label string, active bool, ) *PaginatorLink`

NewPaginatorLink instantiates a new PaginatorLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatorLinkWithDefaults

`func NewPaginatorLinkWithDefaults() *PaginatorLink`

NewPaginatorLinkWithDefaults instantiates a new PaginatorLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PaginatorLink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PaginatorLink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PaginatorLink) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PaginatorLink) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PaginatorLink) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PaginatorLink) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetLabel

`func (o *PaginatorLink) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PaginatorLink) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PaginatorLink) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetActive

`func (o *PaginatorLink) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PaginatorLink) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PaginatorLink) SetActive(v bool)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


