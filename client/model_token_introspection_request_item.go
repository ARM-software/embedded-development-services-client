/*
 * Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

/*
Solar API

This API provides a RESTful interface to all the Solar services e.g. looking for boards, building projects, etc. - This API uses Hypermedia as the Engine of Application State (HATEOAS) to drive the discovery and provide   affordances. - Discovery is possible by following links from the well known root resource. While this specification lists   all supported endpoints, it is only recommended that these are hard coded into a client if code generation is   being used. Otherwise, it is recommended that the discovery mechanisms present in the resources (affordances)   are used exclusively. - Affordances are links which indicate whether an action is currently possible, this is significantly different from   whether the service supports an action in general. This specification defines what actions could be possible,   but only by checking the affordances returned by the API in the returned resources, can a client determine whether   this action is currently possible or available for the current user. For example:   - An operation to modify a resource could be defined in this specification, but the user may lack the appropriate     privileges. In that situation, the affordance link would not be present in the resource when read. Therefore,     the client can infer that it is not possible to edit this resource and present appropriate information to the     user.   - An operation to delete a resource could be defined and be possible in some circumstances. The specification     describes that the delete is supported and how to use it, but the affordance describes whether it is currently     possible. The logic in the API may dictate that if the resource was in use (perhaps it is a running job or used     by another resource), then it will not be possible to delete that resource as it would result in a conflicted     state. - It is strongly encouraged that affordances are used by all clients, even those using code generation. This has the   ability to both improve robustness and the user experience by decoupling the client and server. For example, if for   some reason the criteria for deleting a resource changes, the logic is only implemented in the server and there is   no need to update the logic in the client as it is driven by the affordances. - The format used for the resources is the Hypertext Application Language (HAL), which includes the definition   of links and embedded resources. 

API version: 1.1.1
Contact: support@arm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TokenIntrospectionRequestItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenIntrospectionRequestItem{}

// TokenIntrospectionRequestItem The introspection request is based on https://www.rfc-editor.org/rfc/rfc7662#section-2.1
type TokenIntrospectionRequestItem struct {
	// The string value of the token.
	Token string `json:"token"`
	// A hint about the type of the token submitted for introspection.
	TokenTypeHint *string `json:"token_type_hint,omitempty"`
}

type _TokenIntrospectionRequestItem TokenIntrospectionRequestItem

// NewTokenIntrospectionRequestItem instantiates a new TokenIntrospectionRequestItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenIntrospectionRequestItem(token string) *TokenIntrospectionRequestItem {
	this := TokenIntrospectionRequestItem{}
	this.Token = token
	var tokenTypeHint string = "access_token"
	this.TokenTypeHint = &tokenTypeHint
	return &this
}

// NewTokenIntrospectionRequestItemWithDefaults instantiates a new TokenIntrospectionRequestItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenIntrospectionRequestItemWithDefaults() *TokenIntrospectionRequestItem {
	this := TokenIntrospectionRequestItem{}
	var tokenTypeHint string = "access_token"
	this.TokenTypeHint = &tokenTypeHint
	return &this
}

// GetToken returns the Token field value
func (o *TokenIntrospectionRequestItem) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *TokenIntrospectionRequestItem) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *TokenIntrospectionRequestItem) SetToken(v string) {
	o.Token = v
}

// GetTokenTypeHint returns the TokenTypeHint field value if set, zero value otherwise.
func (o *TokenIntrospectionRequestItem) GetTokenTypeHint() string {
	if o == nil || IsNil(o.TokenTypeHint) {
		var ret string
		return ret
	}
	return *o.TokenTypeHint
}

// GetTokenTypeHintOk returns a tuple with the TokenTypeHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIntrospectionRequestItem) GetTokenTypeHintOk() (*string, bool) {
	if o == nil || IsNil(o.TokenTypeHint) {
		return nil, false
	}
	return o.TokenTypeHint, true
}

// HasTokenTypeHint returns a boolean if a field has been set.
func (o *TokenIntrospectionRequestItem) HasTokenTypeHint() bool {
	if o != nil && !IsNil(o.TokenTypeHint) {
		return true
	}

	return false
}

// SetTokenTypeHint gets a reference to the given string and assigns it to the TokenTypeHint field.
func (o *TokenIntrospectionRequestItem) SetTokenTypeHint(v string) {
	o.TokenTypeHint = &v
}

func (o TokenIntrospectionRequestItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenIntrospectionRequestItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	if !IsNil(o.TokenTypeHint) {
		toSerialize["token_type_hint"] = o.TokenTypeHint
	}
	return toSerialize, nil
}

func (o *TokenIntrospectionRequestItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTokenIntrospectionRequestItem := _TokenIntrospectionRequestItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenIntrospectionRequestItem)

	if err != nil {
		return err
	}

	*o = TokenIntrospectionRequestItem(varTokenIntrospectionRequestItem)

	return err
}

type NullableTokenIntrospectionRequestItem struct {
	value *TokenIntrospectionRequestItem
	isSet bool
}

func (v NullableTokenIntrospectionRequestItem) Get() *TokenIntrospectionRequestItem {
	return v.value
}

func (v *NullableTokenIntrospectionRequestItem) Set(val *TokenIntrospectionRequestItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenIntrospectionRequestItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenIntrospectionRequestItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenIntrospectionRequestItem(val *TokenIntrospectionRequestItem) *NullableTokenIntrospectionRequestItem {
	return &NullableTokenIntrospectionRequestItem{value: val, isSet: true}
}

func (v NullableTokenIntrospectionRequestItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenIntrospectionRequestItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


