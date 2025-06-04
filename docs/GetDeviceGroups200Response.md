# GetDeviceGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetDeviceGroups200ResponseDataInner**](GetDeviceGroups200ResponseDataInner.md) |  | [optional] 
**Links** | Pointer to [**GetDeviceGroups200ResponseLinks**](GetDeviceGroups200ResponseLinks.md) |  | [optional] 
**Meta** | Pointer to [**GetDeviceGroups200ResponseMeta**](GetDeviceGroups200ResponseMeta.md) |  | [optional] 

## Methods

### NewGetDeviceGroups200Response

`func NewGetDeviceGroups200Response() *GetDeviceGroups200Response`

NewGetDeviceGroups200Response instantiates a new GetDeviceGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceGroups200ResponseWithDefaults

`func NewGetDeviceGroups200ResponseWithDefaults() *GetDeviceGroups200Response`

NewGetDeviceGroups200ResponseWithDefaults instantiates a new GetDeviceGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetDeviceGroups200Response) GetData() []GetDeviceGroups200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDeviceGroups200Response) GetDataOk() (*[]GetDeviceGroups200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDeviceGroups200Response) SetData(v []GetDeviceGroups200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetDeviceGroups200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetDeviceGroups200Response) GetLinks() GetDeviceGroups200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetDeviceGroups200Response) GetLinksOk() (*GetDeviceGroups200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetDeviceGroups200Response) SetLinks(v GetDeviceGroups200ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetDeviceGroups200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *GetDeviceGroups200Response) GetMeta() GetDeviceGroups200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDeviceGroups200Response) GetMetaOk() (*GetDeviceGroups200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDeviceGroups200Response) SetMeta(v GetDeviceGroups200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetDeviceGroups200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


