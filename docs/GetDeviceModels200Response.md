# GetDeviceModels200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetApps200ResponseDataInnerDeviceModelsInner**](GetApps200ResponseDataInnerDeviceModelsInner.md) |  | [optional] 
**Links** | Pointer to [**GetDeviceModels200ResponseLinks**](GetDeviceModels200ResponseLinks.md) |  | [optional] 
**Meta** | Pointer to [**GetDeviceModels200ResponseMeta**](GetDeviceModels200ResponseMeta.md) |  | [optional] 

## Methods

### NewGetDeviceModels200Response

`func NewGetDeviceModels200Response() *GetDeviceModels200Response`

NewGetDeviceModels200Response instantiates a new GetDeviceModels200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceModels200ResponseWithDefaults

`func NewGetDeviceModels200ResponseWithDefaults() *GetDeviceModels200Response`

NewGetDeviceModels200ResponseWithDefaults instantiates a new GetDeviceModels200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetDeviceModels200Response) GetData() []GetApps200ResponseDataInnerDeviceModelsInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDeviceModels200Response) GetDataOk() (*[]GetApps200ResponseDataInnerDeviceModelsInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDeviceModels200Response) SetData(v []GetApps200ResponseDataInnerDeviceModelsInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetDeviceModels200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetDeviceModels200Response) GetLinks() GetDeviceModels200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetDeviceModels200Response) GetLinksOk() (*GetDeviceModels200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetDeviceModels200Response) SetLinks(v GetDeviceModels200ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetDeviceModels200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *GetDeviceModels200Response) GetMeta() GetDeviceModels200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDeviceModels200Response) GetMetaOk() (*GetDeviceModels200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDeviceModels200Response) SetMeta(v GetDeviceModels200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetDeviceModels200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


