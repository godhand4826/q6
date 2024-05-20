package queue

import (
	"testing"

	"q6/lib/util"

	"github.com/stretchr/testify/assert"
)

func TestMinQueueItem(t *testing.T) {
	var q = NewQ(func(v1, v2 *int) bool {
		return *v1 < *v2
	})
	q.PushItem(NewItem(util.AsRef(0)))
	q.PushItem(NewItem(util.AsRef(1)))
	q.PushItem(NewItem(util.AsRef(3)))
	q.PushItem(NewItem(util.AsRef(4)))
	q.PushItem(NewItem(util.AsRef(2)))

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
	q.PushItem(NewItem(util.AsRef(0)))
	q.PushItem(NewItem(util.AsRef(1)))
	q.PushItem(NewItem(util.AsRef(3)))
	q.PushItem(NewItem(util.AsRef(4)))
	q.PushItem(NewItem(util.AsRef(2)))

	assert.Equal(t, 4, *q.PopItem().Value)
	assert.Equal(t, 3, *q.PopItem().Value)
	assert.Equal(t, 2, *q.PopItem().Value)
	assert.Equal(t, 1, *q.PopItem().Value)
	assert.Equal(t, 0, *q.PopItem().Value)
}

func TestNew(t *testing.T) {
	var q = NewQ(func(v1, v2 *int) bool {
		return *v1 < *v2
	}, util.AsRef(1), util.AsRef(2))
	assert.Equal(t, 2, q.Len())
	assert.Equal(t, 1, *q.PopItem().Value)
	assert.NotNil(t, 2, *q.PopItem().Value)
}

func TestPushPopItem(t *testing.T) {
	var q = NewQ(func(v1, v2 *int) bool {
		return *v1 < *v2
	})
	q.PushItem(NewItem(util.AsRef(2)))
	q.PushItem(NewItem(util.AsRef(1)))
	assert.Equal(t, 2, q.Len())
	assert.Equal(t, 1, *q.PopItem().Value)
	assert.Equal(t, 2, *q.PopItem().Value)
}

func TestPushPoPItems(t *testing.T) {
	var q = NewQ(func(v1, v2 *int) bool {
		return *v1 < *v2
	})
	q.PushItems(NewItem(util.AsRef(2)), NewItem(util.AsRef(1)))
	assert.Equal(t, 2, q.Len())
	items := q.PopItems(2)
	assert.Equal(t, 1, *items[0].Value)
	assert.Equal(t, 2, *items[1].Value)
}

func TestRemove(t *testing.T) {
	var q = NewQ(func(v1, v2 *int) bool {
		return *v1 < *v2
	}, util.AsRef(1), util.AsRef(2))
	assert.Equal(t, 2, q.Len())
	q.Remove(1)
	assert.Equal(t, 1, *q.PopItem().Value)
}

func TestFix(t *testing.T) {
	var q = NewQ(func(v1, v2 *int) bool {
		return *v1 < *v2
	}, util.AsRef(1), util.AsRef(2))
	assert.Equal(t, 2, q.Len())
	q.Items[0].Value = util.AsRef(3)
	q.Fix(0)
	assert.Equal(t, 2, *q.PopItem().Value)
	assert.Equal(t, 3, *q.PopItem().Value)
}

func TestPeekItem(t *testing.T) {
	var q = NewQ(func(v1, v2 *int) bool {
		return *v1 < *v2
	})
	q.PushItem(NewItem(util.AsRef(2)))
	assert.Equal(t, 2, *q.PeekItem().Value)
}

func TestGetItemValue(t *testing.T) {
	var item = NewItem(util.AsRef(1))
	assert.Equal(t, 1, *GetItemValue(item))
}
