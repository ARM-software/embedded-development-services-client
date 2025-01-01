<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# UserItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Me** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewUserItemLinks

`func NewUserItemLinks(self HalLinkData, ) *UserItemLinks`

NewUserItemLinks instantiates a new UserItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserItemLinksWithDefaults

`func NewUserItemLinksWithDefaults() *UserItemLinks`

NewUserItemLinksWithDefaults instantiates a new UserItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMe

`func (o *UserItemLinks) GetMe() HalLinkData`

GetMe returns the Me field if non-nil, zero value otherwise.

### GetMeOk

`func (o *UserItemLinks) GetMeOk() (*HalLinkData, bool)`

GetMeOk returns a tuple with the Me field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMe

`func (o *UserItemLinks) SetMe(v HalLinkData)`

SetMe sets Me field to given value.

### HasMe

`func (o *UserItemLinks) HasMe() bool`

HasMe returns a boolean if a field has been set.

### GetSelf

`func (o *UserItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *UserItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *UserItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


