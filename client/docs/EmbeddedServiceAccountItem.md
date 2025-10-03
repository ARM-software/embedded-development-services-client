<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# EmbeddedServiceAccountItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**[]ServiceAccountItem**](ServiceAccountItem.md) |  | [optional] 

## Methods

### NewEmbeddedServiceAccountItem

`func NewEmbeddedServiceAccountItem() *EmbeddedServiceAccountItem`

NewEmbeddedServiceAccountItem instantiates a new EmbeddedServiceAccountItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedServiceAccountItemWithDefaults

`func NewEmbeddedServiceAccountItemWithDefaults() *EmbeddedServiceAccountItem`

NewEmbeddedServiceAccountItemWithDefaults instantiates a new EmbeddedServiceAccountItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *EmbeddedServiceAccountItem) GetItem() []ServiceAccountItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *EmbeddedServiceAccountItem) GetItemOk() (*[]ServiceAccountItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *EmbeddedServiceAccountItem) SetItem(v []ServiceAccountItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *EmbeddedServiceAccountItem) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


