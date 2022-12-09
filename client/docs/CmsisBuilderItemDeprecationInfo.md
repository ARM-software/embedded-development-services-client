<!--
Copyright (C) 2020-2022 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# CmsisBuilderItemDeprecationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] [readonly] 
**Issued** | Pointer to **time.Time** |  | [optional] [readonly] 
**Removal** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCmsisBuilderItemDeprecationInfo

`func NewCmsisBuilderItemDeprecationInfo() *CmsisBuilderItemDeprecationInfo`

NewCmsisBuilderItemDeprecationInfo instantiates a new CmsisBuilderItemDeprecationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmsisBuilderItemDeprecationInfoWithDefaults

`func NewCmsisBuilderItemDeprecationInfoWithDefaults() *CmsisBuilderItemDeprecationInfo`

NewCmsisBuilderItemDeprecationInfoWithDefaults instantiates a new CmsisBuilderItemDeprecationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CmsisBuilderItemDeprecationInfo) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CmsisBuilderItemDeprecationInfo) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CmsisBuilderItemDeprecationInfo) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CmsisBuilderItemDeprecationInfo) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetIssued

`func (o *CmsisBuilderItemDeprecationInfo) GetIssued() time.Time`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *CmsisBuilderItemDeprecationInfo) GetIssuedOk() (*time.Time, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *CmsisBuilderItemDeprecationInfo) SetIssued(v time.Time)`

SetIssued sets Issued field to given value.

### HasIssued

`func (o *CmsisBuilderItemDeprecationInfo) HasIssued() bool`

HasIssued returns a boolean if a field has been set.

### GetRemoval

`func (o *CmsisBuilderItemDeprecationInfo) GetRemoval() time.Time`

GetRemoval returns the Removal field if non-nil, zero value otherwise.

### GetRemovalOk

`func (o *CmsisBuilderItemDeprecationInfo) GetRemovalOk() (*time.Time, bool)`

GetRemovalOk returns a tuple with the Removal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoval

`func (o *CmsisBuilderItemDeprecationInfo) SetRemoval(v time.Time)`

SetRemoval sets Removal field to given value.

### HasRemoval

`func (o *CmsisBuilderItemDeprecationInfo) HasRemoval() bool`

HasRemoval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


