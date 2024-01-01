<!--
Copyright (C) 2020-2024 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# AdditionalTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the tool. | 
**Name** | **string** | Unique ID of the tool. | 
**Title** | **string** | Human readable name of the tool. | 
**Version** | **NullableString** | Version of the tool. | 

## Methods

### NewAdditionalTool

`func NewAdditionalTool(description string, name string, title string, version NullableString, ) *AdditionalTool`

NewAdditionalTool instantiates a new AdditionalTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalToolWithDefaults

`func NewAdditionalToolWithDefaults() *AdditionalTool`

NewAdditionalToolWithDefaults instantiates a new AdditionalTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AdditionalTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdditionalTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdditionalTool) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *AdditionalTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdditionalTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdditionalTool) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AdditionalTool) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdditionalTool) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdditionalTool) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVersion

`func (o *AdditionalTool) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdditionalTool) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdditionalTool) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *AdditionalTool) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *AdditionalTool) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


