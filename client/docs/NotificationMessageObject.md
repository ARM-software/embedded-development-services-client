<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# NotificationMessageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ctime** | Pointer to **time.Time** | The time and date at which the message was created. | [optional] [readonly] 
**Message** | **string** | The text of the message. | [readonly] 
**Severity** | Pointer to **string** | Severity of the message. | [optional] [readonly] 
**Source** | Pointer to **string** | The source of the message, typically this could be the name of the service it was originated from. | [optional] [readonly] 

## Methods

### NewNotificationMessageObject

`func NewNotificationMessageObject(message string, ) *NotificationMessageObject`

NewNotificationMessageObject instantiates a new NotificationMessageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationMessageObjectWithDefaults

`func NewNotificationMessageObjectWithDefaults() *NotificationMessageObject`

NewNotificationMessageObjectWithDefaults instantiates a new NotificationMessageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCtime

`func (o *NotificationMessageObject) GetCtime() time.Time`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *NotificationMessageObject) GetCtimeOk() (*time.Time, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *NotificationMessageObject) SetCtime(v time.Time)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *NotificationMessageObject) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetMessage

`func (o *NotificationMessageObject) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationMessageObject) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationMessageObject) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSeverity

`func (o *NotificationMessageObject) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NotificationMessageObject) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NotificationMessageObject) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *NotificationMessageObject) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *NotificationMessageObject) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NotificationMessageObject) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NotificationMessageObject) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *NotificationMessageObject) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


