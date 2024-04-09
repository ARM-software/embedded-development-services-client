/*
 * Copyright (C) 2020-2024 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package client

import (
	"errors"
	"fmt"
)

// ============================================================================================
// This extends ArtefactManagerItem and ArtefactManagerCollection definitions
// ============================================================================================

// FetchType returns the resource type
func (o *ArtefactManagerItem) FetchType() string {
	return "ArtefactManagerItem"
}

// FetchLinks returns the resource links if present
func (o *ArtefactManagerItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *ArtefactManagerItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *ArtefactManagerItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewArtefactManagerModel returns a model.
func NewArtefactManagerModel() IModel {
	return NewArtefactManagerItemWithDefaults()
}

// ArtefactManagerIterator defines an iterator over a collection.
type ArtefactManagerIterator struct {
	elements     []ArtefactManagerItem
	currentIndex int
}

func (m *ArtefactManagerIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *ArtefactManagerIterator) GetNext() (item any, err error) {
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

func NewArtefactManagerIterator(elements []ArtefactManagerItem) (IIterator, error) {
	return &ArtefactManagerIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *ArtefactManagerCollection) FetchType() string {
	return "ArtefactManagerCollection page"
}

// FetchLinks returns the resource links if present
func (o *ArtefactManagerCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *ArtefactManagerCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *ArtefactManagerCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *ArtefactManagerCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *ArtefactManagerCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewArtefactManagerIterator(embedded.GetItem())
	}
	links, err := o.FetchLinks()
	if err != nil {
		return nil, err
	}
	l, ok := links.(HalCollectionLinks)
	if !ok {
		return nil, fmt.Errorf("wrong link type [%T]; expected [HalCollectionLinks]", links)
	}
	return NewHalLinkDataIterator(l.GetItem())
}

func (o *ArtefactManagerCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewArtefactManagerCollection returns a page.
func NewArtefactManagerCollectionCollection() IStaticPage {
	return NewArtefactManagerCollectionWithDefaults()
}
