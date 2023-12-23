/*
 * Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
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

// checks if the BuildJobItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildJobItem{}

// BuildJobItem This resource allows a Build Job to be configured when it is created, such as defining the project to build. When the build job is read, it will include the current status of the build and links to other available resources, such as build messages and build artefacts.
type BuildJobItem struct {
	Links NullableBuildJobItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	// The number of build steps that have been completed so far. Please note: - This value also includes additional service orchestration steps, that are outside the core build process,   so may differ from the build progress indicated within build messages. - This value will only be available after the build has been started.
	BuildStepsCompleted NullableInt32 `json:"buildStepsCompleted"`
	// The total number of steps that will need to be performed to complete the build. Please note: - This value also includes additional service orchestration steps, that are outside the core build process,   so may differ from the build progress indicated within build messages. - This value will only be available after the build has been started.
	BuildStepsTotal NullableInt32 `json:"buildStepsTotal"`
	// The maximum time (in seconds) that the build will be allowed to run. After the timeout has expired the build will be aborted and reported as a failure. The timeout does not include any time the request spent being queued, waiting for the build to be started.
	BuildTimeout *int32 `json:"buildTimeout,omitempty"`
	// Whether to run a clean build.
	CleanBuild *bool `json:"cleanBuild,omitempty"`
	// Build context for jobs that require it.
	Context NullableString `json:"context,omitempty"`
	// True when the build job has completed (this does necessarily indicate success).
	Done bool `json:"done"`
	// True if there was an error in the build service while attempting the build.
	Error bool `json:"error"`
	// True if the compilation or link failed (this should be used in conjunction with the `done` property).
	Failure bool `json:"failure"`
	// Unique ID of the Build Job.
	Name string `json:"name"`
	// CMSIS project to build or being built.
	Project string `json:"project"`
	// True if job is currently queued and waiting to be processed. Otherwise, the job is either currently being processed or ended.
	Queued *bool `json:"queued,omitempty"`
	// A summary status of the job. Note: this value should not be relied upon to determine whether a job has completed, succeeded or failed as this list may change as state machine evolves. Use resource appropriate flags instead.
	Status string `json:"status"`
	// True if the build job was successful (this should be used in conjunction with the `done` property).
	Success bool `json:"success"`
	// Optional human readable name of the CMSIS build job.
	Title NullableString `json:"title,omitempty"`
	// Workspace name where the CMSIS project is present. If not set, the default user's workspace will be used.
	Workspace NullableString `json:"workspace,omitempty"`
}

type _BuildJobItem BuildJobItem

// NewBuildJobItem instantiates a new BuildJobItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildJobItem(links NullableBuildJobItemLinks, metadata NullableCommonMetadata, buildStepsCompleted NullableInt32, buildStepsTotal NullableInt32, done bool, error_ bool, failure bool, name string, project string, status string, success bool) *BuildJobItem {
	this := BuildJobItem{}
	this.Links = links
	this.Metadata = metadata
	this.BuildStepsCompleted = buildStepsCompleted
	this.BuildStepsTotal = buildStepsTotal
	var buildTimeout int32 = 300
	this.BuildTimeout = &buildTimeout
	var cleanBuild bool = false
	this.CleanBuild = &cleanBuild
	this.Done = done
	this.Error = error_
	this.Failure = failure
	this.Name = name
	this.Project = project
	this.Status = status
	this.Success = success
	return &this
}

// NewBuildJobItemWithDefaults instantiates a new BuildJobItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildJobItemWithDefaults() *BuildJobItem {
	this := BuildJobItem{}
	var buildTimeout int32 = 300
	this.BuildTimeout = &buildTimeout
	var cleanBuild bool = false
	this.CleanBuild = &cleanBuild
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for BuildJobItemLinks will be returned
func (o *BuildJobItem) GetLinks() BuildJobItemLinks {
	if o == nil || o.Links.Get() == nil {
		var ret BuildJobItemLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuildJobItem) GetLinksOk() (*BuildJobItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *BuildJobItem) SetLinks(v BuildJobItemLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *BuildJobItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuildJobItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *BuildJobItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetBuildStepsCompleted returns the BuildStepsCompleted field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *BuildJobItem) GetBuildStepsCompleted() int32 {
	if o == nil || o.BuildStepsCompleted.Get() == nil {
		var ret int32
		return ret
	}

	return *o.BuildStepsCompleted.Get()
}

// GetBuildStepsCompletedOk returns a tuple with the BuildStepsCompleted field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuildJobItem) GetBuildStepsCompletedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildStepsCompleted.Get(), o.BuildStepsCompleted.IsSet()
}

// SetBuildStepsCompleted sets field value
func (o *BuildJobItem) SetBuildStepsCompleted(v int32) {
	o.BuildStepsCompleted.Set(&v)
}

// GetBuildStepsTotal returns the BuildStepsTotal field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *BuildJobItem) GetBuildStepsTotal() int32 {
	if o == nil || o.BuildStepsTotal.Get() == nil {
		var ret int32
		return ret
	}

	return *o.BuildStepsTotal.Get()
}

// GetBuildStepsTotalOk returns a tuple with the BuildStepsTotal field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuildJobItem) GetBuildStepsTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildStepsTotal.Get(), o.BuildStepsTotal.IsSet()
}

// SetBuildStepsTotal sets field value
func (o *BuildJobItem) SetBuildStepsTotal(v int32) {
	o.BuildStepsTotal.Set(&v)
}

// GetBuildTimeout returns the BuildTimeout field value if set, zero value otherwise.
func (o *BuildJobItem) GetBuildTimeout() int32 {
	if o == nil || IsNil(o.BuildTimeout) {
		var ret int32
		return ret
	}
	return *o.BuildTimeout
}

// GetBuildTimeoutOk returns a tuple with the BuildTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildJobItem) GetBuildTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.BuildTimeout) {
		return nil, false
	}
	return o.BuildTimeout, true
}

// HasBuildTimeout returns a boolean if a field has been set.
func (o *BuildJobItem) HasBuildTimeout() bool {
	if o != nil && !IsNil(o.BuildTimeout) {
		return true
	}

	return false
}

// SetBuildTimeout gets a reference to the given int32 and assigns it to the BuildTimeout field.
func (o *BuildJobItem) SetBuildTimeout(v int32) {
	o.BuildTimeout = &v
}

// GetCleanBuild returns the CleanBuild field value if set, zero value otherwise.
func (o *BuildJobItem) GetCleanBuild() bool {
	if o == nil || IsNil(o.CleanBuild) {
		var ret bool
		return ret
	}
	return *o.CleanBuild
}

// GetCleanBuildOk returns a tuple with the CleanBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildJobItem) GetCleanBuildOk() (*bool, bool) {
	if o == nil || IsNil(o.CleanBuild) {
		return nil, false
	}
	return o.CleanBuild, true
}

// HasCleanBuild returns a boolean if a field has been set.
func (o *BuildJobItem) HasCleanBuild() bool {
	if o != nil && !IsNil(o.CleanBuild) {
		return true
	}

	return false
}

// SetCleanBuild gets a reference to the given bool and assigns it to the CleanBuild field.
func (o *BuildJobItem) SetCleanBuild(v bool) {
	o.CleanBuild = &v
}

// GetContext returns the Context field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BuildJobItem) GetContext() string {
	if o == nil || IsNil(o.Context.Get()) {
		var ret string
		return ret
	}
	return *o.Context.Get()
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuildJobItem) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Context.Get(), o.Context.IsSet()
}

// HasContext returns a boolean if a field has been set.
func (o *BuildJobItem) HasContext() bool {
	if o != nil && o.Context.IsSet() {
		return true
	}

	return false
}

// SetContext gets a reference to the given NullableString and assigns it to the Context field.
func (o *BuildJobItem) SetContext(v string) {
	o.Context.Set(&v)
}
// SetContextNil sets the value for Context to be an explicit nil
func (o *BuildJobItem) SetContextNil() {
	o.Context.Set(nil)
}

// UnsetContext ensures that no value is present for Context, not even an explicit nil
func (o *BuildJobItem) UnsetContext() {
	o.Context.Unset()
}

// GetDone returns the Done field value
func (o *BuildJobItem) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *BuildJobItem) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *BuildJobItem) SetDone(v bool) {
	o.Done = v
}

// GetError returns the Error field value
func (o *BuildJobItem) GetError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *BuildJobItem) GetErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *BuildJobItem) SetError(v bool) {
	o.Error = v
}

// GetFailure returns the Failure field value
func (o *BuildJobItem) GetFailure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Failure
}

// GetFailureOk returns a tuple with the Failure field value
// and a boolean to check if the value has been set.
func (o *BuildJobItem) GetFailureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failure, true
}

// SetFailure sets field value
func (o *BuildJobItem) SetFailure(v bool) {
	o.Failure = v
}

// GetName returns the Name field value
func (o *BuildJobItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BuildJobItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BuildJobItem) SetName(v string) {
	o.Name = v
}

// GetProject returns the Project field value
func (o *BuildJobItem) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *BuildJobItem) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *BuildJobItem) SetProject(v string) {
	o.Project = v
}

// GetQueued returns the Queued field value if set, zero value otherwise.
func (o *BuildJobItem) GetQueued() bool {
	if o == nil || IsNil(o.Queued) {
		var ret bool
		return ret
	}
	return *o.Queued
}

// GetQueuedOk returns a tuple with the Queued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildJobItem) GetQueuedOk() (*bool, bool) {
	if o == nil || IsNil(o.Queued) {
		return nil, false
	}
	return o.Queued, true
}

// HasQueued returns a boolean if a field has been set.
func (o *BuildJobItem) HasQueued() bool {
	if o != nil && !IsNil(o.Queued) {
		return true
	}

	return false
}

// SetQueued gets a reference to the given bool and assigns it to the Queued field.
func (o *BuildJobItem) SetQueued(v bool) {
	o.Queued = &v
}

// GetStatus returns the Status field value
func (o *BuildJobItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BuildJobItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BuildJobItem) SetStatus(v string) {
	o.Status = v
}

// GetSuccess returns the Success field value
func (o *BuildJobItem) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *BuildJobItem) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *BuildJobItem) SetSuccess(v bool) {
	o.Success = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BuildJobItem) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuildJobItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *BuildJobItem) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *BuildJobItem) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *BuildJobItem) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *BuildJobItem) UnsetTitle() {
	o.Title.Unset()
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BuildJobItem) GetWorkspace() string {
	if o == nil || IsNil(o.Workspace.Get()) {
		var ret string
		return ret
	}
	return *o.Workspace.Get()
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuildJobItem) GetWorkspaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Workspace.Get(), o.Workspace.IsSet()
}

// HasWorkspace returns a boolean if a field has been set.
func (o *BuildJobItem) HasWorkspace() bool {
	if o != nil && o.Workspace.IsSet() {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given NullableString and assigns it to the Workspace field.
func (o *BuildJobItem) SetWorkspace(v string) {
	o.Workspace.Set(&v)
}
// SetWorkspaceNil sets the value for Workspace to be an explicit nil
func (o *BuildJobItem) SetWorkspaceNil() {
	o.Workspace.Set(nil)
}

// UnsetWorkspace ensures that no value is present for Workspace, not even an explicit nil
func (o *BuildJobItem) UnsetWorkspace() {
	o.Workspace.Unset()
}

func (o BuildJobItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildJobItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links.Get()
	toSerialize["_metadata"] = o.Metadata.Get()
	toSerialize["buildStepsCompleted"] = o.BuildStepsCompleted.Get()
	toSerialize["buildStepsTotal"] = o.BuildStepsTotal.Get()
	if !IsNil(o.BuildTimeout) {
		toSerialize["buildTimeout"] = o.BuildTimeout
	}
	if !IsNil(o.CleanBuild) {
		toSerialize["cleanBuild"] = o.CleanBuild
	}
	if o.Context.IsSet() {
		toSerialize["context"] = o.Context.Get()
	}
	toSerialize["done"] = o.Done
	toSerialize["error"] = o.Error
	toSerialize["failure"] = o.Failure
	toSerialize["name"] = o.Name
	toSerialize["project"] = o.Project
	if !IsNil(o.Queued) {
		toSerialize["queued"] = o.Queued
	}
	toSerialize["status"] = o.Status
	toSerialize["success"] = o.Success
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Workspace.IsSet() {
		toSerialize["workspace"] = o.Workspace.Get()
	}
	return toSerialize, nil
}

func (o *BuildJobItem) UnmarshalJSON(data []byte) (err error) {
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
		"project",
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

	varBuildJobItem := _BuildJobItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuildJobItem)

	if err != nil {
		return err
	}

	*o = BuildJobItem(varBuildJobItem)

	return err
}

type NullableBuildJobItem struct {
	value *BuildJobItem
	isSet bool
}

func (v NullableBuildJobItem) Get() *BuildJobItem {
	return v.value
}

func (v *NullableBuildJobItem) Set(val *BuildJobItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildJobItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildJobItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildJobItem(val *BuildJobItem) *NullableBuildJobItem {
	return &NullableBuildJobItem{value: val, isSet: true}
}

func (v NullableBuildJobItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildJobItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


