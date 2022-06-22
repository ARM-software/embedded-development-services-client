# VhtInstanceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableVhtInstanceItemLinks**](VhtInstanceItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Error** | **bool** | True if an error occurred with the instance (e.g. network outage, etc.). | [readonly] 
**InstanceTimeout** | Pointer to **int32** | The maximum time (in minutes) that the instance will be available for. After the timeout has expired the instance will be stopped and released. The timeout does not include any time the request spent being queued, waiting for the instance to be started. | [optional] [default to 60]
**Name** | **string** | Unique ID of the VHT instance. | [readonly] 
**Ready** | **bool** | True if an instance has all the artefacts necessary to carry out run jobs. | [readonly] 
**Requested** | **bool** | True if an instance has been requested. | [readonly] 
**Status** | **string** | A summary status of the VHT instance.  Note: this value should not be relied upon to determine whether a VHT is ready as this list may change as state machine evolves.  Use resource appropriate flags instead. | [readonly] 
**Terminated** | **bool** | True when the instance has been terminated (this necessarily indicates that the instance is no longer available). | [readonly] 
**Title** | Pointer to **NullableString** | Optional human readable name of the VHT instance. | [optional] 

## Methods

### NewVhtInstanceItem

`func NewVhtInstanceItem(links NullableVhtInstanceItemLinks, metadata NullableCommonMetadata, error_ bool, name string, ready bool, requested bool, status string, terminated bool, ) *VhtInstanceItem`

NewVhtInstanceItem instantiates a new VhtInstanceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVhtInstanceItemWithDefaults

`func NewVhtInstanceItemWithDefaults() *VhtInstanceItem`

NewVhtInstanceItemWithDefaults instantiates a new VhtInstanceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *VhtInstanceItem) GetLinks() VhtInstanceItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VhtInstanceItem) GetLinksOk() (*VhtInstanceItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VhtInstanceItem) SetLinks(v VhtInstanceItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *VhtInstanceItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *VhtInstanceItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *VhtInstanceItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VhtInstanceItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VhtInstanceItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *VhtInstanceItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *VhtInstanceItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetError

`func (o *VhtInstanceItem) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VhtInstanceItem) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VhtInstanceItem) SetError(v bool)`

SetError sets Error field to given value.


### GetInstanceTimeout

`func (o *VhtInstanceItem) GetInstanceTimeout() int32`

GetInstanceTimeout returns the InstanceTimeout field if non-nil, zero value otherwise.

### GetInstanceTimeoutOk

`func (o *VhtInstanceItem) GetInstanceTimeoutOk() (*int32, bool)`

GetInstanceTimeoutOk returns a tuple with the InstanceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTimeout

`func (o *VhtInstanceItem) SetInstanceTimeout(v int32)`

SetInstanceTimeout sets InstanceTimeout field to given value.

### HasInstanceTimeout

`func (o *VhtInstanceItem) HasInstanceTimeout() bool`

HasInstanceTimeout returns a boolean if a field has been set.

### GetName

`func (o *VhtInstanceItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VhtInstanceItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VhtInstanceItem) SetName(v string)`

SetName sets Name field to given value.


### GetReady

`func (o *VhtInstanceItem) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *VhtInstanceItem) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *VhtInstanceItem) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetRequested

`func (o *VhtInstanceItem) GetRequested() bool`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *VhtInstanceItem) GetRequestedOk() (*bool, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *VhtInstanceItem) SetRequested(v bool)`

SetRequested sets Requested field to given value.


### GetStatus

`func (o *VhtInstanceItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VhtInstanceItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VhtInstanceItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTerminated

`func (o *VhtInstanceItem) GetTerminated() bool`

GetTerminated returns the Terminated field if non-nil, zero value otherwise.

### GetTerminatedOk

`func (o *VhtInstanceItem) GetTerminatedOk() (*bool, bool)`

GetTerminatedOk returns a tuple with the Terminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminated

`func (o *VhtInstanceItem) SetTerminated(v bool)`

SetTerminated sets Terminated field to given value.


### GetTitle

`func (o *VhtInstanceItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VhtInstanceItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VhtInstanceItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VhtInstanceItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *VhtInstanceItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *VhtInstanceItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


