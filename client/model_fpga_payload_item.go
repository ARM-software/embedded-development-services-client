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

// checks if the FPGAPayloadItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FPGAPayloadItem{}

// FPGAPayloadItem struct for FPGAPayloadItem
type FPGAPayloadItem struct {
	Links NullableFPGAPayloadItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	// Unique ID of this FPGA Payload.
	Name string `json:"name"`
	// Status of the payload. A payload is only ready to be used by an FPGA once it has been processed.
	Status string `json:"status"`
	// True when the payload supports direct connection.
	SupportConnection bool `json:"supportConnection"`
	// Human readable name of the FPGA Payload.
	Title string `json:"title"`
	// The upload location to upload the payload files from. This value will be returned from the upload session creation.
	UploadLocation string `json:"uploadLocation"`
}

type _FPGAPayloadItem FPGAPayloadItem

// NewFPGAPayloadItem instantiates a new FPGAPayloadItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFPGAPayloadItem(links NullableFPGAPayloadItemLinks, metadata NullableCommonMetadata, name string, status string, supportConnection bool, title string, uploadLocation string) *FPGAPayloadItem {
	this := FPGAPayloadItem{}
	this.Links = links
	this.Metadata = metadata
	this.Name = name
	this.Status = status
	this.SupportConnection = supportConnection
	this.Title = title
	this.UploadLocation = uploadLocation
	return &this
}

// NewFPGAPayloadItemWithDefaults instantiates a new FPGAPayloadItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFPGAPayloadItemWithDefaults() *FPGAPayloadItem {
	this := FPGAPayloadItem{}
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for FPGAPayloadItemLinks will be returned
func (o *FPGAPayloadItem) GetLinks() FPGAPayloadItemLinks {
	if o == nil || o.Links.Get() == nil {
		var ret FPGAPayloadItemLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPGAPayloadItem) GetLinksOk() (*FPGAPayloadItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *FPGAPayloadItem) SetLinks(v FPGAPayloadItemLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *FPGAPayloadItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FPGAPayloadItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *FPGAPayloadItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetName returns the Name field value
func (o *FPGAPayloadItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FPGAPayloadItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FPGAPayloadItem) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *FPGAPayloadItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FPGAPayloadItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FPGAPayloadItem) SetStatus(v string) {
	o.Status = v
}

// GetSupportConnection returns the SupportConnection field value
func (o *FPGAPayloadItem) GetSupportConnection() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SupportConnection
}

// GetSupportConnectionOk returns a tuple with the SupportConnection field value
// and a boolean to check if the value has been set.
func (o *FPGAPayloadItem) GetSupportConnectionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportConnection, true
}

// SetSupportConnection sets field value
func (o *FPGAPayloadItem) SetSupportConnection(v bool) {
	o.SupportConnection = v
}

// GetTitle returns the Title field value
func (o *FPGAPayloadItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *FPGAPayloadItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *FPGAPayloadItem) SetTitle(v string) {
	o.Title = v
}

// GetUploadLocation returns the UploadLocation field value
func (o *FPGAPayloadItem) GetUploadLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadLocation
}

// GetUploadLocationOk returns a tuple with the UploadLocation field value
// and a boolean to check if the value has been set.
func (o *FPGAPayloadItem) GetUploadLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadLocation, true
}

// SetUploadLocation sets field value
func (o *FPGAPayloadItem) SetUploadLocation(v string) {
	o.UploadLocation = v
}

func (o FPGAPayloadItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FPGAPayloadItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links.Get()
	toSerialize["_metadata"] = o.Metadata.Get()
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["supportConnection"] = o.SupportConnection
	toSerialize["title"] = o.Title
	toSerialize["uploadLocation"] = o.UploadLocation
	return toSerialize, nil
}

func (o *FPGAPayloadItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_links",
		"_metadata",
		"name",
		"status",
		"supportConnection",
		"title",
		"uploadLocation",
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

	varFPGAPayloadItem := _FPGAPayloadItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFPGAPayloadItem)

	if err != nil {
		return err
	}

	*o = FPGAPayloadItem(varFPGAPayloadItem)

	return err
}

type NullableFPGAPayloadItem struct {
	value *FPGAPayloadItem
	isSet bool
}

func (v NullableFPGAPayloadItem) Get() *FPGAPayloadItem {
	return v.value
}

func (v *NullableFPGAPayloadItem) Set(val *FPGAPayloadItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFPGAPayloadItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFPGAPayloadItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFPGAPayloadItem(val *FPGAPayloadItem) *NullableFPGAPayloadItem {
	return &NullableFPGAPayloadItem{value: val, isSet: true}
}

func (v NullableFPGAPayloadItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFPGAPayloadItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


