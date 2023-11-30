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

// checks if the IntellisenseJobItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntellisenseJobItem{}

// IntellisenseJobItem This resource allows an Intellisense Job to be configured when it is created, such as defining the project to handle. When the job is read, it will include the current status of the job and links to other available resources, such as messages and artefacts.
type IntellisenseJobItem struct {
	Links NullableIntellisenseJobItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	// The number of build steps that have been completed so far. Please note: - This value also includes additional service orchestration steps, that are outside the core build process,   so may differ from the build progress indicated within build messages.  - This value will only be available after the build has been started.
	BuildStepsCompleted NullableInt32 `json:"buildStepsCompleted"`
	// The total number of steps that will need to be performed to complete the build. Please note: - This value also includes additional service orchestration steps, that are outside the core build process,   so may differ from the build progress indicated within build messages.  - This value will only be available after the build has been started.
	BuildStepsTotal NullableInt32 `json:"buildStepsTotal"`
	// Build context for jobs that require it.
	Context NullableString `json:"context,omitempty"`
	// True when the job has completed (this does necessarily indicate success).
	Done bool `json:"done"`
	// True if there was an error in the build service while attempting the job.
	Error bool `json:"error"`
	// True if the compilation database generation failed (this should be used in conjunction with the `done` property).
	Failure bool `json:"failure"`
	// The maximum time (in seconds) that the job will be allowed to run. After the timeout has expired the job will be aborted and reported as a failure. The timeout does not include any time the request spent being queued, waiting for the job to be started.
	JobTimeout *int32 `json:"jobTimeout,omitempty"`
	// Unique ID of the Intellisense Job.
	Name string `json:"name"`
	// Path to packs repository to replace value in compilation database.
	Packs string `json:"packs"`
	// CMSIS project to handle or being handled.
	Project string `json:"project"`
	// True if job is currently queued and waiting to be processed. Otherwise, the job is either currently being processed or ended.
	Queued *bool `json:"queued,omitempty"`
	// A summary status of the job. Note: this value should not be relied upon to determine whether a job has completed, succeeded or failed as this list may change as state machine evolves. Use resource appropriate flags instead.
	Status string `json:"status"`
	// True if the job was successful (this should be used in conjunction with the `done` property).
	Success bool `json:"success"`
	// Optional human readable name of the CMSIS Intellisense job.
	Title NullableString `json:"title,omitempty"`
	// Path to toolchain binaries to replace value in compilation database.
	Toolchain string `json:"toolchain"`
	// Path to toolchain headers to replace value in compilation database.
	ToolchainHeaders string `json:"toolchainHeaders"`
	// Path to user's workspace to replace value in compilation database.
	Workspace string `json:"workspace"`
}

type _IntellisenseJobItem IntellisenseJobItem

// NewIntellisenseJobItem instantiates a new IntellisenseJobItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntellisenseJobItem(links NullableIntellisenseJobItemLinks, metadata NullableCommonMetadata, buildStepsCompleted NullableInt32, buildStepsTotal NullableInt32, done bool, error_ bool, failure bool, name string, packs string, project string, status string, success bool, toolchain string, toolchainHeaders string, workspace string) *IntellisenseJobItem {
	this := IntellisenseJobItem{}
	this.Links = links
	this.Metadata = metadata
	this.BuildStepsCompleted = buildStepsCompleted
	this.BuildStepsTotal = buildStepsTotal
	this.Done = done
	this.Error = error_
	this.Failure = failure
	var jobTimeout int32 = 300
	this.JobTimeout = &jobTimeout
	this.Name = name
	this.Packs = packs
	this.Project = project
	this.Status = status
	this.Success = success
	this.Toolchain = toolchain
	this.ToolchainHeaders = toolchainHeaders
	this.Workspace = workspace
	return &this
}

// NewIntellisenseJobItemWithDefaults instantiates a new IntellisenseJobItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntellisenseJobItemWithDefaults() *IntellisenseJobItem {
	this := IntellisenseJobItem{}
	var jobTimeout int32 = 300
	this.JobTimeout = &jobTimeout
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for IntellisenseJobItemLinks will be returned
func (o *IntellisenseJobItem) GetLinks() IntellisenseJobItemLinks {
	if o == nil || o.Links.Get() == nil {
		var ret IntellisenseJobItemLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntellisenseJobItem) GetLinksOk() (*IntellisenseJobItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *IntellisenseJobItem) SetLinks(v IntellisenseJobItemLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *IntellisenseJobItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntellisenseJobItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *IntellisenseJobItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetBuildStepsCompleted returns the BuildStepsCompleted field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *IntellisenseJobItem) GetBuildStepsCompleted() int32 {
	if o == nil || o.BuildStepsCompleted.Get() == nil {
		var ret int32
		return ret
	}

	return *o.BuildStepsCompleted.Get()
}

// GetBuildStepsCompletedOk returns a tuple with the BuildStepsCompleted field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntellisenseJobItem) GetBuildStepsCompletedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildStepsCompleted.Get(), o.BuildStepsCompleted.IsSet()
}

