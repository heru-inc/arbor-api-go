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
**BatteryTemperatureC** | **int32** |  | 
**IpAddress** | **string** |  | 
**SignalStrength** | **int32** |  | 
**FrequencyMhz** | **int32** |  | 
**LinkSpeedMbps** | **int32** |  | 
**LastLocationLatitude** | **float64** |  | 
**LastLocationLongitude** | **float64** |  | 
**LastLocationAt** | **Time** |  | 

## Methods

### NewDeviceWithSimplifiedDeviceGroup

`func NewDeviceWithSimplifiedDeviceGroup(id string, title string, serialNumber string, model string, tags []string, lastCommunicatedAt Time, isOnline bool, clientVersion string, launcherVersion string, enrollmentDate Time, systemVersion string, osVersion string, ssid string, macAddress string, randomizedMacAddress string, storageSpaceFreeGb float64, storageSpaceTotalGb float64, batteryHealth string, batteryCharging bool, batteryPercentage int32, batteryTemperatureC int32, ipAddress string, signalStrength int32, frequencyMhz int32, linkSpeedMbps int32, lastLocationLatitude float64, lastLocationLongitude float64, lastLocationAt Time, ) *DeviceWithSimplifiedDeviceGroup`

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

`func (o *DeviceWithSimplifiedDeviceGroup) GetLastCommunicatedAt() Time`

GetLastCommunicatedAt returns the LastCommunicatedAt field if non-nil, zero value otherwise.

### GetLastCommunicatedAtOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetLastCommunicatedAtOk() (*Time, bool)`

GetLastCommunicatedAtOk returns a tuple with the LastCommunicatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommunicatedAt

`func (o *DeviceWithSimplifiedDeviceGroup) SetLastCommunicatedAt(v Time)`

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


### GetClientVersion

`func (o *DeviceWithSimplifiedDeviceGroup) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *DeviceWithSimplifiedDeviceGroup) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.


### GetLauncherVersion

`func (o *DeviceWithSimplifiedDeviceGroup) GetLauncherVersion() string`

GetLauncherVersion returns the LauncherVersion field if non-nil, zero value otherwise.

### GetLauncherVersionOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetLauncherVersionOk() (*string, bool)`

GetLauncherVersionOk returns a tuple with the LauncherVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLauncherVersion

`func (o *DeviceWithSimplifiedDeviceGroup) SetLauncherVersion(v string)`

SetLauncherVersion sets LauncherVersion field to given value.


### GetEnrollmentDate

`func (o *DeviceWithSimplifiedDeviceGroup) GetEnrollmentDate() Time`

GetEnrollmentDate returns the EnrollmentDate field if non-nil, zero value otherwise.

### GetEnrollmentDateOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetEnrollmentDateOk() (*Time, bool)`

GetEnrollmentDateOk returns a tuple with the EnrollmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentDate

`func (o *DeviceWithSimplifiedDeviceGroup) SetEnrollmentDate(v Time)`

SetEnrollmentDate sets EnrollmentDate field to given value.


### GetSystemVersion

`func (o *DeviceWithSimplifiedDeviceGroup) GetSystemVersion() string`

GetSystemVersion returns the SystemVersion field if non-nil, zero value otherwise.

### GetSystemVersionOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetSystemVersionOk() (*string, bool)`

GetSystemVersionOk returns a tuple with the SystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemVersion

`func (o *DeviceWithSimplifiedDeviceGroup) SetSystemVersion(v string)`

SetSystemVersion sets SystemVersion field to given value.


### GetOsVersion

`func (o *DeviceWithSimplifiedDeviceGroup) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DeviceWithSimplifiedDeviceGroup) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.


### GetSsid

`func (o *DeviceWithSimplifiedDeviceGroup) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *DeviceWithSimplifiedDeviceGroup) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetMacAddress

