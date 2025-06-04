# InviteUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Sso** | Pointer to **bool** | Should the user be required to authenticate using the organization&#39;s configured IdP? | [optional] 
**OrganizationRoleId** | Pointer to **string** | Required when &#x60;groupRoleId&#x60; &amp; &#x60;groupIds&#x60; are not provided. | [optional] 
**GroupRoleId** | Pointer to **string** | Required when &#x60;organizationRoleId&#x60; is not provided. | [optional] 
**GroupIds** | Pointer to **[]string** | Required when &#x60;organizationRoleId&#x60; is not provided. | [optional] 

## Methods

### NewInviteUserRequest

`func NewInviteUserRequest(email string, ) *InviteUserRequest`

NewInviteUserRequest instantiates a new InviteUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUserRequestWithDefaults

`func NewInviteUserRequestWithDefaults() *InviteUserRequest`

NewInviteUserRequestWithDefaults instantiates a new InviteUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetSso

`func (o *InviteUserRequest) GetSso() bool`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *InviteUserRequest) GetSsoOk() (*bool, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *InviteUserRequest) SetSso(v bool)`

SetSso sets Sso field to given value.

### HasSso

`func (o *InviteUserRequest) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetOrganizationRoleId

`func (o *InviteUserRequest) GetOrganizationRoleId() string`

GetOrganizationRoleId returns the OrganizationRoleId field if non-nil, zero value otherwise.

### GetOrganizationRoleIdOk

`func (o *InviteUserRequest) GetOrganizationRoleIdOk() (*string, bool)`

GetOrganizationRoleIdOk returns a tuple with the OrganizationRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoleId

`func (o *InviteUserRequest) SetOrganizationRoleId(v string)`

SetOrganizationRoleId sets OrganizationRoleId field to given value.

### HasOrganizationRoleId

`func (o *InviteUserRequest) HasOrganizationRoleId() bool`

HasOrganizationRoleId returns a boolean if a field has been set.

### GetGroupRoleId

`func (o *InviteUserRequest) GetGroupRoleId() string`

GetGroupRoleId returns the GroupRoleId field if non-nil, zero value otherwise.

### GetGroupRoleIdOk

`func (o *InviteUserRequest) GetGroupRoleIdOk() (*string, bool)`

GetGroupRoleIdOk returns a tuple with the GroupRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRoleId

`func (o *InviteUserRequest) SetGroupRoleId(v string)`

SetGroupRoleId sets GroupRoleId field to given value.

### HasGroupRoleId

`func (o *InviteUserRequest) HasGroupRoleId() bool`

HasGroupRoleId returns a boolean if a field has been set.

### GetGroupIds

`func (o *InviteUserRequest) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *InviteUserRequest) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *InviteUserRequest) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *InviteUserRequest) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


