<!--
Copyright (C) 2020-2024 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# WorkspaceDetailsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableWorkspaceDetailsItemLinks**](WorkspaceDetailsItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Name** | **string** | Unique ID of the Workspace details. | [readonly] 
**Path** | **string** | relative path of the workspace | [readonly] 
**SizeInByte** | Pointer to **int64** | size on disk (in Byte) used by this workspace | [optional] [readonly] 
**Title** | Pointer to **NullableString** | Optional human readable name of the Workspace details. | [optional] 

## Methods

### NewWorkspaceDetailsItem

`func NewWorkspaceDetailsItem(links NullableWorkspaceDetailsItemLinks, metadata NullableCommonMetadata, name string, path string, ) *WorkspaceDetailsItem`

NewWorkspaceDetailsItem instantiates a new WorkspaceDetailsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceDetailsItemWithDefaults

`func NewWorkspaceDetailsItemWithDefaults() *WorkspaceDetailsItem`

NewWorkspaceDetailsItemWithDefaults instantiates a new WorkspaceDetailsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *WorkspaceDetailsItem) GetLinks() WorkspaceDetailsItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkspaceDetailsItem) GetLinksOk() (*WorkspaceDetailsItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkspaceDetailsItem) SetLinks(v WorkspaceDetailsItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *WorkspaceDetailsItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *WorkspaceDetailsItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *WorkspaceDetailsItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WorkspaceDetailsItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WorkspaceDetailsItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *WorkspaceDetailsItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WorkspaceDetailsItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *WorkspaceDetailsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceDetailsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceDetailsItem) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *WorkspaceDetailsItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WorkspaceDetailsItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WorkspaceDetailsItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetSizeInByte

`func (o *WorkspaceDetailsItem) GetSizeInByte() int64`

GetSizeInByte returns the SizeInByte field if non-nil, zero value otherwise.

### GetSizeInByteOk

`func (o *WorkspaceDetailsItem) GetSizeInByteOk() (*int64, bool)`

GetSizeInByteOk returns a tuple with the SizeInByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInByte

`func (o *WorkspaceDetailsItem) SetSizeInByte(v int64)`

SetSizeInByte sets SizeInByte field to given value.

### HasSizeInByte

`func (o *WorkspaceDetailsItem) HasSizeInByte() bool`

HasSizeInByte returns a boolean if a field has been set.

### GetTitle

`func (o *WorkspaceDetailsItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkspaceDetailsItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkspaceDetailsItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkspaceDetailsItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WorkspaceDetailsItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WorkspaceDetailsItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


