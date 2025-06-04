# \DevicesAPI

All URIs are relative to *https://api.xrdm.app/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFileToDevice**](DevicesAPI.md#AddFileToDevice) | **Post** /devices/{deviceId}/files | 
[**AddReleaseChannelToDevice**](DevicesAPI.md#AddReleaseChannelToDevice) | **Post** /devices/{deviceId}/release-channels | 
[**CheckFingerprint**](DevicesAPI.md#CheckFingerprint) | **Post** /devices/{deviceId}/fingerprint | 
[**GetDevice**](DevicesAPI.md#GetDevice) | **Get** /devices/{deviceId} | 
[**GetDevices**](DevicesAPI.md#GetDevices) | **Get** /devices | 
[**LaunchApp**](DevicesAPI.md#LaunchApp) | **Post** /devices/{deviceId}/launch/{appId} | 
[**RebootDevice**](DevicesAPI.md#RebootDevice) | **Post** /devices/{deviceId}/reboot | 
[**RemoveFileFromDevice**](DevicesAPI.md#RemoveFileFromDevice) | **Delete** /devices/{deviceId}/files | 
[**RemoveReleaseChannelFromDevice**](DevicesAPI.md#RemoveReleaseChannelFromDevice) | **Delete** /devices/{deviceId}/release-channels | 
[**UpdateDevice**](DevicesAPI.md#UpdateDevice) | **Put** /devices/{deviceId} | 



## AddFileToDevice

> AddFileToDevice(ctx, deviceId).AddFileToGroupRequest(addFileToGroupRequest).Execute()





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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an device.
	addFileToGroupRequest := *openapiclient.NewAddFileToGroupRequest("123e4567-e89b-12d3-a456-426614174000") // AddFileToGroupRequest | The details of the file you want to add to or remove from the device. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.AddFileToDevice(context.Background(), deviceId).AddFileToGroupRequest(addFileToGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AddFileToDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The ID of an device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFileToDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addFileToGroupRequest** | [**AddFileToGroupRequest**](AddFileToGroupRequest.md) | The details of the file you want to add to or remove from the device. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddReleaseChannelToDevice

> AddReleaseChannelToDevice(ctx, deviceId).AddReleaseChannelToDeviceGroupRequest(addReleaseChannelToDeviceGroupRequest).Execute()





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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an device.
	addReleaseChannelToDeviceGroupRequest := *openapiclient.NewAddReleaseChannelToDeviceGroupRequest("123e4567-e89b-12d3-a456-426614174000") // AddReleaseChannelToDeviceGroupRequest | The details of the release channel you want to add to or remove from the device. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.AddReleaseChannelToDevice(context.Background(), deviceId).AddReleaseChannelToDeviceGroupRequest(addReleaseChannelToDeviceGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.AddReleaseChannelToDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The ID of an device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddReleaseChannelToDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addReleaseChannelToDeviceGroupRequest** | [**AddReleaseChannelToDeviceGroupRequest**](AddReleaseChannelToDeviceGroupRequest.md) | The details of the release channel you want to add to or remove from the device. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckFingerprint

> CheckFingerprint(ctx, deviceId).CheckFingerprintRequest(checkFingerprintRequest).Execute()





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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an device.
	checkFingerprintRequest := *openapiclient.NewCheckFingerprintRequest("d41d8cd98f00b204e9800998ecf8427e") // CheckFingerprintRequest | The fingerprint to check. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.CheckFingerprint(context.Background(), deviceId).CheckFingerprintRequest(checkFingerprintRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.CheckFingerprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The ID of an device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckFingerprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkFingerprintRequest** | [**CheckFingerprintRequest**](CheckFingerprintRequest.md) | The fingerprint to check. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevice

> GetDevices200ResponseDataInner GetDevice(ctx, deviceId).Execute()





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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an device.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevice`: GetDevices200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The ID of an device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDevices200ResponseDataInner**](GetDevices200ResponseDataInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevices

> GetDevices200Response GetDevices(ctx).PerPage(perPage).Page(page).Execute()





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
	resp, r, err := apiClient.DevicesAPI.GetDevices(context.Background()).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevices`: GetDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of items to return per page. | [default to 10]
 **page** | **int32** | The page number to return. | [default to 1]

### Return type

[**GetDevices200Response**](GetDevices200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchApp

> LaunchApp(ctx, deviceId, appId).Execute()





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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an device.
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.LaunchApp(context.Background(), deviceId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.LaunchApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The ID of an device. | 
**appId** | **string** | The ID of an app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootDevice

> RebootDevice(ctx, deviceId).Execute()





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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an device.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.RebootDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.RebootDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The ID of an device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFileFromDevice

> RemoveFileFromDevice(ctx, deviceId).AddFileToGroupRequest(addFileToGroupRequest).Execute()





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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an device.
	addFileToGroupRequest := *openapiclient.NewAddFileToGroupRequest("123e4567-e89b-12d3-a456-426614174000") // AddFileToGroupRequest | The details of the file you want to add to or remove from the device. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.RemoveFileFromDevice(context.Background(), deviceId).AddFileToGroupRequest(addFileToGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.RemoveFileFromDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The ID of an device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFileFromDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addFileToGroupRequest** | [**AddFileToGroupRequest**](AddFileToGroupRequest.md) | The details of the file you want to add to or remove from the device. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveReleaseChannelFromDevice

> RemoveReleaseChannelFromDevice(ctx, deviceId).AddReleaseChannelToDeviceGroupRequest(addReleaseChannelToDeviceGroupRequest).Execute()





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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an device.
	addReleaseChannelToDeviceGroupRequest := *openapiclient.NewAddReleaseChannelToDeviceGroupRequest("123e4567-e89b-12d3-a456-426614174000") // AddReleaseChannelToDeviceGroupRequest | The details of the release channel you want to add to or remove from the device. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.RemoveReleaseChannelFromDevice(context.Background(), deviceId).AddReleaseChannelToDeviceGroupRequest(addReleaseChannelToDeviceGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.RemoveReleaseChannelFromDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The ID of an device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveReleaseChannelFromDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addReleaseChannelToDeviceGroupRequest** | [**AddReleaseChannelToDeviceGroupRequest**](AddReleaseChannelToDeviceGroupRequest.md) | The details of the release channel you want to add to or remove from the device. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> GetDevices200ResponseDataInner UpdateDevice(ctx, deviceId).UpdateDeviceRequest(updateDeviceRequest).Execute()





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
	deviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of an device.
	updateDeviceRequest := *openapiclient.NewUpdateDeviceRequest() // UpdateDeviceRequest | The device fields to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.UpdateDevice(context.Background(), deviceId).UpdateDeviceRequest(updateDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.UpdateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDevice`: GetDevices200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.UpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The ID of an device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceRequest** | [**UpdateDeviceRequest**](UpdateDeviceRequest.md) | The device fields to update. | 

### Return type

[**GetDevices200ResponseDataInner**](GetDevices200ResponseDataInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

