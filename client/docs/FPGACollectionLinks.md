<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGACollectionLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alternate** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Create** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Embedded** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**First** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Item** | Pointer to [**[]HalLinkData**](HalLinkData.md) |  | [optional] 
**Last** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Next** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Prev** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 
**Simple** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 

## Methods

### NewFPGACollectionLinks

`func NewFPGACollectionLinks(self HalLinkData, ) *FPGACollectionLinks`

NewFPGACollectionLinks instantiates a new FPGACollectionLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGACollectionLinksWithDefaults

`func NewFPGACollectionLinksWithDefaults() *FPGACollectionLinks`

NewFPGACollectionLinksWithDefaults instantiates a new FPGACollectionLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternate

`func (o *FPGACollectionLinks) GetAlternate() HalLinkData`

GetAlternate returns the Alternate field if non-nil, zero value otherwise.

### GetAlternateOk

`func (o *FPGACollectionLinks) GetAlternateOk() (*HalLinkData, bool)`

GetAlternateOk returns a tuple with the Alternate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternate

`func (o *FPGACollectionLinks) SetAlternate(v HalLinkData)`

SetAlternate sets Alternate field to given value.

### HasAlternate

`func (o *FPGACollectionLinks) HasAlternate() bool`

HasAlternate returns a boolean if a field has been set.

### GetCreate

`func (o *FPGACollectionLinks) GetCreate() HalLinkData`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *FPGACollectionLinks) GetCreateOk() (*HalLinkData, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *FPGACollectionLinks) SetCreate(v HalLinkData)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *FPGACollectionLinks) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetEmbedded

`func (o *FPGACollectionLinks) GetEmbedded() HalLinkData`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *FPGACollectionLinks) GetEmbeddedOk() (*HalLinkData, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *FPGACollectionLinks) SetEmbedded(v HalLinkData)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *FPGACollectionLinks) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetFirst

`func (o *FPGACollectionLinks) GetFirst() HalLinkData`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *FPGACollectionLinks) GetFirstOk() (*HalLinkData, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *FPGACollectionLinks) SetFirst(v HalLinkData)`

SetFirst sets First field to given value.

### HasFirst

`func (o *FPGACollectionLinks) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetItem

`func (o *FPGACollectionLinks) GetItem() []HalLinkData`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *FPGACollectionLinks) GetItemOk() (*[]HalLinkData, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *FPGACollectionLinks) SetItem(v []HalLinkData)`

SetItem sets Item field to given value.

### HasItem

`func (o *FPGACollectionLinks) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetLast

`func (o *FPGACollectionLinks) GetLast() HalLinkData`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *FPGACollectionLinks) GetLastOk() (*HalLinkData, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *FPGACollectionLinks) SetLast(v HalLinkData)`

SetLast sets Last field to given value.

### HasLast

`func (o *FPGACollectionLinks) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNext

`func (o *FPGACollectionLinks) GetNext() HalLinkData`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FPGACollectionLinks) GetNextOk() (*HalLinkData, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FPGACollectionLinks) SetNext(v HalLinkData)`

SetNext sets Next field to given value.

### HasNext

`func (o *FPGACollectionLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *FPGACollectionLinks) GetPrev() HalLinkData`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *FPGACollectionLinks) GetPrevOk() (*HalLinkData, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *FPGACollectionLinks) SetPrev(v HalLinkData)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *FPGACollectionLinks) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetSelf

`func (o *FPGACollectionLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *FPGACollectionLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *FPGACollectionLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.


### GetSimple

`func (o *FPGACollectionLinks) GetSimple() HalLinkData`

GetSimple returns the Simple field if non-nil, zero value otherwise.

### GetSimpleOk

`func (o *FPGACollectionLinks) GetSimpleOk() (*HalLinkData, bool)`

GetSimpleOk returns a tuple with the Simple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimple

`func (o *FPGACollectionLinks) SetSimple(v HalLinkData)`

SetSimple sets Simple field to given value.

### HasSimple

`func (o *FPGACollectionLinks) HasSimple() bool`

HasSimple returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


