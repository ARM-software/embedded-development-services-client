/*
 * Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
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

// ============================================================================================
// This extends BuildJobItem and BuildJobCollection definitions
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
// This extends CmsisBuilderItem and CmsisBuilderCollection definitions
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
// This extends CmsisIntellisenseItem and CmsisIntellisenseCollection definitions
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
// This extends FPGAAdminItem and FPGAAdminCollection definitions
// ============================================================================================

// FetchType returns the resource type
func (o *FPGAAdminItem) FetchType() string {
	return "FPGAAdminItem"
}

// FetchLinks returns the resource links if present
func (o *FPGAAdminItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *FPGAAdminItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *FPGAAdminItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewFPGAAdminModel returns a model.
func NewFPGAAdminModel() IModel {
	return NewFPGAAdminItemWithDefaults()
}

// FPGAAdminIterator defines an iterator over a collection.
type FPGAAdminIterator struct {
	elements     []FPGAAdminItem
	currentIndex int
}

func (m *FPGAAdminIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *FPGAAdminIterator) GetNext() (item any, err error) {
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

func NewFPGAAdminIterator(elements []FPGAAdminItem) (IIterator, error) {
	return &FPGAAdminIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *FPGAAdminCollection) FetchType() string {
	return "FPGAAdminCollection page"
}

// FetchLinks returns the resource links if present
func (o *FPGAAdminCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *FPGAAdminCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *FPGAAdminCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *FPGAAdminCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *FPGAAdminCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewFPGAAdminIterator(embedded.GetItem())
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

func (o *FPGAAdminCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewFPGAAdminCollection returns a page.
func NewFPGAAdminCollectionCollection() IStaticPage {
	return NewFPGAAdminCollectionWithDefaults()
}

// ============================================================================================
// This extends FPGAConnectionItem and FPGAConnectionCollection definitions
// ============================================================================================

// FetchType returns the resource type
func (o *FPGAConnectionItem) FetchType() string {
	return "FPGAConnectionItem"
}

// FetchLinks returns the resource links if present
func (o *FPGAConnectionItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *FPGAConnectionItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *FPGAConnectionItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewFPGAConnectionModel returns a model.
func NewFPGAConnectionModel() IModel {
	return NewFPGAConnectionItemWithDefaults()
}

// FPGAConnectionIterator defines an iterator over a collection.
type FPGAConnectionIterator struct {
	elements     []FPGAConnectionItem
	currentIndex int
}

func (m *FPGAConnectionIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *FPGAConnectionIterator) GetNext() (item any, err error) {
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

func NewFPGAConnectionIterator(elements []FPGAConnectionItem) (IIterator, error) {
	return &FPGAConnectionIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *FPGAConnectionCollection) FetchType() string {
	return "FPGAConnectionCollection page"
}

// FetchLinks returns the resource links if present
func (o *FPGAConnectionCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *FPGAConnectionCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *FPGAConnectionCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *FPGAConnectionCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *FPGAConnectionCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewFPGAConnectionIterator(embedded.GetItem())
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

func (o *FPGAConnectionCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewFPGAConnectionCollection returns a page.
func NewFPGAConnectionCollectionCollection() IStaticPage {
	return NewFPGAConnectionCollectionWithDefaults()
}

// ============================================================================================
// This extends FPGAItem and FPGACollection definitions
// ============================================================================================

// FetchType returns the resource type
func (o *FPGAItem) FetchType() string {
	return "FPGAItem"
}

// FetchLinks returns the resource links if present
func (o *FPGAItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *FPGAItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *FPGAItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewFPGAModel returns a model.
func NewFPGAModel() IModel {
	return NewFPGAItemWithDefaults()
}

// FPGAIterator defines an iterator over a collection.
type FPGAIterator struct {
	elements     []FPGAItem
	currentIndex int
}

func (m *FPGAIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *FPGAIterator) GetNext() (item any, err error) {
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

func NewFPGAIterator(elements []FPGAItem) (IIterator, error) {
	return &FPGAIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *FPGACollection) FetchType() string {
	return "FPGACollection page"
}

// FetchLinks returns the resource links if present
func (o *FPGACollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *FPGACollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *FPGACollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *FPGACollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *FPGACollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewFPGAIterator(embedded.GetItem())
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

func (o *FPGACollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewFPGACollection returns a page.
func NewFPGACollectionCollection() IStaticPage {
	return NewFPGACollectionWithDefaults()
}

// ============================================================================================
// This extends FPGAJobItem and FPGAJobCollection definitions
// ============================================================================================

// FetchType returns the resource type
func (o *FPGAJobItem) FetchType() string {
	return "FPGAJobItem"
}

// FetchLinks returns the resource links if present
func (o *FPGAJobItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *FPGAJobItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *FPGAJobItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewFPGAJobModel returns a model.
func NewFPGAJobModel() IModel {
	return NewFPGAJobItemWithDefaults()
}

// FPGAJobIterator defines an iterator over a collection.
type FPGAJobIterator struct {
	elements     []FPGAJobItem
	currentIndex int
}

func (m *FPGAJobIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *FPGAJobIterator) GetNext() (item any, err error) {
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

func NewFPGAJobIterator(elements []FPGAJobItem) (IIterator, error) {
	return &FPGAJobIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *FPGAJobCollection) FetchType() string {
	return "FPGAJobCollection page"
}

// FetchLinks returns the resource links if present
func (o *FPGAJobCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *FPGAJobCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *FPGAJobCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *FPGAJobCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *FPGAJobCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewFPGAJobIterator(embedded.GetItem())
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

func (o *FPGAJobCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewFPGAJobCollection returns a page.
func NewFPGAJobCollectionCollection() IStaticPage {
	return NewFPGAJobCollectionWithDefaults()
}

// ============================================================================================
// This extends FPGAPayloadItem and FPGAPayloadCollection definitions
// ============================================================================================

// FetchType returns the resource type
func (o *FPGAPayloadItem) FetchType() string {
	return "FPGAPayloadItem"
}

// FetchLinks returns the resource links if present
func (o *FPGAPayloadItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *FPGAPayloadItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *FPGAPayloadItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewFPGAPayloadModel returns a model.
func NewFPGAPayloadModel() IModel {
	return NewFPGAPayloadItemWithDefaults()
}

// FPGAPayloadIterator defines an iterator over a collection.
type FPGAPayloadIterator struct {
	elements     []FPGAPayloadItem
	currentIndex int
}

func (m *FPGAPayloadIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *FPGAPayloadIterator) GetNext() (item any, err error) {
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

func NewFPGAPayloadIterator(elements []FPGAPayloadItem) (IIterator, error) {
	return &FPGAPayloadIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *FPGAPayloadCollection) FetchType() string {
	return "FPGAPayloadCollection page"
}

// FetchLinks returns the resource links if present
func (o *FPGAPayloadCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *FPGAPayloadCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *FPGAPayloadCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *FPGAPayloadCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *FPGAPayloadCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewFPGAPayloadIterator(embedded.GetItem())
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

func (o *FPGAPayloadCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewFPGAPayloadCollection returns a page.
func NewFPGAPayloadCollectionCollection() IStaticPage {
	return NewFPGAPayloadCollectionWithDefaults()
}

// ============================================================================================
// This extends GenericWorkJobItem and GenericWorkJobCollection definitions
// ============================================================================================

// FetchType returns the resource type
func (o *GenericWorkJobItem) FetchType() string {
	return "GenericWorkJobItem"
}

// FetchLinks returns the resource links if present
func (o *GenericWorkJobItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *GenericWorkJobItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *GenericWorkJobItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewGenericWorkJobModel returns a model.
func NewGenericWorkJobModel() IModel {
	return NewGenericWorkJobItemWithDefaults()
}

// GenericWorkJobIterator defines an iterator over a collection.
type GenericWorkJobIterator struct {
	elements     []GenericWorkJobItem
	currentIndex int
}

func (m *GenericWorkJobIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *GenericWorkJobIterator) GetNext() (item any, err error) {
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

func NewGenericWorkJobIterator(elements []GenericWorkJobItem) (IIterator, error) {
	return &GenericWorkJobIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *GenericWorkJobCollection) FetchType() string {
	return "GenericWorkJobCollection page"
}

// FetchLinks returns the resource links if present
func (o *GenericWorkJobCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *GenericWorkJobCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *GenericWorkJobCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *GenericWorkJobCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *GenericWorkJobCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewGenericWorkJobIterator(embedded.GetItem())
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

func (o *GenericWorkJobCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewGenericWorkJobCollection returns a page.
func NewGenericWorkJobCollectionCollection() IStaticPage {
	return NewGenericWorkJobCollectionWithDefaults()
}

// ============================================================================================
// This extends GenericWorkerItem and GenericWorkerCollection definitions
// ============================================================================================

// FetchType returns the resource type
func (o *GenericWorkerItem) FetchType() string {
	return "GenericWorkerItem"
}

// FetchLinks returns the resource links if present
func (o *GenericWorkerItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *GenericWorkerItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *GenericWorkerItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewGenericWorkerModel returns a model.
func NewGenericWorkerModel() IModel {
	return NewGenericWorkerItemWithDefaults()
}

// GenericWorkerIterator defines an iterator over a collection.
type GenericWorkerIterator struct {
	elements     []GenericWorkerItem
	currentIndex int
}

func (m *GenericWorkerIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *GenericWorkerIterator) GetNext() (item any, err error) {
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

func NewGenericWorkerIterator(elements []GenericWorkerItem) (IIterator, error) {
	return &GenericWorkerIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *GenericWorkerCollection) FetchType() string {
	return "GenericWorkerCollection page"
}

// FetchLinks returns the resource links if present
func (o *GenericWorkerCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *GenericWorkerCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *GenericWorkerCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *GenericWorkerCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *GenericWorkerCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewGenericWorkerIterator(embedded.GetItem())
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

func (o *GenericWorkerCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewGenericWorkerCollection returns a page.
func NewGenericWorkerCollectionCollection() IStaticPage {
	return NewGenericWorkerCollectionWithDefaults()
}

// ============================================================================================
// This extends InstancePermissionItem and InstancePermissionCollection definitions
// ============================================================================================

// FetchType returns the resource type
func (o *InstancePermissionItem) FetchType() string {
	return "InstancePermissionItem"
}

// FetchLinks returns the resource links if present
func (o *InstancePermissionItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *InstancePermissionItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *InstancePermissionItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewInstancePermissionModel returns a model.
func NewInstancePermissionModel() IModel {
	return NewInstancePermissionItemWithDefaults()
}

// InstancePermissionIterator defines an iterator over a collection.
type InstancePermissionIterator struct {
	elements     []InstancePermissionItem
	currentIndex int
}

func (m *InstancePermissionIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *InstancePermissionIterator) GetNext() (item any, err error) {
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

func NewInstancePermissionIterator(elements []InstancePermissionItem) (IIterator, error) {
	return &InstancePermissionIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *InstancePermissionCollection) FetchType() string {
	return "InstancePermissionCollection page"
}

// FetchLinks returns the resource links if present
func (o *InstancePermissionCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *InstancePermissionCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *InstancePermissionCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *InstancePermissionCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *InstancePermissionCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewInstancePermissionIterator(embedded.GetItem())
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

func (o *InstancePermissionCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewInstancePermissionCollection returns a page.
func NewInstancePermissionCollectionCollection() IStaticPage {
	return NewInstancePermissionCollectionWithDefaults()
}

// ============================================================================================
// This extends IntellisenseJobItem and IntellisenseJobCollection definitions
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
// This extends UserItem and UserCollection definitions
// ============================================================================================

// FetchType returns the resource type
func (o *UserItem) FetchType() string {
	return "UserItem"
}

// FetchLinks returns the resource links if present
func (o *UserItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *UserItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *UserItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

// NewUserModel returns a model.
func NewUserModel() IModel {
	return NewUserItemWithDefaults()
}

// UserIterator defines an iterator over a collection.
type UserIterator struct {
	elements     []UserItem
	currentIndex int
}

func (m *UserIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *UserIterator) GetNext() (item any, err error) {
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

func NewUserIterator(elements []UserItem) (IIterator, error) {
	return &UserIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}

// FetchType returns the resource type
func (o *UserCollection) FetchType() string {
	return "UserCollection page"
}

// FetchLinks returns the resource links if present
func (o *UserCollection) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *UserCollection) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *UserCollection) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *UserCollection) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *UserCollection) GetItemIterator() (IIterator, error) {
	if o.HasEmbedded() {
		embedded := o.GetEmbedded()
		return NewUserIterator(embedded.GetItem())
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

func (o *UserCollection) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewUserCollection returns a page.
func NewUserCollectionCollection() IStaticPage {
	return NewUserCollectionWithDefaults()
}

// ============================================================================================
// This extends VhtInstanceItem and VhtInstanceCollection definitions
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
// This extends VhtItem and VhtCollection definitions
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
// This extends VhtRunJobItem and VhtRunJobCollection definitions
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
// This extends WorkspaceItem and WorkspaceCollection definitions
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
// This extends WorkspaceSourceItem and WorkspaceSourceCollection definitions
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
// This extends MessageObject definitions
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
// This extends NotificationMessageObject definitions
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
// This extends BuildMessageItem definitions
// ============================================================================================

// FetchType returns the resource type
func (o *BuildMessageItem) FetchType() string {
	return "Notification messages"
}

// FetchLinks returns the resource links if present
func (o *BuildMessageItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *BuildMessageItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *BuildMessageItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *BuildMessageItem) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *BuildMessageItem) HasFuture() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasFuture()
	}
	return false
}

func (o *BuildMessageItem) GetItemIterator() (IIterator, error) {
	return NewMessageIterator(o.GetMessages())
}

func (o *BuildMessageItem) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewBuildMessageItemMessageStream returns a message stream.
func NewBuildMessageItemMessageStream() IMessageStream {
	return NewBuildMessageItemWithDefaults()
}

// ============================================================================================
// This extends IntellisenseMessageItem definitions
// ============================================================================================

// FetchType returns the resource type
func (o *IntellisenseMessageItem) FetchType() string {
	return "Notification messages"
}

// FetchLinks returns the resource links if present
func (o *IntellisenseMessageItem) FetchLinks() (links any, err error) {
	if !o.Links.IsSet() {
		err = errors.New("missing links")
		return
	}
	links = o.GetLinks()
	return
}

// FetchName returns the resource name if present, or else an error
func (o *IntellisenseMessageItem) FetchName() (string, error) {
	return o.GetName(), nil
}

// FetchTitle returns the resource title if present, or else an error
func (o *IntellisenseMessageItem) FetchTitle() (string, error) {
	return o.GetTitle(), nil
}

func (o *IntellisenseMessageItem) HasNext() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasNext()
	}
	return false
}

func (o *IntellisenseMessageItem) HasFuture() bool {
	if links, has := o.GetLinksOk(); has {
		return links.HasFuture()
	}
	return false
}

func (o *IntellisenseMessageItem) GetItemIterator() (IIterator, error) {
	return NewMessageIterator(o.GetMessages())
}

func (o *IntellisenseMessageItem) GetItemCount() (count int64, err error) {
	m, ok := o.GetMetadataOk()
	if !ok {
		err = fmt.Errorf("missing metadata: %v", o)
		return
	}
	count = int64(m.GetCount())
	return
}

// NewIntellisenseMessageItemMessageStream returns a message stream.
func NewIntellisenseMessageItemMessageStream() IMessageStream {
	return NewIntellisenseMessageItemWithDefaults()
}

// ============================================================================================
// This extends NotificationFeed definitions
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
