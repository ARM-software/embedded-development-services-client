<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# WorkspaceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TTL** | Pointer to **NullableInt64** | The TTL (time to live in seconds) describing how long the workspace will be still available for. | [optional] [default to 259200]
**Links** | [**NullableWorkspaceItemLinks**](WorkspaceItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Ephemeral** | Pointer to **bool** | True when the workspace has a TTL | [optional] [readonly] 
**Name** | **string** | Unique ID of the Workspace. | [readonly] 
**Title** | Pointer to **NullableString** | Optional human readable name of the Workspace. | [optional] 

## Methods

### NewWorkspaceItem

`func NewWorkspaceItem(links NullableWorkspaceItemLinks, metadata NullableCommonMetadata, name string, ) *WorkspaceItem`

NewWorkspaceItem instantiates a new WorkspaceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceItemWithDefaults

`func NewWorkspaceItemWithDefaults() *WorkspaceItem`

NewWorkspaceItemWithDefaults instantiates a new WorkspaceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTTL

`func (o *WorkspaceItem) GetTTL() int64`

GetTTL returns the TTL field if non-nil, zero value otherwise.

### GetTTLOk

`func (o *WorkspaceItem) GetTTLOk() (*int64, bool)`

GetTTLOk returns a tuple with the TTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTTL

`func (o *WorkspaceItem) SetTTL(v int64)`

SetTTL sets TTL field to given value.

### HasTTL

`func (o *WorkspaceItem) HasTTL() bool`

HasTTL returns a boolean if a field has been set.

### SetTTLNil

`func (o *WorkspaceItem) SetTTLNil(b bool)`

 SetTTLNil sets the value for TTL to be an explicit nil

### UnsetTTL
`func (o *WorkspaceItem) UnsetTTL()`

UnsetTTL ensures that no value is present for TTL, not even an explicit nil
### GetLinks

`func (o *WorkspaceItem) GetLinks() WorkspaceItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkspaceItem) GetLinksOk() (*WorkspaceItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkspaceItem) SetLinks(v WorkspaceItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *WorkspaceItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *WorkspaceItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *WorkspaceItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WorkspaceItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WorkspaceItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *WorkspaceItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WorkspaceItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetEphemeral

`func (o *WorkspaceItem) GetEphemeral() bool`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *WorkspaceItem) GetEphemeralOk() (*bool, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *WorkspaceItem) SetEphemeral(v bool)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *WorkspaceItem) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetName

`func (o *WorkspaceItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceItem) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *WorkspaceItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkspaceItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkspaceItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkspaceItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WorkspaceItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WorkspaceItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


