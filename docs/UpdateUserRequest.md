# UpdateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**OrganizationRoleId** | Pointer to **NullableString** |  | [optional] 
**GroupRoleId** | Pointer to **NullableString** | Required when &#x60;groupIds&#x60; is provided. | [optional] 
**GroupIds** | Pointer to **[]string** | Required when &#x60;groupRoleId&#x60; is provided. | [optional] 

## Methods

### NewUpdateUserRequest

`func NewUpdateUserRequest() *UpdateUserRequest`

NewUpdateUserRequest instantiates a new UpdateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestWithDefaults

`func NewUpdateUserRequestWithDefaults() *UpdateUserRequest`

NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UpdateUserRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateUserRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateUserRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateUserRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UpdateUserRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateUserRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateUserRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateUserRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetOrganizationRoleId

`func (o *UpdateUserRequest) GetOrganizationRoleId() string`

GetOrganizationRoleId returns the OrganizationRoleId field if non-nil, zero value otherwise.

### GetOrganizationRoleIdOk

`func (o *UpdateUserRequest) GetOrganizationRoleIdOk() (*string, bool)`

GetOrganizationRoleIdOk returns a tuple with the OrganizationRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoleId

`func (o *UpdateUserRequest) SetOrganizationRoleId(v string)`

SetOrganizationRoleId sets OrganizationRoleId field to given value.

### HasOrganizationRoleId

`func (o *UpdateUserRequest) HasOrganizationRoleId() bool`

HasOrganizationRoleId returns a boolean if a field has been set.

### SetOrganizationRoleIdNil

`func (o *UpdateUserRequest) SetOrganizationRoleIdNil(b bool)`

 SetOrganizationRoleIdNil sets the value for OrganizationRoleId to be an explicit nil

### UnsetOrganizationRoleId
`func (o *UpdateUserRequest) UnsetOrganizationRoleId()`

UnsetOrganizationRoleId ensures that no value is present for OrganizationRoleId, not even an explicit nil
### GetGroupRoleId

`func (o *UpdateUserRequest) GetGroupRoleId() string`

GetGroupRoleId returns the GroupRoleId field if non-nil, zero value otherwise.

### GetGroupRoleIdOk

`func (o *UpdateUserRequest) GetGroupRoleIdOk() (*string, bool)`

GetGroupRoleIdOk returns a tuple with the GroupRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRoleId

`func (o *UpdateUserRequest) SetGroupRoleId(v string)`

SetGroupRoleId sets GroupRoleId field to given value.

### HasGroupRoleId

`func (o *UpdateUserRequest) HasGroupRoleId() bool`

HasGroupRoleId returns a boolean if a field has been set.

### SetGroupRoleIdNil

`func (o *UpdateUserRequest) SetGroupRoleIdNil(b bool)`

 SetGroupRoleIdNil sets the value for GroupRoleId to be an explicit nil

### UnsetGroupRoleId
`func (o *UpdateUserRequest) UnsetGroupRoleId()`

UnsetGroupRoleId ensures that no value is present for GroupRoleId, not even an explicit nil
### GetGroupIds

`func (o *UpdateUserRequest) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *UpdateUserRequest) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *UpdateUserRequest) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *UpdateUserRequest) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


