# GetDevices200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**DeviceGroup** | Pointer to **NullableString** |  | [optional] 
**Group** | Pointer to [**GetDeviceGroups200ResponseDataInner**](GetDeviceGroups200ResponseDataInner.md) |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**LastCommunicatedAt** | Pointer to **NullableString** |  | [optional] 
**IsOnline** | Pointer to **NullableBool** |  | [optional] 
**ClientVersion** | Pointer to **NullableString** |  | [optional] 
**LauncherVersion** | Pointer to **NullableString** |  | [optional] 
**EnrollmentDate** | Pointer to **NullableString** |  | [optional] 
**SystemVersion** | Pointer to **NullableString** |  | [optional] 
**OsVersion** | Pointer to **NullableString** |  | [optional] 
**Ssid** | Pointer to **NullableString** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 
**RandomizedMacAddress** | Pointer to **NullableString** |  | [optional] 
**StorageSpaceFreeGb** | Pointer to **NullableFloat32** |  | [optional] 
**StorageSpaceTotalGb** | Pointer to **NullableFloat32** |  | [optional] 
**BatteryHealth** | Pointer to **NullableString** |  | [optional] 
**BatteryCharging** | Pointer to **NullableBool** |  | [optional] 
**BatteryPercentage** | Pointer to **NullableInt32** |  | [optional] 
**BatteryTemperatureC** | Pointer to **NullableFloat32** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**SignalStrength** | Pointer to **NullableFloat32** |  | [optional] 
**FrequencyMhz** | Pointer to **NullableInt32** |  | [optional] 
**LinkSpeedMbps** | Pointer to **NullableInt32** |  | [optional] 
**LastLocationLatitude** | Pointer to **NullableFloat32** |  | [optional] 
**LastLocationLongitude** | Pointer to **NullableFloat32** |  | [optional] 
**LastLocationAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetDevices200ResponseDataInner

`func NewGetDevices200ResponseDataInner() *GetDevices200ResponseDataInner`

NewGetDevices200ResponseDataInner instantiates a new GetDevices200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDevices200ResponseDataInnerWithDefaults

`func NewGetDevices200ResponseDataInnerWithDefaults() *GetDevices200ResponseDataInner`

NewGetDevices200ResponseDataInnerWithDefaults instantiates a new GetDevices200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDevices200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDevices200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDevices200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetDevices200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *GetDevices200ResponseDataInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetDevices200ResponseDataInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetDevices200ResponseDataInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetDevices200ResponseDataInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSerialNumber

