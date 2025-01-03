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
**LastCommunicatedAt** | **Time** |  | 
**IsOnline** | **bool** |  | 
**ClientVersion** | **string** |  | 
**LauncherVersion** | **string** |  | 
**EnrollmentDate** | **Time** |  | 
**SystemVersion** | **string** |  | 
**OsVersion** | **string** |  | 
**Ssid** | **string** |  | 
**MacAddress** | **string** |  | 
**RandomizedMacAddress** | **string** |  | 
**StorageSpaceFreeGb** | **float64** |  | 
**StorageSpaceTotalGb** | **float64** |  | 
**BatteryHealth** | **string** |  | 
**BatteryCharging** | **bool** |  | 
**BatteryPercentage** | **int32** |  | 
**BatteryTemperatureC** | **float64** |  | 
**IpAddress** | **string** |  | 
**SignalStrength** | **int32** |  | 
**FrequencyMhz** | **int32** |  | 
**LinkSpeedMbps** | **int32** |  | 
**LastLocationLatitude** | **NullableString** |  | 
**LastLocationLongitude** | **NullableString** |  | 
**LastLocationAt** | **NullableTime** |  | 

## Methods

### NewDevice

`func NewDevice(id string, title string, serialNumber string, model string, tags []string, lastCommunicatedAt Time, isOnline bool, clientVersion string, launcherVersion string, enrollmentDate Time, systemVersion string, osVersion string, ssid string, macAddress string, randomizedMacAddress string, storageSpaceFreeGb float64, storageSpaceTotalGb float64, batteryHealth string, batteryCharging bool, batteryPercentage int32, batteryTemperatureC float64, ipAddress string, signalStrength int32, frequencyMhz int32, linkSpeedMbps int32, lastLocationLatitude NullableString, lastLocationLongitude NullableString, lastLocationAt NullableTime, ) *Device`

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

`func (o *Device) GetLastCommunicatedAt() Time`

GetLastCommunicatedAt returns the LastCommunicatedAt field if non-nil, zero value otherwise.

### GetLastCommunicatedAtOk

`func (o *Device) GetLastCommunicatedAtOk() (*Time, bool)`

GetLastCommunicatedAtOk returns a tuple with the LastCommunicatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommunicatedAt

`func (o *Device) SetLastCommunicatedAt(v Time)`

SetLastCommunicatedAt sets LastCommunicatedAt field to given value.


### GetIsOnline

`func (o *Device) GetIsOnline() bool`

GetIsOnline returns the IsOnline field if non-nil, zero value otherwise.

### GetIsOnlineOk

`func (o *Device) GetIsOnlineOk() (*bool, bool)`

GetIsOnlineOk returns a tuple with the IsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnline

`func (o *Device) SetIsOnline(v bool)`

SetIsOnline sets IsOnline field to given value.


### GetClientVersion

`func (o *Device) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *Device) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *Device) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.


### GetLauncherVersion

`func (o *Device) GetLauncherVersion() string`

GetLauncherVersion returns the LauncherVersion field if non-nil, zero value otherwise.

### GetLauncherVersionOk

`func (o *Device) GetLauncherVersionOk() (*string, bool)`

GetLauncherVersionOk returns a tuple with the LauncherVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLauncherVersion

`func (o *Device) SetLauncherVersion(v string)`

SetLauncherVersion sets LauncherVersion field to given value.


### GetEnrollmentDate

`func (o *Device) GetEnrollmentDate() Time`

GetEnrollmentDate returns the EnrollmentDate field if non-nil, zero value otherwise.

### GetEnrollmentDateOk

`func (o *Device) GetEnrollmentDateOk() (*Time, bool)`

GetEnrollmentDateOk returns a tuple with the EnrollmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentDate

`func (o *Device) SetEnrollmentDate(v Time)`

SetEnrollmentDate sets EnrollmentDate field to given value.


### GetSystemVersion

`func (o *Device) GetSystemVersion() string`

GetSystemVersion returns the SystemVersion field if non-nil, zero value otherwise.

### GetSystemVersionOk

`func (o *Device) GetSystemVersionOk() (*string, bool)`

GetSystemVersionOk returns a tuple with the SystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemVersion

`func (o *Device) SetSystemVersion(v string)`

SetSystemVersion sets SystemVersion field to given value.


### GetOsVersion

`func (o *Device) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *Device) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *Device) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetSsid

