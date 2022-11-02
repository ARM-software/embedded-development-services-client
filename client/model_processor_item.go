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

// ProcessorItem struct for ProcessorItem
type ProcessorItem struct {
	// Name of the processor core
	Core string `json:"core"`
	// Core Revision of the Processor
	CoreRevision NullableString `json:"core_revision,omitempty"`
	// CortexM Vendor Extension support for the Processor
	CortexmVectorExtensions NullableString `json:"cortexm_vector_extensions,omitempty"`
	// Digital Signal Processor of the Processor
	DigitalSignalProcessor NullableString `json:"digital_signal_processor,omitempty"`
	// Endian for the Processor
	Endian NullableString `json:"endian,omitempty"`
	// Floating Point Unit of the Processor
	FloatingPointUnit NullableString `json:"floating_point_unit,omitempty"`
	// Maximum Core Clock Frequency of the Processor
	MaxCoreClockFrequency NullableInt32 `json:"max_core_clock_frequency,omitempty"`
	// Memory Protection Unit of the Processor
	MemoryProtectionUnit NullableString `json:"memory_protection_unit,omitempty"`
	// Human-readable name of the processor
	Title NullableString `json:"title,omitempty"`
	// Trust Zone of the Processor
	TrustZone NullableString `json:"trust_zone,omitempty"`
	// Number of processing units in a symmetric multi-processor core.
	Units NullableInt32 `json:"units,omitempty"`
}

// NewProcessorItem instantiates a new ProcessorItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorItem(core string) *ProcessorItem {
	this := ProcessorItem{}
	this.Core = core
	return &this
}

// NewProcessorItemWithDefaults instantiates a new ProcessorItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorItemWithDefaults() *ProcessorItem {
	this := ProcessorItem{}
	return &this
}

// GetCore returns the Core field value
func (o *ProcessorItem) GetCore() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Core
}

// GetCoreOk returns a tuple with the Core field value
// and a boolean to check if the value has been set.
func (o *ProcessorItem) GetCoreOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Core, true
}

// SetCore sets field value
func (o *ProcessorItem) SetCore(v string) {
	o.Core = v
}

// GetCoreRevision returns the CoreRevision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorItem) GetCoreRevision() string {
	if o == nil || isNil(o.CoreRevision.Get()) {
		var ret string
		return ret
	}
	return *o.CoreRevision.Get()
}

// GetCoreRevisionOk returns a tuple with the CoreRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorItem) GetCoreRevisionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CoreRevision.Get(), o.CoreRevision.IsSet()
}

// HasCoreRevision returns a boolean if a field has been set.
func (o *ProcessorItem) HasCoreRevision() bool {
	if o != nil && o.CoreRevision.IsSet() {
		return true
	}

	return false
}

// SetCoreRevision gets a reference to the given NullableString and assigns it to the CoreRevision field.
func (o *ProcessorItem) SetCoreRevision(v string) {
	o.CoreRevision.Set(&v)
}
// SetCoreRevisionNil sets the value for CoreRevision to be an explicit nil
func (o *ProcessorItem) SetCoreRevisionNil() {
	o.CoreRevision.Set(nil)
}

// UnsetCoreRevision ensures that no value is present for CoreRevision, not even an explicit nil
func (o *ProcessorItem) UnsetCoreRevision() {
	o.CoreRevision.Unset()
}

// GetCortexmVectorExtensions returns the CortexmVectorExtensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorItem) GetCortexmVectorExtensions() string {
	if o == nil || isNil(o.CortexmVectorExtensions.Get()) {
		var ret string
		return ret
	}
	return *o.CortexmVectorExtensions.Get()
}

// GetCortexmVectorExtensionsOk returns a tuple with the CortexmVectorExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorItem) GetCortexmVectorExtensionsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CortexmVectorExtensions.Get(), o.CortexmVectorExtensions.IsSet()
}

