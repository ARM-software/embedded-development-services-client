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

// IntellisenseMessageItem struct for IntellisenseMessageItem
type IntellisenseMessageItem struct {
	Links NullableHalFeedLinks `json:"_links"`
	Metadata NullablePagingMetadata `json:"_metadata"`
	// A list of build messages.
	Messages []MessageObject `json:"messages"`
	// ID of the Build Notification item.
	Name string `json:"name"`
	// Human readable name of the Build Notification item.
	Title *string `json:"title,omitempty"`
}

// NewIntellisenseMessageItem instantiates a new IntellisenseMessageItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntellisenseMessageItem(links NullableHalFeedLinks, metadata NullablePagingMetadata, messages []MessageObject, name string) *IntellisenseMessageItem {
	this := IntellisenseMessageItem{}
	this.Links = links
	this.Metadata = metadata
	this.Messages = messages
	this.Name = name
	return &this
}

// NewIntellisenseMessageItemWithDefaults instantiates a new IntellisenseMessageItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntellisenseMessageItemWithDefaults() *IntellisenseMessageItem {
	this := IntellisenseMessageItem{}
	return &this
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for HalFeedLinks will be returned
func (o *IntellisenseMessageItem) GetLinks() HalFeedLinks {
	if o == nil || o.Links.Get() == nil {
		var ret HalFeedLinks
		return ret
	}

	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntellisenseMessageItem) GetLinksOk() (*HalFeedLinks, bool) {
	if o == nil {
    return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// SetLinks sets field value
func (o *IntellisenseMessageItem) SetLinks(v HalFeedLinks) {
	o.Links.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for PagingMetadata will be returned
func (o *IntellisenseMessageItem) GetMetadata() PagingMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret PagingMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntellisenseMessageItem) GetMetadataOk() (*PagingMetadata, bool) {
	if o == nil {
    return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *IntellisenseMessageItem) SetMetadata(v PagingMetadata) {
	o.Metadata.Set(&v)
}

// GetMessages returns the Messages field value
// If the value is explicit nil, the zero value for []MessageObject will be returned
func (o *IntellisenseMessageItem) GetMessages() []MessageObject {
	if o == nil {
		var ret []MessageObject
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntellisenseMessageItem) GetMessagesOk() ([]MessageObject, bool) {
	if o == nil || isNil(o.Messages) {
    return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *IntellisenseMessageItem) SetMessages(v []MessageObject) {
	o.Messages = v
}

// GetName returns the Name field value
func (o *IntellisenseMessageItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IntellisenseMessageItem) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IntellisenseMessageItem) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *IntellisenseMessageItem) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntellisenseMessageItem) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *IntellisenseMessageItem) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *IntellisenseMessageItem) SetTitle(v string) {
	o.Title = &v
}

func (o IntellisenseMessageItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_links"] = o.Links.Get()
	}
	if true {
		toSerialize["_metadata"] = o.Metadata.Get()
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableIntellisenseMessageItem struct {
	value *IntellisenseMessageItem
	isSet bool
}

func (v NullableIntellisenseMessageItem) Get() *IntellisenseMessageItem {
	return v.value
}

func (v *NullableIntellisenseMessageItem) Set(val *IntellisenseMessageItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIntellisenseMessageItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIntellisenseMessageItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntellisenseMessageItem(val *IntellisenseMessageItem) *NullableIntellisenseMessageItem {
	return &NullableIntellisenseMessageItem{value: val, isSet: true}
}

func (v NullableIntellisenseMessageItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntellisenseMessageItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


