# ProcessorItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Core** | **string** | Name of the processor core | 
**CoreRevision** | Pointer to **NullableString** | Core Revision of the Processor | [optional] 
**CortexmVectorExtensions** | Pointer to **NullableString** | CortexM Vendor Extension support for the Processor | [optional] 
**DigitalSignalProcessor** | Pointer to **NullableString** | Digital Signal Processor of the Processor | [optional] 
**Endian** | Pointer to **NullableString** | Endian for the Processor | [optional] 
**FloatingPointUnit** | Pointer to **NullableString** | Floating Point Unit of the Processor | [optional] 
**MaxCoreClockFrequency** | Pointer to **NullableInt32** | Maximum Core Clock Frequency of the Processor | [optional] 
**MemoryProtectionUnit** | Pointer to **NullableString** | Memory Protection Unit of the Processor | [optional] 
**Title** | Pointer to **NullableString** | Human-readable name of the processor | [optional] 
**TrustZone** | Pointer to **NullableString** | Trust Zone of the Processor | [optional] 
**Units** | Pointer to **NullableInt32** | Number of processing units in a symmetric multi-processor core. | [optional] 

## Methods

### NewProcessorItem

`func NewProcessorItem(core string, ) *ProcessorItem`

NewProcessorItem instantiates a new ProcessorItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorItemWithDefaults

`func NewProcessorItemWithDefaults() *ProcessorItem`

NewProcessorItemWithDefaults instantiates a new ProcessorItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCore

