<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# ServiceAccountItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Delete** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Edit** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewServiceAccountItemLinks

`func NewServiceAccountItemLinks(self HalLinkData, ) *ServiceAccountItemLinks`

NewServiceAccountItemLinks instantiates a new ServiceAccountItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountItemLinksWithDefaults

`func NewServiceAccountItemLinksWithDefaults() *ServiceAccountItemLinks`

NewServiceAccountItemLinksWithDefaults instantiates a new ServiceAccountItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *ServiceAccountItemLinks) GetCollection() HalLinkData`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *ServiceAccountItemLinks) GetCollectionOk() (*HalLinkData, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *ServiceAccountItemLinks) SetCollection(v HalLinkData)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *ServiceAccountItemLinks) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetDelete

`func (o *ServiceAccountItemLinks) GetDelete() HalLinkData`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ServiceAccountItemLinks) GetDeleteOk() (*HalLinkData, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ServiceAccountItemLinks) SetDelete(v HalLinkData)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ServiceAccountItemLinks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetEdit

`func (o *ServiceAccountItemLinks) GetEdit() HalLinkData`

GetEdit returns the Edit field if non-nil, zero value otherwise.

### GetEditOk

`func (o *ServiceAccountItemLinks) GetEditOk() (*HalLinkData, bool)`

GetEditOk returns a tuple with the Edit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdit

`func (o *ServiceAccountItemLinks) SetEdit(v HalLinkData)`

SetEdit sets Edit field to given value.

### HasEdit

`func (o *ServiceAccountItemLinks) HasEdit() bool`

HasEdit returns a boolean if a field has been set.

### GetSelf

`func (o *ServiceAccountItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ServiceAccountItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ServiceAccountItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


