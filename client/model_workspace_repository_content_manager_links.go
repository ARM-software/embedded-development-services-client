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

// checks if the WorkspaceRepositoryContentManagerLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkspaceRepositoryContentManagerLinks{}

// WorkspaceRepositoryContentManagerLinks links to manage the workspace contant
type WorkspaceRepositoryContentManagerLinks struct {
	Clear *HalLinkData `json:"clear,omitempty"`
	Edit *HalLinkData `json:"edit,omitempty"`
	Related HalLinkData `json:"related"`
	Self HalLinkData `json:"self"`
}

type _WorkspaceRepositoryContentManagerLinks WorkspaceRepositoryContentManagerLinks

// NewWorkspaceRepositoryContentManagerLinks instantiates a new WorkspaceRepositoryContentManagerLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceRepositoryContentManagerLinks(related HalLinkData, self HalLinkData) *WorkspaceRepositoryContentManagerLinks {
	this := WorkspaceRepositoryContentManagerLinks{}
	this.Related = related
	this.Self = self
	return &this
}

// NewWorkspaceRepositoryContentManagerLinksWithDefaults instantiates a new WorkspaceRepositoryContentManagerLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceRepositoryContentManagerLinksWithDefaults() *WorkspaceRepositoryContentManagerLinks {
	this := WorkspaceRepositoryContentManagerLinks{}
	return &this
}

// GetClear returns the Clear field value if set, zero value otherwise.
func (o *WorkspaceRepositoryContentManagerLinks) GetClear() HalLinkData {
	if o == nil || IsNil(o.Clear) {
		var ret HalLinkData
		return ret
	}
	return *o.Clear
}

// GetClearOk returns a tuple with the Clear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceRepositoryContentManagerLinks) GetClearOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Clear) {
		return nil, false
	}
	return o.Clear, true
}

// HasClear returns a boolean if a field has been set.
func (o *WorkspaceRepositoryContentManagerLinks) HasClear() bool {
	if o != nil && !IsNil(o.Clear) {
		return true
	}

	return false
}

// SetClear gets a reference to the given HalLinkData and assigns it to the Clear field.
func (o *WorkspaceRepositoryContentManagerLinks) SetClear(v HalLinkData) {
	o.Clear = &v
}

// GetEdit returns the Edit field value if set, zero value otherwise.
func (o *WorkspaceRepositoryContentManagerLinks) GetEdit() HalLinkData {
	if o == nil || IsNil(o.Edit) {
		var ret HalLinkData
		return ret
	}
	return *o.Edit
}

// GetEditOk returns a tuple with the Edit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceRepositoryContentManagerLinks) GetEditOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Edit) {
		return nil, false
	}
	return o.Edit, true
}

// HasEdit returns a boolean if a field has been set.
func (o *WorkspaceRepositoryContentManagerLinks) HasEdit() bool {
	if o != nil && !IsNil(o.Edit) {
		return true
	}

	return false
}

// SetEdit gets a reference to the given HalLinkData and assigns it to the Edit field.
func (o *WorkspaceRepositoryContentManagerLinks) SetEdit(v HalLinkData) {
	o.Edit = &v
}

// GetRelated returns the Related field value
func (o *WorkspaceRepositoryContentManagerLinks) GetRelated() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
func (o *WorkspaceRepositoryContentManagerLinks) GetRelatedOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Related, true
}

// SetRelated sets field value
func (o *WorkspaceRepositoryContentManagerLinks) SetRelated(v HalLinkData) {
	o.Related = v
}

// GetSelf returns the Self field value
func (o *WorkspaceRepositoryContentManagerLinks) GetSelf() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *WorkspaceRepositoryContentManagerLinks) GetSelfOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *WorkspaceRepositoryContentManagerLinks) SetSelf(v HalLinkData) {
	o.Self = v
}

func (o WorkspaceRepositoryContentManagerLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkspaceRepositoryContentManagerLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clear) {
		toSerialize["clear"] = o.Clear
	}
	if !IsNil(o.Edit) {
		toSerialize["edit"] = o.Edit
	}
	toSerialize["related"] = o.Related
	toSerialize["self"] = o.Self
	return toSerialize, nil
}

func (o *WorkspaceRepositoryContentManagerLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"related",
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

	varWorkspaceRepositoryContentManagerLinks := _WorkspaceRepositoryContentManagerLinks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkspaceRepositoryContentManagerLinks)

	if err != nil {
		return err
	}

	*o = WorkspaceRepositoryContentManagerLinks(varWorkspaceRepositoryContentManagerLinks)

	return err
}

type NullableWorkspaceRepositoryContentManagerLinks struct {
	value *WorkspaceRepositoryContentManagerLinks
	isSet bool
}

func (v NullableWorkspaceRepositoryContentManagerLinks) Get() *WorkspaceRepositoryContentManagerLinks {
	return v.value
}

func (v *NullableWorkspaceRepositoryContentManagerLinks) Set(val *WorkspaceRepositoryContentManagerLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceRepositoryContentManagerLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceRepositoryContentManagerLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceRepositoryContentManagerLinks(val *WorkspaceRepositoryContentManagerLinks) *NullableWorkspaceRepositoryContentManagerLinks {
	return &NullableWorkspaceRepositoryContentManagerLinks{value: val, isSet: true}
}

func (v NullableWorkspaceRepositoryContentManagerLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceRepositoryContentManagerLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