`func (o *DeviceWithSimplifiedDeviceGroup) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *DeviceWithSimplifiedDeviceGroup) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetRandomizedMacAddress

`func (o *DeviceWithSimplifiedDeviceGroup) GetRandomizedMacAddress() string`

GetRandomizedMacAddress returns the RandomizedMacAddress field if non-nil, zero value otherwise.

### GetRandomizedMacAddressOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetRandomizedMacAddressOk() (*string, bool)`

GetRandomizedMacAddressOk returns a tuple with the RandomizedMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizedMacAddress

`func (o *DeviceWithSimplifiedDeviceGroup) SetRandomizedMacAddress(v string)`

SetRandomizedMacAddress sets RandomizedMacAddress field to given value.


### GetStorageSpaceFreeGb

`func (o *DeviceWithSimplifiedDeviceGroup) GetStorageSpaceFreeGb() float64`

GetStorageSpaceFreeGb returns the StorageSpaceFreeGb field if non-nil, zero value otherwise.

### GetStorageSpaceFreeGbOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetStorageSpaceFreeGbOk() (*float64, bool)`

GetStorageSpaceFreeGbOk returns a tuple with the StorageSpaceFreeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpaceFreeGb

`func (o *DeviceWithSimplifiedDeviceGroup) SetStorageSpaceFreeGb(v float64)`

SetStorageSpaceFreeGb sets StorageSpaceFreeGb field to given value.


### GetStorageSpaceTotalGb

`func (o *DeviceWithSimplifiedDeviceGroup) GetStorageSpaceTotalGb() float64`

GetStorageSpaceTotalGb returns the StorageSpaceTotalGb field if non-nil, zero value otherwise.

### GetStorageSpaceTotalGbOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetStorageSpaceTotalGbOk() (*float64, bool)`

GetStorageSpaceTotalGbOk returns a tuple with the StorageSpaceTotalGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpaceTotalGb

`func (o *DeviceWithSimplifiedDeviceGroup) SetStorageSpaceTotalGb(v float64)`

SetStorageSpaceTotalGb sets StorageSpaceTotalGb field to given value.


### GetBatteryHealth

`func (o *DeviceWithSimplifiedDeviceGroup) GetBatteryHealth() string`

GetBatteryHealth returns the BatteryHealth field if non-nil, zero value otherwise.

### GetBatteryHealthOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetBatteryHealthOk() (*string, bool)`

GetBatteryHealthOk returns a tuple with the BatteryHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryHealth

`func (o *DeviceWithSimplifiedDeviceGroup) SetBatteryHealth(v string)`

SetBatteryHealth sets BatteryHealth field to given value.


### GetBatteryCharging

`func (o *DeviceWithSimplifiedDeviceGroup) GetBatteryCharging() bool`

GetBatteryCharging returns the BatteryCharging field if non-nil, zero value otherwise.

### GetBatteryChargingOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetBatteryChargingOk() (*bool, bool)`

GetBatteryChargingOk returns a tuple with the BatteryCharging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryCharging

`func (o *DeviceWithSimplifiedDeviceGroup) SetBatteryCharging(v bool)`

SetBatteryCharging sets BatteryCharging field to given value.


### GetBatteryPercentage

`func (o *DeviceWithSimplifiedDeviceGroup) GetBatteryPercentage() int32`

GetBatteryPercentage returns the BatteryPercentage field if non-nil, zero value otherwise.

### GetBatteryPercentageOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetBatteryPercentageOk() (*int32, bool)`

GetBatteryPercentageOk returns a tuple with the BatteryPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryPercentage

`func (o *DeviceWithSimplifiedDeviceGroup) SetBatteryPercentage(v int32)`

SetBatteryPercentage sets BatteryPercentage field to given value.


### GetBatteryTemperatureC

`func (o *DeviceWithSimplifiedDeviceGroup) GetBatteryTemperatureC() int32`

GetBatteryTemperatureC returns the BatteryTemperatureC field if non-nil, zero value otherwise.

### GetBatteryTemperatureCOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetBatteryTemperatureCOk() (*int32, bool)`

