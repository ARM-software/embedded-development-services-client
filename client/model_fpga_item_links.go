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
	"bytes"
	"fmt"
)

// checks if the FPGAItemLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPGAItemLinks{}

// FPGAItemLinks The `create` link (if present) provides the URI where a new job can be spawned The `current` link (if present) provides the URI to the resource describing the job currently running on the FPGA The `prev` link (if present) provides the URI to collection of past jobs which were run on the FPGA The `next` link (if present) provides the URI to collection of queued jobs which will be run on the FPGA The `collection` link (if present) provides the URI to the collection in which this FPGA is an `item`
type FPGAItemLinks struct {
	Collection *HalLinkData `json:"collection,omitempty"`
	Create *HalLinkData `json:"create,omitempty"`
	Current *HalLinkData `json:"current,omitempty"`
	Next *HalLinkData `json:"next,omitempty"`
	Prev *HalLinkData `json:"prev,omitempty"`
	Self HalLinkData `json:"self"`
}

type _FPGAItemLinks FPGAItemLinks

// NewFPGAItemLinks instantiates a new FPGAItemLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPGAItemLinks(self HalLinkData) *FPGAItemLinks {
	this := FPGAItemLinks{}
	this.Self = self
	return &this
}

// NewFPGAItemLinksWithDefaults instantiates a new FPGAItemLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPGAItemLinksWithDefaults() *FPGAItemLinks {
	this := FPGAItemLinks{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *FPGAItemLinks) GetCollection() HalLinkData {
	if o == nil || IsNil(o.Collection) {
		var ret HalLinkData
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FPGAItemLinks) GetCollectionOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *FPGAItemLinks) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given HalLinkData and assigns it to the Collection field.
func (o *FPGAItemLinks) SetCollection(v HalLinkData) {
	o.Collection = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *FPGAItemLinks) GetCreate() HalLinkData {
	if o == nil || IsNil(o.Create) {
		var ret HalLinkData
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FPGAItemLinks) GetCreateOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *FPGAItemLinks) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given HalLinkData and assigns it to the Create field.
func (o *FPGAItemLinks) SetCreate(v HalLinkData) {
	o.Create = &v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *FPGAItemLinks) GetCurrent() HalLinkData {
	if o == nil || IsNil(o.Current) {
		var ret HalLinkData
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FPGAItemLinks) GetCurrentOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *FPGAItemLinks) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given HalLinkData and assigns it to the Current field.
func (o *FPGAItemLinks) SetCurrent(v HalLinkData) {
	o.Current = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *FPGAItemLinks) GetNext() HalLinkData {
	if o == nil || IsNil(o.Next) {
		var ret HalLinkData
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FPGAItemLinks) GetNextOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *FPGAItemLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given HalLinkData and assigns it to the Next field.
func (o *FPGAItemLinks) SetNext(v HalLinkData) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *FPGAItemLinks) GetPrev() HalLinkData {
	if o == nil || IsNil(o.Prev) {
		var ret HalLinkData
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FPGAItemLinks) GetPrevOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *FPGAItemLinks) HasPrev() bool {
	if o != nil && !IsNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given HalLinkData and assigns it to the Prev field.
func (o *FPGAItemLinks) SetPrev(v HalLinkData) {
	o.Prev = &v
}

// GetSelf returns the Self field value
func (o *FPGAItemLinks) GetSelf() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *FPGAItemLinks) GetSelfOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *FPGAItemLinks) SetSelf(v HalLinkData) {
	o.Self = v
}

func (o FPGAItemLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPGAItemLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
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

func (o *FPGAItemLinks) UnmarshalJSON(data []byte) (err error) {
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

	varFPGAItemLinks := _FPGAItemLinks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPGAItemLinks)

	if err != nil {
		return err
	}

	*o = FPGAItemLinks(varFPGAItemLinks)

	return err
}

type NullableFPGAItemLinks struct {
	value *FPGAItemLinks
	isSet bool
}

func (v NullableFPGAItemLinks) Get() *FPGAItemLinks {
	return v.value
}

func (v *NullableFPGAItemLinks) Set(val *FPGAItemLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableFPGAItemLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableFPGAItemLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPGAItemLinks(val *FPGAItemLinks) *NullableFPGAItemLinks {
	return &NullableFPGAItemLinks{value: val, isSet: true}
}

func (v NullableFPGAItemLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPGAItemLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


