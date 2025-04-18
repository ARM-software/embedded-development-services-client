<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# InstancePermissionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableInstancePermissionItemLinks**](InstancePermissionItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Name** | **string** | The unique identifier of the resource instance. | 
**Operations** | **[]string** | The list of allowed CRUDL operations. | 
**Title** | Pointer to **string** | Optional description of the permission. | [optional] 

## Methods

### NewInstancePermissionItem

`func NewInstancePermissionItem(links NullableInstancePermissionItemLinks, metadata NullableCommonMetadata, name string, operations []string, ) *InstancePermissionItem`

NewInstancePermissionItem instantiates a new InstancePermissionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancePermissionItemWithDefaults

`func NewInstancePermissionItemWithDefaults() *InstancePermissionItem`

NewInstancePermissionItemWithDefaults instantiates a new InstancePermissionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *InstancePermissionItem) GetLinks() InstancePermissionItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InstancePermissionItem) GetLinksOk() (*InstancePermissionItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InstancePermissionItem) SetLinks(v InstancePermissionItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *InstancePermissionItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *InstancePermissionItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *InstancePermissionItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InstancePermissionItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InstancePermissionItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *InstancePermissionItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *InstancePermissionItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *InstancePermissionItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstancePermissionItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstancePermissionItem) SetName(v string)`

SetName sets Name field to given value.


### GetOperations

`func (o *InstancePermissionItem) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *InstancePermissionItem) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *InstancePermissionItem) SetOperations(v []string)`

SetOperations sets Operations field to given value.


### SetOperationsNil

`func (o *InstancePermissionItem) SetOperationsNil(b bool)`

 SetOperationsNil sets the value for Operations to be an explicit nil

### UnsetOperations
`func (o *InstancePermissionItem) UnsetOperations()`

UnsetOperations ensures that no value is present for Operations, not even an explicit nil
### GetTitle

`func (o *InstancePermissionItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InstancePermissionItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InstancePermissionItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InstancePermissionItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


