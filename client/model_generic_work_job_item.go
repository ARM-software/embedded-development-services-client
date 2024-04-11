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

// checks if the GenericWorkJobItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericWorkJobItem{}

// GenericWorkJobItem This resource allows an generic work job to be configured when it is created, such as defining the project to handle. When the job is read, it will include the current status of the job and links to other available resources, such as messages and artefacts.
type GenericWorkJobItem struct {
	Links NullableGenericWorkJobItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	// Configuration map for jobs that require it. These could be environment variables. This is implementation dependent.
	Configuration map[string]string `json:"configuration,omitempty"`
	// True when the job has completed (this does necessarily indicate success).
	Done bool `json:"done"`
	// True if there was an error in the service while attempting the job.
	Error bool `json:"error"`
	// True if the job failed (this should be used in conjunction with the `done` property).
	Failure bool `json:"failure"`
	// The number of execution steps that have been completed so far. Please note: - This value also includes additional service orchestration steps, that are outside the core job being    executed by the process, so may differ from the build progress indicated within build messages.  - This value will only be available after the job has been started.
	JobStepsCompleted NullableInt32 `json:"jobStepsCompleted"`
	// The total number of steps that will need to be performed to complete the job. Please note: - This value also includes additional service orchestration steps, that are outside the core job being    executed by the process, so may differ from the build progress indicated within build messages.  - This value will only be available after the job has been started.
	JobStepsTotal NullableInt32 `json:"jobStepsTotal"`
	// The maximum time (in seconds) that the job will be allowed to run. After the timeout has expired the job will be aborted and reported as a failure. The timeout does not include any time the request spent being queued, waiting for the job to be started.
	JobTimeout *int32 `json:"jobTimeout,omitempty"`
	// Unique ID of the Generic Work Job.
	Name string `json:"name"`
	// project to handle or being handled.
	Project *string `json:"project,omitempty"`
	// True if job is currently queued and waiting to be processed. Otherwise, the job is either currently being processed or ended.
	Queued *bool `json:"queued,omitempty"`
	// A summary status of the job. Note: this value should not be relied upon to determine whether a job has completed, succeeded or failed as this list may change as state machine evolves. Use resource appropriate flags instead.
	Status string `json:"status"`
	// True if the job was successful (this should be used in conjunction with the `done` property).
	Success bool `json:"success"`
	// Optional human readable name of the generic work job.
	Title NullableString `json:"title,omitempty"`
	// Path to user's workspace.
	Workspace *string `json:"workspace,omitempty"`
}

type _GenericWorkJobItem GenericWorkJobItem

// NewGenericWorkJobItem instantiates a new GenericWorkJobItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericWorkJobItem(links NullableGenericWorkJobItemLinks, metadata NullableCommonMetadata, done bool, error_ bool, failure bool, jobStepsCompleted NullableInt32, jobStepsTotal NullableInt32, name string, status string, success bool) *GenericWorkJobItem {
	this := GenericWorkJobItem{}
	this.Links = links
	this.Metadata = metadata
	this.Done = done
	this.Error = error_
	this.Failure = failure
	this.JobStepsCompleted = jobStepsCompleted
	this.JobStepsTotal = jobStepsTotal
	var jobTimeout int32 = 300
	this.JobTimeout = &jobTimeout
	this.Name = name
	this.Status = status
	this.Success = success
	return &this
}

// NewGenericWorkJobItemWithDefaults instantiates a new GenericWorkJobItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericWorkJobItemWithDefaults() *GenericWorkJobItem {
	this := GenericWorkJobItem{}
	var jobTimeout int32 = 300
	this.JobTimeout = &jobTimeout
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for GenericWorkJobItemLinks will be returned
func (o *GenericWorkJobItem) GetLinks() GenericWorkJobItemLinks {
	if o == nil || o.Links.Get() == nil {
		var ret GenericWorkJobItemLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericWorkJobItem) GetLinksOk() (*GenericWorkJobItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *GenericWorkJobItem) SetLinks(v GenericWorkJobItemLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *GenericWorkJobItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericWorkJobItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *GenericWorkJobItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenericWorkJobItem) GetConfiguration() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericWorkJobItem) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return &o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *GenericWorkJobItem) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *GenericWorkJobItem) SetConfiguration(v map[string]string) {
	o.Configuration = v
}

// GetDone returns the Done field value
func (o *GenericWorkJobItem) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *GenericWorkJobItem) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *GenericWorkJobItem) SetDone(v bool) {
	o.Done = v
}

// GetError returns the Error field value
func (o *GenericWorkJobItem) GetError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *GenericWorkJobItem) GetErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *GenericWorkJobItem) SetError(v bool) {
	o.Error = v
}

// GetFailure returns the Failure field value
func (o *GenericWorkJobItem) GetFailure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Failure
}

// GetFailureOk returns a tuple with the Failure field value
// and a boolean to check if the value has been set.
func (o *GenericWorkJobItem) GetFailureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failure, true
}

// SetFailure sets field value
func (o *GenericWorkJobItem) SetFailure(v bool) {
	o.Failure = v
}

// GetJobStepsCompleted returns the JobStepsCompleted field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *GenericWorkJobItem) GetJobStepsCompleted() int32 {
	if o == nil || o.JobStepsCompleted.Get() == nil {
		var ret int32
		return ret
	}

	return *o.JobStepsCompleted.Get()
}

// GetJobStepsCompletedOk returns a tuple with the JobStepsCompleted field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericWorkJobItem) GetJobStepsCompletedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobStepsCompleted.Get(), o.JobStepsCompleted.IsSet()
}

// SetJobStepsCompleted sets field value
func (o *GenericWorkJobItem) SetJobStepsCompleted(v int32) {
	o.JobStepsCompleted.Set(&v)
}

