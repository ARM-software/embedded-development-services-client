<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAJobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableFPGAJobItemLinks**](FPGAJobItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Configuration** | Pointer to **map[string]string** | Configuration map for jobs that require it. These could be environment variables. This is job implementation dependent and job documentation should describe it. | [optional] 
**Connected** | **bool** | True when there is an active connection to the application running on the FPGA. If the job does not support connection, this flag will never be true. | [readonly] 
**Done** | **bool** | True when the job has completed (this does not necessarily indicate success). | [readonly] 
**Error** | **bool** | True if there was an error in the service while attempting the job. | [readonly] 
**Failure** | **bool** | True if the job failed (this should be used in conjunction with the &#x60;done&#x60; property). | [readonly] 
**Name** | **string** | Unique ID of the FPGA job. | [readonly] 
**Queued** | **bool** | True if job is currently queued and waiting to be processed. Otherwise, the job is either currently being processed or ended. | [readonly] 
**ReadyForConnection** | **bool** | True when the application running on the FPGA is ready to handle connections. If the job does not support connection, this flag will never be true. | [readonly] 
**Status** | **string** | A summary status of the job. Note: this value should not be relied upon to determine whether a job has completed, succeeded or failed as this list may change as state machine evolves. Use resource appropriate flags instead. | [readonly] 
**StepsCompleted** | **NullableInt32** | The number of steps that have been completed so far. Please note: - This value also includes additional service orchestration steps, that are outside the core process,   so may differ from the job progress indicated within job messages. - This value will only be available after the job has been started. | [readonly] 
**StepsTotal** | **NullableInt32** | The total number of steps that will need to be performed to complete the job. Please note: - This value also includes additional service orchestration steps, that are outside the core process,   so may differ from the job progress indicated within job messages. - This value will only be available after the job has been started. | [readonly] 
**Success** | **bool** | True if the job was successful (this should be used in conjunction with the &#x60;done&#x60; property). | [readonly] 
**SupportConnection** | **bool** | True when the job allows direct connection to the job instance (application running on the FPGA). | [readonly] 
**Target** | [**FPGATargetID**](FPGATargetID.md) |  | 
**Timeout** | Pointer to **int64** | The maximum time (in seconds) that the job will be allowed to run. After the timeout has expired the job will be aborted and reported as a failure. The timeout does not include any time the request spent being queued, waiting for the job to be started. | [optional] [default to 300]
**Title** | Pointer to **NullableString** | Optional human-readable name of the FPGA job. | [optional] 
**Workload** | [**FPGAWorkload**](FPGAWorkload.md) |  | 

## Methods

### NewFPGAJobItem

`func NewFPGAJobItem(links NullableFPGAJobItemLinks, metadata NullableCommonMetadata, connected bool, done bool, error_ bool, failure bool, name string, queued bool, readyForConnection bool, status string, stepsCompleted NullableInt32, stepsTotal NullableInt32, success bool, supportConnection bool, target FPGATargetID, workload FPGAWorkload, ) *FPGAJobItem`

NewFPGAJobItem instantiates a new FPGAJobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAJobItemWithDefaults

`func NewFPGAJobItemWithDefaults() *FPGAJobItem`

NewFPGAJobItemWithDefaults instantiates a new FPGAJobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FPGAJobItem) GetLinks() FPGAJobItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FPGAJobItem) GetLinksOk() (*FPGAJobItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FPGAJobItem) SetLinks(v FPGAJobItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *FPGAJobItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *FPGAJobItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *FPGAJobItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FPGAJobItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FPGAJobItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *FPGAJobItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *FPGAJobItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetConfiguration

`func (o *FPGAJobItem) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *FPGAJobItem) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *FPGAJobItem) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *FPGAJobItem) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *FPGAJobItem) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *FPGAJobItem) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetConnected

`func (o *FPGAJobItem) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *FPGAJobItem) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *FPGAJobItem) SetConnected(v bool)`

SetConnected sets Connected field to given value.


### GetDone

`func (o *FPGAJobItem) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *FPGAJobItem) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *FPGAJobItem) SetDone(v bool)`

SetDone sets Done field to given value.


### GetError

`func (o *FPGAJobItem) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FPGAJobItem) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FPGAJobItem) SetError(v bool)`

SetError sets Error field to given value.


### GetFailure

`func (o *FPGAJobItem) GetFailure() bool`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *FPGAJobItem) GetFailureOk() (*bool, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *FPGAJobItem) SetFailure(v bool)`

