<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# Engine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EngineType**](EngineType.md) |  | 

## Methods

### NewEngine

`func NewEngine(type_ EngineType, ) *Engine`

NewEngine instantiates a new Engine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineWithDefaults

`func NewEngineWithDefaults() *Engine`

NewEngineWithDefaults instantiates a new Engine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Engine) GetType() EngineType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Engine) GetTypeOk() (*EngineType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Engine) SetType(v EngineType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


