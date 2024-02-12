# PresignUploadAppVersionValidationErrorBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Errors** | [**PresignUploadAppVersionValidationErrorBodyErrors**](PresignUploadAppVersionValidationErrorBodyErrors.md) |  | 

## Methods

### NewPresignUploadAppVersionValidationErrorBody

`func NewPresignUploadAppVersionValidationErrorBody(message string, errors PresignUploadAppVersionValidationErrorBodyErrors, ) *PresignUploadAppVersionValidationErrorBody`

NewPresignUploadAppVersionValidationErrorBody instantiates a new PresignUploadAppVersionValidationErrorBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresignUploadAppVersionValidationErrorBodyWithDefaults

`func NewPresignUploadAppVersionValidationErrorBodyWithDefaults() *PresignUploadAppVersionValidationErrorBody`

NewPresignUploadAppVersionValidationErrorBodyWithDefaults instantiates a new PresignUploadAppVersionValidationErrorBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PresignUploadAppVersionValidationErrorBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PresignUploadAppVersionValidationErrorBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PresignUploadAppVersionValidationErrorBody) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *PresignUploadAppVersionValidationErrorBody) GetErrors() PresignUploadAppVersionValidationErrorBodyErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PresignUploadAppVersionValidationErrorBody) GetErrorsOk() (*PresignUploadAppVersionValidationErrorBodyErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PresignUploadAppVersionValidationErrorBody) SetErrors(v PresignUploadAppVersionValidationErrorBodyErrors)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


