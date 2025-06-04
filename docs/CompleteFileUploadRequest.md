# CompleteFileUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**UploadId** | **string** |  | 
**Parts** | [**[]CompleteAppVersionUploadRequestPartsInner**](CompleteAppVersionUploadRequestPartsInner.md) |  | 
**ConflictStrategy** | Pointer to **string** |  | [optional] 

## Methods

### NewCompleteFileUploadRequest

`func NewCompleteFileUploadRequest(key string, uploadId string, parts []CompleteAppVersionUploadRequestPartsInner, ) *CompleteFileUploadRequest`

NewCompleteFileUploadRequest instantiates a new CompleteFileUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteFileUploadRequestWithDefaults

`func NewCompleteFileUploadRequestWithDefaults() *CompleteFileUploadRequest`

NewCompleteFileUploadRequestWithDefaults instantiates a new CompleteFileUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CompleteFileUploadRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CompleteFileUploadRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CompleteFileUploadRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetUploadId

`func (o *CompleteFileUploadRequest) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *CompleteFileUploadRequest) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *CompleteFileUploadRequest) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.


### GetParts

`func (o *CompleteFileUploadRequest) GetParts() []CompleteAppVersionUploadRequestPartsInner`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *CompleteFileUploadRequest) GetPartsOk() (*[]CompleteAppVersionUploadRequestPartsInner, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *CompleteFileUploadRequest) SetParts(v []CompleteAppVersionUploadRequestPartsInner)`

SetParts sets Parts field to given value.


### GetConflictStrategy

`func (o *CompleteFileUploadRequest) GetConflictStrategy() string`

GetConflictStrategy returns the ConflictStrategy field if non-nil, zero value otherwise.

### GetConflictStrategyOk

`func (o *CompleteFileUploadRequest) GetConflictStrategyOk() (*string, bool)`

GetConflictStrategyOk returns a tuple with the ConflictStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictStrategy

`func (o *CompleteFileUploadRequest) SetConflictStrategy(v string)`

SetConflictStrategy sets ConflictStrategy field to given value.

### HasConflictStrategy

`func (o *CompleteFileUploadRequest) HasConflictStrategy() bool`

HasConflictStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


