# AppsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppWithDeviceModels**](AppWithDeviceModels.md) |  | 
**Links** | [**Links**](Links.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewAppsResponse

`func NewAppsResponse(data []AppWithDeviceModels, links Links, meta Meta, ) *AppsResponse`

NewAppsResponse instantiates a new AppsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsResponseWithDefaults

`func NewAppsResponseWithDefaults() *AppsResponse`

NewAppsResponseWithDefaults instantiates a new AppsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppsResponse) GetData() []AppWithDeviceModels`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppsResponse) GetDataOk() (*[]AppWithDeviceModels, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppsResponse) SetData(v []AppWithDeviceModels)`

SetData sets Data field to given value.


### GetLinks

`func (o *AppsResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppsResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppsResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppsResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppsResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppsResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


