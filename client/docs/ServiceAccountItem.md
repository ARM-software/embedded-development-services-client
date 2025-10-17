<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# ServiceAccountItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableServiceAccountItemLinks**](ServiceAccountItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**CreatedBy** | **string** | ID of the user who created this service account. | [readonly] 
**Name** | **string** | Unique ID of the service account. | [readonly] 
**Title** | Pointer to **string** | Human readable name of the service account. | [optional] 

## Methods

### NewServiceAccountItem

`func NewServiceAccountItem(links NullableServiceAccountItemLinks, metadata NullableCommonMetadata, createdBy string, name string, ) *ServiceAccountItem`

NewServiceAccountItem instantiates a new ServiceAccountItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountItemWithDefaults

`func NewServiceAccountItemWithDefaults() *ServiceAccountItem`

NewServiceAccountItemWithDefaults instantiates a new ServiceAccountItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ServiceAccountItem) GetLinks() ServiceAccountItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceAccountItem) GetLinksOk() (*ServiceAccountItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceAccountItem) SetLinks(v ServiceAccountItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *ServiceAccountItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ServiceAccountItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *ServiceAccountItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceAccountItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceAccountItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *ServiceAccountItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ServiceAccountItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedBy

`func (o *ServiceAccountItem) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServiceAccountItem) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServiceAccountItem) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetName

`func (o *ServiceAccountItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountItem) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ServiceAccountItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceAccountItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceAccountItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ServiceAccountItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