`func (o *Device) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *Device) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *Device) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetMacAddress

`func (o *Device) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Device) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *Device) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetRandomizedMacAddress

`func (o *Device) GetRandomizedMacAddress() string`

GetRandomizedMacAddress returns the RandomizedMacAddress field if non-nil, zero value otherwise.

### GetRandomizedMacAddressOk

`func (o *Device) GetRandomizedMacAddressOk() (*string, bool)`

GetRandomizedMacAddressOk returns a tuple with the RandomizedMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizedMacAddress

`func (o *Device) SetRandomizedMacAddress(v string)`

SetRandomizedMacAddress sets RandomizedMacAddress field to given value.


### GetStorageSpaceFreeGb

`func (o *Device) GetStorageSpaceFreeGb() float64`

GetStorageSpaceFreeGb returns the StorageSpaceFreeGb field if non-nil, zero value otherwise.

### GetStorageSpaceFreeGbOk

`func (o *Device) GetStorageSpaceFreeGbOk() (*float64, bool)`

GetStorageSpaceFreeGbOk returns a tuple with the StorageSpaceFreeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpaceFreeGb

`func (o *Device) SetStorageSpaceFreeGb(v float64)`

SetStorageSpaceFreeGb sets StorageSpaceFreeGb field to given value.


### GetStorageSpaceTotalGb

`func (o *Device) GetStorageSpaceTotalGb() float64`

GetStorageSpaceTotalGb returns the StorageSpaceTotalGb field if non-nil, zero value otherwise.

### GetStorageSpaceTotalGbOk

`func (o *Device) GetStorageSpaceTotalGbOk() (*float64, bool)`

GetStorageSpaceTotalGbOk returns a tuple with the StorageSpaceTotalGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpaceTotalGb

`func (o *Device) SetStorageSpaceTotalGb(v float64)`

SetStorageSpaceTotalGb sets StorageSpaceTotalGb field to given value.


### GetBatteryHealth

`func (o *Device) GetBatteryHealth() string`

GetBatteryHealth returns the BatteryHealth field if non-nil, zero value otherwise.

### GetBatteryHealthOk

`func (o *Device) GetBatteryHealthOk() (*string, bool)`

GetBatteryHealthOk returns a tuple with the BatteryHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryHealth

`func (o *Device) SetBatteryHealth(v string)`

SetBatteryHealth sets BatteryHealth field to given value.


### GetBatteryCharging

`func (o *Device) GetBatteryCharging() bool`

GetBatteryCharging returns the BatteryCharging field if non-nil, zero value otherwise.

### GetBatteryChargingOk

`func (o *Device) GetBatteryChargingOk() (*bool, bool)`

GetBatteryChargingOk returns a tuple with the BatteryCharging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryCharging

`func (o *Device) SetBatteryCharging(v bool)`

SetBatteryCharging sets BatteryCharging field to given value.


### GetBatteryPercentage

`func (o *Device) GetBatteryPercentage() int32`

GetBatteryPercentage returns the BatteryPercentage field if non-nil, zero value otherwise.

### GetBatteryPercentageOk

`func (o *Device) GetBatteryPercentageOk() (*int32, bool)`

GetBatteryPercentageOk returns a tuple with the BatteryPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryPercentage

`func (o *Device) SetBatteryPercentage(v int32)`

SetBatteryPercentage sets BatteryPercentage field to given value.


### GetBatteryTemperatureC

`func (o *Device) GetBatteryTemperatureC() float64`

GetBatteryTemperatureC returns the BatteryTemperatureC field if non-nil, zero value otherwise.

### GetBatteryTemperatureCOk

`func (o *Device) GetBatteryTemperatureCOk() (*float64, bool)`

GetBatteryTemperatureCOk returns a tuple with the BatteryTemperatureC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryTemperatureC

`func (o *Device) SetBatteryTemperatureC(v float64)`

SetBatteryTemperatureC sets BatteryTemperatureC field to given value.


### GetIpAddress

