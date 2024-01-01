<!--
Copyright (C) 2020-2024 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# CommonMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ctime** | **time.Time** | Creation Time: UTC date and time (in RFC3339 format) when the resource was created. If this is a system created resource, this will be a fixed time unaffected by user actions. | [readonly] 
**Etime** | Pointer to **NullableTime** | Expiry Time: UTC date and time (in RFC3339 format) when the resource will be removed automatically by the system and become unavailable. | [optional] [readonly] 
**Mtime** | **time.Time** | Last Modification Time: UTC date and time (in RFC3339 format) when the resource was last updated. For a resource that cannot be modified this will be the same as &#x60;ctime&#x60;. | [readonly] 

## Methods

### NewCommonMetadata

`func NewCommonMetadata(ctime time.Time, mtime time.Time, ) *CommonMetadata`

NewCommonMetadata instantiates a new CommonMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonMetadataWithDefaults

`func NewCommonMetadataWithDefaults() *CommonMetadata`

NewCommonMetadataWithDefaults instantiates a new CommonMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCtime

`func (o *CommonMetadata) GetCtime() time.Time`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *CommonMetadata) GetCtimeOk() (*time.Time, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *CommonMetadata) SetCtime(v time.Time)`

SetCtime sets Ctime field to given value.


### GetEtime

`func (o *CommonMetadata) GetEtime() time.Time`

GetEtime returns the Etime field if non-nil, zero value otherwise.

### GetEtimeOk

`func (o *CommonMetadata) GetEtimeOk() (*time.Time, bool)`

GetEtimeOk returns a tuple with the Etime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtime

`func (o *CommonMetadata) SetEtime(v time.Time)`

SetEtime sets Etime field to given value.

### HasEtime

`func (o *CommonMetadata) HasEtime() bool`

HasEtime returns a boolean if a field has been set.

### SetEtimeNil

`func (o *CommonMetadata) SetEtimeNil(b bool)`

 SetEtimeNil sets the value for Etime to be an explicit nil

### UnsetEtime
`func (o *CommonMetadata) UnsetEtime()`

UnsetEtime ensures that no value is present for Etime, not even an explicit nil
### GetMtime

`func (o *CommonMetadata) GetMtime() time.Time`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *CommonMetadata) GetMtimeOk() (*time.Time, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *CommonMetadata) SetMtime(v time.Time)`

SetMtime sets Mtime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


