<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# HalSimpleCollectionLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alternate** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**First** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Last** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Next** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Prev** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewHalSimpleCollectionLinks

`func NewHalSimpleCollectionLinks(self HalLinkData, ) *HalSimpleCollectionLinks`

NewHalSimpleCollectionLinks instantiates a new HalSimpleCollectionLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHalSimpleCollectionLinksWithDefaults

`func NewHalSimpleCollectionLinksWithDefaults() *HalSimpleCollectionLinks`

NewHalSimpleCollectionLinksWithDefaults instantiates a new HalSimpleCollectionLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternate

`func (o *HalSimpleCollectionLinks) GetAlternate() HalLinkData`

GetAlternate returns the Alternate field if non-nil, zero value otherwise.

### GetAlternateOk

`func (o *HalSimpleCollectionLinks) GetAlternateOk() (*HalLinkData, bool)`

GetAlternateOk returns a tuple with the Alternate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternate

`func (o *HalSimpleCollectionLinks) SetAlternate(v HalLinkData)`

SetAlternate sets Alternate field to given value.

### HasAlternate

`func (o *HalSimpleCollectionLinks) HasAlternate() bool`

HasAlternate returns a boolean if a field has been set.

### GetFirst

`func (o *HalSimpleCollectionLinks) GetFirst() HalLinkData`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *HalSimpleCollectionLinks) GetFirstOk() (*HalLinkData, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *HalSimpleCollectionLinks) SetFirst(v HalLinkData)`

SetFirst sets First field to given value.

### HasFirst

`func (o *HalSimpleCollectionLinks) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *HalSimpleCollectionLinks) GetLast() HalLinkData`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *HalSimpleCollectionLinks) GetLastOk() (*HalLinkData, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *HalSimpleCollectionLinks) SetLast(v HalLinkData)`

SetLast sets Last field to given value.

### HasLast

`func (o *HalSimpleCollectionLinks) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNext

`func (o *HalSimpleCollectionLinks) GetNext() HalLinkData`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *HalSimpleCollectionLinks) GetNextOk() (*HalLinkData, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *HalSimpleCollectionLinks) SetNext(v HalLinkData)`

SetNext sets Next field to given value.

### HasNext

`func (o *HalSimpleCollectionLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *HalSimpleCollectionLinks) GetPrev() HalLinkData`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *HalSimpleCollectionLinks) GetPrevOk() (*HalLinkData, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *HalSimpleCollectionLinks) SetPrev(v HalLinkData)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *HalSimpleCollectionLinks) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetSelf

`func (o *HalSimpleCollectionLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *HalSimpleCollectionLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *HalSimpleCollectionLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


