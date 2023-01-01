<!--
Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# ArtefactManagerLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clear** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Collection** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Download** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Related** | [**HalLinkData**](HalLinkData.md) |  | 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 
**Upload** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 

## Methods

### NewArtefactManagerLinks

`func NewArtefactManagerLinks(related HalLinkData, self HalLinkData, ) *ArtefactManagerLinks`

NewArtefactManagerLinks instantiates a new ArtefactManagerLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtefactManagerLinksWithDefaults

`func NewArtefactManagerLinksWithDefaults() *ArtefactManagerLinks`

NewArtefactManagerLinksWithDefaults instantiates a new ArtefactManagerLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClear

`func (o *ArtefactManagerLinks) GetClear() HalLinkData`

GetClear returns the Clear field if non-nil, zero value otherwise.

### GetClearOk

`func (o *ArtefactManagerLinks) GetClearOk() (*HalLinkData, bool)`

GetClearOk returns a tuple with the Clear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClear

`func (o *ArtefactManagerLinks) SetClear(v HalLinkData)`

SetClear sets Clear field to given value.

### HasClear

`func (o *ArtefactManagerLinks) HasClear() bool`

HasClear returns a boolean if a field has been set.

### GetCollection

`func (o *ArtefactManagerLinks) GetCollection() HalLinkData`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *ArtefactManagerLinks) GetCollectionOk() (*HalLinkData, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *ArtefactManagerLinks) SetCollection(v HalLinkData)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *ArtefactManagerLinks) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetDownload

`func (o *ArtefactManagerLinks) GetDownload() HalLinkData`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ArtefactManagerLinks) GetDownloadOk() (*HalLinkData, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ArtefactManagerLinks) SetDownload(v HalLinkData)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *ArtefactManagerLinks) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetRelated

`func (o *ArtefactManagerLinks) GetRelated() HalLinkData`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *ArtefactManagerLinks) GetRelatedOk() (*HalLinkData, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *ArtefactManagerLinks) SetRelated(v HalLinkData)`

SetRelated sets Related field to given value.


### GetSelf

`func (o *ArtefactManagerLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ArtefactManagerLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ArtefactManagerLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.


### GetUpload

`func (o *ArtefactManagerLinks) GetUpload() HalLinkData`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *ArtefactManagerLinks) GetUploadOk() (*HalLinkData, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *ArtefactManagerLinks) SetUpload(v HalLinkData)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *ArtefactManagerLinks) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


