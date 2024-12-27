# \InviteUsersAPI

All URIs are relative to *https://api.xrdm.app/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InviteUsersInviteUser**](InviteUsersAPI.md#InviteUsersInviteUser) | **Post** /invite-users | 



## InviteUsersInviteUser

> InviteUsersInviteUser(ctx).InviteUsersInviteUserRequest(inviteUsersInviteUserRequest).Execute()





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
	inviteUsersInviteUserRequest := *openapiclient.NewInviteUsersInviteUserRequest("Email_example") // InviteUsersInviteUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InviteUsersAPI.InviteUsersInviteUser(context.Background()).InviteUsersInviteUserRequest(inviteUsersInviteUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InviteUsersAPI.InviteUsersInviteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteUsersInviteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteUsersInviteUserRequest** | [**InviteUsersInviteUserRequest**](InviteUsersInviteUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

