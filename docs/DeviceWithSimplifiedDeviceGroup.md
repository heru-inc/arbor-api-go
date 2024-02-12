# DeviceWithSimplifiedDeviceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**SerialNumber** | **string** |  | 
**DeviceGroup** | Pointer to [**SimplifiedDeviceGroup**](SimplifiedDeviceGroup.md) |  | [optional] 
**Model** | **string** |  | 
**Tags** | **[]string** |  | 
**LastCommunicatedAt** | **time.Time** |  | 
**IsOnline** | **bool** |  | 

## Methods

### NewDeviceWithSimplifiedDeviceGroup

`func NewDeviceWithSimplifiedDeviceGroup(id string, title string, serialNumber string, model string, tags []string, lastCommunicatedAt time.Time, isOnline bool, ) *DeviceWithSimplifiedDeviceGroup`

NewDeviceWithSimplifiedDeviceGroup instantiates a new DeviceWithSimplifiedDeviceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithSimplifiedDeviceGroupWithDefaults

`func NewDeviceWithSimplifiedDeviceGroupWithDefaults() *DeviceWithSimplifiedDeviceGroup`

NewDeviceWithSimplifiedDeviceGroupWithDefaults instantiates a new DeviceWithSimplifiedDeviceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceWithSimplifiedDeviceGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceWithSimplifiedDeviceGroup) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *DeviceWithSimplifiedDeviceGroup) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeviceWithSimplifiedDeviceGroup) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSerialNumber

`func (o *DeviceWithSimplifiedDeviceGroup) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DeviceWithSimplifiedDeviceGroup) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetDeviceGroup

`func (o *DeviceWithSimplifiedDeviceGroup) GetDeviceGroup() SimplifiedDeviceGroup`

GetDeviceGroup returns the DeviceGroup field if non-nil, zero value otherwise.

### GetDeviceGroupOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetDeviceGroupOk() (*SimplifiedDeviceGroup, bool)`

GetDeviceGroupOk returns a tuple with the DeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroup

`func (o *DeviceWithSimplifiedDeviceGroup) SetDeviceGroup(v SimplifiedDeviceGroup)`

SetDeviceGroup sets DeviceGroup field to given value.

### HasDeviceGroup

`func (o *DeviceWithSimplifiedDeviceGroup) HasDeviceGroup() bool`

HasDeviceGroup returns a boolean if a field has been set.

### GetModel

`func (o *DeviceWithSimplifiedDeviceGroup) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceWithSimplifiedDeviceGroup) SetModel(v string)`

SetModel sets Model field to given value.


### GetTags

`func (o *DeviceWithSimplifiedDeviceGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceWithSimplifiedDeviceGroup) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetLastCommunicatedAt

`func (o *DeviceWithSimplifiedDeviceGroup) GetLastCommunicatedAt() time.Time`

GetLastCommunicatedAt returns the LastCommunicatedAt field if non-nil, zero value otherwise.

### GetLastCommunicatedAtOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetLastCommunicatedAtOk() (*time.Time, bool)`

GetLastCommunicatedAtOk returns a tuple with the LastCommunicatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommunicatedAt

`func (o *DeviceWithSimplifiedDeviceGroup) SetLastCommunicatedAt(v time.Time)`

SetLastCommunicatedAt sets LastCommunicatedAt field to given value.


### GetIsOnline

`func (o *DeviceWithSimplifiedDeviceGroup) GetIsOnline() bool`

GetIsOnline returns the IsOnline field if non-nil, zero value otherwise.

### GetIsOnlineOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetIsOnlineOk() (*bool, bool)`

GetIsOnlineOk returns a tuple with the IsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnline

`func (o *DeviceWithSimplifiedDeviceGroup) SetIsOnline(v bool)`

SetIsOnline sets IsOnline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


