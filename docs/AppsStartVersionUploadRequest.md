# AppsStartVersionUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** |  | 
**ReleaseChannelId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAppsStartVersionUploadRequest

`func NewAppsStartVersionUploadRequest(filename string, ) *AppsStartVersionUploadRequest`

NewAppsStartVersionUploadRequest instantiates a new AppsStartVersionUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsStartVersionUploadRequestWithDefaults

`func NewAppsStartVersionUploadRequestWithDefaults() *AppsStartVersionUploadRequest`

NewAppsStartVersionUploadRequestWithDefaults instantiates a new AppsStartVersionUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *AppsStartVersionUploadRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AppsStartVersionUploadRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AppsStartVersionUploadRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetReleaseChannelId

`func (o *AppsStartVersionUploadRequest) GetReleaseChannelId() string`

GetReleaseChannelId returns the ReleaseChannelId field if non-nil, zero value otherwise.

### GetReleaseChannelIdOk

`func (o *AppsStartVersionUploadRequest) GetReleaseChannelIdOk() (*string, bool)`

GetReleaseChannelIdOk returns a tuple with the ReleaseChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseChannelId

`func (o *AppsStartVersionUploadRequest) SetReleaseChannelId(v string)`

SetReleaseChannelId sets ReleaseChannelId field to given value.

### HasReleaseChannelId

`func (o *AppsStartVersionUploadRequest) HasReleaseChannelId() bool`

HasReleaseChannelId returns a boolean if a field has been set.

### SetReleaseChannelIdNil

`func (o *AppsStartVersionUploadRequest) SetReleaseChannelIdNil(b bool)`

 SetReleaseChannelIdNil sets the value for ReleaseChannelId to be an explicit nil

### UnsetReleaseChannelId
`func (o *AppsStartVersionUploadRequest) UnsetReleaseChannelId()`

UnsetReleaseChannelId ensures that no value is present for ReleaseChannelId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


