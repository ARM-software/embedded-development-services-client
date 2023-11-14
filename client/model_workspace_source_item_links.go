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
	"fmt"
)

// checks if the WorkspaceSourceItemLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkspaceSourceItemLinks{}

// WorkspaceSourceItemLinks The `create` link (if present) provides the URI to request a Workspace.
type WorkspaceSourceItemLinks struct {
	Collection *HalLinkData `json:"collection,omitempty"`
	Create *HalLinkData `json:"create,omitempty"`
	Self HalLinkData `json:"self"`
}

type _WorkspaceSourceItemLinks WorkspaceSourceItemLinks

// NewWorkspaceSourceItemLinks instantiates a new WorkspaceSourceItemLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceSourceItemLinks(self HalLinkData) *WorkspaceSourceItemLinks {
	this := WorkspaceSourceItemLinks{}
	this.Self = self
	return &this
}

// NewWorkspaceSourceItemLinksWithDefaults instantiates a new WorkspaceSourceItemLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceSourceItemLinksWithDefaults() *WorkspaceSourceItemLinks {
	this := WorkspaceSourceItemLinks{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *WorkspaceSourceItemLinks) GetCollection() HalLinkData {
	if o == nil || IsNil(o.Collection) {
		var ret HalLinkData
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceSourceItemLinks) GetCollectionOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *WorkspaceSourceItemLinks) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given HalLinkData and assigns it to the Collection field.
func (o *WorkspaceSourceItemLinks) SetCollection(v HalLinkData) {
	o.Collection = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *WorkspaceSourceItemLinks) GetCreate() HalLinkData {
	if o == nil || IsNil(o.Create) {
		var ret HalLinkData
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceSourceItemLinks) GetCreateOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *WorkspaceSourceItemLinks) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given HalLinkData and assigns it to the Create field.
func (o *WorkspaceSourceItemLinks) SetCreate(v HalLinkData) {
	o.Create = &v
}

// GetSelf returns the Self field value
func (o *WorkspaceSourceItemLinks) GetSelf() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSourceItemLinks) GetSelfOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *WorkspaceSourceItemLinks) SetSelf(v HalLinkData) {
	o.Self = v
}

func (o WorkspaceSourceItemLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkspaceSourceItemLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	toSerialize["self"] = o.Self
	return toSerialize, nil
}

func (o *WorkspaceSourceItemLinks) UnmarshalJSON(bytes []byte) (err error) {
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

	varWorkspaceSourceItemLinks := _WorkspaceSourceItemLinks{}

	err = json.Unmarshal(bytes, &varWorkspaceSourceItemLinks)

	if err != nil {
		return err
	}

	*o = WorkspaceSourceItemLinks(varWorkspaceSourceItemLinks)

	return err
}

type NullableWorkspaceSourceItemLinks struct {
	value *WorkspaceSourceItemLinks
	isSet bool
}

func (v NullableWorkspaceSourceItemLinks) Get() *WorkspaceSourceItemLinks {
	return v.value
}

func (v *NullableWorkspaceSourceItemLinks) Set(val *WorkspaceSourceItemLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceSourceItemLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceSourceItemLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceSourceItemLinks(val *WorkspaceSourceItemLinks) *NullableWorkspaceSourceItemLinks {
	return &NullableWorkspaceSourceItemLinks{value: val, isSet: true}
}

func (v NullableWorkspaceSourceItemLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceSourceItemLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


