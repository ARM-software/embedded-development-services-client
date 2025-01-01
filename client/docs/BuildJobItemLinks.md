<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# BuildJobItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artefacts** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Cancel** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Collection** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Delete** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Describedby** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Details** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Outputs** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Related** | [**HalLinkData**](HalLinkData.md) |  | 
**Retain** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewBuildJobItemLinks

`func NewBuildJobItemLinks(related HalLinkData, self HalLinkData, ) *BuildJobItemLinks`

NewBuildJobItemLinks instantiates a new BuildJobItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildJobItemLinksWithDefaults

`func NewBuildJobItemLinksWithDefaults() *BuildJobItemLinks`

NewBuildJobItemLinksWithDefaults instantiates a new BuildJobItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtefacts

`func (o *BuildJobItemLinks) GetArtefacts() HalLinkData`

GetArtefacts returns the Artefacts field if non-nil, zero value otherwise.

### GetArtefactsOk

`func (o *BuildJobItemLinks) GetArtefactsOk() (*HalLinkData, bool)`

GetArtefactsOk returns a tuple with the Artefacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtefacts

`func (o *BuildJobItemLinks) SetArtefacts(v HalLinkData)`

SetArtefacts sets Artefacts field to given value.

### HasArtefacts

`func (o *BuildJobItemLinks) HasArtefacts() bool`

HasArtefacts returns a boolean if a field has been set.

### GetCancel

`func (o *BuildJobItemLinks) GetCancel() HalLinkData`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *BuildJobItemLinks) GetCancelOk() (*HalLinkData, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *BuildJobItemLinks) SetCancel(v HalLinkData)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *BuildJobItemLinks) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### GetCollection

`func (o *BuildJobItemLinks) GetCollection() HalLinkData`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *BuildJobItemLinks) GetCollectionOk() (*HalLinkData, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *BuildJobItemLinks) SetCollection(v HalLinkData)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *BuildJobItemLinks) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetDelete

`func (o *BuildJobItemLinks) GetDelete() HalLinkData`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *BuildJobItemLinks) GetDeleteOk() (*HalLinkData, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *BuildJobItemLinks) SetDelete(v HalLinkData)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *BuildJobItemLinks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDescribedby

`func (o *BuildJobItemLinks) GetDescribedby() HalLinkData`

GetDescribedby returns the Describedby field if non-nil, zero value otherwise.

### GetDescribedbyOk

`func (o *BuildJobItemLinks) GetDescribedbyOk() (*HalLinkData, bool)`

GetDescribedbyOk returns a tuple with the Describedby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribedby

`func (o *BuildJobItemLinks) SetDescribedby(v HalLinkData)`

SetDescribedby sets Describedby field to given value.

### HasDescribedby

`func (o *BuildJobItemLinks) HasDescribedby() bool`

HasDescribedby returns a boolean if a field has been set.

### GetDetails

`func (o *BuildJobItemLinks) GetDetails() HalLinkData`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BuildJobItemLinks) GetDetailsOk() (*HalLinkData, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BuildJobItemLinks) SetDetails(v HalLinkData)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BuildJobItemLinks) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetOutputs

`func (o *BuildJobItemLinks) GetOutputs() HalLinkData`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *BuildJobItemLinks) GetOutputsOk() (*HalLinkData, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *BuildJobItemLinks) SetOutputs(v HalLinkData)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *BuildJobItemLinks) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetRelated

`func (o *BuildJobItemLinks) GetRelated() HalLinkData`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *BuildJobItemLinks) GetRelatedOk() (*HalLinkData, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *BuildJobItemLinks) SetRelated(v HalLinkData)`

SetRelated sets Related field to given value.


### GetRetain

`func (o *BuildJobItemLinks) GetRetain() HalLinkData`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *BuildJobItemLinks) GetRetainOk() (*HalLinkData, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *BuildJobItemLinks) SetRetain(v HalLinkData)`

SetRetain sets Retain field to given value.

### HasRetain

`func (o *BuildJobItemLinks) HasRetain() bool`

HasRetain returns a boolean if a field has been set.

### GetSelf

`func (o *BuildJobItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *BuildJobItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *BuildJobItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


