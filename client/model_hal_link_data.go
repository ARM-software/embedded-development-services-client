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

// checks if the HalLinkData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HalLinkData{}

// HalLinkData The base HAL hyperlink object.
type HalLinkData struct {
	// Indicates that the link is to be deprecated (i.e. removed) at a future date. Its value is a URL
	Deprecation *string `json:"deprecation,omitempty"`
	// link URI or URI Template
	Href string `json:"href"`
	// Indicates the language of the target resource (as defined by [RFC5988]).
	Hreflang *string `json:"hreflang,omitempty"`
	// key for selecting Link Objects which share the same relation type.
	Name *string `json:"name,omitempty"`
	// URI that hints about the profile (as defined by [I-D.wilde-profile-link]) of the target resource.
	Profile *string `json:"profile,omitempty"`
	// States whether Link Object's \"href\" is a URI Template
	Templated *bool `json:"templated,omitempty"`
	// Human-readable identifier labelling the link.
	Title *string `json:"title,omitempty"`
	// hint to indicate the media type
	Type *string `json:"type,omitempty"`
}

type _HalLinkData HalLinkData

// NewHalLinkData instantiates a new HalLinkData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHalLinkData(href string) *HalLinkData {
	this := HalLinkData{}
	this.Href = href
	var templated bool = false
	this.Templated = &templated
	return &this
}

// NewHalLinkDataWithDefaults instantiates a new HalLinkData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHalLinkDataWithDefaults() *HalLinkData {
	this := HalLinkData{}
	var templated bool = false
	this.Templated = &templated
	return &this
}

// GetDeprecation returns the Deprecation field value if set, zero value otherwise.
func (o *HalLinkData) GetDeprecation() string {
	if o == nil || IsNil(o.Deprecation) {
		var ret string
		return ret
	}
	return *o.Deprecation
}

// GetDeprecationOk returns a tuple with the Deprecation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalLinkData) GetDeprecationOk() (*string, bool) {
	if o == nil || IsNil(o.Deprecation) {
		return nil, false
	}
	return o.Deprecation, true
}

// HasDeprecation returns a boolean if a field has been set.
func (o *HalLinkData) HasDeprecation() bool {
	if o != nil && !IsNil(o.Deprecation) {
		return true
	}

	return false
}

// SetDeprecation gets a reference to the given string and assigns it to the Deprecation field.
func (o *HalLinkData) SetDeprecation(v string) {
	o.Deprecation = &v
}

// GetHref returns the Href field value
func (o *HalLinkData) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *HalLinkData) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *HalLinkData) SetHref(v string) {
	o.Href = v
}

// GetHreflang returns the Hreflang field value if set, zero value otherwise.
func (o *HalLinkData) GetHreflang() string {
	if o == nil || IsNil(o.Hreflang) {
		var ret string
		return ret
	}
	return *o.Hreflang
}

// GetHreflangOk returns a tuple with the Hreflang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalLinkData) GetHreflangOk() (*string, bool) {
	if o == nil || IsNil(o.Hreflang) {
		return nil, false
	}
	return o.Hreflang, true
}

// HasHreflang returns a boolean if a field has been set.
func (o *HalLinkData) HasHreflang() bool {
	if o != nil && !IsNil(o.Hreflang) {
		return true
	}

	return false
}

// SetHreflang gets a reference to the given string and assigns it to the Hreflang field.
func (o *HalLinkData) SetHreflang(v string) {
	o.Hreflang = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HalLinkData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalLinkData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HalLinkData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HalLinkData) SetName(v string) {
	o.Name = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *HalLinkData) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalLinkData) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *HalLinkData) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *HalLinkData) SetProfile(v string) {
	o.Profile = &v
}

// GetTemplated returns the Templated field value if set, zero value otherwise.
func (o *HalLinkData) GetTemplated() bool {
	if o == nil || IsNil(o.Templated) {
		var ret bool
		return ret
	}
	return *o.Templated
}

// GetTemplatedOk returns a tuple with the Templated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalLinkData) GetTemplatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Templated) {
		return nil, false
	}
	return o.Templated, true
}

// HasTemplated returns a boolean if a field has been set.
func (o *HalLinkData) HasTemplated() bool {
	if o != nil && !IsNil(o.Templated) {
		return true
	}

	return false
}

// SetTemplated gets a reference to the given bool and assigns it to the Templated field.
func (o *HalLinkData) SetTemplated(v bool) {
	o.Templated = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *HalLinkData) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalLinkData) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *HalLinkData) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *HalLinkData) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HalLinkData) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HalLinkData) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HalLinkData) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HalLinkData) SetType(v string) {
	o.Type = &v
}

func (o HalLinkData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HalLinkData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deprecation) {
		toSerialize["deprecation"] = o.Deprecation
	}
	toSerialize["href"] = o.Href
	if !IsNil(o.Hreflang) {
		toSerialize["hreflang"] = o.Hreflang
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Templated) {
		toSerialize["templated"] = o.Templated
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *HalLinkData) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
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

	varHalLinkData := _HalLinkData{}

	err = json.Unmarshal(bytes, &varHalLinkData)

	if err != nil {
		return err
	}

	*o = HalLinkData(varHalLinkData)

	return err
}

type NullableHalLinkData struct {
	value *HalLinkData
	isSet bool
}

func (v NullableHalLinkData) Get() *HalLinkData {
	return v.value
}

func (v *NullableHalLinkData) Set(val *HalLinkData) {
	v.value = val
	v.isSet = true
}

func (v NullableHalLinkData) IsSet() bool {
	return v.isSet
}

func (v *NullableHalLinkData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHalLinkData(val *HalLinkData) *NullableHalLinkData {
	return &NullableHalLinkData{value: val, isSet: true}
}

func (v NullableHalLinkData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHalLinkData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


