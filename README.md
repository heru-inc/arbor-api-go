# Go API client for openapi

This API provides a RESTful interface to interact with your organization's data.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/heru-inc/arbor-api-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.xrdm.app/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*APIInfoAPI* | [**APIInfoInfo**](docs/APIInfoAPI.md#apiinfoinfo) | **Get** / | 
*AppsAPI* | [**AppsApp**](docs/AppsAPI.md#appsapp) | **Get** /apps/{appId} | 
*AppsAPI* | [**AppsApps**](docs/AppsAPI.md#appsapps) | **Get** /apps | 
*AppsAPI* | [**AppsCompleteVersionUpload**](docs/AppsAPI.md#appscompleteversionupload) | **Post** /apps/{appId}/versions/{versionId}/complete | 
*AppsAPI* | [**AppsPreSignVersionUploadUrl**](docs/AppsAPI.md#appspresignversionuploadurl) | **Post** /apps/{appId}/versions/{versionId}/pre-sign | 
*AppsAPI* | [**AppsReleaseChannel**](docs/AppsAPI.md#appsreleasechannel) | **Get** /apps/{appId}/release-channels/{releaseChannelId} | 
*AppsAPI* | [**AppsReleaseChannels**](docs/AppsAPI.md#appsreleasechannels) | **Get** /apps/{appId}/release-channels | 
*AppsAPI* | [**AppsStartVersionUpload**](docs/AppsAPI.md#appsstartversionupload) | **Post** /apps/{appId}/versions | 
*AppsAPI* | [**AppsUpdateApp**](docs/AppsAPI.md#appsupdateapp) | **Put** /apps/{appId} | 
*AppsAPI* | [**AppsUpdateReleaseChannel**](docs/AppsAPI.md#appsupdatereleasechannel) | **Put** /apps/{appId}/release-channels/{releaseChannelId} | 
*AppsAPI* | [**AppsVersions**](docs/AppsAPI.md#appsversions) | **Get** /apps/{appId}/versions | 
*DeviceGroupsAPI* | [**DeviceGroupsGroups**](docs/DeviceGroupsAPI.md#devicegroupsgroups) | **Get** /device-groups | 
*DeviceModelsAPI* | [**DeviceModelsDeviceModel**](docs/DeviceModelsAPI.md#devicemodelsdevicemodel) | **Get** /device-models/{deviceModelId} | 
*DeviceModelsAPI* | [**DeviceModelsDeviceModels**](docs/DeviceModelsAPI.md#devicemodelsdevicemodels) | **Get** /device-models | 
*DevicesAPI* | [**DevicesDevice**](docs/DevicesAPI.md#devicesdevice) | **Get** /devices/{deviceId} | 
*DevicesAPI* | [**DevicesDevices**](docs/DevicesAPI.md#devicesdevices) | **Get** /devices | 
*DevicesAPI* | [**DevicesLaunchApp**](docs/DevicesAPI.md#deviceslaunchapp) | **Post** /devices/{deviceId}/launch/{appId} | 
*DevicesAPI* | [**DevicesReboot**](docs/DevicesAPI.md#devicesreboot) | **Post** /devices/{deviceId}/reboot | 
*DevicesAPI* | [**DevicesUpdateDevice**](docs/DevicesAPI.md#devicesupdatedevice) | **Put** /devices/{deviceId} | 
*FilesAPI* | [**FilesFile**](docs/FilesAPI.md#filesfile) | **Get** /files/{fileId} | 
*FilesAPI* | [**FilesFiles**](docs/FilesAPI.md#filesfiles) | **Get** /files | 


## Documentation For Models

 - [App](docs/App.md)
 - [AppVersion](docs/AppVersion.md)
 - [AppVersionStatus](docs/AppVersionStatus.md)
 - [AppWithDeviceModels](docs/AppWithDeviceModels.md)
 - [AppsCompleteVersionUploadRequest](docs/AppsCompleteVersionUploadRequest.md)
 - [AppsPreSignVersionUploadUrlRequest](docs/AppsPreSignVersionUploadUrlRequest.md)
 - [AppsResponse](docs/AppsResponse.md)
 - [AppsStartVersionUploadRequest](docs/AppsStartVersionUploadRequest.md)
 - [AppsUpdateAppRequest](docs/AppsUpdateAppRequest.md)
 - [AppsUpdateReleaseChannelRequest](docs/AppsUpdateReleaseChannelRequest.md)
 - [BadRequestErrorBody](docs/BadRequestErrorBody.md)
 - [CompleteUploadAppVersionValidationErrorBody](docs/CompleteUploadAppVersionValidationErrorBody.md)
 - [CompleteUploadAppVersionValidationErrorBodyErrors](docs/CompleteUploadAppVersionValidationErrorBodyErrors.md)
 - [CompleteVersionUploadRequestPart](docs/CompleteVersionUploadRequestPart.md)
 - [Device](docs/Device.md)
 - [DeviceGroup](docs/DeviceGroup.md)
 - [DeviceGroupsResponse](docs/DeviceGroupsResponse.md)
 - [DeviceModel](docs/DeviceModel.md)
 - [DeviceModelsResponse](docs/DeviceModelsResponse.md)
 - [DeviceWithSimplifiedDeviceGroup](docs/DeviceWithSimplifiedDeviceGroup.md)
 - [DevicesResponse](docs/DevicesResponse.md)
 - [DevicesUpdateDeviceRequest](docs/DevicesUpdateDeviceRequest.md)
 - [File](docs/File.md)
 - [FileWithDeviceIds](docs/FileWithDeviceIds.md)
 - [FilesResponse](docs/FilesResponse.md)
 - [ForbiddenErrorBody](docs/ForbiddenErrorBody.md)
 - [Links](docs/Links.md)
 - [Meta](docs/Meta.md)
 - [NotFoundErrorBody](docs/NotFoundErrorBody.md)
 - [PaginatorLink](docs/PaginatorLink.md)
 - [PreSignVersionUploadUrlResponsePart](docs/PreSignVersionUploadUrlResponsePart.md)
 - [PresignUploadAppVersionValidationErrorBody](docs/PresignUploadAppVersionValidationErrorBody.md)
 - [PresignUploadAppVersionValidationErrorBodyErrors](docs/PresignUploadAppVersionValidationErrorBodyErrors.md)
 - [ReleaseChannel](docs/ReleaseChannel.md)
 - [ReleaseChannelWithVersion](docs/ReleaseChannelWithVersion.md)
 - [ReleaseChannelsResponse](docs/ReleaseChannelsResponse.md)
 - [RootResponse](docs/RootResponse.md)
 - [SimplifiedDeviceGroup](docs/SimplifiedDeviceGroup.md)
 - [StartUploadAppVersionValidationErrorBody](docs/StartUploadAppVersionValidationErrorBody.md)
 - [StartUploadAppVersionValidationErrorBodyErrors](docs/StartUploadAppVersionValidationErrorBodyErrors.md)
 - [StartVersionUploadResponse](docs/StartVersionUploadResponse.md)
 - [UpdateAppValidationErrorBody](docs/UpdateAppValidationErrorBody.md)
 - [UpdateAppValidationErrorBodyErrors](docs/UpdateAppValidationErrorBodyErrors.md)
 - [UpdateDeviceValidationErrorBody](docs/UpdateDeviceValidationErrorBody.md)
 - [UpdateDeviceValidationErrorBodyErrors](docs/UpdateDeviceValidationErrorBodyErrors.md)
 - [UpdateReleaseChannelValidationErrorBody](docs/UpdateReleaseChannelValidationErrorBody.md)
 - [UpdateReleaseChannelValidationErrorBodyErrors](docs/UpdateReleaseChannelValidationErrorBodyErrors.md)
 - [VersionsResponse](docs/VersionsResponse.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



