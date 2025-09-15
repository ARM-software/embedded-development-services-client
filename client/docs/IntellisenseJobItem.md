<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# IntellisenseJobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableIntellisenseJobItemLinks**](IntellisenseJobItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**BuildStepsCompleted** | **NullableInt32** | The number of steps that have been completed so far. Please note: - This value also includes additional service orchestration steps, that are outside the core process,   so may differ from the job progress indicated within job messages. - This value will only be available after the job has been started. | [readonly] 
**BuildStepsTotal** | **NullableInt32** | The total number of steps that will need to be performed to complete the job. Please note: - This value also includes additional service orchestration steps, that are outside the core process,   so may differ from the job progress indicated within job messages. - This value will only be available after the job has been started. | [readonly] 
**Context** | Pointer to **NullableString** | Build context for jobs that require it. | [optional] 
**Done** | **bool** | True when the job has completed (this does not necessarily indicate success). | [readonly] 
**Error** | **bool** | True if there was an error in the service while attempting the job. | [readonly] 
**Failure** | **bool** | True if the job failed (this should be used in conjunction with the &#x60;done&#x60; property). | [readonly] 
**JobTimeout** | Pointer to **int32** | The maximum time (in seconds) that the job will be allowed to run. After the timeout has expired the job will be aborted and reported as a failure. The timeout does not include any time the request spent being queued, waiting for the job to be started. | [optional] [default to 90]
**Name** | **string** | Unique ID of the Intellisense Job. | [readonly] 
**Packs** | **string** | Path to packs repository to replace value in compilation database. | 
**Project** | **string** | CMSIS project to handle or being handled. | 
**Queued** | **bool** | True if job is currently queued and waiting to be processed. Otherwise, the job is either currently being processed or ended. | [readonly] 
**Status** | **string** | A summary status of the job. Note: this value should not be relied upon to determine whether a job has completed, succeeded or failed as this list may change as state machine evolves. Use resource appropriate flags instead. | [readonly] 
**Success** | **bool** | True if the job was successful (this should be used in conjunction with the &#x60;done&#x60; property). | [readonly] 
**Suspended** | **bool** | True if job has been cancelled or an order to halt it has been received. | [readonly] 
**Title** | Pointer to **NullableString** | Optional human readable name of the CMSIS Intellisense job. | [optional] 
**Toolchain** | **string** | Path to toolchain binaries to replace value in compilation database. | 
**ToolchainHeaders** | **string** | Path to toolchain headers to replace value in compilation database. | 
**Workspace** | **string** | Path to user&#39;s workspace to replace value in compilation database. | 

## Methods

### NewIntellisenseJobItem

`func NewIntellisenseJobItem(links NullableIntellisenseJobItemLinks, metadata NullableCommonMetadata, buildStepsCompleted NullableInt32, buildStepsTotal NullableInt32, done bool, error_ bool, failure bool, name string, packs string, project string, queued bool, status string, success bool, suspended bool, toolchain string, toolchainHeaders string, workspace string, ) *IntellisenseJobItem`

NewIntellisenseJobItem instantiates a new IntellisenseJobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntellisenseJobItemWithDefaults

`func NewIntellisenseJobItemWithDefaults() *IntellisenseJobItem`

NewIntellisenseJobItemWithDefaults instantiates a new IntellisenseJobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *IntellisenseJobItem) GetLinks() IntellisenseJobItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IntellisenseJobItem) GetLinksOk() (*IntellisenseJobItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IntellisenseJobItem) SetLinks(v IntellisenseJobItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *IntellisenseJobItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *IntellisenseJobItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *IntellisenseJobItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IntellisenseJobItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IntellisenseJobItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *IntellisenseJobItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *IntellisenseJobItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetBuildStepsCompleted

`func (o *IntellisenseJobItem) GetBuildStepsCompleted() int32`

GetBuildStepsCompleted returns the BuildStepsCompleted field if non-nil, zero value otherwise.

### GetBuildStepsCompletedOk

`func (o *IntellisenseJobItem) GetBuildStepsCompletedOk() (*int32, bool)`

GetBuildStepsCompletedOk returns a tuple with the BuildStepsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStepsCompleted

`func (o *IntellisenseJobItem) SetBuildStepsCompleted(v int32)`

SetBuildStepsCompleted sets BuildStepsCompleted field to given value.


### SetBuildStepsCompletedNil

`func (o *IntellisenseJobItem) SetBuildStepsCompletedNil(b bool)`

 SetBuildStepsCompletedNil sets the value for BuildStepsCompleted to be an explicit nil

### UnsetBuildStepsCompleted
`func (o *IntellisenseJobItem) UnsetBuildStepsCompleted()`

UnsetBuildStepsCompleted ensures that no value is present for BuildStepsCompleted, not even an explicit nil
### GetBuildStepsTotal

`func (o *IntellisenseJobItem) GetBuildStepsTotal() int32`

GetBuildStepsTotal returns the BuildStepsTotal field if non-nil, zero value otherwise.

### GetBuildStepsTotalOk

`func (o *IntellisenseJobItem) GetBuildStepsTotalOk() (*int32, bool)`

GetBuildStepsTotalOk returns a tuple with the BuildStepsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStepsTotal

`func (o *IntellisenseJobItem) SetBuildStepsTotal(v int32)`

SetBuildStepsTotal sets BuildStepsTotal field to given value.


### SetBuildStepsTotalNil

`func (o *IntellisenseJobItem) SetBuildStepsTotalNil(b bool)`

 SetBuildStepsTotalNil sets the value for BuildStepsTotal to be an explicit nil

### UnsetBuildStepsTotal
`func (o *IntellisenseJobItem) UnsetBuildStepsTotal()`

UnsetBuildStepsTotal ensures that no value is present for BuildStepsTotal, not even an explicit nil
### GetContext

`func (o *IntellisenseJobItem) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *IntellisenseJobItem) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *IntellisenseJobItem) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *IntellisenseJobItem) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *IntellisenseJobItem) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *IntellisenseJobItem) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetDone

`func (o *IntellisenseJobItem) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *IntellisenseJobItem) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *IntellisenseJobItem) SetDone(v bool)`

SetDone sets Done field to given value.


### GetError

`func (o *IntellisenseJobItem) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IntellisenseJobItem) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IntellisenseJobItem) SetError(v bool)`

SetError sets Error field to given value.


### GetFailure

`func (o *IntellisenseJobItem) GetFailure() bool`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *IntellisenseJobItem) GetFailureOk() (*bool, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *IntellisenseJobItem) SetFailure(v bool)`

