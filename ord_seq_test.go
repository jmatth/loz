package loz_test

import (
	"fmt"
	"testing"

	"github.com/jmatth/loz"
	"github.com/stretchr/testify/assert"
)

func ExampleOrdSeq_Max() {
	iter := loz.OrdSeq[int](loz.IterSlice([]int{3, 2, 5, -3, 0}))
	maxVal, err := iter.Max()
	fmt.Printf("%v, %v", maxVal, err)
	// Output: 5, <nil>
}

func TestOrdSeqMaxEmpty(t *testing.T) {
	iter := loz.OrdSeq[int](loz.IterSlice([]int{}))
	maxVal, err := iter.Max()
	assert.ErrorIs(t, err, loz.EmptySeqErr)
	assert.Equal(t, 0, maxVal)
}

func ExampleOrdSeq_Min() {
	iter := loz.OrdSeq[int](loz.IterSlice([]int{3, 2, 5, -3, 0}))
	minVal, err := iter.Min()
	fmt.Printf("%v, %v", minVal, err)
	// Output: -3, <nil>
}

func TestOrdSeqMinEmpty(t *testing.T) {
	iter := loz.OrdSeq[int](loz.IterSlice([]int{}))
	maxVal, err := iter.Min()
	assert.ErrorIs(t, err, loz.EmptySeqErr)
	assert.Equal(t, 0, maxVal)
}
