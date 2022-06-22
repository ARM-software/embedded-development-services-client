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

// CmsisIntellisenseItemDeprecationInfo Further information on the deprecation status.
type CmsisIntellisenseItemDeprecationInfo struct {
	Comment *string `json:"comment,omitempty"`
	Issued *time.Time `json:"issued,omitempty"`
	Removal *time.Time `json:"removal,omitempty"`
}

// NewCmsisIntellisenseItemDeprecationInfo instantiates a new CmsisIntellisenseItemDeprecationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmsisIntellisenseItemDeprecationInfo() *CmsisIntellisenseItemDeprecationInfo {
	this := CmsisIntellisenseItemDeprecationInfo{}
	return &this
}

// NewCmsisIntellisenseItemDeprecationInfoWithDefaults instantiates a new CmsisIntellisenseItemDeprecationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmsisIntellisenseItemDeprecationInfoWithDefaults() *CmsisIntellisenseItemDeprecationInfo {
	this := CmsisIntellisenseItemDeprecationInfo{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CmsisIntellisenseItemDeprecationInfo) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsisIntellisenseItemDeprecationInfo) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CmsisIntellisenseItemDeprecationInfo) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CmsisIntellisenseItemDeprecationInfo) SetComment(v string) {
	o.Comment = &v
}

// GetIssued returns the Issued field value if set, zero value otherwise.
func (o *CmsisIntellisenseItemDeprecationInfo) GetIssued() time.Time {
	if o == nil || o.Issued == nil {
		var ret time.Time
		return ret
	}
	return *o.Issued
}

// GetIssuedOk returns a tuple with the Issued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsisIntellisenseItemDeprecationInfo) GetIssuedOk() (*time.Time, bool) {
	if o == nil || o.Issued == nil {
		return nil, false
	}
	return o.Issued, true
}

// HasIssued returns a boolean if a field has been set.
func (o *CmsisIntellisenseItemDeprecationInfo) HasIssued() bool {
	if o != nil && o.Issued != nil {
		return true
	}

	return false
}

// SetIssued gets a reference to the given time.Time and assigns it to the Issued field.
func (o *CmsisIntellisenseItemDeprecationInfo) SetIssued(v time.Time) {
	o.Issued = &v
}

// GetRemoval returns the Removal field value if set, zero value otherwise.
func (o *CmsisIntellisenseItemDeprecationInfo) GetRemoval() time.Time {
	if o == nil || o.Removal == nil {
		var ret time.Time
		return ret
	}
	return *o.Removal
}

// GetRemovalOk returns a tuple with the Removal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmsisIntellisenseItemDeprecationInfo) GetRemovalOk() (*time.Time, bool) {
	if o == nil || o.Removal == nil {
		return nil, false
	}
	return o.Removal, true
}

// HasRemoval returns a boolean if a field has been set.
func (o *CmsisIntellisenseItemDeprecationInfo) HasRemoval() bool {
	if o != nil && o.Removal != nil {
		return true
	}

	return false
}

// SetRemoval gets a reference to the given time.Time and assigns it to the Removal field.
func (o *CmsisIntellisenseItemDeprecationInfo) SetRemoval(v time.Time) {
	o.Removal = &v
}

func (o CmsisIntellisenseItemDeprecationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Issued != nil {
		toSerialize["issued"] = o.Issued
	}
	if o.Removal != nil {
		toSerialize["removal"] = o.Removal
	}
	return json.Marshal(toSerialize)
}

type NullableCmsisIntellisenseItemDeprecationInfo struct {
	value *CmsisIntellisenseItemDeprecationInfo
	isSet bool
}

func (v NullableCmsisIntellisenseItemDeprecationInfo) Get() *CmsisIntellisenseItemDeprecationInfo {
	return v.value
}

func (v *NullableCmsisIntellisenseItemDeprecationInfo) Set(val *CmsisIntellisenseItemDeprecationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCmsisIntellisenseItemDeprecationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCmsisIntellisenseItemDeprecationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmsisIntellisenseItemDeprecationInfo(val *CmsisIntellisenseItemDeprecationInfo) *NullableCmsisIntellisenseItemDeprecationInfo {
	return &NullableCmsisIntellisenseItemDeprecationInfo{value: val, isSet: true}
}

func (v NullableCmsisIntellisenseItemDeprecationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmsisIntellisenseItemDeprecationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


