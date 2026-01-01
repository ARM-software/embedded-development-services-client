<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# AccessTokenItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Delete** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewAccessTokenItemLinks

`func NewAccessTokenItemLinks(self HalLinkData, ) *AccessTokenItemLinks`

NewAccessTokenItemLinks instantiates a new AccessTokenItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenItemLinksWithDefaults

`func NewAccessTokenItemLinksWithDefaults() *AccessTokenItemLinks`

NewAccessTokenItemLinksWithDefaults instantiates a new AccessTokenItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *AccessTokenItemLinks) GetAuthor() HalLinkData`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AccessTokenItemLinks) GetAuthorOk() (*HalLinkData, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AccessTokenItemLinks) SetAuthor(v HalLinkData)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AccessTokenItemLinks) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetDelete

`func (o *AccessTokenItemLinks) GetDelete() HalLinkData`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *AccessTokenItemLinks) GetDeleteOk() (*HalLinkData, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *AccessTokenItemLinks) SetDelete(v HalLinkData)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *AccessTokenItemLinks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetSelf

`func (o *AccessTokenItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AccessTokenItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AccessTokenItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


