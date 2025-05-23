<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# UserCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**EmbeddedUserItem**](EmbeddedUserItem.md) |  | [optional] 
**Links** | [**NullableHalOnlyEmbeddableCollectionLinks**](HalOnlyEmbeddableCollectionLinks.md) |  | 
**Metadata** | [**NullableCollectionMetadata**](CollectionMetadata.md) |  | 
**Name** | **string** | ID of the Collection. | [readonly] 
**Title** | **string** | Human readable title of the collection. | [readonly] 

## Methods

### NewUserCollection

`func NewUserCollection(links NullableHalOnlyEmbeddableCollectionLinks, metadata NullableCollectionMetadata, name string, title string, ) *UserCollection`

NewUserCollection instantiates a new UserCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCollectionWithDefaults

`func NewUserCollectionWithDefaults() *UserCollection`

NewUserCollectionWithDefaults instantiates a new UserCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *UserCollection) GetEmbedded() EmbeddedUserItem`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *UserCollection) GetEmbeddedOk() (*EmbeddedUserItem, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *UserCollection) SetEmbedded(v EmbeddedUserItem)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *UserCollection) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *UserCollection) GetLinks() HalOnlyEmbeddableCollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserCollection) GetLinksOk() (*HalOnlyEmbeddableCollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserCollection) SetLinks(v HalOnlyEmbeddableCollectionLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *UserCollection) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *UserCollection) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *UserCollection) GetMetadata() CollectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserCollection) GetMetadataOk() (*CollectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserCollection) SetMetadata(v CollectionMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *UserCollection) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *UserCollection) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *UserCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserCollection) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *UserCollection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserCollection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserCollection) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


