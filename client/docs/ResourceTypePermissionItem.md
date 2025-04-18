<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# ResourceTypePermissionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**ResourceTypePermissionItemLinks**](ResourceTypePermissionItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Name** | **string** | The name of the resource type. | 
**Operations** | **[]string** | The list of allowed CRUDL operations. | 
**Title** | Pointer to **string** | Optional description of the permission. | [optional] 

## Methods

### NewResourceTypePermissionItem

`func NewResourceTypePermissionItem(links ResourceTypePermissionItemLinks, metadata NullableCommonMetadata, name string, operations []string, ) *ResourceTypePermissionItem`

NewResourceTypePermissionItem instantiates a new ResourceTypePermissionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTypePermissionItemWithDefaults

`func NewResourceTypePermissionItemWithDefaults() *ResourceTypePermissionItem`

NewResourceTypePermissionItemWithDefaults instantiates a new ResourceTypePermissionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ResourceTypePermissionItem) GetLinks() ResourceTypePermissionItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceTypePermissionItem) GetLinksOk() (*ResourceTypePermissionItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceTypePermissionItem) SetLinks(v ResourceTypePermissionItemLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *ResourceTypePermissionItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceTypePermissionItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceTypePermissionItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *ResourceTypePermissionItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ResourceTypePermissionItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *ResourceTypePermissionItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceTypePermissionItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceTypePermissionItem) SetName(v string)`

SetName sets Name field to given value.


### GetOperations

`func (o *ResourceTypePermissionItem) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ResourceTypePermissionItem) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ResourceTypePermissionItem) SetOperations(v []string)`

SetOperations sets Operations field to given value.


### SetOperationsNil

`func (o *ResourceTypePermissionItem) SetOperationsNil(b bool)`

 SetOperationsNil sets the value for Operations to be an explicit nil

### UnsetOperations
`func (o *ResourceTypePermissionItem) UnsetOperations()`

UnsetOperations ensures that no value is present for Operations, not even an explicit nil
### GetTitle

`func (o *ResourceTypePermissionItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResourceTypePermissionItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResourceTypePermissionItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResourceTypePermissionItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