// HasCortexmVectorExtensions returns a boolean if a field has been set.
func (o *ProcessorItem) HasCortexmVectorExtensions() bool {
	if o != nil && o.CortexmVectorExtensions.IsSet() {
		return true
	}

	return false
}

// SetCortexmVectorExtensions gets a reference to the given NullableString and assigns it to the CortexmVectorExtensions field.
func (o *ProcessorItem) SetCortexmVectorExtensions(v string) {
	o.CortexmVectorExtensions.Set(&v)
}
// SetCortexmVectorExtensionsNil sets the value for CortexmVectorExtensions to be an explicit nil
func (o *ProcessorItem) SetCortexmVectorExtensionsNil() {
	o.CortexmVectorExtensions.Set(nil)
}

// UnsetCortexmVectorExtensions ensures that no value is present for CortexmVectorExtensions, not even an explicit nil
func (o *ProcessorItem) UnsetCortexmVectorExtensions() {
	o.CortexmVectorExtensions.Unset()
}

// GetDigitalSignalProcessor returns the DigitalSignalProcessor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorItem) GetDigitalSignalProcessor() string {
	if o == nil || isNil(o.DigitalSignalProcessor.Get()) {
		var ret string
		return ret
	}
	return *o.DigitalSignalProcessor.Get()
}

// GetDigitalSignalProcessorOk returns a tuple with the DigitalSignalProcessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorItem) GetDigitalSignalProcessorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DigitalSignalProcessor.Get(), o.DigitalSignalProcessor.IsSet()
}

// HasDigitalSignalProcessor returns a boolean if a field has been set.
func (o *ProcessorItem) HasDigitalSignalProcessor() bool {
	if o != nil && o.DigitalSignalProcessor.IsSet() {
		return true
	}

	return false
}

// SetDigitalSignalProcessor gets a reference to the given NullableString and assigns it to the DigitalSignalProcessor field.
func (o *ProcessorItem) SetDigitalSignalProcessor(v string) {
	o.DigitalSignalProcessor.Set(&v)
}
// SetDigitalSignalProcessorNil sets the value for DigitalSignalProcessor to be an explicit nil
func (o *ProcessorItem) SetDigitalSignalProcessorNil() {
	o.DigitalSignalProcessor.Set(nil)
}

// UnsetDigitalSignalProcessor ensures that no value is present for DigitalSignalProcessor, not even an explicit nil
func (o *ProcessorItem) UnsetDigitalSignalProcessor() {
	o.DigitalSignalProcessor.Unset()
}

// GetEndian returns the Endian field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorItem) GetEndian() string {
	if o == nil || isNil(o.Endian.Get()) {
		var ret string
		return ret
	}
	return *o.Endian.Get()
}

// GetEndianOk returns a tuple with the Endian field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorItem) GetEndianOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Endian.Get(), o.Endian.IsSet()
}

// HasEndian returns a boolean if a field has been set.
func (o *ProcessorItem) HasEndian() bool {
	if o != nil && o.Endian.IsSet() {
		return true
	}

	return false
}

// SetEndian gets a reference to the given NullableString and assigns it to the Endian field.
func (o *ProcessorItem) SetEndian(v string) {
	o.Endian.Set(&v)
}
// SetEndianNil sets the value for Endian to be an explicit nil
func (o *ProcessorItem) SetEndianNil() {
	o.Endian.Set(nil)
}

// UnsetEndian ensures that no value is present for Endian, not even an explicit nil
func (o *ProcessorItem) UnsetEndian() {
	o.Endian.Unset()
}

// GetFloatingPointUnit returns the FloatingPointUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorItem) GetFloatingPointUnit() string {
	if o == nil || isNil(o.FloatingPointUnit.Get()) {
		var ret string
		return ret
	}
	return *o.FloatingPointUnit.Get()
}

// GetFloatingPointUnitOk returns a tuple with the FloatingPointUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorItem) GetFloatingPointUnitOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.FloatingPointUnit.Get(), o.FloatingPointUnit.IsSet()
}

