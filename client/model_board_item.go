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

// BoardItem struct for BoardItem
type BoardItem struct {
	Links BoardItemLinks `json:"_links"`
	Metadata NullableCommonMetadata `json:"_metadata"`
	// Array of debug interfaces for the Board.
	DebugInterfaces []ADebugInterface `json:"debug_interfaces,omitempty"`
	// Description of the Board.
	Description string `json:"description"`
	// Array of features for the Board.
	Features []BoardItemFeaturesInner `json:"features,omitempty"`
	// Unique ID of the Board.
	Id string `json:"id"`
	// Array of mounted devices for the Board.
	MountedDevices []DeviceItem `json:"mounted_devices,omitempty"`
	// Name of the Board.
	Name string `json:"name"`
	// Revision of the Board.
	Revision string `json:"revision"`
	// Brief summary of the Board.
	Summary string `json:"summary"`
}

// NewBoardItem instantiates a new BoardItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoardItem(links BoardItemLinks, metadata NullableCommonMetadata, description string, id string, name string, revision string, summary string) *BoardItem {
	this := BoardItem{}
	this.Links = links
	this.Metadata = metadata
	this.Description = description
	this.Id = id
	this.Name = name
	this.Revision = revision
	this.Summary = summary
	return &this
}

// NewBoardItemWithDefaults instantiates a new BoardItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoardItemWithDefaults() *BoardItem {
	this := BoardItem{}
	return &this
}

// GetLinks returns the Links field value
func (o *BoardItem) GetLinks() BoardItemLinks {
	if o == nil {
		var ret BoardItemLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BoardItem) GetLinksOk() (*BoardItemLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BoardItem) SetLinks(v BoardItemLinks) {
	o.Links = v
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for CommonMetadata will be returned
func (o *BoardItem) GetMetadata() CommonMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret CommonMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoardItem) GetMetadataOk() (*CommonMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *BoardItem) SetMetadata(v CommonMetadata) {
	o.Metadata.Set(&v)
}

// GetDebugInterfaces returns the DebugInterfaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoardItem) GetDebugInterfaces() []ADebugInterface {
	if o == nil {
		var ret []ADebugInterface
		return ret
	}
	return o.DebugInterfaces
}

// GetDebugInterfacesOk returns a tuple with the DebugInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoardItem) GetDebugInterfacesOk() ([]ADebugInterface, bool) {
	if o == nil || o.DebugInterfaces == nil {
		return nil, false
	}
	return o.DebugInterfaces, true
}

// HasDebugInterfaces returns a boolean if a field has been set.
func (o *BoardItem) HasDebugInterfaces() bool {
	if o != nil && o.DebugInterfaces != nil {
		return true
	}

	return false
}

// SetDebugInterfaces gets a reference to the given []ADebugInterface and assigns it to the DebugInterfaces field.
func (o *BoardItem) SetDebugInterfaces(v []ADebugInterface) {
	o.DebugInterfaces = v
}

// GetDescription returns the Description field value
func (o *BoardItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *BoardItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *BoardItem) SetDescription(v string) {
	o.Description = v
}

// GetFeatures returns the Features field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoardItem) GetFeatures() []BoardItemFeaturesInner {
	if o == nil {
		var ret []BoardItemFeaturesInner
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoardItem) GetFeaturesOk() ([]BoardItemFeaturesInner, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *BoardItem) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []BoardItemFeaturesInner and assigns it to the Features field.
func (o *BoardItem) SetFeatures(v []BoardItemFeaturesInner) {
	o.Features = v
}

// GetId returns the Id field value
func (o *BoardItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BoardItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BoardItem) SetId(v string) {
	o.Id = v
}

// GetMountedDevices returns the MountedDevices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoardItem) GetMountedDevices() []DeviceItem {
	if o == nil {
		var ret []DeviceItem
		return ret
	}
	return o.MountedDevices
}

// GetMountedDevicesOk returns a tuple with the MountedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoardItem) GetMountedDevicesOk() ([]DeviceItem, bool) {
	if o == nil || o.MountedDevices == nil {
		return nil, false
	}
	return o.MountedDevices, true
}

// HasMountedDevices returns a boolean if a field has been set.
func (o *BoardItem) HasMountedDevices() bool {
	if o != nil && o.MountedDevices != nil {
		return true
	}

	return false
}

// SetMountedDevices gets a reference to the given []DeviceItem and assigns it to the MountedDevices field.
func (o *BoardItem) SetMountedDevices(v []DeviceItem) {
	o.MountedDevices = v
}

// GetName returns the Name field value
func (o *BoardItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BoardItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BoardItem) SetName(v string) {
	o.Name = v
}

// GetRevision returns the Revision field value
func (o *BoardItem) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *BoardItem) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *BoardItem) SetRevision(v string) {
	o.Revision = v
}

// GetSummary returns the Summary field value
func (o *BoardItem) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *BoardItem) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *BoardItem) SetSummary(v string) {
	o.Summary = v
}

func (o BoardItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["_metadata"] = o.Metadata.Get()
	}
	if o.DebugInterfaces != nil {
		toSerialize["debug_interfaces"] = o.DebugInterfaces
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.MountedDevices != nil {
		toSerialize["mounted_devices"] = o.MountedDevices
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["revision"] = o.Revision
	}
	if true {
		toSerialize["summary"] = o.Summary
	}
	return json.Marshal(toSerialize)
}

type NullableBoardItem struct {
	value *BoardItem
	isSet bool
}

func (v NullableBoardItem) Get() *BoardItem {
	return v.value
}

func (v *NullableBoardItem) Set(val *BoardItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBoardItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBoardItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoardItem(val *BoardItem) *NullableBoardItem {
	return &NullableBoardItem{value: val, isSet: true}
}

func (v NullableBoardItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoardItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

