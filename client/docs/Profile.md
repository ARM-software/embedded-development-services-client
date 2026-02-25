<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Unique identifier for the profile. | [optional] 
**Mapping** | Pointer to **NullableString** | Mapping name for report tools. | [optional] 
**Name** | **string** | A reference to the profile name used within the analyser. | 

## Methods

### NewProfile

`func NewProfile(name string, ) *Profile`

NewProfile instantiates a new Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileWithDefaults

`func NewProfileWithDefaults() *Profile`

NewProfileWithDefaults instantiates a new Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Profile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Profile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Profile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Profile) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Profile) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Profile) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMapping

`func (o *Profile) GetMapping() string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *Profile) GetMappingOk() (*string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *Profile) SetMapping(v string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *Profile) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### SetMappingNil

`func (o *Profile) SetMappingNil(b bool)`

 SetMappingNil sets the value for Mapping to be an explicit nil

### UnsetMapping
`func (o *Profile) UnsetMapping()`

UnsetMapping ensures that no value is present for Mapping, not even an explicit nil
### GetName

`func (o *Profile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Profile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Profile) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


