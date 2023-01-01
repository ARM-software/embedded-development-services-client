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
)

// ADebugInterface struct for ADebugInterface
type ADebugInterface struct {
	// Adapter for the interface
	Adapter *string `json:"adapter,omitempty"`
	// Connector for the interface
	Connector *string `json:"connector,omitempty"`
}

// NewADebugInterface instantiates a new ADebugInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewADebugInterface() *ADebugInterface {
	this := ADebugInterface{}
	return &this
}

// NewADebugInterfaceWithDefaults instantiates a new ADebugInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewADebugInterfaceWithDefaults() *ADebugInterface {
	this := ADebugInterface{}
	return &this
}

// GetAdapter returns the Adapter field value if set, zero value otherwise.
func (o *ADebugInterface) GetAdapter() string {
	if o == nil || isNil(o.Adapter) {
		var ret string
		return ret
	}
	return *o.Adapter
}

// GetAdapterOk returns a tuple with the Adapter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADebugInterface) GetAdapterOk() (*string, bool) {
	if o == nil || isNil(o.Adapter) {
    return nil, false
	}
	return o.Adapter, true
}

// HasAdapter returns a boolean if a field has been set.
func (o *ADebugInterface) HasAdapter() bool {
	if o != nil && !isNil(o.Adapter) {
		return true
	}

	return false
}

// SetAdapter gets a reference to the given string and assigns it to the Adapter field.
func (o *ADebugInterface) SetAdapter(v string) {
	o.Adapter = &v
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *ADebugInterface) GetConnector() string {
	if o == nil || isNil(o.Connector) {
		var ret string
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADebugInterface) GetConnectorOk() (*string, bool) {
	if o == nil || isNil(o.Connector) {
    return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *ADebugInterface) HasConnector() bool {
	if o != nil && !isNil(o.Connector) {
		return true
	}

	return false
}

// SetConnector gets a reference to the given string and assigns it to the Connector field.
func (o *ADebugInterface) SetConnector(v string) {
	o.Connector = &v
}

func (o ADebugInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Adapter) {
		toSerialize["adapter"] = o.Adapter
	}
	if !isNil(o.Connector) {
		toSerialize["connector"] = o.Connector
	}
	return json.Marshal(toSerialize)
}

type NullableADebugInterface struct {
	value *ADebugInterface
	isSet bool
}

func (v NullableADebugInterface) Get() *ADebugInterface {
	return v.value
}

func (v *NullableADebugInterface) Set(val *ADebugInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableADebugInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableADebugInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableADebugInterface(val *ADebugInterface) *NullableADebugInterface {
	return &NullableADebugInterface{value: val, isSet: true}
}

func (v NullableADebugInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableADebugInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


