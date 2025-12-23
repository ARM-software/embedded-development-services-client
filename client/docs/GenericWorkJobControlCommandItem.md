<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# GenericWorkJobControlCommandItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **map[string]string** | Optional arguments supplied with a job control command. | [optional] 
**Name** | **string** | Name of a control command within a namespace. | 
**Namespace** | **string** | The logical scope of a control command. | 

## Methods

### NewGenericWorkJobControlCommandItem

`func NewGenericWorkJobControlCommandItem(name string, namespace string, ) *GenericWorkJobControlCommandItem`

NewGenericWorkJobControlCommandItem instantiates a new GenericWorkJobControlCommandItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericWorkJobControlCommandItemWithDefaults

`func NewGenericWorkJobControlCommandItemWithDefaults() *GenericWorkJobControlCommandItem`

NewGenericWorkJobControlCommandItemWithDefaults instantiates a new GenericWorkJobControlCommandItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *GenericWorkJobControlCommandItem) GetArgs() map[string]string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *GenericWorkJobControlCommandItem) GetArgsOk() (*map[string]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *GenericWorkJobControlCommandItem) SetArgs(v map[string]string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *GenericWorkJobControlCommandItem) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *GenericWorkJobControlCommandItem) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *GenericWorkJobControlCommandItem) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetName

`func (o *GenericWorkJobControlCommandItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenericWorkJobControlCommandItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenericWorkJobControlCommandItem) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *GenericWorkJobControlCommandItem) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GenericWorkJobControlCommandItem) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GenericWorkJobControlCommandItem) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


