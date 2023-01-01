/*
 * Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

// Package client defines an HTTP client for communicating with the web services.
// It includes the definition of request/response types as well as provides helpers for calling specific helpers.
package client

// *************************************************************************************
// NOTE: this file is not generated.
// It defines generic models.
// *************************************************************************************

type IModel interface {
	// FetchType returns the resource type
	FetchType() string
	// FetchLinks returns the resource links if present
	FetchLinks() (links any, err error)
	// FetchName returns the resource name if present, or else an error
	FetchName() (string, error)
	// FetchTitle returns the resource title if present, or else an error
	FetchTitle() (string, error)
}

// The following definitions match what is in golang-utils for [pagination](https://github.com/ARM-software/golang-utils/blob/master/utils/collection/pagination/interfaces.go)

type IIterator interface {
	// HasNext returns whether there are more items available or not.
	HasNext() bool
	// GetNext returns the next item.
	GetNext() (*interface{}, error)
}

// IStaticPage defines a generic page for a collection.
type IStaticPage interface {
	// HasNext states whether more pages are accessible.
	HasNext() bool
	// GetItemIterator returns a new iterator over the page's items.
	GetItemIterator() (IIterator, error)
	// GetItemCount returns the number of items in this page
	GetItemCount() (int64, error)
}

// IMessageStream defines a page for a collection which does not have any known ending.
type IMessageStream interface {
	IStaticPage
	// HasFuture states whether there may be future items.
	HasFuture() bool
}
