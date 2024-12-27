# \DevicesAPI

All URIs are relative to *https://api.xrdm.app/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesAddReleaseChannel**](DevicesAPI.md#DevicesAddReleaseChannel) | **Post** /devices/{deviceId}/release-channels | 
[**DevicesCheckFingerprint**](DevicesAPI.md#DevicesCheckFingerprint) | **Post** /devices/{deviceId}/fingerprint | 
[**DevicesDevice**](DevicesAPI.md#DevicesDevice) | **Get** /devices/{deviceId} | 
[**DevicesDevices**](DevicesAPI.md#DevicesDevices) | **Get** /devices | 
[**DevicesLaunchApp**](DevicesAPI.md#DevicesLaunchApp) | **Post** /devices/{deviceId}/launch/{appId} | 
[**DevicesReboot**](DevicesAPI.md#DevicesReboot) | **Post** /devices/{deviceId}/reboot | 
[**DevicesRemoveReleaseChannel**](DevicesAPI.md#DevicesRemoveReleaseChannel) | **Delete** /devices/{deviceId}/release-channels | 
[**DevicesUpdateDevice**](DevicesAPI.md#DevicesUpdateDevice) | **Put** /devices/{deviceId} | 



## DevicesAddReleaseChannel

> DevicesAddReleaseChannel(ctx, deviceId).DeviceGroupsAddReleaseChannelRequest(deviceGroupsAddReleaseChannelRequest).Execute()





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
	deviceId := "1ca42d4c-558a-4b50-a1e4-4b2fc7f70adc" // string | 
	deviceGroupsAddReleaseChannelRequest := *openapiclient.NewDeviceGroupsAddReleaseChannelRequest("9d6a9043-e469-4239-9cc5-147fb20905d9") // DeviceGroupsAddReleaseChannelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.DevicesAddReleaseChannel(context.Background(), deviceId).DeviceGroupsAddReleaseChannelRequest(deviceGroupsAddReleaseChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DevicesAddReleaseChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesAddReleaseChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceGroupsAddReleaseChannelRequest** | [**DeviceGroupsAddReleaseChannelRequest**](DeviceGroupsAddReleaseChannelRequest.md) |  | 

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


## DevicesCheckFingerprint

> DevicesCheckFingerprint(ctx, deviceId).DevicesCheckFingerprintRequest(devicesCheckFingerprintRequest).Execute()





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
	deviceId := "1ca42d4c-558a-4b50-a1e4-4b2fc7f70adc" // string | 
	devicesCheckFingerprintRequest := *openapiclient.NewDevicesCheckFingerprintRequest("a1b2c3d4e5f6g7h8i9j0") // DevicesCheckFingerprintRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.DevicesCheckFingerprint(context.Background(), deviceId).DevicesCheckFingerprintRequest(devicesCheckFingerprintRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DevicesCheckFingerprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesCheckFingerprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **devicesCheckFingerprintRequest** | [**DevicesCheckFingerprintRequest**](DevicesCheckFingerprintRequest.md) |  | 

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


## DevicesDevice

> DeviceWithSimplifiedDeviceGroup DevicesDevice(ctx, deviceId).Execute()





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
	deviceId := "9d6a9043-e469-4239-9cc5-147fb20905d9" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.DevicesDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DevicesDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesDevice`: DeviceWithSimplifiedDeviceGroup
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.DevicesDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceWithSimplifiedDeviceGroup**](DeviceWithSimplifiedDeviceGroup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesDevices

> DevicesResponse DevicesDevices(ctx).PerPage(perPage).Page(page).Execute()





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
	resp, r, err := apiClient.DevicesAPI.DevicesDevices(context.Background()).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DevicesDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesDevices`: DevicesResponse
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.DevicesDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**DevicesResponse**](DevicesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesLaunchApp

> DevicesLaunchApp(ctx, deviceId, appId).Execute()





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
	deviceId := "1ca42d4c-558a-4b50-a1e4-4b2fc7f70adc" // string | 
	appId := "9d6a9043-e469-4239-9cc5-147fb20905d9" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.DevicesLaunchApp(context.Background(), deviceId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DevicesLaunchApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesLaunchAppRequest struct via the builder pattern


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


## DevicesReboot

> DevicesReboot(ctx, deviceId).Execute()





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
	deviceId := "1ca42d4c-558a-4b50-a1e4-4b2fc7f70adc" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.DevicesReboot(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DevicesReboot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesRebootRequest struct via the builder pattern


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


## DevicesRemoveReleaseChannel

> DevicesRemoveReleaseChannel(ctx, deviceId).DeviceGroupsAddReleaseChannelRequest(deviceGroupsAddReleaseChannelRequest).Execute()





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
	deviceId := "1ca42d4c-558a-4b50-a1e4-4b2fc7f70adc" // string | 
	deviceGroupsAddReleaseChannelRequest := *openapiclient.NewDeviceGroupsAddReleaseChannelRequest("9d6a9043-e469-4239-9cc5-147fb20905d9") // DeviceGroupsAddReleaseChannelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.DevicesRemoveReleaseChannel(context.Background(), deviceId).DeviceGroupsAddReleaseChannelRequest(deviceGroupsAddReleaseChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DevicesRemoveReleaseChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesRemoveReleaseChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceGroupsAddReleaseChannelRequest** | [**DeviceGroupsAddReleaseChannelRequest**](DeviceGroupsAddReleaseChannelRequest.md) |  | 

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


## DevicesUpdateDevice

> DeviceWithSimplifiedDeviceGroup DevicesUpdateDevice(ctx, deviceId).DevicesUpdateDeviceRequest(devicesUpdateDeviceRequest).Execute()





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
	deviceId := "9d6a9043-e469-4239-9cc5-147fb20905d9" // string | 
	devicesUpdateDeviceRequest := *openapiclient.NewDevicesUpdateDeviceRequest() // DevicesUpdateDeviceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.DevicesUpdateDevice(context.Background(), deviceId).DevicesUpdateDeviceRequest(devicesUpdateDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DevicesUpdateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesUpdateDevice`: DeviceWithSimplifiedDeviceGroup
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.DevicesUpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **devicesUpdateDeviceRequest** | [**DevicesUpdateDeviceRequest**](DevicesUpdateDeviceRequest.md) |  | 

### Return type

[**DeviceWithSimplifiedDeviceGroup**](DeviceWithSimplifiedDeviceGroup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

