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

// checks if the WorkspaceRepositoryContentManager type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkspaceRepositoryContentManager{}

// WorkspaceRepositoryContentManager workspace controller resource when the content of the workspace is defined in a repository
type WorkspaceRepositoryContentManager struct {
	Links NullableWorkspaceRepositoryContentManagerLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	// Number of commits to fetch. 0 indicates all history for all branches and tags
	FetchDepth NullableInt32 `json:"fetchDepth,omitempty"`
	// ID of the repository.
	Name string `json:"name"`
	// The branch, tag or SHA to checkout
	Reference NullableString `json:"reference"`
	// repository URL to clone it from
	RepositoryUrl string `json:"repositoryUrl"`
	// Source Control Management system to use
	Scm string `json:"scm"`
	// Whether to checkout submodules: `true` to checkout submodules recursively
	Submodules *bool `json:"submodules,omitempty"`
	// Optional human readable name of the repository.
	Title NullableString `json:"title,omitempty"`
}

// NewWorkspaceRepositoryContentManager instantiates a new WorkspaceRepositoryContentManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceRepositoryContentManager(links NullableWorkspaceRepositoryContentManagerLinks, metadata NullableCommonMetadata, name string, reference NullableString, repositoryUrl string, scm string) *WorkspaceRepositoryContentManager {
	this := WorkspaceRepositoryContentManager{}
	this.Links = links
	this.Metadata = metadata
	var fetchDepth int32 = 1
	this.FetchDepth = *NewNullableInt32(&fetchDepth)
	this.Name = name
	this.Reference = reference
	this.RepositoryUrl = repositoryUrl
	this.Scm = scm
	var submodules bool = true
	this.Submodules = &submodules
	return &this
}

// NewWorkspaceRepositoryContentManagerWithDefaults instantiates a new WorkspaceRepositoryContentManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceRepositoryContentManagerWithDefaults() *WorkspaceRepositoryContentManager {
	this := WorkspaceRepositoryContentManager{}
	var fetchDepth int32 = 1
	this.FetchDepth = *NewNullableInt32(&fetchDepth)
	var submodules bool = true
	this.Submodules = &submodules
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for WorkspaceRepositoryContentManagerLinks will be returned
func (o *WorkspaceRepositoryContentManager) GetLinks() WorkspaceRepositoryContentManagerLinks {
	if o == nil || o.Links.Get() == nil {
		var ret WorkspaceRepositoryContentManagerLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkspaceRepositoryContentManager) GetLinksOk() (*WorkspaceRepositoryContentManagerLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *WorkspaceRepositoryContentManager) SetLinks(v WorkspaceRepositoryContentManagerLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *WorkspaceRepositoryContentManager) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkspaceRepositoryContentManager) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *WorkspaceRepositoryContentManager) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetFetchDepth returns the FetchDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkspaceRepositoryContentManager) GetFetchDepth() int32 {
	if o == nil || IsNil(o.FetchDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.FetchDepth.Get()
}

// GetFetchDepthOk returns a tuple with the FetchDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkspaceRepositoryContentManager) GetFetchDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FetchDepth.Get(), o.FetchDepth.IsSet()
}

// HasFetchDepth returns a boolean if a field has been set.
func (o *WorkspaceRepositoryContentManager) HasFetchDepth() bool {
	if o != nil && o.FetchDepth.IsSet() {
		return true
	}

	return false
}

// SetFetchDepth gets a reference to the given NullableInt32 and assigns it to the FetchDepth field.
func (o *WorkspaceRepositoryContentManager) SetFetchDepth(v int32) {
	o.FetchDepth.Set(&v)
}
// SetFetchDepthNil sets the value for FetchDepth to be an explicit nil
func (o *WorkspaceRepositoryContentManager) SetFetchDepthNil() {
	o.FetchDepth.Set(nil)
}

// UnsetFetchDepth ensures that no value is present for FetchDepth, not even an explicit nil
func (o *WorkspaceRepositoryContentManager) UnsetFetchDepth() {
	o.FetchDepth.Unset()
}

