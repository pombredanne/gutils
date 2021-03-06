package gutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func TestCircularArray(t *testing.T) {
	circ := NewCircularArray(10)

	assert.Equal(
		t, circ.Dump(), "CircularArray {size: 10, start: 0, end: 0, buffer: []}",
	)

	assert.Equal(t, circ.Capacity(), uint(10))
	assert.Equal(t, circ.Length(), uint(0))

	v, err := circ.Pop()
	assert.Equal(t, err, Empty)
	assert.Equal(t, v, nil)

	v, err = circ.PopNewest()
	assert.Equal(t, err, Empty)
	assert.Equal(t, v, nil)

	v, err = circ.PeekNewest()
	assert.Equal(t, err, Empty)
	assert.Equal(t, v, nil)

	v, err = circ.PeekOldest()
	assert.Equal(t, err, Empty)
	assert.Equal(t, v, nil)

	v, err = circ.Ith(0)
	assert.Equal(t, err, Empty)
	assert.Equal(t, v, nil)

	// one element in the array
	o, d := circ.Push(10)
	assert.Equal(
		t, circ.Dump(),
		"CircularArray {size: 10, start: 0, end: 1, buffer: [10]}",
	)
	assert.Nil(t, o)
	assert.False(t, d)

	v, err = circ.PeekNewest()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 10)
	assert.Equal(
		t, circ.Dump(),
		"CircularArray {size: 10, start: 0, end: 1, buffer: [10]}",
	)

	v, err = circ.PeekOldest()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 10)
	assert.Equal(
		t, circ.Dump(),
		"CircularArray {size: 10, start: 0, end: 1, buffer: [10]}",
	)

	v, err = circ.Ith(0)
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 10)
	assert.Equal(
		t, circ.Dump(),
		"CircularArray {size: 10, start: 0, end: 1, buffer: [10]}",
	)

	v, err = circ.Ith(1)
	assert.Equal(t, err, IndexOutOfBounds)
	assert.Equal(t, v, nil)
	assert.Equal(
		t, circ.Dump(),
		"CircularArray {size: 10, start: 0, end: 1, buffer: [10]}",
	)

	v, err = circ.Pop()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 10)
	assert.Equal(
		t, circ.Dump(),
		"CircularArray {size: 10, start: 0, end: 0, buffer: []}",
	)

	circ = NewCircularArray(10)
	o, d = circ.Push(10)
	assert.Nil(t, o)
	assert.False(t, d)

	v, err = circ.PopNewest()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 10)
	assert.Equal(
		t, "CircularArray {size: 10, start: 0, end: 0, buffer: []}",
		circ.Dump(),
	)

	o, d = circ.Push(20)
	assert.Nil(t, o)
	assert.False(t, d)

	assert.Equal(
		t, "CircularArray {size: 10, start: 0, end: 1, buffer: [20]}",
		circ.Dump(),
	)

	v, err = circ.PeekNewest()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 20)
	assert.Equal(
		t, "CircularArray {size: 10, start: 0, end: 1, buffer: [20]}",
		circ.Dump(),
	)

	v, err = circ.PeekOldest()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 20)
	assert.Equal(
		t, "CircularArray {size: 10, start: 0, end: 1, buffer: [20]}",
		circ.Dump(),
	)

	v, err = circ.Ith(0)
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 20)
	assert.Equal(
		t, "CircularArray {size: 10, start: 0, end: 1, buffer: [20]}",
		circ.Dump(),
	)

	v, err = circ.Ith(1)
	assert.Equal(t, err, IndexOutOfBounds)
	assert.Equal(t, v, nil)
	assert.Equal(
		t, "CircularArray {size: 10, start: 0, end: 1, buffer: [20]}",
		circ.Dump(),
	)

	v, err = circ.Pop()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 20)
	assert.Equal(
		t, "CircularArray {size: 10, start: 0, end: 0, buffer: []}",
		circ.Dump(),
	)

	o, d  = circ.Push(30)
	assert.Nil(t, o)
	assert.False(t, d)

	o, d  = circ.Push(31)
	assert.Nil(t, o)
	assert.False(t, d)

	o, d  = circ.Push(32)
	assert.Nil(t, o)
	assert.False(t, d)

	o, d  = circ.Push(33)
	assert.Nil(t, o)
	assert.False(t, d)

	o, d  = circ.Push(34)
	assert.Nil(t, o)
	assert.False(t, d)

	o, d  = circ.Push(35)
	assert.Nil(t, o)
	assert.False(t, d)

	o, d  = circ.Push(36)
	assert.Nil(t, o)
	assert.False(t, d)

	o, d  = circ.Push(37)
	assert.Nil(t, o)
	assert.False(t, d)

	o, d  = circ.Push(38)
	assert.Nil(t, o)
	assert.False(t, d)

	o, d  = circ.Push(39)
	assert.Nil(t, o)
	assert.False(t, d)

	o, d  = circ.Push(40)
	assert.Equal(t, o, 30)
	assert.True(t, d)

	assert.Equal(
		t,
		"CircularArray {size: 10, start: 1, end: 11, buffer: [40 31 32 33 34 35 36 37 38 39]}",
		circ.Dump(),
	)

	v, err = circ.Pop()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 31)
	assert.Equal(
		t,
		"CircularArray {size: 10, start: 2, end: 11, buffer: [40 31 32 33 34 35 36 37 38 39]}",
		circ.Dump(),
	)

	circ.Pop() // 32
	circ.Pop()
	circ.Pop()
	circ.Pop()
	circ.Pop()
	circ.Pop()
	circ.Pop()
	circ.Pop()

	v, err = circ.Pop()

	assert.Equal(t, err, nil)
	assert.Equal(t, v, 40)
	assert.Equal(
		t,
		"CircularArray {size: 10, start: 0, end: 0, buffer: []}",
		circ.Dump(),
	)

	v, err = circ.Pop()
	assert.Equal(t, err, Empty)
	assert.Equal(t, v, nil)
	assert.Equal(
		t,
		"CircularArray {size: 10, start: 0, end: 0, buffer: []}",
		circ.Dump(),
	)

	o, d = circ.Push(50)
	assert.Nil(t, o)
	assert.False(t, d)

	o, d = circ.Push(51)
	assert.Nil(t, o)
	assert.False(t, d)

	o, d = circ.Push(52)
	assert.Nil(t, o)
	assert.False(t, d)

	o, d = circ.Push(53)
	assert.Nil(t, o)
	assert.False(t, d)

	assert.Equal(
		t,
		"CircularArray {size: 10, start: 0, end: 4, buffer: [50 51 52 53]}",
		circ.Dump(),
	)

	v, err = circ.PopNewest()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 53)

	v, err = circ.PeekNewest()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 52)
	assert.Equal(
		t, circ.Dump(),
		"CircularArray {size: 10, start: 0, end: 3, buffer: [50 51 52 53]}",
		circ.Dump(),
	)

	v, err = circ.PeekOldest()
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 50)
	assert.Equal(
		t, circ.Dump(),
		"CircularArray {size: 10, start: 0, end: 3, buffer: [50 51 52 53]}",
		circ.Dump(),
	)

	v, err = circ.Ith(0)
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 50)
	assert.Equal(
		t, circ.Dump(),
		"CircularArray {size: 10, start: 0, end: 3, buffer: [50 51 52 53]}",
		circ.Dump(),
	)

	v, err = circ.Ith(2)
	assert.Equal(t, err, nil)
	assert.Equal(t, v, 52)
	assert.Equal(
		t, circ.Dump(),
		"CircularArray {size: 10, start: 0, end: 3, buffer: [50 51 52 53]}",
		circ.Dump(),
	)

	v, err = circ.Ith(3)
	assert.Equal(t, err, IndexOutOfBounds)
	assert.Equal(t, v, nil)
	assert.Equal(
		t, circ.Dump(),
		"CircularArray {size: 10, start: 0, end: 3, buffer: [50 51 52 53]}",
		circ.Dump(),
	)

	circ.PopNewest()
	circ.PopNewest()
	circ.PopNewest()
	assert.Equal(t, uint(0), circ.Length())

	v, err = circ.PopNewest()
	assert.Equal(t, err, Empty)
	assert.Equal(t, v, nil)

	for i := 100; i < 110; i++ {
		o, d := circ.Push(i)
		assert.Nil(t, o)
		assert.False(t, d)
	}

	for i := 110; i < 200; i++ {
		o, d := circ.Push(i)
		fmt.Println(i, o.(int), int(i-10))
		assert.Equal(t, o.(int), int(i-10))
		assert.True(t, d)
	}

	assert.Equal(t, uint(10), circ.Length())

	for i := 199; i >= 190; i-- {
		v, err = circ.PopNewest()
		assert.Equal(t, nil, err)
		assert.Equal(t, i, v)
	}

	v, err = circ.PopNewest()
	assert.Equal(t, Empty, err)
	assert.Equal(t, nil, v)
}

func TestCircularBufferArray(t *testing.T) {
	circ := NewCircularBufferArray(10)

	assert.Equal(
		t, circ.Dump(), "CircularArray {size: 10, start: 0, end: 0, buffer: []}",
	)

	assert.Equal(t, circ.Capacity(), uint(10))
	assert.Equal(t, circ.Length(), uint(0))

	o, d := circ.Push([]byte("foo"))
	assert.Nil(t, o)
	assert.False(t, d)

	v, err := circ.Pop()

	assert.Equal(t, []byte("foo"), v)
	assert.Equal(t, nil, err)

	v, err = circ.Pop()

	assert.Equal(t, []byte(nil), v)
	assert.Equal(t, Empty, err)
}
