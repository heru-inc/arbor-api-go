# \DeviceGroupsAPI

All URIs are relative to *https://api.xrdm.app/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceGroupsAddReleaseChannel**](DeviceGroupsAPI.md#DeviceGroupsAddReleaseChannel) | **Post** /device-groups/{deviceGroupId}/release-channels | 
[**DeviceGroupsGroups**](DeviceGroupsAPI.md#DeviceGroupsGroups) | **Get** /device-groups | 
[**DeviceGroupsRemoveReleaseChannel**](DeviceGroupsAPI.md#DeviceGroupsRemoveReleaseChannel) | **Delete** /device-groups/{deviceGroupId}/release-channels | 



## DeviceGroupsAddReleaseChannel

> DeviceGroupsAddReleaseChannel(ctx, deviceGroupId).DeviceGroupsAddReleaseChannelRequest(deviceGroupsAddReleaseChannelRequest).Execute()





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
	deviceGroupId := "1ca42d4c-558a-4b50-a1e4-4b2fc7f70adc" // string | 
	deviceGroupsAddReleaseChannelRequest := *openapiclient.NewDeviceGroupsAddReleaseChannelRequest("9d6a9043-e469-4239-9cc5-147fb20905d9") // DeviceGroupsAddReleaseChannelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceGroupsAPI.DeviceGroupsAddReleaseChannel(context.Background(), deviceGroupId).DeviceGroupsAddReleaseChannelRequest(deviceGroupsAddReleaseChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceGroupsAPI.DeviceGroupsAddReleaseChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceGroupsAddReleaseChannelRequest struct via the builder pattern


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


## DeviceGroupsGroups

> DeviceGroupsResponse DeviceGroupsGroups(ctx).PerPage(perPage).Page(page).Execute()





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
	resp, r, err := apiClient.DeviceGroupsAPI.DeviceGroupsGroups(context.Background()).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceGroupsAPI.DeviceGroupsGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceGroupsGroups`: DeviceGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceGroupsAPI.DeviceGroupsGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceGroupsGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**DeviceGroupsResponse**](DeviceGroupsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceGroupsRemoveReleaseChannel

> DeviceGroupsRemoveReleaseChannel(ctx, deviceGroupId).DeviceGroupsAddReleaseChannelRequest(deviceGroupsAddReleaseChannelRequest).Execute()





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
	deviceGroupId := "1ca42d4c-558a-4b50-a1e4-4b2fc7f70adc" // string | 
	deviceGroupsAddReleaseChannelRequest := *openapiclient.NewDeviceGroupsAddReleaseChannelRequest("9d6a9043-e469-4239-9cc5-147fb20905d9") // DeviceGroupsAddReleaseChannelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceGroupsAPI.DeviceGroupsRemoveReleaseChannel(context.Background(), deviceGroupId).DeviceGroupsAddReleaseChannelRequest(deviceGroupsAddReleaseChannelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceGroupsAPI.DeviceGroupsRemoveReleaseChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceGroupsRemoveReleaseChannelRequest struct via the builder pattern


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

