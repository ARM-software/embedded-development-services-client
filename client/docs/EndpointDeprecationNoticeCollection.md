<!--
Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# EndpointDeprecationNoticeCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**EmbeddedEndpointDeprecationNoticeItems**](EmbeddedEndpointDeprecationNoticeItems.md) |  | [optional] 
**Links** | [**NullableHalSimpleCollectionLinks**](HalSimpleCollectionLinks.md) |  | 
**Metadata** | [**NullablePagingMetadata**](PagingMetadata.md) |  | 
**Name** | **string** | ID of the Collection. | [readonly] 
**Title** | **string** | Human readable title of the collection. | [readonly] 

## Methods

### NewEndpointDeprecationNoticeCollection

`func NewEndpointDeprecationNoticeCollection(links NullableHalSimpleCollectionLinks, metadata NullablePagingMetadata, name string, title string, ) *EndpointDeprecationNoticeCollection`

NewEndpointDeprecationNoticeCollection instantiates a new EndpointDeprecationNoticeCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointDeprecationNoticeCollectionWithDefaults

`func NewEndpointDeprecationNoticeCollectionWithDefaults() *EndpointDeprecationNoticeCollection`

NewEndpointDeprecationNoticeCollectionWithDefaults instantiates a new EndpointDeprecationNoticeCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *EndpointDeprecationNoticeCollection) GetEmbedded() EmbeddedEndpointDeprecationNoticeItems`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *EndpointDeprecationNoticeCollection) GetEmbeddedOk() (*EmbeddedEndpointDeprecationNoticeItems, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *EndpointDeprecationNoticeCollection) SetEmbedded(v EmbeddedEndpointDeprecationNoticeItems)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *EndpointDeprecationNoticeCollection) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *EndpointDeprecationNoticeCollection) GetLinks() HalSimpleCollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EndpointDeprecationNoticeCollection) GetLinksOk() (*HalSimpleCollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EndpointDeprecationNoticeCollection) SetLinks(v HalSimpleCollectionLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *EndpointDeprecationNoticeCollection) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *EndpointDeprecationNoticeCollection) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *EndpointDeprecationNoticeCollection) GetMetadata() PagingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EndpointDeprecationNoticeCollection) GetMetadataOk() (*PagingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EndpointDeprecationNoticeCollection) SetMetadata(v PagingMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *EndpointDeprecationNoticeCollection) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EndpointDeprecationNoticeCollection) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *EndpointDeprecationNoticeCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointDeprecationNoticeCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointDeprecationNoticeCollection) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *EndpointDeprecationNoticeCollection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EndpointDeprecationNoticeCollection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EndpointDeprecationNoticeCollection) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


