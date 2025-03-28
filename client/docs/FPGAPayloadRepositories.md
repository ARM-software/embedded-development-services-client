<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAPayloadRepositories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **[]string** | A list of repositories to create that will contain FPGA payloads | [optional] 
**Existing** | Pointer to **[]string** | A list of existing repositories that will contain FPGA payloads | [optional] 

## Methods

### NewFPGAPayloadRepositories

`func NewFPGAPayloadRepositories() *FPGAPayloadRepositories`

NewFPGAPayloadRepositories instantiates a new FPGAPayloadRepositories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAPayloadRepositoriesWithDefaults

`func NewFPGAPayloadRepositoriesWithDefaults() *FPGAPayloadRepositories`

NewFPGAPayloadRepositoriesWithDefaults instantiates a new FPGAPayloadRepositories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *FPGAPayloadRepositories) GetCreate() []string`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *FPGAPayloadRepositories) GetCreateOk() (*[]string, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *FPGAPayloadRepositories) SetCreate(v []string)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *FPGAPayloadRepositories) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetExisting

`func (o *FPGAPayloadRepositories) GetExisting() []string`

GetExisting returns the Existing field if non-nil, zero value otherwise.

### GetExistingOk

`func (o *FPGAPayloadRepositories) GetExistingOk() (*[]string, bool)`

GetExistingOk returns a tuple with the Existing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExisting

`func (o *FPGAPayloadRepositories) SetExisting(v []string)`

SetExisting sets Existing field to given value.

### HasExisting

`func (o *FPGAPayloadRepositories) HasExisting() bool`

HasExisting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


