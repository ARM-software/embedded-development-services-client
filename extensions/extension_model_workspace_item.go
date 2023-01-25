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
func (o *WorkspaceItem) FetchType() string {
	return "Workspace"
}

// FetchLinks returns the resource links if present
func (o *WorkspaceItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *WorkspaceItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *WorkspaceItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewWorkspaceModel returns a model.
func NewWorkspaceModel() IModel {
	return NewWorkspaceItemWithDefaults()
}

// WorkspaceIterator defines an iterator over a workspace collection.
type WorkspaceIterator struct {
	elements     []WorkspaceItem
	currentIndex int
}

func (m *WorkspaceIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *WorkspaceIterator) GetNext() (item interface{}, err error) {
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

func NewWorkspaceIterator(elements []WorkspaceItem) (IIterator, error) {
	return &WorkspaceIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}
