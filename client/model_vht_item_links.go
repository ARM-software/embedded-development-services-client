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

// VhtItemLinks The `create` link (if present) provides the URI to request a VHT instance.
type VhtItemLinks struct {
	Collection *HalLinkData `json:"collection,omitempty"`
	Create *HalLinkData `json:"create,omitempty"`
	Instance *HalLinkData `json:"instance,omitempty"`
	Self HalLinkData `json:"self"`
}

// NewVhtItemLinks instantiates a new VhtItemLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVhtItemLinks(self HalLinkData) *VhtItemLinks {
	this := VhtItemLinks{}
	this.Self = self
	return &this
}

// NewVhtItemLinksWithDefaults instantiates a new VhtItemLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVhtItemLinksWithDefaults() *VhtItemLinks {
	this := VhtItemLinks{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *VhtItemLinks) GetCollection() HalLinkData {
	if o == nil || isNil(o.Collection) {
		var ret HalLinkData
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtItemLinks) GetCollectionOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Collection) {
    return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *VhtItemLinks) HasCollection() bool {
	if o != nil && !isNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given HalLinkData and assigns it to the Collection field.
func (o *VhtItemLinks) SetCollection(v HalLinkData) {
	o.Collection = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *VhtItemLinks) GetCreate() HalLinkData {
	if o == nil || isNil(o.Create) {
		var ret HalLinkData
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtItemLinks) GetCreateOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Create) {
    return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *VhtItemLinks) HasCreate() bool {
	if o != nil && !isNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given HalLinkData and assigns it to the Create field.
func (o *VhtItemLinks) SetCreate(v HalLinkData) {
	o.Create = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *VhtItemLinks) GetInstance() HalLinkData {
	if o == nil || isNil(o.Instance) {
		var ret HalLinkData
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtItemLinks) GetInstanceOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Instance) {
    return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *VhtItemLinks) HasInstance() bool {
	if o != nil && !isNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given HalLinkData and assigns it to the Instance field.
func (o *VhtItemLinks) SetInstance(v HalLinkData) {
	o.Instance = &v
}

// GetSelf returns the Self field value
func (o *VhtItemLinks) GetSelf() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *VhtItemLinks) GetSelfOk() (*HalLinkData, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *VhtItemLinks) SetSelf(v HalLinkData) {
	o.Self = v
}

func (o VhtItemLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !isNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !isNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if true {
		toSerialize["self"] = o.Self
	}
	return json.Marshal(toSerialize)
}

type NullableVhtItemLinks struct {
	value *VhtItemLinks
	isSet bool
}

func (v NullableVhtItemLinks) Get() *VhtItemLinks {
	return v.value
}

func (v *NullableVhtItemLinks) Set(val *VhtItemLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableVhtItemLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableVhtItemLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVhtItemLinks(val *VhtItemLinks) *NullableVhtItemLinks {
	return &NullableVhtItemLinks{value: val, isSet: true}
}

func (v NullableVhtItemLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVhtItemLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


