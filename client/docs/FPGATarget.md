<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGATarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | a human-readable description of the target | [optional] 
**Details** | Pointer to **map[string]string** | Extra metadata describing FPGA targets | [optional] 
**Identifier** | [**FPGATargetID**](FPGATargetID.md) |  | 
**Title** | Pointer to **string** | a human-readable name for the target | [optional] 

## Methods

### NewFPGATarget

`func NewFPGATarget(identifier FPGATargetID, ) *FPGATarget`

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
### GetDetails

`func (o *FPGATarget) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *FPGATarget) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *FPGATarget) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *FPGATarget) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetIdentifier

`func (o *FPGATarget) GetIdentifier() FPGATargetID`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *FPGATarget) GetIdentifierOk() (*FPGATargetID, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *FPGATarget) SetIdentifier(v FPGATargetID)`

SetIdentifier sets Identifier field to given value.


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


