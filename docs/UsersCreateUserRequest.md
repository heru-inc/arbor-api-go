# UsersCreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Email** | **string** |  | 
**OrganizationRoleId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUsersCreateUserRequest

`func NewUsersCreateUserRequest(firstName string, lastName string, email string, ) *UsersCreateUserRequest`

NewUsersCreateUserRequest instantiates a new UsersCreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersCreateUserRequestWithDefaults

`func NewUsersCreateUserRequestWithDefaults() *UsersCreateUserRequest`

NewUsersCreateUserRequestWithDefaults instantiates a new UsersCreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UsersCreateUserRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UsersCreateUserRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UsersCreateUserRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UsersCreateUserRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UsersCreateUserRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UsersCreateUserRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *UsersCreateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UsersCreateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UsersCreateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOrganizationRoleId

`func (o *UsersCreateUserRequest) GetOrganizationRoleId() string`

GetOrganizationRoleId returns the OrganizationRoleId field if non-nil, zero value otherwise.

### GetOrganizationRoleIdOk

`func (o *UsersCreateUserRequest) GetOrganizationRoleIdOk() (*string, bool)`

GetOrganizationRoleIdOk returns a tuple with the OrganizationRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoleId

`func (o *UsersCreateUserRequest) SetOrganizationRoleId(v string)`

SetOrganizationRoleId sets OrganizationRoleId field to given value.

### HasOrganizationRoleId

`func (o *UsersCreateUserRequest) HasOrganizationRoleId() bool`

HasOrganizationRoleId returns a boolean if a field has been set.

### SetOrganizationRoleIdNil

`func (o *UsersCreateUserRequest) SetOrganizationRoleIdNil(b bool)`

 SetOrganizationRoleIdNil sets the value for OrganizationRoleId to be an explicit nil

### UnsetOrganizationRoleId
`func (o *UsersCreateUserRequest) UnsetOrganizationRoleId()`

UnsetOrganizationRoleId ensures that no value is present for OrganizationRoleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


