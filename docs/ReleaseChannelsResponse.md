# ReleaseChannelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ReleaseChannel**](ReleaseChannel.md) |  | 
**Links** | [**Links**](Links.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewReleaseChannelsResponse

`func NewReleaseChannelsResponse(data []ReleaseChannel, links Links, meta Meta, ) *ReleaseChannelsResponse`

NewReleaseChannelsResponse instantiates a new ReleaseChannelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseChannelsResponseWithDefaults

`func NewReleaseChannelsResponseWithDefaults() *ReleaseChannelsResponse`

NewReleaseChannelsResponseWithDefaults instantiates a new ReleaseChannelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ReleaseChannelsResponse) GetData() []ReleaseChannel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReleaseChannelsResponse) GetDataOk() (*[]ReleaseChannel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReleaseChannelsResponse) SetData(v []ReleaseChannel)`

SetData sets Data field to given value.


### GetLinks

`func (o *ReleaseChannelsResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReleaseChannelsResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReleaseChannelsResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *ReleaseChannelsResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ReleaseChannelsResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ReleaseChannelsResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


