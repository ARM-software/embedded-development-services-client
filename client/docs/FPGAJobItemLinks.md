<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAJobItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artefacts** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Cancel** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Collection** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Connect** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Delete** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Details** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Related** | [**HalLinkData**](HalLinkData.md) |  | 
**Retain** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewFPGAJobItemLinks

`func NewFPGAJobItemLinks(related HalLinkData, self HalLinkData, ) *FPGAJobItemLinks`

NewFPGAJobItemLinks instantiates a new FPGAJobItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAJobItemLinksWithDefaults

`func NewFPGAJobItemLinksWithDefaults() *FPGAJobItemLinks`

NewFPGAJobItemLinksWithDefaults instantiates a new FPGAJobItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtefacts

`func (o *FPGAJobItemLinks) GetArtefacts() HalLinkData`

GetArtefacts returns the Artefacts field if non-nil, zero value otherwise.

### GetArtefactsOk

`func (o *FPGAJobItemLinks) GetArtefactsOk() (*HalLinkData, bool)`

GetArtefactsOk returns a tuple with the Artefacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtefacts

`func (o *FPGAJobItemLinks) SetArtefacts(v HalLinkData)`

SetArtefacts sets Artefacts field to given value.

### HasArtefacts

`func (o *FPGAJobItemLinks) HasArtefacts() bool`

HasArtefacts returns a boolean if a field has been set.

### GetCancel

`func (o *FPGAJobItemLinks) GetCancel() HalLinkData`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *FPGAJobItemLinks) GetCancelOk() (*HalLinkData, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *FPGAJobItemLinks) SetCancel(v HalLinkData)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *FPGAJobItemLinks) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### GetCollection

`func (o *FPGAJobItemLinks) GetCollection() HalLinkData`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *FPGAJobItemLinks) GetCollectionOk() (*HalLinkData, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *FPGAJobItemLinks) SetCollection(v HalLinkData)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *FPGAJobItemLinks) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetConnect

`func (o *FPGAJobItemLinks) GetConnect() HalLinkData`

GetConnect returns the Connect field if non-nil, zero value otherwise.

### GetConnectOk

`func (o *FPGAJobItemLinks) GetConnectOk() (*HalLinkData, bool)`

GetConnectOk returns a tuple with the Connect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnect

`func (o *FPGAJobItemLinks) SetConnect(v HalLinkData)`

SetConnect sets Connect field to given value.

### HasConnect

`func (o *FPGAJobItemLinks) HasConnect() bool`

HasConnect returns a boolean if a field has been set.

### GetDelete

`func (o *FPGAJobItemLinks) GetDelete() HalLinkData`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *FPGAJobItemLinks) GetDeleteOk() (*HalLinkData, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *FPGAJobItemLinks) SetDelete(v HalLinkData)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *FPGAJobItemLinks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDetails

`func (o *FPGAJobItemLinks) GetDetails() HalLinkData`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *FPGAJobItemLinks) GetDetailsOk() (*HalLinkData, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *FPGAJobItemLinks) SetDetails(v HalLinkData)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *FPGAJobItemLinks) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRelated

`func (o *FPGAJobItemLinks) GetRelated() HalLinkData`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *FPGAJobItemLinks) GetRelatedOk() (*HalLinkData, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *FPGAJobItemLinks) SetRelated(v HalLinkData)`

SetRelated sets Related field to given value.


### GetRetain

`func (o *FPGAJobItemLinks) GetRetain() HalLinkData`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *FPGAJobItemLinks) GetRetainOk() (*HalLinkData, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *FPGAJobItemLinks) SetRetain(v HalLinkData)`

SetRetain sets Retain field to given value.

### HasRetain

`func (o *FPGAJobItemLinks) HasRetain() bool`

HasRetain returns a boolean if a field has been set.

### GetSelf

`func (o *FPGAJobItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *FPGAJobItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *FPGAJobItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


