# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**SerialNumber** | **string** |  | 
**DeviceGroup** | Pointer to **NullableString** |  | [optional] 
**Model** | **string** |  | 
**Tags** | **[]string** |  | 
**LastCommunicatedAt** | **time.Time** |  | 

## Methods

### NewDevice

`func NewDevice(id string, title string, serialNumber string, model string, tags []string, lastCommunicatedAt time.Time, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *Device) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Device) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Device) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSerialNumber

`func (o *Device) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Device) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Device) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetDeviceGroup

`func (o *Device) GetDeviceGroup() string`

GetDeviceGroup returns the DeviceGroup field if non-nil, zero value otherwise.

### GetDeviceGroupOk

`func (o *Device) GetDeviceGroupOk() (*string, bool)`

GetDeviceGroupOk returns a tuple with the DeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroup

`func (o *Device) SetDeviceGroup(v string)`

SetDeviceGroup sets DeviceGroup field to given value.

### HasDeviceGroup

`func (o *Device) HasDeviceGroup() bool`

HasDeviceGroup returns a boolean if a field has been set.

### SetDeviceGroupNil

`func (o *Device) SetDeviceGroupNil(b bool)`

 SetDeviceGroupNil sets the value for DeviceGroup to be an explicit nil

### UnsetDeviceGroup
`func (o *Device) UnsetDeviceGroup()`

UnsetDeviceGroup ensures that no value is present for DeviceGroup, not even an explicit nil
### GetModel

`func (o *Device) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Device) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Device) SetModel(v string)`

SetModel sets Model field to given value.


### GetTags

`func (o *Device) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Device) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Device) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetLastCommunicatedAt

`func (o *Device) GetLastCommunicatedAt() time.Time`

GetLastCommunicatedAt returns the LastCommunicatedAt field if non-nil, zero value otherwise.

### GetLastCommunicatedAtOk

`func (o *Device) GetLastCommunicatedAtOk() (*time.Time, bool)`

GetLastCommunicatedAtOk returns a tuple with the LastCommunicatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommunicatedAt

`func (o *Device) SetLastCommunicatedAt(v time.Time)`

SetLastCommunicatedAt sets LastCommunicatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


