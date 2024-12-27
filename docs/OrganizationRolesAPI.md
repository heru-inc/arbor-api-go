# \OrganizationRolesAPI

All URIs are relative to *https://api.xrdm.app/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationRolesOrganizationRole**](OrganizationRolesAPI.md#OrganizationRolesOrganizationRole) | **Get** /organization-roles/{organizationRoleId} | 
[**OrganizationRolesOrganizationRoles**](OrganizationRolesAPI.md#OrganizationRolesOrganizationRoles) | **Get** /organization-roles | 



## OrganizationRolesOrganizationRole

> OrganizationRole OrganizationRolesOrganizationRole(ctx, organizationRoleId).Execute()





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
	organizationRoleId := "660c96fc-8775-49ae-9bf0-863c37cf528f" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationRolesAPI.OrganizationRolesOrganizationRole(context.Background(), organizationRoleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.OrganizationRolesOrganizationRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationRolesOrganizationRole`: OrganizationRole
	fmt.Fprintf(os.Stdout, "Response from `OrganizationRolesAPI.OrganizationRolesOrganizationRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationRoleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationRolesOrganizationRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationRole**](OrganizationRole.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationRolesOrganizationRoles

> OrganizationRolesResponse OrganizationRolesOrganizationRoles(ctx).PerPage(perPage).Page(page).Execute()





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
	resp, r, err := apiClient.OrganizationRolesAPI.OrganizationRolesOrganizationRoles(context.Background()).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationRolesAPI.OrganizationRolesOrganizationRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationRolesOrganizationRoles`: OrganizationRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationRolesAPI.OrganizationRolesOrganizationRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationRolesOrganizationRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**OrganizationRolesResponse**](OrganizationRolesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

