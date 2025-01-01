<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# GenericWorkerItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableGenericWorkerItemLinks**](GenericWorkerItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Deprecated** | **bool** | True if this worker is scheduled to be removed from the service. | [readonly] 
**DeprecationInfo** | Pointer to [**DeprecationInfo**](DeprecationInfo.md) |  | [optional] 
**Name** | **string** | Unique ID of this worker | [readonly] 
**PrimaryToolsType** | **string** | Generic worker binary | [readonly] 
**PrimaryToolsVersion** | **string** | Version of the primary tools in use, as specified by the tools creators. | [readonly] 
**Title** | **string** | Human readable name of the worker. | [readonly] 
**ToolchainType** | Pointer to **string** | Generic worker required toolchain | [optional] [readonly] 
**ToolchainVersion** | Pointer to **string** | The version of the tool chain in use as specified by the tool chain supplier. | [optional] [readonly] 

## Methods

### NewGenericWorkerItem

`func NewGenericWorkerItem(links NullableGenericWorkerItemLinks, metadata NullableCommonMetadata, deprecated bool, name string, primaryToolsType string, primaryToolsVersion string, title string, ) *GenericWorkerItem`

NewGenericWorkerItem instantiates a new GenericWorkerItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericWorkerItemWithDefaults

`func NewGenericWorkerItemWithDefaults() *GenericWorkerItem`

NewGenericWorkerItemWithDefaults instantiates a new GenericWorkerItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GenericWorkerItem) GetLinks() GenericWorkerItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GenericWorkerItem) GetLinksOk() (*GenericWorkerItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GenericWorkerItem) SetLinks(v GenericWorkerItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *GenericWorkerItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *GenericWorkerItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *GenericWorkerItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GenericWorkerItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GenericWorkerItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *GenericWorkerItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GenericWorkerItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDeprecated

`func (o *GenericWorkerItem) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *GenericWorkerItem) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *GenericWorkerItem) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.


### GetDeprecationInfo

`func (o *GenericWorkerItem) GetDeprecationInfo() DeprecationInfo`

GetDeprecationInfo returns the DeprecationInfo field if non-nil, zero value otherwise.

### GetDeprecationInfoOk

`func (o *GenericWorkerItem) GetDeprecationInfoOk() (*DeprecationInfo, bool)`

GetDeprecationInfoOk returns a tuple with the DeprecationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationInfo

`func (o *GenericWorkerItem) SetDeprecationInfo(v DeprecationInfo)`

SetDeprecationInfo sets DeprecationInfo field to given value.

### HasDeprecationInfo

`func (o *GenericWorkerItem) HasDeprecationInfo() bool`

HasDeprecationInfo returns a boolean if a field has been set.

### GetName

`func (o *GenericWorkerItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenericWorkerItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenericWorkerItem) SetName(v string)`

SetName sets Name field to given value.


### GetPrimaryToolsType

`func (o *GenericWorkerItem) GetPrimaryToolsType() string`

GetPrimaryToolsType returns the PrimaryToolsType field if non-nil, zero value otherwise.

### GetPrimaryToolsTypeOk

`func (o *GenericWorkerItem) GetPrimaryToolsTypeOk() (*string, bool)`

GetPrimaryToolsTypeOk returns a tuple with the PrimaryToolsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryToolsType

`func (o *GenericWorkerItem) SetPrimaryToolsType(v string)`

SetPrimaryToolsType sets PrimaryToolsType field to given value.


### GetPrimaryToolsVersion

`func (o *GenericWorkerItem) GetPrimaryToolsVersion() string`

GetPrimaryToolsVersion returns the PrimaryToolsVersion field if non-nil, zero value otherwise.

### GetPrimaryToolsVersionOk

`func (o *GenericWorkerItem) GetPrimaryToolsVersionOk() (*string, bool)`

GetPrimaryToolsVersionOk returns a tuple with the PrimaryToolsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryToolsVersion

`func (o *GenericWorkerItem) SetPrimaryToolsVersion(v string)`

SetPrimaryToolsVersion sets PrimaryToolsVersion field to given value.


### GetTitle

`func (o *GenericWorkerItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GenericWorkerItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GenericWorkerItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetToolchainType

`func (o *GenericWorkerItem) GetToolchainType() string`

GetToolchainType returns the ToolchainType field if non-nil, zero value otherwise.

### GetToolchainTypeOk

`func (o *GenericWorkerItem) GetToolchainTypeOk() (*string, bool)`

GetToolchainTypeOk returns a tuple with the ToolchainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolchainType

`func (o *GenericWorkerItem) SetToolchainType(v string)`

SetToolchainType sets ToolchainType field to given value.

### HasToolchainType

`func (o *GenericWorkerItem) HasToolchainType() bool`

HasToolchainType returns a boolean if a field has been set.

### GetToolchainVersion

`func (o *GenericWorkerItem) GetToolchainVersion() string`

GetToolchainVersion returns the ToolchainVersion field if non-nil, zero value otherwise.

### GetToolchainVersionOk

`func (o *GenericWorkerItem) GetToolchainVersionOk() (*string, bool)`

GetToolchainVersionOk returns a tuple with the ToolchainVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolchainVersion

`func (o *GenericWorkerItem) SetToolchainVersion(v string)`

SetToolchainVersion sets ToolchainVersion field to given value.

### HasToolchainVersion

`func (o *GenericWorkerItem) HasToolchainVersion() bool`

HasToolchainVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


