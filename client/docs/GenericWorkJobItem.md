<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# GenericWorkJobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableGenericWorkJobItemLinks**](GenericWorkJobItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Configuration** | Pointer to **map[string]string** | Configuration map for jobs that require it. These could be environment variables. This is job implementation dependent and job documentation should describe it. | [optional] 
**Done** | **bool** | True when the job has completed (this does necessarily indicate success). | [readonly] 
**Error** | **bool** | True if there was an error in the service while attempting the job. | [readonly] 
**Failure** | **bool** | True if the job failed (this should be used in conjunction with the &#x60;done&#x60; property). | [readonly] 
**JobStepsCompleted** | **NullableInt32** | The number of execution steps that have been completed so far. Please note: - This value also includes additional service orchestration steps, that are outside the core job being    executed by the process, so may differ from the build progress indicated within build messages.  - This value will only be available after the job has been started. | [readonly] 
**JobStepsTotal** | **NullableInt32** | The total number of steps that will need to be performed to complete the job. Please note: - This value also includes additional service orchestration steps, that are outside the core job being    executed by the process, so may differ from the build progress indicated within build messages.  - This value will only be available after the job has been started. | [readonly] 
**JobTimeout** | Pointer to **int32** | The maximum time (in seconds) that the job will be allowed to run. After the timeout has expired the job will be aborted and reported as a failure. The timeout does not include any time the request spent being queued, waiting for the job to be started. | [optional] [default to 300]
**Name** | **string** | Unique ID of the Generic Work Job. | [readonly] 
**Project** | Pointer to **string** | Path in the workspace to the project to handle or being handled. | [optional] 
**Queued** | Pointer to **bool** | True if job is currently queued and waiting to be processed. Otherwise, the job is either currently being processed or ended. | [optional] [readonly] 
**Status** | **string** | A summary status of the job. Note: this value should not be relied upon to determine whether a job has completed, succeeded or failed as this list may change as state machine evolves. Use resource appropriate flags instead. | [readonly] 
**Success** | **bool** | True if the job was successful (this should be used in conjunction with the &#x60;done&#x60; property). | [readonly] 
**Title** | Pointer to **NullableString** | Optional human readable name of the generic work job. | [optional] 
**Workspace** | Pointer to **NullableString** | Workspace name where the project is present. If not set, the default user&#39;s workspace will be used. | [optional] 

## Methods

### NewGenericWorkJobItem

`func NewGenericWorkJobItem(links NullableGenericWorkJobItemLinks, metadata NullableCommonMetadata, done bool, error_ bool, failure bool, jobStepsCompleted NullableInt32, jobStepsTotal NullableInt32, name string, status string, success bool, ) *GenericWorkJobItem`

NewGenericWorkJobItem instantiates a new GenericWorkJobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericWorkJobItemWithDefaults

`func NewGenericWorkJobItemWithDefaults() *GenericWorkJobItem`

NewGenericWorkJobItemWithDefaults instantiates a new GenericWorkJobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GenericWorkJobItem) GetLinks() GenericWorkJobItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GenericWorkJobItem) GetLinksOk() (*GenericWorkJobItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GenericWorkJobItem) SetLinks(v GenericWorkJobItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *GenericWorkJobItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *GenericWorkJobItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *GenericWorkJobItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GenericWorkJobItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GenericWorkJobItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *GenericWorkJobItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GenericWorkJobItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetConfiguration

`func (o *GenericWorkJobItem) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GenericWorkJobItem) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GenericWorkJobItem) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GenericWorkJobItem) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *GenericWorkJobItem) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *GenericWorkJobItem) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetDone

`func (o *GenericWorkJobItem) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *GenericWorkJobItem) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *GenericWorkJobItem) SetDone(v bool)`

SetDone sets Done field to given value.


### GetError

`func (o *GenericWorkJobItem) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GenericWorkJobItem) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GenericWorkJobItem) SetError(v bool)`

SetError sets Error field to given value.


### GetFailure

`func (o *GenericWorkJobItem) GetFailure() bool`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *GenericWorkJobItem) GetFailureOk() (*bool, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *GenericWorkJobItem) SetFailure(v bool)`

