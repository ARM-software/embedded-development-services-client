<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# PermissionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Operations** | Pointer to **[]string** | The list of allowed CRUDL operations. | [optional] 
**ResourceInstance** | **string** | The unique identifier of the resource instance. | 
**ResourceType** | **string** | The type of resource for which permission is being checked. This should only refer to items and not collections. | 
**UserName** | **string** | The unique identifier of the user requesting access. | 
**UserToken** | Pointer to **string** | The API token of the user requesting access. This can be a JWT or an internal access token. | [optional] 

## Methods

### NewPermissionItem

`func NewPermissionItem(metadata NullableCommonMetadata, resourceInstance string, resourceType string, userName string, ) *PermissionItem`

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

`func (o *PermissionItem) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *PermissionItem) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *PermissionItem) SetOperations(v []string)`

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


### GetUserName

`func (o *PermissionItem) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *PermissionItem) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *PermissionItem) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetUserToken

`func (o *PermissionItem) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *PermissionItem) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *PermissionItem) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *PermissionItem) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


