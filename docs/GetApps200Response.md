# GetApps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetApps200ResponseDataInner**](GetApps200ResponseDataInner.md) |  | [optional] 
**Links** | Pointer to [**GetApps200ResponseLinks**](GetApps200ResponseLinks.md) |  | [optional] 
**Meta** | Pointer to [**GetApps200ResponseMeta**](GetApps200ResponseMeta.md) |  | [optional] 

## Methods

### NewGetApps200Response

`func NewGetApps200Response() *GetApps200Response`

NewGetApps200Response instantiates a new GetApps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApps200ResponseWithDefaults

`func NewGetApps200ResponseWithDefaults() *GetApps200Response`

NewGetApps200ResponseWithDefaults instantiates a new GetApps200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetApps200Response) GetData() []GetApps200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetApps200Response) GetDataOk() (*[]GetApps200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetApps200Response) SetData(v []GetApps200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetApps200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetApps200Response) GetLinks() GetApps200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetApps200Response) GetLinksOk() (*GetApps200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetApps200Response) SetLinks(v GetApps200ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetApps200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *GetApps200Response) GetMeta() GetApps200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetApps200Response) GetMetaOk() (*GetApps200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetApps200Response) SetMeta(v GetApps200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetApps200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


