<!--
Copyright (C) 2020-2022 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# ADebugInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adapter** | Pointer to **string** | Adapter for the interface | [optional] 
**Connector** | Pointer to **string** | Connector for the interface | [optional] 

## Methods

### NewADebugInterface

`func NewADebugInterface() *ADebugInterface`

NewADebugInterface instantiates a new ADebugInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADebugInterfaceWithDefaults

`func NewADebugInterfaceWithDefaults() *ADebugInterface`

NewADebugInterfaceWithDefaults instantiates a new ADebugInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdapter

`func (o *ADebugInterface) GetAdapter() string`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *ADebugInterface) GetAdapterOk() (*string, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *ADebugInterface) SetAdapter(v string)`

SetAdapter sets Adapter field to given value.

### HasAdapter

`func (o *ADebugInterface) HasAdapter() bool`

HasAdapter returns a boolean if a field has been set.

### GetConnector

`func (o *ADebugInterface) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ADebugInterface) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ADebugInterface) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *ADebugInterface) HasConnector() bool`

HasConnector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


