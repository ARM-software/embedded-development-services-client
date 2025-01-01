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

// checks if the WorkspaceSourceItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkspaceSourceItem{}

// WorkspaceSourceItem struct for WorkspaceSourceItem
type WorkspaceSourceItem struct {
	Links NullableWorkspaceSourceItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	// True if this workspace source is scheduled to be no longer supported by the service.
	Deprecated bool `json:"deprecated"`
	DeprecationInfo *DeprecationInfo `json:"deprecationInfo,omitempty"`
	// More details about this workspace source type.
	Description *string `json:"description,omitempty"`
	// Unique ID of the Workspace Source.
	Name string `json:"name"`
	// Human readable name of the Workspace source.
	Title string `json:"title"`
}

type _WorkspaceSourceItem WorkspaceSourceItem

// NewWorkspaceSourceItem instantiates a new WorkspaceSourceItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceSourceItem(links NullableWorkspaceSourceItemLinks, metadata NullableCommonMetadata, deprecated bool, name string, title string) *WorkspaceSourceItem {
	this := WorkspaceSourceItem{}
	this.Links = links
	this.Metadata = metadata
	this.Deprecated = deprecated
	this.Name = name
	this.Title = title
	return &this
}

// NewWorkspaceSourceItemWithDefaults instantiates a new WorkspaceSourceItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceSourceItemWithDefaults() *WorkspaceSourceItem {
	this := WorkspaceSourceItem{}
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for WorkspaceSourceItemLinks will be returned
func (o *WorkspaceSourceItem) GetLinks() WorkspaceSourceItemLinks {
	if o == nil || o.Links.Get() == nil {
		var ret WorkspaceSourceItemLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkspaceSourceItem) GetLinksOk() (*WorkspaceSourceItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *WorkspaceSourceItem) SetLinks(v WorkspaceSourceItemLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *WorkspaceSourceItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkspaceSourceItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *WorkspaceSourceItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetDeprecated returns the Deprecated field value
func (o *WorkspaceSourceItem) GetDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSourceItem) GetDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deprecated, true
}

// SetDeprecated sets field value
func (o *WorkspaceSourceItem) SetDeprecated(v bool) {
	o.Deprecated = v
}

// GetDeprecationInfo returns the DeprecationInfo field value if set, zero value otherwise.
func (o *WorkspaceSourceItem) GetDeprecationInfo() DeprecationInfo {
	if o == nil || IsNil(o.DeprecationInfo) {
		var ret DeprecationInfo
		return ret
	}
	return *o.DeprecationInfo
}

// GetDeprecationInfoOk returns a tuple with the DeprecationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceSourceItem) GetDeprecationInfoOk() (*DeprecationInfo, bool) {
	if o == nil || IsNil(o.DeprecationInfo) {
		return nil, false
	}
	return o.DeprecationInfo, true
}

// HasDeprecationInfo returns a boolean if a field has been set.
func (o *WorkspaceSourceItem) HasDeprecationInfo() bool {
	if o != nil && !IsNil(o.DeprecationInfo) {
		return true
	}

	return false
}

// SetDeprecationInfo gets a reference to the given DeprecationInfo and assigns it to the DeprecationInfo field.
func (o *WorkspaceSourceItem) SetDeprecationInfo(v DeprecationInfo) {
	o.DeprecationInfo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkspaceSourceItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceSourceItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkspaceSourceItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkspaceSourceItem) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *WorkspaceSourceItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSourceItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkspaceSourceItem) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *WorkspaceSourceItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSourceItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *WorkspaceSourceItem) SetTitle(v string) {
	o.Title = v
}

func (o WorkspaceSourceItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkspaceSourceItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links.Get()
	toSerialize["_metadata"] = o.Metadata.Get()
	toSerialize["deprecated"] = o.Deprecated
	if !IsNil(o.DeprecationInfo) {
		toSerialize["deprecationInfo"] = o.DeprecationInfo
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *WorkspaceSourceItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_links",
		"_metadata",
		"deprecated",
		"name",
		"title",
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

	varWorkspaceSourceItem := _WorkspaceSourceItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkspaceSourceItem)

	if err != nil {
		return err
	}

	*o = WorkspaceSourceItem(varWorkspaceSourceItem)

	return err
}

type NullableWorkspaceSourceItem struct {
	value *WorkspaceSourceItem
	isSet bool
}

func (v NullableWorkspaceSourceItem) Get() *WorkspaceSourceItem {
	return v.value
}

func (v *NullableWorkspaceSourceItem) Set(val *WorkspaceSourceItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceSourceItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceSourceItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceSourceItem(val *WorkspaceSourceItem) *NullableWorkspaceSourceItem {
	return &NullableWorkspaceSourceItem{value: val, isSet: true}
}

func (v NullableWorkspaceSourceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceSourceItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


