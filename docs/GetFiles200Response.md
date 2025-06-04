# GetFiles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetFiles200ResponseDataInner**](GetFiles200ResponseDataInner.md) |  | [optional] 
**Links** | Pointer to [**GetFiles200ResponseLinks**](GetFiles200ResponseLinks.md) |  | [optional] 
**Meta** | Pointer to [**GetFiles200ResponseMeta**](GetFiles200ResponseMeta.md) |  | [optional] 

## Methods

### NewGetFiles200Response

`func NewGetFiles200Response() *GetFiles200Response`

NewGetFiles200Response instantiates a new GetFiles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiles200ResponseWithDefaults

`func NewGetFiles200ResponseWithDefaults() *GetFiles200Response`

NewGetFiles200ResponseWithDefaults instantiates a new GetFiles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetFiles200Response) GetData() []GetFiles200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetFiles200Response) GetDataOk() (*[]GetFiles200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetFiles200Response) SetData(v []GetFiles200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetFiles200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetFiles200Response) GetLinks() GetFiles200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetFiles200Response) GetLinksOk() (*GetFiles200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetFiles200Response) SetLinks(v GetFiles200ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetFiles200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *GetFiles200Response) GetMeta() GetFiles200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetFiles200Response) GetMetaOk() (*GetFiles200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetFiles200Response) SetMeta(v GetFiles200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetFiles200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


