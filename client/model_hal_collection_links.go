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

// checks if the HalCollectionLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HalCollectionLinks{}

// HalCollectionLinks These hypermedia links inside a collection resource allow contained resources to be discovered and for large collections to be paged through. - The `item` link relation contains a list of resources contained in the collection. If there is a `next` or   `previous` link present then not all the resources in the collection are currently being shown. - The `first` link points to the first page of a feed resource, this will not be be present if the entire   resource has been returned in the current page. - The `last` link points to the last page of a feed resource, this will not be be present if the entire   resource has been returned in the current page. - The `prev` link points to the previous page in the resource, this will not be present if the current page is   the first (or only page) page. - The `next` link points to the next page in the resource, this will not be present if the current page is the   last (or only page) page. - The `simple` link relation will be present in embedded representations to retrieve a non embedded   representation of the current context. - The `alternate` link relation is a simple templated URI to allow page selection by a client. - The `embedded`  link relation is an embedded representation of the current (non-embedded context).
type HalCollectionLinks struct {
	Alternate *HalLinkData `json:"alternate,omitempty"`
	Embedded *HalLinkData `json:"embedded,omitempty"`
	First *HalLinkData `json:"first,omitempty"`
	Item []HalLinkData `json:"item,omitempty"`
	Last *HalLinkData `json:"last,omitempty"`
	Next *HalLinkData `json:"next,omitempty"`
	Prev *HalLinkData `json:"prev,omitempty"`
	Self HalLinkData `json:"self"`
	Simple *HalLinkData `json:"simple,omitempty"`
}

// NewHalCollectionLinks instantiates a new HalCollectionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHalCollectionLinks(self HalLinkData) *HalCollectionLinks {
	this := HalCollectionLinks{}
	this.Self = self
	return &this
}

// NewHalCollectionLinksWithDefaults instantiates a new HalCollectionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHalCollectionLinksWithDefaults() *HalCollectionLinks {
	this := HalCollectionLinks{}
	return &this
}

// GetAlternate returns the Alternate field value if set, zero value otherwise.
func (o *HalCollectionLinks) GetAlternate() HalLinkData {
	if o == nil || isNil(o.Alternate) {
		var ret HalLinkData
		return ret
	}
	return *o.Alternate
}

// GetAlternateOk returns a tuple with the Alternate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalCollectionLinks) GetAlternateOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Alternate) {
		return nil, false
	}
	return o.Alternate, true
}

// HasAlternate returns a boolean if a field has been set.
func (o *HalCollectionLinks) HasAlternate() bool {
	if o != nil && !isNil(o.Alternate) {
		return true
	}

	return false
}

// SetAlternate gets a reference to the given HalLinkData and assigns it to the Alternate field.
func (o *HalCollectionLinks) SetAlternate(v HalLinkData) {
	o.Alternate = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *HalCollectionLinks) GetEmbedded() HalLinkData {
	if o == nil || isNil(o.Embedded) {
		var ret HalLinkData
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalCollectionLinks) GetEmbeddedOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *HalCollectionLinks) HasEmbedded() bool {
	if o != nil && !isNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given HalLinkData and assigns it to the Embedded field.
func (o *HalCollectionLinks) SetEmbedded(v HalLinkData) {
	o.Embedded = &v
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *HalCollectionLinks) GetFirst() HalLinkData {
	if o == nil || isNil(o.First) {
		var ret HalLinkData
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalCollectionLinks) GetFirstOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *HalCollectionLinks) HasFirst() bool {
	if o != nil && !isNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given HalLinkData and assigns it to the First field.
func (o *HalCollectionLinks) SetFirst(v HalLinkData) {
	o.First = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *HalCollectionLinks) GetItem() []HalLinkData {
	if o == nil || isNil(o.Item) {
		var ret []HalLinkData
		return ret
	}
	return o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalCollectionLinks) GetItemOk() ([]HalLinkData, bool) {
	if o == nil || isNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *HalCollectionLinks) HasItem() bool {
	if o != nil && !isNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given []HalLinkData and assigns it to the Item field.
func (o *HalCollectionLinks) SetItem(v []HalLinkData) {
	o.Item = v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *HalCollectionLinks) GetLast() HalLinkData {
	if o == nil || isNil(o.Last) {
		var ret HalLinkData
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalCollectionLinks) GetLastOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *HalCollectionLinks) HasLast() bool {
	if o != nil && !isNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given HalLinkData and assigns it to the Last field.
func (o *HalCollectionLinks) SetLast(v HalLinkData) {
	o.Last = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *HalCollectionLinks) GetNext() HalLinkData {
	if o == nil || isNil(o.Next) {
		var ret HalLinkData
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalCollectionLinks) GetNextOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *HalCollectionLinks) HasNext() bool {
	if o != nil && !isNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given HalLinkData and assigns it to the Next field.
func (o *HalCollectionLinks) SetNext(v HalLinkData) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *HalCollectionLinks) GetPrev() HalLinkData {
	if o == nil || isNil(o.Prev) {
		var ret HalLinkData
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalCollectionLinks) GetPrevOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *HalCollectionLinks) HasPrev() bool {
	if o != nil && !isNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given HalLinkData and assigns it to the Prev field.
func (o *HalCollectionLinks) SetPrev(v HalLinkData) {
	o.Prev = &v
}

// GetSelf returns the Self field value
func (o *HalCollectionLinks) GetSelf() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *HalCollectionLinks) GetSelfOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *HalCollectionLinks) SetSelf(v HalLinkData) {
	o.Self = v
}

// GetSimple returns the Simple field value if set, zero value otherwise.
func (o *HalCollectionLinks) GetSimple() HalLinkData {
	if o == nil || isNil(o.Simple) {
		var ret HalLinkData
		return ret
	}
	return *o.Simple
}

// GetSimpleOk returns a tuple with the Simple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalCollectionLinks) GetSimpleOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Simple) {
		return nil, false
	}
	return o.Simple, true
}

// HasSimple returns a boolean if a field has been set.
func (o *HalCollectionLinks) HasSimple() bool {
	if o != nil && !isNil(o.Simple) {
		return true
	}

	return false
}

// SetSimple gets a reference to the given HalLinkData and assigns it to the Simple field.
func (o *HalCollectionLinks) SetSimple(v HalLinkData) {
	o.Simple = &v
}

func (o HalCollectionLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HalCollectionLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Alternate) {
		toSerialize["alternate"] = o.Alternate
	}
	if !isNil(o.Embedded) {
		toSerialize["embedded"] = o.Embedded
	}
	if !isNil(o.First) {
		toSerialize["first"] = o.First
	}
	if !isNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	if !isNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !isNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !isNil(o.Prev) {
		toSerialize["prev"] = o.Prev
	}
	toSerialize["self"] = o.Self
	if !isNil(o.Simple) {
		toSerialize["simple"] = o.Simple
	}
	return toSerialize, nil
}

type NullableHalCollectionLinks struct {
	value *HalCollectionLinks
	isSet bool
}

func (v NullableHalCollectionLinks) Get() *HalCollectionLinks {
	return v.value
}

func (v *NullableHalCollectionLinks) Set(val *HalCollectionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableHalCollectionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableHalCollectionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHalCollectionLinks(val *HalCollectionLinks) *NullableHalCollectionLinks {
	return &NullableHalCollectionLinks{value: val, isSet: true}
}

func (v NullableHalCollectionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHalCollectionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


