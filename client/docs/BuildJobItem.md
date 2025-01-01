<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# BuildJobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableBuildJobItemLinks**](BuildJobItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**BuildStepsCompleted** | **NullableInt32** | The number of build steps that have been completed so far. Please note: - This value also includes additional service orchestration steps, that are outside the core build process,   so may differ from the build progress indicated within build messages. - This value will only be available after the build has been started. | [readonly] 
**BuildStepsTotal** | **NullableInt32** | The total number of steps that will need to be performed to complete the build. Please note: - This value also includes additional service orchestration steps, that are outside the core build process,   so may differ from the build progress indicated within build messages. - This value will only be available after the build has been started. | [readonly] 
**BuildTimeout** | Pointer to **int32** | The maximum time (in seconds) that the build will be allowed to run. After the timeout has expired the build will be aborted and reported as a failure. The timeout does not include any time the request spent being queued, waiting for the build to be started. | [optional] [default to 300]
**CleanBuild** | Pointer to **bool** | Whether to run a clean build. | [optional] [default to false]
**Context** | Pointer to **NullableString** | Build context for jobs that require it. | [optional] 
**Done** | **bool** | True when the build job has completed (this does necessarily indicate success). | [readonly] 
**Error** | **bool** | True if there was an error in the build service while attempting the build. | [readonly] 
**Failure** | **bool** | True if the compilation or link failed (this should be used in conjunction with the &#x60;done&#x60; property). | [readonly] 
**Name** | **string** | Unique ID of the Build Job. | [readonly] 
**Project** | **string** | CMSIS project to build or being built. | 
**Queued** | Pointer to **bool** | True if job is currently queued and waiting to be processed. Otherwise, the job is either currently being processed or ended. | [optional] [readonly] 
**Status** | **string** | A summary status of the job. Note: this value should not be relied upon to determine whether a job has completed, succeeded or failed as this list may change as state machine evolves. Use resource appropriate flags instead. | [readonly] 
**Success** | **bool** | True if the build job was successful (this should be used in conjunction with the &#x60;done&#x60; property). | [readonly] 
**Title** | Pointer to **NullableString** | Optional human readable name of the CMSIS build job. | [optional] 
**Workspace** | Pointer to **NullableString** | Workspace name where the CMSIS project is present. If not set, the default user&#39;s workspace will be used. | [optional] 

## Methods

### NewBuildJobItem

`func NewBuildJobItem(links NullableBuildJobItemLinks, metadata NullableCommonMetadata, buildStepsCompleted NullableInt32, buildStepsTotal NullableInt32, done bool, error_ bool, failure bool, name string, project string, status string, success bool, ) *BuildJobItem`

NewBuildJobItem instantiates a new BuildJobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildJobItemWithDefaults

`func NewBuildJobItemWithDefaults() *BuildJobItem`

NewBuildJobItemWithDefaults instantiates a new BuildJobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *BuildJobItem) GetLinks() BuildJobItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BuildJobItem) GetLinksOk() (*BuildJobItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BuildJobItem) SetLinks(v BuildJobItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *BuildJobItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *BuildJobItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *BuildJobItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BuildJobItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BuildJobItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *BuildJobItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BuildJobItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetBuildStepsCompleted

`func (o *BuildJobItem) GetBuildStepsCompleted() int32`

GetBuildStepsCompleted returns the BuildStepsCompleted field if non-nil, zero value otherwise.

### GetBuildStepsCompletedOk

`func (o *BuildJobItem) GetBuildStepsCompletedOk() (*int32, bool)`

GetBuildStepsCompletedOk returns a tuple with the BuildStepsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStepsCompleted

`func (o *BuildJobItem) SetBuildStepsCompleted(v int32)`

SetBuildStepsCompleted sets BuildStepsCompleted field to given value.


### SetBuildStepsCompletedNil

`func (o *BuildJobItem) SetBuildStepsCompletedNil(b bool)`

 SetBuildStepsCompletedNil sets the value for BuildStepsCompleted to be an explicit nil

### UnsetBuildStepsCompleted
`func (o *BuildJobItem) UnsetBuildStepsCompleted()`

UnsetBuildStepsCompleted ensures that no value is present for BuildStepsCompleted, not even an explicit nil
### GetBuildStepsTotal

`func (o *BuildJobItem) GetBuildStepsTotal() int32`

GetBuildStepsTotal returns the BuildStepsTotal field if non-nil, zero value otherwise.

### GetBuildStepsTotalOk

`func (o *BuildJobItem) GetBuildStepsTotalOk() (*int32, bool)`

GetBuildStepsTotalOk returns a tuple with the BuildStepsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStepsTotal

`func (o *BuildJobItem) SetBuildStepsTotal(v int32)`

SetBuildStepsTotal sets BuildStepsTotal field to given value.


### SetBuildStepsTotalNil

`func (o *BuildJobItem) SetBuildStepsTotalNil(b bool)`

 SetBuildStepsTotalNil sets the value for BuildStepsTotal to be an explicit nil

### UnsetBuildStepsTotal
`func (o *BuildJobItem) UnsetBuildStepsTotal()`

UnsetBuildStepsTotal ensures that no value is present for BuildStepsTotal, not even an explicit nil
### GetBuildTimeout

`func (o *BuildJobItem) GetBuildTimeout() int32`

GetBuildTimeout returns the BuildTimeout field if non-nil, zero value otherwise.

### GetBuildTimeoutOk

`func (o *BuildJobItem) GetBuildTimeoutOk() (*int32, bool)`

GetBuildTimeoutOk returns a tuple with the BuildTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTimeout

`func (o *BuildJobItem) SetBuildTimeout(v int32)`

SetBuildTimeout sets BuildTimeout field to given value.

### HasBuildTimeout

`func (o *BuildJobItem) HasBuildTimeout() bool`

HasBuildTimeout returns a boolean if a field has been set.

### GetCleanBuild

`func (o *BuildJobItem) GetCleanBuild() bool`

GetCleanBuild returns the CleanBuild field if non-nil, zero value otherwise.

### GetCleanBuildOk

`func (o *BuildJobItem) GetCleanBuildOk() (*bool, bool)`

GetCleanBuildOk returns a tuple with the CleanBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanBuild

`func (o *BuildJobItem) SetCleanBuild(v bool)`

SetCleanBuild sets CleanBuild field to given value.

### HasCleanBuild

`func (o *BuildJobItem) HasCleanBuild() bool`

HasCleanBuild returns a boolean if a field has been set.

### GetContext

`func (o *BuildJobItem) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *BuildJobItem) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *BuildJobItem) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *BuildJobItem) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *BuildJobItem) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *BuildJobItem) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetDone

`func (o *BuildJobItem) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *BuildJobItem) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *BuildJobItem) SetDone(v bool)`

SetDone sets Done field to given value.


### GetError

`func (o *BuildJobItem) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BuildJobItem) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BuildJobItem) SetError(v bool)`

SetError sets Error field to given value.


### GetFailure

`func (o *BuildJobItem) GetFailure() bool`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *BuildJobItem) GetFailureOk() (*bool, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *BuildJobItem) SetFailure(v bool)`

SetFailure sets Failure field to given value.


### GetName

`func (o *BuildJobItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuildJobItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuildJobItem) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *BuildJobItem) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BuildJobItem) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BuildJobItem) SetProject(v string)`

SetProject sets Project field to given value.


### GetQueued

`func (o *BuildJobItem) GetQueued() bool`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *BuildJobItem) GetQueuedOk() (*bool, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *BuildJobItem) SetQueued(v bool)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *BuildJobItem) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetStatus

`func (o *BuildJobItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BuildJobItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BuildJobItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSuccess

`func (o *BuildJobItem) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BuildJobItem) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BuildJobItem) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetTitle

`func (o *BuildJobItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BuildJobItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BuildJobItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BuildJobItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BuildJobItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BuildJobItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetWorkspace

`func (o *BuildJobItem) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *BuildJobItem) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *BuildJobItem) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *BuildJobItem) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### SetWorkspaceNil

`func (o *BuildJobItem) SetWorkspaceNil(b bool)`

 SetWorkspaceNil sets the value for Workspace to be an explicit nil

### UnsetWorkspace
`func (o *BuildJobItem) UnsetWorkspace()`

UnsetWorkspace ensures that no value is present for Workspace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


