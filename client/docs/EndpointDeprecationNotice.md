<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# EndpointDeprecationNotice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableEndpointDeprecationNoticeLinks**](EndpointDeprecationNoticeLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Deprecation** | [**DeprecationInfo**](DeprecationInfo.md) |  | 
**Endpoint** | [**HalLinkData**](HalLinkData.md) |  | 
**Replacement** | Pointer to [**HalLinkData**](HalLinkData.md) |  | [optional] 

## Methods

### NewEndpointDeprecationNotice

`func NewEndpointDeprecationNotice(links NullableEndpointDeprecationNoticeLinks, metadata NullableCommonMetadata, deprecation DeprecationInfo, endpoint HalLinkData, ) *EndpointDeprecationNotice`

NewEndpointDeprecationNotice instantiates a new EndpointDeprecationNotice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointDeprecationNoticeWithDefaults

`func NewEndpointDeprecationNoticeWithDefaults() *EndpointDeprecationNotice`

NewEndpointDeprecationNoticeWithDefaults instantiates a new EndpointDeprecationNotice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EndpointDeprecationNotice) GetLinks() EndpointDeprecationNoticeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EndpointDeprecationNotice) GetLinksOk() (*EndpointDeprecationNoticeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EndpointDeprecationNotice) SetLinks(v EndpointDeprecationNoticeLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *EndpointDeprecationNotice) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *EndpointDeprecationNotice) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *EndpointDeprecationNotice) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EndpointDeprecationNotice) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EndpointDeprecationNotice) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *EndpointDeprecationNotice) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EndpointDeprecationNotice) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDeprecation

`func (o *EndpointDeprecationNotice) GetDeprecation() DeprecationInfo`

GetDeprecation returns the Deprecation field if non-nil, zero value otherwise.

### GetDeprecationOk

`func (o *EndpointDeprecationNotice) GetDeprecationOk() (*DeprecationInfo, bool)`

GetDeprecationOk returns a tuple with the Deprecation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecation

`func (o *EndpointDeprecationNotice) SetDeprecation(v DeprecationInfo)`

SetDeprecation sets Deprecation field to given value.


### GetEndpoint

`func (o *EndpointDeprecationNotice) GetEndpoint() HalLinkData`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *EndpointDeprecationNotice) GetEndpointOk() (*HalLinkData, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *EndpointDeprecationNotice) SetEndpoint(v HalLinkData)`

SetEndpoint sets Endpoint field to given value.


### GetReplacement

`func (o *EndpointDeprecationNotice) GetReplacement() HalLinkData`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *EndpointDeprecationNotice) GetReplacementOk() (*HalLinkData, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *EndpointDeprecationNotice) SetReplacement(v HalLinkData)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *EndpointDeprecationNotice) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


