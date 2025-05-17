<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAWorkload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | **string** | Identifier of the payload in which the project is present. | 
**Project** | **string** | Path to user&#39;s project (e.g. binary, zip) to load and run onto the FPGA. | 
**Repository** | Pointer to **NullableString** | Identifier of the repository in which the payload is stored. | [optional] 

## Methods

### NewFPGAWorkload

`func NewFPGAWorkload(payload string, project string, ) *FPGAWorkload`

NewFPGAWorkload instantiates a new FPGAWorkload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAWorkloadWithDefaults

`func NewFPGAWorkloadWithDefaults() *FPGAWorkload`

NewFPGAWorkloadWithDefaults instantiates a new FPGAWorkload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *FPGAWorkload) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *FPGAWorkload) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *FPGAWorkload) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetProject

`func (o *FPGAWorkload) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FPGAWorkload) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FPGAWorkload) SetProject(v string)`

SetProject sets Project field to given value.


### GetRepository

`func (o *FPGAWorkload) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *FPGAWorkload) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *FPGAWorkload) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *FPGAWorkload) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *FPGAWorkload) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *FPGAWorkload) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


