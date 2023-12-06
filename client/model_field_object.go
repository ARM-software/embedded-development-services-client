/*
 * Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

/*
Solar API

This API provides a RESTful interface to all the Solar services e.g. looking for boards, building projects, etc. - This API uses Hypermedia as the Engine of Application State (HATEOAS) to drive the discovery and provide   affordances. - Discovery is possible by following links from the well known root resource. While this specification lists   all supported endpoints, it is only recommended that these are hard coded into a client if code generation is   being used. Otherwise, it is recommended that the discovery mechanisms present in the resources (affordances)   are used exclusively. - Affordances are links which indicate whether an action is currently possible, this is significantly different from   whether the service supports an action in general. This specification defines what actions could be possible,   but only by checking the affordances returned by the API in the returned resources, can a client determine whether   this action is currently possible or available for the current user. For example:   - An operation to modify a resource could be defined in this specification, but the user may lack the appropriate     privileges. In that situation, the affordance link would not be present in the resource when read. Therefore,     the client can infer that it is not possible to edit this resource and present appropriate information to the     user.   - An operation to delete a resource could be defined and be possible in some circumstances. The specification     describes that the delete is supported and how to use it, but the affordance describes whether it is currently     possible. The logic in the API may dictate that if the resource was in use (perhaps it is a running job or used     by another resource), then it will not be possible to delete that resource as it would result in a conflicted     state. - It is strongly encouraged that affordances are used by all clients, even those using code generation. This has the   ability to both improve robustness and the user experience by decoupling the client and server. For example, if for   some reason the criteria for deleting a resource changes, the logic is only implemented in the server and there is   no need to update the logic in the client as it is driven by the affordances. - The format used for the resources is the Hypertext Application Language (HAL), which includes the definition   of links and embedded resources. 

API version: 1.1.1
Contact: support@arm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the FieldObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldObject{}

// FieldObject struct for FieldObject
type FieldObject struct {
	// Name of the field name in the request which has failed validation.
	FieldName string `json:"fieldName"`
	// Field name, possibly including the path of the field which caused the error.
	FieldPath *string `json:"fieldPath,omitempty"`
	// A human readable message, which should provide an explanation and possible corrective actions.
	Message string `json:"message"`
}

type _FieldObject FieldObject

// NewFieldObject instantiates a new FieldObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldObject(fieldName string, message string) *FieldObject {
	this := FieldObject{}
	this.FieldName = fieldName
	this.Message = message
	return &this
}

// NewFieldObjectWithDefaults instantiates a new FieldObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldObjectWithDefaults() *FieldObject {
	this := FieldObject{}
	return &this
}

// GetFieldName returns the FieldName field value
func (o *FieldObject) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *FieldObject) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *FieldObject) SetFieldName(v string) {
	o.FieldName = v
}

// GetFieldPath returns the FieldPath field value if set, zero value otherwise.
func (o *FieldObject) GetFieldPath() string {
	if o == nil || IsNil(o.FieldPath) {
		var ret string
		return ret
	}
	return *o.FieldPath
}

// GetFieldPathOk returns a tuple with the FieldPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldObject) GetFieldPathOk() (*string, bool) {
	if o == nil || IsNil(o.FieldPath) {
		return nil, false
	}
	return o.FieldPath, true
}

// HasFieldPath returns a boolean if a field has been set.
func (o *FieldObject) HasFieldPath() bool {
	if o != nil && !IsNil(o.FieldPath) {
		return true
	}

	return false
}

// SetFieldPath gets a reference to the given string and assigns it to the FieldPath field.
func (o *FieldObject) SetFieldPath(v string) {
	o.FieldPath = &v
}

// GetMessage returns the Message field value
func (o *FieldObject) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *FieldObject) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *FieldObject) SetMessage(v string) {
	o.Message = v
}

func (o FieldObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fieldName"] = o.FieldName
	if !IsNil(o.FieldPath) {
		toSerialize["fieldPath"] = o.FieldPath
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *FieldObject) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fieldName",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFieldObject := _FieldObject{}

	err = json.Unmarshal(bytes, &varFieldObject)

	if err != nil {
		return err
	}

	*o = FieldObject(varFieldObject)

	return err
}

type NullableFieldObject struct {
	value *FieldObject
	isSet bool
}

func (v NullableFieldObject) Get() *FieldObject {
	return v.value
}

func (v *NullableFieldObject) Set(val *FieldObject) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldObject) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldObject(val *FieldObject) *NullableFieldObject {
	return &NullableFieldObject{value: val, isSet: true}
}

func (v NullableFieldObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


