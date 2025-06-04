# \DeviceModelsAPI

All URIs are relative to *https://api.xrdm.app/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceModel**](DeviceModelsAPI.md#GetDeviceModel) | **Get** /device-models/{deviceModelId} | 
[**GetDeviceModels**](DeviceModelsAPI.md#GetDeviceModels) | **Get** /device-models | 



## GetDeviceModel

> GetApps200ResponseDataInnerDeviceModelsInner GetDeviceModel(ctx, deviceModelId).Execute()





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
	deviceModelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a device model.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceModelsAPI.GetDeviceModel(context.Background(), deviceModelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceModelsAPI.GetDeviceModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceModel`: GetApps200ResponseDataInnerDeviceModelsInner
	fmt.Fprintf(os.Stdout, "Response from `DeviceModelsAPI.GetDeviceModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceModelId** | **string** | The ID of a device model. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetApps200ResponseDataInnerDeviceModelsInner**](GetApps200ResponseDataInnerDeviceModelsInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceModels

> GetDeviceModels200Response GetDeviceModels(ctx).PerPage(perPage).Page(page).Execute()





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
	resp, r, err := apiClient.DeviceModelsAPI.GetDeviceModels(context.Background()).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceModelsAPI.GetDeviceModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceModels`: GetDeviceModels200Response
	fmt.Fprintf(os.Stdout, "Response from `DeviceModelsAPI.GetDeviceModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of items to return per page. | [default to 10]
 **page** | **int32** | The page number to return. | [default to 1]

### Return type

[**GetDeviceModels200Response**](GetDeviceModels200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

