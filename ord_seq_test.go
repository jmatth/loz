package loz_test

import (
	"fmt"
	"testing"

	"github.com/jmatth/loz"
	"github.com/stretchr/testify/assert"
)

func ExampleOrdSeq_Max() {
	iter := loz.OrdSeq[int](loz.IterSlice([]int{3, 2, 5, -3, 0}))
	max, err := iter.Max()
	fmt.Printf("%v, %v", max, err)
	// Output: 5, <nil>
}

func TestOrdSeqMaxEmpty(t *testing.T) {
	iter := loz.OrdSeq[int](loz.IterSlice([]int{}))
	max, err := iter.Max()
	assert.ErrorIs(t, err, loz.EmptySeqErr)
	assert.Equal(t, 0, max)
}

func ExampleOrdSeq_Min() {
	iter := loz.OrdSeq[int](loz.IterSlice([]int{3, 2, 5, -3, 0}))
	min, err := iter.Min()
	fmt.Printf("%v, %v", min, err)
	// Output: -3, <nil>
}

func TestOrdSeqMinEmpty(t *testing.T) {
	iter := loz.OrdSeq[int](loz.IterSlice([]int{}))
	max, err := iter.Min()
	assert.ErrorIs(t, err, loz.EmptySeqErr)
	assert.Equal(t, 0, max)
}