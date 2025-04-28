package loz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleSeq_TakeWhile() {
	result := Values([]int{2, 4, 5, 6, 8}).
		TakeWhile(func(n int) bool { return n%2 == 0 }).
		ToSlice()
	fmt.Printf("%v", result)
	// Output: [2 4]
}

func ExampleSeq_SkipWhile() {
	result := Values([]int{2, 4, 5, 6, 8}).
		SkipWhile(func(n int) bool { return n%2 == 0 }).
		ToSlice()
	fmt.Printf("%v", result)
	// Output: [5 6 8]
}

func ExampleSeq_Filter() {
	filteredSlice := Values([]bool{true, false, true, false, true}).
		Filter(
			func(b bool) bool {
				return !b
			}).ToSlice()
	fmt.Printf("%v", filteredSlice)
	// Output: [false false]
}

func ExampleSeq_Skip() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	skipped := Values(nums).Skip(3).ToSlice()
	fmt.Printf("%v", skipped)
	// Output: [4 5 6 7 8 9]
}

func ExampleSeq_Take() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	took := Values(nums).Take(3).ToSlice()
	fmt.Printf("%v", took)
	// Output: [1 2 3]
}

func ExampleSeq_Map() {
	nums := []int{1, 2, 3}
	doubled := Values(nums).Map(func(n int) int { return n * 2 }).ToSlice()
	fmt.Printf("%v", doubled)
	// Output: [2 4 6]
}

func ExampleSeq_Any() {
	isEven := func(n int) bool {
		return n%2 == 0
	}
	isBig := func(n int) bool {
		return n > 9_000
	}

	nums := []int{1, 3, 7, 9_001}
	anyEven := Values(nums).Any(isEven)
	anyBig := Values(nums).Any(isBig)
	fmt.Printf("%v, %v", anyEven, anyBig)
	// Output: false, true
}

func ExampleSeq_None() {
	isEven := func(n int) bool {
		return n%2 == 0
	}
	isBig := func(n int) bool {
		return n > 9_000
	}

	nums := []int{1, 3, 7, 9_001}
	anyEven := Values(nums).None(isEven)
	anyBig := Values(nums).None(isBig)
	fmt.Printf("%v, %v", anyEven, anyBig)
	// Output: true, false
}

func ExampleSeq_Every() {
	isOdd := func(n int) bool {
		return n%2 != 0
	}
	isBig := func(n int) bool {
		return n > 9_000
	}

	nums := []int{1, 3, 7, 9_001}
	anyEven := Values(nums).Every(isOdd)
	anyBig := Values(nums).Every(isBig)
	fmt.Printf("%v, %v", anyEven, anyBig)
	// Output: true, false
}

func ExampleSeq_Expand() {
	expander := func(n int) Seq[int] {
		return func(yield func(int) bool) {
			for i := range n {
				if !yield(i + 1) {
					break
				}
			}
		}
	}

	nums := []int{1, 2, 3, 0, 5}
	expanded := Values(nums).Expand(expander).ToSlice()
	fmt.Printf("%v", expanded)
	// Output: [1 1 2 1 2 3 1 2 3 4 5]
}

func ExampleSeq_First() {
	first, err := Values([]int{}).First()
	fmt.Printf("%v, %v\n", first, err)
	first, err = Values([]int{1, 2, 3}).First()
	fmt.Printf("%v, %v", first, err)
	// Output: 0, First called on empty Seq
	// 1, <nil>
}

func ExampleSeq_Last() {
	first, err := Values([]int{}).Last()
	fmt.Printf("%v, %v\n", first, err)
	first, err = Values([]int{1, 2, 3}).Last()
	fmt.Printf("%v, %v", first, err)
	// Output: 0, Last called on empty Seq
	// 3, <nil>
}

func ExampleSeq_Fold() {
	mult := func(a, b int) int {
		return a * b
	}
	foldEmpty := Values([]int{}).Fold(100, mult)
	foldVals := Values([]int{2, 4, 2}).Fold(1, mult)
	fmt.Printf("%v, %v", foldEmpty, foldVals)
	// Output: 100, 16
}

func ExampleSeq_Reduce() {
	mult := func(a, b int) int {
		return a * b
	}
	reduced, err := Values([]int{}).Reduce(mult)
	fmt.Printf("%v, %v\n", reduced, err)
	reduced, err = Values([]int{2, 4, 2}).Reduce(mult)
	fmt.Printf("%v, %v\n", reduced, err)
	// Output: 0, Reduce called on empty Seq
	// 16, <nil>
}

func ExampleSeq_ForEach() {
	Values([]int{1, 2, 3}).ForEach(func(i int) {
		fmt.Printf("(%v)", i)
	})
	// Output: (1)(2)(3)
}

func TestSkipAll(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	skipped := Values(nums).Skip(100).ToSlice()
	assert.Empty(t, skipped)
}

func TestSkipAndTake(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	took := Values(nums).Skip(3).Take(3).ToSlice()
	assert.Equal(t, []int{4, 5, 6}, took)
}

func TestRepeatCalls(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	baseSeq := Values(nums).Skip(3)
	took := baseSeq.Take(3).ToSlice()
	skipped := baseSeq.Skip(3).ToSlice()
	assert.Equal(t, []int{4, 5, 6}, took)
	assert.Equal(t, []int{7, 8, 9}, skipped)
}
