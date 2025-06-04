# CompleteAppVersionUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**UploadId** | **string** |  | 
**Parts** | [**[]CompleteAppVersionUploadRequestPartsInner**](CompleteAppVersionUploadRequestPartsInner.md) |  | 
**VersionName** | Pointer to **NullableString** |  | [optional] 
**ReleaseNotes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCompleteAppVersionUploadRequest

`func NewCompleteAppVersionUploadRequest(key string, uploadId string, parts []CompleteAppVersionUploadRequestPartsInner, ) *CompleteAppVersionUploadRequest`

NewCompleteAppVersionUploadRequest instantiates a new CompleteAppVersionUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteAppVersionUploadRequestWithDefaults

`func NewCompleteAppVersionUploadRequestWithDefaults() *CompleteAppVersionUploadRequest`

NewCompleteAppVersionUploadRequestWithDefaults instantiates a new CompleteAppVersionUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CompleteAppVersionUploadRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CompleteAppVersionUploadRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CompleteAppVersionUploadRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetUploadId

`func (o *CompleteAppVersionUploadRequest) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *CompleteAppVersionUploadRequest) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *CompleteAppVersionUploadRequest) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.


### GetParts

`func (o *CompleteAppVersionUploadRequest) GetParts() []CompleteAppVersionUploadRequestPartsInner`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *CompleteAppVersionUploadRequest) GetPartsOk() (*[]CompleteAppVersionUploadRequestPartsInner, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *CompleteAppVersionUploadRequest) SetParts(v []CompleteAppVersionUploadRequestPartsInner)`

SetParts sets Parts field to given value.


### GetVersionName

`func (o *CompleteAppVersionUploadRequest) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *CompleteAppVersionUploadRequest) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *CompleteAppVersionUploadRequest) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *CompleteAppVersionUploadRequest) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### SetVersionNameNil

`func (o *CompleteAppVersionUploadRequest) SetVersionNameNil(b bool)`

 SetVersionNameNil sets the value for VersionName to be an explicit nil

### UnsetVersionName
`func (o *CompleteAppVersionUploadRequest) UnsetVersionName()`

UnsetVersionName ensures that no value is present for VersionName, not even an explicit nil
### GetReleaseNotes

`func (o *CompleteAppVersionUploadRequest) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *CompleteAppVersionUploadRequest) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *CompleteAppVersionUploadRequest) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.

### HasReleaseNotes

`func (o *CompleteAppVersionUploadRequest) HasReleaseNotes() bool`

HasReleaseNotes returns a boolean if a field has been set.

### SetReleaseNotesNil

`func (o *CompleteAppVersionUploadRequest) SetReleaseNotesNil(b bool)`

 SetReleaseNotesNil sets the value for ReleaseNotes to be an explicit nil

### UnsetReleaseNotes
`func (o *CompleteAppVersionUploadRequest) UnsetReleaseNotes()`

UnsetReleaseNotes ensures that no value is present for ReleaseNotes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


