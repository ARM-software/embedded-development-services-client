<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAPayloadOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique ID of the user that owns this FPGA Payload. | [readonly] 
**Title** | **string** | The human readable name of the user that owns this FPGA Payload. | [readonly] 

## Methods

### NewFPGAPayloadOwner

`func NewFPGAPayloadOwner(name string, title string, ) *FPGAPayloadOwner`

NewFPGAPayloadOwner instantiates a new FPGAPayloadOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAPayloadOwnerWithDefaults

`func NewFPGAPayloadOwnerWithDefaults() *FPGAPayloadOwner`

NewFPGAPayloadOwnerWithDefaults instantiates a new FPGAPayloadOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FPGAPayloadOwner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FPGAPayloadOwner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FPGAPayloadOwner) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *FPGAPayloadOwner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FPGAPayloadOwner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FPGAPayloadOwner) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


