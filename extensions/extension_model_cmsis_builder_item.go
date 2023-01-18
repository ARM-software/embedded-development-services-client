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
func (o *CmsisBuilderItem) FetchType() string {
	return "Builder"
}

// FetchLinks returns the resource links if present
func (o *CmsisBuilderItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *CmsisBuilderItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *CmsisBuilderItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewBuilderModel returns a model.
func NewBuilderModel() IModel {
	return NewCmsisBuilderItemWithDefaults()
}

// BuilderIterator defines an iterator over a message collection.
type BuilderIterator struct {
	elements     []CmsisBuilderItem
	currentIndex int
}

func (m *BuilderIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *BuilderIterator) GetNext() (item interface{}, err error) {
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

func NewBuildersIterator(elements []CmsisBuilderItem) (IIterator, error) {
	return &BuilderIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}
