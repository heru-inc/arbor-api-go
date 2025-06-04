# GetAppVersions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetAppVersions200ResponseDataInner**](GetAppVersions200ResponseDataInner.md) |  | [optional] 
**Links** | Pointer to [**GetAppVersions200ResponseLinks**](GetAppVersions200ResponseLinks.md) |  | [optional] 
**Meta** | Pointer to [**GetAppVersions200ResponseMeta**](GetAppVersions200ResponseMeta.md) |  | [optional] 

## Methods

### NewGetAppVersions200Response

`func NewGetAppVersions200Response() *GetAppVersions200Response`

NewGetAppVersions200Response instantiates a new GetAppVersions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppVersions200ResponseWithDefaults

`func NewGetAppVersions200ResponseWithDefaults() *GetAppVersions200Response`

NewGetAppVersions200ResponseWithDefaults instantiates a new GetAppVersions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAppVersions200Response) GetData() []GetAppVersions200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAppVersions200Response) GetDataOk() (*[]GetAppVersions200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAppVersions200Response) SetData(v []GetAppVersions200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetAppVersions200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetAppVersions200Response) GetLinks() GetAppVersions200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAppVersions200Response) GetLinksOk() (*GetAppVersions200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAppVersions200Response) SetLinks(v GetAppVersions200ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetAppVersions200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *GetAppVersions200Response) GetMeta() GetAppVersions200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAppVersions200Response) GetMetaOk() (*GetAppVersions200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAppVersions200Response) SetMeta(v GetAppVersions200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetAppVersions200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


