# GetApps200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**DeviceModels** | Pointer to [**[]GetApps200ResponseDataInnerDeviceModelsInner**](GetApps200ResponseDataInnerDeviceModelsInner.md) |  | [optional] 
**PackageName** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to [**NullableGetApps200ResponseDataInnerIcon**](GetApps200ResponseDataInnerIcon.md) |  | [optional] 
**LatestAvailableVersion** | Pointer to **NullableString** |  | [optional] 
**OwnerOrganization** | Pointer to [**GetApps200ResponseDataInnerOwnerOrganization**](GetApps200ResponseDataInnerOwnerOrganization.md) |  | [optional] 

## Methods

### NewGetApps200ResponseDataInner

`func NewGetApps200ResponseDataInner() *GetApps200ResponseDataInner`

NewGetApps200ResponseDataInner instantiates a new GetApps200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApps200ResponseDataInnerWithDefaults

`func NewGetApps200ResponseDataInnerWithDefaults() *GetApps200ResponseDataInner`

NewGetApps200ResponseDataInnerWithDefaults instantiates a new GetApps200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApps200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApps200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApps200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetApps200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *GetApps200ResponseDataInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetApps200ResponseDataInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetApps200ResponseDataInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetApps200ResponseDataInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *GetApps200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetApps200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetApps200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetApps200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetApps200ResponseDataInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetApps200ResponseDataInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTags

`func (o *GetApps200ResponseDataInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetApps200ResponseDataInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetApps200ResponseDataInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetApps200ResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDeviceModels

`func (o *GetApps200ResponseDataInner) GetDeviceModels() []GetApps200ResponseDataInnerDeviceModelsInner`

GetDeviceModels returns the DeviceModels field if non-nil, zero value otherwise.

### GetDeviceModelsOk

`func (o *GetApps200ResponseDataInner) GetDeviceModelsOk() (*[]GetApps200ResponseDataInnerDeviceModelsInner, bool)`

GetDeviceModelsOk returns a tuple with the DeviceModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModels

`func (o *GetApps200ResponseDataInner) SetDeviceModels(v []GetApps200ResponseDataInnerDeviceModelsInner)`

SetDeviceModels sets DeviceModels field to given value.

### HasDeviceModels

`func (o *GetApps200ResponseDataInner) HasDeviceModels() bool`

HasDeviceModels returns a boolean if a field has been set.

### GetPackageName

`func (o *GetApps200ResponseDataInner) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *GetApps200ResponseDataInner) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *GetApps200ResponseDataInner) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *GetApps200ResponseDataInner) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### SetPackageNameNil

`func (o *GetApps200ResponseDataInner) SetPackageNameNil(b bool)`

 SetPackageNameNil sets the value for PackageName to be an explicit nil

### UnsetPackageName
`func (o *GetApps200ResponseDataInner) UnsetPackageName()`

UnsetPackageName ensures that no value is present for PackageName, not even an explicit nil
### GetIcon

`func (o *GetApps200ResponseDataInner) GetIcon() GetApps200ResponseDataInnerIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GetApps200ResponseDataInner) GetIconOk() (*GetApps200ResponseDataInnerIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GetApps200ResponseDataInner) SetIcon(v GetApps200ResponseDataInnerIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GetApps200ResponseDataInner) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *GetApps200ResponseDataInner) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *GetApps200ResponseDataInner) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetLatestAvailableVersion

`func (o *GetApps200ResponseDataInner) GetLatestAvailableVersion() string`

GetLatestAvailableVersion returns the LatestAvailableVersion field if non-nil, zero value otherwise.

### GetLatestAvailableVersionOk

`func (o *GetApps200ResponseDataInner) GetLatestAvailableVersionOk() (*string, bool)`

GetLatestAvailableVersionOk returns a tuple with the LatestAvailableVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAvailableVersion

`func (o *GetApps200ResponseDataInner) SetLatestAvailableVersion(v string)`

SetLatestAvailableVersion sets LatestAvailableVersion field to given value.

### HasLatestAvailableVersion

`func (o *GetApps200ResponseDataInner) HasLatestAvailableVersion() bool`

HasLatestAvailableVersion returns a boolean if a field has been set.

### SetLatestAvailableVersionNil

`func (o *GetApps200ResponseDataInner) SetLatestAvailableVersionNil(b bool)`

 SetLatestAvailableVersionNil sets the value for LatestAvailableVersion to be an explicit nil

### UnsetLatestAvailableVersion
`func (o *GetApps200ResponseDataInner) UnsetLatestAvailableVersion()`

UnsetLatestAvailableVersion ensures that no value is present for LatestAvailableVersion, not even an explicit nil
### GetOwnerOrganization

`func (o *GetApps200ResponseDataInner) GetOwnerOrganization() GetApps200ResponseDataInnerOwnerOrganization`

GetOwnerOrganization returns the OwnerOrganization field if non-nil, zero value otherwise.

### GetOwnerOrganizationOk

`func (o *GetApps200ResponseDataInner) GetOwnerOrganizationOk() (*GetApps200ResponseDataInnerOwnerOrganization, bool)`

GetOwnerOrganizationOk returns a tuple with the OwnerOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOrganization

`func (o *GetApps200ResponseDataInner) SetOwnerOrganization(v GetApps200ResponseDataInnerOwnerOrganization)`

SetOwnerOrganization sets OwnerOrganization field to given value.

### HasOwnerOrganization

`func (o *GetApps200ResponseDataInner) HasOwnerOrganization() bool`

HasOwnerOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


