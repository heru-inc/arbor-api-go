# \GroupsAPI

All URIs are relative to *https://api.xrdm.app/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFileToGroup**](GroupsAPI.md#AddFileToGroup) | **Post** /groups/{groupId}/files | 
[**AddReleaseChannelToDeviceGroup**](GroupsAPI.md#AddReleaseChannelToDeviceGroup) | **Post** /device-groups/{groupId}/release-channels | 
[**AddReleaseChannelToGroup**](GroupsAPI.md#AddReleaseChannelToGroup) | **Post** /groups/{groupId}/release-channels | 
[**GetDeviceGroups**](GroupsAPI.md#GetDeviceGroups) | **Get** /device-groups | 
[**GetGroupReleaseChannels**](GroupsAPI.md#GetGroupReleaseChannels) | **Get** /groups/{groupId}/release-channels | 
[**GetGroups**](GroupsAPI.md#GetGroups) | **Get** /groups | 
[**RemoveFileFromGroup**](GroupsAPI.md#RemoveFileFromGroup) | **Delete** /groups/{groupId}/files | 
[**RemoveReleaseChannelFromDeviceGroup**](GroupsAPI.md#RemoveReleaseChannelFromDeviceGroup) | **Delete** /device-groups/{groupId}/release-channels | 
[**RemoveReleaseChannelFromGroup**](GroupsAPI.md#RemoveReleaseChannelFromGroup) | **Delete** /groups/{groupId}/release-channels | 



## AddFileToGroup

> AddFileToGroup(ctx, groupId).AddFileToGroupRequest(addFileToGroupRequest).Execute()





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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a group.
	addFileToGroupRequest := *openapiclient.NewAddFileToGroupRequest("123e4567-e89b-12d3-a456-426614174000") // AddFileToGroupRequest | The details of the file you want to add to or remove from the group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.AddFileToGroup(context.Background(), groupId).AddFileToGroupRequest(addFileToGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddFileToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of a group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFileToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addFileToGroupRequest** | [**AddFileToGroupRequest**](AddFileToGroupRequest.md) | The details of the file you want to add to or remove from the group. | 

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


## AddReleaseChannelToDeviceGroup

> AddReleaseChannelToDeviceGroup(ctx, groupId).AddReleaseChannelToDeviceGroupRequest(addReleaseChannelToDeviceGroupRequest).Execute()





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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a group.
	addReleaseChannelToDeviceGroupRequest := *openapiclient.NewAddReleaseChannelToDeviceGroupRequest("123e4567-e89b-12d3-a456-426614174000") // AddReleaseChannelToDeviceGroupRequest | The details of the release channel you want to add to or remove from the group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.AddReleaseChannelToDeviceGroup(context.Background(), groupId).AddReleaseChannelToDeviceGroupRequest(addReleaseChannelToDeviceGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddReleaseChannelToDeviceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of a group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddReleaseChannelToDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addReleaseChannelToDeviceGroupRequest** | [**AddReleaseChannelToDeviceGroupRequest**](AddReleaseChannelToDeviceGroupRequest.md) | The details of the release channel you want to add to or remove from the group. | 

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


## AddReleaseChannelToGroup

> AddReleaseChannelToGroup(ctx, groupId).AddReleaseChannelToDeviceGroupRequest(addReleaseChannelToDeviceGroupRequest).Execute()





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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a group.
	addReleaseChannelToDeviceGroupRequest := *openapiclient.NewAddReleaseChannelToDeviceGroupRequest("123e4567-e89b-12d3-a456-426614174000") // AddReleaseChannelToDeviceGroupRequest | The details of the release channel you want to add to or remove from the group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.AddReleaseChannelToGroup(context.Background(), groupId).AddReleaseChannelToDeviceGroupRequest(addReleaseChannelToDeviceGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddReleaseChannelToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of a group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddReleaseChannelToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addReleaseChannelToDeviceGroupRequest** | [**AddReleaseChannelToDeviceGroupRequest**](AddReleaseChannelToDeviceGroupRequest.md) | The details of the release channel you want to add to or remove from the group. | 

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


## GetDeviceGroups

> GetDeviceGroups200Response GetDeviceGroups(ctx).PerPage(perPage).Page(page).Execute()





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
	resp, r, err := apiClient.GroupsAPI.GetDeviceGroups(context.Background()).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetDeviceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceGroups`: GetDeviceGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetDeviceGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of items to return per page. | [default to 10]
 **page** | **int32** | The page number to return. | [default to 1]

### Return type

[**GetDeviceGroups200Response**](GetDeviceGroups200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupReleaseChannels

> GetGroupReleaseChannels200Response GetGroupReleaseChannels(ctx, groupId).Execute()





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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupReleaseChannels(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupReleaseChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupReleaseChannels`: GetGroupReleaseChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupReleaseChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of a group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupReleaseChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetGroupReleaseChannels200Response**](GetGroupReleaseChannels200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> GetDeviceGroups200Response GetGroups(ctx).PerPage(perPage).Page(page).Execute()





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
	resp, r, err := apiClient.GroupsAPI.GetGroups(context.Background()).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroups`: GetDeviceGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of items to return per page. | [default to 10]
 **page** | **int32** | The page number to return. | [default to 1]

### Return type

[**GetDeviceGroups200Response**](GetDeviceGroups200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFileFromGroup

> RemoveFileFromGroup(ctx, groupId).AddFileToGroupRequest(addFileToGroupRequest).Execute()





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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a group.
	addFileToGroupRequest := *openapiclient.NewAddFileToGroupRequest("123e4567-e89b-12d3-a456-426614174000") // AddFileToGroupRequest | The details of the file you want to add to or remove from the group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.RemoveFileFromGroup(context.Background(), groupId).AddFileToGroupRequest(addFileToGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.RemoveFileFromGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of a group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFileFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addFileToGroupRequest** | [**AddFileToGroupRequest**](AddFileToGroupRequest.md) | The details of the file you want to add to or remove from the group. | 

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


## RemoveReleaseChannelFromDeviceGroup

> RemoveReleaseChannelFromDeviceGroup(ctx, groupId).AddReleaseChannelToDeviceGroupRequest(addReleaseChannelToDeviceGroupRequest).Execute()





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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a group.
	addReleaseChannelToDeviceGroupRequest := *openapiclient.NewAddReleaseChannelToDeviceGroupRequest("123e4567-e89b-12d3-a456-426614174000") // AddReleaseChannelToDeviceGroupRequest | The details of the release channel you want to add to or remove from the group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.RemoveReleaseChannelFromDeviceGroup(context.Background(), groupId).AddReleaseChannelToDeviceGroupRequest(addReleaseChannelToDeviceGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.RemoveReleaseChannelFromDeviceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of a group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveReleaseChannelFromDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addReleaseChannelToDeviceGroupRequest** | [**AddReleaseChannelToDeviceGroupRequest**](AddReleaseChannelToDeviceGroupRequest.md) | The details of the release channel you want to add to or remove from the group. | 

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


## RemoveReleaseChannelFromGroup

> RemoveReleaseChannelFromGroup(ctx, groupId).AddReleaseChannelToDeviceGroupRequest(addReleaseChannelToDeviceGroupRequest).Execute()





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
	groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a group.
	addReleaseChannelToDeviceGroupRequest := *openapiclient.NewAddReleaseChannelToDeviceGroupRequest("123e4567-e89b-12d3-a456-426614174000") // AddReleaseChannelToDeviceGroupRequest | The details of the release channel you want to add to or remove from the group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.RemoveReleaseChannelFromGroup(context.Background(), groupId).AddReleaseChannelToDeviceGroupRequest(addReleaseChannelToDeviceGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.RemoveReleaseChannelFromGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of a group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveReleaseChannelFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addReleaseChannelToDeviceGroupRequest** | [**AddReleaseChannelToDeviceGroupRequest**](AddReleaseChannelToDeviceGroupRequest.md) | The details of the release channel you want to add to or remove from the group. | 

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

