<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# InstancePermissionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]PermissionOperation**](PermissionOperation.md) | The list of allowed CRUDL operations. | 
**ResourceInstance** | **string** | The unique identifier of the resource instance. | 

## Methods

### NewInstancePermissionItem

`func NewInstancePermissionItem(operations []PermissionOperation, resourceInstance string, ) *InstancePermissionItem`

NewInstancePermissionItem instantiates a new InstancePermissionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancePermissionItemWithDefaults

`func NewInstancePermissionItemWithDefaults() *InstancePermissionItem`

NewInstancePermissionItemWithDefaults instantiates a new InstancePermissionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *InstancePermissionItem) GetOperations() []PermissionOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *InstancePermissionItem) GetOperationsOk() (*[]PermissionOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *InstancePermissionItem) SetOperations(v []PermissionOperation)`

SetOperations sets Operations field to given value.


### SetOperationsNil

`func (o *InstancePermissionItem) SetOperationsNil(b bool)`

 SetOperationsNil sets the value for Operations to be an explicit nil

### UnsetOperations
`func (o *InstancePermissionItem) UnsetOperations()`

UnsetOperations ensures that no value is present for Operations, not even an explicit nil
### GetResourceInstance

`func (o *InstancePermissionItem) GetResourceInstance() string`

GetResourceInstance returns the ResourceInstance field if non-nil, zero value otherwise.

### GetResourceInstanceOk

`func (o *InstancePermissionItem) GetResourceInstanceOk() (*string, bool)`

GetResourceInstanceOk returns a tuple with the ResourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInstance

`func (o *InstancePermissionItem) SetResourceInstance(v string)`

SetResourceInstance sets ResourceInstance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