`func (o *GetDevices200ResponseDataInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *GetDevices200ResponseDataInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *GetDevices200ResponseDataInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *GetDevices200ResponseDataInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetDeviceGroup

`func (o *GetDevices200ResponseDataInner) GetDeviceGroup() string`

GetDeviceGroup returns the DeviceGroup field if non-nil, zero value otherwise.

### GetDeviceGroupOk

`func (o *GetDevices200ResponseDataInner) GetDeviceGroupOk() (*string, bool)`

GetDeviceGroupOk returns a tuple with the DeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroup

`func (o *GetDevices200ResponseDataInner) SetDeviceGroup(v string)`

SetDeviceGroup sets DeviceGroup field to given value.

### HasDeviceGroup

`func (o *GetDevices200ResponseDataInner) HasDeviceGroup() bool`

HasDeviceGroup returns a boolean if a field has been set.

### SetDeviceGroupNil

`func (o *GetDevices200ResponseDataInner) SetDeviceGroupNil(b bool)`

 SetDeviceGroupNil sets the value for DeviceGroup to be an explicit nil

### UnsetDeviceGroup
`func (o *GetDevices200ResponseDataInner) UnsetDeviceGroup()`

UnsetDeviceGroup ensures that no value is present for DeviceGroup, not even an explicit nil
### GetGroup

`func (o *GetDevices200ResponseDataInner) GetGroup() GetDeviceGroups200ResponseDataInner`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetDevices200ResponseDataInner) GetGroupOk() (*GetDeviceGroups200ResponseDataInner, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetDevices200ResponseDataInner) SetGroup(v GetDeviceGroups200ResponseDataInner)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetDevices200ResponseDataInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetModel

`func (o *GetDevices200ResponseDataInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetDevices200ResponseDataInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetDevices200ResponseDataInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetDevices200ResponseDataInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTags

`func (o *GetDevices200ResponseDataInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetDevices200ResponseDataInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetDevices200ResponseDataInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetDevices200ResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLastCommunicatedAt

`func (o *GetDevices200ResponseDataInner) GetLastCommunicatedAt() string`

GetLastCommunicatedAt returns the LastCommunicatedAt field if non-nil, zero value otherwise.

### GetLastCommunicatedAtOk

`func (o *GetDevices200ResponseDataInner) GetLastCommunicatedAtOk() (*string, bool)`

GetLastCommunicatedAtOk returns a tuple with the LastCommunicatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommunicatedAt

`func (o *GetDevices200ResponseDataInner) SetLastCommunicatedAt(v string)`

SetLastCommunicatedAt sets LastCommunicatedAt field to given value.

### HasLastCommunicatedAt

`func (o *GetDevices200ResponseDataInner) HasLastCommunicatedAt() bool`

HasLastCommunicatedAt returns a boolean if a field has been set.

### SetLastCommunicatedAtNil

`func (o *GetDevices200ResponseDataInner) SetLastCommunicatedAtNil(b bool)`

 SetLastCommunicatedAtNil sets the value for LastCommunicatedAt to be an explicit nil

### UnsetLastCommunicatedAt
`func (o *GetDevices200ResponseDataInner) UnsetLastCommunicatedAt()`

UnsetLastCommunicatedAt ensures that no value is present for LastCommunicatedAt, not even an explicit nil
### GetIsOnline

`func (o *GetDevices200ResponseDataInner) GetIsOnline() bool`

GetIsOnline returns the IsOnline field if non-nil, zero value otherwise.

### GetIsOnlineOk

`func (o *GetDevices200ResponseDataInner) GetIsOnlineOk() (*bool, bool)`

GetIsOnlineOk returns a tuple with the IsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnline

`func (o *GetDevices200ResponseDataInner) SetIsOnline(v bool)`

SetIsOnline sets IsOnline field to given value.

### HasIsOnline

`func (o *GetDevices200ResponseDataInner) HasIsOnline() bool`

HasIsOnline returns a boolean if a field has been set.

### SetIsOnlineNil

`func (o *GetDevices200ResponseDataInner) SetIsOnlineNil(b bool)`

 SetIsOnlineNil sets the value for IsOnline to be an explicit nil

### UnsetIsOnline
`func (o *GetDevices200ResponseDataInner) UnsetIsOnline()`

UnsetIsOnline ensures that no value is present for IsOnline, not even an explicit nil
### GetClientVersion

`func (o *GetDevices200ResponseDataInner) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *GetDevices200ResponseDataInner) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *GetDevices200ResponseDataInner) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.

### HasClientVersion

`func (o *GetDevices200ResponseDataInner) HasClientVersion() bool`

HasClientVersion returns a boolean if a field has been set.

### SetClientVersionNil

`func (o *GetDevices200ResponseDataInner) SetClientVersionNil(b bool)`

 SetClientVersionNil sets the value for ClientVersion to be an explicit nil

### UnsetClientVersion
`func (o *GetDevices200ResponseDataInner) UnsetClientVersion()`

UnsetClientVersion ensures that no value is present for ClientVersion, not even an explicit nil
### GetLauncherVersion

`func (o *GetDevices200ResponseDataInner) GetLauncherVersion() string`

GetLauncherVersion returns the LauncherVersion field if non-nil, zero value otherwise.

### GetLauncherVersionOk

`func (o *GetDevices200ResponseDataInner) GetLauncherVersionOk() (*string, bool)`

GetLauncherVersionOk returns a tuple with the LauncherVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLauncherVersion

`func (o *GetDevices200ResponseDataInner) SetLauncherVersion(v string)`

SetLauncherVersion sets LauncherVersion field to given value.

### HasLauncherVersion

`func (o *GetDevices200ResponseDataInner) HasLauncherVersion() bool`

HasLauncherVersion returns a boolean if a field has been set.

