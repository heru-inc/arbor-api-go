# UpdateAppValidationErrorBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Errors** | [**UpdateAppValidationErrorBodyErrors**](UpdateAppValidationErrorBodyErrors.md) |  | 

## Methods

### NewUpdateAppValidationErrorBody

`func NewUpdateAppValidationErrorBody(message string, errors UpdateAppValidationErrorBodyErrors, ) *UpdateAppValidationErrorBody`

NewUpdateAppValidationErrorBody instantiates a new UpdateAppValidationErrorBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppValidationErrorBodyWithDefaults

`func NewUpdateAppValidationErrorBodyWithDefaults() *UpdateAppValidationErrorBody`

NewUpdateAppValidationErrorBodyWithDefaults instantiates a new UpdateAppValidationErrorBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateAppValidationErrorBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateAppValidationErrorBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateAppValidationErrorBody) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *UpdateAppValidationErrorBody) GetErrors() UpdateAppValidationErrorBodyErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UpdateAppValidationErrorBody) GetErrorsOk() (*UpdateAppValidationErrorBodyErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UpdateAppValidationErrorBody) SetErrors(v UpdateAppValidationErrorBodyErrors)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


