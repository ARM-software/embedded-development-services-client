<!--
Copyright (C) 2020-2024 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# VhtRunJobItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancel** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Collection** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Delete** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Details** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Related** | [**HalLinkData**](HalLinkData.md) |  | 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewVhtRunJobItemLinks

`func NewVhtRunJobItemLinks(related HalLinkData, self HalLinkData, ) *VhtRunJobItemLinks`

NewVhtRunJobItemLinks instantiates a new VhtRunJobItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVhtRunJobItemLinksWithDefaults

`func NewVhtRunJobItemLinksWithDefaults() *VhtRunJobItemLinks`

NewVhtRunJobItemLinksWithDefaults instantiates a new VhtRunJobItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancel

`func (o *VhtRunJobItemLinks) GetCancel() HalLinkData`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *VhtRunJobItemLinks) GetCancelOk() (*HalLinkData, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *VhtRunJobItemLinks) SetCancel(v HalLinkData)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *VhtRunJobItemLinks) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### GetCollection

`func (o *VhtRunJobItemLinks) GetCollection() HalLinkData`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *VhtRunJobItemLinks) GetCollectionOk() (*HalLinkData, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *VhtRunJobItemLinks) SetCollection(v HalLinkData)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *VhtRunJobItemLinks) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetDelete

`func (o *VhtRunJobItemLinks) GetDelete() HalLinkData`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *VhtRunJobItemLinks) GetDeleteOk() (*HalLinkData, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *VhtRunJobItemLinks) SetDelete(v HalLinkData)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *VhtRunJobItemLinks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDetails

`func (o *VhtRunJobItemLinks) GetDetails() HalLinkData`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VhtRunJobItemLinks) GetDetailsOk() (*HalLinkData, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VhtRunJobItemLinks) SetDetails(v HalLinkData)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VhtRunJobItemLinks) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRelated

`func (o *VhtRunJobItemLinks) GetRelated() HalLinkData`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *VhtRunJobItemLinks) GetRelatedOk() (*HalLinkData, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *VhtRunJobItemLinks) SetRelated(v HalLinkData)`

SetRelated sets Related field to given value.


### GetSelf

`func (o *VhtRunJobItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *VhtRunJobItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *VhtRunJobItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