### SetLauncherVersionNil

`func (o *GetDevices200ResponseDataInner) SetLauncherVersionNil(b bool)`

 SetLauncherVersionNil sets the value for LauncherVersion to be an explicit nil

### UnsetLauncherVersion
`func (o *GetDevices200ResponseDataInner) UnsetLauncherVersion()`

UnsetLauncherVersion ensures that no value is present for LauncherVersion, not even an explicit nil
### GetEnrollmentDate

`func (o *GetDevices200ResponseDataInner) GetEnrollmentDate() string`

GetEnrollmentDate returns the EnrollmentDate field if non-nil, zero value otherwise.

### GetEnrollmentDateOk

`func (o *GetDevices200ResponseDataInner) GetEnrollmentDateOk() (*string, bool)`

GetEnrollmentDateOk returns a tuple with the EnrollmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentDate

`func (o *GetDevices200ResponseDataInner) SetEnrollmentDate(v string)`

SetEnrollmentDate sets EnrollmentDate field to given value.

### HasEnrollmentDate

`func (o *GetDevices200ResponseDataInner) HasEnrollmentDate() bool`

HasEnrollmentDate returns a boolean if a field has been set.

### SetEnrollmentDateNil

`func (o *GetDevices200ResponseDataInner) SetEnrollmentDateNil(b bool)`

 SetEnrollmentDateNil sets the value for EnrollmentDate to be an explicit nil

### UnsetEnrollmentDate
`func (o *GetDevices200ResponseDataInner) UnsetEnrollmentDate()`

UnsetEnrollmentDate ensures that no value is present for EnrollmentDate, not even an explicit nil
### GetSystemVersion

`func (o *GetDevices200ResponseDataInner) GetSystemVersion() string`

GetSystemVersion returns the SystemVersion field if non-nil, zero value otherwise.

### GetSystemVersionOk

`func (o *GetDevices200ResponseDataInner) GetSystemVersionOk() (*string, bool)`

GetSystemVersionOk returns a tuple with the SystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemVersion

`func (o *GetDevices200ResponseDataInner) SetSystemVersion(v string)`

SetSystemVersion sets SystemVersion field to given value.

### HasSystemVersion

`func (o *GetDevices200ResponseDataInner) HasSystemVersion() bool`

HasSystemVersion returns a boolean if a field has been set.

### SetSystemVersionNil

`func (o *GetDevices200ResponseDataInner) SetSystemVersionNil(b bool)`

 SetSystemVersionNil sets the value for SystemVersion to be an explicit nil

### UnsetSystemVersion
`func (o *GetDevices200ResponseDataInner) UnsetSystemVersion()`

UnsetSystemVersion ensures that no value is present for SystemVersion, not even an explicit nil
### GetOsVersion

`func (o *GetDevices200ResponseDataInner) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *GetDevices200ResponseDataInner) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *GetDevices200ResponseDataInner) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *GetDevices200ResponseDataInner) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersionNil

`func (o *GetDevices200ResponseDataInner) SetOsVersionNil(b bool)`

 SetOsVersionNil sets the value for OsVersion to be an explicit nil

### UnsetOsVersion
`func (o *GetDevices200ResponseDataInner) UnsetOsVersion()`

UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
### GetSsid

`func (o *GetDevices200ResponseDataInner) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *GetDevices200ResponseDataInner) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *GetDevices200ResponseDataInner) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *GetDevices200ResponseDataInner) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### SetSsidNil

`func (o *GetDevices200ResponseDataInner) SetSsidNil(b bool)`

 SetSsidNil sets the value for Ssid to be an explicit nil

### UnsetSsid
`func (o *GetDevices200ResponseDataInner) UnsetSsid()`

UnsetSsid ensures that no value is present for Ssid, not even an explicit nil
### GetMacAddress

