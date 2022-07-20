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

// BoardItemLinks The links for a BoardItem.
type BoardItemLinks struct {
	Collection HalLinkData `json:"collection"`
	// Links to external services.
	Curies []HalLinkData `json:"curies,omitempty"`
	Self HalLinkData `json:"self"`
}

// NewBoardItemLinks instantiates a new BoardItemLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoardItemLinks(collection HalLinkData, self HalLinkData) *BoardItemLinks {
	this := BoardItemLinks{}
	this.Collection = collection
	this.Self = self
	return &this
}

// NewBoardItemLinksWithDefaults instantiates a new BoardItemLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoardItemLinksWithDefaults() *BoardItemLinks {
	this := BoardItemLinks{}
	return &this
}

// GetCollection returns the Collection field value
func (o *BoardItemLinks) GetCollection() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *BoardItemLinks) GetCollectionOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *BoardItemLinks) SetCollection(v HalLinkData) {
	o.Collection = v
}

// GetCuries returns the Curies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoardItemLinks) GetCuries() []HalLinkData {
	if o == nil {
		var ret []HalLinkData
		return ret
	}
	return o.Curies
}

// GetCuriesOk returns a tuple with the Curies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoardItemLinks) GetCuriesOk() ([]HalLinkData, bool) {
	if o == nil || o.Curies == nil {
		return nil, false
	}
	return o.Curies, true
}

// HasCuries returns a boolean if a field has been set.
func (o *BoardItemLinks) HasCuries() bool {
	if o != nil && o.Curies != nil {
		return true
	}

	return false
}

// SetCuries gets a reference to the given []HalLinkData and assigns it to the Curies field.
func (o *BoardItemLinks) SetCuries(v []HalLinkData) {
	o.Curies = v
}

// GetSelf returns the Self field value
func (o *BoardItemLinks) GetSelf() HalLinkData {
	if o == nil {
		var ret HalLinkData
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *BoardItemLinks) GetSelfOk() (*HalLinkData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *BoardItemLinks) SetSelf(v HalLinkData) {
	o.Self = v
}

func (o BoardItemLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["collection"] = o.Collection
	}
	if o.Curies != nil {
		toSerialize["curies"] = o.Curies
	}
	if true {
		toSerialize["self"] = o.Self
	}
	return json.Marshal(toSerialize)
}

type NullableBoardItemLinks struct {
	value *BoardItemLinks
	isSet bool
}

func (v NullableBoardItemLinks) Get() *BoardItemLinks {
	return v.value
}

func (v *NullableBoardItemLinks) Set(val *BoardItemLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableBoardItemLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableBoardItemLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoardItemLinks(val *BoardItemLinks) *NullableBoardItemLinks {
	return &NullableBoardItemLinks{value: val, isSet: true}
}

func (v NullableBoardItemLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoardItemLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


