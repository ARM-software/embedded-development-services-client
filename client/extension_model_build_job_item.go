/*
 * Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

// Package client defines an HTTP client for communicating with the web services.
// It includes the definition of request/response types as well as provides helpers for calling specific helpers.
package client

import (
	"errors"
)

// *************************************************************************************
// NOTE: this file is not generated.
// It extends the generated models.
// *************************************************************************************

// FetchType returns the resource type
func (o *BuildJobItem) FetchType() string {
	return "Build Job"
}

// FetchLinks returns the resource links if present
func (o *BuildJobItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *BuildJobItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *BuildJobItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewBuildJobModel returns a model.
func NewBuildJobModel() IModel {
	return NewBuildJobItemWithDefaults()
}
