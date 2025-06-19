<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAAdminCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**EmbeddedFPGAAdminItem**](EmbeddedFPGAAdminItem.md) |  | [optional] 
**Links** | [**NullableHalCollectionLinks**](HalCollectionLinks.md) |  | 
**Metadata** | [**NullablePagingMetadata**](PagingMetadata.md) |  | 
**Name** | **string** | ID of the Collection. | [readonly] 
**Title** | **string** | Human readable title of the collection. | [readonly] 

## Methods

### NewFPGAAdminCollection

`func NewFPGAAdminCollection(links NullableHalCollectionLinks, metadata NullablePagingMetadata, name string, title string, ) *FPGAAdminCollection`

NewFPGAAdminCollection instantiates a new FPGAAdminCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAAdminCollectionWithDefaults

`func NewFPGAAdminCollectionWithDefaults() *FPGAAdminCollection`

NewFPGAAdminCollectionWithDefaults instantiates a new FPGAAdminCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *FPGAAdminCollection) GetEmbedded() EmbeddedFPGAAdminItem`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *FPGAAdminCollection) GetEmbeddedOk() (*EmbeddedFPGAAdminItem, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *FPGAAdminCollection) SetEmbedded(v EmbeddedFPGAAdminItem)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *FPGAAdminCollection) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *FPGAAdminCollection) GetLinks() HalCollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FPGAAdminCollection) GetLinksOk() (*HalCollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FPGAAdminCollection) SetLinks(v HalCollectionLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *FPGAAdminCollection) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *FPGAAdminCollection) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *FPGAAdminCollection) GetMetadata() PagingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FPGAAdminCollection) GetMetadataOk() (*PagingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FPGAAdminCollection) SetMetadata(v PagingMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *FPGAAdminCollection) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *FPGAAdminCollection) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *FPGAAdminCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FPGAAdminCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FPGAAdminCollection) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *FPGAAdminCollection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FPGAAdminCollection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FPGAAdminCollection) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


