# UpdateDeviceValidationErrorBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Errors** | [**UpdateDeviceValidationErrorBodyErrors**](UpdateDeviceValidationErrorBodyErrors.md) |  | 

## Methods

### NewUpdateDeviceValidationErrorBody

`func NewUpdateDeviceValidationErrorBody(message string, errors UpdateDeviceValidationErrorBodyErrors, ) *UpdateDeviceValidationErrorBody`

NewUpdateDeviceValidationErrorBody instantiates a new UpdateDeviceValidationErrorBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceValidationErrorBodyWithDefaults

`func NewUpdateDeviceValidationErrorBodyWithDefaults() *UpdateDeviceValidationErrorBody`

NewUpdateDeviceValidationErrorBodyWithDefaults instantiates a new UpdateDeviceValidationErrorBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateDeviceValidationErrorBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateDeviceValidationErrorBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateDeviceValidationErrorBody) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *UpdateDeviceValidationErrorBody) GetErrors() UpdateDeviceValidationErrorBodyErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UpdateDeviceValidationErrorBody) GetErrorsOk() (*UpdateDeviceValidationErrorBodyErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UpdateDeviceValidationErrorBody) SetErrors(v UpdateDeviceValidationErrorBodyErrors)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


