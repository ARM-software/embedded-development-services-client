<!--
Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# WorkspaceRepositoryContentManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableWorkspaceRepositoryContentManagerLinks**](WorkspaceRepositoryContentManagerLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**FetchDepth** | Pointer to **NullableInt32** | Number of commits to fetch. 0 indicates all history for all branches and tags | [optional] [default to 1]
**Name** | **string** | ID of the repository. | [readonly] 
**Reference** | **NullableString** | The branch, tag or SHA to checkout | 
**RepositoryUrl** | **string** | repository URL to clone it from | 
**Scm** | **string** | Source Control Management system to use | 
**Submodules** | Pointer to **bool** | Whether to checkout submodules: &#x60;true&#x60; to checkout submodules recursively | [optional] [default to true]
**Title** | Pointer to **NullableString** | Optional human readable name of the repository. | [optional] 

## Methods

### NewWorkspaceRepositoryContentManager

`func NewWorkspaceRepositoryContentManager(links NullableWorkspaceRepositoryContentManagerLinks, metadata NullableCommonMetadata, name string, reference NullableString, repositoryUrl string, scm string, ) *WorkspaceRepositoryContentManager`

NewWorkspaceRepositoryContentManager instantiates a new WorkspaceRepositoryContentManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceRepositoryContentManagerWithDefaults

`func NewWorkspaceRepositoryContentManagerWithDefaults() *WorkspaceRepositoryContentManager`

NewWorkspaceRepositoryContentManagerWithDefaults instantiates a new WorkspaceRepositoryContentManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *WorkspaceRepositoryContentManager) GetLinks() WorkspaceRepositoryContentManagerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkspaceRepositoryContentManager) GetLinksOk() (*WorkspaceRepositoryContentManagerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkspaceRepositoryContentManager) SetLinks(v WorkspaceRepositoryContentManagerLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *WorkspaceRepositoryContentManager) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *WorkspaceRepositoryContentManager) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *WorkspaceRepositoryContentManager) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WorkspaceRepositoryContentManager) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WorkspaceRepositoryContentManager) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *WorkspaceRepositoryContentManager) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WorkspaceRepositoryContentManager) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFetchDepth

`func (o *WorkspaceRepositoryContentManager) GetFetchDepth() int32`

GetFetchDepth returns the FetchDepth field if non-nil, zero value otherwise.

### GetFetchDepthOk

`func (o *WorkspaceRepositoryContentManager) GetFetchDepthOk() (*int32, bool)`

GetFetchDepthOk returns a tuple with the FetchDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchDepth

`func (o *WorkspaceRepositoryContentManager) SetFetchDepth(v int32)`

SetFetchDepth sets FetchDepth field to given value.

### HasFetchDepth

`func (o *WorkspaceRepositoryContentManager) HasFetchDepth() bool`

HasFetchDepth returns a boolean if a field has been set.

### SetFetchDepthNil

`func (o *WorkspaceRepositoryContentManager) SetFetchDepthNil(b bool)`

 SetFetchDepthNil sets the value for FetchDepth to be an explicit nil

### UnsetFetchDepth
`func (o *WorkspaceRepositoryContentManager) UnsetFetchDepth()`

UnsetFetchDepth ensures that no value is present for FetchDepth, not even an explicit nil
### GetName

`func (o *WorkspaceRepositoryContentManager) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceRepositoryContentManager) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceRepositoryContentManager) SetName(v string)`

SetName sets Name field to given value.


### GetReference

`func (o *WorkspaceRepositoryContentManager) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *WorkspaceRepositoryContentManager) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *WorkspaceRepositoryContentManager) SetReference(v string)`

SetReference sets Reference field to given value.


### SetReferenceNil

`func (o *WorkspaceRepositoryContentManager) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *WorkspaceRepositoryContentManager) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetRepositoryUrl

`func (o *WorkspaceRepositoryContentManager) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *WorkspaceRepositoryContentManager) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *WorkspaceRepositoryContentManager) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.


### GetScm

`func (o *WorkspaceRepositoryContentManager) GetScm() string`

GetScm returns the Scm field if non-nil, zero value otherwise.

### GetScmOk

`func (o *WorkspaceRepositoryContentManager) GetScmOk() (*string, bool)`

GetScmOk returns a tuple with the Scm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScm

`func (o *WorkspaceRepositoryContentManager) SetScm(v string)`

SetScm sets Scm field to given value.


### GetSubmodules

`func (o *WorkspaceRepositoryContentManager) GetSubmodules() bool`

GetSubmodules returns the Submodules field if non-nil, zero value otherwise.

### GetSubmodulesOk

`func (o *WorkspaceRepositoryContentManager) GetSubmodulesOk() (*bool, bool)`

GetSubmodulesOk returns a tuple with the Submodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmodules

`func (o *WorkspaceRepositoryContentManager) SetSubmodules(v bool)`

SetSubmodules sets Submodules field to given value.

### HasSubmodules

`func (o *WorkspaceRepositoryContentManager) HasSubmodules() bool`

HasSubmodules returns a boolean if a field has been set.

### GetTitle

`func (o *WorkspaceRepositoryContentManager) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkspaceRepositoryContentManager) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkspaceRepositoryContentManager) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkspaceRepositoryContentManager) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WorkspaceRepositoryContentManager) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WorkspaceRepositoryContentManager) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