SetFailure sets Failure field to given value.


### GetJobTimeout

`func (o *IntellisenseJobItem) GetJobTimeout() int32`

GetJobTimeout returns the JobTimeout field if non-nil, zero value otherwise.

### GetJobTimeoutOk

`func (o *IntellisenseJobItem) GetJobTimeoutOk() (*int32, bool)`

GetJobTimeoutOk returns a tuple with the JobTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTimeout

`func (o *IntellisenseJobItem) SetJobTimeout(v int32)`

SetJobTimeout sets JobTimeout field to given value.

### HasJobTimeout

`func (o *IntellisenseJobItem) HasJobTimeout() bool`

HasJobTimeout returns a boolean if a field has been set.

### GetName

`func (o *IntellisenseJobItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntellisenseJobItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntellisenseJobItem) SetName(v string)`

SetName sets Name field to given value.


### GetPacks

`func (o *IntellisenseJobItem) GetPacks() string`

GetPacks returns the Packs field if non-nil, zero value otherwise.

### GetPacksOk

`func (o *IntellisenseJobItem) GetPacksOk() (*string, bool)`

GetPacksOk returns a tuple with the Packs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacks

`func (o *IntellisenseJobItem) SetPacks(v string)`

SetPacks sets Packs field to given value.


### GetProject

`func (o *IntellisenseJobItem) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *IntellisenseJobItem) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *IntellisenseJobItem) SetProject(v string)`

SetProject sets Project field to given value.


### GetQueued

`func (o *IntellisenseJobItem) GetQueued() bool`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *IntellisenseJobItem) GetQueuedOk() (*bool, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *IntellisenseJobItem) SetQueued(v bool)`

SetQueued sets Queued field to given value.


### GetStatus

`func (o *IntellisenseJobItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntellisenseJobItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntellisenseJobItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSuccess

`func (o *IntellisenseJobItem) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *IntellisenseJobItem) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *IntellisenseJobItem) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetSuspended

`func (o *IntellisenseJobItem) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *IntellisenseJobItem) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *IntellisenseJobItem) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.


### GetTitle

`func (o *IntellisenseJobItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IntellisenseJobItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IntellisenseJobItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IntellisenseJobItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *IntellisenseJobItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *IntellisenseJobItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetToolchain

`func (o *IntellisenseJobItem) GetToolchain() string`

GetToolchain returns the Toolchain field if non-nil, zero value otherwise.

### GetToolchainOk

`func (o *IntellisenseJobItem) GetToolchainOk() (*string, bool)`

GetToolchainOk returns a tuple with the Toolchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolchain

`func (o *IntellisenseJobItem) SetToolchain(v string)`

SetToolchain sets Toolchain field to given value.


### GetToolchainHeaders

`func (o *IntellisenseJobItem) GetToolchainHeaders() string`

GetToolchainHeaders returns the ToolchainHeaders field if non-nil, zero value otherwise.

### GetToolchainHeadersOk

`func (o *IntellisenseJobItem) GetToolchainHeadersOk() (*string, bool)`

GetToolchainHeadersOk returns a tuple with the ToolchainHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolchainHeaders

`func (o *IntellisenseJobItem) SetToolchainHeaders(v string)`

SetToolchainHeaders sets ToolchainHeaders field to given value.


### GetWorkspace

`func (o *IntellisenseJobItem) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *IntellisenseJobItem) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *IntellisenseJobItem) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


