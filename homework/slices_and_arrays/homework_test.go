package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values   []int
	capacity int
	first    int // индекс первого заполненного элемента
	length   int // сколько элементов заполнено
}

func NewCircularQueue(capacity int) CircularQueue {
	return CircularQueue{
		values:   make([]int, capacity),
		capacity: capacity,
		first:    0,
		length:   0,
	}
}

// [_,_,_]
// first 0, length 0, cap 3
// [0,_,_]
// first 0, length 1, cap 3
// [0,1,_]
// first 0, length 2, cap 3
// [0,1,2]
// first 0, length 3, cap 3
// [_,1,2]
// first 1, length 2, cap 3
// [_,_,2]
// first 2, length 1, cap 3
// [3,_,2]
// first 2, length 2, cap 3
// [3,4,2]
// first 2, length 3, cap 3

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	i := q.first + q.length
	if i >= q.capacity {
		// Если вышли за правую границу, то идём по кругу с начала
		i = i - q.capacity
	}

	q.values[i] = value
	q.length++

	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	// Двигаем индексы, сами элементы в срезе не удаляем
	q.first++
	q.length--
	if q.first == q.capacity {
		// Если вышли за правую границу, то идём на начало
		q.first = 0
	}

	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.first]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}

	i := q.first + q.length - 1
	if i >= q.capacity {
		// Если вышли за правую границу, то идём по кругу с начала
		i = i - q.capacity
	}

	return q.values[i]
}

func (q *CircularQueue) Empty() bool {
	return q.length == 0
}

func (q *CircularQueue) Full() bool {
	return q.length == q.capacity
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 3, queue.Back())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
}
