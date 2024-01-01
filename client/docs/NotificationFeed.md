<!--
Copyright (C) 2020-2024 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# NotificationFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableHalFeedLinks**](HalFeedLinks.md) |  | 
**Metadata** | [**NullablePagingMetadata**](PagingMetadata.md) |  | 
**Messages** | [**[]NotificationMessageObject**](NotificationMessageObject.md) | A list of messages. | [readonly] 
**Name** | **string** | ID of the notification item. | [readonly] 
**Title** | Pointer to **string** | Human readable name of the notification item. | [optional] 

## Methods

### NewNotificationFeed

`func NewNotificationFeed(links NullableHalFeedLinks, metadata NullablePagingMetadata, messages []NotificationMessageObject, name string, ) *NotificationFeed`

NewNotificationFeed instantiates a new NotificationFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationFeedWithDefaults

`func NewNotificationFeedWithDefaults() *NotificationFeed`

NewNotificationFeedWithDefaults instantiates a new NotificationFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *NotificationFeed) GetLinks() HalFeedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NotificationFeed) GetLinksOk() (*HalFeedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NotificationFeed) SetLinks(v HalFeedLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *NotificationFeed) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *NotificationFeed) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *NotificationFeed) GetMetadata() PagingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationFeed) GetMetadataOk() (*PagingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationFeed) SetMetadata(v PagingMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *NotificationFeed) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *NotificationFeed) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMessages

`func (o *NotificationFeed) GetMessages() []NotificationMessageObject`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *NotificationFeed) GetMessagesOk() (*[]NotificationMessageObject, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *NotificationFeed) SetMessages(v []NotificationMessageObject)`

SetMessages sets Messages field to given value.


### SetMessagesNil

`func (o *NotificationFeed) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *NotificationFeed) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetName

`func (o *NotificationFeed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationFeed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationFeed) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NotificationFeed) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NotificationFeed) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NotificationFeed) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NotificationFeed) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


