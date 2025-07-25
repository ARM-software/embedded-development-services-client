<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableFPGAItemLinks**](FPGAItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**AccessibleRepositories** | Pointer to **[]string** | A list of repositories accessible from the FPGA host | [optional] [readonly] 
**AgentAuthToken** | Pointer to **string** | Hybrid RSA AES encoded token for the agent managing this specific FPGA host | [optional] 
**AgentHost** | Pointer to **string** | The name of the host machine where the agent controlling this FPGA is running | [optional] 
**Busy** | **bool** | True if this FPGA is currently busy running a job. If this case, the &#x60;current&#x60; link should be accessible to give more details about the on-going job. | [readonly] 
**Deprecated** | **bool** | True if this FPGA is scheduled to be removed from the service. | [readonly] 
**DeprecationInfo** | Pointer to [**DeprecationInfo**](DeprecationInfo.md) |  | [optional] 
**Description** | **string** | Description of the FPGA configuration | [readonly] 
**EntitlementGroups** | Pointer to **[]string** | A list of entitlement groups the FPGA is associated with. | [optional] 
**ExtraMetadata** | Pointer to **map[string]string** | Extra metadata describing FPGAs. | [optional] 
**Name** | **string** | Unique ID of this FPGA | [readonly] 
**ReadUARTConnections** | Pointer to [**[]FPGAUARTConnection**](FPGAUARTConnection.md) | A list of UART connections that will be read from during a job | [optional] 
**Ready** | **bool** | True if this FPGA is ready to accept jobs to run. | [readonly] 
**Repositories** | Pointer to [**FPGAPayloadRepositories**](FPGAPayloadRepositories.md) |  | [optional] 
**SupportedTargets** | [**[]FPGATarget**](FPGATarget.md) | A list of targets supported for this FPGA | 
**Title** | **string** | Human readable name of the FPGA. | 
**WorkerAuthToken** | Pointer to **string** | Hybrid RSA AES encoded worker authorisation token to use for authorising with the boundary controller | [optional] 
**WriteUARTConnection** | Pointer to [**FPGAUARTConnection**](FPGAUARTConnection.md) |  | [optional] 

## Methods

### NewFPGAItem

`func NewFPGAItem(links NullableFPGAItemLinks, metadata NullableCommonMetadata, busy bool, deprecated bool, description string, name string, ready bool, supportedTargets []FPGATarget, title string, ) *FPGAItem`

NewFPGAItem instantiates a new FPGAItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAItemWithDefaults

`func NewFPGAItemWithDefaults() *FPGAItem`

NewFPGAItemWithDefaults instantiates a new FPGAItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FPGAItem) GetLinks() FPGAItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FPGAItem) GetLinksOk() (*FPGAItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FPGAItem) SetLinks(v FPGAItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *FPGAItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *FPGAItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *FPGAItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FPGAItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FPGAItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *FPGAItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *FPGAItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAccessibleRepositories

`func (o *FPGAItem) GetAccessibleRepositories() []string`

GetAccessibleRepositories returns the AccessibleRepositories field if non-nil, zero value otherwise.

### GetAccessibleRepositoriesOk

`func (o *FPGAItem) GetAccessibleRepositoriesOk() (*[]string, bool)`

GetAccessibleRepositoriesOk returns a tuple with the AccessibleRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibleRepositories

`func (o *FPGAItem) SetAccessibleRepositories(v []string)`

SetAccessibleRepositories sets AccessibleRepositories field to given value.

### HasAccessibleRepositories

`func (o *FPGAItem) HasAccessibleRepositories() bool`

HasAccessibleRepositories returns a boolean if a field has been set.

### GetAgentAuthToken

`func (o *FPGAItem) GetAgentAuthToken() string`

GetAgentAuthToken returns the AgentAuthToken field if non-nil, zero value otherwise.

### GetAgentAuthTokenOk

`func (o *FPGAItem) GetAgentAuthTokenOk() (*string, bool)`

GetAgentAuthTokenOk returns a tuple with the AgentAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAuthToken

`func (o *FPGAItem) SetAgentAuthToken(v string)`

SetAgentAuthToken sets AgentAuthToken field to given value.

### HasAgentAuthToken

`func (o *FPGAItem) HasAgentAuthToken() bool`

HasAgentAuthToken returns a boolean if a field has been set.

### GetAgentHost

`func (o *FPGAItem) GetAgentHost() string`

GetAgentHost returns the AgentHost field if non-nil, zero value otherwise.

### GetAgentHostOk

`func (o *FPGAItem) GetAgentHostOk() (*string, bool)`

GetAgentHostOk returns a tuple with the AgentHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentHost

`func (o *FPGAItem) SetAgentHost(v string)`

SetAgentHost sets AgentHost field to given value.

### HasAgentHost

`func (o *FPGAItem) HasAgentHost() bool`

HasAgentHost returns a boolean if a field has been set.

### GetBusy

`func (o *FPGAItem) GetBusy() bool`

GetBusy returns the Busy field if non-nil, zero value otherwise.

### GetBusyOk

`func (o *FPGAItem) GetBusyOk() (*bool, bool)`

GetBusyOk returns a tuple with the Busy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusy

`func (o *FPGAItem) SetBusy(v bool)`

SetBusy sets Busy field to given value.


### GetDeprecated

`func (o *FPGAItem) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *FPGAItem) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *FPGAItem) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.


### GetDeprecationInfo

`func (o *FPGAItem) GetDeprecationInfo() DeprecationInfo`

GetDeprecationInfo returns the DeprecationInfo field if non-nil, zero value otherwise.

### GetDeprecationInfoOk

`func (o *FPGAItem) GetDeprecationInfoOk() (*DeprecationInfo, bool)`

GetDeprecationInfoOk returns a tuple with the DeprecationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationInfo

`func (o *FPGAItem) SetDeprecationInfo(v DeprecationInfo)`

SetDeprecationInfo sets DeprecationInfo field to given value.

### HasDeprecationInfo

`func (o *FPGAItem) HasDeprecationInfo() bool`

HasDeprecationInfo returns a boolean if a field has been set.

### GetDescription

`func (o *FPGAItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FPGAItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FPGAItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEntitlementGroups

`func (o *FPGAItem) GetEntitlementGroups() []string`

GetEntitlementGroups returns the EntitlementGroups field if non-nil, zero value otherwise.

### GetEntitlementGroupsOk

`func (o *FPGAItem) GetEntitlementGroupsOk() (*[]string, bool)`

GetEntitlementGroupsOk returns a tuple with the EntitlementGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementGroups

`func (o *FPGAItem) SetEntitlementGroups(v []string)`

SetEntitlementGroups sets EntitlementGroups field to given value.

### HasEntitlementGroups

`func (o *FPGAItem) HasEntitlementGroups() bool`

HasEntitlementGroups returns a boolean if a field has been set.

### GetExtraMetadata

`func (o *FPGAItem) GetExtraMetadata() map[string]string`

GetExtraMetadata returns the ExtraMetadata field if non-nil, zero value otherwise.

### GetExtraMetadataOk

`func (o *FPGAItem) GetExtraMetadataOk() (*map[string]string, bool)`

GetExtraMetadataOk returns a tuple with the ExtraMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraMetadata

`func (o *FPGAItem) SetExtraMetadata(v map[string]string)`

SetExtraMetadata sets ExtraMetadata field to given value.

### HasExtraMetadata

`func (o *FPGAItem) HasExtraMetadata() bool`

HasExtraMetadata returns a boolean if a field has been set.

### GetName

`func (o *FPGAItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FPGAItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FPGAItem) SetName(v string)`

SetName sets Name field to given value.


### GetReadUARTConnections

`func (o *FPGAItem) GetReadUARTConnections() []FPGAUARTConnection`

GetReadUARTConnections returns the ReadUARTConnections field if non-nil, zero value otherwise.

### GetReadUARTConnectionsOk

`func (o *FPGAItem) GetReadUARTConnectionsOk() (*[]FPGAUARTConnection, bool)`

GetReadUARTConnectionsOk returns a tuple with the ReadUARTConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadUARTConnections

`func (o *FPGAItem) SetReadUARTConnections(v []FPGAUARTConnection)`

SetReadUARTConnections sets ReadUARTConnections field to given value.

### HasReadUARTConnections

`func (o *FPGAItem) HasReadUARTConnections() bool`

HasReadUARTConnections returns a boolean if a field has been set.

### GetReady

`func (o *FPGAItem) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *FPGAItem) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *FPGAItem) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetRepositories

`func (o *FPGAItem) GetRepositories() FPGAPayloadRepositories`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *FPGAItem) GetRepositoriesOk() (*FPGAPayloadRepositories, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *FPGAItem) SetRepositories(v FPGAPayloadRepositories)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *FPGAItem) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetSupportedTargets

`func (o *FPGAItem) GetSupportedTargets() []FPGATarget`

GetSupportedTargets returns the SupportedTargets field if non-nil, zero value otherwise.

### GetSupportedTargetsOk

`func (o *FPGAItem) GetSupportedTargetsOk() (*[]FPGATarget, bool)`

GetSupportedTargetsOk returns a tuple with the SupportedTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTargets

`func (o *FPGAItem) SetSupportedTargets(v []FPGATarget)`

SetSupportedTargets sets SupportedTargets field to given value.


### GetTitle

`func (o *FPGAItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FPGAItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FPGAItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetWorkerAuthToken

`func (o *FPGAItem) GetWorkerAuthToken() string`

GetWorkerAuthToken returns the WorkerAuthToken field if non-nil, zero value otherwise.

### GetWorkerAuthTokenOk

`func (o *FPGAItem) GetWorkerAuthTokenOk() (*string, bool)`

GetWorkerAuthTokenOk returns a tuple with the WorkerAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerAuthToken

`func (o *FPGAItem) SetWorkerAuthToken(v string)`

SetWorkerAuthToken sets WorkerAuthToken field to given value.

### HasWorkerAuthToken

`func (o *FPGAItem) HasWorkerAuthToken() bool`

HasWorkerAuthToken returns a boolean if a field has been set.

### GetWriteUARTConnection

`func (o *FPGAItem) GetWriteUARTConnection() FPGAUARTConnection`

GetWriteUARTConnection returns the WriteUARTConnection field if non-nil, zero value otherwise.

### GetWriteUARTConnectionOk

`func (o *FPGAItem) GetWriteUARTConnectionOk() (*FPGAUARTConnection, bool)`

GetWriteUARTConnectionOk returns a tuple with the WriteUARTConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteUARTConnection

`func (o *FPGAItem) SetWriteUARTConnection(v FPGAUARTConnection)`

SetWriteUARTConnection sets WriteUARTConnection field to given value.

### HasWriteUARTConnection

`func (o *FPGAItem) HasWriteUARTConnection() bool`

HasWriteUARTConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


