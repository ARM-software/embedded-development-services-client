<!--
Copyright (C) 2020-2024 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# PATItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TTL** | **int64** | The TTL (time to live in seconds) describing how long the personal access token will be alive for. | 
**Links** | [**PATItemLinks**](PATItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**LastUsed** | **time.Time** | UTC date and time when the token was last used. | [readonly] 
**Name** | **string** | Unique ID of the personal access token. | [readonly] 
**Secret** | Pointer to **NullableString** | The personal access token. | [optional] [readonly] 
**Title** | **string** | Human readable name of the personal access token. | 

## Methods

### NewPATItem

`func NewPATItem(tTL int64, links PATItemLinks, metadata NullableCommonMetadata, lastUsed time.Time, name string, title string, ) *PATItem`

NewPATItem instantiates a new PATItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATItemWithDefaults

`func NewPATItemWithDefaults() *PATItem`

NewPATItemWithDefaults instantiates a new PATItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTTL

`func (o *PATItem) GetTTL() int64`

GetTTL returns the TTL field if non-nil, zero value otherwise.

### GetTTLOk

`func (o *PATItem) GetTTLOk() (*int64, bool)`

GetTTLOk returns a tuple with the TTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTTL

`func (o *PATItem) SetTTL(v int64)`

SetTTL sets TTL field to given value.


### GetLinks

`func (o *PATItem) GetLinks() PATItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PATItem) GetLinksOk() (*PATItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PATItem) SetLinks(v PATItemLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *PATItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *PATItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetLastUsed

`func (o *PATItem) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *PATItem) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *PATItem) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.


### GetName

`func (o *PATItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATItem) SetName(v string)`

SetName sets Name field to given value.


### GetSecret

`func (o *PATItem) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PATItem) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PATItem) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PATItem) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *PATItem) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *PATItem) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetTitle

`func (o *PATItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PATItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PATItem) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


