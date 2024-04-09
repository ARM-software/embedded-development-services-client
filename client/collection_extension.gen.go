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
// This extends ArtefactManagerItem and ArtefactManagerCollection with pagination methods
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

// ============================================================================================
// This extends BuildJobItem and BuildJobCollection with pagination methods
// ============================================================================================

// FetchType returns the resource type
func (o *BuildJobItem) FetchType() string {
	return "BuildJobItem"
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

// BuildJobIterator defines an iterator over a collection.
type BuildJobIterator struct {
	elements     []BuildJobItem
	currentIndex int
}

func (m *BuildJobIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *BuildJobIterator) GetNext() (item any, err error) {
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

func NewBuildJobIterator(elements []BuildJobItem) (IIterator, error) {
	return &BuildJobIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *BuildJobCollection) FetchType() string {
	return "BuildJobCollection page"
}

// FetchLinks returns the resource links if present
func (o *BuildJobCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *BuildJobCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *BuildJobCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *BuildJobCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *BuildJobCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewBuildJobIterator(embedded.GetItem())
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

func (o *BuildJobCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewBuildJobCollection returns a page.
func NewBuildJobCollectionCollection() IStaticPage {
	return NewBuildJobCollectionWithDefaults()
}

// ============================================================================================
// This extends CmsisBuilderItem and CmsisBuilderCollection with pagination methods
// ============================================================================================

// FetchType returns the resource type
func (o *CmsisBuilderItem) FetchType() string {
	return "CmsisBuilderItem"
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

// NewCmsisBuilderModel returns a model.
func NewCmsisBuilderModel() IModel {
	return NewCmsisBuilderItemWithDefaults()
}

// CmsisBuilderIterator defines an iterator over a collection.
type CmsisBuilderIterator struct {
	elements     []CmsisBuilderItem
	currentIndex int
}

func (m *CmsisBuilderIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *CmsisBuilderIterator) GetNext() (item any, err error) {
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

func NewCmsisBuilderIterator(elements []CmsisBuilderItem) (IIterator, error) {
	return &CmsisBuilderIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *CmsisBuilderCollection) FetchType() string {
	return "CmsisBuilderCollection page"
}

// FetchLinks returns the resource links if present
func (o *CmsisBuilderCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *CmsisBuilderCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *CmsisBuilderCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *CmsisBuilderCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *CmsisBuilderCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewCmsisBuilderIterator(embedded.GetItem())
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

func (o *CmsisBuilderCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewCmsisBuilderCollection returns a page.
func NewCmsisBuilderCollectionCollection() IStaticPage {
	return NewCmsisBuilderCollectionWithDefaults()
}

// ============================================================================================
// This extends CmsisIntellisenseItem and CmsisIntellisenseCollection with pagination methods
// ============================================================================================

// FetchType returns the resource type
func (o *CmsisIntellisenseItem) FetchType() string {
	return "CmsisIntellisenseItem"
}

// FetchLinks returns the resource links if present
func (o *CmsisIntellisenseItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *CmsisIntellisenseItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *CmsisIntellisenseItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewCmsisIntellisenseModel returns a model.
func NewCmsisIntellisenseModel() IModel {
	return NewCmsisIntellisenseItemWithDefaults()
}

// CmsisIntellisenseIterator defines an iterator over a collection.
type CmsisIntellisenseIterator struct {
	elements     []CmsisIntellisenseItem
	currentIndex int
}

func (m *CmsisIntellisenseIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *CmsisIntellisenseIterator) GetNext() (item any, err error) {
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

func NewCmsisIntellisenseIterator(elements []CmsisIntellisenseItem) (IIterator, error) {
	return &CmsisIntellisenseIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *CmsisIntellisenseCollection) FetchType() string {
	return "CmsisIntellisenseCollection page"
}

// FetchLinks returns the resource links if present
func (o *CmsisIntellisenseCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *CmsisIntellisenseCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *CmsisIntellisenseCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *CmsisIntellisenseCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *CmsisIntellisenseCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewCmsisIntellisenseIterator(embedded.GetItem())
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

func (o *CmsisIntellisenseCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewCmsisIntellisenseCollection returns a page.
func NewCmsisIntellisenseCollectionCollection() IStaticPage {
	return NewCmsisIntellisenseCollectionWithDefaults()
}

// ============================================================================================
// This extends IntellisenseJobItem and IntellisenseJobCollection with pagination methods
// ============================================================================================

// FetchType returns the resource type
func (o *IntellisenseJobItem) FetchType() string {
	return "IntellisenseJobItem"
}

// FetchLinks returns the resource links if present
func (o *IntellisenseJobItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *IntellisenseJobItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *IntellisenseJobItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewIntellisenseJobModel returns a model.
func NewIntellisenseJobModel() IModel {
	return NewIntellisenseJobItemWithDefaults()
}

// IntellisenseJobIterator defines an iterator over a collection.
type IntellisenseJobIterator struct {
	elements     []IntellisenseJobItem
	currentIndex int
}

func (m *IntellisenseJobIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *IntellisenseJobIterator) GetNext() (item any, err error) {
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

func NewIntellisenseJobIterator(elements []IntellisenseJobItem) (IIterator, error) {
	return &IntellisenseJobIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *IntellisenseJobCollection) FetchType() string {
	return "IntellisenseJobCollection page"
}

// FetchLinks returns the resource links if present
func (o *IntellisenseJobCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *IntellisenseJobCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *IntellisenseJobCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *IntellisenseJobCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *IntellisenseJobCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewIntellisenseJobIterator(embedded.GetItem())
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

func (o *IntellisenseJobCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewIntellisenseJobCollection returns a page.
func NewIntellisenseJobCollectionCollection() IStaticPage {
	return NewIntellisenseJobCollectionWithDefaults()
}

// ============================================================================================
// This extends VhtInstanceItem and VhtInstanceCollection with pagination methods
// ============================================================================================

// FetchType returns the resource type
func (o *VhtInstanceItem) FetchType() string {
	return "VhtInstanceItem"
}

// FetchLinks returns the resource links if present
func (o *VhtInstanceItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *VhtInstanceItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *VhtInstanceItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewVhtInstanceModel returns a model.
func NewVhtInstanceModel() IModel {
	return NewVhtInstanceItemWithDefaults()
}

// VhtInstanceIterator defines an iterator over a collection.
type VhtInstanceIterator struct {
	elements     []VhtInstanceItem
	currentIndex int
}

func (m *VhtInstanceIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *VhtInstanceIterator) GetNext() (item any, err error) {
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

func NewVhtInstanceIterator(elements []VhtInstanceItem) (IIterator, error) {
	return &VhtInstanceIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *VhtInstanceCollection) FetchType() string {
	return "VhtInstanceCollection page"
}

// FetchLinks returns the resource links if present
func (o *VhtInstanceCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *VhtInstanceCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *VhtInstanceCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *VhtInstanceCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *VhtInstanceCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewVhtInstanceIterator(embedded.GetItem())
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

func (o *VhtInstanceCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewVhtInstanceCollection returns a page.
func NewVhtInstanceCollectionCollection() IStaticPage {
	return NewVhtInstanceCollectionWithDefaults()
}

// ============================================================================================
// This extends VhtItem and VhtCollection with pagination methods
// ============================================================================================

// FetchType returns the resource type
func (o *VhtItem) FetchType() string {
	return "VhtItem"
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

// VhtIterator defines an iterator over a collection.
type VhtIterator struct {
	elements     []VhtItem
	currentIndex int
}

func (m *VhtIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *VhtIterator) GetNext() (item any, err error) {
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

// FetchType returns the resource type
func (o *VhtCollection) FetchType() string {
	return "VhtCollection page"
}

// FetchLinks returns the resource links if present
func (o *VhtCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *VhtCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *VhtCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *VhtCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *VhtCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewVhtIterator(embedded.GetItem())
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

func (o *VhtCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewVhtCollection returns a page.
func NewVhtCollectionCollection() IStaticPage {
	return NewVhtCollectionWithDefaults()
}

// ============================================================================================
// This extends VhtRunJobItem and VhtRunJobCollection with pagination methods
// ============================================================================================

// FetchType returns the resource type
func (o *VhtRunJobItem) FetchType() string {
	return "VhtRunJobItem"
}

// FetchLinks returns the resource links if present
func (o *VhtRunJobItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *VhtRunJobItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *VhtRunJobItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewVhtRunJobModel returns a model.
func NewVhtRunJobModel() IModel {
	return NewVhtRunJobItemWithDefaults()
}

// VhtRunJobIterator defines an iterator over a collection.
type VhtRunJobIterator struct {
	elements     []VhtRunJobItem
	currentIndex int
}

func (m *VhtRunJobIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *VhtRunJobIterator) GetNext() (item any, err error) {
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

func NewVhtRunJobIterator(elements []VhtRunJobItem) (IIterator, error) {
	return &VhtRunJobIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *VhtRunJobCollection) FetchType() string {
	return "VhtRunJobCollection page"
}

// FetchLinks returns the resource links if present
func (o *VhtRunJobCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *VhtRunJobCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *VhtRunJobCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *VhtRunJobCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *VhtRunJobCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewVhtRunJobIterator(embedded.GetItem())
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

func (o *VhtRunJobCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewVhtRunJobCollection returns a page.
func NewVhtRunJobCollectionCollection() IStaticPage {
	return NewVhtRunJobCollectionWithDefaults()
}

// ============================================================================================
// This extends WorkspaceItem and WorkspaceCollection with pagination methods
// ============================================================================================

// FetchType returns the resource type
func (o *WorkspaceItem) FetchType() string {
	return "WorkspaceItem"
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

// WorkspaceIterator defines an iterator over a collection.
type WorkspaceIterator struct {
	elements     []WorkspaceItem
	currentIndex int
}

func (m *WorkspaceIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *WorkspaceIterator) GetNext() (item any, err error) {
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

// FetchType returns the resource type
func (o *WorkspaceCollection) FetchType() string {
	return "WorkspaceCollection page"
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

// NewWorkspaceCollection returns a page.
func NewWorkspaceCollectionCollection() IStaticPage {
	return NewWorkspaceCollectionWithDefaults()
}

// ============================================================================================
// This extends WorkspaceSourceItem and WorkspaceSourceCollection with pagination methods
// ============================================================================================

// FetchType returns the resource type
func (o *WorkspaceSourceItem) FetchType() string {
	return "WorkspaceSourceItem"
}

// FetchLinks returns the resource links if present
func (o *WorkspaceSourceItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *WorkspaceSourceItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *WorkspaceSourceItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewWorkspaceSourceModel returns a model.
func NewWorkspaceSourceModel() IModel {
	return NewWorkspaceSourceItemWithDefaults()
}

// WorkspaceSourceIterator defines an iterator over a collection.
type WorkspaceSourceIterator struct {
	elements     []WorkspaceSourceItem
	currentIndex int
}

func (m *WorkspaceSourceIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *WorkspaceSourceIterator) GetNext() (item any, err error) {
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

func NewWorkspaceSourceIterator(elements []WorkspaceSourceItem) (IIterator, error) {
	return &WorkspaceSourceIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *WorkspaceSourceCollection) FetchType() string {
	return "WorkspaceSourceCollection page"
}

// FetchLinks returns the resource links if present
func (o *WorkspaceSourceCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *WorkspaceSourceCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *WorkspaceSourceCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *WorkspaceSourceCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *WorkspaceSourceCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewWorkspaceSourceIterator(embedded.GetItem())
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

func (o *WorkspaceSourceCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewWorkspaceSourceCollection returns a page.
func NewWorkspaceSourceCollectionCollection() IStaticPage {
	return NewWorkspaceSourceCollectionWithDefaults()
}

// ============================================================================================
// This extends MessageObject with pagination methods
// ============================================================================================

// FetchType returns the resource type
func (o *MessageObject) FetchType() string {
	return "Message"
}

// MessageIterator defines an iterator over a collection.
type MessageIterator struct {
	elements     []MessageObject
	currentIndex int
}

func (m *MessageIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *MessageIterator) GetNext() (item any, err error) {
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

func NewMessageIterator(elements []MessageObject) (IIterator, error) {
	return &MessageIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// ============================================================================================
// This extends NotificationMessageObject with pagination methods
// ============================================================================================

// FetchType returns the resource type
func (o *NotificationMessageObject) FetchType() string {
	return "Message"
}

// NotificationMessageIterator defines an iterator over a collection.
type NotificationMessageIterator struct {
	elements     []NotificationMessageObject
	currentIndex int
}

func (m *NotificationMessageIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *NotificationMessageIterator) GetNext() (item any, err error) {
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

func NewNotificationMessageIterator(elements []NotificationMessageObject) (IIterator, error) {
	return &NotificationMessageIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// ============================================================================================
// This extends NotificationFeed with pagination methods
// ============================================================================================

// FetchType returns the resource type
func (o *NotificationFeed) FetchType() string {
	return "Notification messages"
}

// FetchLinks returns the resource links if present
func (o *NotificationFeed) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *NotificationFeed) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *NotificationFeed) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *NotificationFeed) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *NotificationFeed) HasFuture() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasFuture()
	}
	return false
}

func (o *NotificationFeed) GetItemIterator() (IIterator, error) {
	return NewNotificationMessageIterator(o.GetMessages())
}

func (o *NotificationFeed) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewNotificationFeedMessageStream returns a message stream.
func NewNotificationFeedMessageStream() IMessageStream {
	return NewNotificationFeedWithDefaults()
}
