<!--
Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# ArtefactManagerItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableArtefactManagerLinks**](ArtefactManagerLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Category** | **NullableString** | Category the artefact belongs to. | 
**ContentMediaType** | **string** | Media type of the artefact according to https://www.iana.org/assignments/media-types/media-types.xhtml | 
**Description** | **string** | Description of what the artefact is. | 
**Hash** | **NullableString** | Hash of the artefact (sha256) | 
**Mandatory** | Pointer to **NullableBool** | Whether the artefact is required or not for a workflow. | [optional] 
**MaxSize** | **int64** | Maximum size in bytes accepted for this artefact. | 
**Name** | **string** | ID of the artefact as set by the system. | 
**Size** | Pointer to **int64** | Size in bytes of this artefact. | [optional] 
**Title** | Pointer to **NullableString** | Optional human readable name of the artefact. | [optional] 

## Methods

### NewArtefactManagerItem

`func NewArtefactManagerItem(links NullableArtefactManagerLinks, metadata NullableCommonMetadata, category NullableString, contentMediaType string, description string, hash NullableString, maxSize int64, name string, ) *ArtefactManagerItem`

NewArtefactManagerItem instantiates a new ArtefactManagerItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtefactManagerItemWithDefaults

`func NewArtefactManagerItemWithDefaults() *ArtefactManagerItem`

NewArtefactManagerItemWithDefaults instantiates a new ArtefactManagerItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ArtefactManagerItem) GetLinks() ArtefactManagerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ArtefactManagerItem) GetLinksOk() (*ArtefactManagerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ArtefactManagerItem) SetLinks(v ArtefactManagerLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *ArtefactManagerItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ArtefactManagerItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *ArtefactManagerItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArtefactManagerItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArtefactManagerItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *ArtefactManagerItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ArtefactManagerItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCategory

`func (o *ArtefactManagerItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ArtefactManagerItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ArtefactManagerItem) SetCategory(v string)`

SetCategory sets Category field to given value.


### SetCategoryNil

`func (o *ArtefactManagerItem) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ArtefactManagerItem) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetContentMediaType

`func (o *ArtefactManagerItem) GetContentMediaType() string`

GetContentMediaType returns the ContentMediaType field if non-nil, zero value otherwise.

### GetContentMediaTypeOk

`func (o *ArtefactManagerItem) GetContentMediaTypeOk() (*string, bool)`

GetContentMediaTypeOk returns a tuple with the ContentMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentMediaType

`func (o *ArtefactManagerItem) SetContentMediaType(v string)`

SetContentMediaType sets ContentMediaType field to given value.


### GetDescription

`func (o *ArtefactManagerItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtefactManagerItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtefactManagerItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHash

`func (o *ArtefactManagerItem) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ArtefactManagerItem) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ArtefactManagerItem) SetHash(v string)`

SetHash sets Hash field to given value.


### SetHashNil

`func (o *ArtefactManagerItem) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *ArtefactManagerItem) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetMandatory

`func (o *ArtefactManagerItem) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *ArtefactManagerItem) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *ArtefactManagerItem) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *ArtefactManagerItem) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### SetMandatoryNil

`func (o *ArtefactManagerItem) SetMandatoryNil(b bool)`

 SetMandatoryNil sets the value for Mandatory to be an explicit nil

### UnsetMandatory
`func (o *ArtefactManagerItem) UnsetMandatory()`

UnsetMandatory ensures that no value is present for Mandatory, not even an explicit nil
### GetMaxSize

`func (o *ArtefactManagerItem) GetMaxSize() int64`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *ArtefactManagerItem) GetMaxSizeOk() (*int64, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *ArtefactManagerItem) SetMaxSize(v int64)`

SetMaxSize sets MaxSize field to given value.


### GetName

`func (o *ArtefactManagerItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtefactManagerItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtefactManagerItem) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *ArtefactManagerItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ArtefactManagerItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ArtefactManagerItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ArtefactManagerItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTitle

`func (o *ArtefactManagerItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArtefactManagerItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArtefactManagerItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ArtefactManagerItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ArtefactManagerItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ArtefactManagerItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


