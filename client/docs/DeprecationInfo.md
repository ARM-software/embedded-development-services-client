<!--
Copyright (C) 2020-2022 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# DeprecationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | further explanation of the rationale behind the deprecation | 
**Issued** | **time.Time** | time when the deprecation notice was first issued | 
**Removal** | **time.Time** | time when the removal will be effective | 

## Methods

### NewDeprecationInfo

`func NewDeprecationInfo(comment string, issued time.Time, removal time.Time, ) *DeprecationInfo`

NewDeprecationInfo instantiates a new DeprecationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecationInfoWithDefaults

`func NewDeprecationInfoWithDefaults() *DeprecationInfo`

NewDeprecationInfoWithDefaults instantiates a new DeprecationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *DeprecationInfo) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DeprecationInfo) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DeprecationInfo) SetComment(v string)`

SetComment sets Comment field to given value.


### GetIssued

`func (o *DeprecationInfo) GetIssued() time.Time`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *DeprecationInfo) GetIssuedOk() (*time.Time, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *DeprecationInfo) SetIssued(v time.Time)`

SetIssued sets Issued field to given value.


### GetRemoval

`func (o *DeprecationInfo) GetRemoval() time.Time`

GetRemoval returns the Removal field if non-nil, zero value otherwise.

### GetRemovalOk

`func (o *DeprecationInfo) GetRemovalOk() (*time.Time, bool)`

GetRemovalOk returns a tuple with the Removal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoval

`func (o *DeprecationInfo) SetRemoval(v time.Time)`

SetRemoval sets Removal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


