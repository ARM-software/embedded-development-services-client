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

// checks if the VhtCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VhtCollection{}

// VhtCollection This collection resource follows the common pattern of linking to contained resources. Optionally, rather than linking to other resources, it can embed then into the collection to reduce the number of round trips to the server (at the expense of caching). In file system terms, it is similar to a directory but only contains links to (or embeds) a single type of resource.
type VhtCollection struct {
	Embedded *EmbeddedVhtItems `json:"_embedded,omitempty"`
	Links NullableHalCollectionLinks `json:"_links"`
	Metadata NullablePagingMetadata `json:"_metadata"`
	// ID of the Collection.
	Name string `json:"name"`
	// Human readable title of the collection.
	Title string `json:"title"`
}

type _VhtCollection VhtCollection

// NewVhtCollection instantiates a new VhtCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVhtCollection(links NullableHalCollectionLinks, metadata NullablePagingMetadata, name string, title string) *VhtCollection {
	this := VhtCollection{}
	this.Links = links
	this.Metadata = metadata
	this.Name = name
	this.Title = title
	return &this
}

// NewVhtCollectionWithDefaults instantiates a new VhtCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVhtCollectionWithDefaults() *VhtCollection {
	this := VhtCollection{}
	return &this
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *VhtCollection) GetEmbedded() EmbeddedVhtItems {
	if o == nil || IsNil(o.Embedded) {
		var ret EmbeddedVhtItems
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VhtCollection) GetEmbeddedOk() (*EmbeddedVhtItems, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *VhtCollection) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given EmbeddedVhtItems and assigns it to the Embedded field.
func (o *VhtCollection) SetEmbedded(v EmbeddedVhtItems) {
	o.Embedded = &v
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for HalCollectionLinks will be returned
func (o *VhtCollection) GetLinks() HalCollectionLinks {
	if o == nil || o.Links.Get() == nil {
		var ret HalCollectionLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VhtCollection) GetLinksOk() (*HalCollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *VhtCollection) SetLinks(v HalCollectionLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for PagingMetadata will be returned
func (o *VhtCollection) GetMetadata() PagingMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret PagingMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VhtCollection) GetMetadataOk() (*PagingMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *VhtCollection) SetMetadata(v PagingMetadata) {
	o.Metadata.Set(&v)
}

// GetName returns the Name field value
func (o *VhtCollection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VhtCollection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VhtCollection) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *VhtCollection) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *VhtCollection) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *VhtCollection) SetTitle(v string) {
	o.Title = v
}

func (o VhtCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VhtCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	toSerialize["_links"] = o.Links.Get()
	toSerialize["_metadata"] = o.Metadata.Get()
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *VhtCollection) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_links",
		"_metadata",
		"name",
		"title",
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

	varVhtCollection := _VhtCollection{}

	err = json.Unmarshal(bytes, &varVhtCollection)

	if err != nil {
		return err
	}

	*o = VhtCollection(varVhtCollection)

	return err
}

type NullableVhtCollection struct {
	value *VhtCollection
	isSet bool
}

func (v NullableVhtCollection) Get() *VhtCollection {
	return v.value
}

func (v *NullableVhtCollection) Set(val *VhtCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableVhtCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableVhtCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVhtCollection(val *VhtCollection) *NullableVhtCollection {
	return &NullableVhtCollection{value: val, isSet: true}
}

func (v NullableVhtCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVhtCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


