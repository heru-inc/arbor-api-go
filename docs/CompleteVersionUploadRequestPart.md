# CompleteVersionUploadRequestPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartNumber** | **int32** |  | 
**ETag** | **string** |  | 

## Methods

### NewCompleteVersionUploadRequestPart

`func NewCompleteVersionUploadRequestPart(partNumber int32, eTag string, ) *CompleteVersionUploadRequestPart`

NewCompleteVersionUploadRequestPart instantiates a new CompleteVersionUploadRequestPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteVersionUploadRequestPartWithDefaults

`func NewCompleteVersionUploadRequestPartWithDefaults() *CompleteVersionUploadRequestPart`

NewCompleteVersionUploadRequestPartWithDefaults instantiates a new CompleteVersionUploadRequestPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartNumber

`func (o *CompleteVersionUploadRequestPart) GetPartNumber() int32`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *CompleteVersionUploadRequestPart) GetPartNumberOk() (*int32, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *CompleteVersionUploadRequestPart) SetPartNumber(v int32)`

SetPartNumber sets PartNumber field to given value.


### GetETag

`func (o *CompleteVersionUploadRequestPart) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *CompleteVersionUploadRequestPart) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *CompleteVersionUploadRequestPart) SetETag(v string)`

SetETag sets ETag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


