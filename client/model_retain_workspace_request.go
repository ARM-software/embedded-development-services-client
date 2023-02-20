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

// checks if the RetainWorkspaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetainWorkspaceRequest{}

// RetainWorkspaceRequest struct for RetainWorkspaceRequest
type RetainWorkspaceRequest struct {
	// Time to live (in seconds), i.e. The the number of seconds from the current time (when this request is received) for which to keep the workspace and associated resources.
	Ttl NullableInt64 `json:"ttl,omitempty"`
}

// NewRetainWorkspaceRequest instantiates a new RetainWorkspaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetainWorkspaceRequest() *RetainWorkspaceRequest {
	this := RetainWorkspaceRequest{}
	return &this
}

// NewRetainWorkspaceRequestWithDefaults instantiates a new RetainWorkspaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetainWorkspaceRequestWithDefaults() *RetainWorkspaceRequest {
	this := RetainWorkspaceRequest{}
	return &this
}

// GetTtl returns the Ttl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RetainWorkspaceRequest) GetTtl() int64 {
	if o == nil || IsNil(o.Ttl.Get()) {
		var ret int64
		return ret
	}
	return *o.Ttl.Get()
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetainWorkspaceRequest) GetTtlOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ttl.Get(), o.Ttl.IsSet()
}

// HasTtl returns a boolean if a field has been set.
func (o *RetainWorkspaceRequest) HasTtl() bool {
	if o != nil && o.Ttl.IsSet() {
		return true
	}

	return false
}

// SetTtl gets a reference to the given NullableInt64 and assigns it to the Ttl field.
func (o *RetainWorkspaceRequest) SetTtl(v int64) {
	o.Ttl.Set(&v)
}
// SetTtlNil sets the value for Ttl to be an explicit nil
func (o *RetainWorkspaceRequest) SetTtlNil() {
	o.Ttl.Set(nil)
}

// UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
func (o *RetainWorkspaceRequest) UnsetTtl() {
	o.Ttl.Unset()
}

func (o RetainWorkspaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetainWorkspaceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ttl.IsSet() {
		toSerialize["ttl"] = o.Ttl.Get()
	}
	return toSerialize, nil
}

type NullableRetainWorkspaceRequest struct {
	value *RetainWorkspaceRequest
	isSet bool
}

func (v NullableRetainWorkspaceRequest) Get() *RetainWorkspaceRequest {
	return v.value
}

func (v *NullableRetainWorkspaceRequest) Set(val *RetainWorkspaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRetainWorkspaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRetainWorkspaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetainWorkspaceRequest(val *RetainWorkspaceRequest) *NullableRetainWorkspaceRequest {
	return &NullableRetainWorkspaceRequest{value: val, isSet: true}
}

func (v NullableRetainWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetainWorkspaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