// GetJobStepsTotal returns the JobStepsTotal field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *GenericWorkJobItem) GetJobStepsTotal() int32 {
	if o == nil || o.JobStepsTotal.Get() == nil {
		var ret int32
		return ret
	}

	return *o.JobStepsTotal.Get()
}

// GetJobStepsTotalOk returns a tuple with the JobStepsTotal field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericWorkJobItem) GetJobStepsTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobStepsTotal.Get(), o.JobStepsTotal.IsSet()
}

// SetJobStepsTotal sets field value
func (o *GenericWorkJobItem) SetJobStepsTotal(v int32) {
	o.JobStepsTotal.Set(&v)
}

// GetJobTimeout returns the JobTimeout field value if set, zero value otherwise.
func (o *GenericWorkJobItem) GetJobTimeout() int32 {
	if o == nil || IsNil(o.JobTimeout) {
		var ret int32
		return ret
	}
	return *o.JobTimeout
}

// GetJobTimeoutOk returns a tuple with the JobTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericWorkJobItem) GetJobTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.JobTimeout) {
		return nil, false
	}
	return o.JobTimeout, true
}

// HasJobTimeout returns a boolean if a field has been set.
func (o *GenericWorkJobItem) HasJobTimeout() bool {
	if o != nil && !IsNil(o.JobTimeout) {
		return true
	}

	return false
}

// SetJobTimeout gets a reference to the given int32 and assigns it to the JobTimeout field.
func (o *GenericWorkJobItem) SetJobTimeout(v int32) {
	o.JobTimeout = &v
}

// GetName returns the Name field value
func (o *GenericWorkJobItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GenericWorkJobItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GenericWorkJobItem) SetName(v string) {
	o.Name = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *GenericWorkJobItem) GetProject() string {
	if o == nil || IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericWorkJobItem) GetProjectOk() (*string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *GenericWorkJobItem) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *GenericWorkJobItem) SetProject(v string) {
	o.Project = &v
}

// GetQueued returns the Queued field value if set, zero value otherwise.
func (o *GenericWorkJobItem) GetQueued() bool {
	if o == nil || IsNil(o.Queued) {
		var ret bool
		return ret
	}
	return *o.Queued
}

// GetQueuedOk returns a tuple with the Queued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericWorkJobItem) GetQueuedOk() (*bool, bool) {
	if o == nil || IsNil(o.Queued) {
		return nil, false
	}
	return o.Queued, true
}

// HasQueued returns a boolean if a field has been set.
func (o *GenericWorkJobItem) HasQueued() bool {
	if o != nil && !IsNil(o.Queued) {
		return true
	}

	return false
}

// SetQueued gets a reference to the given bool and assigns it to the Queued field.
func (o *GenericWorkJobItem) SetQueued(v bool) {
	o.Queued = &v
}

// GetStatus returns the Status field value
func (o *GenericWorkJobItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GenericWorkJobItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GenericWorkJobItem) SetStatus(v string) {
	o.Status = v
}

// GetSuccess returns the Success field value
func (o *GenericWorkJobItem) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *GenericWorkJobItem) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *GenericWorkJobItem) SetSuccess(v bool) {
	o.Success = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenericWorkJobItem) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericWorkJobItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *GenericWorkJobItem) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *GenericWorkJobItem) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *GenericWorkJobItem) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *GenericWorkJobItem) UnsetTitle() {
	o.Title.Unset()
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *GenericWorkJobItem) GetWorkspace() string {
	if o == nil || IsNil(o.Workspace) {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericWorkJobItem) GetWorkspaceOk() (*string, bool) {
	if o == nil || IsNil(o.Workspace) {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *GenericWorkJobItem) HasWorkspace() bool {
	if o != nil && !IsNil(o.Workspace) {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *GenericWorkJobItem) SetWorkspace(v string) {
	o.Workspace = &v
}

func (o GenericWorkJobItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericWorkJobItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links.Get()
	toSerialize["_metadata"] = o.Metadata.Get()
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	toSerialize["done"] = o.Done
	toSerialize["error"] = o.Error
	toSerialize["failure"] = o.Failure
	toSerialize["jobStepsCompleted"] = o.JobStepsCompleted.Get()
	toSerialize["jobStepsTotal"] = o.JobStepsTotal.Get()
	if !IsNil(o.JobTimeout) {
		toSerialize["jobTimeout"] = o.JobTimeout
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Queued) {
		toSerialize["queued"] = o.Queued
	}
	toSerialize["status"] = o.Status
	toSerialize["success"] = o.Success
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.Workspace) {
		toSerialize["workspace"] = o.Workspace
	}
	return toSerialize, nil
}

func (o *GenericWorkJobItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_links",
		"_metadata",
		"done",
		"error",
		"failure",
		"jobStepsCompleted",
		"jobStepsTotal",
		"name",
		"status",
		"success",
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

	varGenericWorkJobItem := _GenericWorkJobItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGenericWorkJobItem)

	if err != nil {
		return err
	}

	*o = GenericWorkJobItem(varGenericWorkJobItem)

	return err
}

type NullableGenericWorkJobItem struct {
	value *GenericWorkJobItem
	isSet bool
}

func (v NullableGenericWorkJobItem) Get() *GenericWorkJobItem {
	return v.value
}

func (v *NullableGenericWorkJobItem) Set(val *GenericWorkJobItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericWorkJobItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericWorkJobItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericWorkJobItem(val *GenericWorkJobItem) *NullableGenericWorkJobItem {
	return &NullableGenericWorkJobItem{value: val, isSet: true}
}

func (v NullableGenericWorkJobItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericWorkJobItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


