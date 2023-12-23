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

// checks if the VhtItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VhtItem{}

// VhtItem struct for VhtItem
type VhtItem struct {
	Links NullableVhtItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	AdditionalTools []AdditionalTool `json:"additionalTools,omitempty"`
	// The board Target.
	Board string `json:"board"`
	// The revision of the board Target.
	BoardRevision *string `json:"boardRevision,omitempty"`
	// True if this VHT is scheduled to be removed from the service.
	Deprecated bool `json:"deprecated"`
	DeprecationInfo *DeprecationInfo `json:"deprecationInfo,omitempty"`
	// The model used
	Model NullableString `json:"model,omitempty"`
	// Unique ID of the VHT.
	Name string `json:"name"`
	// Human readable name of the VHT.
	Title string `json:"title"`
	// The vendor who supplies the VHT.
	Vendor string `json:"vendor"`
	VirtualInterfaces []VirtualInterface `json:"virtualInterfaces,omitempty"`
}

type _VhtItem VhtItem

// NewVhtItem instantiates a new VhtItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVhtItem(links NullableVhtItemLinks, metadata NullableCommonMetadata, board string, deprecated bool, name string, title string, vendor string) *VhtItem {
	this := VhtItem{}
	this.Links = links
	this.Metadata = metadata
	this.Board = board
	this.Deprecated = deprecated
	this.Name = name
	this.Title = title
	this.Vendor = vendor
	return &this
}

// NewVhtItemWithDefaults instantiates a new VhtItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVhtItemWithDefaults() *VhtItem {
	this := VhtItem{}
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for VhtItemLinks will be returned
func (o *VhtItem) GetLinks() VhtItemLinks {
	if o == nil || o.Links.Get() == nil {
		var ret VhtItemLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VhtItem) GetLinksOk() (*VhtItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *VhtItem) SetLinks(v VhtItemLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *VhtItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VhtItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *VhtItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetAdditionalTools returns the AdditionalTools field value if set, zero value otherwise.
func (o *VhtItem) GetAdditionalTools() []AdditionalTool {
	if o == nil || IsNil(o.AdditionalTools) {
		var ret []AdditionalTool
		return ret
	}
	return o.AdditionalTools
}

// GetAdditionalToolsOk returns a tuple with the AdditionalTools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtItem) GetAdditionalToolsOk() ([]AdditionalTool, bool) {
	if o == nil || IsNil(o.AdditionalTools) {
		return nil, false
	}
	return o.AdditionalTools, true
}

// HasAdditionalTools returns a boolean if a field has been set.
func (o *VhtItem) HasAdditionalTools() bool {
	if o != nil && !IsNil(o.AdditionalTools) {
		return true
	}

	return false
}

// SetAdditionalTools gets a reference to the given []AdditionalTool and assigns it to the AdditionalTools field.
func (o *VhtItem) SetAdditionalTools(v []AdditionalTool) {
	o.AdditionalTools = v
}

// GetBoard returns the Board field value
func (o *VhtItem) GetBoard() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Board
}

// GetBoardOk returns a tuple with the Board field value
// and a boolean to check if the value has been set.
func (o *VhtItem) GetBoardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Board, true
}

// SetBoard sets field value
func (o *VhtItem) SetBoard(v string) {
	o.Board = v
}

// GetBoardRevision returns the BoardRevision field value if set, zero value otherwise.
func (o *VhtItem) GetBoardRevision() string {
	if o == nil || IsNil(o.BoardRevision) {
		var ret string
		return ret
	}
	return *o.BoardRevision
}

// GetBoardRevisionOk returns a tuple with the BoardRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtItem) GetBoardRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.BoardRevision) {
		return nil, false
	}
	return o.BoardRevision, true
}

// HasBoardRevision returns a boolean if a field has been set.
func (o *VhtItem) HasBoardRevision() bool {
	if o != nil && !IsNil(o.BoardRevision) {
		return true
	}

	return false
}

// SetBoardRevision gets a reference to the given string and assigns it to the BoardRevision field.
func (o *VhtItem) SetBoardRevision(v string) {
	o.BoardRevision = &v
}

// GetDeprecated returns the Deprecated field value
func (o *VhtItem) GetDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value
// and a boolean to check if the value has been set.
func (o *VhtItem) GetDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deprecated, true
}

// SetDeprecated sets field value
func (o *VhtItem) SetDeprecated(v bool) {
	o.Deprecated = v
}

