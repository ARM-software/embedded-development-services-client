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
func (o *WorkspaceCollection) FetchType() string {
	return "Workspace collection page"
}

// FetchLinks returns the resource links if present
func (o *WorkspaceCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *WorkspaceCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *WorkspaceCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *WorkspaceCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *WorkspaceCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewWorkspaceIterator(embedded.GetItem())
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

func (o *WorkspaceCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewWorkspacesCollection returns a page.
func NewWorkspacesCollection() IStaticPage {
	return NewWorkspaceCollectionWithDefaults()
}
