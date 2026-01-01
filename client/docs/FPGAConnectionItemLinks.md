<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAConnectionItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connect** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Related** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewFPGAConnectionItemLinks

`func NewFPGAConnectionItemLinks(self HalLinkData, ) *FPGAConnectionItemLinks`

NewFPGAConnectionItemLinks instantiates a new FPGAConnectionItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAConnectionItemLinksWithDefaults

`func NewFPGAConnectionItemLinksWithDefaults() *FPGAConnectionItemLinks`

NewFPGAConnectionItemLinksWithDefaults instantiates a new FPGAConnectionItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnect

`func (o *FPGAConnectionItemLinks) GetConnect() HalLinkData`

GetConnect returns the Connect field if non-nil, zero value otherwise.

### GetConnectOk

`func (o *FPGAConnectionItemLinks) GetConnectOk() (*HalLinkData, bool)`

GetConnectOk returns a tuple with the Connect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnect

`func (o *FPGAConnectionItemLinks) SetConnect(v HalLinkData)`

SetConnect sets Connect field to given value.

### HasConnect

`func (o *FPGAConnectionItemLinks) HasConnect() bool`

HasConnect returns a boolean if a field has been set.

### GetRelated

`func (o *FPGAConnectionItemLinks) GetRelated() HalLinkData`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *FPGAConnectionItemLinks) GetRelatedOk() (*HalLinkData, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *FPGAConnectionItemLinks) SetRelated(v HalLinkData)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *FPGAConnectionItemLinks) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetSelf

`func (o *FPGAConnectionItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *FPGAConnectionItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *FPGAConnectionItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


