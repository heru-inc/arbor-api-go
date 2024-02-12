# FilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]File**](File.md) |  | 
**Links** | [**Links**](Links.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewFilesResponse

`func NewFilesResponse(data []File, links Links, meta Meta, ) *FilesResponse`

NewFilesResponse instantiates a new FilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesResponseWithDefaults

`func NewFilesResponseWithDefaults() *FilesResponse`

NewFilesResponseWithDefaults instantiates a new FilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FilesResponse) GetData() []File`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FilesResponse) GetDataOk() (*[]File, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FilesResponse) SetData(v []File)`

SetData sets Data field to given value.


### GetLinks

`func (o *FilesResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FilesResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FilesResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *FilesResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FilesResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FilesResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


