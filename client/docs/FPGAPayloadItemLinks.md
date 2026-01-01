<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAPayloadItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | [**HalLinkData**](HalLinkData.md) |  | 
**Delete** | [**HalLinkData**](HalLinkData.md) |  | 
**Related** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewFPGAPayloadItemLinks

`func NewFPGAPayloadItemLinks(collection HalLinkData, delete HalLinkData, self HalLinkData, ) *FPGAPayloadItemLinks`

NewFPGAPayloadItemLinks instantiates a new FPGAPayloadItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAPayloadItemLinksWithDefaults

`func NewFPGAPayloadItemLinksWithDefaults() *FPGAPayloadItemLinks`

NewFPGAPayloadItemLinksWithDefaults instantiates a new FPGAPayloadItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *FPGAPayloadItemLinks) GetCollection() HalLinkData`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *FPGAPayloadItemLinks) GetCollectionOk() (*HalLinkData, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *FPGAPayloadItemLinks) SetCollection(v HalLinkData)`

SetCollection sets Collection field to given value.


### GetDelete

`func (o *FPGAPayloadItemLinks) GetDelete() HalLinkData`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *FPGAPayloadItemLinks) GetDeleteOk() (*HalLinkData, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *FPGAPayloadItemLinks) SetDelete(v HalLinkData)`

SetDelete sets Delete field to given value.


### GetRelated

`func (o *FPGAPayloadItemLinks) GetRelated() HalLinkData`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *FPGAPayloadItemLinks) GetRelatedOk() (*HalLinkData, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *FPGAPayloadItemLinks) SetRelated(v HalLinkData)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *FPGAPayloadItemLinks) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetSelf

`func (o *FPGAPayloadItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *FPGAPayloadItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *FPGAPayloadItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


