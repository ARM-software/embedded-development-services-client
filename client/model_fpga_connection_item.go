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

// checks if the FPGAConnectionItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPGAConnectionItem{}

// FPGAConnectionItem struct for FPGAConnectionItem
type FPGAConnectionItem struct {
	Links NullableFPGAConnectionItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	// Number of current connections using this channel.
	Count *int32 `json:"count,omitempty"`
	// The identifier of the connection.
	Name string `json:"name"`
	// Whether the connection is ready to use.
	Ready *bool `json:"ready,omitempty"`
	// Status of the connection
	Status *string `json:"status,omitempty"`
	// Human readable description of the connection
	Title *string `json:"title,omitempty"`
}

type _FPGAConnectionItem FPGAConnectionItem

// NewFPGAConnectionItem instantiates a new FPGAConnectionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPGAConnectionItem(links NullableFPGAConnectionItemLinks, metadata NullableCommonMetadata, name string) *FPGAConnectionItem {
	this := FPGAConnectionItem{}
	this.Links = links
	this.Metadata = metadata
	this.Name = name
	return &this
}

// NewFPGAConnectionItemWithDefaults instantiates a new FPGAConnectionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPGAConnectionItemWithDefaults() *FPGAConnectionItem {
	this := FPGAConnectionItem{}
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for FPGAConnectionItemLinks will be returned
func (o *FPGAConnectionItem) GetLinks() FPGAConnectionItemLinks {
	if o == nil || o.Links.Get() == nil {
		var ret FPGAConnectionItemLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPGAConnectionItem) GetLinksOk() (*FPGAConnectionItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *FPGAConnectionItem) SetLinks(v FPGAConnectionItemLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *FPGAConnectionItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPGAConnectionItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *FPGAConnectionItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *FPGAConnectionItem) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FPGAConnectionItem) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *FPGAConnectionItem) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *FPGAConnectionItem) SetCount(v int32) {
	o.Count = &v
}

// GetName returns the Name field value
func (o *FPGAConnectionItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FPGAConnectionItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FPGAConnectionItem) SetName(v string) {
	o.Name = v
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *FPGAConnectionItem) GetReady() bool {
	if o == nil || IsNil(o.Ready) {
		var ret bool
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FPGAConnectionItem) GetReadyOk() (*bool, bool) {
	if o == nil || IsNil(o.Ready) {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *FPGAConnectionItem) HasReady() bool {
	if o != nil && !IsNil(o.Ready) {
		return true
	}

	return false
}

// SetReady gets a reference to the given bool and assigns it to the Ready field.
func (o *FPGAConnectionItem) SetReady(v bool) {
	o.Ready = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FPGAConnectionItem) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FPGAConnectionItem) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FPGAConnectionItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FPGAConnectionItem) SetStatus(v string) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *FPGAConnectionItem) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FPGAConnectionItem) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *FPGAConnectionItem) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *FPGAConnectionItem) SetTitle(v string) {
	o.Title = &v
}

func (o FPGAConnectionItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPGAConnectionItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links.Get()
	toSerialize["_metadata"] = o.Metadata.Get()
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Ready) {
		toSerialize["ready"] = o.Ready
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

func (o *FPGAConnectionItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_links",
		"_metadata",
		"name",
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

	varFPGAConnectionItem := _FPGAConnectionItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPGAConnectionItem)

	if err != nil {
		return err
	}

	*o = FPGAConnectionItem(varFPGAConnectionItem)

	return err
}

type NullableFPGAConnectionItem struct {
	value *FPGAConnectionItem
	isSet bool
}

func (v NullableFPGAConnectionItem) Get() *FPGAConnectionItem {
	return v.value
}

func (v *NullableFPGAConnectionItem) Set(val *FPGAConnectionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFPGAConnectionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFPGAConnectionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPGAConnectionItem(val *FPGAConnectionItem) *NullableFPGAConnectionItem {
	return &NullableFPGAConnectionItem{value: val, isSet: true}
}

func (v NullableFPGAConnectionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPGAConnectionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


