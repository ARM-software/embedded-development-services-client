<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# TokenIntrospectionResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Active** | **bool** | Boolean indicator of whether or not the presented token is currently active. | 
**PolicyName** | Pointer to **string** | Name of the policy. | [optional] 

## Methods

### NewTokenIntrospectionResponseItem

`func NewTokenIntrospectionResponseItem(metadata NullableCommonMetadata, active bool, ) *TokenIntrospectionResponseItem`

NewTokenIntrospectionResponseItem instantiates a new TokenIntrospectionResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenIntrospectionResponseItemWithDefaults

`func NewTokenIntrospectionResponseItemWithDefaults() *TokenIntrospectionResponseItem`

NewTokenIntrospectionResponseItemWithDefaults instantiates a new TokenIntrospectionResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *TokenIntrospectionResponseItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TokenIntrospectionResponseItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TokenIntrospectionResponseItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *TokenIntrospectionResponseItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *TokenIntrospectionResponseItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetActive

`func (o *TokenIntrospectionResponseItem) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TokenIntrospectionResponseItem) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TokenIntrospectionResponseItem) SetActive(v bool)`

SetActive sets Active field to given value.


### GetPolicyName

`func (o *TokenIntrospectionResponseItem) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *TokenIntrospectionResponseItem) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *TokenIntrospectionResponseItem) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *TokenIntrospectionResponseItem) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


