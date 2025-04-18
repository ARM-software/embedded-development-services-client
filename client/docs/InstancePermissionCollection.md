<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# InstancePermissionCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**EmbeddedInstancePermissionItem**](EmbeddedInstancePermissionItem.md) |  | [optional] 
**Links** | [**NullableHalCollectionLinks**](HalCollectionLinks.md) |  | 
**Metadata** | [**NullablePagingMetadata**](PagingMetadata.md) |  | 
**Name** | **string** | ID of the Collection. | [readonly] 
**Title** | **string** | Human readable title of the collection. | [readonly] 

## Methods

### NewInstancePermissionCollection

`func NewInstancePermissionCollection(links NullableHalCollectionLinks, metadata NullablePagingMetadata, name string, title string, ) *InstancePermissionCollection`

NewInstancePermissionCollection instantiates a new InstancePermissionCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancePermissionCollectionWithDefaults

`func NewInstancePermissionCollectionWithDefaults() *InstancePermissionCollection`

NewInstancePermissionCollectionWithDefaults instantiates a new InstancePermissionCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *InstancePermissionCollection) GetEmbedded() EmbeddedInstancePermissionItem`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *InstancePermissionCollection) GetEmbeddedOk() (*EmbeddedInstancePermissionItem, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *InstancePermissionCollection) SetEmbedded(v EmbeddedInstancePermissionItem)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *InstancePermissionCollection) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *InstancePermissionCollection) GetLinks() HalCollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InstancePermissionCollection) GetLinksOk() (*HalCollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InstancePermissionCollection) SetLinks(v HalCollectionLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *InstancePermissionCollection) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *InstancePermissionCollection) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *InstancePermissionCollection) GetMetadata() PagingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InstancePermissionCollection) GetMetadataOk() (*PagingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InstancePermissionCollection) SetMetadata(v PagingMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *InstancePermissionCollection) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *InstancePermissionCollection) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *InstancePermissionCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstancePermissionCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstancePermissionCollection) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *InstancePermissionCollection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InstancePermissionCollection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InstancePermissionCollection) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


