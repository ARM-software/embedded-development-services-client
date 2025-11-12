<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# ServiceAccessTokenItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableAccessTokenItemLinks**](AccessTokenItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**CreatedBy** | **string** | ID of the user who created this access token. | [readonly] 
**LastUsed** | **time.Time** | When this token was last used to authenticate a request. | [readonly] 
**Name** | **string** | Unique ID of the service access token. | [readonly] 
**Secret** | Pointer to **string** | The access token. This field will only be returned by the service upon creation and the secret will not be re-retrievable. | [optional] [readonly] 
**SecretHint** | **string** | The four character hint of the access token secret. | [readonly] 
**Title** | **string** | Human readable name of the service access token. | 

## Methods

### NewServiceAccessTokenItem

`func NewServiceAccessTokenItem(links NullableAccessTokenItemLinks, metadata NullableCommonMetadata, createdBy string, lastUsed time.Time, name string, secretHint string, title string, ) *ServiceAccessTokenItem`

NewServiceAccessTokenItem instantiates a new ServiceAccessTokenItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccessTokenItemWithDefaults

`func NewServiceAccessTokenItemWithDefaults() *ServiceAccessTokenItem`

NewServiceAccessTokenItemWithDefaults instantiates a new ServiceAccessTokenItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ServiceAccessTokenItem) GetLinks() AccessTokenItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceAccessTokenItem) GetLinksOk() (*AccessTokenItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceAccessTokenItem) SetLinks(v AccessTokenItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *ServiceAccessTokenItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ServiceAccessTokenItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *ServiceAccessTokenItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceAccessTokenItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceAccessTokenItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *ServiceAccessTokenItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ServiceAccessTokenItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedBy

`func (o *ServiceAccessTokenItem) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServiceAccessTokenItem) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServiceAccessTokenItem) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastUsed

`func (o *ServiceAccessTokenItem) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *ServiceAccessTokenItem) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *ServiceAccessTokenItem) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.


### GetName

`func (o *ServiceAccessTokenItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccessTokenItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccessTokenItem) SetName(v string)`

SetName sets Name field to given value.


### GetSecret

`func (o *ServiceAccessTokenItem) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ServiceAccessTokenItem) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ServiceAccessTokenItem) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ServiceAccessTokenItem) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSecretHint

`func (o *ServiceAccessTokenItem) GetSecretHint() string`

GetSecretHint returns the SecretHint field if non-nil, zero value otherwise.

### GetSecretHintOk

`func (o *ServiceAccessTokenItem) GetSecretHintOk() (*string, bool)`

GetSecretHintOk returns a tuple with the SecretHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretHint

`func (o *ServiceAccessTokenItem) SetSecretHint(v string)`

SetSecretHint sets SecretHint field to given value.


### GetTitle

`func (o *ServiceAccessTokenItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceAccessTokenItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceAccessTokenItem) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


