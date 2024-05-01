<!--
Copyright (C) 2020-2024 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# PATCreationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TTL** | **int64** | The TTL (time to live in seconds) describing how long the personal access token will be alive for. | 
**Links** | [**PATItemLinks**](PATItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Name** | **string** | Unique ID of the personal access token. | [readonly] 
**Secret** | **string** | The personal access token. | [readonly] 
**Title** | **string** | Human readable name of the personal access token. | 

## Methods

### NewPATCreationItem

`func NewPATCreationItem(tTL int64, links PATItemLinks, metadata NullableCommonMetadata, name string, secret string, title string, ) *PATCreationItem`

NewPATCreationItem instantiates a new PATCreationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCreationItemWithDefaults

`func NewPATCreationItemWithDefaults() *PATCreationItem`

NewPATCreationItemWithDefaults instantiates a new PATCreationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTTL

`func (o *PATCreationItem) GetTTL() int64`

GetTTL returns the TTL field if non-nil, zero value otherwise.

### GetTTLOk

`func (o *PATCreationItem) GetTTLOk() (*int64, bool)`

GetTTLOk returns a tuple with the TTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTTL

`func (o *PATCreationItem) SetTTL(v int64)`

SetTTL sets TTL field to given value.


### GetLinks

`func (o *PATCreationItem) GetLinks() PATItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PATCreationItem) GetLinksOk() (*PATItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PATCreationItem) SetLinks(v PATItemLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *PATCreationItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCreationItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCreationItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *PATCreationItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCreationItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *PATCreationItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCreationItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCreationItem) SetName(v string)`

SetName sets Name field to given value.


### GetSecret

`func (o *PATCreationItem) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PATCreationItem) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PATCreationItem) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetTitle

`func (o *PATCreationItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PATCreationItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PATCreationItem) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


