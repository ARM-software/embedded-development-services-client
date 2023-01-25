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
func (o *VhtItem) FetchType() string {
	return "Build Job"
}

// FetchLinks returns the resource links if present
func (o *VhtItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *VhtItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *VhtItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewVhtModel returns a model.
func NewVhtModel() IModel {
	return NewVhtItemWithDefaults()
}

// VhtIterator defines an iterator over a virtual hardware target collection.
type VhtIterator struct {
	elements     []VhtItem
	currentIndex int
}

func (m *VhtIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *VhtIterator) GetNext() (item interface{}, err error) {
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

func NewVhtIterator(elements []VhtItem) (IIterator, error) {
	return &VhtIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}
