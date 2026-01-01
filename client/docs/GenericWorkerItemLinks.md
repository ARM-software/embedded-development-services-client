<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# GenericWorkerItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Create** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewGenericWorkerItemLinks

`func NewGenericWorkerItemLinks(self HalLinkData, ) *GenericWorkerItemLinks`

NewGenericWorkerItemLinks instantiates a new GenericWorkerItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericWorkerItemLinksWithDefaults

`func NewGenericWorkerItemLinksWithDefaults() *GenericWorkerItemLinks`

NewGenericWorkerItemLinksWithDefaults instantiates a new GenericWorkerItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *GenericWorkerItemLinks) GetCollection() HalLinkData`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *GenericWorkerItemLinks) GetCollectionOk() (*HalLinkData, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *GenericWorkerItemLinks) SetCollection(v HalLinkData)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *GenericWorkerItemLinks) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetCreate

`func (o *GenericWorkerItemLinks) GetCreate() HalLinkData`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *GenericWorkerItemLinks) GetCreateOk() (*HalLinkData, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *GenericWorkerItemLinks) SetCreate(v HalLinkData)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *GenericWorkerItemLinks) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetSelf

`func (o *GenericWorkerItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *GenericWorkerItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *GenericWorkerItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


