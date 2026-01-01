<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# TokenIntrospectionRequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | The string value of the token. | 
**TokenTypeHint** | Pointer to **string** | A hint about the type of the token submitted for introspection. | [optional] [default to "access_token"]

## Methods

### NewTokenIntrospectionRequestItem

`func NewTokenIntrospectionRequestItem(token string, ) *TokenIntrospectionRequestItem`

NewTokenIntrospectionRequestItem instantiates a new TokenIntrospectionRequestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenIntrospectionRequestItemWithDefaults

`func NewTokenIntrospectionRequestItemWithDefaults() *TokenIntrospectionRequestItem`

NewTokenIntrospectionRequestItemWithDefaults instantiates a new TokenIntrospectionRequestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *TokenIntrospectionRequestItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenIntrospectionRequestItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenIntrospectionRequestItem) SetToken(v string)`

SetToken sets Token field to given value.


### GetTokenTypeHint

`func (o *TokenIntrospectionRequestItem) GetTokenTypeHint() string`

GetTokenTypeHint returns the TokenTypeHint field if non-nil, zero value otherwise.

### GetTokenTypeHintOk

`func (o *TokenIntrospectionRequestItem) GetTokenTypeHintOk() (*string, bool)`

GetTokenTypeHintOk returns a tuple with the TokenTypeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTypeHint

`func (o *TokenIntrospectionRequestItem) SetTokenTypeHint(v string)`

SetTokenTypeHint sets TokenTypeHint field to given value.

### HasTokenTypeHint

`func (o *TokenIntrospectionRequestItem) HasTokenTypeHint() bool`

HasTokenTypeHint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


