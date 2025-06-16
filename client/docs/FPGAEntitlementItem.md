<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAEntitlementItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Fpga** | **string** | The identifier of the FPGA resource. | 
**Name** | **string** | The unique identifier of the resource instance. | 
**Users** | **[]string** | A list of user IDs permitted to access the FPGA. | 

## Methods

### NewFPGAEntitlementItem

`func NewFPGAEntitlementItem(metadata NullableCommonMetadata, fpga string, name string, users []string, ) *FPGAEntitlementItem`

NewFPGAEntitlementItem instantiates a new FPGAEntitlementItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAEntitlementItemWithDefaults

`func NewFPGAEntitlementItemWithDefaults() *FPGAEntitlementItem`

NewFPGAEntitlementItemWithDefaults instantiates a new FPGAEntitlementItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *FPGAEntitlementItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FPGAEntitlementItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FPGAEntitlementItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *FPGAEntitlementItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *FPGAEntitlementItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFpga

`func (o *FPGAEntitlementItem) GetFpga() string`

GetFpga returns the Fpga field if non-nil, zero value otherwise.

### GetFpgaOk

`func (o *FPGAEntitlementItem) GetFpgaOk() (*string, bool)`

GetFpgaOk returns a tuple with the Fpga field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpga

`func (o *FPGAEntitlementItem) SetFpga(v string)`

SetFpga sets Fpga field to given value.


### GetName

`func (o *FPGAEntitlementItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FPGAEntitlementItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FPGAEntitlementItem) SetName(v string)`

SetName sets Name field to given value.


### GetUsers

`func (o *FPGAEntitlementItem) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *FPGAEntitlementItem) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *FPGAEntitlementItem) SetUsers(v []string)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


