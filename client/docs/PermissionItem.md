<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# PermissionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Operations** | Pointer to [**[]PermissionOperation**](PermissionOperation.md) | The list of allowed CRUDL operations. | [optional] 
**OwnerOperations** | Pointer to [**[]PermissionOperation**](PermissionOperation.md) | The list of allowed CRUDL operations for a resource owner. A resource owner is usually the user who created the resource, or someone who has been given ownership responsibilities—such as managing the the lifecycle of a resource. | [optional] [readonly] 
**ResourceInstance** | Pointer to **string** | The unique identifier of the resource instance. | [optional] 
**ResourceType** | **string** | The type of resource for which permission is being checked. This should only refer to items and not collections. | 
**Token** | **string** | The API token of the user requesting access. This can be a JWT or an internal access token. | 

## Methods

### NewPermissionItem

`func NewPermissionItem(metadata NullableCommonMetadata, resourceType string, token string, ) *PermissionItem`

NewPermissionItem instantiates a new PermissionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionItemWithDefaults

`func NewPermissionItemWithDefaults() *PermissionItem`

NewPermissionItemWithDefaults instantiates a new PermissionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *PermissionItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PermissionItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PermissionItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *PermissionItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PermissionItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOperations

`func (o *PermissionItem) GetOperations() []PermissionOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *PermissionItem) GetOperationsOk() (*[]PermissionOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *PermissionItem) SetOperations(v []PermissionOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *PermissionItem) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### SetOperationsNil

`func (o *PermissionItem) SetOperationsNil(b bool)`

 SetOperationsNil sets the value for Operations to be an explicit nil

### UnsetOperations
`func (o *PermissionItem) UnsetOperations()`

UnsetOperations ensures that no value is present for Operations, not even an explicit nil
### GetOwnerOperations

`func (o *PermissionItem) GetOwnerOperations() []PermissionOperation`

GetOwnerOperations returns the OwnerOperations field if non-nil, zero value otherwise.

### GetOwnerOperationsOk

`func (o *PermissionItem) GetOwnerOperationsOk() (*[]PermissionOperation, bool)`

GetOwnerOperationsOk returns a tuple with the OwnerOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOperations

`func (o *PermissionItem) SetOwnerOperations(v []PermissionOperation)`

SetOwnerOperations sets OwnerOperations field to given value.

### HasOwnerOperations

`func (o *PermissionItem) HasOwnerOperations() bool`

HasOwnerOperations returns a boolean if a field has been set.

### SetOwnerOperationsNil

`func (o *PermissionItem) SetOwnerOperationsNil(b bool)`

 SetOwnerOperationsNil sets the value for OwnerOperations to be an explicit nil

### UnsetOwnerOperations
`func (o *PermissionItem) UnsetOwnerOperations()`

UnsetOwnerOperations ensures that no value is present for OwnerOperations, not even an explicit nil
### GetResourceInstance

`func (o *PermissionItem) GetResourceInstance() string`

GetResourceInstance returns the ResourceInstance field if non-nil, zero value otherwise.

### GetResourceInstanceOk

`func (o *PermissionItem) GetResourceInstanceOk() (*string, bool)`

GetResourceInstanceOk returns a tuple with the ResourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInstance

`func (o *PermissionItem) SetResourceInstance(v string)`

SetResourceInstance sets ResourceInstance field to given value.

### HasResourceInstance

`func (o *PermissionItem) HasResourceInstance() bool`

HasResourceInstance returns a boolean if a field has been set.

### GetResourceType

`func (o *PermissionItem) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *PermissionItem) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *PermissionItem) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetToken

`func (o *PermissionItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PermissionItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PermissionItem) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


