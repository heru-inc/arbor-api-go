# \DeviceModelsAPI

All URIs are relative to *https://api.xrdm.app/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceModelsDeviceModel**](DeviceModelsAPI.md#DeviceModelsDeviceModel) | **Get** /device-models/{deviceModelId} | 
[**DeviceModelsDeviceModels**](DeviceModelsAPI.md#DeviceModelsDeviceModels) | **Get** /device-models | 



## DeviceModelsDeviceModel

> DeviceModel DeviceModelsDeviceModel(ctx, deviceModelId).Execute()





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
	deviceModelId := "660c96fc-8775-49ae-9bf0-863c37cf528f" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceModelsAPI.DeviceModelsDeviceModel(context.Background(), deviceModelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceModelsAPI.DeviceModelsDeviceModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceModelsDeviceModel`: DeviceModel
	fmt.Fprintf(os.Stdout, "Response from `DeviceModelsAPI.DeviceModelsDeviceModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceModelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceModelsDeviceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceModel**](DeviceModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceModelsDeviceModels

> DeviceModelsResponse DeviceModelsDeviceModels(ctx).PerPage(perPage).Page(page).Execute()





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
	resp, r, err := apiClient.DeviceModelsAPI.DeviceModelsDeviceModels(context.Background()).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceModelsAPI.DeviceModelsDeviceModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceModelsDeviceModels`: DeviceModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceModelsAPI.DeviceModelsDeviceModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceModelsDeviceModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**DeviceModelsResponse**](DeviceModelsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

