<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# VhtInstanceCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**EmbeddedVhtInstanceItems**](EmbeddedVhtInstanceItems.md) |  | [optional] 
**Links** | [**NullableHalCollectionLinks**](HalCollectionLinks.md) |  | 
**Metadata** | [**NullablePagingMetadata**](PagingMetadata.md) |  | 
**Name** | **string** | ID of the Collection. | [readonly] 
**Title** | **string** | Human readable title of the collection. | [readonly] 

## Methods

### NewVhtInstanceCollection

`func NewVhtInstanceCollection(links NullableHalCollectionLinks, metadata NullablePagingMetadata, name string, title string, ) *VhtInstanceCollection`

NewVhtInstanceCollection instantiates a new VhtInstanceCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVhtInstanceCollectionWithDefaults

`func NewVhtInstanceCollectionWithDefaults() *VhtInstanceCollection`

NewVhtInstanceCollectionWithDefaults instantiates a new VhtInstanceCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *VhtInstanceCollection) GetEmbedded() EmbeddedVhtInstanceItems`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *VhtInstanceCollection) GetEmbeddedOk() (*EmbeddedVhtInstanceItems, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *VhtInstanceCollection) SetEmbedded(v EmbeddedVhtInstanceItems)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *VhtInstanceCollection) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *VhtInstanceCollection) GetLinks() HalCollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VhtInstanceCollection) GetLinksOk() (*HalCollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VhtInstanceCollection) SetLinks(v HalCollectionLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *VhtInstanceCollection) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *VhtInstanceCollection) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *VhtInstanceCollection) GetMetadata() PagingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VhtInstanceCollection) GetMetadataOk() (*PagingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VhtInstanceCollection) SetMetadata(v PagingMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *VhtInstanceCollection) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *VhtInstanceCollection) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *VhtInstanceCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VhtInstanceCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VhtInstanceCollection) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *VhtInstanceCollection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VhtInstanceCollection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VhtInstanceCollection) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


