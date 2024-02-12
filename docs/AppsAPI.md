# \AppsAPI

All URIs are relative to *https://api.xrdm.app/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsApp**](AppsAPI.md#AppsApp) | **Get** /apps/{appId} | 
[**AppsApps**](AppsAPI.md#AppsApps) | **Get** /apps | 
[**AppsCompleteVersionUpload**](AppsAPI.md#AppsCompleteVersionUpload) | **Post** /apps/{appId}/versions/{versionId}/complete | 
[**AppsPreSignVersionUploadUrl**](AppsAPI.md#AppsPreSignVersionUploadUrl) | **Post** /apps/{appId}/versions/{versionId}/pre-sign | 
[**AppsReleaseChannel**](AppsAPI.md#AppsReleaseChannel) | **Get** /apps/{appId}/release-channels/{releaseChannelId} | 
[**AppsReleaseChannels**](AppsAPI.md#AppsReleaseChannels) | **Get** /apps/{appId}/release-channels | 
[**AppsStartVersionUpload**](AppsAPI.md#AppsStartVersionUpload) | **Post** /apps/{appId}/versions | 
[**AppsUpdateApp**](AppsAPI.md#AppsUpdateApp) | **Put** /apps/{appId} | 
[**AppsUpdateReleaseChannel**](AppsAPI.md#AppsUpdateReleaseChannel) | **Put** /apps/{appId}/release-channels/{releaseChannelId} | 
[**AppsVersions**](AppsAPI.md#AppsVersions) | **Get** /apps/{appId}/versions | 



## AppsApp

> AppWithDeviceModels AppsApp(ctx, appId).Execute()





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
	appId := "9d6a9043-e469-4239-9cc5-147fb20905d9" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsApp(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsApp`: AppWithDeviceModels
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppWithDeviceModels**](AppWithDeviceModels.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsApps

> AppsResponse AppsApps(ctx).PerPage(perPage).Page(page).Execute()





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
	perPage := int32(2) // int32 |  (optional)
	page := int32(1) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsApps(context.Background()).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsApps`: AppsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**AppsResponse**](AppsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsCompleteVersionUpload

> AppVersion AppsCompleteVersionUpload(ctx, appId, versionId).AppsCompleteVersionUploadRequest(appsCompleteVersionUploadRequest).Execute()





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
	appId := "9d6a9043-e469-4239-9cc5-147fb20905d9" // string | 
	versionId := "342c2a9f-559c-4fee-9d43-51ec4c4c55f9" // string | 
	appsCompleteVersionUploadRequest := *openapiclient.NewAppsCompleteVersionUploadRequest("path/to/MyApp.apk", "ABC1234567890", []openapiclient.CompleteVersionUploadRequestPart{*openapiclient.NewCompleteVersionUploadRequestPart(int32(123), "ETag_example")}) // AppsCompleteVersionUploadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsCompleteVersionUpload(context.Background(), appId, versionId).AppsCompleteVersionUploadRequest(appsCompleteVersionUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsCompleteVersionUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsCompleteVersionUpload`: AppVersion
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsCompleteVersionUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsCompleteVersionUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appsCompleteVersionUploadRequest** | [**AppsCompleteVersionUploadRequest**](AppsCompleteVersionUploadRequest.md) |  | 

### Return type

[**AppVersion**](AppVersion.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsPreSignVersionUploadUrl

> []PreSignVersionUploadUrlResponsePart AppsPreSignVersionUploadUrl(ctx, appId, versionId).AppsPreSignVersionUploadUrlRequest(appsPreSignVersionUploadUrlRequest).Execute()





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
	appId := "9d6a9043-e469-4239-9cc5-147fb20905d9" // string | 
	versionId := "342c2a9f-559c-4fee-9d43-51ec4c4c55f9" // string | 
	appsPreSignVersionUploadUrlRequest := *openapiclient.NewAppsPreSignVersionUploadUrlRequest("path/to/MyApp.apk", "ABC1234567890", []int32{int32(123)}) // AppsPreSignVersionUploadUrlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsPreSignVersionUploadUrl(context.Background(), appId, versionId).AppsPreSignVersionUploadUrlRequest(appsPreSignVersionUploadUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsPreSignVersionUploadUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsPreSignVersionUploadUrl`: []PreSignVersionUploadUrlResponsePart
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsPreSignVersionUploadUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsPreSignVersionUploadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appsPreSignVersionUploadUrlRequest** | [**AppsPreSignVersionUploadUrlRequest**](AppsPreSignVersionUploadUrlRequest.md) |  | 

### Return type

[**[]PreSignVersionUploadUrlResponsePart**](PreSignVersionUploadUrlResponsePart.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsReleaseChannel

> ReleaseChannelWithVersion AppsReleaseChannel(ctx, appId, releaseChannelId).Execute()





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
	appId := "9d6a9043-e469-4239-9cc5-147fb20905d9" // string | 
	releaseChannelId := "9d6a9043-e469-4239-9cc5-147fb20905d9" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsReleaseChannel(context.Background(), appId, releaseChannelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsReleaseChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsReleaseChannel`: ReleaseChannelWithVersion
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsReleaseChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**releaseChannelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsReleaseChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ReleaseChannelWithVersion**](ReleaseChannelWithVersion.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsReleaseChannels

> ReleaseChannelsResponse AppsReleaseChannels(ctx, appId).PerPage(perPage).Page(page).Execute()





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
	appId := "9d6a9043-e469-4239-9cc5-147fb20905d9" // string | 
	perPage := int32(2) // int32 |  (optional)
	page := int32(1) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsReleaseChannels(context.Background(), appId).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsReleaseChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsReleaseChannels`: ReleaseChannelsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsReleaseChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsReleaseChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**ReleaseChannelsResponse**](ReleaseChannelsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsStartVersionUpload

> StartVersionUploadResponse AppsStartVersionUpload(ctx, appId).AppsStartVersionUploadRequest(appsStartVersionUploadRequest).Execute()





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
	appId := "9d6a9043-e469-4239-9cc5-147fb20905d9" // string | 
	appsStartVersionUploadRequest := *openapiclient.NewAppsStartVersionUploadRequest("MyApp-v1.apk") // AppsStartVersionUploadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsStartVersionUpload(context.Background(), appId).AppsStartVersionUploadRequest(appsStartVersionUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsStartVersionUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsStartVersionUpload`: StartVersionUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsStartVersionUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsStartVersionUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appsStartVersionUploadRequest** | [**AppsStartVersionUploadRequest**](AppsStartVersionUploadRequest.md) |  | 

### Return type

[**StartVersionUploadResponse**](StartVersionUploadResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsUpdateApp

> AppWithDeviceModels AppsUpdateApp(ctx, appId).AppsUpdateAppRequest(appsUpdateAppRequest).Execute()





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
	appId := "db163242-cc1a-48ff-a6b1-db3415ea60e2" // string | 
	appsUpdateAppRequest := *openapiclient.NewAppsUpdateAppRequest() // AppsUpdateAppRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsUpdateApp(context.Background(), appId).AppsUpdateAppRequest(appsUpdateAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsUpdateApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsUpdateApp`: AppWithDeviceModels
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsUpdateApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsUpdateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appsUpdateAppRequest** | [**AppsUpdateAppRequest**](AppsUpdateAppRequest.md) |  | 

### Return type

[**AppWithDeviceModels**](AppWithDeviceModels.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsUpdateReleaseChannel

> ReleaseChannelWithVersion AppsUpdateReleaseChannel(ctx, appId, releaseChannelId).AppsUpdateReleaseChannelRequest(appsUpdateReleaseChannelRequest).Execute()





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
	appId := "9d6a9043-e469-4239-9cc5-147fb20905d9" // string | 
	releaseChannelId := "9d6a9043-e469-4239-9cc5-147fb20905d9" // string | 
	appsUpdateReleaseChannelRequest := *openapiclient.NewAppsUpdateReleaseChannelRequest("81ec530f-61f3-4fe3-9842-7517c050062c") // AppsUpdateReleaseChannelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsUpdateReleaseChannel(context.Background(), appId, releaseChannelId).AppsUpdateReleaseChannelRequest(appsUpdateReleaseChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsUpdateReleaseChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsUpdateReleaseChannel`: ReleaseChannelWithVersion
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsUpdateReleaseChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**releaseChannelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsUpdateReleaseChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appsUpdateReleaseChannelRequest** | [**AppsUpdateReleaseChannelRequest**](AppsUpdateReleaseChannelRequest.md) |  | 

### Return type

[**ReleaseChannelWithVersion**](ReleaseChannelWithVersion.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsVersions

> VersionsResponse AppsVersions(ctx, appId).PerPage(perPage).Page(page).Execute()





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
	appId := "9d6a9043-e469-4239-9cc5-147fb20905d9" // string | 
	perPage := int32(10) // int32 |  (optional)
	page := int32(1) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsVersions(context.Background(), appId).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsVersions`: VersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**VersionsResponse**](VersionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

