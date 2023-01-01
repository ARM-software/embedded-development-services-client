/*
 * Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

// Package client defines an HTTP client for communicating with the web services.
// It includes the definition of request/response types as well as provides helpers for calling specific helpers.
package client

import (
	"errors"
	"fmt"
)

// *************************************************************************************
// NOTE: this file is not generated.
// It extends the generated models.
// *************************************************************************************

// FetchType returns the resource type
func (o *BuildMessageItem) FetchType() string {
	return "Build Messages"
}

// FetchLinks returns the resource links if present
func (o *BuildMessageItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *BuildMessageItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *BuildMessageItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *BuildMessageItem) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *BuildMessageItem) HasFuture() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasFuture()
	}
	return false
}

func (o *BuildMessageItem) GetItemIterator() (IIterator, error) {
	return NewMessagesIterator(o.GetMessages())
}

func (o *BuildMessageItem) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewBuildMessageStream returns a message stream.
func NewBuildMessageStream() IMessageStream {
	return NewBuildMessageItemWithDefaults()
}
