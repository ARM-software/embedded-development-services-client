/*
 * Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
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
	"fmt"
)

// ToolchainTypes The type of toolchain that will be used to compile and link the project.  * `GCC` - GNU Tools for Arm Embedded Processors. Refer to <a href=\"https://developer.arm.com/tools-and-software/open-source-software/developer-tools/gnu-toolchain/gnu-rm\">Arm GCC</a>.  * `G++` - Code Sourcery GCC compiler for C and C++ (is now Mentor Graphics CodeBench).  * `AC5` - Arm Compiler for C and C++ Major Version 5. Refer to <a href=\"https://developer.arm.com/tools-and-software/embedded/arm-compiler/documentation/version-5\">Arm Compiler</a>. Due to incompatible command line syntax version 5 and 6 are listed as separate compilers.  * `AC6` - Arm Compiler for C and C++ Major Version 6. Refer to <a href=\"https://developer.arm.com/tools-and-software/embedded/arm-compiler/documentation\">Arm Compiler</a>. Due to incompatible command line syntax version 5 and 6 are listed as separate compilers.  * `IAR` - IAR compiler for C and C++.  * `Tasking` - TASKING compiler for C and C++.  * `GHS` - Green Hills Software compiler for C, C++, and EC++.
type ToolchainTypes string

// List of ToolchainTypes
const (
	AC5 ToolchainTypes = "AC5"
	AC6 ToolchainTypes = "AC6"
	G ToolchainTypes = "G++"
	GCC ToolchainTypes = "GCC"
	GHS ToolchainTypes = "GHS"
	IAR ToolchainTypes = "IAR"
	TASKING ToolchainTypes = "Tasking"
)

// All allowed values of ToolchainTypes enum
var AllowedToolchainTypesEnumValues = []ToolchainTypes{
	"AC5",
	"AC6",
	"G++",
	"GCC",
	"GHS",
	"IAR",
	"Tasking",
}

func (v *ToolchainTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ToolchainTypes(value)
	for _, existing := range AllowedToolchainTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ToolchainTypes", value)
}

// NewToolchainTypesFromValue returns a pointer to a valid ToolchainTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolchainTypesFromValue(v string) (*ToolchainTypes, error) {
	ev := ToolchainTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolchainTypes: valid values are %v", v, AllowedToolchainTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolchainTypes) IsValid() bool {
	for _, existing := range AllowedToolchainTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToolchainTypes value
func (v ToolchainTypes) Ptr() *ToolchainTypes {
	return &v
}

type NullableToolchainTypes struct {
	value *ToolchainTypes
	isSet bool
}

func (v NullableToolchainTypes) Get() *ToolchainTypes {
	return v.value
}

func (v *NullableToolchainTypes) Set(val *ToolchainTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableToolchainTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableToolchainTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolchainTypes(val *ToolchainTypes) *NullableToolchainTypes {
	return &NullableToolchainTypes{value: val, isSet: true}
}

func (v NullableToolchainTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolchainTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

