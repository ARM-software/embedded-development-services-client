<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# InstancePermissionItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Related** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewInstancePermissionItemLinks

`func NewInstancePermissionItemLinks(self HalLinkData, ) *InstancePermissionItemLinks`

NewInstancePermissionItemLinks instantiates a new InstancePermissionItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancePermissionItemLinksWithDefaults

`func NewInstancePermissionItemLinksWithDefaults() *InstancePermissionItemLinks`

NewInstancePermissionItemLinksWithDefaults instantiates a new InstancePermissionItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *InstancePermissionItemLinks) GetCollection() HalLinkData`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *InstancePermissionItemLinks) GetCollectionOk() (*HalLinkData, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *InstancePermissionItemLinks) SetCollection(v HalLinkData)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *InstancePermissionItemLinks) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetRelated

`func (o *InstancePermissionItemLinks) GetRelated() HalLinkData`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *InstancePermissionItemLinks) GetRelatedOk() (*HalLinkData, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *InstancePermissionItemLinks) SetRelated(v HalLinkData)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *InstancePermissionItemLinks) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetSelf

`func (o *InstancePermissionItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *InstancePermissionItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *InstancePermissionItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