SetFailure sets Failure field to given value.


### GetJobStepsCompleted

`func (o *GenericWorkJobItem) GetJobStepsCompleted() int32`

GetJobStepsCompleted returns the JobStepsCompleted field if non-nil, zero value otherwise.

### GetJobStepsCompletedOk

`func (o *GenericWorkJobItem) GetJobStepsCompletedOk() (*int32, bool)`

GetJobStepsCompletedOk returns a tuple with the JobStepsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStepsCompleted

`func (o *GenericWorkJobItem) SetJobStepsCompleted(v int32)`

SetJobStepsCompleted sets JobStepsCompleted field to given value.


### SetJobStepsCompletedNil

`func (o *GenericWorkJobItem) SetJobStepsCompletedNil(b bool)`

 SetJobStepsCompletedNil sets the value for JobStepsCompleted to be an explicit nil

### UnsetJobStepsCompleted
`func (o *GenericWorkJobItem) UnsetJobStepsCompleted()`

UnsetJobStepsCompleted ensures that no value is present for JobStepsCompleted, not even an explicit nil
### GetJobStepsTotal

`func (o *GenericWorkJobItem) GetJobStepsTotal() int32`

GetJobStepsTotal returns the JobStepsTotal field if non-nil, zero value otherwise.

### GetJobStepsTotalOk

`func (o *GenericWorkJobItem) GetJobStepsTotalOk() (*int32, bool)`

GetJobStepsTotalOk returns a tuple with the JobStepsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStepsTotal

`func (o *GenericWorkJobItem) SetJobStepsTotal(v int32)`

SetJobStepsTotal sets JobStepsTotal field to given value.


### SetJobStepsTotalNil

`func (o *GenericWorkJobItem) SetJobStepsTotalNil(b bool)`

 SetJobStepsTotalNil sets the value for JobStepsTotal to be an explicit nil

### UnsetJobStepsTotal
`func (o *GenericWorkJobItem) UnsetJobStepsTotal()`

UnsetJobStepsTotal ensures that no value is present for JobStepsTotal, not even an explicit nil
### GetJobTimeout

`func (o *GenericWorkJobItem) GetJobTimeout() int32`

GetJobTimeout returns the JobTimeout field if non-nil, zero value otherwise.

### GetJobTimeoutOk

`func (o *GenericWorkJobItem) GetJobTimeoutOk() (*int32, bool)`

GetJobTimeoutOk returns a tuple with the JobTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTimeout

`func (o *GenericWorkJobItem) SetJobTimeout(v int32)`

SetJobTimeout sets JobTimeout field to given value.

### HasJobTimeout

`func (o *GenericWorkJobItem) HasJobTimeout() bool`

HasJobTimeout returns a boolean if a field has been set.

### GetName

`func (o *GenericWorkJobItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenericWorkJobItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenericWorkJobItem) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *GenericWorkJobItem) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GenericWorkJobItem) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GenericWorkJobItem) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GenericWorkJobItem) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetQueued

`func (o *GenericWorkJobItem) GetQueued() bool`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *GenericWorkJobItem) GetQueuedOk() (*bool, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *GenericWorkJobItem) SetQueued(v bool)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *GenericWorkJobItem) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetStatus

`func (o *GenericWorkJobItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GenericWorkJobItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GenericWorkJobItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSuccess

`func (o *GenericWorkJobItem) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GenericWorkJobItem) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GenericWorkJobItem) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetTitle

`func (o *GenericWorkJobItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GenericWorkJobItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GenericWorkJobItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GenericWorkJobItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *GenericWorkJobItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *GenericWorkJobItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetWorkspace

`func (o *GenericWorkJobItem) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *GenericWorkJobItem) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *GenericWorkJobItem) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *GenericWorkJobItem) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### SetWorkspaceNil

`func (o *GenericWorkJobItem) SetWorkspaceNil(b bool)`

 SetWorkspaceNil sets the value for Workspace to be an explicit nil

### UnsetWorkspace
`func (o *GenericWorkJobItem) UnsetWorkspace()`

UnsetWorkspace ensures that no value is present for Workspace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


