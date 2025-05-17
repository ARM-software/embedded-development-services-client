<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGATarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | a human-readable description of the target | [optional] 
**HardwareImage** | **string** | identifier of the hardware image | 
**SoftwareImage** | **string** | identifier of the software image | 
**Title** | Pointer to **string** | a human-readable name for the target | [optional] 

## Methods

### NewFPGATarget

`func NewFPGATarget(hardwareImage string, softwareImage string, ) *FPGATarget`

NewFPGATarget instantiates a new FPGATarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGATargetWithDefaults

`func NewFPGATargetWithDefaults() *FPGATarget`

NewFPGATargetWithDefaults instantiates a new FPGATarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FPGATarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FPGATarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FPGATarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FPGATarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FPGATarget) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FPGATarget) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHardwareImage

`func (o *FPGATarget) GetHardwareImage() string`

GetHardwareImage returns the HardwareImage field if non-nil, zero value otherwise.

### GetHardwareImageOk

`func (o *FPGATarget) GetHardwareImageOk() (*string, bool)`

GetHardwareImageOk returns a tuple with the HardwareImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareImage

`func (o *FPGATarget) SetHardwareImage(v string)`

SetHardwareImage sets HardwareImage field to given value.


### GetSoftwareImage

`func (o *FPGATarget) GetSoftwareImage() string`

GetSoftwareImage returns the SoftwareImage field if non-nil, zero value otherwise.

### GetSoftwareImageOk

`func (o *FPGATarget) GetSoftwareImageOk() (*string, bool)`

GetSoftwareImageOk returns a tuple with the SoftwareImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImage

`func (o *FPGATarget) SetSoftwareImage(v string)`

SetSoftwareImage sets SoftwareImage field to given value.


### GetTitle

`func (o *FPGATarget) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FPGATarget) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FPGATarget) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FPGATarget) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