`func (o *GetDevices200ResponseDataInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *GetDevices200ResponseDataInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *GetDevices200ResponseDataInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *GetDevices200ResponseDataInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *GetDevices200ResponseDataInner) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *GetDevices200ResponseDataInner) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetRandomizedMacAddress

`func (o *GetDevices200ResponseDataInner) GetRandomizedMacAddress() string`

GetRandomizedMacAddress returns the RandomizedMacAddress field if non-nil, zero value otherwise.

### GetRandomizedMacAddressOk

`func (o *GetDevices200ResponseDataInner) GetRandomizedMacAddressOk() (*string, bool)`

GetRandomizedMacAddressOk returns a tuple with the RandomizedMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomizedMacAddress

`func (o *GetDevices200ResponseDataInner) SetRandomizedMacAddress(v string)`

SetRandomizedMacAddress sets RandomizedMacAddress field to given value.

### HasRandomizedMacAddress

`func (o *GetDevices200ResponseDataInner) HasRandomizedMacAddress() bool`

HasRandomizedMacAddress returns a boolean if a field has been set.

### SetRandomizedMacAddressNil

`func (o *GetDevices200ResponseDataInner) SetRandomizedMacAddressNil(b bool)`

 SetRandomizedMacAddressNil sets the value for RandomizedMacAddress to be an explicit nil

### UnsetRandomizedMacAddress
`func (o *GetDevices200ResponseDataInner) UnsetRandomizedMacAddress()`

UnsetRandomizedMacAddress ensures that no value is present for RandomizedMacAddress, not even an explicit nil
### GetStorageSpaceFreeGb

`func (o *GetDevices200ResponseDataInner) GetStorageSpaceFreeGb() float32`

GetStorageSpaceFreeGb returns the StorageSpaceFreeGb field if non-nil, zero value otherwise.

### GetStorageSpaceFreeGbOk

`func (o *GetDevices200ResponseDataInner) GetStorageSpaceFreeGbOk() (*float32, bool)`

GetStorageSpaceFreeGbOk returns a tuple with the StorageSpaceFreeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpaceFreeGb

`func (o *GetDevices200ResponseDataInner) SetStorageSpaceFreeGb(v float32)`

SetStorageSpaceFreeGb sets StorageSpaceFreeGb field to given value.

### HasStorageSpaceFreeGb

`func (o *GetDevices200ResponseDataInner) HasStorageSpaceFreeGb() bool`

HasStorageSpaceFreeGb returns a boolean if a field has been set.

### SetStorageSpaceFreeGbNil

`func (o *GetDevices200ResponseDataInner) SetStorageSpaceFreeGbNil(b bool)`

 SetStorageSpaceFreeGbNil sets the value for StorageSpaceFreeGb to be an explicit nil

### UnsetStorageSpaceFreeGb
`func (o *GetDevices200ResponseDataInner) UnsetStorageSpaceFreeGb()`

UnsetStorageSpaceFreeGb ensures that no value is present for StorageSpaceFreeGb, not even an explicit nil
### GetStorageSpaceTotalGb

`func (o *GetDevices200ResponseDataInner) GetStorageSpaceTotalGb() float32`

GetStorageSpaceTotalGb returns the StorageSpaceTotalGb field if non-nil, zero value otherwise.

### GetStorageSpaceTotalGbOk

`func (o *GetDevices200ResponseDataInner) GetStorageSpaceTotalGbOk() (*float32, bool)`

GetStorageSpaceTotalGbOk returns a tuple with the StorageSpaceTotalGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpaceTotalGb

`func (o *GetDevices200ResponseDataInner) SetStorageSpaceTotalGb(v float32)`

SetStorageSpaceTotalGb sets StorageSpaceTotalGb field to given value.

### HasStorageSpaceTotalGb

`func (o *GetDevices200ResponseDataInner) HasStorageSpaceTotalGb() bool`

HasStorageSpaceTotalGb returns a boolean if a field has been set.

### SetStorageSpaceTotalGbNil

`func (o *GetDevices200ResponseDataInner) SetStorageSpaceTotalGbNil(b bool)`

 SetStorageSpaceTotalGbNil sets the value for StorageSpaceTotalGb to be an explicit nil

### UnsetStorageSpaceTotalGb
`func (o *GetDevices200ResponseDataInner) UnsetStorageSpaceTotalGb()`

UnsetStorageSpaceTotalGb ensures that no value is present for StorageSpaceTotalGb, not even an explicit nil
### GetBatteryHealth

`func (o *GetDevices200ResponseDataInner) GetBatteryHealth() string`

GetBatteryHealth returns the BatteryHealth field if non-nil, zero value otherwise.

### GetBatteryHealthOk

`func (o *GetDevices200ResponseDataInner) GetBatteryHealthOk() (*string, bool)`

GetBatteryHealthOk returns a tuple with the BatteryHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryHealth

`func (o *GetDevices200ResponseDataInner) SetBatteryHealth(v string)`

SetBatteryHealth sets BatteryHealth field to given value.

### HasBatteryHealth

`func (o *GetDevices200ResponseDataInner) HasBatteryHealth() bool`

HasBatteryHealth returns a boolean if a field has been set.

### SetBatteryHealthNil

`func (o *GetDevices200ResponseDataInner) SetBatteryHealthNil(b bool)`

 SetBatteryHealthNil sets the value for BatteryHealth to be an explicit nil

### UnsetBatteryHealth
`func (o *GetDevices200ResponseDataInner) UnsetBatteryHealth()`

UnsetBatteryHealth ensures that no value is present for BatteryHealth, not even an explicit nil
### GetBatteryCharging

`func (o *GetDevices200ResponseDataInner) GetBatteryCharging() bool`

GetBatteryCharging returns the BatteryCharging field if non-nil, zero value otherwise.

### GetBatteryChargingOk

`func (o *GetDevices200ResponseDataInner) GetBatteryChargingOk() (*bool, bool)`

GetBatteryChargingOk returns a tuple with the BatteryCharging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryCharging

`func (o *GetDevices200ResponseDataInner) SetBatteryCharging(v bool)`

SetBatteryCharging sets BatteryCharging field to given value.

### HasBatteryCharging

`func (o *GetDevices200ResponseDataInner) HasBatteryCharging() bool`

HasBatteryCharging returns a boolean if a field has been set.

### SetBatteryChargingNil

`func (o *GetDevices200ResponseDataInner) SetBatteryChargingNil(b bool)`

 SetBatteryChargingNil sets the value for BatteryCharging to be an explicit nil

### UnsetBatteryCharging
`func (o *GetDevices200ResponseDataInner) UnsetBatteryCharging()`

UnsetBatteryCharging ensures that no value is present for BatteryCharging, not even an explicit nil
### GetBatteryPercentage

`func (o *GetDevices200ResponseDataInner) GetBatteryPercentage() int32`

GetBatteryPercentage returns the BatteryPercentage field if non-nil, zero value otherwise.

### GetBatteryPercentageOk

`func (o *GetDevices200ResponseDataInner) GetBatteryPercentageOk() (*int32, bool)`

GetBatteryPercentageOk returns a tuple with the BatteryPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryPercentage

`func (o *GetDevices200ResponseDataInner) SetBatteryPercentage(v int32)`

SetBatteryPercentage sets BatteryPercentage field to given value.

### HasBatteryPercentage

`func (o *GetDevices200ResponseDataInner) HasBatteryPercentage() bool`

HasBatteryPercentage returns a boolean if a field has been set.

### SetBatteryPercentageNil

`func (o *GetDevices200ResponseDataInner) SetBatteryPercentageNil(b bool)`

 SetBatteryPercentageNil sets the value for BatteryPercentage to be an explicit nil

### UnsetBatteryPercentage
`func (o *GetDevices200ResponseDataInner) UnsetBatteryPercentage()`

UnsetBatteryPercentage ensures that no value is present for BatteryPercentage, not even an explicit nil
### GetBatteryTemperatureC

`func (o *GetDevices200ResponseDataInner) GetBatteryTemperatureC() float32`

GetBatteryTemperatureC returns the BatteryTemperatureC field if non-nil, zero value otherwise.

### GetBatteryTemperatureCOk

`func (o *GetDevices200ResponseDataInner) GetBatteryTemperatureCOk() (*float32, bool)`

GetBatteryTemperatureCOk returns a tuple with the BatteryTemperatureC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryTemperatureC

`func (o *GetDevices200ResponseDataInner) SetBatteryTemperatureC(v float32)`

SetBatteryTemperatureC sets BatteryTemperatureC field to given value.

### HasBatteryTemperatureC

`func (o *GetDevices200ResponseDataInner) HasBatteryTemperatureC() bool`

HasBatteryTemperatureC returns a boolean if a field has been set.

### SetBatteryTemperatureCNil

`func (o *GetDevices200ResponseDataInner) SetBatteryTemperatureCNil(b bool)`

 SetBatteryTemperatureCNil sets the value for BatteryTemperatureC to be an explicit nil

### UnsetBatteryTemperatureC
`func (o *GetDevices200ResponseDataInner) UnsetBatteryTemperatureC()`

UnsetBatteryTemperatureC ensures that no value is present for BatteryTemperatureC, not even an explicit nil
### GetIpAddress

`func (o *GetDevices200ResponseDataInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GetDevices200ResponseDataInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GetDevices200ResponseDataInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GetDevices200ResponseDataInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *GetDevices200ResponseDataInner) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *GetDevices200ResponseDataInner) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetSignalStrength

