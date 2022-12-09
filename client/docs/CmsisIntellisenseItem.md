<!--
Copyright (C) 2020-2022 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# CmsisIntellisenseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableCmsisIntellisenseItemLinks**](CmsisIntellisenseItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**BuildToolsType** | **string** | The build system that will be used to build the project. | [readonly] 
**BuildToolsVersion** | **string** | Version of the build tools in use, as specified by the tools creators. | [readonly] 
**Deprecated** | **bool** | True if this CMSIS Builder is scheduled to be removed from the service. | [readonly] 
**DeprecationInfo** | Pointer to [**NullableCmsisIntellisenseItemDeprecationInfo**](CmsisIntellisenseItemDeprecationInfo.md) |  | [optional] 
**Name** | **string** | Unique ID of the CMSIS builder. | [readonly] 
**Title** | **string** | Human readable name of the CMSIS builder. | [readonly] 
**ToolchainType** | [**ToolchainTypes**](ToolchainTypes.md) |  | 
**ToolchainVersion** | **string** | The version of the toolchain in use as specified by the toolchain supplier. | [readonly] 

## Methods

### NewCmsisIntellisenseItem

`func NewCmsisIntellisenseItem(links NullableCmsisIntellisenseItemLinks, metadata NullableCommonMetadata, buildToolsType string, buildToolsVersion string, deprecated bool, name string, title string, toolchainType ToolchainTypes, toolchainVersion string, ) *CmsisIntellisenseItem`

NewCmsisIntellisenseItem instantiates a new CmsisIntellisenseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmsisIntellisenseItemWithDefaults

`func NewCmsisIntellisenseItemWithDefaults() *CmsisIntellisenseItem`

NewCmsisIntellisenseItemWithDefaults instantiates a new CmsisIntellisenseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CmsisIntellisenseItem) GetLinks() CmsisIntellisenseItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CmsisIntellisenseItem) GetLinksOk() (*CmsisIntellisenseItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CmsisIntellisenseItem) SetLinks(v CmsisIntellisenseItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *CmsisIntellisenseItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *CmsisIntellisenseItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *CmsisIntellisenseItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CmsisIntellisenseItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CmsisIntellisenseItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *CmsisIntellisenseItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CmsisIntellisenseItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetBuildToolsType

`func (o *CmsisIntellisenseItem) GetBuildToolsType() string`

GetBuildToolsType returns the BuildToolsType field if non-nil, zero value otherwise.

### GetBuildToolsTypeOk

`func (o *CmsisIntellisenseItem) GetBuildToolsTypeOk() (*string, bool)`

GetBuildToolsTypeOk returns a tuple with the BuildToolsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildToolsType

`func (o *CmsisIntellisenseItem) SetBuildToolsType(v string)`

SetBuildToolsType sets BuildToolsType field to given value.


### GetBuildToolsVersion

`func (o *CmsisIntellisenseItem) GetBuildToolsVersion() string`

GetBuildToolsVersion returns the BuildToolsVersion field if non-nil, zero value otherwise.

### GetBuildToolsVersionOk

`func (o *CmsisIntellisenseItem) GetBuildToolsVersionOk() (*string, bool)`

GetBuildToolsVersionOk returns a tuple with the BuildToolsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildToolsVersion

`func (o *CmsisIntellisenseItem) SetBuildToolsVersion(v string)`

SetBuildToolsVersion sets BuildToolsVersion field to given value.


### GetDeprecated

`func (o *CmsisIntellisenseItem) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *CmsisIntellisenseItem) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *CmsisIntellisenseItem) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.


### GetDeprecationInfo

`func (o *CmsisIntellisenseItem) GetDeprecationInfo() CmsisIntellisenseItemDeprecationInfo`

GetDeprecationInfo returns the DeprecationInfo field if non-nil, zero value otherwise.

### GetDeprecationInfoOk

`func (o *CmsisIntellisenseItem) GetDeprecationInfoOk() (*CmsisIntellisenseItemDeprecationInfo, bool)`

GetDeprecationInfoOk returns a tuple with the DeprecationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationInfo

`func (o *CmsisIntellisenseItem) SetDeprecationInfo(v CmsisIntellisenseItemDeprecationInfo)`

SetDeprecationInfo sets DeprecationInfo field to given value.

### HasDeprecationInfo

`func (o *CmsisIntellisenseItem) HasDeprecationInfo() bool`

HasDeprecationInfo returns a boolean if a field has been set.

### SetDeprecationInfoNil

`func (o *CmsisIntellisenseItem) SetDeprecationInfoNil(b bool)`

 SetDeprecationInfoNil sets the value for DeprecationInfo to be an explicit nil

### UnsetDeprecationInfo
`func (o *CmsisIntellisenseItem) UnsetDeprecationInfo()`

UnsetDeprecationInfo ensures that no value is present for DeprecationInfo, not even an explicit nil
### GetName

`func (o *CmsisIntellisenseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CmsisIntellisenseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CmsisIntellisenseItem) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CmsisIntellisenseItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CmsisIntellisenseItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CmsisIntellisenseItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetToolchainType

`func (o *CmsisIntellisenseItem) GetToolchainType() ToolchainTypes`

GetToolchainType returns the ToolchainType field if non-nil, zero value otherwise.

### GetToolchainTypeOk

`func (o *CmsisIntellisenseItem) GetToolchainTypeOk() (*ToolchainTypes, bool)`

GetToolchainTypeOk returns a tuple with the ToolchainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolchainType

`func (o *CmsisIntellisenseItem) SetToolchainType(v ToolchainTypes)`

SetToolchainType sets ToolchainType field to given value.


### GetToolchainVersion

`func (o *CmsisIntellisenseItem) GetToolchainVersion() string`

GetToolchainVersion returns the ToolchainVersion field if non-nil, zero value otherwise.

### GetToolchainVersionOk

`func (o *CmsisIntellisenseItem) GetToolchainVersionOk() (*string, bool)`

GetToolchainVersionOk returns a tuple with the ToolchainVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolchainVersion

`func (o *CmsisIntellisenseItem) SetToolchainVersion(v string)`

SetToolchainVersion sets ToolchainVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


