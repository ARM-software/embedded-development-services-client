<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# EntitlementsListFeedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableHalFeedLinks**](HalFeedLinks.md) |  | 
**Metadata** | [**NullablePagingMetadata**](PagingMetadata.md) |  | 
**Items** | [**[]InstancePermissionItem**](InstancePermissionItem.md) | A list of resource instances and the operations a user is allowed to perform on each. | 
**OwnerOperations** | Pointer to [**[]PermissionOperation**](PermissionOperation.md) | The list of allowed CRUDL operations for a resource owner. A resource owner is usually the user who created the resource, or someone who has been given ownership responsibilities—such as managing the the lifecycle of a resource. | [optional] [readonly] 

## Methods

### NewEntitlementsListFeedItem

`func NewEntitlementsListFeedItem(links NullableHalFeedLinks, metadata NullablePagingMetadata, items []InstancePermissionItem, ) *EntitlementsListFeedItem`

NewEntitlementsListFeedItem instantiates a new EntitlementsListFeedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsListFeedItemWithDefaults

`func NewEntitlementsListFeedItemWithDefaults() *EntitlementsListFeedItem`

NewEntitlementsListFeedItemWithDefaults instantiates a new EntitlementsListFeedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EntitlementsListFeedItem) GetLinks() HalFeedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntitlementsListFeedItem) GetLinksOk() (*HalFeedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntitlementsListFeedItem) SetLinks(v HalFeedLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *EntitlementsListFeedItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *EntitlementsListFeedItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *EntitlementsListFeedItem) GetMetadata() PagingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementsListFeedItem) GetMetadataOk() (*PagingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementsListFeedItem) SetMetadata(v PagingMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *EntitlementsListFeedItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EntitlementsListFeedItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetItems

`func (o *EntitlementsListFeedItem) GetItems() []InstancePermissionItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EntitlementsListFeedItem) GetItemsOk() (*[]InstancePermissionItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EntitlementsListFeedItem) SetItems(v []InstancePermissionItem)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *EntitlementsListFeedItem) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *EntitlementsListFeedItem) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetOwnerOperations

`func (o *EntitlementsListFeedItem) GetOwnerOperations() []PermissionOperation`

GetOwnerOperations returns the OwnerOperations field if non-nil, zero value otherwise.

### GetOwnerOperationsOk

`func (o *EntitlementsListFeedItem) GetOwnerOperationsOk() (*[]PermissionOperation, bool)`

GetOwnerOperationsOk returns a tuple with the OwnerOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOperations

`func (o *EntitlementsListFeedItem) SetOwnerOperations(v []PermissionOperation)`

SetOwnerOperations sets OwnerOperations field to given value.

### HasOwnerOperations

`func (o *EntitlementsListFeedItem) HasOwnerOperations() bool`

HasOwnerOperations returns a boolean if a field has been set.

### SetOwnerOperationsNil

`func (o *EntitlementsListFeedItem) SetOwnerOperationsNil(b bool)`

 SetOwnerOperationsNil sets the value for OwnerOperations to be an explicit nil

### UnsetOwnerOperations
`func (o *EntitlementsListFeedItem) UnsetOwnerOperations()`

UnsetOwnerOperations ensures that no value is present for OwnerOperations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


