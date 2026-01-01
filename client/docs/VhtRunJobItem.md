<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# VhtRunJobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableVhtRunJobItemLinks**](VhtRunJobItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Done** | **bool** | True when the job has completed. | [readonly] 
**Error** | **bool** | True if the job exited with an error code (this should be used in conjunction with the &#x60;done&#x60; property). | [readonly] 
**JobTimeout** | Pointer to **int32** | The maximum time (in seconds) that the job will be allowed to run. After the timeout has expired the job will be aborted and reported as a failure. The timeout does not include any time the request spent being queued, waiting for the job to be started. | [optional] [default to 300]
**Name** | **string** | Unique ID of the VHT run Job. | [readonly] 
**Queued** | **bool** | True if job is currently queued and waiting to be processed. Otherwise, the job is either currently being processed or ended. | [readonly] 
**Status** | **string** | A summary status of the VHT run job.  Note: this value should not be relied upon to determine whether a job has completed, started or errored as this list may change as state machine evolves.  Use resource appropriate flags instead. | [readonly] 
**SystemError** | **bool** | True if there was a system error in the service while running the job (this should be used in conjunction with the &#x60;done&#x60; property). | [readonly] 
**TimedOut** | **bool** | True if the job timed out (this should be used in conjunction with the &#x60;done&#x60; property). | [readonly] 
**Title** | Pointer to **NullableString** | Optional human readable name of the VHT run job. | [optional] 

## Methods

### NewVhtRunJobItem

`func NewVhtRunJobItem(links NullableVhtRunJobItemLinks, metadata NullableCommonMetadata, done bool, error_ bool, name string, queued bool, status string, systemError bool, timedOut bool, ) *VhtRunJobItem`

NewVhtRunJobItem instantiates a new VhtRunJobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVhtRunJobItemWithDefaults

`func NewVhtRunJobItemWithDefaults() *VhtRunJobItem`

NewVhtRunJobItemWithDefaults instantiates a new VhtRunJobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *VhtRunJobItem) GetLinks() VhtRunJobItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VhtRunJobItem) GetLinksOk() (*VhtRunJobItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VhtRunJobItem) SetLinks(v VhtRunJobItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *VhtRunJobItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *VhtRunJobItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *VhtRunJobItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VhtRunJobItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VhtRunJobItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *VhtRunJobItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *VhtRunJobItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDone

`func (o *VhtRunJobItem) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *VhtRunJobItem) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *VhtRunJobItem) SetDone(v bool)`

SetDone sets Done field to given value.


### GetError

`func (o *VhtRunJobItem) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VhtRunJobItem) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VhtRunJobItem) SetError(v bool)`

SetError sets Error field to given value.


### GetJobTimeout

`func (o *VhtRunJobItem) GetJobTimeout() int32`

GetJobTimeout returns the JobTimeout field if non-nil, zero value otherwise.

### GetJobTimeoutOk

`func (o *VhtRunJobItem) GetJobTimeoutOk() (*int32, bool)`

GetJobTimeoutOk returns a tuple with the JobTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTimeout

`func (o *VhtRunJobItem) SetJobTimeout(v int32)`

SetJobTimeout sets JobTimeout field to given value.

### HasJobTimeout

`func (o *VhtRunJobItem) HasJobTimeout() bool`

HasJobTimeout returns a boolean if a field has been set.

### GetName

`func (o *VhtRunJobItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VhtRunJobItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VhtRunJobItem) SetName(v string)`

SetName sets Name field to given value.


### GetQueued

`func (o *VhtRunJobItem) GetQueued() bool`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *VhtRunJobItem) GetQueuedOk() (*bool, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *VhtRunJobItem) SetQueued(v bool)`

SetQueued sets Queued field to given value.


### GetStatus

`func (o *VhtRunJobItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VhtRunJobItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VhtRunJobItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSystemError

`func (o *VhtRunJobItem) GetSystemError() bool`

GetSystemError returns the SystemError field if non-nil, zero value otherwise.

### GetSystemErrorOk

`func (o *VhtRunJobItem) GetSystemErrorOk() (*bool, bool)`

GetSystemErrorOk returns a tuple with the SystemError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemError

`func (o *VhtRunJobItem) SetSystemError(v bool)`

SetSystemError sets SystemError field to given value.


### GetTimedOut

`func (o *VhtRunJobItem) GetTimedOut() bool`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *VhtRunJobItem) GetTimedOutOk() (*bool, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *VhtRunJobItem) SetTimedOut(v bool)`

SetTimedOut sets TimedOut field to given value.


### GetTitle

`func (o *VhtRunJobItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VhtRunJobItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VhtRunJobItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VhtRunJobItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *VhtRunJobItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *VhtRunJobItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


