<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# IntellisenseMessageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableHalFeedLinks**](HalFeedLinks.md) |  | 
**Metadata** | [**NullablePagingMetadata**](PagingMetadata.md) |  | 
**Messages** | [**[]MessageObject**](MessageObject.md) | A list of build messages. | [readonly] 
**Name** | **string** | ID of the Build Notification item. | [readonly] 
**Title** | Pointer to **string** | Human readable name of the Build Notification item. | [optional] 

## Methods

### NewIntellisenseMessageItem

`func NewIntellisenseMessageItem(links NullableHalFeedLinks, metadata NullablePagingMetadata, messages []MessageObject, name string, ) *IntellisenseMessageItem`

NewIntellisenseMessageItem instantiates a new IntellisenseMessageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntellisenseMessageItemWithDefaults

`func NewIntellisenseMessageItemWithDefaults() *IntellisenseMessageItem`

NewIntellisenseMessageItemWithDefaults instantiates a new IntellisenseMessageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *IntellisenseMessageItem) GetLinks() HalFeedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IntellisenseMessageItem) GetLinksOk() (*HalFeedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IntellisenseMessageItem) SetLinks(v HalFeedLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *IntellisenseMessageItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *IntellisenseMessageItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *IntellisenseMessageItem) GetMetadata() PagingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IntellisenseMessageItem) GetMetadataOk() (*PagingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IntellisenseMessageItem) SetMetadata(v PagingMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *IntellisenseMessageItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *IntellisenseMessageItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMessages

`func (o *IntellisenseMessageItem) GetMessages() []MessageObject`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *IntellisenseMessageItem) GetMessagesOk() (*[]MessageObject, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *IntellisenseMessageItem) SetMessages(v []MessageObject)`

SetMessages sets Messages field to given value.


### SetMessagesNil

`func (o *IntellisenseMessageItem) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *IntellisenseMessageItem) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetName

`func (o *IntellisenseMessageItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntellisenseMessageItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntellisenseMessageItem) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *IntellisenseMessageItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IntellisenseMessageItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IntellisenseMessageItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IntellisenseMessageItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


