<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# RetainWorkspaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | Pointer to **NullableInt64** | Time to live (in seconds), i.e. The the number of seconds from the current time (when this request is received) for which to keep the workspace and associated resources. | [optional] 

## Methods

### NewRetainWorkspaceRequest

`func NewRetainWorkspaceRequest() *RetainWorkspaceRequest`

NewRetainWorkspaceRequest instantiates a new RetainWorkspaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetainWorkspaceRequestWithDefaults

`func NewRetainWorkspaceRequestWithDefaults() *RetainWorkspaceRequest`

NewRetainWorkspaceRequestWithDefaults instantiates a new RetainWorkspaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *RetainWorkspaceRequest) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RetainWorkspaceRequest) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RetainWorkspaceRequest) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RetainWorkspaceRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *RetainWorkspaceRequest) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *RetainWorkspaceRequest) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


