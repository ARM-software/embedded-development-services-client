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
	"fmt"
)

// checks if the CmsisBuilderItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CmsisBuilderItem{}

// CmsisBuilderItem struct for CmsisBuilderItem
type CmsisBuilderItem struct {
	Links NullableCmsisBuilderItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	BuildToolsType BuildToolTypes `json:"buildToolsType"`
	// Version of the build tools in use, as specified by the tools creators.
	BuildToolsVersion string `json:"buildToolsVersion"`
	// True if this CMSIS Builder is scheduled to be removed from the service.
	Deprecated bool `json:"deprecated"`
	DeprecationInfo *DeprecationInfo `json:"deprecationInfo,omitempty"`
	// Unique ID of the CMSIS builder.
	Name string `json:"name"`
	// Human readable name of the CMSIS builder.
	Title string `json:"title"`
	ToolchainType ToolchainTypes `json:"toolchainType"`
	// The version of the toolchain in use as specified by the toolchain supplier.
	ToolchainVersion string `json:"toolchainVersion"`
}

type _CmsisBuilderItem CmsisBuilderItem

// NewCmsisBuilderItem instantiates a new CmsisBuilderItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmsisBuilderItem(links NullableCmsisBuilderItemLinks, metadata NullableCommonMetadata, buildToolsType BuildToolTypes, buildToolsVersion string, deprecated bool, name string, title string, toolchainType ToolchainTypes, toolchainVersion string) *CmsisBuilderItem {
	this := CmsisBuilderItem{}
	this.Links = links
	this.Metadata = metadata
	this.BuildToolsType = buildToolsType
	this.BuildToolsVersion = buildToolsVersion
	this.Deprecated = deprecated
	this.Name = name
	this.Title = title
	this.ToolchainType = toolchainType
	this.ToolchainVersion = toolchainVersion
	return &this
}

// NewCmsisBuilderItemWithDefaults instantiates a new CmsisBuilderItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmsisBuilderItemWithDefaults() *CmsisBuilderItem {
	this := CmsisBuilderItem{}
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for CmsisBuilderItemLinks will be returned
func (o *CmsisBuilderItem) GetLinks() CmsisBuilderItemLinks {
	if o == nil || o.Links.Get() == nil {
		var ret CmsisBuilderItemLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CmsisBuilderItem) GetLinksOk() (*CmsisBuilderItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *CmsisBuilderItem) SetLinks(v CmsisBuilderItemLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *CmsisBuilderItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CmsisBuilderItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *CmsisBuilderItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetBuildToolsType returns the BuildToolsType field value
func (o *CmsisBuilderItem) GetBuildToolsType() BuildToolTypes {
	if o == nil {
		var ret BuildToolTypes
		return ret
	}

	return o.BuildToolsType
}

// GetBuildToolsTypeOk returns a tuple with the BuildToolsType field value
// and a boolean to check if the value has been set.
func (o *CmsisBuilderItem) GetBuildToolsTypeOk() (*BuildToolTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildToolsType, true
}

// SetBuildToolsType sets field value
func (o *CmsisBuilderItem) SetBuildToolsType(v BuildToolTypes) {
	o.BuildToolsType = v
}

// GetBuildToolsVersion returns the BuildToolsVersion field value
func (o *CmsisBuilderItem) GetBuildToolsVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildToolsVersion
}

// GetBuildToolsVersionOk returns a tuple with the BuildToolsVersion field value
// and a boolean to check if the value has been set.
func (o *CmsisBuilderItem) GetBuildToolsVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildToolsVersion, true
}

// SetBuildToolsVersion sets field value
func (o *CmsisBuilderItem) SetBuildToolsVersion(v string) {
	o.BuildToolsVersion = v
}

// GetDeprecated returns the Deprecated field value
func (o *CmsisBuilderItem) GetDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value
// and a boolean to check if the value has been set.
func (o *CmsisBuilderItem) GetDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deprecated, true
}

// SetDeprecated sets field value
func (o *CmsisBuilderItem) SetDeprecated(v bool) {
	o.Deprecated = v
}

