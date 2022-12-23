package client

import (
	"errors"
)

//*************************************************************************************
// NOTE: this file is not generated.
// It extends the generated models.
//*************************************************************************************

// FetchType returns the resource type
func (o *BuildJobItem) FetchType() string {
	return "Build Job"
}

// FetchLinks returns the resource links if present
func (o *BuildJobItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *BuildJobItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *BuildJobItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewBuildJobModel returns a model.
func NewBuildJobModel() IModel {
	return NewBuildJobItemWithDefaults()
}
