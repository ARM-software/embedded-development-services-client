<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# AnalyserItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableAnalyserItemLinks**](AnalyserItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Clock** | Pointer to [**ClockInfo**](ClockInfo.md) |  | [optional] 
**Description** | Pointer to **string** | A detailed description of what the analyser does and its capabilities. | [optional] 
**Engine** | [**Engine**](Engine.md) |  | 
**ExecutionFlavour** | Pointer to **string** | The targeted execution flavour. | [optional] 
**HardwareTargets** | [**[]HardwareTarget**](HardwareTarget.md) | List of hardware targets that this analyser can evaluate across supported execution modes. | 
**Ip** | Pointer to [**AnalyserIP**](AnalyserIP.md) |  | [optional] 
**Name** | **string** | Unique system identifier. | [readonly] 
**PartnerSpecific** | Pointer to [**PartnerSpecific**](PartnerSpecific.md) |  | [optional] 
**Profile** | [**Profile**](Profile.md) |  | 
**SupportedModelFormats** | [**[]ModelFormat**](ModelFormat.md) | List of model formats supported by this analyser. | 
**Title** | **string** | A human-friendly display title for the analyser. | 
**Tool** | [**Tool**](Tool.md) |  | 
**Version** | **string** | Semantic version of the analyser. | 

## Methods

### NewAnalyserItem

`func NewAnalyserItem(links NullableAnalyserItemLinks, metadata NullableCommonMetadata, engine Engine, hardwareTargets []HardwareTarget, name string, profile Profile, supportedModelFormats []ModelFormat, title string, tool Tool, version string, ) *AnalyserItem`

NewAnalyserItem instantiates a new AnalyserItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyserItemWithDefaults

`func NewAnalyserItemWithDefaults() *AnalyserItem`

NewAnalyserItemWithDefaults instantiates a new AnalyserItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AnalyserItem) GetLinks() AnalyserItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AnalyserItem) GetLinksOk() (*AnalyserItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AnalyserItem) SetLinks(v AnalyserItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *AnalyserItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AnalyserItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *AnalyserItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AnalyserItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AnalyserItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *AnalyserItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AnalyserItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetClock

`func (o *AnalyserItem) GetClock() ClockInfo`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *AnalyserItem) GetClockOk() (*ClockInfo, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *AnalyserItem) SetClock(v ClockInfo)`

SetClock sets Clock field to given value.

### HasClock

`func (o *AnalyserItem) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetDescription

`func (o *AnalyserItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AnalyserItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AnalyserItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AnalyserItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEngine

`func (o *AnalyserItem) GetEngine() Engine`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *AnalyserItem) GetEngineOk() (*Engine, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *AnalyserItem) SetEngine(v Engine)`

SetEngine sets Engine field to given value.


### GetExecutionFlavour

`func (o *AnalyserItem) GetExecutionFlavour() string`

GetExecutionFlavour returns the ExecutionFlavour field if non-nil, zero value otherwise.

### GetExecutionFlavourOk

`func (o *AnalyserItem) GetExecutionFlavourOk() (*string, bool)`

GetExecutionFlavourOk returns a tuple with the ExecutionFlavour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionFlavour

`func (o *AnalyserItem) SetExecutionFlavour(v string)`

SetExecutionFlavour sets ExecutionFlavour field to given value.

### HasExecutionFlavour

`func (o *AnalyserItem) HasExecutionFlavour() bool`

HasExecutionFlavour returns a boolean if a field has been set.

### GetHardwareTargets

`func (o *AnalyserItem) GetHardwareTargets() []HardwareTarget`

GetHardwareTargets returns the HardwareTargets field if non-nil, zero value otherwise.

### GetHardwareTargetsOk

`func (o *AnalyserItem) GetHardwareTargetsOk() (*[]HardwareTarget, bool)`

GetHardwareTargetsOk returns a tuple with the HardwareTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareTargets

`func (o *AnalyserItem) SetHardwareTargets(v []HardwareTarget)`

SetHardwareTargets sets HardwareTargets field to given value.


### GetIp

`func (o *AnalyserItem) GetIp() AnalyserIP`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AnalyserItem) GetIpOk() (*AnalyserIP, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AnalyserItem) SetIp(v AnalyserIP)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AnalyserItem) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetName

`func (o *AnalyserItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnalyserItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnalyserItem) SetName(v string)`

SetName sets Name field to given value.


### GetPartnerSpecific

`func (o *AnalyserItem) GetPartnerSpecific() PartnerSpecific`

GetPartnerSpecific returns the PartnerSpecific field if non-nil, zero value otherwise.

### GetPartnerSpecificOk

`func (o *AnalyserItem) GetPartnerSpecificOk() (*PartnerSpecific, bool)`

GetPartnerSpecificOk returns a tuple with the PartnerSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerSpecific

`func (o *AnalyserItem) SetPartnerSpecific(v PartnerSpecific)`

SetPartnerSpecific sets PartnerSpecific field to given value.

### HasPartnerSpecific

`func (o *AnalyserItem) HasPartnerSpecific() bool`

HasPartnerSpecific returns a boolean if a field has been set.

### GetProfile

`func (o *AnalyserItem) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AnalyserItem) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AnalyserItem) SetProfile(v Profile)`

SetProfile sets Profile field to given value.


### GetSupportedModelFormats

`func (o *AnalyserItem) GetSupportedModelFormats() []ModelFormat`

GetSupportedModelFormats returns the SupportedModelFormats field if non-nil, zero value otherwise.

### GetSupportedModelFormatsOk

`func (o *AnalyserItem) GetSupportedModelFormatsOk() (*[]ModelFormat, bool)`

GetSupportedModelFormatsOk returns a tuple with the SupportedModelFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModelFormats

`func (o *AnalyserItem) SetSupportedModelFormats(v []ModelFormat)`

SetSupportedModelFormats sets SupportedModelFormats field to given value.


### GetTitle

`func (o *AnalyserItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AnalyserItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AnalyserItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTool

`func (o *AnalyserItem) GetTool() Tool`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *AnalyserItem) GetToolOk() (*Tool, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *AnalyserItem) SetTool(v Tool)`

SetTool sets Tool field to given value.


### GetVersion

`func (o *AnalyserItem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AnalyserItem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AnalyserItem) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


