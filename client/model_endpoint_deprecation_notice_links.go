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

// checks if the EndpointDeprecationNoticeLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointDeprecationNoticeLinks{}

// EndpointDeprecationNoticeLinks The `collection` link corresponds to the collection containing this notice
type EndpointDeprecationNoticeLinks struct {
	Collection *HalLinkData `json:"collection,omitempty"`
	Self HalLinkData `json:"self"`
}

type _EndpointDeprecationNoticeLinks EndpointDeprecationNoticeLinks

// NewEndpointDeprecationNoticeLinks instantiates a new EndpointDeprecationNoticeLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointDeprecationNoticeLinks(self HalLinkData) *EndpointDeprecationNoticeLinks {
	this := EndpointDeprecationNoticeLinks{}
	this.Self = self
	return &this
}

// NewEndpointDeprecationNoticeLinksWithDefaults instantiates a new EndpointDeprecationNoticeLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointDeprecationNoticeLinksWithDefaults() *EndpointDeprecationNoticeLinks {
	this := EndpointDeprecationNoticeLinks{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *EndpointDeprecationNoticeLinks) GetCollection() HalLinkData {
	if o == nil || IsNil(o.Collection) {
		var ret HalLinkData
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointDeprecationNoticeLinks) GetCollectionOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *EndpointDeprecationNoticeLinks) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given HalLinkData and assigns it to the Collection field.
func (o *EndpointDeprecationNoticeLinks) SetCollection(v HalLinkData) {
	o.Collection = &v
}

// GetSelf returns the Self field value
func (o *EndpointDeprecationNoticeLinks) GetSelf() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *EndpointDeprecationNoticeLinks) GetSelfOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *EndpointDeprecationNoticeLinks) SetSelf(v HalLinkData) {
	o.Self = v
}

func (o EndpointDeprecationNoticeLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointDeprecationNoticeLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	toSerialize["self"] = o.Self
	return toSerialize, nil
}

func (o *EndpointDeprecationNoticeLinks) UnmarshalJSON(data []byte) (err error) {
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

	varEndpointDeprecationNoticeLinks := _EndpointDeprecationNoticeLinks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointDeprecationNoticeLinks)

	if err != nil {
		return err
	}

	*o = EndpointDeprecationNoticeLinks(varEndpointDeprecationNoticeLinks)

	return err
}

type NullableEndpointDeprecationNoticeLinks struct {
	value *EndpointDeprecationNoticeLinks
	isSet bool
}

func (v NullableEndpointDeprecationNoticeLinks) Get() *EndpointDeprecationNoticeLinks {
	return v.value
}

func (v *NullableEndpointDeprecationNoticeLinks) Set(val *EndpointDeprecationNoticeLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointDeprecationNoticeLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointDeprecationNoticeLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointDeprecationNoticeLinks(val *EndpointDeprecationNoticeLinks) *NullableEndpointDeprecationNoticeLinks {
	return &NullableEndpointDeprecationNoticeLinks{value: val, isSet: true}
}

func (v NullableEndpointDeprecationNoticeLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointDeprecationNoticeLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


