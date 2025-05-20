<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGATargetID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraKeys** | Pointer to **map[string]string** | Additional keys which may be required to differentiate between targets when software and hardware images are not enough. | [optional] 
**HardwareImage** | **string** | identifier of the hardware image | 
**SoftwareImage** | **string** | identifier of the software image | 

## Methods

### NewFPGATargetID

`func NewFPGATargetID(hardwareImage string, softwareImage string, ) *FPGATargetID`

NewFPGATargetID instantiates a new FPGATargetID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGATargetIDWithDefaults

`func NewFPGATargetIDWithDefaults() *FPGATargetID`

NewFPGATargetIDWithDefaults instantiates a new FPGATargetID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraKeys

`func (o *FPGATargetID) GetExtraKeys() map[string]string`

GetExtraKeys returns the ExtraKeys field if non-nil, zero value otherwise.

### GetExtraKeysOk

`func (o *FPGATargetID) GetExtraKeysOk() (*map[string]string, bool)`

GetExtraKeysOk returns a tuple with the ExtraKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraKeys

`func (o *FPGATargetID) SetExtraKeys(v map[string]string)`

SetExtraKeys sets ExtraKeys field to given value.

### HasExtraKeys

`func (o *FPGATargetID) HasExtraKeys() bool`

HasExtraKeys returns a boolean if a field has been set.

### SetExtraKeysNil

`func (o *FPGATargetID) SetExtraKeysNil(b bool)`

 SetExtraKeysNil sets the value for ExtraKeys to be an explicit nil

### UnsetExtraKeys
`func (o *FPGATargetID) UnsetExtraKeys()`

UnsetExtraKeys ensures that no value is present for ExtraKeys, not even an explicit nil
### GetHardwareImage

`func (o *FPGATargetID) GetHardwareImage() string`

GetHardwareImage returns the HardwareImage field if non-nil, zero value otherwise.

### GetHardwareImageOk

`func (o *FPGATargetID) GetHardwareImageOk() (*string, bool)`

GetHardwareImageOk returns a tuple with the HardwareImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareImage

`func (o *FPGATargetID) SetHardwareImage(v string)`

SetHardwareImage sets HardwareImage field to given value.


### GetSoftwareImage

`func (o *FPGATargetID) GetSoftwareImage() string`

GetSoftwareImage returns the SoftwareImage field if non-nil, zero value otherwise.

### GetSoftwareImageOk

`func (o *FPGATargetID) GetSoftwareImageOk() (*string, bool)`

GetSoftwareImageOk returns a tuple with the SoftwareImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImage

`func (o *FPGATargetID) SetSoftwareImage(v string)`

SetSoftwareImage sets SoftwareImage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