// GetName returns the Name field value
func (o *WorkspaceRepositoryContentManager) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkspaceRepositoryContentManager) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkspaceRepositoryContentManager) SetName(v string) {
	o.Name = v
}

// GetReference returns the Reference field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WorkspaceRepositoryContentManager) GetReference() string {
	if o == nil || o.Reference.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reference.Get()
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkspaceRepositoryContentManager) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reference.Get(), o.Reference.IsSet()
}

// SetReference sets field value
func (o *WorkspaceRepositoryContentManager) SetReference(v string) {
	o.Reference.Set(&v)
}

// GetRepositoryUrl returns the RepositoryUrl field value
func (o *WorkspaceRepositoryContentManager) GetRepositoryUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value
// and a boolean to check if the value has been set.
func (o *WorkspaceRepositoryContentManager) GetRepositoryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryUrl, true
}

// SetRepositoryUrl sets field value
func (o *WorkspaceRepositoryContentManager) SetRepositoryUrl(v string) {
	o.RepositoryUrl = v
}

// GetScm returns the Scm field value
func (o *WorkspaceRepositoryContentManager) GetScm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scm
}

// GetScmOk returns a tuple with the Scm field value
// and a boolean to check if the value has been set.
func (o *WorkspaceRepositoryContentManager) GetScmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scm, true
}

// SetScm sets field value
func (o *WorkspaceRepositoryContentManager) SetScm(v string) {
	o.Scm = v
}

// GetSubmodules returns the Submodules field value if set, zero value otherwise.
func (o *WorkspaceRepositoryContentManager) GetSubmodules() bool {
	if o == nil || IsNil(o.Submodules) {
		var ret bool
		return ret
	}
	return *o.Submodules
}

// GetSubmodulesOk returns a tuple with the Submodules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceRepositoryContentManager) GetSubmodulesOk() (*bool, bool) {
	if o == nil || IsNil(o.Submodules) {
		return nil, false
	}
	return o.Submodules, true
}

// HasSubmodules returns a boolean if a field has been set.
func (o *WorkspaceRepositoryContentManager) HasSubmodules() bool {
	if o != nil && !IsNil(o.Submodules) {
		return true
	}

	return false
}

// SetSubmodules gets a reference to the given bool and assigns it to the Submodules field.
func (o *WorkspaceRepositoryContentManager) SetSubmodules(v bool) {
	o.Submodules = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkspaceRepositoryContentManager) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkspaceRepositoryContentManager) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkspaceRepositoryContentManager) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *WorkspaceRepositoryContentManager) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *WorkspaceRepositoryContentManager) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *WorkspaceRepositoryContentManager) UnsetTitle() {
	o.Title.Unset()
}

func (o WorkspaceRepositoryContentManager) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkspaceRepositoryContentManager) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links.Get()
	toSerialize["_metadata"] = o.Metadata.Get()
	if o.FetchDepth.IsSet() {
		toSerialize["fetchDepth"] = o.FetchDepth.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["reference"] = o.Reference.Get()
	toSerialize["repositoryUrl"] = o.RepositoryUrl
	toSerialize["scm"] = o.Scm
	if !IsNil(o.Submodules) {
		toSerialize["submodules"] = o.Submodules
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	return toSerialize, nil
}

type NullableWorkspaceRepositoryContentManager struct {
	value *WorkspaceRepositoryContentManager
	isSet bool
}

func (v NullableWorkspaceRepositoryContentManager) Get() *WorkspaceRepositoryContentManager {
	return v.value
}

func (v *NullableWorkspaceRepositoryContentManager) Set(val *WorkspaceRepositoryContentManager) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceRepositoryContentManager) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceRepositoryContentManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceRepositoryContentManager(val *WorkspaceRepositoryContentManager) *NullableWorkspaceRepositoryContentManager {
	return &NullableWorkspaceRepositoryContentManager{value: val, isSet: true}
}

func (v NullableWorkspaceRepositoryContentManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceRepositoryContentManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


