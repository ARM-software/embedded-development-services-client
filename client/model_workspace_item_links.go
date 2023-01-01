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

// WorkspaceItemLinks The `related` link indicates the source of the workspace content. The `artefacts` links to content manager.
type WorkspaceItemLinks struct {
	Artefacts *HalLinkData `json:"artefacts,omitempty"`
	Collection *HalLinkData `json:"collection,omitempty"`
	Delete *HalLinkData `json:"delete,omitempty"`
	Details *HalLinkData `json:"details,omitempty"`
	Related HalLinkData `json:"related"`
	Retain *HalLinkData `json:"retain,omitempty"`
	Self HalLinkData `json:"self"`
}

// NewWorkspaceItemLinks instantiates a new WorkspaceItemLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceItemLinks(related HalLinkData, self HalLinkData) *WorkspaceItemLinks {
	this := WorkspaceItemLinks{}
	this.Related = related
	this.Self = self
	return &this
}

// NewWorkspaceItemLinksWithDefaults instantiates a new WorkspaceItemLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceItemLinksWithDefaults() *WorkspaceItemLinks {
	this := WorkspaceItemLinks{}
	return &this
}

// GetArtefacts returns the Artefacts field value if set, zero value otherwise.
func (o *WorkspaceItemLinks) GetArtefacts() HalLinkData {
	if o == nil || isNil(o.Artefacts) {
		var ret HalLinkData
		return ret
	}
	return *o.Artefacts
}

// GetArtefactsOk returns a tuple with the Artefacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceItemLinks) GetArtefactsOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Artefacts) {
    return nil, false
	}
	return o.Artefacts, true
}

// HasArtefacts returns a boolean if a field has been set.
func (o *WorkspaceItemLinks) HasArtefacts() bool {
	if o != nil && !isNil(o.Artefacts) {
		return true
	}

	return false
}

// SetArtefacts gets a reference to the given HalLinkData and assigns it to the Artefacts field.
func (o *WorkspaceItemLinks) SetArtefacts(v HalLinkData) {
	o.Artefacts = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *WorkspaceItemLinks) GetCollection() HalLinkData {
	if o == nil || isNil(o.Collection) {
		var ret HalLinkData
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceItemLinks) GetCollectionOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Collection) {
    return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *WorkspaceItemLinks) HasCollection() bool {
	if o != nil && !isNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given HalLinkData and assigns it to the Collection field.
func (o *WorkspaceItemLinks) SetCollection(v HalLinkData) {
	o.Collection = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *WorkspaceItemLinks) GetDelete() HalLinkData {
	if o == nil || isNil(o.Delete) {
		var ret HalLinkData
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceItemLinks) GetDeleteOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Delete) {
    return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *WorkspaceItemLinks) HasDelete() bool {
	if o != nil && !isNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given HalLinkData and assigns it to the Delete field.
func (o *WorkspaceItemLinks) SetDelete(v HalLinkData) {
	o.Delete = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *WorkspaceItemLinks) GetDetails() HalLinkData {
	if o == nil || isNil(o.Details) {
		var ret HalLinkData
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceItemLinks) GetDetailsOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *WorkspaceItemLinks) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given HalLinkData and assigns it to the Details field.
func (o *WorkspaceItemLinks) SetDetails(v HalLinkData) {
	o.Details = &v
}

// GetRelated returns the Related field value
func (o *WorkspaceItemLinks) GetRelated() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
func (o *WorkspaceItemLinks) GetRelatedOk() (*HalLinkData, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Related, true
}

// SetRelated sets field value
func (o *WorkspaceItemLinks) SetRelated(v HalLinkData) {
	o.Related = v
}

// GetRetain returns the Retain field value if set, zero value otherwise.
func (o *WorkspaceItemLinks) GetRetain() HalLinkData {
	if o == nil || isNil(o.Retain) {
		var ret HalLinkData
		return ret
	}
	return *o.Retain
}

// GetRetainOk returns a tuple with the Retain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceItemLinks) GetRetainOk() (*HalLinkData, bool) {
	if o == nil || isNil(o.Retain) {
    return nil, false
	}
	return o.Retain, true
}

// HasRetain returns a boolean if a field has been set.
func (o *WorkspaceItemLinks) HasRetain() bool {
	if o != nil && !isNil(o.Retain) {
		return true
	}

	return false
}

// SetRetain gets a reference to the given HalLinkData and assigns it to the Retain field.
func (o *WorkspaceItemLinks) SetRetain(v HalLinkData) {
	o.Retain = &v
}

// GetSelf returns the Self field value
func (o *WorkspaceItemLinks) GetSelf() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *WorkspaceItemLinks) GetSelfOk() (*HalLinkData, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *WorkspaceItemLinks) SetSelf(v HalLinkData) {
	o.Self = v
}

func (o WorkspaceItemLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Artefacts) {
		toSerialize["artefacts"] = o.Artefacts
	}
	if !isNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !isNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if true {
		toSerialize["related"] = o.Related
	}
	if !isNil(o.Retain) {
		toSerialize["retain"] = o.Retain
	}
	if true {
		toSerialize["self"] = o.Self
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspaceItemLinks struct {
	value *WorkspaceItemLinks
	isSet bool
}

func (v NullableWorkspaceItemLinks) Get() *WorkspaceItemLinks {
	return v.value
}

func (v *NullableWorkspaceItemLinks) Set(val *WorkspaceItemLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceItemLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceItemLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceItemLinks(val *WorkspaceItemLinks) *NullableWorkspaceItemLinks {
	return &NullableWorkspaceItemLinks{value: val, isSet: true}
}

func (v NullableWorkspaceItemLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceItemLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


