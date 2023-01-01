<!--
Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# MessageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ctime** | Pointer to **time.Time** | The time and date at which the message was created. | [optional] [readonly] 
**Message** | **string** | The text of the build message. | [readonly] 
**Severity** | Pointer to **string** | Severity of the message. | [optional] [readonly] 
**Source** | Pointer to **string** | The source of the message, typically this could be the build service itself or some component of the build tools, such as the compiler or linker. | [optional] [readonly] 

## Methods

### NewMessageObject

`func NewMessageObject(message string, ) *MessageObject`

NewMessageObject instantiates a new MessageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageObjectWithDefaults

`func NewMessageObjectWithDefaults() *MessageObject`

NewMessageObjectWithDefaults instantiates a new MessageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCtime

`func (o *MessageObject) GetCtime() time.Time`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *MessageObject) GetCtimeOk() (*time.Time, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *MessageObject) SetCtime(v time.Time)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *MessageObject) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetMessage

`func (o *MessageObject) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MessageObject) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MessageObject) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSeverity

`func (o *MessageObject) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *MessageObject) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *MessageObject) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *MessageObject) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *MessageObject) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MessageObject) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MessageObject) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MessageObject) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


