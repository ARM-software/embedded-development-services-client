/*
 * Copyright (C) 2020-2024 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package client

import (
	"fmt"

	"github.com/ARM-software/golang-utils/utils/collection/pagination"
	"github.com/ARM-software/golang-utils/utils/commonerrors"
)

type staticPageWrapper struct {
	page IStaticPage
}

func (s *staticPageWrapper) HasNext() bool {
	return s.page.HasNext()
}

func (s *staticPageWrapper) GetItemCount() (int64, error) {
	return s.page.GetItemCount()
}

func (s *staticPageWrapper) GetItemIterator() (pagination.IIterator, error) {
	i, err := s.page.GetItemIterator()
	if err != nil {
		return nil, err
	}
	i2, ok := i.(pagination.IIterator)
	if !ok {
		return nil, fmt.Errorf("%w: could not convert client.IIterator to pagination.IIterator", commonerrors.ErrMarshalling)
	}
	return i2, nil
}

func MapIStaticPage(staticPage IStaticPage) pagination.IStaticPage {
	if staticPage == nil {
		return nil
	}
	return &staticPageWrapper{staticPage}
}

type staticPageStreamWrapper struct {
	pagination.IStaticPage
	stream IMessageStream
}

func (s *staticPageStreamWrapper) HasFuture() bool {
	return s.stream.HasFuture()
}

func MapIMessageStream(messageStream IMessageStream) pagination.IStaticPageStream {
	if messageStream == nil {
		return nil
	}
	return &staticPageStreamWrapper{IStaticPage: MapIStaticPage(messageStream), stream: messageStream}
}
