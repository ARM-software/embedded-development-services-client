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
	"time"
)

// DeprecationInfo Additional information about the deprecation status.
type DeprecationInfo struct {
	// further explanation of the rationale behind the deprecation
	Comment string `json:"comment"`
	// time when the deprecation notice was first issued
	Issued time.Time `json:"issued"`
	// time when the removal will be effective
	Removal time.Time `json:"removal"`
}

// NewDeprecationInfo instantiates a new DeprecationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecationInfo(comment string, issued time.Time, removal time.Time) *DeprecationInfo {
	this := DeprecationInfo{}
	this.Comment = comment
	this.Issued = issued
	this.Removal = removal
	return &this
}

// NewDeprecationInfoWithDefaults instantiates a new DeprecationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecationInfoWithDefaults() *DeprecationInfo {
	this := DeprecationInfo{}
	return &this
}

// GetComment returns the Comment field value
func (o *DeprecationInfo) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *DeprecationInfo) GetCommentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *DeprecationInfo) SetComment(v string) {
	o.Comment = v
}

// GetIssued returns the Issued field value
func (o *DeprecationInfo) GetIssued() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Issued
}

// GetIssuedOk returns a tuple with the Issued field value
// and a boolean to check if the value has been set.
func (o *DeprecationInfo) GetIssuedOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Issued, true
}

// SetIssued sets field value
func (o *DeprecationInfo) SetIssued(v time.Time) {
	o.Issued = v
}

// GetRemoval returns the Removal field value
func (o *DeprecationInfo) GetRemoval() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Removal
}

// GetRemovalOk returns a tuple with the Removal field value
// and a boolean to check if the value has been set.
func (o *DeprecationInfo) GetRemovalOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Removal, true
}

// SetRemoval sets field value
func (o *DeprecationInfo) SetRemoval(v time.Time) {
	o.Removal = v
}

func (o DeprecationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["issued"] = o.Issued
	}
	if true {
		toSerialize["removal"] = o.Removal
	}
	return json.Marshal(toSerialize)
}

type NullableDeprecationInfo struct {
	value *DeprecationInfo
	isSet bool
}

func (v NullableDeprecationInfo) Get() *DeprecationInfo {
	return v.value
}

func (v *NullableDeprecationInfo) Set(val *DeprecationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecationInfo(val *DeprecationInfo) *NullableDeprecationInfo {
	return &NullableDeprecationInfo{value: val, isSet: true}
}

func (v NullableDeprecationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


