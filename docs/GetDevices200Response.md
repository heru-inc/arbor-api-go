# GetDevices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetDevices200ResponseDataInner**](GetDevices200ResponseDataInner.md) |  | [optional] 
**Links** | Pointer to [**GetDevices200ResponseLinks**](GetDevices200ResponseLinks.md) |  | [optional] 
**Meta** | Pointer to [**GetDevices200ResponseMeta**](GetDevices200ResponseMeta.md) |  | [optional] 

## Methods

### NewGetDevices200Response

`func NewGetDevices200Response() *GetDevices200Response`

NewGetDevices200Response instantiates a new GetDevices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDevices200ResponseWithDefaults

`func NewGetDevices200ResponseWithDefaults() *GetDevices200Response`

NewGetDevices200ResponseWithDefaults instantiates a new GetDevices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetDevices200Response) GetData() []GetDevices200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDevices200Response) GetDataOk() (*[]GetDevices200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDevices200Response) SetData(v []GetDevices200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetDevices200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetDevices200Response) GetLinks() GetDevices200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetDevices200Response) GetLinksOk() (*GetDevices200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetDevices200Response) SetLinks(v GetDevices200ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetDevices200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *GetDevices200Response) GetMeta() GetDevices200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDevices200Response) GetMetaOk() (*GetDevices200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDevices200Response) SetMeta(v GetDevices200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetDevices200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