SetFailure sets Failure field to given value.


### GetName

`func (o *FPGAJobItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FPGAJobItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FPGAJobItem) SetName(v string)`

SetName sets Name field to given value.


### GetQueued

`func (o *FPGAJobItem) GetQueued() bool`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *FPGAJobItem) GetQueuedOk() (*bool, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *FPGAJobItem) SetQueued(v bool)`

SetQueued sets Queued field to given value.


### GetReadyForConnection

`func (o *FPGAJobItem) GetReadyForConnection() bool`

GetReadyForConnection returns the ReadyForConnection field if non-nil, zero value otherwise.

### GetReadyForConnectionOk

`func (o *FPGAJobItem) GetReadyForConnectionOk() (*bool, bool)`

GetReadyForConnectionOk returns a tuple with the ReadyForConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyForConnection

`func (o *FPGAJobItem) SetReadyForConnection(v bool)`

SetReadyForConnection sets ReadyForConnection field to given value.


### GetStatus

`func (o *FPGAJobItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FPGAJobItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FPGAJobItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStepsCompleted

`func (o *FPGAJobItem) GetStepsCompleted() int32`

GetStepsCompleted returns the StepsCompleted field if non-nil, zero value otherwise.

### GetStepsCompletedOk

`func (o *FPGAJobItem) GetStepsCompletedOk() (*int32, bool)`

GetStepsCompletedOk returns a tuple with the StepsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsCompleted

`func (o *FPGAJobItem) SetStepsCompleted(v int32)`

SetStepsCompleted sets StepsCompleted field to given value.


### SetStepsCompletedNil

`func (o *FPGAJobItem) SetStepsCompletedNil(b bool)`

 SetStepsCompletedNil sets the value for StepsCompleted to be an explicit nil

### UnsetStepsCompleted
`func (o *FPGAJobItem) UnsetStepsCompleted()`

UnsetStepsCompleted ensures that no value is present for StepsCompleted, not even an explicit nil
### GetStepsTotal

`func (o *FPGAJobItem) GetStepsTotal() int32`

GetStepsTotal returns the StepsTotal field if non-nil, zero value otherwise.

### GetStepsTotalOk

`func (o *FPGAJobItem) GetStepsTotalOk() (*int32, bool)`

GetStepsTotalOk returns a tuple with the StepsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsTotal

`func (o *FPGAJobItem) SetStepsTotal(v int32)`

SetStepsTotal sets StepsTotal field to given value.


### SetStepsTotalNil

`func (o *FPGAJobItem) SetStepsTotalNil(b bool)`

 SetStepsTotalNil sets the value for StepsTotal to be an explicit nil

### UnsetStepsTotal
`func (o *FPGAJobItem) UnsetStepsTotal()`

UnsetStepsTotal ensures that no value is present for StepsTotal, not even an explicit nil
### GetSuccess

`func (o *FPGAJobItem) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FPGAJobItem) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FPGAJobItem) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetSupportConnection

`func (o *FPGAJobItem) GetSupportConnection() bool`

GetSupportConnection returns the SupportConnection field if non-nil, zero value otherwise.

### GetSupportConnectionOk

`func (o *FPGAJobItem) GetSupportConnectionOk() (*bool, bool)`

GetSupportConnectionOk returns a tuple with the SupportConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportConnection

`func (o *FPGAJobItem) SetSupportConnection(v bool)`

SetSupportConnection sets SupportConnection field to given value.


### GetTarget

`func (o *FPGAJobItem) GetTarget() FPGATargetID`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *FPGAJobItem) GetTargetOk() (*FPGATargetID, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *FPGAJobItem) SetTarget(v FPGATargetID)`

SetTarget sets Target field to given value.


### GetTimeout

`func (o *FPGAJobItem) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *FPGAJobItem) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *FPGAJobItem) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *FPGAJobItem) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTitle

`func (o *FPGAJobItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FPGAJobItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FPGAJobItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FPGAJobItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FPGAJobItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FPGAJobItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetWorkload

`func (o *FPGAJobItem) GetWorkload() FPGAWorkload`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *FPGAJobItem) GetWorkloadOk() (*FPGAWorkload, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *FPGAJobItem) SetWorkload(v FPGAWorkload)`

SetWorkload sets Workload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


