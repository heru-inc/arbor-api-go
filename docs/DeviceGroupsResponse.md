# DeviceGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DeviceGroup**](DeviceGroup.md) |  | 
**Links** | [**Links**](Links.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewDeviceGroupsResponse

`func NewDeviceGroupsResponse(data []DeviceGroup, links Links, meta Meta, ) *DeviceGroupsResponse`

NewDeviceGroupsResponse instantiates a new DeviceGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceGroupsResponseWithDefaults

`func NewDeviceGroupsResponseWithDefaults() *DeviceGroupsResponse`

NewDeviceGroupsResponseWithDefaults instantiates a new DeviceGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DeviceGroupsResponse) GetData() []DeviceGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeviceGroupsResponse) GetDataOk() (*[]DeviceGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeviceGroupsResponse) SetData(v []DeviceGroup)`

SetData sets Data field to given value.


### GetLinks

`func (o *DeviceGroupsResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceGroupsResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceGroupsResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *DeviceGroupsResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DeviceGroupsResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DeviceGroupsResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


