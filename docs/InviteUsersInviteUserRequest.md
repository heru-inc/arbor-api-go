# InviteUsersInviteUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**OrganizationRoleId** | Pointer to **NullableString** |  | [optional] 
**Sso** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewInviteUsersInviteUserRequest

`func NewInviteUsersInviteUserRequest(email string, ) *InviteUsersInviteUserRequest`

NewInviteUsersInviteUserRequest instantiates a new InviteUsersInviteUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUsersInviteUserRequestWithDefaults

`func NewInviteUsersInviteUserRequestWithDefaults() *InviteUsersInviteUserRequest`

NewInviteUsersInviteUserRequestWithDefaults instantiates a new InviteUsersInviteUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteUsersInviteUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteUsersInviteUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteUsersInviteUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOrganizationRoleId

`func (o *InviteUsersInviteUserRequest) GetOrganizationRoleId() string`

GetOrganizationRoleId returns the OrganizationRoleId field if non-nil, zero value otherwise.

### GetOrganizationRoleIdOk

`func (o *InviteUsersInviteUserRequest) GetOrganizationRoleIdOk() (*string, bool)`

GetOrganizationRoleIdOk returns a tuple with the OrganizationRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoleId

`func (o *InviteUsersInviteUserRequest) SetOrganizationRoleId(v string)`

SetOrganizationRoleId sets OrganizationRoleId field to given value.

### HasOrganizationRoleId

`func (o *InviteUsersInviteUserRequest) HasOrganizationRoleId() bool`

HasOrganizationRoleId returns a boolean if a field has been set.

### SetOrganizationRoleIdNil

`func (o *InviteUsersInviteUserRequest) SetOrganizationRoleIdNil(b bool)`

 SetOrganizationRoleIdNil sets the value for OrganizationRoleId to be an explicit nil

### UnsetOrganizationRoleId
`func (o *InviteUsersInviteUserRequest) UnsetOrganizationRoleId()`

UnsetOrganizationRoleId ensures that no value is present for OrganizationRoleId, not even an explicit nil
### GetSso

`func (o *InviteUsersInviteUserRequest) GetSso() bool`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *InviteUsersInviteUserRequest) GetSsoOk() (*bool, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *InviteUsersInviteUserRequest) SetSso(v bool)`

SetSso sets Sso field to given value.

### HasSso

`func (o *InviteUsersInviteUserRequest) HasSso() bool`

HasSso returns a boolean if a field has been set.

### SetSsoNil

`func (o *InviteUsersInviteUserRequest) SetSsoNil(b bool)`

 SetSsoNil sets the value for Sso to be an explicit nil

### UnsetSso
`func (o *InviteUsersInviteUserRequest) UnsetSso()`

UnsetSso ensures that no value is present for Sso, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


