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
func (o *MessageObject) FetchType() string {
	return "Message"
}

// MessageIterator defines an iterator over a message collection.
type MessageIterator struct {
	elements     []MessageObject
	currentIndex int
}

func (m *MessageIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *MessageIterator) GetNext() (item *interface{}, err error) {
	if m.currentIndex < 0 {
		err = errors.New("incorrect element index")
		return
	}
	if !m.HasNext() {
		err = errors.New("no more items")
		return
	}
	element := interface{}(m.elements[m.currentIndex])
	item = &element
	m.currentIndex++
	return
}

func NewMessagesIterator(elements []MessageObject) (IIterator, error) {
	return &MessageIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}
