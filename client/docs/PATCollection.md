<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# PATCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**EmbeddedPATItem**](EmbeddedPATItem.md) |  | [optional] 
**Links** | [**NullableHalOnlyEmbeddableCollectionLinks**](HalOnlyEmbeddableCollectionLinks.md) |  | 
**Metadata** | [**NullableCollectionMetadata**](CollectionMetadata.md) |  | 
**Name** | **string** | ID of the Collection. | [readonly] 
**Title** | **string** | Human readable title of the collection. | [readonly] 

## Methods

### NewPATCollection

`func NewPATCollection(links NullableHalOnlyEmbeddableCollectionLinks, metadata NullableCollectionMetadata, name string, title string, ) *PATCollection`

NewPATCollection instantiates a new PATCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCollectionWithDefaults

`func NewPATCollectionWithDefaults() *PATCollection`

NewPATCollectionWithDefaults instantiates a new PATCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *PATCollection) GetEmbedded() EmbeddedPATItem`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *PATCollection) GetEmbeddedOk() (*EmbeddedPATItem, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *PATCollection) SetEmbedded(v EmbeddedPATItem)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *PATCollection) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *PATCollection) GetLinks() HalOnlyEmbeddableCollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PATCollection) GetLinksOk() (*HalOnlyEmbeddableCollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PATCollection) SetLinks(v HalOnlyEmbeddableCollectionLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *PATCollection) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *PATCollection) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *PATCollection) GetMetadata() CollectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCollection) GetMetadataOk() (*CollectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCollection) SetMetadata(v CollectionMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *PATCollection) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCollection) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *PATCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCollection) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *PATCollection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PATCollection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PATCollection) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


