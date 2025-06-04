# UpdateDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**DeviceGroupId** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateDeviceRequest

`func NewUpdateDeviceRequest() *UpdateDeviceRequest`

NewUpdateDeviceRequest instantiates a new UpdateDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceRequestWithDefaults

`func NewUpdateDeviceRequestWithDefaults() *UpdateDeviceRequest`

NewUpdateDeviceRequestWithDefaults instantiates a new UpdateDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateDeviceRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateDeviceRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateDeviceRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateDeviceRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDeviceGroupId

`func (o *UpdateDeviceRequest) GetDeviceGroupId() string`

GetDeviceGroupId returns the DeviceGroupId field if non-nil, zero value otherwise.

### GetDeviceGroupIdOk

`func (o *UpdateDeviceRequest) GetDeviceGroupIdOk() (*string, bool)`

GetDeviceGroupIdOk returns a tuple with the DeviceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroupId

`func (o *UpdateDeviceRequest) SetDeviceGroupId(v string)`

SetDeviceGroupId sets DeviceGroupId field to given value.

### HasDeviceGroupId

`func (o *UpdateDeviceRequest) HasDeviceGroupId() bool`

HasDeviceGroupId returns a boolean if a field has been set.

### SetDeviceGroupIdNil

`func (o *UpdateDeviceRequest) SetDeviceGroupIdNil(b bool)`

 SetDeviceGroupIdNil sets the value for DeviceGroupId to be an explicit nil

### UnsetDeviceGroupId
`func (o *UpdateDeviceRequest) UnsetDeviceGroupId()`

UnsetDeviceGroupId ensures that no value is present for DeviceGroupId, not even an explicit nil
### GetTags

`func (o *UpdateDeviceRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateDeviceRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateDeviceRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateDeviceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


