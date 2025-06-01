package loz_test

import (
	"fmt"
	"testing"

	"github.com/jmatth/loz"
	"github.com/stretchr/testify/assert"
)

func ExampleRange() {
	nums := loz.Range(3).CollectSlice()
	fmt.Printf("%v", nums)
	// Output: [0 1 2]
}

func TestRangeFrom_InvalidInput(t *testing.T) {
	nums := loz.RangeFrom(3, 3).CollectSlice()
	assert.Empty(t, nums)
	nums = (loz.RangeFrom(4, 3)).CollectSlice()
	assert.Empty(t, nums)
}

func ExampleRangeInterval() {
	nums := loz.RangeInterval(2, 11, 2).CollectSlice()
	fmt.Printf("%v", nums)
	// Output: [2 4 6 8 10]
}

func ExampleRangeInterval_negativeInterval() {
	nums := loz.RangeInterval(1, -6, -2).CollectSlice()
	fmt.Printf("%v", nums)
	// Output: [1 -1 -3 -5]
}

func ExampleRangeInterval_floats() {
	nums := loz.RangeInterval(1.0, 3, 0.5).CollectSlice()
	fmt.Printf("%v", nums)
	// Output: [1 1.5 2 2.5]
}

func TestRangeInterval_InvalidInput(t *testing.T) {
	nums := loz.RangeInterval(2, 2, 1).CollectSlice()
	assert.Empty(t, nums)
	nums = loz.RangeInterval(2, 1, 1).CollectSlice()
	assert.Empty(t, nums)
	nums = loz.RangeInterval(0, 10, -1).CollectSlice()
	assert.Empty(t, nums)
	nums = loz.RangeInterval(0, 10, 0).CollectSlice()
	assert.Empty(t, nums)
}

func ExampleRangeFrom() {
	nums := loz.RangeFrom(3, 6).CollectSlice()
	fmt.Printf("%v", nums)
	// Output: [3 4 5]
}

func ExampleNumSeq_Sum() {
	iter := loz.Range(4)
	fmt.Printf("%v", iter.Sum())
	// Output: 6
}

func ExampleNumSeq_Sum_empty() {
	iter := loz.Range(0)
	fmt.Printf("%v: %v", iter.CollectSlice(), iter.Sum())
	// Output: []: 0
}

func ExampleNumSeq_Sum_negatives() {
	iter := loz.RangeInterval(0, -4, -1)
	fmt.Printf("%v", iter.Sum())
	// Output: -6
}
