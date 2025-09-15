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
**Owner** | [**FPGAPayloadOwner**](FPGAPayloadOwner.md) |  | 
**Status** | **string** | Status of the payload. A payload is only ready to be used by an FPGA once it has been processed. | [readonly] 
**SupportConnection** | **bool** | True when the payload supports direct connection. | 
**Title** | **string** | Human readable name of the FPGA Payload. | 
**UploadJob** | Pointer to **NullableString** | The unique ID of the upload job that processed this payload. This will be null if the payload has not been processed yet. | [optional] [readonly] 
**UploadLocation** | **string** | The upload location to upload the payload files from. This value will be returned from the upload session creation. | 

## Methods

### NewFPGAPayloadItem

`func NewFPGAPayloadItem(links NullableFPGAPayloadItemLinks, metadata NullableCommonMetadata, name string, owner FPGAPayloadOwner, status string, supportConnection bool, title string, uploadLocation string, ) *FPGAPayloadItem`

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


### GetOwner

`func (o *FPGAPayloadItem) GetOwner() FPGAPayloadOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FPGAPayloadItem) GetOwnerOk() (*FPGAPayloadOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FPGAPayloadItem) SetOwner(v FPGAPayloadOwner)`

SetOwner sets Owner field to given value.


### GetStatus

`func (o *FPGAPayloadItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FPGAPayloadItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FPGAPayloadItem) SetStatus(v string)`

SetStatus sets Status field to given value.


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


### GetUploadJob

`func (o *FPGAPayloadItem) GetUploadJob() string`

GetUploadJob returns the UploadJob field if non-nil, zero value otherwise.

### GetUploadJobOk

`func (o *FPGAPayloadItem) GetUploadJobOk() (*string, bool)`

GetUploadJobOk returns a tuple with the UploadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadJob

`func (o *FPGAPayloadItem) SetUploadJob(v string)`

SetUploadJob sets UploadJob field to given value.

### HasUploadJob

`func (o *FPGAPayloadItem) HasUploadJob() bool`

HasUploadJob returns a boolean if a field has been set.

### SetUploadJobNil

`func (o *FPGAPayloadItem) SetUploadJobNil(b bool)`

 SetUploadJobNil sets the value for UploadJob to be an explicit nil

### UnsetUploadJob
`func (o *FPGAPayloadItem) UnsetUploadJob()`

UnsetUploadJob ensures that no value is present for UploadJob, not even an explicit nil
### GetUploadLocation

`func (o *FPGAPayloadItem) GetUploadLocation() string`

GetUploadLocation returns the UploadLocation field if non-nil, zero value otherwise.

### GetUploadLocationOk

`func (o *FPGAPayloadItem) GetUploadLocationOk() (*string, bool)`

GetUploadLocationOk returns a tuple with the UploadLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadLocation

`func (o *FPGAPayloadItem) SetUploadLocation(v string)`

SetUploadLocation sets UploadLocation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


