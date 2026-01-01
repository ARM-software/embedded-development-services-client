<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# BuildMessageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableHalFeedLinks**](HalFeedLinks.md) |  | 
**Metadata** | [**NullablePagingMetadata**](PagingMetadata.md) |  | 
**Messages** | [**[]MessageObject**](MessageObject.md) | A list of build messages. | [readonly] 
**Name** | **string** | ID of the Build Notification item. | [readonly] 
**Title** | Pointer to **string** | Human readable name of the Build Notification item. | [optional] 

## Methods

### NewBuildMessageItem

`func NewBuildMessageItem(links NullableHalFeedLinks, metadata NullablePagingMetadata, messages []MessageObject, name string, ) *BuildMessageItem`

NewBuildMessageItem instantiates a new BuildMessageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildMessageItemWithDefaults

`func NewBuildMessageItemWithDefaults() *BuildMessageItem`

NewBuildMessageItemWithDefaults instantiates a new BuildMessageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *BuildMessageItem) GetLinks() HalFeedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BuildMessageItem) GetLinksOk() (*HalFeedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BuildMessageItem) SetLinks(v HalFeedLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *BuildMessageItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *BuildMessageItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *BuildMessageItem) GetMetadata() PagingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BuildMessageItem) GetMetadataOk() (*PagingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BuildMessageItem) SetMetadata(v PagingMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *BuildMessageItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BuildMessageItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMessages

`func (o *BuildMessageItem) GetMessages() []MessageObject`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *BuildMessageItem) GetMessagesOk() (*[]MessageObject, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *BuildMessageItem) SetMessages(v []MessageObject)`

SetMessages sets Messages field to given value.


### SetMessagesNil

`func (o *BuildMessageItem) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *BuildMessageItem) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetName

`func (o *BuildMessageItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuildMessageItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuildMessageItem) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *BuildMessageItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BuildMessageItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BuildMessageItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BuildMessageItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