// GetDeprecationInfo returns the DeprecationInfo field value if set, zero value otherwise.
func (o *CmsisBuilderItem) GetDeprecationInfo() DeprecationInfo {
	if o == nil || IsNil(o.DeprecationInfo) {
		var ret DeprecationInfo
		return ret
	}
	return *o.DeprecationInfo
}

// GetDeprecationInfoOk returns a tuple with the DeprecationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsisBuilderItem) GetDeprecationInfoOk() (*DeprecationInfo, bool) {
	if o == nil || IsNil(o.DeprecationInfo) {
		return nil, false
	}
	return o.DeprecationInfo, true
}

// HasDeprecationInfo returns a boolean if a field has been set.
func (o *CmsisBuilderItem) HasDeprecationInfo() bool {
	if o != nil && !IsNil(o.DeprecationInfo) {
		return true
	}

	return false
}

// SetDeprecationInfo gets a reference to the given DeprecationInfo and assigns it to the DeprecationInfo field.
func (o *CmsisBuilderItem) SetDeprecationInfo(v DeprecationInfo) {
	o.DeprecationInfo = &v
}

// GetName returns the Name field value
func (o *CmsisBuilderItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CmsisBuilderItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CmsisBuilderItem) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *CmsisBuilderItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CmsisBuilderItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CmsisBuilderItem) SetTitle(v string) {
	o.Title = v
}

// GetToolchainType returns the ToolchainType field value
func (o *CmsisBuilderItem) GetToolchainType() ToolchainTypes {
	if o == nil {
		var ret ToolchainTypes
		return ret
	}

	return o.ToolchainType
}

// GetToolchainTypeOk returns a tuple with the ToolchainType field value
// and a boolean to check if the value has been set.
func (o *CmsisBuilderItem) GetToolchainTypeOk() (*ToolchainTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolchainType, true
}

// SetToolchainType sets field value
func (o *CmsisBuilderItem) SetToolchainType(v ToolchainTypes) {
	o.ToolchainType = v
}

// GetToolchainVersion returns the ToolchainVersion field value
func (o *CmsisBuilderItem) GetToolchainVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToolchainVersion
}

// GetToolchainVersionOk returns a tuple with the ToolchainVersion field value
// and a boolean to check if the value has been set.
func (o *CmsisBuilderItem) GetToolchainVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolchainVersion, true
}

// SetToolchainVersion sets field value
func (o *CmsisBuilderItem) SetToolchainVersion(v string) {
	o.ToolchainVersion = v
}

func (o CmsisBuilderItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmsisBuilderItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links.Get()
	toSerialize["_metadata"] = o.Metadata.Get()
	toSerialize["buildToolsType"] = o.BuildToolsType
	toSerialize["buildToolsVersion"] = o.BuildToolsVersion
	toSerialize["deprecated"] = o.Deprecated
	if !IsNil(o.DeprecationInfo) {
		toSerialize["deprecationInfo"] = o.DeprecationInfo
	}
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	toSerialize["toolchainType"] = o.ToolchainType
	toSerialize["toolchainVersion"] = o.ToolchainVersion
	return toSerialize, nil
}

func (o *CmsisBuilderItem) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_links",
		"_metadata",
		"buildToolsType",
		"buildToolsVersion",
		"deprecated",
		"name",
		"title",
		"toolchainType",
		"toolchainVersion",
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

	varCmsisBuilderItem := _CmsisBuilderItem{}

	err = json.Unmarshal(bytes, &varCmsisBuilderItem)

	if err != nil {
		return err
	}

	*o = CmsisBuilderItem(varCmsisBuilderItem)

	return err
}

type NullableCmsisBuilderItem struct {
	value *CmsisBuilderItem
	isSet bool
}

func (v NullableCmsisBuilderItem) Get() *CmsisBuilderItem {
	return v.value
}

func (v *NullableCmsisBuilderItem) Set(val *CmsisBuilderItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCmsisBuilderItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCmsisBuilderItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmsisBuilderItem(val *CmsisBuilderItem) *NullableCmsisBuilderItem {
	return &NullableCmsisBuilderItem{value: val, isSet: true}
}

func (v NullableCmsisBuilderItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmsisBuilderItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


