package heaply_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viridovics/heaply"
)

func TestMaxHeapTest(t *testing.T) {
	maxHeap := heaply.NewHeap(func(a, b int) int { return a - b })

	assert.True(t, maxHeap.Size() == 0)

	_, ok := maxHeap.Pop()
	assert.False(t, ok)
	_, ok = maxHeap.Top()
	assert.False(t, ok)

	rounds := 77777
	i := 0
	for i < rounds {
		maxHeap.Push(rand.Int())
		i++
	}

	prev := math.MaxInt

	i = 0
	for i < rounds {
		maxTop, ok := maxHeap.Top()
		assert.True(t, ok)
		assert.True(t, prev >= maxTop, fmt.Sprintf("%d, %d", prev, maxTop))
		maxPop, ok := maxHeap.Pop()
		assert.True(t, ok)
		assert.Equal(t, maxTop, maxPop)
		i++
		prev = maxTop
	}
}

func TestExampleMaxHeapTest(t *testing.T) {
	maxHeap := heaply.NewHeap(func(a, b int) int { return a - b })

	maxHeap.Push(10)
	maxHeap.Push(5)
	maxHeap.Push(20)
	maxHeap.Push(1)
	maxHeap.Push(15)
	maxHeap.Push(8)

	max, ok := maxHeap.Pop()
	assert.True(t, ok)
	assert.Equal(t, 20, max)

	max, ok = maxHeap.Top()
	assert.True(t, ok)
	assert.Equal(t, 15, max)
}
