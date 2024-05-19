package main

import "container/heap"

type LessFn[V any] func(*V, *V) bool

type Queue[V any] struct {
	Items  []*Item[V]
	lessFn LessFn[V]
}

type Item[V any] struct {
	Value *V
	Index int
}

func NewItem[V any](value *V) *Item[V] {
	return &Item[V]{
		Value: value,
		Index: -1,
	}
}

func GetItemValue[V any](item *Item[V]) *V {
	return item.Value
}

func NewQ[V any](lessFn LessFn[V], values ...*V) *Queue[V] {
	q := &Queue[V]{
		lessFn: lessFn,
	}

	for _, v := range values {
		q.PushItem(NewItem(v))
	}

	q.Init()

	return q
}

// PushItem pushes an item onto the queue.
func (q *Queue[V]) PushItem(item *Item[V]) {
	heap.Push(q, item)
}

// PushItems pushes items onto the queue.
func (q *Queue[V]) PushItems(value ...*Item[V]) {
	for _, v := range value {
		q.PushItem(v)
	}
}

// PopItem pops the minimum item from the queue.
func (q *Queue[V]) PopItem() *Item[V] {
	return heap.Pop(q).(*Item[V])
}

// PopItems pops n items from the queue.
func (q *Queue[V]) PopItems(n int) []*Item[V] {
	items := make([]*Item[V], 0, n)
	for i := 0; i < n; i++ {
		items = append(items, q.PopItem())
	}
	return items
}

// Init initializes the queue.
func (q *Queue[V]) Init() {
	heap.Init(q)
}

// Remove removes an item by index
func (q *Queue[V]) Remove(i int) {
	heap.Remove(q, i)
}

// Fix fix the heap ordering for the item at index i
func (q *Queue[V]) Fix(i int) {
	heap.Fix(q, i)
}

// Peek peek the top item of the queue
func (q *Queue[V]) PeekItem() *Item[V] {
	return q.Items[0]
}

// ---------- Start implement heap.Interface ----------

// Implement heap.Interface
var _ heap.Interface = (*Queue[int])(nil)

// Len implements heap.Interface.
func (q *Queue[V]) Len() int {
	return len(q.Items)
}

// Less implements heap.Interface.
func (q *Queue[V]) Less(i int, j int) bool {
	return q.lessFn(q.Items[i].Value, q.Items[j].Value)
}

// Pop implements heap.Interface.
func (q *Queue[V]) Pop() any {
	old := q.Items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	q.Items = old[0 : n-1]
	return item
}

// Push implements heap.Interface.
func (q *Queue[V]) Push(x any) {
	n := len(q.Items)
	item := x.(*Item[V])
	item.Index = n
	q.Items = append(q.Items, item)
}

// Swap implements heap.Interface.
func (q *Queue[V]) Swap(i int, j int) {
	q.Items[i], q.Items[j] = q.Items[j], q.Items[i]
	q.Items[i].Index = i
	q.Items[j].Index = j
}

// ---------- End implement heap.Interface ----------
