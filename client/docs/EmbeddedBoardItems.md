<!--
Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# EmbeddedBoardItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**[]BoardItem**](BoardItem.md) |  | [optional] 

## Methods

### NewEmbeddedBoardItems

`func NewEmbeddedBoardItems() *EmbeddedBoardItems`

NewEmbeddedBoardItems instantiates a new EmbeddedBoardItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedBoardItemsWithDefaults

`func NewEmbeddedBoardItemsWithDefaults() *EmbeddedBoardItems`

NewEmbeddedBoardItemsWithDefaults instantiates a new EmbeddedBoardItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *EmbeddedBoardItems) GetItem() []BoardItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *EmbeddedBoardItems) GetItemOk() (*[]BoardItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *EmbeddedBoardItems) SetItem(v []BoardItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *EmbeddedBoardItems) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


