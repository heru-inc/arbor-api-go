# PreSignFileUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**UploadId** | **string** |  | 
**PartNumbers** | **[]int32** |  | 

## Methods

### NewPreSignFileUploadRequest

`func NewPreSignFileUploadRequest(key string, uploadId string, partNumbers []int32, ) *PreSignFileUploadRequest`

NewPreSignFileUploadRequest instantiates a new PreSignFileUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreSignFileUploadRequestWithDefaults

`func NewPreSignFileUploadRequestWithDefaults() *PreSignFileUploadRequest`

NewPreSignFileUploadRequestWithDefaults instantiates a new PreSignFileUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PreSignFileUploadRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PreSignFileUploadRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PreSignFileUploadRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetUploadId

`func (o *PreSignFileUploadRequest) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *PreSignFileUploadRequest) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *PreSignFileUploadRequest) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.


### GetPartNumbers

`func (o *PreSignFileUploadRequest) GetPartNumbers() []int32`

GetPartNumbers returns the PartNumbers field if non-nil, zero value otherwise.

### GetPartNumbersOk

`func (o *PreSignFileUploadRequest) GetPartNumbersOk() (*[]int32, bool)`

GetPartNumbersOk returns a tuple with the PartNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumbers

`func (o *PreSignFileUploadRequest) SetPartNumbers(v []int32)`

SetPartNumbers sets PartNumbers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


