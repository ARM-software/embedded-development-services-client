/*
 * Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

/*
Solar API

This API provides a RESTful interface to all the Solar services e.g. looking for boards, building projects, etc. - This API uses Hypermedia as the Engine of Application State (HATEOAS) to drive the discovery and provide   affordances. - Discovery is possible by following links from the well known root resource. While this specification lists   all supported endpoints, it is only recommended that these are hard coded into a client if code generation is   being used. Otherwise, it is recommended that the discovery mechanisms present in the resources (affordances)   are used exclusively. - Affordances are links which indicate whether an action is currently possible, this is significantly different from   whether the service supports an action in general. This specification defines what actions could be possible,   but only by checking the affordances returned by the API in the returned resources, can a client determine whether   this action is currently possible or available for the current user. For example:   - An operation to modify a resource could be defined in this specification, but the user may lack the appropriate     privileges. In that situation, the affordance link would not be present in the resource when read. Therefore,     the client can infer that it is not possible to edit this resource and present appropriate information to the     user.   - An operation to delete a resource could be defined and be possible in some circumstances. The specification     describes that the delete is supported and how to use it, but the affordance describes whether it is currently     possible. The logic in the API may dictate that if the resource was in use (perhaps it is a running job or used     by another resource), then it will not be possible to delete that resource as it would result in a conflicted     state. - It is strongly encouraged that affordances are used by all clients, even those using code generation. This has the   ability to both improve robustness and the user experience by decoupling the client and server. For example, if for   some reason the criteria for deleting a resource changes, the logic is only implemented in the server and there is   no need to update the logic in the client as it is driven by the affordances. - The format used for the resources is the Hypertext Application Language (HAL), which includes the definition   of links and embedded resources. 

API version: 1.0.0
Contact: support@arm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the MessageObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageObject{}

// MessageObject struct for MessageObject
type MessageObject struct {
	// The time and date at which the message was created.
	Ctime *time.Time `json:"ctime,omitempty"`
	// The text of the build message.
	Message string `json:"message"`
	// Severity of the message.
	Severity *string `json:"severity,omitempty"`
	// The source of the message, typically this could be the build service itself or some component of the build tools, such as the compiler or linker.
	Source *string `json:"source,omitempty"`
}

// NewMessageObject instantiates a new MessageObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageObject(message string) *MessageObject {
	this := MessageObject{}
	this.Message = message
	return &this
}

// NewMessageObjectWithDefaults instantiates a new MessageObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageObjectWithDefaults() *MessageObject {
	this := MessageObject{}
	return &this
}

// GetCtime returns the Ctime field value if set, zero value otherwise.
func (o *MessageObject) GetCtime() time.Time {
	if o == nil || IsNil(o.Ctime) {
		var ret time.Time
		return ret
	}
	return *o.Ctime
}

// GetCtimeOk returns a tuple with the Ctime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageObject) GetCtimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ctime) {
		return nil, false
	}
	return o.Ctime, true
}

// HasCtime returns a boolean if a field has been set.
func (o *MessageObject) HasCtime() bool {
	if o != nil && !IsNil(o.Ctime) {
		return true
	}

	return false
}

// SetCtime gets a reference to the given time.Time and assigns it to the Ctime field.
func (o *MessageObject) SetCtime(v time.Time) {
	o.Ctime = &v
}

// GetMessage returns the Message field value
func (o *MessageObject) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *MessageObject) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *MessageObject) SetMessage(v string) {
	o.Message = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *MessageObject) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageObject) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *MessageObject) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *MessageObject) SetSeverity(v string) {
	o.Severity = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *MessageObject) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageObject) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MessageObject) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *MessageObject) SetSource(v string) {
	o.Source = &v
}

func (o MessageObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: ctime is readOnly
	// skip: message is readOnly
	// skip: severity is readOnly
	// skip: source is readOnly
	return toSerialize, nil
}

type NullableMessageObject struct {
	value *MessageObject
	isSet bool
}

func (v NullableMessageObject) Get() *MessageObject {
	return v.value
}

func (v *NullableMessageObject) Set(val *MessageObject) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageObject) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageObject(val *MessageObject) *NullableMessageObject {
	return &NullableMessageObject{value: val, isSet: true}
}

func (v NullableMessageObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


