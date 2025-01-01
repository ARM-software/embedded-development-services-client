<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# WorkspaceItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artefacts** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Collection** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Delete** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Details** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Related** | [**HalLinkData**](HalLinkData.md) |  | 
**Retain** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewWorkspaceItemLinks

`func NewWorkspaceItemLinks(related HalLinkData, self HalLinkData, ) *WorkspaceItemLinks`

NewWorkspaceItemLinks instantiates a new WorkspaceItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceItemLinksWithDefaults

`func NewWorkspaceItemLinksWithDefaults() *WorkspaceItemLinks`

NewWorkspaceItemLinksWithDefaults instantiates a new WorkspaceItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtefacts

`func (o *WorkspaceItemLinks) GetArtefacts() HalLinkData`

GetArtefacts returns the Artefacts field if non-nil, zero value otherwise.

### GetArtefactsOk

`func (o *WorkspaceItemLinks) GetArtefactsOk() (*HalLinkData, bool)`

GetArtefactsOk returns a tuple with the Artefacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtefacts

`func (o *WorkspaceItemLinks) SetArtefacts(v HalLinkData)`

SetArtefacts sets Artefacts field to given value.

### HasArtefacts

`func (o *WorkspaceItemLinks) HasArtefacts() bool`

HasArtefacts returns a boolean if a field has been set.

### GetCollection

`func (o *WorkspaceItemLinks) GetCollection() HalLinkData`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *WorkspaceItemLinks) GetCollectionOk() (*HalLinkData, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *WorkspaceItemLinks) SetCollection(v HalLinkData)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *WorkspaceItemLinks) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetDelete

`func (o *WorkspaceItemLinks) GetDelete() HalLinkData`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *WorkspaceItemLinks) GetDeleteOk() (*HalLinkData, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *WorkspaceItemLinks) SetDelete(v HalLinkData)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *WorkspaceItemLinks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDetails

`func (o *WorkspaceItemLinks) GetDetails() HalLinkData`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *WorkspaceItemLinks) GetDetailsOk() (*HalLinkData, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *WorkspaceItemLinks) SetDetails(v HalLinkData)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *WorkspaceItemLinks) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRelated

`func (o *WorkspaceItemLinks) GetRelated() HalLinkData`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *WorkspaceItemLinks) GetRelatedOk() (*HalLinkData, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *WorkspaceItemLinks) SetRelated(v HalLinkData)`

SetRelated sets Related field to given value.


### GetRetain

`func (o *WorkspaceItemLinks) GetRetain() HalLinkData`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *WorkspaceItemLinks) GetRetainOk() (*HalLinkData, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *WorkspaceItemLinks) SetRetain(v HalLinkData)`

SetRetain sets Retain field to given value.

### HasRetain

`func (o *WorkspaceItemLinks) HasRetain() bool`

HasRetain returns a boolean if a field has been set.

### GetSelf

`func (o *WorkspaceItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *WorkspaceItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *WorkspaceItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


