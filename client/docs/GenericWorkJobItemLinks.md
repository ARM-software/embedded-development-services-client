<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# GenericWorkJobItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artefacts** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Cancel** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Collection** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Delete** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Details** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Related** | [**HalLinkData**](HalLinkData.md) |  | 
**Retain** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 
**Sunset** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 

## Methods

### NewGenericWorkJobItemLinks

`func NewGenericWorkJobItemLinks(related HalLinkData, self HalLinkData, ) *GenericWorkJobItemLinks`

NewGenericWorkJobItemLinks instantiates a new GenericWorkJobItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericWorkJobItemLinksWithDefaults

`func NewGenericWorkJobItemLinksWithDefaults() *GenericWorkJobItemLinks`

NewGenericWorkJobItemLinksWithDefaults instantiates a new GenericWorkJobItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtefacts

`func (o *GenericWorkJobItemLinks) GetArtefacts() HalLinkData`

GetArtefacts returns the Artefacts field if non-nil, zero value otherwise.

### GetArtefactsOk

`func (o *GenericWorkJobItemLinks) GetArtefactsOk() (*HalLinkData, bool)`

GetArtefactsOk returns a tuple with the Artefacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtefacts

`func (o *GenericWorkJobItemLinks) SetArtefacts(v HalLinkData)`

SetArtefacts sets Artefacts field to given value.

### HasArtefacts

`func (o *GenericWorkJobItemLinks) HasArtefacts() bool`

HasArtefacts returns a boolean if a field has been set.

### GetCancel

`func (o *GenericWorkJobItemLinks) GetCancel() HalLinkData`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *GenericWorkJobItemLinks) GetCancelOk() (*HalLinkData, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *GenericWorkJobItemLinks) SetCancel(v HalLinkData)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *GenericWorkJobItemLinks) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### GetCollection

`func (o *GenericWorkJobItemLinks) GetCollection() HalLinkData`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *GenericWorkJobItemLinks) GetCollectionOk() (*HalLinkData, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *GenericWorkJobItemLinks) SetCollection(v HalLinkData)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *GenericWorkJobItemLinks) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetDelete

`func (o *GenericWorkJobItemLinks) GetDelete() HalLinkData`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *GenericWorkJobItemLinks) GetDeleteOk() (*HalLinkData, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *GenericWorkJobItemLinks) SetDelete(v HalLinkData)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *GenericWorkJobItemLinks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDetails

`func (o *GenericWorkJobItemLinks) GetDetails() HalLinkData`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GenericWorkJobItemLinks) GetDetailsOk() (*HalLinkData, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GenericWorkJobItemLinks) SetDetails(v HalLinkData)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GenericWorkJobItemLinks) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRelated

`func (o *GenericWorkJobItemLinks) GetRelated() HalLinkData`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *GenericWorkJobItemLinks) GetRelatedOk() (*HalLinkData, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *GenericWorkJobItemLinks) SetRelated(v HalLinkData)`

SetRelated sets Related field to given value.


### GetRetain

`func (o *GenericWorkJobItemLinks) GetRetain() HalLinkData`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *GenericWorkJobItemLinks) GetRetainOk() (*HalLinkData, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *GenericWorkJobItemLinks) SetRetain(v HalLinkData)`

SetRetain sets Retain field to given value.

### HasRetain

`func (o *GenericWorkJobItemLinks) HasRetain() bool`

HasRetain returns a boolean if a field has been set.

### GetSelf

`func (o *GenericWorkJobItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *GenericWorkJobItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *GenericWorkJobItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.


### GetSunset

`func (o *GenericWorkJobItemLinks) GetSunset() HalLinkData`

GetSunset returns the Sunset field if non-nil, zero value otherwise.

### GetSunsetOk

`func (o *GenericWorkJobItemLinks) GetSunsetOk() (*HalLinkData, bool)`

GetSunsetOk returns a tuple with the Sunset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunset

`func (o *GenericWorkJobItemLinks) SetSunset(v HalLinkData)`

SetSunset sets Sunset field to given value.

### HasSunset

`func (o *GenericWorkJobItemLinks) HasSunset() bool`

HasSunset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


