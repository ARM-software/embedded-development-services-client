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

// checks if the VhtRunJobItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VhtRunJobItem{}

// VhtRunJobItem When the job is ready, it will include the current status of the job and links to other available resources, such as messages.
type VhtRunJobItem struct {
	Links NullableVhtRunJobItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	// True when the job has completed.
	Done bool `json:"done"`
	// True if the job exited with an error code (this should be used in conjunction with the `done` property).
	Error bool `json:"error"`
	// The maximum time (in seconds) that the job will be allowed to run. After the timeout has expired the job will be aborted and reported as a failure. The timeout does not include any time the request spent being queued, waiting for the job to be started.
	JobTimeout *int32 `json:"jobTimeout,omitempty"`
	// Unique ID of the VHT run Job.
	Name string `json:"name"`
	// True if job is currently queued and waiting to be processed. Otherwise, the job is either currently being processed or ended.
	Queued bool `json:"queued"`
	// A summary status of the VHT run job.  Note: this value should not be relied upon to determine whether a job has completed, started or errored as this list may change as state machine evolves.  Use resource appropriate flags instead.
	Status string `json:"status"`
	// True if there was a system error in the service while running the job (this should be used in conjunction with the `done` property).
	SystemError bool `json:"systemError"`
	// True if the job timed out (this should be used in conjunction with the `done` property).
	TimedOut bool `json:"timedOut"`
	// Optional human readable name of the VHT run job.
	Title NullableString `json:"title,omitempty"`
}

type _VhtRunJobItem VhtRunJobItem

// NewVhtRunJobItem instantiates a new VhtRunJobItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVhtRunJobItem(links NullableVhtRunJobItemLinks, metadata NullableCommonMetadata, done bool, error_ bool, name string, queued bool, status string, systemError bool, timedOut bool) *VhtRunJobItem {
	this := VhtRunJobItem{}
	this.Links = links
	this.Metadata = metadata
	this.Done = done
	this.Error = error_
	var jobTimeout int32 = 300
	this.JobTimeout = &jobTimeout
	this.Name = name
	this.Queued = queued
	this.Status = status
	this.SystemError = systemError
	this.TimedOut = timedOut
	return &this
}

// NewVhtRunJobItemWithDefaults instantiates a new VhtRunJobItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVhtRunJobItemWithDefaults() *VhtRunJobItem {
	this := VhtRunJobItem{}
	var jobTimeout int32 = 300
	this.JobTimeout = &jobTimeout
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for VhtRunJobItemLinks will be returned
func (o *VhtRunJobItem) GetLinks() VhtRunJobItemLinks {
	if o == nil || o.Links.Get() == nil {
		var ret VhtRunJobItemLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VhtRunJobItem) GetLinksOk() (*VhtRunJobItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *VhtRunJobItem) SetLinks(v VhtRunJobItemLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *VhtRunJobItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VhtRunJobItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *VhtRunJobItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetDone returns the Done field value
func (o *VhtRunJobItem) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *VhtRunJobItem) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *VhtRunJobItem) SetDone(v bool) {
	o.Done = v
}

// GetError returns the Error field value
func (o *VhtRunJobItem) GetError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *VhtRunJobItem) GetErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *VhtRunJobItem) SetError(v bool) {
	o.Error = v
}

// GetJobTimeout returns the JobTimeout field value if set, zero value otherwise.
func (o *VhtRunJobItem) GetJobTimeout() int32 {
	if o == nil || IsNil(o.JobTimeout) {
		var ret int32
		return ret
	}
	return *o.JobTimeout
}

// GetJobTimeoutOk returns a tuple with the JobTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtRunJobItem) GetJobTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.JobTimeout) {
		return nil, false
	}
	return o.JobTimeout, true
}

// HasJobTimeout returns a boolean if a field has been set.
func (o *VhtRunJobItem) HasJobTimeout() bool {
	if o != nil && !IsNil(o.JobTimeout) {
		return true
	}

	return false
}

// SetJobTimeout gets a reference to the given int32 and assigns it to the JobTimeout field.
func (o *VhtRunJobItem) SetJobTimeout(v int32) {
	o.JobTimeout = &v
}

// GetName returns the Name field value
func (o *VhtRunJobItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VhtRunJobItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VhtRunJobItem) SetName(v string) {
	o.Name = v
}

// GetQueued returns the Queued field value
func (o *VhtRunJobItem) GetQueued() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Queued
}

// GetQueuedOk returns a tuple with the Queued field value
// and a boolean to check if the value has been set.
func (o *VhtRunJobItem) GetQueuedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queued, true
}

// SetQueued sets field value
func (o *VhtRunJobItem) SetQueued(v bool) {
	o.Queued = v
}

// GetStatus returns the Status field value
func (o *VhtRunJobItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VhtRunJobItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VhtRunJobItem) SetStatus(v string) {
	o.Status = v
}

// GetSystemError returns the SystemError field value
func (o *VhtRunJobItem) GetSystemError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SystemError
}

// GetSystemErrorOk returns a tuple with the SystemError field value
// and a boolean to check if the value has been set.
func (o *VhtRunJobItem) GetSystemErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemError, true
}

// SetSystemError sets field value
func (o *VhtRunJobItem) SetSystemError(v bool) {
	o.SystemError = v
}

// GetTimedOut returns the TimedOut field value
func (o *VhtRunJobItem) GetTimedOut() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TimedOut
}

// GetTimedOutOk returns a tuple with the TimedOut field value
// and a boolean to check if the value has been set.
func (o *VhtRunJobItem) GetTimedOutOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimedOut, true
}

// SetTimedOut sets field value
func (o *VhtRunJobItem) SetTimedOut(v bool) {
	o.TimedOut = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VhtRunJobItem) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VhtRunJobItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *VhtRunJobItem) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *VhtRunJobItem) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *VhtRunJobItem) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *VhtRunJobItem) UnsetTitle() {
	o.Title.Unset()
}

func (o VhtRunJobItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VhtRunJobItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links.Get()
	toSerialize["_metadata"] = o.Metadata.Get()
	toSerialize["done"] = o.Done
	toSerialize["error"] = o.Error
	if !IsNil(o.JobTimeout) {
		toSerialize["jobTimeout"] = o.JobTimeout
	}
	toSerialize["name"] = o.Name
	toSerialize["queued"] = o.Queued
	toSerialize["status"] = o.Status
	toSerialize["systemError"] = o.SystemError
	toSerialize["timedOut"] = o.TimedOut
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	return toSerialize, nil
}

func (o *VhtRunJobItem) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_links",
		"_metadata",
		"done",
		"error",
		"name",
		"queued",
		"status",
		"systemError",
		"timedOut",
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

	varVhtRunJobItem := _VhtRunJobItem{}

	err = json.Unmarshal(bytes, &varVhtRunJobItem)

	if err != nil {
		return err
	}

	*o = VhtRunJobItem(varVhtRunJobItem)

	return err
}

type NullableVhtRunJobItem struct {
	value *VhtRunJobItem
	isSet bool
}

func (v NullableVhtRunJobItem) Get() *VhtRunJobItem {
	return v.value
}

func (v *NullableVhtRunJobItem) Set(val *VhtRunJobItem) {
	v.value = val
	v.isSet = true
}

func (v NullableVhtRunJobItem) IsSet() bool {
	return v.isSet
}

func (v *NullableVhtRunJobItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVhtRunJobItem(val *VhtRunJobItem) *NullableVhtRunJobItem {
	return &NullableVhtRunJobItem{value: val, isSet: true}
}

func (v NullableVhtRunJobItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVhtRunJobItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


