<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# WorkspaceSourceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableWorkspaceSourceItemLinks**](WorkspaceSourceItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Deprecated** | **bool** | True if this workspace source is scheduled to be no longer supported by the service. | [readonly] 
**DeprecationInfo** | Pointer to [**DeprecationInfo**](DeprecationInfo.md) |  | [optional] 
**Description** | Pointer to **string** | More details about this workspace source type. | [optional] [readonly] 
**Name** | **string** | Unique ID of the Workspace Source. | [readonly] 
**Title** | **string** | Human readable name of the Workspace source. | [readonly] 

## Methods

### NewWorkspaceSourceItem

`func NewWorkspaceSourceItem(links NullableWorkspaceSourceItemLinks, metadata NullableCommonMetadata, deprecated bool, name string, title string, ) *WorkspaceSourceItem`

NewWorkspaceSourceItem instantiates a new WorkspaceSourceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceSourceItemWithDefaults

`func NewWorkspaceSourceItemWithDefaults() *WorkspaceSourceItem`

NewWorkspaceSourceItemWithDefaults instantiates a new WorkspaceSourceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *WorkspaceSourceItem) GetLinks() WorkspaceSourceItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkspaceSourceItem) GetLinksOk() (*WorkspaceSourceItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkspaceSourceItem) SetLinks(v WorkspaceSourceItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *WorkspaceSourceItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *WorkspaceSourceItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *WorkspaceSourceItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WorkspaceSourceItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WorkspaceSourceItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *WorkspaceSourceItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WorkspaceSourceItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDeprecated

`func (o *WorkspaceSourceItem) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *WorkspaceSourceItem) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *WorkspaceSourceItem) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.


### GetDeprecationInfo

`func (o *WorkspaceSourceItem) GetDeprecationInfo() DeprecationInfo`

GetDeprecationInfo returns the DeprecationInfo field if non-nil, zero value otherwise.

### GetDeprecationInfoOk

`func (o *WorkspaceSourceItem) GetDeprecationInfoOk() (*DeprecationInfo, bool)`

GetDeprecationInfoOk returns a tuple with the DeprecationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationInfo

`func (o *WorkspaceSourceItem) SetDeprecationInfo(v DeprecationInfo)`

SetDeprecationInfo sets DeprecationInfo field to given value.

### HasDeprecationInfo

`func (o *WorkspaceSourceItem) HasDeprecationInfo() bool`

HasDeprecationInfo returns a boolean if a field has been set.

### GetDescription

`func (o *WorkspaceSourceItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspaceSourceItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspaceSourceItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkspaceSourceItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkspaceSourceItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceSourceItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceSourceItem) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *WorkspaceSourceItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkspaceSourceItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkspaceSourceItem) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


