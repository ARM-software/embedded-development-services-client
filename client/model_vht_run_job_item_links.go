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

// checks if the VhtRunJobItemLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VhtRunJobItemLinks{}

// VhtRunJobItemLinks The `related` link indicates the VHT instance on which the job is run. The `details` links to a resource that provides details of progress (messages).
type VhtRunJobItemLinks struct {
	Cancel *HalLinkData `json:"cancel,omitempty"`
	Collection *HalLinkData `json:"collection,omitempty"`
	Delete *HalLinkData `json:"delete,omitempty"`
	Details *HalLinkData `json:"details,omitempty"`
	Related HalLinkData `json:"related"`
	Self HalLinkData `json:"self"`
}

type _VhtRunJobItemLinks VhtRunJobItemLinks

// NewVhtRunJobItemLinks instantiates a new VhtRunJobItemLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVhtRunJobItemLinks(related HalLinkData, self HalLinkData) *VhtRunJobItemLinks {
	this := VhtRunJobItemLinks{}
	this.Related = related
	this.Self = self
	return &this
}

// NewVhtRunJobItemLinksWithDefaults instantiates a new VhtRunJobItemLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVhtRunJobItemLinksWithDefaults() *VhtRunJobItemLinks {
	this := VhtRunJobItemLinks{}
	return &this
}

// GetCancel returns the Cancel field value if set, zero value otherwise.
func (o *VhtRunJobItemLinks) GetCancel() HalLinkData {
	if o == nil || IsNil(o.Cancel) {
		var ret HalLinkData
		return ret
	}
	return *o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtRunJobItemLinks) GetCancelOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Cancel) {
		return nil, false
	}
	return o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *VhtRunJobItemLinks) HasCancel() bool {
	if o != nil && !IsNil(o.Cancel) {
		return true
	}

	return false
}

// SetCancel gets a reference to the given HalLinkData and assigns it to the Cancel field.
func (o *VhtRunJobItemLinks) SetCancel(v HalLinkData) {
	o.Cancel = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *VhtRunJobItemLinks) GetCollection() HalLinkData {
	if o == nil || IsNil(o.Collection) {
		var ret HalLinkData
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtRunJobItemLinks) GetCollectionOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *VhtRunJobItemLinks) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given HalLinkData and assigns it to the Collection field.
func (o *VhtRunJobItemLinks) SetCollection(v HalLinkData) {
	o.Collection = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *VhtRunJobItemLinks) GetDelete() HalLinkData {
	if o == nil || IsNil(o.Delete) {
		var ret HalLinkData
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtRunJobItemLinks) GetDeleteOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *VhtRunJobItemLinks) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given HalLinkData and assigns it to the Delete field.
func (o *VhtRunJobItemLinks) SetDelete(v HalLinkData) {
	o.Delete = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *VhtRunJobItemLinks) GetDetails() HalLinkData {
	if o == nil || IsNil(o.Details) {
		var ret HalLinkData
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtRunJobItemLinks) GetDetailsOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *VhtRunJobItemLinks) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given HalLinkData and assigns it to the Details field.
func (o *VhtRunJobItemLinks) SetDetails(v HalLinkData) {
	o.Details = &v
}

// GetRelated returns the Related field value
func (o *VhtRunJobItemLinks) GetRelated() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
func (o *VhtRunJobItemLinks) GetRelatedOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Related, true
}

// SetRelated sets field value
func (o *VhtRunJobItemLinks) SetRelated(v HalLinkData) {
	o.Related = v
}

// GetSelf returns the Self field value
func (o *VhtRunJobItemLinks) GetSelf() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *VhtRunJobItemLinks) GetSelfOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *VhtRunJobItemLinks) SetSelf(v HalLinkData) {
	o.Self = v
}

func (o VhtRunJobItemLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VhtRunJobItemLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cancel) {
		toSerialize["cancel"] = o.Cancel
	}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
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

func (o *VhtRunJobItemLinks) UnmarshalJSON(bytes []byte) (err error) {
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

	varVhtRunJobItemLinks := _VhtRunJobItemLinks{}

	err = json.Unmarshal(bytes, &varVhtRunJobItemLinks)

	if err != nil {
		return err
	}

	*o = VhtRunJobItemLinks(varVhtRunJobItemLinks)

	return err
}

type NullableVhtRunJobItemLinks struct {
	value *VhtRunJobItemLinks
	isSet bool
}

func (v NullableVhtRunJobItemLinks) Get() *VhtRunJobItemLinks {
	return v.value
}

func (v *NullableVhtRunJobItemLinks) Set(val *VhtRunJobItemLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableVhtRunJobItemLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableVhtRunJobItemLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVhtRunJobItemLinks(val *VhtRunJobItemLinks) *NullableVhtRunJobItemLinks {
	return &NullableVhtRunJobItemLinks{value: val, isSet: true}
}

func (v NullableVhtRunJobItemLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVhtRunJobItemLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