// SetBuildStepsCompleted sets field value
func (o *IntellisenseJobItem) SetBuildStepsCompleted(v int32) {
	o.BuildStepsCompleted.Set(&v)
}

// GetBuildStepsTotal returns the BuildStepsTotal field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *IntellisenseJobItem) GetBuildStepsTotal() int32 {
	if o == nil || o.BuildStepsTotal.Get() == nil {
		var ret int32
		return ret
	}

	return *o.BuildStepsTotal.Get()
}

// GetBuildStepsTotalOk returns a tuple with the BuildStepsTotal field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntellisenseJobItem) GetBuildStepsTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildStepsTotal.Get(), o.BuildStepsTotal.IsSet()
}

// SetBuildStepsTotal sets field value
func (o *IntellisenseJobItem) SetBuildStepsTotal(v int32) {
	o.BuildStepsTotal.Set(&v)
}

// GetContext returns the Context field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntellisenseJobItem) GetContext() string {
	if o == nil || IsNil(o.Context.Get()) {
		var ret string
		return ret
	}
	return *o.Context.Get()
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntellisenseJobItem) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Context.Get(), o.Context.IsSet()
}

// HasContext returns a boolean if a field has been set.
func (o *IntellisenseJobItem) HasContext() bool {
	if o != nil && o.Context.IsSet() {
		return true
	}

	return false
}

// SetContext gets a reference to the given NullableString and assigns it to the Context field.
func (o *IntellisenseJobItem) SetContext(v string) {
	o.Context.Set(&v)
}
// SetContextNil sets the value for Context to be an explicit nil
func (o *IntellisenseJobItem) SetContextNil() {
	o.Context.Set(nil)
}

// UnsetContext ensures that no value is present for Context, not even an explicit nil
func (o *IntellisenseJobItem) UnsetContext() {
	o.Context.Unset()
}

// GetDone returns the Done field value
func (o *IntellisenseJobItem) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *IntellisenseJobItem) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *IntellisenseJobItem) SetDone(v bool) {
	o.Done = v
}

// GetError returns the Error field value
func (o *IntellisenseJobItem) GetError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *IntellisenseJobItem) GetErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *IntellisenseJobItem) SetError(v bool) {
	o.Error = v
}

// GetFailure returns the Failure field value
func (o *IntellisenseJobItem) GetFailure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Failure
}

// GetFailureOk returns a tuple with the Failure field value
// and a boolean to check if the value has been set.
func (o *IntellisenseJobItem) GetFailureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failure, true
}

// SetFailure sets field value
func (o *IntellisenseJobItem) SetFailure(v bool) {
	o.Failure = v
}

// GetJobTimeout returns the JobTimeout field value if set, zero value otherwise.
func (o *IntellisenseJobItem) GetJobTimeout() int32 {
	if o == nil || IsNil(o.JobTimeout) {
		var ret int32
		return ret
	}
	return *o.JobTimeout
}

// GetJobTimeoutOk returns a tuple with the JobTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntellisenseJobItem) GetJobTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.JobTimeout) {
		return nil, false
	}
	return o.JobTimeout, true
}

// HasJobTimeout returns a boolean if a field has been set.
func (o *IntellisenseJobItem) HasJobTimeout() bool {
	if o != nil && !IsNil(o.JobTimeout) {
		return true
	}

	return false
}

// SetJobTimeout gets a reference to the given int32 and assigns it to the JobTimeout field.
func (o *IntellisenseJobItem) SetJobTimeout(v int32) {
	o.JobTimeout = &v
}

// GetName returns the Name field value
func (o *IntellisenseJobItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IntellisenseJobItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IntellisenseJobItem) SetName(v string) {
	o.Name = v
}

// GetPacks returns the Packs field value
func (o *IntellisenseJobItem) GetPacks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Packs
}

// GetPacksOk returns a tuple with the Packs field value
// and a boolean to check if the value has been set.
func (o *IntellisenseJobItem) GetPacksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Packs, true
}

// SetPacks sets field value
func (o *IntellisenseJobItem) SetPacks(v string) {
	o.Packs = v
}

// GetProject returns the Project field value
func (o *IntellisenseJobItem) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *IntellisenseJobItem) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *IntellisenseJobItem) SetProject(v string) {
	o.Project = v
}

// GetQueued returns the Queued field value if set, zero value otherwise.
func (o *IntellisenseJobItem) GetQueued() bool {
	if o == nil || IsNil(o.Queued) {
		var ret bool
		return ret
	}
	return *o.Queued
}

