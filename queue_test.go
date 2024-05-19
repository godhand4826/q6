package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinQueueItem(t *testing.T) {
	var q = NewQ(func(v1, v2 *int) bool {
		return *v1 < *v2
	})
	q.PushItem(NewItem(AsRef(0)))
	q.PushItem(NewItem(AsRef(1)))
	q.PushItem(NewItem(AsRef(3)))
	q.PushItem(NewItem(AsRef(4)))
	q.PushItem(NewItem(AsRef(2)))

	assert.Equal(t, 0, *q.PopItem().Value)
	assert.Equal(t, 1, *q.PopItem().Value)
	assert.Equal(t, 2, *q.PopItem().Value)
	assert.Equal(t, 3, *q.PopItem().Value)
	assert.Equal(t, 4, *q.PopItem().Value)
}

func TestMaxQueueItem(t *testing.T) {
	var q = NewQ(func(v1, v2 *int) bool {
		return *v1 > *v2
	})
	q.PushItem(NewItem(AsRef(0)))
	q.PushItem(NewItem(AsRef(1)))
	q.PushItem(NewItem(AsRef(3)))
	q.PushItem(NewItem(AsRef(4)))
	q.PushItem(NewItem(AsRef(2)))

	assert.Equal(t, 4, *q.PopItem().Value)
	assert.Equal(t, 3, *q.PopItem().Value)
	assert.Equal(t, 2, *q.PopItem().Value)
	assert.Equal(t, 1, *q.PopItem().Value)
	assert.Equal(t, 0, *q.PopItem().Value)
}
