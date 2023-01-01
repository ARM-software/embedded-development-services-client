<!--
Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# HalLinkData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deprecation** | Pointer to **string** | Indicates that the link is to be deprecated (i.e. removed) at a future date. Its value is a URL | [optional] 
**Href** | **string** | link URI or URI Template | 
**Hreflang** | Pointer to **string** | Indicates the language of the target resource (as defined by [RFC5988]). | [optional] 
**Name** | Pointer to **string** | key for selecting Link Objects which share the same relation type. | [optional] 
**Profile** | Pointer to **string** | URI that hints about the profile (as defined by [I-D.wilde-profile-link]) of the target resource. | [optional] 
**Templated** | Pointer to **bool** | States whether Link Object&#39;s \&quot;href\&quot; is a URI Template | [optional] [default to false]
**Title** | Pointer to **string** | Human-readable identifier labelling the link. | [optional] 
**Type** | Pointer to **string** | hint to indicate the media type | [optional] 

## Methods

### NewHalLinkData

`func NewHalLinkData(href string, ) *HalLinkData`

NewHalLinkData instantiates a new HalLinkData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHalLinkDataWithDefaults

`func NewHalLinkDataWithDefaults() *HalLinkData`

NewHalLinkDataWithDefaults instantiates a new HalLinkData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeprecation

`func (o *HalLinkData) GetDeprecation() string`

GetDeprecation returns the Deprecation field if non-nil, zero value otherwise.

### GetDeprecationOk

`func (o *HalLinkData) GetDeprecationOk() (*string, bool)`

GetDeprecationOk returns a tuple with the Deprecation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecation

`func (o *HalLinkData) SetDeprecation(v string)`

SetDeprecation sets Deprecation field to given value.

### HasDeprecation

`func (o *HalLinkData) HasDeprecation() bool`

HasDeprecation returns a boolean if a field has been set.

### GetHref

`func (o *HalLinkData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *HalLinkData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *HalLinkData) SetHref(v string)`

SetHref sets Href field to given value.


### GetHreflang

`func (o *HalLinkData) GetHreflang() string`

GetHreflang returns the Hreflang field if non-nil, zero value otherwise.

### GetHreflangOk

`func (o *HalLinkData) GetHreflangOk() (*string, bool)`

GetHreflangOk returns a tuple with the Hreflang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHreflang

`func (o *HalLinkData) SetHreflang(v string)`

SetHreflang sets Hreflang field to given value.

### HasHreflang

`func (o *HalLinkData) HasHreflang() bool`

HasHreflang returns a boolean if a field has been set.

### GetName

`func (o *HalLinkData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HalLinkData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HalLinkData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HalLinkData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfile

`func (o *HalLinkData) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *HalLinkData) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *HalLinkData) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *HalLinkData) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetTemplated

`func (o *HalLinkData) GetTemplated() bool`

GetTemplated returns the Templated field if non-nil, zero value otherwise.

### GetTemplatedOk

`func (o *HalLinkData) GetTemplatedOk() (*bool, bool)`

GetTemplatedOk returns a tuple with the Templated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplated

`func (o *HalLinkData) SetTemplated(v bool)`

SetTemplated sets Templated field to given value.

### HasTemplated

`func (o *HalLinkData) HasTemplated() bool`

HasTemplated returns a boolean if a field has been set.

### GetTitle

`func (o *HalLinkData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HalLinkData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HalLinkData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HalLinkData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *HalLinkData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HalLinkData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HalLinkData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HalLinkData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


