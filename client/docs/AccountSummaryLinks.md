<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# AccountSummaryLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Describedby** | [**HalLinkData**](HalLinkData.md) |  | 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewAccountSummaryLinks

`func NewAccountSummaryLinks(describedby HalLinkData, self HalLinkData, ) *AccountSummaryLinks`

NewAccountSummaryLinks instantiates a new AccountSummaryLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSummaryLinksWithDefaults

`func NewAccountSummaryLinksWithDefaults() *AccountSummaryLinks`

NewAccountSummaryLinksWithDefaults instantiates a new AccountSummaryLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescribedby

`func (o *AccountSummaryLinks) GetDescribedby() HalLinkData`

GetDescribedby returns the Describedby field if non-nil, zero value otherwise.

### GetDescribedbyOk

`func (o *AccountSummaryLinks) GetDescribedbyOk() (*HalLinkData, bool)`

GetDescribedbyOk returns a tuple with the Describedby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribedby

`func (o *AccountSummaryLinks) SetDescribedby(v HalLinkData)`

SetDescribedby sets Describedby field to given value.


### GetSelf

`func (o *AccountSummaryLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AccountSummaryLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AccountSummaryLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


