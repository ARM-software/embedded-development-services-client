<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# EntitlementsListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | Pointer to [**[]PermissionOperation**](PermissionOperation.md) | The list of CRUDL operations to check. | [optional] 
**ResourceType** | **string** | The type of resource for which permission is being checked. This should only refer to items and not collections. | 
**Token** | **string** | The API token of the user requesting access. This can be a JWT or an internal access token. | 

## Methods

### NewEntitlementsListRequest

`func NewEntitlementsListRequest(resourceType string, token string, ) *EntitlementsListRequest`

NewEntitlementsListRequest instantiates a new EntitlementsListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsListRequestWithDefaults

`func NewEntitlementsListRequestWithDefaults() *EntitlementsListRequest`

NewEntitlementsListRequestWithDefaults instantiates a new EntitlementsListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *EntitlementsListRequest) GetOperations() []PermissionOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *EntitlementsListRequest) GetOperationsOk() (*[]PermissionOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *EntitlementsListRequest) SetOperations(v []PermissionOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *EntitlementsListRequest) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### SetOperationsNil

`func (o *EntitlementsListRequest) SetOperationsNil(b bool)`

 SetOperationsNil sets the value for Operations to be an explicit nil

### UnsetOperations
`func (o *EntitlementsListRequest) UnsetOperations()`

UnsetOperations ensures that no value is present for Operations, not even an explicit nil
### GetResourceType

`func (o *EntitlementsListRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *EntitlementsListRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *EntitlementsListRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetToken

`func (o *EntitlementsListRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EntitlementsListRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EntitlementsListRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


