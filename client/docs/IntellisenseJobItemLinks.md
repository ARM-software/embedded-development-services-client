<!--
Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# IntellisenseJobItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artefacts** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Cancel** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Collection** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Delete** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Details** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Outputs** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Related** | [**HalLinkData**](HalLinkData.md) |  | 
**Retain** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewIntellisenseJobItemLinks

`func NewIntellisenseJobItemLinks(related HalLinkData, self HalLinkData, ) *IntellisenseJobItemLinks`

NewIntellisenseJobItemLinks instantiates a new IntellisenseJobItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntellisenseJobItemLinksWithDefaults

`func NewIntellisenseJobItemLinksWithDefaults() *IntellisenseJobItemLinks`

NewIntellisenseJobItemLinksWithDefaults instantiates a new IntellisenseJobItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtefacts

`func (o *IntellisenseJobItemLinks) GetArtefacts() HalLinkData`

GetArtefacts returns the Artefacts field if non-nil, zero value otherwise.

### GetArtefactsOk

`func (o *IntellisenseJobItemLinks) GetArtefactsOk() (*HalLinkData, bool)`

GetArtefactsOk returns a tuple with the Artefacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtefacts

`func (o *IntellisenseJobItemLinks) SetArtefacts(v HalLinkData)`

SetArtefacts sets Artefacts field to given value.

### HasArtefacts

`func (o *IntellisenseJobItemLinks) HasArtefacts() bool`

HasArtefacts returns a boolean if a field has been set.

### GetCancel

`func (o *IntellisenseJobItemLinks) GetCancel() HalLinkData`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *IntellisenseJobItemLinks) GetCancelOk() (*HalLinkData, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *IntellisenseJobItemLinks) SetCancel(v HalLinkData)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *IntellisenseJobItemLinks) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### GetCollection

`func (o *IntellisenseJobItemLinks) GetCollection() HalLinkData`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *IntellisenseJobItemLinks) GetCollectionOk() (*HalLinkData, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *IntellisenseJobItemLinks) SetCollection(v HalLinkData)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *IntellisenseJobItemLinks) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetDelete

`func (o *IntellisenseJobItemLinks) GetDelete() HalLinkData`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *IntellisenseJobItemLinks) GetDeleteOk() (*HalLinkData, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *IntellisenseJobItemLinks) SetDelete(v HalLinkData)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *IntellisenseJobItemLinks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDetails

`func (o *IntellisenseJobItemLinks) GetDetails() HalLinkData`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *IntellisenseJobItemLinks) GetDetailsOk() (*HalLinkData, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *IntellisenseJobItemLinks) SetDetails(v HalLinkData)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *IntellisenseJobItemLinks) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetOutputs

`func (o *IntellisenseJobItemLinks) GetOutputs() HalLinkData`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *IntellisenseJobItemLinks) GetOutputsOk() (*HalLinkData, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *IntellisenseJobItemLinks) SetOutputs(v HalLinkData)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *IntellisenseJobItemLinks) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetRelated

`func (o *IntellisenseJobItemLinks) GetRelated() HalLinkData`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *IntellisenseJobItemLinks) GetRelatedOk() (*HalLinkData, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *IntellisenseJobItemLinks) SetRelated(v HalLinkData)`

SetRelated sets Related field to given value.


### GetRetain

`func (o *IntellisenseJobItemLinks) GetRetain() HalLinkData`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *IntellisenseJobItemLinks) GetRetainOk() (*HalLinkData, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *IntellisenseJobItemLinks) SetRetain(v HalLinkData)`

SetRetain sets Retain field to given value.

### HasRetain

`func (o *IntellisenseJobItemLinks) HasRetain() bool`

HasRetain returns a boolean if a field has been set.

### GetSelf

`func (o *IntellisenseJobItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *IntellisenseJobItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *IntellisenseJobItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


