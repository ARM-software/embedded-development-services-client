<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAConnectionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**FPGAConnectionItemLinks**](FPGAConnectionItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Count** | Pointer to **int32** | Number of current connections using this channel. | [optional] 
**Name** | **string** | The identifier of the connection. | 
**Ready** | Pointer to **bool** | Whether the connection is ready to use. | [optional] 
**Status** | Pointer to **string** | Status of the connection | [optional] 
**Title** | Pointer to **string** | Human readable description of the connection | [optional] 

## Methods

### NewFPGAConnectionItem

`func NewFPGAConnectionItem(links FPGAConnectionItemLinks, metadata NullableCommonMetadata, name string, ) *FPGAConnectionItem`

NewFPGAConnectionItem instantiates a new FPGAConnectionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAConnectionItemWithDefaults

`func NewFPGAConnectionItemWithDefaults() *FPGAConnectionItem`

NewFPGAConnectionItemWithDefaults instantiates a new FPGAConnectionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FPGAConnectionItem) GetLinks() FPGAConnectionItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FPGAConnectionItem) GetLinksOk() (*FPGAConnectionItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FPGAConnectionItem) SetLinks(v FPGAConnectionItemLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *FPGAConnectionItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FPGAConnectionItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FPGAConnectionItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *FPGAConnectionItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *FPGAConnectionItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCount

`func (o *FPGAConnectionItem) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FPGAConnectionItem) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FPGAConnectionItem) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FPGAConnectionItem) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetName

`func (o *FPGAConnectionItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FPGAConnectionItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FPGAConnectionItem) SetName(v string)`

SetName sets Name field to given value.


### GetReady

`func (o *FPGAConnectionItem) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *FPGAConnectionItem) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *FPGAConnectionItem) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *FPGAConnectionItem) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetStatus

`func (o *FPGAConnectionItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FPGAConnectionItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FPGAConnectionItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FPGAConnectionItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *FPGAConnectionItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FPGAConnectionItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FPGAConnectionItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FPGAConnectionItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


