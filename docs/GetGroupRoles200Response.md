# GetGroupRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetUsers200ResponseDataInnerOrganizationRole**](GetUsers200ResponseDataInnerOrganizationRole.md) |  | [optional] 
**Links** | Pointer to [**GetGroupRoles200ResponseLinks**](GetGroupRoles200ResponseLinks.md) |  | [optional] 
**Meta** | Pointer to [**GetGroupRoles200ResponseMeta**](GetGroupRoles200ResponseMeta.md) |  | [optional] 

## Methods

### NewGetGroupRoles200Response

`func NewGetGroupRoles200Response() *GetGroupRoles200Response`

NewGetGroupRoles200Response instantiates a new GetGroupRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroupRoles200ResponseWithDefaults

`func NewGetGroupRoles200ResponseWithDefaults() *GetGroupRoles200Response`

NewGetGroupRoles200ResponseWithDefaults instantiates a new GetGroupRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetGroupRoles200Response) GetData() []GetUsers200ResponseDataInnerOrganizationRole`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetGroupRoles200Response) GetDataOk() (*[]GetUsers200ResponseDataInnerOrganizationRole, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetGroupRoles200Response) SetData(v []GetUsers200ResponseDataInnerOrganizationRole)`

SetData sets Data field to given value.

### HasData

`func (o *GetGroupRoles200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetGroupRoles200Response) GetLinks() GetGroupRoles200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetGroupRoles200Response) GetLinksOk() (*GetGroupRoles200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetGroupRoles200Response) SetLinks(v GetGroupRoles200ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetGroupRoles200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *GetGroupRoles200Response) GetMeta() GetGroupRoles200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetGroupRoles200Response) GetMetaOk() (*GetGroupRoles200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetGroupRoles200Response) SetMeta(v GetGroupRoles200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetGroupRoles200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