`func (o *GetDevices200ResponseDataInner) GetSignalStrength() float32`

GetSignalStrength returns the SignalStrength field if non-nil, zero value otherwise.

### GetSignalStrengthOk

`func (o *GetDevices200ResponseDataInner) GetSignalStrengthOk() (*float32, bool)`

GetSignalStrengthOk returns a tuple with the SignalStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalStrength

`func (o *GetDevices200ResponseDataInner) SetSignalStrength(v float32)`

SetSignalStrength sets SignalStrength field to given value.

### HasSignalStrength

`func (o *GetDevices200ResponseDataInner) HasSignalStrength() bool`

HasSignalStrength returns a boolean if a field has been set.

### SetSignalStrengthNil

`func (o *GetDevices200ResponseDataInner) SetSignalStrengthNil(b bool)`

 SetSignalStrengthNil sets the value for SignalStrength to be an explicit nil

### UnsetSignalStrength
`func (o *GetDevices200ResponseDataInner) UnsetSignalStrength()`

UnsetSignalStrength ensures that no value is present for SignalStrength, not even an explicit nil
### GetFrequencyMhz

`func (o *GetDevices200ResponseDataInner) GetFrequencyMhz() int32`

GetFrequencyMhz returns the FrequencyMhz field if non-nil, zero value otherwise.

