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
	"time"
	"bytes"
	"fmt"
)

// checks if the PATItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATItem{}

// PATItem A Personal Access Token Item.
type PATItem struct {
	Links PATItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	// UTC date and time when the token was last used.
	LastUsed time.Time `json:"lastUsed"`
	// Unique ID of the personal access token.
	Name string `json:"name"`
	// Human readable name of the personal access token.
	Title string `json:"title"`
}

type _PATItem PATItem

// NewPATItem instantiates a new PATItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATItem(links PATItemLinks, metadata NullableCommonMetadata, lastUsed time.Time, name string, title string) *PATItem {
	this := PATItem{}
	this.Links = links
	this.Metadata = metadata
	this.LastUsed = lastUsed
	this.Name = name
	this.Title = title
	return &this
}

// NewPATItemWithDefaults instantiates a new PATItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATItemWithDefaults() *PATItem {
	this := PATItem{}
	return &this
}

// GetLinks returns the Links field value
func (o *PATItem) GetLinks() PATItemLinks {
	if o == nil {
		var ret PATItemLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PATItem) GetLinksOk() (*PATItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PATItem) SetLinks(v PATItemLinks) {
	o.Links = v
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *PATItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *PATItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetLastUsed returns the LastUsed field value
func (o *PATItem) GetLastUsed() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUsed
}

// GetLastUsedOk returns a tuple with the LastUsed field value
// and a boolean to check if the value has been set.
func (o *PATItem) GetLastUsedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUsed, true
}

// SetLastUsed sets field value
func (o *PATItem) SetLastUsed(v time.Time) {
	o.LastUsed = v
}

// GetName returns the Name field value
func (o *PATItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PATItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PATItem) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *PATItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PATItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PATItem) SetTitle(v string) {
	o.Title = v
}

func (o PATItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links
	toSerialize["_metadata"] = o.Metadata.Get()
	toSerialize["lastUsed"] = o.LastUsed
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *PATItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_links",
		"_metadata",
		"lastUsed",
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

	varPATItem := _PATItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPATItem)

	if err != nil {
		return err
	}

	*o = PATItem(varPATItem)

	return err
}

type NullablePATItem struct {
	value *PATItem
	isSet bool
}

func (v NullablePATItem) Get() *PATItem {
	return v.value
}

func (v *NullablePATItem) Set(val *PATItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePATItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePATItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATItem(val *PATItem) *NullablePATItem {
	return &NullablePATItem{value: val, isSet: true}
}

func (v NullablePATItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


