# GetOrganizationRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetUsers200ResponseDataInnerOrganizationRole**](GetUsers200ResponseDataInnerOrganizationRole.md) |  | [optional] 
**Links** | Pointer to [**GetOrganizationRoles200ResponseLinks**](GetOrganizationRoles200ResponseLinks.md) |  | [optional] 
**Meta** | Pointer to [**GetOrganizationRoles200ResponseMeta**](GetOrganizationRoles200ResponseMeta.md) |  | [optional] 

## Methods

### NewGetOrganizationRoles200Response

`func NewGetOrganizationRoles200Response() *GetOrganizationRoles200Response`

NewGetOrganizationRoles200Response instantiates a new GetOrganizationRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationRoles200ResponseWithDefaults

`func NewGetOrganizationRoles200ResponseWithDefaults() *GetOrganizationRoles200Response`

NewGetOrganizationRoles200ResponseWithDefaults instantiates a new GetOrganizationRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetOrganizationRoles200Response) GetData() []GetUsers200ResponseDataInnerOrganizationRole`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOrganizationRoles200Response) GetDataOk() (*[]GetUsers200ResponseDataInnerOrganizationRole, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOrganizationRoles200Response) SetData(v []GetUsers200ResponseDataInnerOrganizationRole)`

SetData sets Data field to given value.

### HasData

`func (o *GetOrganizationRoles200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GetOrganizationRoles200Response) GetLinks() GetOrganizationRoles200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetOrganizationRoles200Response) GetLinksOk() (*GetOrganizationRoles200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetOrganizationRoles200Response) SetLinks(v GetOrganizationRoles200ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetOrganizationRoles200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *GetOrganizationRoles200Response) GetMeta() GetOrganizationRoles200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetOrganizationRoles200Response) GetMetaOk() (*GetOrganizationRoles200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetOrganizationRoles200Response) SetMeta(v GetOrganizationRoles200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetOrganizationRoles200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


