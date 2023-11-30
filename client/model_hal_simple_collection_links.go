/*
 * Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

/*
Solar API

This API provides a RESTful interface to all the Solar services e.g. looking for boards, building projects, etc. - This API uses Hypermedia as the Engine of Application State (HATEOAS) to drive the discovery and provide   affordances. - Discovery is possible by following links from the well known root resource. While this specification lists   all supported endpoints, it is only recommended that these are hard coded into a client if code generation is   being used. Otherwise, it is recommended that the discovery mechanisms present in the resources (affordances)   are used exclusively. - Affordances are links which indicate whether an action is currently possible, this is significantly different from   whether the service supports an action in general. This specification defines what actions could be possible,   but only by checking the affordances returned by the API in the returned resources, can a client determine whether   this action is currently possible or available for the current user. For example:   - An operation to modify a resource could be defined in this specification, but the user may lack the appropriate     privileges. In that situation, the affordance link would not be present in the resource when read. Therefore,     the client can infer that it is not possible to edit this resource and present appropriate information to the     user.   - An operation to delete a resource could be defined and be possible in some circumstances. The specification     describes that the delete is supported and how to use it, but the affordance describes whether it is currently     possible. The logic in the API may dictate that if the resource was in use (perhaps it is a running job or used     by another resource), then it will not be possible to delete that resource as it would result in a conflicted     state. - It is strongly encouraged that affordances are used by all clients, even those using code generation. This has the   ability to both improve robustness and the user experience by decoupling the client and server. For example, if for   some reason the criteria for deleting a resource changes, the logic is only implemented in the server and there is   no need to update the logic in the client as it is driven by the affordances. - The format used for the resources is the Hypertext Application Language (HAL), which includes the definition   of links and embedded resources. 

API version: 1.1.0
Contact: support@arm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the HalSimpleCollectionLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HalSimpleCollectionLinks{}

// HalSimpleCollectionLinks These hypermedia links inside a feed resource allow the contents of the resource to be paged. A simple collection only contains the minimum links needed to traverse a collection. - The `first` link points to the first page of a feed resource, this will not be be present if the entire   resource has been returned in the current page. - The `last` link points to the last page of a feed resource, this will not be be present if the entire   resource has been returned in the current page. - The `prev` link points to the previous page in the resource, this will not be present if the current page is   the first (or only page) page. - The `next` link points to the next page in the resource, this will not be present if the current page is the   last (or only page) page.
type HalSimpleCollectionLinks struct {
	First *HalLinkData `json:"first,omitempty"`
	Last *HalLinkData `json:"last,omitempty"`
	Next *HalLinkData `json:"next,omitempty"`
	Prev *HalLinkData `json:"prev,omitempty"`
	Self HalLinkData `json:"self"`
}

type _HalSimpleCollectionLinks HalSimpleCollectionLinks

// NewHalSimpleCollectionLinks instantiates a new HalSimpleCollectionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHalSimpleCollectionLinks(self HalLinkData) *HalSimpleCollectionLinks {
	this := HalSimpleCollectionLinks{}
	this.Self = self
	return &this
}

// NewHalSimpleCollectionLinksWithDefaults instantiates a new HalSimpleCollectionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHalSimpleCollectionLinksWithDefaults() *HalSimpleCollectionLinks {
	this := HalSimpleCollectionLinks{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *HalSimpleCollectionLinks) GetFirst() HalLinkData {
	if o == nil || IsNil(o.First) {
		var ret HalLinkData
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalSimpleCollectionLinks) GetFirstOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *HalSimpleCollectionLinks) HasFirst() bool {
	if o != nil && !IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given HalLinkData and assigns it to the First field.
func (o *HalSimpleCollectionLinks) SetFirst(v HalLinkData) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *HalSimpleCollectionLinks) GetLast() HalLinkData {
	if o == nil || IsNil(o.Last) {
		var ret HalLinkData
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalSimpleCollectionLinks) GetLastOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *HalSimpleCollectionLinks) HasLast() bool {
	if o != nil && !IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given HalLinkData and assigns it to the Last field.
func (o *HalSimpleCollectionLinks) SetLast(v HalLinkData) {
	o.Last = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *HalSimpleCollectionLinks) GetNext() HalLinkData {
	if o == nil || IsNil(o.Next) {
		var ret HalLinkData
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalSimpleCollectionLinks) GetNextOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *HalSimpleCollectionLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given HalLinkData and assigns it to the Next field.
func (o *HalSimpleCollectionLinks) SetNext(v HalLinkData) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *HalSimpleCollectionLinks) GetPrev() HalLinkData {
	if o == nil || IsNil(o.Prev) {
		var ret HalLinkData
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalSimpleCollectionLinks) GetPrevOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *HalSimpleCollectionLinks) HasPrev() bool {
	if o != nil && !IsNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given HalLinkData and assigns it to the Prev field.
func (o *HalSimpleCollectionLinks) SetPrev(v HalLinkData) {
	o.Prev = &v
}

// GetSelf returns the Self field value
func (o *HalSimpleCollectionLinks) GetSelf() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *HalSimpleCollectionLinks) GetSelfOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *HalSimpleCollectionLinks) SetSelf(v HalLinkData) {
	o.Self = v
}

func (o HalSimpleCollectionLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HalSimpleCollectionLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.First) {
		toSerialize["first"] = o.First
	}
	if !IsNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Prev) {
		toSerialize["prev"] = o.Prev
	}
	toSerialize["self"] = o.Self
	return toSerialize, nil
}

func (o *HalSimpleCollectionLinks) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
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

	varHalSimpleCollectionLinks := _HalSimpleCollectionLinks{}

	err = json.Unmarshal(bytes, &varHalSimpleCollectionLinks)

	if err != nil {
		return err
	}

	*o = HalSimpleCollectionLinks(varHalSimpleCollectionLinks)

	return err
}

type NullableHalSimpleCollectionLinks struct {
	value *HalSimpleCollectionLinks
	isSet bool
}

func (v NullableHalSimpleCollectionLinks) Get() *HalSimpleCollectionLinks {
	return v.value
}

func (v *NullableHalSimpleCollectionLinks) Set(val *HalSimpleCollectionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableHalSimpleCollectionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableHalSimpleCollectionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHalSimpleCollectionLinks(val *HalSimpleCollectionLinks) *NullableHalSimpleCollectionLinks {
	return &NullableHalSimpleCollectionLinks{value: val, isSet: true}
}

func (v NullableHalSimpleCollectionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHalSimpleCollectionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