GetBatteryTemperatureCOk returns a tuple with the BatteryTemperatureC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryTemperatureC

`func (o *DeviceWithSimplifiedDeviceGroup) SetBatteryTemperatureC(v int32)`

SetBatteryTemperatureC sets BatteryTemperatureC field to given value.


### GetIpAddress

`func (o *DeviceWithSimplifiedDeviceGroup) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *DeviceWithSimplifiedDeviceGroup) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetSignalStrength

`func (o *DeviceWithSimplifiedDeviceGroup) GetSignalStrength() int32`

GetSignalStrength returns the SignalStrength field if non-nil, zero value otherwise.

### GetSignalStrengthOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetSignalStrengthOk() (*int32, bool)`

GetSignalStrengthOk returns a tuple with the SignalStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalStrength

`func (o *DeviceWithSimplifiedDeviceGroup) SetSignalStrength(v int32)`

SetSignalStrength sets SignalStrength field to given value.


### GetFrequencyMhz

`func (o *DeviceWithSimplifiedDeviceGroup) GetFrequencyMhz() int32`

GetFrequencyMhz returns the FrequencyMhz field if non-nil, zero value otherwise.

### GetFrequencyMhzOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetFrequencyMhzOk() (*int32, bool)`

GetFrequencyMhzOk returns a tuple with the FrequencyMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyMhz

`func (o *DeviceWithSimplifiedDeviceGroup) SetFrequencyMhz(v int32)`

SetFrequencyMhz sets FrequencyMhz field to given value.


### GetLinkSpeedMbps

`func (o *DeviceWithSimplifiedDeviceGroup) GetLinkSpeedMbps() int32`

GetLinkSpeedMbps returns the LinkSpeedMbps field if non-nil, zero value otherwise.

### GetLinkSpeedMbpsOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetLinkSpeedMbpsOk() (*int32, bool)`

GetLinkSpeedMbpsOk returns a tuple with the LinkSpeedMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeedMbps

`func (o *DeviceWithSimplifiedDeviceGroup) SetLinkSpeedMbps(v int32)`

SetLinkSpeedMbps sets LinkSpeedMbps field to given value.


### GetLastLocationLatitude

`func (o *DeviceWithSimplifiedDeviceGroup) GetLastLocationLatitude() float64`

GetLastLocationLatitude returns the LastLocationLatitude field if non-nil, zero value otherwise.

### GetLastLocationLatitudeOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetLastLocationLatitudeOk() (*float64, bool)`

GetLastLocationLatitudeOk returns a tuple with the LastLocationLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocationLatitude

`func (o *DeviceWithSimplifiedDeviceGroup) SetLastLocationLatitude(v float64)`

SetLastLocationLatitude sets LastLocationLatitude field to given value.


### GetLastLocationLongitude

`func (o *DeviceWithSimplifiedDeviceGroup) GetLastLocationLongitude() float64`

GetLastLocationLongitude returns the LastLocationLongitude field if non-nil, zero value otherwise.

### GetLastLocationLongitudeOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetLastLocationLongitudeOk() (*float64, bool)`

GetLastLocationLongitudeOk returns a tuple with the LastLocationLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocationLongitude

`func (o *DeviceWithSimplifiedDeviceGroup) SetLastLocationLongitude(v float64)`

SetLastLocationLongitude sets LastLocationLongitude field to given value.


### GetLastLocationAt

`func (o *DeviceWithSimplifiedDeviceGroup) GetLastLocationAt() Time`

GetLastLocationAt returns the LastLocationAt field if non-nil, zero value otherwise.

### GetLastLocationAtOk

`func (o *DeviceWithSimplifiedDeviceGroup) GetLastLocationAtOk() (*Time, bool)`

GetLastLocationAtOk returns a tuple with the LastLocationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocationAt

`func (o *DeviceWithSimplifiedDeviceGroup) SetLastLocationAt(v Time)`

SetLastLocationAt sets LastLocationAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


