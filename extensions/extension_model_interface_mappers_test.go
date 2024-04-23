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
	t.Run("Normal case", func(t *testing.T) {
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
	})

	t.Run("Nil case", func(t *testing.T) {
		t.Run("IStaticPage", func(t *testing.T) {
			var x IStaticPage = nil
			y := MapIStaticPage(x)
			assert.Nil(t, y)
			assert.Equal(t, x, y)
		})
		t.Run("IMessageStream", func(t *testing.T) {
			var x IMessageStream = nil
			y := MapIMessageStream(x)
			assert.Nil(t, y)
			assert.Equal(t, x, y)
		})
	})
}
