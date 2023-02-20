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

// checks if the WorkspaceDetailsItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkspaceDetailsItem{}

// WorkspaceDetailsItem This is providing more details about the workspace
type WorkspaceDetailsItem struct {
	Links NullableWorkspaceDetailsItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	// Unique ID of the Workspace details.
	Name string `json:"name"`
	// relative path of the workspace
	Path string `json:"path"`
	// size on disk (in Byte) used by this workspace
	SizeInByte *int64 `json:"sizeInByte,omitempty"`
	// Optional human readable name of the Workspace details.
	Title NullableString `json:"title,omitempty"`
}

// NewWorkspaceDetailsItem instantiates a new WorkspaceDetailsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceDetailsItem(links NullableWorkspaceDetailsItemLinks, metadata NullableCommonMetadata, name string, path string) *WorkspaceDetailsItem {
	this := WorkspaceDetailsItem{}
	this.Links = links
	this.Metadata = metadata
	this.Name = name
	this.Path = path
	return &this
}

// NewWorkspaceDetailsItemWithDefaults instantiates a new WorkspaceDetailsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceDetailsItemWithDefaults() *WorkspaceDetailsItem {
	this := WorkspaceDetailsItem{}
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for WorkspaceDetailsItemLinks will be returned
func (o *WorkspaceDetailsItem) GetLinks() WorkspaceDetailsItemLinks {
	if o == nil || o.Links.Get() == nil {
		var ret WorkspaceDetailsItemLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkspaceDetailsItem) GetLinksOk() (*WorkspaceDetailsItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *WorkspaceDetailsItem) SetLinks(v WorkspaceDetailsItemLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *WorkspaceDetailsItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkspaceDetailsItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *WorkspaceDetailsItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetName returns the Name field value
func (o *WorkspaceDetailsItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkspaceDetailsItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkspaceDetailsItem) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value
func (o *WorkspaceDetailsItem) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *WorkspaceDetailsItem) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *WorkspaceDetailsItem) SetPath(v string) {
	o.Path = v
}

// GetSizeInByte returns the SizeInByte field value if set, zero value otherwise.
func (o *WorkspaceDetailsItem) GetSizeInByte() int64 {
	if o == nil || IsNil(o.SizeInByte) {
		var ret int64
		return ret
	}
	return *o.SizeInByte
}

// GetSizeInByteOk returns a tuple with the SizeInByte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceDetailsItem) GetSizeInByteOk() (*int64, bool) {
	if o == nil || IsNil(o.SizeInByte) {
		return nil, false
	}
	return o.SizeInByte, true
}

// HasSizeInByte returns a boolean if a field has been set.
func (o *WorkspaceDetailsItem) HasSizeInByte() bool {
	if o != nil && !IsNil(o.SizeInByte) {
		return true
	}

	return false
}

// SetSizeInByte gets a reference to the given int64 and assigns it to the SizeInByte field.
func (o *WorkspaceDetailsItem) SetSizeInByte(v int64) {
	o.SizeInByte = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkspaceDetailsItem) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkspaceDetailsItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkspaceDetailsItem) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *WorkspaceDetailsItem) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *WorkspaceDetailsItem) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *WorkspaceDetailsItem) UnsetTitle() {
	o.Title.Unset()
}

func (o WorkspaceDetailsItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkspaceDetailsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links.Get()
	toSerialize["_metadata"] = o.Metadata.Get()
	// skip: name is readOnly
	// skip: path is readOnly
	// skip: sizeInByte is readOnly
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	return toSerialize, nil
}

type NullableWorkspaceDetailsItem struct {
	value *WorkspaceDetailsItem
	isSet bool
}

func (v NullableWorkspaceDetailsItem) Get() *WorkspaceDetailsItem {
	return v.value
}

func (v *NullableWorkspaceDetailsItem) Set(val *WorkspaceDetailsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceDetailsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceDetailsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceDetailsItem(val *WorkspaceDetailsItem) *NullableWorkspaceDetailsItem {
	return &NullableWorkspaceDetailsItem{value: val, isSet: true}
}

func (v NullableWorkspaceDetailsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceDetailsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


