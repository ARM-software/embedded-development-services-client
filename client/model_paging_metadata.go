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
	"time"
	"fmt"
)

// checks if the PagingMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagingMetadata{}

// PagingMetadata Information present in each resource that supports paging. This information provides paging status and other information about the resource itself.
type PagingMetadata struct {
	// Paging metadata: The number of items returned in this message.
	Count int32 `json:"count"`
	// Creation Time: UTC date and time (in RFC3339 format) when the resource was created. If this is a system created resource, this will be a fixed time unaffected by user actions.
	Ctime time.Time `json:"ctime"`
	// Expiry Time: UTC date and time (in RFC3339 format) when the resource will be removed automatically by the system and become unavailable.
	Etime NullableTime `json:"etime,omitempty"`
	// Paging metadata: The limit on the number of items to return.
	Limit int32 `json:"limit"`
	// Last Modification Time: UTC date and time (in RFC3339 format) when the resource was last updated. For a resource that cannot be modified this will be the same as `ctime`.
	Mtime time.Time `json:"mtime"`
	// Paging metadata: The index of the first item returned.
	Offset int32 `json:"offset"`
	// Paging metadata: Total number of items that can be paged through.
	Total int32 `json:"total"`
}

type _PagingMetadata PagingMetadata

// NewPagingMetadata instantiates a new PagingMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagingMetadata(count int32, ctime time.Time, limit int32, mtime time.Time, offset int32, total int32) *PagingMetadata {
	this := PagingMetadata{}
	this.Count = count
	this.Ctime = ctime
	this.Limit = limit
	this.Mtime = mtime
	this.Offset = offset
	this.Total = total
	return &this
}

// NewPagingMetadataWithDefaults instantiates a new PagingMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingMetadataWithDefaults() *PagingMetadata {
	this := PagingMetadata{}
	return &this
}

// GetCount returns the Count field value
func (o *PagingMetadata) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PagingMetadata) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PagingMetadata) SetCount(v int32) {
	o.Count = v
}

// GetCtime returns the Ctime field value
func (o *PagingMetadata) GetCtime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Ctime
}

// GetCtimeOk returns a tuple with the Ctime field value
// and a boolean to check if the value has been set.
func (o *PagingMetadata) GetCtimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ctime, true
}

// SetCtime sets field value
func (o *PagingMetadata) SetCtime(v time.Time) {
	o.Ctime = v
}

// GetEtime returns the Etime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PagingMetadata) GetEtime() time.Time {
	if o == nil || IsNil(o.Etime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Etime.Get()
}

// GetEtimeOk returns a tuple with the Etime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PagingMetadata) GetEtimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Etime.Get(), o.Etime.IsSet()
}

// HasEtime returns a boolean if a field has been set.
func (o *PagingMetadata) HasEtime() bool {
	if o != nil && o.Etime.IsSet() {
		return true
	}

	return false
}

// SetEtime gets a reference to the given NullableTime and assigns it to the Etime field.
func (o *PagingMetadata) SetEtime(v time.Time) {
	o.Etime.Set(&v)
}
// SetEtimeNil sets the value for Etime to be an explicit nil
func (o *PagingMetadata) SetEtimeNil() {
	o.Etime.Set(nil)
}

// UnsetEtime ensures that no value is present for Etime, not even an explicit nil
func (o *PagingMetadata) UnsetEtime() {
	o.Etime.Unset()
}

// GetLimit returns the Limit field value
func (o *PagingMetadata) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PagingMetadata) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PagingMetadata) SetLimit(v int32) {
	o.Limit = v
}

// GetMtime returns the Mtime field value
func (o *PagingMetadata) GetMtime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Mtime
}

// GetMtimeOk returns a tuple with the Mtime field value
// and a boolean to check if the value has been set.
func (o *PagingMetadata) GetMtimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mtime, true
}

// SetMtime sets field value
func (o *PagingMetadata) SetMtime(v time.Time) {
	o.Mtime = v
}

// GetOffset returns the Offset field value
func (o *PagingMetadata) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *PagingMetadata) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *PagingMetadata) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *PagingMetadata) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PagingMetadata) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PagingMetadata) SetTotal(v int32) {
	o.Total = v
}

func (o PagingMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagingMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["ctime"] = o.Ctime
	if o.Etime.IsSet() {
		toSerialize["etime"] = o.Etime.Get()
	}
	toSerialize["limit"] = o.Limit
	toSerialize["mtime"] = o.Mtime
	toSerialize["offset"] = o.Offset
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

func (o *PagingMetadata) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"ctime",
		"limit",
		"mtime",
		"offset",
		"total",
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

	varPagingMetadata := _PagingMetadata{}

	err = json.Unmarshal(bytes, &varPagingMetadata)

	if err != nil {
		return err
	}

	*o = PagingMetadata(varPagingMetadata)

	return err
}

type NullablePagingMetadata struct {
	value *PagingMetadata
	isSet bool
}

func (v NullablePagingMetadata) Get() *PagingMetadata {
	return v.value
}

func (v *NullablePagingMetadata) Set(val *PagingMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePagingMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePagingMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagingMetadata(val *PagingMetadata) *NullablePagingMetadata {
	return &NullablePagingMetadata{value: val, isSet: true}
}

func (v NullablePagingMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagingMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