`func (o *ProcessorItem) GetCore() string`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *ProcessorItem) GetCoreOk() (*string, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *ProcessorItem) SetCore(v string)`

SetCore sets Core field to given value.


### GetCoreRevision

`func (o *ProcessorItem) GetCoreRevision() string`

GetCoreRevision returns the CoreRevision field if non-nil, zero value otherwise.

### GetCoreRevisionOk

`func (o *ProcessorItem) GetCoreRevisionOk() (*string, bool)`

GetCoreRevisionOk returns a tuple with the CoreRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreRevision

`func (o *ProcessorItem) SetCoreRevision(v string)`

SetCoreRevision sets CoreRevision field to given value.

### HasCoreRevision

`func (o *ProcessorItem) HasCoreRevision() bool`

HasCoreRevision returns a boolean if a field has been set.

### SetCoreRevisionNil

`func (o *ProcessorItem) SetCoreRevisionNil(b bool)`

 SetCoreRevisionNil sets the value for CoreRevision to be an explicit nil

### UnsetCoreRevision
`func (o *ProcessorItem) UnsetCoreRevision()`

UnsetCoreRevision ensures that no value is present for CoreRevision, not even an explicit nil
### GetCortexmVectorExtensions

`func (o *ProcessorItem) GetCortexmVectorExtensions() string`

GetCortexmVectorExtensions returns the CortexmVectorExtensions field if non-nil, zero value otherwise.

### GetCortexmVectorExtensionsOk

`func (o *ProcessorItem) GetCortexmVectorExtensionsOk() (*string, bool)`

GetCortexmVectorExtensionsOk returns a tuple with the CortexmVectorExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCortexmVectorExtensions

`func (o *ProcessorItem) SetCortexmVectorExtensions(v string)`

SetCortexmVectorExtensions sets CortexmVectorExtensions field to given value.

### HasCortexmVectorExtensions

`func (o *ProcessorItem) HasCortexmVectorExtensions() bool`

HasCortexmVectorExtensions returns a boolean if a field has been set.

### SetCortexmVectorExtensionsNil

`func (o *ProcessorItem) SetCortexmVectorExtensionsNil(b bool)`

 SetCortexmVectorExtensionsNil sets the value for CortexmVectorExtensions to be an explicit nil

### UnsetCortexmVectorExtensions
`func (o *ProcessorItem) UnsetCortexmVectorExtensions()`

UnsetCortexmVectorExtensions ensures that no value is present for CortexmVectorExtensions, not even an explicit nil
### GetDigitalSignalProcessor

`func (o *ProcessorItem) GetDigitalSignalProcessor() string`

GetDigitalSignalProcessor returns the DigitalSignalProcessor field if non-nil, zero value otherwise.

### GetDigitalSignalProcessorOk

`func (o *ProcessorItem) GetDigitalSignalProcessorOk() (*string, bool)`

GetDigitalSignalProcessorOk returns a tuple with the DigitalSignalProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalSignalProcessor

`func (o *ProcessorItem) SetDigitalSignalProcessor(v string)`

SetDigitalSignalProcessor sets DigitalSignalProcessor field to given value.

### HasDigitalSignalProcessor

`func (o *ProcessorItem) HasDigitalSignalProcessor() bool`

HasDigitalSignalProcessor returns a boolean if a field has been set.

### SetDigitalSignalProcessorNil

`func (o *ProcessorItem) SetDigitalSignalProcessorNil(b bool)`

 SetDigitalSignalProcessorNil sets the value for DigitalSignalProcessor to be an explicit nil

### UnsetDigitalSignalProcessor
`func (o *ProcessorItem) UnsetDigitalSignalProcessor()`

UnsetDigitalSignalProcessor ensures that no value is present for DigitalSignalProcessor, not even an explicit nil
### GetEndian

`func (o *ProcessorItem) GetEndian() string`

GetEndian returns the Endian field if non-nil, zero value otherwise.

### GetEndianOk

`func (o *ProcessorItem) GetEndianOk() (*string, bool)`

GetEndianOk returns a tuple with the Endian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndian

`func (o *ProcessorItem) SetEndian(v string)`

SetEndian sets Endian field to given value.

### HasEndian

`func (o *ProcessorItem) HasEndian() bool`

HasEndian returns a boolean if a field has been set.

### SetEndianNil

`func (o *ProcessorItem) SetEndianNil(b bool)`

 SetEndianNil sets the value for Endian to be an explicit nil

### UnsetEndian
`func (o *ProcessorItem) UnsetEndian()`

UnsetEndian ensures that no value is present for Endian, not even an explicit nil
### GetFloatingPointUnit

`func (o *ProcessorItem) GetFloatingPointUnit() string`

GetFloatingPointUnit returns the FloatingPointUnit field if non-nil, zero value otherwise.

### GetFloatingPointUnitOk

`func (o *ProcessorItem) GetFloatingPointUnitOk() (*string, bool)`

GetFloatingPointUnitOk returns a tuple with the FloatingPointUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingPointUnit

`func (o *ProcessorItem) SetFloatingPointUnit(v string)`

SetFloatingPointUnit sets FloatingPointUnit field to given value.

### HasFloatingPointUnit

`func (o *ProcessorItem) HasFloatingPointUnit() bool`

HasFloatingPointUnit returns a boolean if a field has been set.

### SetFloatingPointUnitNil

`func (o *ProcessorItem) SetFloatingPointUnitNil(b bool)`

 SetFloatingPointUnitNil sets the value for FloatingPointUnit to be an explicit nil

### UnsetFloatingPointUnit
`func (o *ProcessorItem) UnsetFloatingPointUnit()`

UnsetFloatingPointUnit ensures that no value is present for FloatingPointUnit, not even an explicit nil
### GetMaxCoreClockFrequency

`func (o *ProcessorItem) GetMaxCoreClockFrequency() int32`

GetMaxCoreClockFrequency returns the MaxCoreClockFrequency field if non-nil, zero value otherwise.

### GetMaxCoreClockFrequencyOk

`func (o *ProcessorItem) GetMaxCoreClockFrequencyOk() (*int32, bool)`

GetMaxCoreClockFrequencyOk returns a tuple with the MaxCoreClockFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCoreClockFrequency

`func (o *ProcessorItem) SetMaxCoreClockFrequency(v int32)`

SetMaxCoreClockFrequency sets MaxCoreClockFrequency field to given value.

### HasMaxCoreClockFrequency

`func (o *ProcessorItem) HasMaxCoreClockFrequency() bool`

HasMaxCoreClockFrequency returns a boolean if a field has been set.

### SetMaxCoreClockFrequencyNil

`func (o *ProcessorItem) SetMaxCoreClockFrequencyNil(b bool)`

 SetMaxCoreClockFrequencyNil sets the value for MaxCoreClockFrequency to be an explicit nil

### UnsetMaxCoreClockFrequency
`func (o *ProcessorItem) UnsetMaxCoreClockFrequency()`

UnsetMaxCoreClockFrequency ensures that no value is present for MaxCoreClockFrequency, not even an explicit nil
### GetMemoryProtectionUnit

`func (o *ProcessorItem) GetMemoryProtectionUnit() string`

GetMemoryProtectionUnit returns the MemoryProtectionUnit field if non-nil, zero value otherwise.

### GetMemoryProtectionUnitOk

`func (o *ProcessorItem) GetMemoryProtectionUnitOk() (*string, bool)`

GetMemoryProtectionUnitOk returns a tuple with the MemoryProtectionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryProtectionUnit

`func (o *ProcessorItem) SetMemoryProtectionUnit(v string)`

SetMemoryProtectionUnit sets MemoryProtectionUnit field to given value.

### HasMemoryProtectionUnit

`func (o *ProcessorItem) HasMemoryProtectionUnit() bool`

HasMemoryProtectionUnit returns a boolean if a field has been set.

### SetMemoryProtectionUnitNil

`func (o *ProcessorItem) SetMemoryProtectionUnitNil(b bool)`

 SetMemoryProtectionUnitNil sets the value for MemoryProtectionUnit to be an explicit nil

### UnsetMemoryProtectionUnit
`func (o *ProcessorItem) UnsetMemoryProtectionUnit()`

UnsetMemoryProtectionUnit ensures that no value is present for MemoryProtectionUnit, not even an explicit nil
### GetTitle

`func (o *ProcessorItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProcessorItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProcessorItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProcessorItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ProcessorItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ProcessorItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTrustZone

`func (o *ProcessorItem) GetTrustZone() string`

GetTrustZone returns the TrustZone field if non-nil, zero value otherwise.

### GetTrustZoneOk

`func (o *ProcessorItem) GetTrustZoneOk() (*string, bool)`

GetTrustZoneOk returns a tuple with the TrustZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustZone

`func (o *ProcessorItem) SetTrustZone(v string)`

SetTrustZone sets TrustZone field to given value.

### HasTrustZone

`func (o *ProcessorItem) HasTrustZone() bool`

HasTrustZone returns a boolean if a field has been set.

### SetTrustZoneNil

`func (o *ProcessorItem) SetTrustZoneNil(b bool)`

 SetTrustZoneNil sets the value for TrustZone to be an explicit nil

### UnsetTrustZone
`func (o *ProcessorItem) UnsetTrustZone()`

UnsetTrustZone ensures that no value is present for TrustZone, not even an explicit nil
### GetUnits

`func (o *ProcessorItem) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ProcessorItem) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ProcessorItem) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ProcessorItem) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnitsNil

`func (o *ProcessorItem) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *ProcessorItem) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


