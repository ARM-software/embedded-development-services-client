<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# AccountSummaryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableAccountSummaryLinks**](AccountSummaryLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**AccountType** | [**AccountType**](AccountType.md) |  | 
**Name** | **string** | The systems unique identifier for the account. | [readonly] 
**Title** | Pointer to **string** | Human readable name for the account. | [optional] 

## Methods

### NewAccountSummaryItem

`func NewAccountSummaryItem(links NullableAccountSummaryLinks, metadata NullableCommonMetadata, accountType AccountType, name string, ) *AccountSummaryItem`

NewAccountSummaryItem instantiates a new AccountSummaryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSummaryItemWithDefaults

`func NewAccountSummaryItemWithDefaults() *AccountSummaryItem`

NewAccountSummaryItemWithDefaults instantiates a new AccountSummaryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AccountSummaryItem) GetLinks() AccountSummaryLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AccountSummaryItem) GetLinksOk() (*AccountSummaryLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AccountSummaryItem) SetLinks(v AccountSummaryLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *AccountSummaryItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AccountSummaryItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *AccountSummaryItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountSummaryItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountSummaryItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *AccountSummaryItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AccountSummaryItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAccountType

`func (o *AccountSummaryItem) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountSummaryItem) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountSummaryItem) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetName

`func (o *AccountSummaryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountSummaryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountSummaryItem) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AccountSummaryItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AccountSummaryItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AccountSummaryItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AccountSummaryItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


