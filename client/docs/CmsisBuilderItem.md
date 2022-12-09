<!--
Copyright (C) 2020-2022 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# CmsisBuilderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableCmsisBuilderItemLinks**](CmsisBuilderItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**BuildToolsType** | [**BuildToolTypes**](BuildToolTypes.md) |  | 
**BuildToolsVersion** | **string** | Version of the build tools in use, as specified by the tools creators. | [readonly] 
**Deprecated** | **bool** | True if this CMSIS Builder is scheduled to be removed from the service. | [readonly] 
**DeprecationInfo** | Pointer to [**NullableCmsisBuilderItemDeprecationInfo**](CmsisBuilderItemDeprecationInfo.md) |  | [optional] 
**Name** | **string** | Unique ID of the CMSIS builder. | [readonly] 
**Title** | **string** | Human readable name of the CMSIS builder. | [readonly] 
**ToolchainType** | [**ToolchainTypes**](ToolchainTypes.md) |  | 
**ToolchainVersion** | **string** | The version of the toolchain in use as specified by the toolchain supplier. | [readonly] 

## Methods

### NewCmsisBuilderItem

`func NewCmsisBuilderItem(links NullableCmsisBuilderItemLinks, metadata NullableCommonMetadata, buildToolsType BuildToolTypes, buildToolsVersion string, deprecated bool, name string, title string, toolchainType ToolchainTypes, toolchainVersion string, ) *CmsisBuilderItem`

NewCmsisBuilderItem instantiates a new CmsisBuilderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmsisBuilderItemWithDefaults

`func NewCmsisBuilderItemWithDefaults() *CmsisBuilderItem`

NewCmsisBuilderItemWithDefaults instantiates a new CmsisBuilderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CmsisBuilderItem) GetLinks() CmsisBuilderItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CmsisBuilderItem) GetLinksOk() (*CmsisBuilderItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CmsisBuilderItem) SetLinks(v CmsisBuilderItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *CmsisBuilderItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *CmsisBuilderItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *CmsisBuilderItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CmsisBuilderItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CmsisBuilderItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *CmsisBuilderItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CmsisBuilderItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetBuildToolsType

`func (o *CmsisBuilderItem) GetBuildToolsType() BuildToolTypes`

GetBuildToolsType returns the BuildToolsType field if non-nil, zero value otherwise.

### GetBuildToolsTypeOk

`func (o *CmsisBuilderItem) GetBuildToolsTypeOk() (*BuildToolTypes, bool)`

GetBuildToolsTypeOk returns a tuple with the BuildToolsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildToolsType

`func (o *CmsisBuilderItem) SetBuildToolsType(v BuildToolTypes)`

SetBuildToolsType sets BuildToolsType field to given value.


### GetBuildToolsVersion

`func (o *CmsisBuilderItem) GetBuildToolsVersion() string`

GetBuildToolsVersion returns the BuildToolsVersion field if non-nil, zero value otherwise.

### GetBuildToolsVersionOk

`func (o *CmsisBuilderItem) GetBuildToolsVersionOk() (*string, bool)`

GetBuildToolsVersionOk returns a tuple with the BuildToolsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildToolsVersion

`func (o *CmsisBuilderItem) SetBuildToolsVersion(v string)`

SetBuildToolsVersion sets BuildToolsVersion field to given value.


### GetDeprecated

`func (o *CmsisBuilderItem) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *CmsisBuilderItem) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *CmsisBuilderItem) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.


### GetDeprecationInfo

`func (o *CmsisBuilderItem) GetDeprecationInfo() CmsisBuilderItemDeprecationInfo`

GetDeprecationInfo returns the DeprecationInfo field if non-nil, zero value otherwise.

### GetDeprecationInfoOk

`func (o *CmsisBuilderItem) GetDeprecationInfoOk() (*CmsisBuilderItemDeprecationInfo, bool)`

GetDeprecationInfoOk returns a tuple with the DeprecationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationInfo

`func (o *CmsisBuilderItem) SetDeprecationInfo(v CmsisBuilderItemDeprecationInfo)`

SetDeprecationInfo sets DeprecationInfo field to given value.

### HasDeprecationInfo

`func (o *CmsisBuilderItem) HasDeprecationInfo() bool`

HasDeprecationInfo returns a boolean if a field has been set.

### SetDeprecationInfoNil

`func (o *CmsisBuilderItem) SetDeprecationInfoNil(b bool)`

 SetDeprecationInfoNil sets the value for DeprecationInfo to be an explicit nil

### UnsetDeprecationInfo
`func (o *CmsisBuilderItem) UnsetDeprecationInfo()`

UnsetDeprecationInfo ensures that no value is present for DeprecationInfo, not even an explicit nil
### GetName

`func (o *CmsisBuilderItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CmsisBuilderItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CmsisBuilderItem) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CmsisBuilderItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CmsisBuilderItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CmsisBuilderItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetToolchainType

`func (o *CmsisBuilderItem) GetToolchainType() ToolchainTypes`

GetToolchainType returns the ToolchainType field if non-nil, zero value otherwise.

### GetToolchainTypeOk

`func (o *CmsisBuilderItem) GetToolchainTypeOk() (*ToolchainTypes, bool)`

GetToolchainTypeOk returns a tuple with the ToolchainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolchainType

`func (o *CmsisBuilderItem) SetToolchainType(v ToolchainTypes)`

SetToolchainType sets ToolchainType field to given value.


### GetToolchainVersion

`func (o *CmsisBuilderItem) GetToolchainVersion() string`

GetToolchainVersion returns the ToolchainVersion field if non-nil, zero value otherwise.

### GetToolchainVersionOk

`func (o *CmsisBuilderItem) GetToolchainVersionOk() (*string, bool)`

GetToolchainVersionOk returns a tuple with the ToolchainVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolchainVersion

`func (o *CmsisBuilderItem) SetToolchainVersion(v string)`

SetToolchainVersion sets ToolchainVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


