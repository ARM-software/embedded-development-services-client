/*
 * Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

/*
Solar API

This API provides a RESTful interface to all the Solar services e.g. looking for boards, building projects, etc. - This API uses Hypermedia as the Engine of Application State (HATEOAS) to drive the discovery and provide   affordances. - Discovery is possible by following links from the well known root resource. While this specification lists   all supported endpoints, it is only recommended that these are hard coded into a client if code generation is   being used. Otherwise, it is recommended that the discovery mechanisms present in the resources (affordances)   are used exclusively. - Affordances are links which indicate whether an action is currently possible, this is significantly different from   whether the service supports an action in general. This specification defines what actions could be possible,   but only by checking the affordances returned by the API in the returned resources, can a client determine whether   this action is currently possible or available for the current user. For example:   - An operation to modify a resource could be defined in this specification, but the user may lack the appropriate     privileges. In that situation, the affordance link would not be present in the resource when read. Therefore,     the client can infer that it is not possible to edit this resource and present appropriate information to the     user.   - An operation to delete a resource could be defined and be possible in some circumstances. The specification     describes that the delete is supported and how to use it, but the affordance describes whether it is currently     possible. The logic in the API may dictate that if the resource was in use (perhaps it is a running job or used     by another resource), then it will not be possible to delete that resource as it would result in a conflicted     state. - It is strongly encouraged that affordances are used by all clients, even those using code generation. This has the   ability to both improve robustness and the user experience by decoupling the client and server. For example, if for   some reason the criteria for deleting a resource changes, the logic is only implemented in the server and there is   no need to update the logic in the client as it is driven by the affordances. - The format used for the resources is the Hypertext Application Language (HAL), which includes the definition   of links and embedded resources. 

API version: 1.0.0
Contact: support@arm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// BuildToolTypes The type of build system that will be used to build the project.  * `CMSIS` - Default CMSIS projects  * `CMSIS_CSOLUTION` - CMSIS <a href=\"https://github.com/Open-CMSIS-Pack/devtools\">csolution</a> for <a href=\"https://github.com/Open-CMSIS-Pack/devtools/blob/main/tools/projmgr/docs/Manual/Overview.md\">csolution projects</a>  * `CMSIS_CBUILD` - CMSIS <a href=\"https://github.com/Open-CMSIS-Pack/cbuild\">cbuild</a> for cprj projects  * `CMAKE` -  Default <a href=\"https://cmake.org/\">CMake</a>  * `CMAKE_NINJA` -  <a href=\"https://cmake.org/\">CMake</a> using the generator for <a href=\"https://cmake.org/cmake/help/latest/generator/Ninja.html\">Ninja</a> and building with <a href=\"https://ninja-build.org/\">Ninja</a>  * `CMAKE_MAKE` -  <a href=\"https://cmake.org/\">CMake</a> using the generator for <a href=\"https://cmake.org/cmake/help/latest/generator/Unix%20Makefiles.html\">Unix Make</a> and building with <a href=\"https://www.gnu.org/software/make/\">GNU Make</a> 
type BuildToolTypes string

// List of BuildToolTypes
const (
	CMAKE BuildToolTypes = "CMAKE"
	CMAKE_MAKE BuildToolTypes = "CMAKE_MAKE"
	CMAKE_NINJA BuildToolTypes = "CMAKE_NINJA"
	CMSIS BuildToolTypes = "CMSIS"
	CMSIS_CBUILD BuildToolTypes = "CMSIS_CBUILD"
	CMSIS_CSOLUTION BuildToolTypes = "CMSIS_CSOLUTION"
)

// All allowed values of BuildToolTypes enum
var AllowedBuildToolTypesEnumValues = []BuildToolTypes{
	"CMAKE",
	"CMAKE_MAKE",
	"CMAKE_NINJA",
	"CMSIS",
	"CMSIS_CBUILD",
	"CMSIS_CSOLUTION",
}

func (v *BuildToolTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BuildToolTypes(value)
	for _, existing := range AllowedBuildToolTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BuildToolTypes", value)
}

// NewBuildToolTypesFromValue returns a pointer to a valid BuildToolTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBuildToolTypesFromValue(v string) (*BuildToolTypes, error) {
	ev := BuildToolTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BuildToolTypes: valid values are %v", v, AllowedBuildToolTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BuildToolTypes) IsValid() bool {
	for _, existing := range AllowedBuildToolTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BuildToolTypes value
func (v BuildToolTypes) Ptr() *BuildToolTypes {
	return &v
}

type NullableBuildToolTypes struct {
	value *BuildToolTypes
	isSet bool
}

func (v NullableBuildToolTypes) Get() *BuildToolTypes {
	return v.value
}

func (v *NullableBuildToolTypes) Set(val *BuildToolTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildToolTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildToolTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildToolTypes(val *BuildToolTypes) *NullableBuildToolTypes {
	return &NullableBuildToolTypes{value: val, isSet: true}
}

func (v NullableBuildToolTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildToolTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

