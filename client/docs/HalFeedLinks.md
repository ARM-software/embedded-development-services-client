<!--
Copyright (C) 2020-2024 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# HalFeedLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alternate** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**First** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Future** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Last** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Next** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Prev** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewHalFeedLinks

`func NewHalFeedLinks(self HalLinkData, ) *HalFeedLinks`

NewHalFeedLinks instantiates a new HalFeedLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHalFeedLinksWithDefaults

`func NewHalFeedLinksWithDefaults() *HalFeedLinks`

NewHalFeedLinksWithDefaults instantiates a new HalFeedLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternate

`func (o *HalFeedLinks) GetAlternate() HalLinkData`

GetAlternate returns the Alternate field if non-nil, zero value otherwise.

### GetAlternateOk

`func (o *HalFeedLinks) GetAlternateOk() (*HalLinkData, bool)`

GetAlternateOk returns a tuple with the Alternate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternate

`func (o *HalFeedLinks) SetAlternate(v HalLinkData)`

SetAlternate sets Alternate field to given value.

### HasAlternate

`func (o *HalFeedLinks) HasAlternate() bool`

HasAlternate returns a boolean if a field has been set.

### GetFirst

`func (o *HalFeedLinks) GetFirst() HalLinkData`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *HalFeedLinks) GetFirstOk() (*HalLinkData, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *HalFeedLinks) SetFirst(v HalLinkData)`

SetFirst sets First field to given value.

### HasFirst

`func (o *HalFeedLinks) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetFuture

`func (o *HalFeedLinks) GetFuture() HalLinkData`

GetFuture returns the Future field if non-nil, zero value otherwise.

### GetFutureOk

`func (o *HalFeedLinks) GetFutureOk() (*HalLinkData, bool)`

GetFutureOk returns a tuple with the Future field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuture

`func (o *HalFeedLinks) SetFuture(v HalLinkData)`

SetFuture sets Future field to given value.

### HasFuture

`func (o *HalFeedLinks) HasFuture() bool`

HasFuture returns a boolean if a field has been set.

### GetLast

`func (o *HalFeedLinks) GetLast() HalLinkData`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *HalFeedLinks) GetLastOk() (*HalLinkData, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *HalFeedLinks) SetLast(v HalLinkData)`

SetLast sets Last field to given value.

### HasLast

`func (o *HalFeedLinks) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetNext

`func (o *HalFeedLinks) GetNext() HalLinkData`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *HalFeedLinks) GetNextOk() (*HalLinkData, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *HalFeedLinks) SetNext(v HalLinkData)`

SetNext sets Next field to given value.

### HasNext

`func (o *HalFeedLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *HalFeedLinks) GetPrev() HalLinkData`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *HalFeedLinks) GetPrevOk() (*HalLinkData, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *HalFeedLinks) SetPrev(v HalLinkData)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *HalFeedLinks) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetSelf

`func (o *HalFeedLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *HalFeedLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *HalFeedLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


