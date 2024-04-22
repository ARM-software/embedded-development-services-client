/*
 * Copyright (C) 2020-2024 Arm Limited or its affiliates and Contributors. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */
package client

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testIterator struct {
	next      any
	itemCount int64
}

type testMessageStream struct {
	i         testIterator
	hasFuture bool
}

func (t *testIterator) HasNext() bool {
	return t.next != nil
}

func (t *testIterator) GetNext() (any, error) {
	return t.next, nil
}

func (t *testIterator) GetItemCount() (int64, error) {
	return t.itemCount, nil
}

func (t *testMessageStream) HasNext() bool {
	return t.i.HasNext()
}

func (t *testMessageStream) GetNext() (any, error) {
	return t.i.GetNext()
}

func (t *testMessageStream) GetItemCount() (int64, error) {
	return t.i.GetItemCount()
}

func (t *testMessageStream) GetItemIterator() (IIterator, error) {
	return &t.i, nil
}

func (t *testMessageStream) HasFuture() bool {
	return t.hasFuture
}

func TestExtensionModelInterfaceMappers(t *testing.T) {
	var x testMessageStream
	err := faker.FakeData(&x)
	require.NoError(t, err)

	y := MapIMessageStream(&x)

	x1, err := x.GetItemCount()
	assert.NoError(t, err)
	y1, err := y.GetItemCount()
	assert.NoError(t, err)
	assert.EqualValues(t, x1, y1)

	x2 := x.HasNext()
	y2 := y.HasNext()
	assert.EqualValues(t, x2, y2)

	x3 := x.HasFuture()
	y3 := y.HasFuture()
	assert.EqualValues(t, x3, y3)

	x4, err := x.GetItemIterator()
	assert.NoError(t, err)
	y4, err := y.GetItemIterator()
	assert.NoError(t, err)
	assert.EqualValues(t, x4, y4)
}