`func (o *Device) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Device) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Device) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetSignalStrength

`func (o *Device) GetSignalStrength() int32`

GetSignalStrength returns the SignalStrength field if non-nil, zero value otherwise.

### GetSignalStrengthOk

`func (o *Device) GetSignalStrengthOk() (*int32, bool)`

GetSignalStrengthOk returns a tuple with the SignalStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalStrength

`func (o *Device) SetSignalStrength(v int32)`

SetSignalStrength sets SignalStrength field to given value.


### GetFrequencyMhz

`func (o *Device) GetFrequencyMhz() int32`

GetFrequencyMhz returns the FrequencyMhz field if non-nil, zero value otherwise.

### GetFrequencyMhzOk

`func (o *Device) GetFrequencyMhzOk() (*int32, bool)`

GetFrequencyMhzOk returns a tuple with the FrequencyMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyMhz

`func (o *Device) SetFrequencyMhz(v int32)`

SetFrequencyMhz sets FrequencyMhz field to given value.


### GetLinkSpeedMbps

`func (o *Device) GetLinkSpeedMbps() int32`

GetLinkSpeedMbps returns the LinkSpeedMbps field if non-nil, zero value otherwise.

### GetLinkSpeedMbpsOk

`func (o *Device) GetLinkSpeedMbpsOk() (*int32, bool)`

GetLinkSpeedMbpsOk returns a tuple with the LinkSpeedMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeedMbps

`func (o *Device) SetLinkSpeedMbps(v int32)`

SetLinkSpeedMbps sets LinkSpeedMbps field to given value.


### GetLastLocationLatitude

`func (o *Device) GetLastLocationLatitude() string`

GetLastLocationLatitude returns the LastLocationLatitude field if non-nil, zero value otherwise.

### GetLastLocationLatitudeOk

`func (o *Device) GetLastLocationLatitudeOk() (*string, bool)`

GetLastLocationLatitudeOk returns a tuple with the LastLocationLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocationLatitude

`func (o *Device) SetLastLocationLatitude(v string)`

SetLastLocationLatitude sets LastLocationLatitude field to given value.


### SetLastLocationLatitudeNil

`func (o *Device) SetLastLocationLatitudeNil(b bool)`

 SetLastLocationLatitudeNil sets the value for LastLocationLatitude to be an explicit nil

### UnsetLastLocationLatitude
`func (o *Device) UnsetLastLocationLatitude()`

UnsetLastLocationLatitude ensures that no value is present for LastLocationLatitude, not even an explicit nil
### GetLastLocationLongitude

`func (o *Device) GetLastLocationLongitude() string`

GetLastLocationLongitude returns the LastLocationLongitude field if non-nil, zero value otherwise.

### GetLastLocationLongitudeOk

`func (o *Device) GetLastLocationLongitudeOk() (*string, bool)`

GetLastLocationLongitudeOk returns a tuple with the LastLocationLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocationLongitude

`func (o *Device) SetLastLocationLongitude(v string)`

SetLastLocationLongitude sets LastLocationLongitude field to given value.


### SetLastLocationLongitudeNil

`func (o *Device) SetLastLocationLongitudeNil(b bool)`

 SetLastLocationLongitudeNil sets the value for LastLocationLongitude to be an explicit nil

### UnsetLastLocationLongitude
`func (o *Device) UnsetLastLocationLongitude()`

UnsetLastLocationLongitude ensures that no value is present for LastLocationLongitude, not even an explicit nil
### GetLastLocationAt

`func (o *Device) GetLastLocationAt() Time`

GetLastLocationAt returns the LastLocationAt field if non-nil, zero value otherwise.

### GetLastLocationAtOk

`func (o *Device) GetLastLocationAtOk() (*Time, bool)`

GetLastLocationAtOk returns a tuple with the LastLocationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocationAt

`func (o *Device) SetLastLocationAt(v Time)`

SetLastLocationAt sets LastLocationAt field to given value.


### SetLastLocationAtNil

`func (o *Device) SetLastLocationAtNil(b bool)`

 SetLastLocationAtNil sets the value for LastLocationAt to be an explicit nil

### UnsetLastLocationAt
`func (o *Device) UnsetLastLocationAt()`

UnsetLastLocationAt ensures that no value is present for LastLocationAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