// HasFloatingPointUnit returns a boolean if a field has been set.
func (o *ProcessorItem) HasFloatingPointUnit() bool {
	if o != nil && o.FloatingPointUnit.IsSet() {
		return true
	}

	return false
}

// SetFloatingPointUnit gets a reference to the given NullableString and assigns it to the FloatingPointUnit field.
func (o *ProcessorItem) SetFloatingPointUnit(v string) {
	o.FloatingPointUnit.Set(&v)
}
// SetFloatingPointUnitNil sets the value for FloatingPointUnit to be an explicit nil
func (o *ProcessorItem) SetFloatingPointUnitNil() {
	o.FloatingPointUnit.Set(nil)
}

// UnsetFloatingPointUnit ensures that no value is present for FloatingPointUnit, not even an explicit nil
func (o *ProcessorItem) UnsetFloatingPointUnit() {
	o.FloatingPointUnit.Unset()
}

// GetMaxCoreClockFrequency returns the MaxCoreClockFrequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorItem) GetMaxCoreClockFrequency() int32 {
	if o == nil || isNil(o.MaxCoreClockFrequency.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxCoreClockFrequency.Get()
}

// GetMaxCoreClockFrequencyOk returns a tuple with the MaxCoreClockFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorItem) GetMaxCoreClockFrequencyOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxCoreClockFrequency.Get(), o.MaxCoreClockFrequency.IsSet()
}

// HasMaxCoreClockFrequency returns a boolean if a field has been set.
func (o *ProcessorItem) HasMaxCoreClockFrequency() bool {
	if o != nil && o.MaxCoreClockFrequency.IsSet() {
		return true
	}

	return false
}

// SetMaxCoreClockFrequency gets a reference to the given NullableInt32 and assigns it to the MaxCoreClockFrequency field.
func (o *ProcessorItem) SetMaxCoreClockFrequency(v int32) {
	o.MaxCoreClockFrequency.Set(&v)
}
// SetMaxCoreClockFrequencyNil sets the value for MaxCoreClockFrequency to be an explicit nil
func (o *ProcessorItem) SetMaxCoreClockFrequencyNil() {
	o.MaxCoreClockFrequency.Set(nil)
}

// UnsetMaxCoreClockFrequency ensures that no value is present for MaxCoreClockFrequency, not even an explicit nil
func (o *ProcessorItem) UnsetMaxCoreClockFrequency() {
	o.MaxCoreClockFrequency.Unset()
}

// GetMemoryProtectionUnit returns the MemoryProtectionUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorItem) GetMemoryProtectionUnit() string {
	if o == nil || isNil(o.MemoryProtectionUnit.Get()) {
		var ret string
		return ret
	}
	return *o.MemoryProtectionUnit.Get()
}

// GetMemoryProtectionUnitOk returns a tuple with the MemoryProtectionUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorItem) GetMemoryProtectionUnitOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MemoryProtectionUnit.Get(), o.MemoryProtectionUnit.IsSet()
}

// HasMemoryProtectionUnit returns a boolean if a field has been set.
func (o *ProcessorItem) HasMemoryProtectionUnit() bool {
	if o != nil && o.MemoryProtectionUnit.IsSet() {
		return true
	}

	return false
}

// SetMemoryProtectionUnit gets a reference to the given NullableString and assigns it to the MemoryProtectionUnit field.
func (o *ProcessorItem) SetMemoryProtectionUnit(v string) {
	o.MemoryProtectionUnit.Set(&v)
}
// SetMemoryProtectionUnitNil sets the value for MemoryProtectionUnit to be an explicit nil
func (o *ProcessorItem) SetMemoryProtectionUnitNil() {
	o.MemoryProtectionUnit.Set(nil)
}

// UnsetMemoryProtectionUnit ensures that no value is present for MemoryProtectionUnit, not even an explicit nil
func (o *ProcessorItem) UnsetMemoryProtectionUnit() {
	o.MemoryProtectionUnit.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorItem) GetTitle() string {
	if o == nil || isNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorItem) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ProcessorItem) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ProcessorItem) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ProcessorItem) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ProcessorItem) UnsetTitle() {
	o.Title.Unset()
}