// GetQueuedOk returns a tuple with the Queued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntellisenseJobItem) GetQueuedOk() (*bool, bool) {
	if o == nil || IsNil(o.Queued) {
		return nil, false
	}
	return o.Queued, true
}

// HasQueued returns a boolean if a field has been set.
func (o *IntellisenseJobItem) HasQueued() bool {
	if o != nil && !IsNil(o.Queued) {
		return true
	}

	return false
}

// SetQueued gets a reference to the given bool and assigns it to the Queued field.
func (o *IntellisenseJobItem) SetQueued(v bool) {
	o.Queued = &v
}

// GetStatus returns the Status field value
func (o *IntellisenseJobItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *IntellisenseJobItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *IntellisenseJobItem) SetStatus(v string) {
	o.Status = v
}

// GetSuccess returns the Success field value
func (o *IntellisenseJobItem) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *IntellisenseJobItem) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *IntellisenseJobItem) SetSuccess(v bool) {
	o.Success = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntellisenseJobItem) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntellisenseJobItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *IntellisenseJobItem) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *IntellisenseJobItem) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *IntellisenseJobItem) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *IntellisenseJobItem) UnsetTitle() {
	o.Title.Unset()
}

// GetToolchain returns the Toolchain field value
func (o *IntellisenseJobItem) GetToolchain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Toolchain
}

// GetToolchainOk returns a tuple with the Toolchain field value
// and a boolean to check if the value has been set.
func (o *IntellisenseJobItem) GetToolchainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Toolchain, true
}

// SetToolchain sets field value
func (o *IntellisenseJobItem) SetToolchain(v string) {
	o.Toolchain = v
}

// GetToolchainHeaders returns the ToolchainHeaders field value
func (o *IntellisenseJobItem) GetToolchainHeaders() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToolchainHeaders
}

// GetToolchainHeadersOk returns a tuple with the ToolchainHeaders field value
// and a boolean to check if the value has been set.
func (o *IntellisenseJobItem) GetToolchainHeadersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolchainHeaders, true
}

// SetToolchainHeaders sets field value
func (o *IntellisenseJobItem) SetToolchainHeaders(v string) {
	o.ToolchainHeaders = v
}

// GetWorkspace returns the Workspace field value
func (o *IntellisenseJobItem) GetWorkspace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value
// and a boolean to check if the value has been set.
func (o *IntellisenseJobItem) GetWorkspaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workspace, true
}

// SetWorkspace sets field value
func (o *IntellisenseJobItem) SetWorkspace(v string) {
	o.Workspace = v
}

func (o IntellisenseJobItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntellisenseJobItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links.Get()
	toSerialize["_metadata"] = o.Metadata.Get()
	toSerialize["buildStepsCompleted"] = o.BuildStepsCompleted.Get()
	toSerialize["buildStepsTotal"] = o.BuildStepsTotal.Get()
	if o.Context.IsSet() {
		toSerialize["context"] = o.Context.Get()
	}
	toSerialize["done"] = o.Done
	toSerialize["error"] = o.Error
	toSerialize["failure"] = o.Failure
	if !IsNil(o.JobTimeout) {
		toSerialize["jobTimeout"] = o.JobTimeout
	}
	toSerialize["name"] = o.Name
	toSerialize["packs"] = o.Packs
	toSerialize["project"] = o.Project
	if !IsNil(o.Queued) {
		toSerialize["queued"] = o.Queued
	}
	toSerialize["status"] = o.Status
	toSerialize["success"] = o.Success
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	toSerialize["toolchain"] = o.Toolchain
	toSerialize["toolchainHeaders"] = o.ToolchainHeaders
	toSerialize["workspace"] = o.Workspace
	return toSerialize, nil
}

func (o *IntellisenseJobItem) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_links",
		"_metadata",
		"buildStepsCompleted",
		"buildStepsTotal",
		"done",
		"error",
		"failure",
		"name",
		"packs",
		"project",
		"status",
		"success",
		"toolchain",
		"toolchainHeaders",
		"workspace",
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

	varIntellisenseJobItem := _IntellisenseJobItem{}

	err = json.Unmarshal(bytes, &varIntellisenseJobItem)

	if err != nil {
		return err
	}

	*o = IntellisenseJobItem(varIntellisenseJobItem)

	return err
}

type NullableIntellisenseJobItem struct {
	value *IntellisenseJobItem
	isSet bool
}

func (v NullableIntellisenseJobItem) Get() *IntellisenseJobItem {
	return v.value
}

func (v *NullableIntellisenseJobItem) Set(val *IntellisenseJobItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIntellisenseJobItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIntellisenseJobItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntellisenseJobItem(val *IntellisenseJobItem) *NullableIntellisenseJobItem {
	return &NullableIntellisenseJobItem{value: val, isSet: true}
}

func (v NullableIntellisenseJobItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntellisenseJobItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


