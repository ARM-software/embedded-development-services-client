/*
 * Copyright (C) 2020-2024 Arm Limited or its affiliates and Contributors. All rights reserved.
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

// checks if the VhtInstanceItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VhtInstanceItem{}

// VhtInstanceItem When the VHT instance is ready, it will include the current status of the instance and links to other available resources, such as instance messages and instance artefacts.
type VhtInstanceItem struct {
	Links NullableVhtInstanceItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	// True if an error occurred with the instance (e.g. network outage, etc.).
	Error bool `json:"error"`
	// The maximum time (in minutes) that the instance will be available for. After the timeout has expired the instance will be stopped and released. The timeout does not include any time the request spent being queued, waiting for the instance to be started.
	InstanceTimeout *int32 `json:"instanceTimeout,omitempty"`
	// Unique ID of the VHT instance.
	Name string `json:"name"`
	// True if an instance has all the artefacts necessary to carry out run jobs.
	Ready bool `json:"ready"`
	// True if an instance has been requested.
	Requested bool `json:"requested"`
	// A summary status of the VHT instance.  Note: this value should not be relied upon to determine whether a VHT is ready as this list may change as state machine evolves.  Use resource appropriate flags instead.
	Status string `json:"status"`
	// True when the instance has been terminated (this necessarily indicates that the instance is no longer available).
	Terminated bool `json:"terminated"`
	// Optional human readable name of the VHT instance.
	Title NullableString `json:"title,omitempty"`
}

type _VhtInstanceItem VhtInstanceItem

// NewVhtInstanceItem instantiates a new VhtInstanceItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVhtInstanceItem(links NullableVhtInstanceItemLinks, metadata NullableCommonMetadata, error_ bool, name string, ready bool, requested bool, status string, terminated bool) *VhtInstanceItem {
	this := VhtInstanceItem{}
	this.Links = links
	this.Metadata = metadata
	this.Error = error_
	var instanceTimeout int32 = 60
	this.InstanceTimeout = &instanceTimeout
	this.Name = name
	this.Ready = ready
	this.Requested = requested
	this.Status = status
	this.Terminated = terminated
	return &this
}

// NewVhtInstanceItemWithDefaults instantiates a new VhtInstanceItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVhtInstanceItemWithDefaults() *VhtInstanceItem {
	this := VhtInstanceItem{}
	var instanceTimeout int32 = 60
	this.InstanceTimeout = &instanceTimeout
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for VhtInstanceItemLinks will be returned
func (o *VhtInstanceItem) GetLinks() VhtInstanceItemLinks {
	if o == nil || o.Links.Get() == nil {
		var ret VhtInstanceItemLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VhtInstanceItem) GetLinksOk() (*VhtInstanceItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *VhtInstanceItem) SetLinks(v VhtInstanceItemLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *VhtInstanceItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VhtInstanceItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *VhtInstanceItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetError returns the Error field value
func (o *VhtInstanceItem) GetError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *VhtInstanceItem) GetErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *VhtInstanceItem) SetError(v bool) {
	o.Error = v
}

// GetInstanceTimeout returns the InstanceTimeout field value if set, zero value otherwise.
func (o *VhtInstanceItem) GetInstanceTimeout() int32 {
	if o == nil || IsNil(o.InstanceTimeout) {
		var ret int32
		return ret
	}
	return *o.InstanceTimeout
}

// GetInstanceTimeoutOk returns a tuple with the InstanceTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtInstanceItem) GetInstanceTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.InstanceTimeout) {
		return nil, false
	}
	return o.InstanceTimeout, true
}

// HasInstanceTimeout returns a boolean if a field has been set.
func (o *VhtInstanceItem) HasInstanceTimeout() bool {
	if o != nil && !IsNil(o.InstanceTimeout) {
		return true
	}

	return false
}

// SetInstanceTimeout gets a reference to the given int32 and assigns it to the InstanceTimeout field.
func (o *VhtInstanceItem) SetInstanceTimeout(v int32) {
	o.InstanceTimeout = &v
}

// GetName returns the Name field value
func (o *VhtInstanceItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VhtInstanceItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VhtInstanceItem) SetName(v string) {
	o.Name = v
}

// GetReady returns the Ready field value
func (o *VhtInstanceItem) GetReady() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ready
}

// GetReadyOk returns a tuple with the Ready field value
// and a boolean to check if the value has been set.
func (o *VhtInstanceItem) GetReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ready, true
}

// SetReady sets field value
func (o *VhtInstanceItem) SetReady(v bool) {
	o.Ready = v
}

// GetRequested returns the Requested field value
func (o *VhtInstanceItem) GetRequested() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value
// and a boolean to check if the value has been set.
func (o *VhtInstanceItem) GetRequestedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requested, true
}

// SetRequested sets field value
func (o *VhtInstanceItem) SetRequested(v bool) {
	o.Requested = v
}

// GetStatus returns the Status field value
func (o *VhtInstanceItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VhtInstanceItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VhtInstanceItem) SetStatus(v string) {
	o.Status = v
}

// GetTerminated returns the Terminated field value
func (o *VhtInstanceItem) GetTerminated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Terminated
}

// GetTerminatedOk returns a tuple with the Terminated field value
// and a boolean to check if the value has been set.
func (o *VhtInstanceItem) GetTerminatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Terminated, true
}

// SetTerminated sets field value
func (o *VhtInstanceItem) SetTerminated(v bool) {
	o.Terminated = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VhtInstanceItem) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VhtInstanceItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *VhtInstanceItem) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *VhtInstanceItem) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *VhtInstanceItem) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *VhtInstanceItem) UnsetTitle() {
	o.Title.Unset()
}

func (o VhtInstanceItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VhtInstanceItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links.Get()
	toSerialize["_metadata"] = o.Metadata.Get()
	toSerialize["error"] = o.Error
	if !IsNil(o.InstanceTimeout) {
		toSerialize["instanceTimeout"] = o.InstanceTimeout
	}
	toSerialize["name"] = o.Name
	toSerialize["ready"] = o.Ready
	toSerialize["requested"] = o.Requested
	toSerialize["status"] = o.Status
	toSerialize["terminated"] = o.Terminated
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	return toSerialize, nil
}

func (o *VhtInstanceItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_links",
		"_metadata",
		"error",
		"name",
		"ready",
		"requested",
		"status",
		"terminated",
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

	varVhtInstanceItem := _VhtInstanceItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVhtInstanceItem)

	if err != nil {
		return err
	}

	*o = VhtInstanceItem(varVhtInstanceItem)

	return err
}

type NullableVhtInstanceItem struct {
	value *VhtInstanceItem
	isSet bool
}

func (v NullableVhtInstanceItem) Get() *VhtInstanceItem {
	return v.value
}

func (v *NullableVhtInstanceItem) Set(val *VhtInstanceItem) {
	v.value = val
	v.isSet = true
}

func (v NullableVhtInstanceItem) IsSet() bool {
	return v.isSet
}

func (v *NullableVhtInstanceItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVhtInstanceItem(val *VhtInstanceItem) *NullableVhtInstanceItem {
	return &NullableVhtInstanceItem{value: val, isSet: true}
}

func (v NullableVhtInstanceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVhtInstanceItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


