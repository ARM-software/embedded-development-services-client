/*
 * Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package client

// ============================================================================================
// This extends BuildJobItem definitions
// ============================================================================================
// HasMessages specifies whether messages are available for this job.
func (o *BuildJobItem) HasMessages() (available bool) {
	if o == nil {
		return
	}
	links := o.GetLinks()
	_, available = links.GetDetailsOk()
	return
}

// HasArtefacts specifies whether artefacts are available for this job.
func (o *BuildJobItem) HasArtefacts() (available bool) {
	if o == nil {
		return
	}
	links := o.GetLinks()
	_, available = links.GetArtefactsOk()
	return
}

// ============================================================================================
// This extends FPGAJobItem definitions
// ============================================================================================
// HasMessages specifies whether messages are available for this job.
func (o *FPGAJobItem) HasMessages() (available bool) {
	if o == nil {
		return
	}
	links := o.GetLinks()
	_, available = links.GetDetailsOk()
	return
}

// HasArtefacts specifies whether artefacts are available for this job.
func (o *FPGAJobItem) HasArtefacts() (available bool) {
	if o == nil {
		return
	}
	links := o.GetLinks()
	_, available = links.GetArtefactsOk()
	return
}

// ============================================================================================
// This extends GenericWorkJobItem definitions
// ============================================================================================
// HasMessages specifies whether messages are available for this job.
func (o *GenericWorkJobItem) HasMessages() (available bool) {
	if o == nil {
		return
	}
	links := o.GetLinks()
	_, available = links.GetDetailsOk()
	return
}

// HasArtefacts specifies whether artefacts are available for this job.
func (o *GenericWorkJobItem) HasArtefacts() (available bool) {
	if o == nil {
		return
	}
	links := o.GetLinks()
	_, available = links.GetArtefactsOk()
	return
}

// ============================================================================================
// This extends IntellisenseJobItem definitions
// ============================================================================================
// HasMessages specifies whether messages are available for this job.
func (o *IntellisenseJobItem) HasMessages() (available bool) {
	if o == nil {
		return
	}
	links := o.GetLinks()
	_, available = links.GetDetailsOk()
	return
}

// HasArtefacts specifies whether artefacts are available for this job.
func (o *IntellisenseJobItem) HasArtefacts() (available bool) {
	if o == nil {
		return
	}
	links := o.GetLinks()
	_, available = links.GetArtefactsOk()
	return
}

// ============================================================================================
// This extends VhtRunJobItem definitions
// ============================================================================================
// HasMessages specifies whether messages are available for this job.
func (o *VhtRunJobItem) HasMessages() (available bool) {
	if o == nil {
		return
	}
	links := o.GetLinks()
	_, available = links.GetDetailsOk()
	return
}

// HasArtefacts specifies whether artefacts are available for this job.
func (o *VhtRunJobItem) HasArtefacts() (available bool) {
	if o == nil {
		return
	}
	links := o.GetLinks()
	_, available = links.GetArtefactsOk()
	return
}