### GetFrequencyMhzOk

`func (o *GetDevices200ResponseDataInner) GetFrequencyMhzOk() (*int32, bool)`

GetFrequencyMhzOk returns a tuple with the FrequencyMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyMhz

`func (o *GetDevices200ResponseDataInner) SetFrequencyMhz(v int32)`

SetFrequencyMhz sets FrequencyMhz field to given value.

### HasFrequencyMhz

`func (o *GetDevices200ResponseDataInner) HasFrequencyMhz() bool`

HasFrequencyMhz returns a boolean if a field has been set.

### SetFrequencyMhzNil

`func (o *GetDevices200ResponseDataInner) SetFrequencyMhzNil(b bool)`

 SetFrequencyMhzNil sets the value for FrequencyMhz to be an explicit nil

### UnsetFrequencyMhz
`func (o *GetDevices200ResponseDataInner) UnsetFrequencyMhz()`

UnsetFrequencyMhz ensures that no value is present for FrequencyMhz, not even an explicit nil
### GetLinkSpeedMbps

`func (o *GetDevices200ResponseDataInner) GetLinkSpeedMbps() int32`

GetLinkSpeedMbps returns the LinkSpeedMbps field if non-nil, zero value otherwise.

### GetLinkSpeedMbpsOk

`func (o *GetDevices200ResponseDataInner) GetLinkSpeedMbpsOk() (*int32, bool)`

GetLinkSpeedMbpsOk returns a tuple with the LinkSpeedMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeedMbps

`func (o *GetDevices200ResponseDataInner) SetLinkSpeedMbps(v int32)`

SetLinkSpeedMbps sets LinkSpeedMbps field to given value.

### HasLinkSpeedMbps

`func (o *GetDevices200ResponseDataInner) HasLinkSpeedMbps() bool`

HasLinkSpeedMbps returns a boolean if a field has been set.

### SetLinkSpeedMbpsNil

