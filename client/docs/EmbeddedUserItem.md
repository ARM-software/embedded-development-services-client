<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# EmbeddedUserItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**[]UserItem**](UserItem.md) |  | [optional] 

## Methods

### NewEmbeddedUserItem

`func NewEmbeddedUserItem() *EmbeddedUserItem`

NewEmbeddedUserItem instantiates a new EmbeddedUserItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedUserItemWithDefaults

`func NewEmbeddedUserItemWithDefaults() *EmbeddedUserItem`

NewEmbeddedUserItemWithDefaults instantiates a new EmbeddedUserItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *EmbeddedUserItem) GetItem() []UserItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *EmbeddedUserItem) GetItemOk() (*[]UserItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *EmbeddedUserItem) SetItem(v []UserItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *EmbeddedUserItem) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


