# GetUsers200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**OrganizationRole** | Pointer to [**GetUsers200ResponseDataInnerOrganizationRole**](GetUsers200ResponseDataInnerOrganizationRole.md) |  | [optional] 
**GroupRoles** | Pointer to [**[]GetUsers200ResponseDataInnerGroupRolesInner**](GetUsers200ResponseDataInnerGroupRolesInner.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUsers200ResponseDataInner

`func NewGetUsers200ResponseDataInner() *GetUsers200ResponseDataInner`

NewGetUsers200ResponseDataInner instantiates a new GetUsers200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsers200ResponseDataInnerWithDefaults

`func NewGetUsers200ResponseDataInnerWithDefaults() *GetUsers200ResponseDataInner`

NewGetUsers200ResponseDataInnerWithDefaults instantiates a new GetUsers200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetUsers200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUsers200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUsers200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetUsers200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *GetUsers200ResponseDataInner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *GetUsers200ResponseDataInner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *GetUsers200ResponseDataInner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *GetUsers200ResponseDataInner) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *GetUsers200ResponseDataInner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *GetUsers200ResponseDataInner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *GetUsers200ResponseDataInner) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *GetUsers200ResponseDataInner) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *GetUsers200ResponseDataInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetUsers200ResponseDataInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetUsers200ResponseDataInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetUsers200ResponseDataInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOrganizationRole

`func (o *GetUsers200ResponseDataInner) GetOrganizationRole() GetUsers200ResponseDataInnerOrganizationRole`

GetOrganizationRole returns the OrganizationRole field if non-nil, zero value otherwise.

### GetOrganizationRoleOk

`func (o *GetUsers200ResponseDataInner) GetOrganizationRoleOk() (*GetUsers200ResponseDataInnerOrganizationRole, bool)`

GetOrganizationRoleOk returns a tuple with the OrganizationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRole

`func (o *GetUsers200ResponseDataInner) SetOrganizationRole(v GetUsers200ResponseDataInnerOrganizationRole)`

SetOrganizationRole sets OrganizationRole field to given value.

### HasOrganizationRole

`func (o *GetUsers200ResponseDataInner) HasOrganizationRole() bool`

HasOrganizationRole returns a boolean if a field has been set.

### GetGroupRoles

`func (o *GetUsers200ResponseDataInner) GetGroupRoles() []GetUsers200ResponseDataInnerGroupRolesInner`

GetGroupRoles returns the GroupRoles field if non-nil, zero value otherwise.

### GetGroupRolesOk

`func (o *GetUsers200ResponseDataInner) GetGroupRolesOk() (*[]GetUsers200ResponseDataInnerGroupRolesInner, bool)`

GetGroupRolesOk returns a tuple with the GroupRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRoles

`func (o *GetUsers200ResponseDataInner) SetGroupRoles(v []GetUsers200ResponseDataInnerGroupRolesInner)`

SetGroupRoles sets GroupRoles field to given value.

### HasGroupRoles

`func (o *GetUsers200ResponseDataInner) HasGroupRoles() bool`

HasGroupRoles returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetUsers200ResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetUsers200ResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetUsers200ResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetUsers200ResponseDataInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetUsers200ResponseDataInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetUsers200ResponseDataInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetUsers200ResponseDataInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetUsers200ResponseDataInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


