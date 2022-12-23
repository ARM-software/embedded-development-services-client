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