// GetTrustZone returns the TrustZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorItem) GetTrustZone() string {
	if o == nil || isNil(o.TrustZone.Get()) {
		var ret string
		return ret
	}
	return *o.TrustZone.Get()
}

// GetTrustZoneOk returns a tuple with the TrustZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorItem) GetTrustZoneOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.TrustZone.Get(), o.TrustZone.IsSet()
}

// HasTrustZone returns a boolean if a field has been set.
func (o *ProcessorItem) HasTrustZone() bool {
	if o != nil && o.TrustZone.IsSet() {
		return true
	}

	return false
}

// SetTrustZone gets a reference to the given NullableString and assigns it to the TrustZone field.
func (o *ProcessorItem) SetTrustZone(v string) {
	o.TrustZone.Set(&v)
}
// SetTrustZoneNil sets the value for TrustZone to be an explicit nil
func (o *ProcessorItem) SetTrustZoneNil() {
	o.TrustZone.Set(nil)
}

// UnsetTrustZone ensures that no value is present for TrustZone, not even an explicit nil
func (o *ProcessorItem) UnsetTrustZone() {
	o.TrustZone.Unset()
}

// GetUnits returns the Units field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcessorItem) GetUnits() int32 {
	if o == nil || isNil(o.Units.Get()) {
		var ret int32
		return ret
	}
	return *o.Units.Get()
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcessorItem) GetUnitsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Units.Get(), o.Units.IsSet()
}

// HasUnits returns a boolean if a field has been set.
func (o *ProcessorItem) HasUnits() bool {
	if o != nil && o.Units.IsSet() {
		return true
	}

	return false
}

// SetUnits gets a reference to the given NullableInt32 and assigns it to the Units field.
func (o *ProcessorItem) SetUnits(v int32) {
	o.Units.Set(&v)
}
// SetUnitsNil sets the value for Units to be an explicit nil
func (o *ProcessorItem) SetUnitsNil() {
	o.Units.Set(nil)
}

// UnsetUnits ensures that no value is present for Units, not even an explicit nil
func (o *ProcessorItem) UnsetUnits() {
	o.Units.Unset()
}

func (o ProcessorItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["core"] = o.Core
	}
	if o.CoreRevision.IsSet() {
		toSerialize["core_revision"] = o.CoreRevision.Get()
	}
	if o.CortexmVectorExtensions.IsSet() {
		toSerialize["cortexm_vector_extensions"] = o.CortexmVectorExtensions.Get()
	}
	if o.DigitalSignalProcessor.IsSet() {
		toSerialize["digital_signal_processor"] = o.DigitalSignalProcessor.Get()
	}
	if o.Endian.IsSet() {
		toSerialize["endian"] = o.Endian.Get()
	}
	if o.FloatingPointUnit.IsSet() {
		toSerialize["floating_point_unit"] = o.FloatingPointUnit.Get()
	}
	if o.MaxCoreClockFrequency.IsSet() {
		toSerialize["max_core_clock_frequency"] = o.MaxCoreClockFrequency.Get()
	}
	if o.MemoryProtectionUnit.IsSet() {
		toSerialize["memory_protection_unit"] = o.MemoryProtectionUnit.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.TrustZone.IsSet() {
		toSerialize["trust_zone"] = o.TrustZone.Get()
	}
	if o.Units.IsSet() {
		toSerialize["units"] = o.Units.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorItem struct {
	value *ProcessorItem
	isSet bool
}

func (v NullableProcessorItem) Get() *ProcessorItem {
	return v.value
}

func (v *NullableProcessorItem) Set(val *ProcessorItem) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorItem) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorItem(val *ProcessorItem) *NullableProcessorItem {
	return &NullableProcessorItem{value: val, isSet: true}
}

func (v NullableProcessorItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


