<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAPayloadItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableFPGAPayloadItemLinks**](FPGAPayloadItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Name** | **string** | Unique ID of this FPGA Payload. | [readonly] 
**SupportConnection** | **bool** | True when the payload supports direct connection. | 
**Title** | **string** | Human readable name of the FPGA Payload. | 
**Workspace** | **string** | The ID of the workspace created for the payload to be uploaded to. This value should be provided to the upload session creation. | [readonly] 

## Methods

### NewFPGAPayloadItem

`func NewFPGAPayloadItem(links NullableFPGAPayloadItemLinks, metadata NullableCommonMetadata, name string, supportConnection bool, title string, workspace string, ) *FPGAPayloadItem`

NewFPGAPayloadItem instantiates a new FPGAPayloadItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAPayloadItemWithDefaults

`func NewFPGAPayloadItemWithDefaults() *FPGAPayloadItem`

NewFPGAPayloadItemWithDefaults instantiates a new FPGAPayloadItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FPGAPayloadItem) GetLinks() FPGAPayloadItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FPGAPayloadItem) GetLinksOk() (*FPGAPayloadItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FPGAPayloadItem) SetLinks(v FPGAPayloadItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *FPGAPayloadItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *FPGAPayloadItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *FPGAPayloadItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FPGAPayloadItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FPGAPayloadItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *FPGAPayloadItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *FPGAPayloadItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *FPGAPayloadItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FPGAPayloadItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FPGAPayloadItem) SetName(v string)`

SetName sets Name field to given value.


### GetSupportConnection

`func (o *FPGAPayloadItem) GetSupportConnection() bool`

GetSupportConnection returns the SupportConnection field if non-nil, zero value otherwise.

### GetSupportConnectionOk

`func (o *FPGAPayloadItem) GetSupportConnectionOk() (*bool, bool)`

GetSupportConnectionOk returns a tuple with the SupportConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportConnection

`func (o *FPGAPayloadItem) SetSupportConnection(v bool)`

SetSupportConnection sets SupportConnection field to given value.


### GetTitle

`func (o *FPGAPayloadItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FPGAPayloadItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FPGAPayloadItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetWorkspace

`func (o *FPGAPayloadItem) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *FPGAPayloadItem) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *FPGAPayloadItem) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


