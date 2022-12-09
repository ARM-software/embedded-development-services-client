<!--
Copyright (C) 2020-2022 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# RetainBuildJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | Pointer to **NullableInt64** | Time to live (in seconds), i.e. The the number of seconds from the current time (when this request is received) for which to keep the job and associated resources. | [optional] 

## Methods

### NewRetainBuildJobRequest

`func NewRetainBuildJobRequest() *RetainBuildJobRequest`

NewRetainBuildJobRequest instantiates a new RetainBuildJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetainBuildJobRequestWithDefaults

`func NewRetainBuildJobRequestWithDefaults() *RetainBuildJobRequest`

NewRetainBuildJobRequestWithDefaults instantiates a new RetainBuildJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *RetainBuildJobRequest) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RetainBuildJobRequest) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RetainBuildJobRequest) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RetainBuildJobRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *RetainBuildJobRequest) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *RetainBuildJobRequest) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


