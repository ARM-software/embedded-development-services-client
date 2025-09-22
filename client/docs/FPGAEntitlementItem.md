<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAEntitlementItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableFPGAEntitlementItemLinks**](FPGAEntitlementItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Fpga** | Pointer to **string** | Deprecated in favour of &#x60;fpgas&#x60;. The identifier of the FPGA resource. | [optional] 
**Fpgas** | **[]string** | A list of FPGA resource identifiers that the entitlement provides access to. | 
**Name** | **string** | The unique identifier of the resource instance. | 
**Repositories** | **[]string** | A list of Repository resource identifiers that the entitlement provides access to. | 
**Repository** | Pointer to **string** | Deprecated in favour of &#x60;repositories&#x60;. The identifier of the Repository resource. | [optional] 
**Users** | **[]string** | A list of user emails permitted to access the FPGA. | 

## Methods

### NewFPGAEntitlementItem

`func NewFPGAEntitlementItem(links NullableFPGAEntitlementItemLinks, metadata NullableCommonMetadata, fpgas []string, name string, repositories []string, users []string, ) *FPGAEntitlementItem`

NewFPGAEntitlementItem instantiates a new FPGAEntitlementItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAEntitlementItemWithDefaults

`func NewFPGAEntitlementItemWithDefaults() *FPGAEntitlementItem`

NewFPGAEntitlementItemWithDefaults instantiates a new FPGAEntitlementItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FPGAEntitlementItem) GetLinks() FPGAEntitlementItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FPGAEntitlementItem) GetLinksOk() (*FPGAEntitlementItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FPGAEntitlementItem) SetLinks(v FPGAEntitlementItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *FPGAEntitlementItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *FPGAEntitlementItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
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

### HasFpga

`func (o *FPGAEntitlementItem) HasFpga() bool`

HasFpga returns a boolean if a field has been set.

### GetFpgas

`func (o *FPGAEntitlementItem) GetFpgas() []string`

GetFpgas returns the Fpgas field if non-nil, zero value otherwise.

### GetFpgasOk

`func (o *FPGAEntitlementItem) GetFpgasOk() (*[]string, bool)`

GetFpgasOk returns a tuple with the Fpgas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpgas

`func (o *FPGAEntitlementItem) SetFpgas(v []string)`

SetFpgas sets Fpgas field to given value.


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


### GetRepositories

`func (o *FPGAEntitlementItem) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *FPGAEntitlementItem) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *FPGAEntitlementItem) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.


### GetRepository

`func (o *FPGAEntitlementItem) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *FPGAEntitlementItem) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *FPGAEntitlementItem) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *FPGAEntitlementItem) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

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


