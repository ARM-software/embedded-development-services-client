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

// checks if the ArtefactManagerLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtefactManagerLinks{}

// ArtefactManagerLinks links to manage artefacts Note: the links for actions on artefacts (i.e. upload, download, clear) should have in their title the title specified by the user for the artefact when uploading it.
type ArtefactManagerLinks struct {
	Clear *HalLinkData `json:"clear,omitempty"`
	Collection *HalLinkData `json:"collection,omitempty"`
	Download *HalLinkData `json:"download,omitempty"`
	Related HalLinkData `json:"related"`
	Self HalLinkData `json:"self"`
	Upload *HalLinkData `json:"upload,omitempty"`
}

type _ArtefactManagerLinks ArtefactManagerLinks

// NewArtefactManagerLinks instantiates a new ArtefactManagerLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtefactManagerLinks(related HalLinkData, self HalLinkData) *ArtefactManagerLinks {
	this := ArtefactManagerLinks{}
	this.Related = related
	this.Self = self
	return &this
}

// NewArtefactManagerLinksWithDefaults instantiates a new ArtefactManagerLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtefactManagerLinksWithDefaults() *ArtefactManagerLinks {
	this := ArtefactManagerLinks{}
	return &this
}

// GetClear returns the Clear field value if set, zero value otherwise.
func (o *ArtefactManagerLinks) GetClear() HalLinkData {
	if o == nil || IsNil(o.Clear) {
		var ret HalLinkData
		return ret
	}
	return *o.Clear
}

// GetClearOk returns a tuple with the Clear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtefactManagerLinks) GetClearOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Clear) {
		return nil, false
	}
	return o.Clear, true
}

// HasClear returns a boolean if a field has been set.
func (o *ArtefactManagerLinks) HasClear() bool {
	if o != nil && !IsNil(o.Clear) {
		return true
	}

	return false
}

// SetClear gets a reference to the given HalLinkData and assigns it to the Clear field.
func (o *ArtefactManagerLinks) SetClear(v HalLinkData) {
	o.Clear = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *ArtefactManagerLinks) GetCollection() HalLinkData {
	if o == nil || IsNil(o.Collection) {
		var ret HalLinkData
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtefactManagerLinks) GetCollectionOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *ArtefactManagerLinks) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given HalLinkData and assigns it to the Collection field.
func (o *ArtefactManagerLinks) SetCollection(v HalLinkData) {
	o.Collection = &v
}

// GetDownload returns the Download field value if set, zero value otherwise.
func (o *ArtefactManagerLinks) GetDownload() HalLinkData {
	if o == nil || IsNil(o.Download) {
		var ret HalLinkData
		return ret
	}
	return *o.Download
}

// GetDownloadOk returns a tuple with the Download field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtefactManagerLinks) GetDownloadOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Download) {
		return nil, false
	}
	return o.Download, true
}

// HasDownload returns a boolean if a field has been set.
func (o *ArtefactManagerLinks) HasDownload() bool {
	if o != nil && !IsNil(o.Download) {
		return true
	}

	return false
}

// SetDownload gets a reference to the given HalLinkData and assigns it to the Download field.
func (o *ArtefactManagerLinks) SetDownload(v HalLinkData) {
	o.Download = &v
}

// GetRelated returns the Related field value
func (o *ArtefactManagerLinks) GetRelated() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
func (o *ArtefactManagerLinks) GetRelatedOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Related, true
}

// SetRelated sets field value
func (o *ArtefactManagerLinks) SetRelated(v HalLinkData) {
	o.Related = v
}

// GetSelf returns the Self field value
func (o *ArtefactManagerLinks) GetSelf() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ArtefactManagerLinks) GetSelfOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *ArtefactManagerLinks) SetSelf(v HalLinkData) {
	o.Self = v
}

// GetUpload returns the Upload field value if set, zero value otherwise.
func (o *ArtefactManagerLinks) GetUpload() HalLinkData {
	if o == nil || IsNil(o.Upload) {
		var ret HalLinkData
		return ret
	}
	return *o.Upload
}

// GetUploadOk returns a tuple with the Upload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtefactManagerLinks) GetUploadOk() (*HalLinkData, bool) {
	if o == nil || IsNil(o.Upload) {
		return nil, false
	}
	return o.Upload, true
}

// HasUpload returns a boolean if a field has been set.
func (o *ArtefactManagerLinks) HasUpload() bool {
	if o != nil && !IsNil(o.Upload) {
		return true
	}

	return false
}

// SetUpload gets a reference to the given HalLinkData and assigns it to the Upload field.
func (o *ArtefactManagerLinks) SetUpload(v HalLinkData) {
	o.Upload = &v
}

func (o ArtefactManagerLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtefactManagerLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clear) {
		toSerialize["clear"] = o.Clear
	}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !IsNil(o.Download) {
		toSerialize["download"] = o.Download
	}
	toSerialize["related"] = o.Related
	toSerialize["self"] = o.Self
	if !IsNil(o.Upload) {
		toSerialize["upload"] = o.Upload
	}
	return toSerialize, nil
}

func (o *ArtefactManagerLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"related",
		"self",
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

	varArtefactManagerLinks := _ArtefactManagerLinks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArtefactManagerLinks)

	if err != nil {
		return err
	}

	*o = ArtefactManagerLinks(varArtefactManagerLinks)

	return err
}

type NullableArtefactManagerLinks struct {
	value *ArtefactManagerLinks
	isSet bool
}

func (v NullableArtefactManagerLinks) Get() *ArtefactManagerLinks {
	return v.value
}

func (v *NullableArtefactManagerLinks) Set(val *ArtefactManagerLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableArtefactManagerLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableArtefactManagerLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtefactManagerLinks(val *ArtefactManagerLinks) *NullableArtefactManagerLinks {
	return &NullableArtefactManagerLinks{value: val, isSet: true}
}

func (v NullableArtefactManagerLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtefactManagerLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


