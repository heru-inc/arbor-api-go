# \AppsAPI

All URIs are relative to *https://api.xrdm.app/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteAppVersionUpload**](AppsAPI.md#CompleteAppVersionUpload) | **Post** /apps/{appId}/versions/{versionId}/complete | 
[**GetApp**](AppsAPI.md#GetApp) | **Get** /apps/{appId} | 
[**GetAppReleaseChannel**](AppsAPI.md#GetAppReleaseChannel) | **Get** /apps/{appId}/release-channels/{releaseChannelId} | 
[**GetAppReleaseChannels**](AppsAPI.md#GetAppReleaseChannels) | **Get** /apps/{appId}/release-channels | 
[**GetAppVersions**](AppsAPI.md#GetAppVersions) | **Get** /apps/{appId}/versions | 
[**GetApps**](AppsAPI.md#GetApps) | **Get** /apps | 
[**InitiateAppVersionUpload**](AppsAPI.md#InitiateAppVersionUpload) | **Post** /apps/{appId}/versions | 
[**PreSignAppVersionUpload**](AppsAPI.md#PreSignAppVersionUpload) | **Post** /apps/{appId}/versions/{versionId}/pre-sign | 
[**ShareReleaseChannel**](AppsAPI.md#ShareReleaseChannel) | **Post** /apps/{appId}/release-channels/{releaseChannelId}/share | 
[**UnshareReleaseChannel**](AppsAPI.md#UnshareReleaseChannel) | **Delete** /apps/{appId}/release-channels/{releaseChannelId}/share | 
[**UpdateApp**](AppsAPI.md#UpdateApp) | **Put** /apps/{appId} | 
[**UpdateReleaseChannel**](AppsAPI.md#UpdateReleaseChannel) | **Put** /apps/{appId}/release-channels/{releaseChannelId} | 



## CompleteAppVersionUpload

> GetAppVersions200ResponseDataInner CompleteAppVersionUpload(ctx, appId, versionId).CompleteAppVersionUploadRequest(completeAppVersionUploadRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/heru-inc/arbor-api-go"
)

func main() {
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app.
	versionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app version.
	completeAppVersionUploadRequest := *openapiclient.NewCompleteAppVersionUploadRequest("xrdm-016da47e-4c49-48ad-9c61-94904e06a226/apps/2f5d24ec-be98-49d3-ac42-cdecbeea40f0/f40d390d-94fc-47a9-952c-8b7321f3eefc/my-app.apk", "OGRhZDQ5NTMtOTQxYS00ZDI3LWE0MzUtOGUyNzEzZWZiYTlmLmY1YmE5NTkyLWIzZDMtNDE1ZS1iOWIxLWM1YWJjYjAxMmY5OQ", []openapiclient.CompleteAppVersionUploadRequestPartsInner{*openapiclient.NewCompleteAppVersionUploadRequestPartsInner(int32(1), "d41d8cd98f00b204e9800998ecf8427e")}) // CompleteAppVersionUploadRequest | The fields to complete the upload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.CompleteAppVersionUpload(context.Background(), appId, versionId).CompleteAppVersionUploadRequest(completeAppVersionUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.CompleteAppVersionUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteAppVersionUpload`: GetAppVersions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.CompleteAppVersionUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of an app. | 
**versionId** | **string** | The ID of an app version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteAppVersionUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **completeAppVersionUploadRequest** | [**CompleteAppVersionUploadRequest**](CompleteAppVersionUploadRequest.md) | The fields to complete the upload. | 

### Return type

[**GetAppVersions200ResponseDataInner**](GetAppVersions200ResponseDataInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApp

> GetApps200ResponseDataInner GetApp(ctx, appId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/heru-inc/arbor-api-go"
)

func main() {
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetApp(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApp`: GetApps200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of an app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetApps200ResponseDataInner**](GetApps200ResponseDataInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppReleaseChannel

> GetAppReleaseChannels200ResponseDataInner GetAppReleaseChannel(ctx, appId, releaseChannelId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/heru-inc/arbor-api-go"
)

func main() {
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app.
	releaseChannelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a release channel.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetAppReleaseChannel(context.Background(), appId, releaseChannelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppReleaseChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppReleaseChannel`: GetAppReleaseChannels200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppReleaseChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of an app. | 
**releaseChannelId** | **string** | The ID of a release channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppReleaseChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAppReleaseChannels200ResponseDataInner**](GetAppReleaseChannels200ResponseDataInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppReleaseChannels

> GetAppReleaseChannels200Response GetAppReleaseChannels(ctx, appId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/heru-inc/arbor-api-go"
)

func main() {
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetAppReleaseChannels(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppReleaseChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppReleaseChannels`: GetAppReleaseChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppReleaseChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of an app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppReleaseChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAppReleaseChannels200Response**](GetAppReleaseChannels200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppVersions

> GetAppVersions200Response GetAppVersions(ctx, appId).PerPage(perPage).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/heru-inc/arbor-api-go"
)

func main() {
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app.
	perPage := int32(56) // int32 | The number of items to return per page. (optional) (default to 10)
	page := int32(56) // int32 | The page number to return. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetAppVersions(context.Background(), appId).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppVersions`: GetAppVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of an app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of items to return per page. | [default to 10]
 **page** | **int32** | The page number to return. | [default to 1]

### Return type

[**GetAppVersions200Response**](GetAppVersions200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApps

> GetApps200Response GetApps(ctx).PerPage(perPage).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/heru-inc/arbor-api-go"
)

func main() {
	perPage := int32(56) // int32 | The number of items to return per page. (optional) (default to 10)
	page := int32(56) // int32 | The page number to return. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetApps(context.Background()).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApps`: GetApps200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of items to return per page. | [default to 10]
 **page** | **int32** | The page number to return. | [default to 1]

### Return type

[**GetApps200Response**](GetApps200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateAppVersionUpload

> InitiateAppVersionUpload200Response InitiateAppVersionUpload(ctx, appId).InitiateAppVersionUploadRequest(initiateAppVersionUploadRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/heru-inc/arbor-api-go"
)

func main() {
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app.
	initiateAppVersionUploadRequest := *openapiclient.NewInitiateAppVersionUploadRequest("my-app.apk") // InitiateAppVersionUploadRequest | Initiate a new app version upload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.InitiateAppVersionUpload(context.Background(), appId).InitiateAppVersionUploadRequest(initiateAppVersionUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.InitiateAppVersionUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateAppVersionUpload`: InitiateAppVersionUpload200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.InitiateAppVersionUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of an app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateAppVersionUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **initiateAppVersionUploadRequest** | [**InitiateAppVersionUploadRequest**](InitiateAppVersionUploadRequest.md) | Initiate a new app version upload. | 

### Return type

[**InitiateAppVersionUpload200Response**](InitiateAppVersionUpload200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreSignAppVersionUpload

> []PreSignAppVersionUpload200ResponseInner PreSignAppVersionUpload(ctx, appId, versionId).PreSignAppVersionUploadRequest(preSignAppVersionUploadRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/heru-inc/arbor-api-go"
)

func main() {
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app.
	versionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app version.
	preSignAppVersionUploadRequest := *openapiclient.NewPreSignAppVersionUploadRequest("xrdm-016da47e-4c49-48ad-9c61-94904e06a226/apps/2f5d24ec-be98-49d3-ac42-cdecbeea40f0/f40d390d-94fc-47a9-952c-8b7321f3eefc/my-app.apk", "OGRhZDQ5NTMtOTQxYS00ZDI3LWE0MzUtOGUyNzEzZWZiYTlmLmY1YmE5NTkyLWIzZDMtNDE1ZS1iOWIxLWM1YWJjYjAxMmY5OQ", []int32{int32(1)}) // PreSignAppVersionUploadRequest | The fields to retrieve presigned URLs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.PreSignAppVersionUpload(context.Background(), appId, versionId).PreSignAppVersionUploadRequest(preSignAppVersionUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.PreSignAppVersionUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreSignAppVersionUpload`: []PreSignAppVersionUpload200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.PreSignAppVersionUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of an app. | 
**versionId** | **string** | The ID of an app version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreSignAppVersionUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **preSignAppVersionUploadRequest** | [**PreSignAppVersionUploadRequest**](PreSignAppVersionUploadRequest.md) | The fields to retrieve presigned URLs. | 

### Return type

[**[]PreSignAppVersionUpload200ResponseInner**](PreSignAppVersionUpload200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareReleaseChannel

> ShareReleaseChannel200Response ShareReleaseChannel(ctx, appId, releaseChannelId).ShareReleaseChannelRequest(shareReleaseChannelRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/heru-inc/arbor-api-go"
)

func main() {
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app.
	releaseChannelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a release channel.
	shareReleaseChannelRequest := *openapiclient.NewShareReleaseChannelRequest("another-organization-inc") // ShareReleaseChannelRequest | The details of the organization with which you want to share/unshare the release channel. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.ShareReleaseChannel(context.Background(), appId, releaseChannelId).ShareReleaseChannelRequest(shareReleaseChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ShareReleaseChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShareReleaseChannel`: ShareReleaseChannel200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.ShareReleaseChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of an app. | 
**releaseChannelId** | **string** | The ID of a release channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareReleaseChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shareReleaseChannelRequest** | [**ShareReleaseChannelRequest**](ShareReleaseChannelRequest.md) | The details of the organization with which you want to share/unshare the release channel. | 

### Return type

[**ShareReleaseChannel200Response**](ShareReleaseChannel200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnshareReleaseChannel

> UnshareReleaseChannel200Response UnshareReleaseChannel(ctx, appId, releaseChannelId).ShareReleaseChannelRequest(shareReleaseChannelRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/heru-inc/arbor-api-go"
)

func main() {
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app.
	releaseChannelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a release channel.
	shareReleaseChannelRequest := *openapiclient.NewShareReleaseChannelRequest("another-organization-inc") // ShareReleaseChannelRequest | The details of the organization with which you want to share/unshare the release channel. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.UnshareReleaseChannel(context.Background(), appId, releaseChannelId).ShareReleaseChannelRequest(shareReleaseChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UnshareReleaseChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnshareReleaseChannel`: UnshareReleaseChannel200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.UnshareReleaseChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of an app. | 
**releaseChannelId** | **string** | The ID of a release channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnshareReleaseChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shareReleaseChannelRequest** | [**ShareReleaseChannelRequest**](ShareReleaseChannelRequest.md) | The details of the organization with which you want to share/unshare the release channel. | 

### Return type

[**UnshareReleaseChannel200Response**](UnshareReleaseChannel200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApp

> GetApps200ResponseDataInner UpdateApp(ctx, appId).UpdateAppRequest(updateAppRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/heru-inc/arbor-api-go"
)

func main() {
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app.
	updateAppRequest := *openapiclient.NewUpdateAppRequest() // UpdateAppRequest | The app fields to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.UpdateApp(context.Background(), appId).UpdateAppRequest(updateAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UpdateApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApp`: GetApps200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.UpdateApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of an app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppRequest** | [**UpdateAppRequest**](UpdateAppRequest.md) | The app fields to update. | 

### Return type

[**GetApps200ResponseDataInner**](GetApps200ResponseDataInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReleaseChannel

> GetAppReleaseChannels200ResponseDataInner UpdateReleaseChannel(ctx, appId, releaseChannelId).UpdateReleaseChannelRequest(updateReleaseChannelRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/heru-inc/arbor-api-go"
)

func main() {
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app.
	releaseChannelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a release channel.
	updateReleaseChannelRequest := *openapiclient.NewUpdateReleaseChannelRequest("123e4567-e89b-12d3-a456-426614174000") // UpdateReleaseChannelRequest | The release channel fields to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.UpdateReleaseChannel(context.Background(), appId, releaseChannelId).UpdateReleaseChannelRequest(updateReleaseChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UpdateReleaseChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReleaseChannel`: GetAppReleaseChannels200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.UpdateReleaseChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The ID of an app. | 
**releaseChannelId** | **string** | The ID of a release channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleaseChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateReleaseChannelRequest** | [**UpdateReleaseChannelRequest**](UpdateReleaseChannelRequest.md) | The release channel fields to update. | 

### Return type

[**GetAppReleaseChannels200ResponseDataInner**](GetAppReleaseChannels200ResponseDataInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

