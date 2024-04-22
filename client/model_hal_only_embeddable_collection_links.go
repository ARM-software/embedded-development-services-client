/*
 * Copyright (C) 2020-2024 Arm Limited or its affiliates and Contributors. All rights reserved.
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
	"bytes"
	"fmt"
)

// checks if the HalOnlyEmbeddableCollectionLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HalOnlyEmbeddableCollectionLinks{}

// HalOnlyEmbeddableCollectionLinks These hypermedia links inside a collection resource allow contained resources to be discovered. - The `item` link relation contains a list of resources contained in the collection. - The `simple` link relation will be present in embedded representations to retrieve a non embedded   representation of the current context. - The `alternate` link relation is a simple templated URI to allow page selection by a client. - The `embedded`  link relation is an embedded representation of the current (non-embedded context).
type HalOnlyEmbeddableCollectionLinks struct {
	Alternate *HalLinkData `json:"alternate,omitempty"`
	Embedded *HalLinkData `json:"embedded,omitempty"`
	Item []HalLinkData `json:"item,omitempty"`
	Self HalLinkData `json:"self"`
	Simple *HalLinkData `json:"simple,omitempty"`
}

type _HalOnlyEmbeddableCollectionLinks HalOnlyEmbeddableCollectionLinks

// NewHalOnlyEmbeddableCollectionLinks instantiates a new HalOnlyEmbeddableCollectionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHalOnlyEmbeddableCollectionLinks(self HalLinkData) *HalOnlyEmbeddableCollectionLinks {
	this := HalOnlyEmbeddableCollectionLinks{}
	this.Self = self
	return &this
}

// NewHalOnlyEmbeddableCollectionLinksWithDefaults instantiates a new HalOnlyEmbeddableCollectionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHalOnlyEmbeddableCollectionLinksWithDefaults() *HalOnlyEmbeddableCollectionLinks {
	this := HalOnlyEmbeddableCollectionLinks{}
	return &this
}

// GetAlternate returns the Alternate field value if set, zero value otherwise.
func (o *HalOnlyEmbeddableCollectionLinks) GetAlternate() HalLinkData {
	if o == nil || IsNil(o.Alternate) {
		var ret HalLinkData
		return ret
	}
	return *o.Alternate
}

// GetAlternateOk returns a tuple with the Alternate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalOnlyEmbeddableCollectionLinks) GetAlternateOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Alternate) {
		return nil, false
	}
	return o.Alternate, true
}

// HasAlternate returns a boolean if a field has been set.
func (o *HalOnlyEmbeddableCollectionLinks) HasAlternate() bool {
	if o != nil && !IsNil(o.Alternate) {
		return true
	}

	return false
}

// SetAlternate gets a reference to the given HalLinkData and assigns it to the Alternate field.
func (o *HalOnlyEmbeddableCollectionLinks) SetAlternate(v HalLinkData) {
	o.Alternate = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *HalOnlyEmbeddableCollectionLinks) GetEmbedded() HalLinkData {
	if o == nil || IsNil(o.Embedded) {
		var ret HalLinkData
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalOnlyEmbeddableCollectionLinks) GetEmbeddedOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *HalOnlyEmbeddableCollectionLinks) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given HalLinkData and assigns it to the Embedded field.
func (o *HalOnlyEmbeddableCollectionLinks) SetEmbedded(v HalLinkData) {
	o.Embedded = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *HalOnlyEmbeddableCollectionLinks) GetItem() []HalLinkData {
	if o == nil || IsNil(o.Item) {
		var ret []HalLinkData
		return ret
	}
	return o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalOnlyEmbeddableCollectionLinks) GetItemOk() ([]HalLinkData, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *HalOnlyEmbeddableCollectionLinks) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given []HalLinkData and assigns it to the Item field.
func (o *HalOnlyEmbeddableCollectionLinks) SetItem(v []HalLinkData) {
	o.Item = v
}

// GetSelf returns the Self field value
func (o *HalOnlyEmbeddableCollectionLinks) GetSelf() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *HalOnlyEmbeddableCollectionLinks) GetSelfOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *HalOnlyEmbeddableCollectionLinks) SetSelf(v HalLinkData) {
	o.Self = v
}

// GetSimple returns the Simple field value if set, zero value otherwise.
func (o *HalOnlyEmbeddableCollectionLinks) GetSimple() HalLinkData {
	if o == nil || IsNil(o.Simple) {
		var ret HalLinkData
		return ret
	}
	return *o.Simple
}

// GetSimpleOk returns a tuple with the Simple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalOnlyEmbeddableCollectionLinks) GetSimpleOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Simple) {
		return nil, false
	}
	return o.Simple, true
}

// HasSimple returns a boolean if a field has been set.
func (o *HalOnlyEmbeddableCollectionLinks) HasSimple() bool {
	if o != nil && !IsNil(o.Simple) {
		return true
	}

	return false
}

// SetSimple gets a reference to the given HalLinkData and assigns it to the Simple field.
func (o *HalOnlyEmbeddableCollectionLinks) SetSimple(v HalLinkData) {
	o.Simple = &v
}

func (o HalOnlyEmbeddableCollectionLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HalOnlyEmbeddableCollectionLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alternate) {
		toSerialize["alternate"] = o.Alternate
	}
	if !IsNil(o.Embedded) {
		toSerialize["embedded"] = o.Embedded
	}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	toSerialize["self"] = o.Self
	if !IsNil(o.Simple) {
		toSerialize["simple"] = o.Simple
	}
	return toSerialize, nil
}

func (o *HalOnlyEmbeddableCollectionLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHalOnlyEmbeddableCollectionLinks := _HalOnlyEmbeddableCollectionLinks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHalOnlyEmbeddableCollectionLinks)

	if err != nil {
		return err
	}

	*o = HalOnlyEmbeddableCollectionLinks(varHalOnlyEmbeddableCollectionLinks)

	return err
}

type NullableHalOnlyEmbeddableCollectionLinks struct {
	value *HalOnlyEmbeddableCollectionLinks
	isSet bool
}

func (v NullableHalOnlyEmbeddableCollectionLinks) Get() *HalOnlyEmbeddableCollectionLinks {
	return v.value
}

func (v *NullableHalOnlyEmbeddableCollectionLinks) Set(val *HalOnlyEmbeddableCollectionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableHalOnlyEmbeddableCollectionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableHalOnlyEmbeddableCollectionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHalOnlyEmbeddableCollectionLinks(val *HalOnlyEmbeddableCollectionLinks) *NullableHalOnlyEmbeddableCollectionLinks {
	return &NullableHalOnlyEmbeddableCollectionLinks{value: val, isSet: true}
}

func (v NullableHalOnlyEmbeddableCollectionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHalOnlyEmbeddableCollectionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

