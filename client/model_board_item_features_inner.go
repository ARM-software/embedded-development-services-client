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
)

// BoardItemFeaturesInner A board feature
type BoardItemFeaturesInner struct {
	// Category of the feature.
	Category *string `json:"category,omitempty"`
	// Default name of the feature, used if no name is defined.
	DefaultName *string `json:"default_name,omitempty"`
	// Any further details about the feature.
	Detail *string `json:"detail,omitempty"`
	// Name of the feature.
	Name *string `json:"name,omitempty"`
	// Type of feature.
	Type *string `json:"type,omitempty"`
}

// NewBoardItemFeaturesInner instantiates a new BoardItemFeaturesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoardItemFeaturesInner() *BoardItemFeaturesInner {
	this := BoardItemFeaturesInner{}
	return &this
}

// NewBoardItemFeaturesInnerWithDefaults instantiates a new BoardItemFeaturesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoardItemFeaturesInnerWithDefaults() *BoardItemFeaturesInner {
	this := BoardItemFeaturesInner{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *BoardItemFeaturesInner) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardItemFeaturesInner) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *BoardItemFeaturesInner) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *BoardItemFeaturesInner) SetCategory(v string) {
	o.Category = &v
}

// GetDefaultName returns the DefaultName field value if set, zero value otherwise.
func (o *BoardItemFeaturesInner) GetDefaultName() string {
	if o == nil || o.DefaultName == nil {
		var ret string
		return ret
	}
	return *o.DefaultName
}

// GetDefaultNameOk returns a tuple with the DefaultName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardItemFeaturesInner) GetDefaultNameOk() (*string, bool) {
	if o == nil || o.DefaultName == nil {
		return nil, false
	}
	return o.DefaultName, true
}

// HasDefaultName returns a boolean if a field has been set.
func (o *BoardItemFeaturesInner) HasDefaultName() bool {
	if o != nil && o.DefaultName != nil {
		return true
	}

	return false
}

// SetDefaultName gets a reference to the given string and assigns it to the DefaultName field.
func (o *BoardItemFeaturesInner) SetDefaultName(v string) {
	o.DefaultName = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *BoardItemFeaturesInner) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardItemFeaturesInner) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *BoardItemFeaturesInner) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *BoardItemFeaturesInner) SetDetail(v string) {
	o.Detail = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BoardItemFeaturesInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardItemFeaturesInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BoardItemFeaturesInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BoardItemFeaturesInner) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BoardItemFeaturesInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardItemFeaturesInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BoardItemFeaturesInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BoardItemFeaturesInner) SetType(v string) {
	o.Type = &v
}

func (o BoardItemFeaturesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.DefaultName != nil {
		toSerialize["default_name"] = o.DefaultName
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBoardItemFeaturesInner struct {
	value *BoardItemFeaturesInner
	isSet bool
}

func (v NullableBoardItemFeaturesInner) Get() *BoardItemFeaturesInner {
	return v.value
}

func (v *NullableBoardItemFeaturesInner) Set(val *BoardItemFeaturesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBoardItemFeaturesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBoardItemFeaturesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoardItemFeaturesInner(val *BoardItemFeaturesInner) *NullableBoardItemFeaturesInner {
	return &NullableBoardItemFeaturesInner{value: val, isSet: true}
}

func (v NullableBoardItemFeaturesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoardItemFeaturesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