// GetDeprecationInfo returns the DeprecationInfo field value if set, zero value otherwise.
func (o *VhtItem) GetDeprecationInfo() DeprecationInfo {
	if o == nil || IsNil(o.DeprecationInfo) {
		var ret DeprecationInfo
		return ret
	}
	return *o.DeprecationInfo
}

// GetDeprecationInfoOk returns a tuple with the DeprecationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtItem) GetDeprecationInfoOk() (*DeprecationInfo, bool) {
	if o == nil || IsNil(o.DeprecationInfo) {
		return nil, false
	}
	return o.DeprecationInfo, true
}

// HasDeprecationInfo returns a boolean if a field has been set.
func (o *VhtItem) HasDeprecationInfo() bool {
	if o != nil && !IsNil(o.DeprecationInfo) {
		return true
	}

	return false
}

// SetDeprecationInfo gets a reference to the given DeprecationInfo and assigns it to the DeprecationInfo field.
func (o *VhtItem) SetDeprecationInfo(v DeprecationInfo) {
	o.DeprecationInfo = &v
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VhtItem) GetModel() string {
	if o == nil || IsNil(o.Model.Get()) {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VhtItem) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *VhtItem) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *VhtItem) SetModel(v string) {
	o.Model.Set(&v)
}
// SetModelNil sets the value for Model to be an explicit nil
func (o *VhtItem) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *VhtItem) UnsetModel() {
	o.Model.Unset()
}

// GetName returns the Name field value
func (o *VhtItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VhtItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VhtItem) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *VhtItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *VhtItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *VhtItem) SetTitle(v string) {
	o.Title = v
}

// GetVendor returns the Vendor field value
func (o *VhtItem) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *VhtItem) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *VhtItem) SetVendor(v string) {
	o.Vendor = v
}

// GetVirtualInterfaces returns the VirtualInterfaces field value if set, zero value otherwise.
func (o *VhtItem) GetVirtualInterfaces() []VirtualInterface {
	if o == nil || IsNil(o.VirtualInterfaces) {
		var ret []VirtualInterface
		return ret
	}
	return o.VirtualInterfaces
}

// GetVirtualInterfacesOk returns a tuple with the VirtualInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtItem) GetVirtualInterfacesOk() ([]VirtualInterface, bool) {
	if o == nil || IsNil(o.VirtualInterfaces) {
		return nil, false
	}
	return o.VirtualInterfaces, true
}

// HasVirtualInterfaces returns a boolean if a field has been set.
func (o *VhtItem) HasVirtualInterfaces() bool {
	if o != nil && !IsNil(o.VirtualInterfaces) {
		return true
	}

	return false
}

// SetVirtualInterfaces gets a reference to the given []VirtualInterface and assigns it to the VirtualInterfaces field.
func (o *VhtItem) SetVirtualInterfaces(v []VirtualInterface) {
	o.VirtualInterfaces = v
}

func (o VhtItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VhtItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links.Get()
	toSerialize["_metadata"] = o.Metadata.Get()
	if !IsNil(o.AdditionalTools) {
		toSerialize["additionalTools"] = o.AdditionalTools
	}
	toSerialize["board"] = o.Board
	if !IsNil(o.BoardRevision) {
		toSerialize["boardRevision"] = o.BoardRevision
	}
	toSerialize["deprecated"] = o.Deprecated
	if !IsNil(o.DeprecationInfo) {
		toSerialize["deprecationInfo"] = o.DeprecationInfo
	}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	toSerialize["vendor"] = o.Vendor
	if !IsNil(o.VirtualInterfaces) {
		toSerialize["virtualInterfaces"] = o.VirtualInterfaces
	}
	return toSerialize, nil
}

func (o *VhtItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_links",
		"_metadata",
		"board",
		"deprecated",
		"name",
		"title",
		"vendor",
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

	varVhtItem := _VhtItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVhtItem)

	if err != nil {
		return err
	}

	*o = VhtItem(varVhtItem)

	return err
}

type NullableVhtItem struct {
	value *VhtItem
	isSet bool
}

func (v NullableVhtItem) Get() *VhtItem {
	return v.value
}

func (v *NullableVhtItem) Set(val *VhtItem) {
	v.value = val
	v.isSet = true
}

func (v NullableVhtItem) IsSet() bool {
	return v.isSet
}

func (v *NullableVhtItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVhtItem(val *VhtItem) *NullableVhtItem {
	return &NullableVhtItem{value: val, isSet: true}
}

func (v NullableVhtItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVhtItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


