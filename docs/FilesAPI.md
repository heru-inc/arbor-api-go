# \FilesAPI

All URIs are relative to *https://api.xrdm.app/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteFileUpload**](FilesAPI.md#CompleteFileUpload) | **Post** /files/{fileId}/complete | 
[**GetFile**](FilesAPI.md#GetFile) | **Get** /files/{fileId} | 
[**GetFiles**](FilesAPI.md#GetFiles) | **Get** /files | 
[**InitiateFileUpload**](FilesAPI.md#InitiateFileUpload) | **Post** /files | 
[**PreSignFileUpload**](FilesAPI.md#PreSignFileUpload) | **Post** /files/{fileId}/pre-sign | 



## CompleteFileUpload

> GetFiles200ResponseDataInner CompleteFileUpload(ctx, fileId).CompleteFileUploadRequest(completeFileUploadRequest).Execute()





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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a file.
	completeFileUploadRequest := *openapiclient.NewCompleteFileUploadRequest("xrdm-016da47e-4c49-48ad-9c61-94904e06a226/files/2f5d24ec-be98-49d3-ac42-cdecbeea40f0/f40d390d-94fc-47a9-952c-8b7321f3eefc/my-file.txt", "OGRhZDQ5NTMtOTQxYS00ZDI3LWE0MzUtOGUyNzEzZWZiYTlmLmY1YmE5NTkyLWIzZDMtNDE1ZS1iOWIxLWM1YWJjYjAxMmY5OQ", []openapiclient.CompleteAppVersionUploadRequestPartsInner{*openapiclient.NewCompleteAppVersionUploadRequestPartsInner(int32(1), "d41d8cd98f00b204e9800998ecf8427e")}) // CompleteFileUploadRequest | The fields to complete the upload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.CompleteFileUpload(context.Background(), fileId).CompleteFileUploadRequest(completeFileUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.CompleteFileUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteFileUpload`: GetFiles200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.CompleteFileUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The ID of a file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteFileUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **completeFileUploadRequest** | [**CompleteFileUploadRequest**](CompleteFileUploadRequest.md) | The fields to complete the upload. | 

### Return type

[**GetFiles200ResponseDataInner**](GetFiles200ResponseDataInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFile

> GetFiles200ResponseDataInner GetFile(ctx, fileId).Execute()





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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFile(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFile`: GetFiles200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The ID of a file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFiles200ResponseDataInner**](GetFiles200ResponseDataInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiles

> GetFiles200Response GetFiles(ctx).PerPage(perPage).Page(page).Execute()





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
	resp, r, err := apiClient.FilesAPI.GetFiles(context.Background()).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiles`: GetFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of items to return per page. | [default to 10]
 **page** | **int32** | The page number to return. | [default to 1]

### Return type

[**GetFiles200Response**](GetFiles200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateFileUpload

> InitiateFileUpload200Response InitiateFileUpload(ctx).InitiateFileUploadRequest(initiateFileUploadRequest).Execute()





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
	initiateFileUploadRequest := *openapiclient.NewInitiateFileUploadRequest("my-app.apk", "/sdcard/my-files/some-sub-directory") // InitiateFileUploadRequest | Initiate a new file upload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.InitiateFileUpload(context.Background()).InitiateFileUploadRequest(initiateFileUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.InitiateFileUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateFileUpload`: InitiateFileUpload200Response
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.InitiateFileUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateFileUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **initiateFileUploadRequest** | [**InitiateFileUploadRequest**](InitiateFileUploadRequest.md) | Initiate a new file upload. | 

### Return type

[**InitiateFileUpload200Response**](InitiateFileUpload200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreSignFileUpload

> []PreSignFileUpload200ResponseInner PreSignFileUpload(ctx, fileId).PreSignFileUploadRequest(preSignFileUploadRequest).Execute()





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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of a file.
	preSignFileUploadRequest := *openapiclient.NewPreSignFileUploadRequest("xrdm-016da47e-4c49-48ad-9c61-94904e06a226/files/2f5d24ec-be98-49d3-ac42-cdecbeea40f0/f40d390d-94fc-47a9-952c-8b7321f3eefc/my-file.txt", "OGRhZDQ5NTMtOTQxYS00ZDI3LWE0MzUtOGUyNzEzZWZiYTlmLmY1YmE5NTkyLWIzZDMtNDE1ZS1iOWIxLWM1YWJjYjAxMmY5OQ", []int32{int32(1)}) // PreSignFileUploadRequest | The fields to retrieve presigned URLs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.PreSignFileUpload(context.Background(), fileId).PreSignFileUploadRequest(preSignFileUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.PreSignFileUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreSignFileUpload`: []PreSignFileUpload200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.PreSignFileUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The ID of a file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreSignFileUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preSignFileUploadRequest** | [**PreSignFileUploadRequest**](PreSignFileUploadRequest.md) | The fields to retrieve presigned URLs. | 

### Return type

[**[]PreSignFileUpload200ResponseInner**](PreSignFileUpload200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