`func (o *GetDevices200ResponseDataInner) SetLinkSpeedMbpsNil(b bool)`

 SetLinkSpeedMbpsNil sets the value for LinkSpeedMbps to be an explicit nil

### UnsetLinkSpeedMbps
`func (o *GetDevices200ResponseDataInner) UnsetLinkSpeedMbps()`

UnsetLinkSpeedMbps ensures that no value is present for LinkSpeedMbps, not even an explicit nil
### GetLastLocationLatitude

`func (o *GetDevices200ResponseDataInner) GetLastLocationLatitude() float32`

GetLastLocationLatitude returns the LastLocationLatitude field if non-nil, zero value otherwise.

### GetLastLocationLatitudeOk

`func (o *GetDevices200ResponseDataInner) GetLastLocationLatitudeOk() (*float32, bool)`

GetLastLocationLatitudeOk returns a tuple with the LastLocationLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocationLatitude

`func (o *GetDevices200ResponseDataInner) SetLastLocationLatitude(v float32)`

SetLastLocationLatitude sets LastLocationLatitude field to given value.

### HasLastLocationLatitude

`func (o *GetDevices200ResponseDataInner) HasLastLocationLatitude() bool`

HasLastLocationLatitude returns a boolean if a field has been set.

### SetLastLocationLatitudeNil

`func (o *GetDevices200ResponseDataInner) SetLastLocationLatitudeNil(b bool)`

 SetLastLocationLatitudeNil sets the value for LastLocationLatitude to be an explicit nil

### UnsetLastLocationLatitude
`func (o *GetDevices200ResponseDataInner) UnsetLastLocationLatitude()`

UnsetLastLocationLatitude ensures that no value is present for LastLocationLatitude, not even an explicit nil
### GetLastLocationLongitude

`func (o *GetDevices200ResponseDataInner) GetLastLocationLongitude() float32`

GetLastLocationLongitude returns the LastLocationLongitude field if non-nil, zero value otherwise.

### GetLastLocationLongitudeOk

`func (o *GetDevices200ResponseDataInner) GetLastLocationLongitudeOk() (*float32, bool)`

GetLastLocationLongitudeOk returns a tuple with the LastLocationLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocationLongitude

`func (o *GetDevices200ResponseDataInner) SetLastLocationLongitude(v float32)`

SetLastLocationLongitude sets LastLocationLongitude field to given value.

### HasLastLocationLongitude

`func (o *GetDevices200ResponseDataInner) HasLastLocationLongitude() bool`

HasLastLocationLongitude returns a boolean if a field has been set.

### SetLastLocationLongitudeNil

`func (o *GetDevices200ResponseDataInner) SetLastLocationLongitudeNil(b bool)`

 SetLastLocationLongitudeNil sets the value for LastLocationLongitude to be an explicit nil

### UnsetLastLocationLongitude
`func (o *GetDevices200ResponseDataInner) UnsetLastLocationLongitude()`

UnsetLastLocationLongitude ensures that no value is present for LastLocationLongitude, not even an explicit nil
### GetLastLocationAt

`func (o *GetDevices200ResponseDataInner) GetLastLocationAt() string`

GetLastLocationAt returns the LastLocationAt field if non-nil, zero value otherwise.

### GetLastLocationAtOk

`func (o *GetDevices200ResponseDataInner) GetLastLocationAtOk() (*string, bool)`

GetLastLocationAtOk returns a tuple with the LastLocationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocationAt

`func (o *GetDevices200ResponseDataInner) SetLastLocationAt(v string)`

SetLastLocationAt sets LastLocationAt field to given value.

### HasLastLocationAt

`func (o *GetDevices200ResponseDataInner) HasLastLocationAt() bool`

HasLastLocationAt returns a boolean if a field has been set.

### SetLastLocationAtNil

`func (o *GetDevices200ResponseDataInner) SetLastLocationAtNil(b bool)`

 SetLastLocationAtNil sets the value for LastLocationAt to be an explicit nil

### UnsetLastLocationAt
`func (o *GetDevices200ResponseDataInner) UnsetLastLocationAt()`

UnsetLastLocationAt ensures that no value is present for LastLocationAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


