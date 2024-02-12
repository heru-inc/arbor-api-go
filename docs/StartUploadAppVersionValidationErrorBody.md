# StartUploadAppVersionValidationErrorBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Errors** | [**StartUploadAppVersionValidationErrorBodyErrors**](StartUploadAppVersionValidationErrorBodyErrors.md) |  | 

## Methods

### NewStartUploadAppVersionValidationErrorBody

`func NewStartUploadAppVersionValidationErrorBody(message string, errors StartUploadAppVersionValidationErrorBodyErrors, ) *StartUploadAppVersionValidationErrorBody`

NewStartUploadAppVersionValidationErrorBody instantiates a new StartUploadAppVersionValidationErrorBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartUploadAppVersionValidationErrorBodyWithDefaults

`func NewStartUploadAppVersionValidationErrorBodyWithDefaults() *StartUploadAppVersionValidationErrorBody`

NewStartUploadAppVersionValidationErrorBodyWithDefaults instantiates a new StartUploadAppVersionValidationErrorBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *StartUploadAppVersionValidationErrorBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StartUploadAppVersionValidationErrorBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StartUploadAppVersionValidationErrorBody) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *StartUploadAppVersionValidationErrorBody) GetErrors() StartUploadAppVersionValidationErrorBodyErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *StartUploadAppVersionValidationErrorBody) GetErrorsOk() (*StartUploadAppVersionValidationErrorBodyErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *StartUploadAppVersionValidationErrorBody) SetErrors(v StartUploadAppVersionValidationErrorBodyErrors)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


