<!--
Copyright (C) 2020-2022 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# BoardItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | [**HalLinkData**](HalLinkData.md) |  | 
**Device** | Pointer to [**[]HalLinkData**](HalLinkData.md) | Links to Devices mounted on the Board. | [optional] 
**Documentation** | Pointer to [**[]HalLinkData**](HalLinkData.md) | Links to any documentation held by the Board. | [optional] 
**Download** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Guide** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Image** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 
**Vendor** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewBoardItemLinks

`func NewBoardItemLinks(collection HalLinkData, self HalLinkData, vendor HalLinkData, ) *BoardItemLinks`

NewBoardItemLinks instantiates a new BoardItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardItemLinksWithDefaults

`func NewBoardItemLinksWithDefaults() *BoardItemLinks`

NewBoardItemLinksWithDefaults instantiates a new BoardItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *BoardItemLinks) GetCollection() HalLinkData`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *BoardItemLinks) GetCollectionOk() (*HalLinkData, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *BoardItemLinks) SetCollection(v HalLinkData)`

SetCollection sets Collection field to given value.


### GetDevice

`func (o *BoardItemLinks) GetDevice() []HalLinkData`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BoardItemLinks) GetDeviceOk() (*[]HalLinkData, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BoardItemLinks) SetDevice(v []HalLinkData)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *BoardItemLinks) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDocumentation

`func (o *BoardItemLinks) GetDocumentation() []HalLinkData`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *BoardItemLinks) GetDocumentationOk() (*[]HalLinkData, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *BoardItemLinks) SetDocumentation(v []HalLinkData)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *BoardItemLinks) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetDownload

`func (o *BoardItemLinks) GetDownload() HalLinkData`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *BoardItemLinks) GetDownloadOk() (*HalLinkData, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *BoardItemLinks) SetDownload(v HalLinkData)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *BoardItemLinks) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetGuide

`func (o *BoardItemLinks) GetGuide() HalLinkData`

GetGuide returns the Guide field if non-nil, zero value otherwise.

### GetGuideOk

`func (o *BoardItemLinks) GetGuideOk() (*HalLinkData, bool)`

GetGuideOk returns a tuple with the Guide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuide

`func (o *BoardItemLinks) SetGuide(v HalLinkData)`

SetGuide sets Guide field to given value.

### HasGuide

`func (o *BoardItemLinks) HasGuide() bool`

HasGuide returns a boolean if a field has been set.

### GetImage

`func (o *BoardItemLinks) GetImage() HalLinkData`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BoardItemLinks) GetImageOk() (*HalLinkData, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BoardItemLinks) SetImage(v HalLinkData)`

SetImage sets Image field to given value.

### HasImage

`func (o *BoardItemLinks) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetSelf

`func (o *BoardItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *BoardItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *BoardItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.


### GetVendor

`func (o *BoardItemLinks) GetVendor() HalLinkData`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *BoardItemLinks) GetVendorOk() (*HalLinkData, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *BoardItemLinks) SetVendor(v HalLinkData)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


