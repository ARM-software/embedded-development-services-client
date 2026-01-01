<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# HalCollectionLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alternate** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Embedded** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**First** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Item** | Pointer to [**[]HalLinkData**](HalLinkData.md) |  | [optional] 
**Last** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Next** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Prev** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 
**Simple** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 

## Methods

### NewHalCollectionLinks

`func NewHalCollectionLinks(self HalLinkData, ) *HalCollectionLinks`

NewHalCollectionLinks instantiates a new HalCollectionLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHalCollectionLinksWithDefaults

`func NewHalCollectionLinksWithDefaults() *HalCollectionLinks`

NewHalCollectionLinksWithDefaults instantiates a new HalCollectionLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternate

`func (o *HalCollectionLinks) GetAlternate() HalLinkData`

GetAlternate returns the Alternate field if non-nil, zero value otherwise.

### GetAlternateOk

`func (o *HalCollectionLinks) GetAlternateOk() (*HalLinkData, bool)`

GetAlternateOk returns a tuple with the Alternate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternate

`func (o *HalCollectionLinks) SetAlternate(v HalLinkData)`

SetAlternate sets Alternate field to given value.

### HasAlternate

`func (o *HalCollectionLinks) HasAlternate() bool`

HasAlternate returns a boolean if a field has been set.

### GetEmbedded

`func (o *HalCollectionLinks) GetEmbedded() HalLinkData`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *HalCollectionLinks) GetEmbeddedOk() (*HalLinkData, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *HalCollectionLinks) SetEmbedded(v HalLinkData)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *HalCollectionLinks) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetFirst

`func (o *HalCollectionLinks) GetFirst() HalLinkData`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *HalCollectionLinks) GetFirstOk() (*HalLinkData, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *HalCollectionLinks) SetFirst(v HalLinkData)`

SetFirst sets First field to given value.

### HasFirst

`func (o *HalCollectionLinks) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetItem

`func (o *HalCollectionLinks) GetItem() []HalLinkData`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *HalCollectionLinks) GetItemOk() (*[]HalLinkData, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *HalCollectionLinks) SetItem(v []HalLinkData)`

SetItem sets Item field to given value.

### HasItem

`func (o *HalCollectionLinks) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetLast

`func (o *HalCollectionLinks) GetLast() HalLinkData`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *HalCollectionLinks) GetLastOk() (*HalLinkData, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *HalCollectionLinks) SetLast(v HalLinkData)`

SetLast sets Last field to given value.

### HasLast

`func (o *HalCollectionLinks) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNext

`func (o *HalCollectionLinks) GetNext() HalLinkData`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *HalCollectionLinks) GetNextOk() (*HalLinkData, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *HalCollectionLinks) SetNext(v HalLinkData)`

SetNext sets Next field to given value.

### HasNext

`func (o *HalCollectionLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *HalCollectionLinks) GetPrev() HalLinkData`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *HalCollectionLinks) GetPrevOk() (*HalLinkData, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *HalCollectionLinks) SetPrev(v HalLinkData)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *HalCollectionLinks) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetSelf

`func (o *HalCollectionLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *HalCollectionLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *HalCollectionLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.


### GetSimple

`func (o *HalCollectionLinks) GetSimple() HalLinkData`

GetSimple returns the Simple field if non-nil, zero value otherwise.

### GetSimpleOk

`func (o *HalCollectionLinks) GetSimpleOk() (*HalLinkData, bool)`

GetSimpleOk returns a tuple with the Simple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimple

`func (o *HalCollectionLinks) SetSimple(v HalLinkData)`

SetSimple sets Simple field to given value.

### HasSimple

`func (o *HalCollectionLinks) HasSimple() bool`

HasSimple returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


