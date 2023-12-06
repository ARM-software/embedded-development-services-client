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

// checks if the VhtInstanceItemLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VhtInstanceItemLinks{}

// VhtInstanceItemLinks The `related` link indicates the VHT defining the instance. The `details` links to a resource that provides output messages of a VHT instance. The `artefacts` links to a collection of artefact managers which are needed for a run job to happen. The `create` link (if present) provides the URI to start a VHT run job.
type VhtInstanceItemLinks struct {
	Artefacts *HalLinkData `json:"artefacts,omitempty"`
	Cancel *HalLinkData `json:"cancel,omitempty"`
	Collection *HalLinkData `json:"collection,omitempty"`
	Create *HalLinkData `json:"create,omitempty"`
	Delete *HalLinkData `json:"delete,omitempty"`
	Details *HalLinkData `json:"details,omitempty"`
	Related HalLinkData `json:"related"`
	Self HalLinkData `json:"self"`
}

type _VhtInstanceItemLinks VhtInstanceItemLinks

// NewVhtInstanceItemLinks instantiates a new VhtInstanceItemLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVhtInstanceItemLinks(related HalLinkData, self HalLinkData) *VhtInstanceItemLinks {
	this := VhtInstanceItemLinks{}
	this.Related = related
	this.Self = self
	return &this
}

// NewVhtInstanceItemLinksWithDefaults instantiates a new VhtInstanceItemLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVhtInstanceItemLinksWithDefaults() *VhtInstanceItemLinks {
	this := VhtInstanceItemLinks{}
	return &this
}

// GetArtefacts returns the Artefacts field value if set, zero value otherwise.
func (o *VhtInstanceItemLinks) GetArtefacts() HalLinkData {
	if o == nil || IsNil(o.Artefacts) {
		var ret HalLinkData
		return ret
	}
	return *o.Artefacts
}

// GetArtefactsOk returns a tuple with the Artefacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtInstanceItemLinks) GetArtefactsOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Artefacts) {
		return nil, false
	}
	return o.Artefacts, true
}

// HasArtefacts returns a boolean if a field has been set.
func (o *VhtInstanceItemLinks) HasArtefacts() bool {
	if o != nil && !IsNil(o.Artefacts) {
		return true
	}

	return false
}

// SetArtefacts gets a reference to the given HalLinkData and assigns it to the Artefacts field.
func (o *VhtInstanceItemLinks) SetArtefacts(v HalLinkData) {
	o.Artefacts = &v
}

// GetCancel returns the Cancel field value if set, zero value otherwise.
func (o *VhtInstanceItemLinks) GetCancel() HalLinkData {
	if o == nil || IsNil(o.Cancel) {
		var ret HalLinkData
		return ret
	}
	return *o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtInstanceItemLinks) GetCancelOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Cancel) {
		return nil, false
	}
	return o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *VhtInstanceItemLinks) HasCancel() bool {
	if o != nil && !IsNil(o.Cancel) {
		return true
	}

	return false
}

// SetCancel gets a reference to the given HalLinkData and assigns it to the Cancel field.
func (o *VhtInstanceItemLinks) SetCancel(v HalLinkData) {
	o.Cancel = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *VhtInstanceItemLinks) GetCollection() HalLinkData {
	if o == nil || IsNil(o.Collection) {
		var ret HalLinkData
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtInstanceItemLinks) GetCollectionOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *VhtInstanceItemLinks) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given HalLinkData and assigns it to the Collection field.
func (o *VhtInstanceItemLinks) SetCollection(v HalLinkData) {
	o.Collection = &v
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *VhtInstanceItemLinks) GetCreate() HalLinkData {
	if o == nil || IsNil(o.Create) {
		var ret HalLinkData
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtInstanceItemLinks) GetCreateOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *VhtInstanceItemLinks) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given HalLinkData and assigns it to the Create field.
func (o *VhtInstanceItemLinks) SetCreate(v HalLinkData) {
	o.Create = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *VhtInstanceItemLinks) GetDelete() HalLinkData {
	if o == nil || IsNil(o.Delete) {
		var ret HalLinkData
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtInstanceItemLinks) GetDeleteOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *VhtInstanceItemLinks) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given HalLinkData and assigns it to the Delete field.
func (o *VhtInstanceItemLinks) SetDelete(v HalLinkData) {
	o.Delete = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *VhtInstanceItemLinks) GetDetails() HalLinkData {
	if o == nil || IsNil(o.Details) {
		var ret HalLinkData
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtInstanceItemLinks) GetDetailsOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *VhtInstanceItemLinks) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given HalLinkData and assigns it to the Details field.
func (o *VhtInstanceItemLinks) SetDetails(v HalLinkData) {
	o.Details = &v
}

// GetRelated returns the Related field value
func (o *VhtInstanceItemLinks) GetRelated() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
func (o *VhtInstanceItemLinks) GetRelatedOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Related, true
}

// SetRelated sets field value
func (o *VhtInstanceItemLinks) SetRelated(v HalLinkData) {
	o.Related = v
}

// GetSelf returns the Self field value
func (o *VhtInstanceItemLinks) GetSelf() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *VhtInstanceItemLinks) GetSelfOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *VhtInstanceItemLinks) SetSelf(v HalLinkData) {
	o.Self = v
}

func (o VhtInstanceItemLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VhtInstanceItemLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Artefacts) {
		toSerialize["artefacts"] = o.Artefacts
	}
	if !IsNil(o.Cancel) {
		toSerialize["cancel"] = o.Cancel
	}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	toSerialize["related"] = o.Related
	toSerialize["self"] = o.Self
	return toSerialize, nil
}

func (o *VhtInstanceItemLinks) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"related",
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

	varVhtInstanceItemLinks := _VhtInstanceItemLinks{}

	err = json.Unmarshal(bytes, &varVhtInstanceItemLinks)

	if err != nil {
		return err
	}

	*o = VhtInstanceItemLinks(varVhtInstanceItemLinks)

	return err
}

type NullableVhtInstanceItemLinks struct {
	value *VhtInstanceItemLinks
	isSet bool
}

func (v NullableVhtInstanceItemLinks) Get() *VhtInstanceItemLinks {
	return v.value
}

func (v *NullableVhtInstanceItemLinks) Set(val *VhtInstanceItemLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableVhtInstanceItemLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableVhtInstanceItemLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVhtInstanceItemLinks(val *VhtInstanceItemLinks) *NullableVhtInstanceItemLinks {
	return &NullableVhtInstanceItemLinks{value: val, isSet: true}
}

func (v NullableVhtInstanceItemLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVhtInstanceItemLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


