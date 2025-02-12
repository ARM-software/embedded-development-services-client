/*
 * Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
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
)

// checks if the EmbeddedGenericWorkJobItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbeddedGenericWorkJobItems{}

// EmbeddedGenericWorkJobItems Embedded resource use the same link relation as a dictionary key, but rather than returning a link to the resource, the resource is instead embedded into the collection resource.
type EmbeddedGenericWorkJobItems struct {
	Item []GenericWorkJobItem `json:"item,omitempty"`
}

// NewEmbeddedGenericWorkJobItems instantiates a new EmbeddedGenericWorkJobItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddedGenericWorkJobItems() *EmbeddedGenericWorkJobItems {
	this := EmbeddedGenericWorkJobItems{}
	return &this
}

// NewEmbeddedGenericWorkJobItemsWithDefaults instantiates a new EmbeddedGenericWorkJobItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedGenericWorkJobItemsWithDefaults() *EmbeddedGenericWorkJobItems {
	this := EmbeddedGenericWorkJobItems{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *EmbeddedGenericWorkJobItems) GetItem() []GenericWorkJobItem {
	if o == nil || IsNil(o.Item) {
		var ret []GenericWorkJobItem
		return ret
	}
	return o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedGenericWorkJobItems) GetItemOk() ([]GenericWorkJobItem, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *EmbeddedGenericWorkJobItems) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given []GenericWorkJobItem and assigns it to the Item field.
func (o *EmbeddedGenericWorkJobItems) SetItem(v []GenericWorkJobItem) {
	o.Item = v
}

func (o EmbeddedGenericWorkJobItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbeddedGenericWorkJobItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	return toSerialize, nil
}

type NullableEmbeddedGenericWorkJobItems struct {
	value *EmbeddedGenericWorkJobItems
	isSet bool
}

func (v NullableEmbeddedGenericWorkJobItems) Get() *EmbeddedGenericWorkJobItems {
	return v.value
}

func (v *NullableEmbeddedGenericWorkJobItems) Set(val *EmbeddedGenericWorkJobItems) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedGenericWorkJobItems) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedGenericWorkJobItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedGenericWorkJobItems(val *EmbeddedGenericWorkJobItems) *NullableEmbeddedGenericWorkJobItems {
	return &NullableEmbeddedGenericWorkJobItems{value: val, isSet: true}
}

func (v NullableEmbeddedGenericWorkJobItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedGenericWorkJobItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


