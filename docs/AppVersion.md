# AppVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Version** | **string** |  | 
**Code** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableString** |  | [optional] 
**Added** | **time.Time** |  | 
**Status** | [**AppVersionStatus**](AppVersionStatus.md) |  | 
**DownloadUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAppVersion

`func NewAppVersion(id string, version string, added time.Time, status AppVersionStatus, ) *AppVersion`

NewAppVersion instantiates a new AppVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppVersionWithDefaults

`func NewAppVersionWithDefaults() *AppVersion`

NewAppVersionWithDefaults instantiates a new AppVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppVersion) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *AppVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppVersion) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetCode

`func (o *AppVersion) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AppVersion) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AppVersion) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AppVersion) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AppVersion) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AppVersion) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetSize

`func (o *AppVersion) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AppVersion) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AppVersion) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *AppVersion) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *AppVersion) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *AppVersion) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetAdded

`func (o *AppVersion) GetAdded() time.Time`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *AppVersion) GetAddedOk() (*time.Time, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *AppVersion) SetAdded(v time.Time)`

SetAdded sets Added field to given value.


### GetStatus

`func (o *AppVersion) GetStatus() AppVersionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppVersion) GetStatusOk() (*AppVersionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppVersion) SetStatus(v AppVersionStatus)`

SetStatus sets Status field to given value.


### GetDownloadUrl

`func (o *AppVersion) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *AppVersion) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *AppVersion) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *AppVersion) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### SetDownloadUrlNil

`func (o *AppVersion) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *AppVersion) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


