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

// checks if the PATCreationItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCreationItem{}

// PATCreationItem A Personal Access Token Item.
type PATCreationItem struct {
	// The TTL (time to live in seconds) describing how long the personal access token will be alive for.
	TTL int64 `json:"TTL"`
	Links PATItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	// Unique ID of the personal access token.
	Name string `json:"name"`
	// The personal access token.
	Secret string `json:"secret"`
	// Human readable name of the personal access token.
	Title string `json:"title"`
}

type _PATCreationItem PATCreationItem

// NewPATCreationItem instantiates a new PATCreationItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCreationItem(tTL int64, links PATItemLinks, metadata NullableCommonMetadata, name string, secret string, title string) *PATCreationItem {
	this := PATCreationItem{}
	this.TTL = tTL
	this.Links = links
	this.Metadata = metadata
	this.Name = name
	this.Secret = secret
	this.Title = title
	return &this
}

// NewPATCreationItemWithDefaults instantiates a new PATCreationItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCreationItemWithDefaults() *PATCreationItem {
	this := PATCreationItem{}
	return &this
}

// GetTTL returns the TTL field value
func (o *PATCreationItem) GetTTL() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TTL
}

// GetTTLOk returns a tuple with the TTL field value
// and a boolean to check if the value has been set.
func (o *PATCreationItem) GetTTLOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TTL, true
}

// SetTTL sets field value
func (o *PATCreationItem) SetTTL(v int64) {
	o.TTL = v
}

// GetLinks returns the Links field value
func (o *PATCreationItem) GetLinks() PATItemLinks {
	if o == nil {
		var ret PATItemLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PATCreationItem) GetLinksOk() (*PATItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PATCreationItem) SetLinks(v PATItemLinks) {
	o.Links = v
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *PATCreationItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCreationItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *PATCreationItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetName returns the Name field value
func (o *PATCreationItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PATCreationItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PATCreationItem) SetName(v string) {
	o.Name = v
}

// GetSecret returns the Secret field value
func (o *PATCreationItem) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *PATCreationItem) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *PATCreationItem) SetSecret(v string) {
	o.Secret = v
}

// GetTitle returns the Title field value
func (o *PATCreationItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PATCreationItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PATCreationItem) SetTitle(v string) {
	o.Title = v
}

func (o PATCreationItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCreationItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["TTL"] = o.TTL
	toSerialize["_links"] = o.Links
	toSerialize["_metadata"] = o.Metadata.Get()
	toSerialize["name"] = o.Name
	toSerialize["secret"] = o.Secret
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *PATCreationItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"TTL",
		"_links",
		"_metadata",
		"name",
		"secret",
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

	varPATCreationItem := _PATCreationItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPATCreationItem)

	if err != nil {
		return err
	}

	*o = PATCreationItem(varPATCreationItem)

	return err
}

type NullablePATCreationItem struct {
	value *PATCreationItem
	isSet bool
}

func (v NullablePATCreationItem) Get() *PATCreationItem {
	return v.value
}

func (v *NullablePATCreationItem) Set(val *PATCreationItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCreationItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCreationItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCreationItem(val *PATCreationItem) *NullablePATCreationItem {
	return &NullablePATCreationItem{value: val, isSet: true}
}

func (v NullablePATCreationItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCreationItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


