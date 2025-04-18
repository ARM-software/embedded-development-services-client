/*
 * Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
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
func (o *HalLinkData) FetchType() string {
	return "HAL link"
}

// HalLinkDataIterator defines an iterator over a message collection.
type HalLinkDataIterator struct {
	elements     []HalLinkData
	currentIndex int
}

func (m *HalLinkDataIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *HalLinkDataIterator) GetNext() (item interface{}, err error) {
	if m.currentIndex < 0 {
		err = errors.New("incorrect element index")
		return
	}
	if !m.HasNext() {
		err = errors.New("no more items")
		return
	}
	element := m.elements[m.currentIndex]
	item = &element
	m.currentIndex++
	return
}

func NewHalLinkDataIterator(elements []HalLinkData) (IIterator, error) {
	return &HalLinkDataIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// *************************************************************************************
// Definitions for making types compatible with Client Utils
// *************************************************************************************
// GetNext returns the Next field value if set, zero value otherwise.
func (o *HalCollectionLinks) GetNextP() *HalLinkData {
	if o == nil || IsNil(o.Next) {
		return nil
	}
	return o.Next
}
