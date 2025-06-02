package loz_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/jmatth/loz"
	"github.com/stretchr/testify/assert"
)

func ExampleSeq_TakeWhile() {
	result := loz.IterSlice([]int{2, 4, 5, 6, 8}).
		TakeWhile(func(n int) bool { return n%2 == 0 }).
		CollectSlice()
	fmt.Printf("%v", result)
	// Output: [2 4]
}

func ExampleSeq_SkipWhile() {
	result := loz.IterSlice([]int{2, 4, 5, 6, 8}).
		SkipWhile(func(n int) bool { return n%2 == 0 }).
		CollectSlice()
	fmt.Printf("%v", result)
	// Output: [5 6 8]
}

func ExampleSeq_Filter() {
	filteredSlice := loz.IterSlice([]bool{true, false, true, false, true}).
		Filter(
			func(b bool) bool {
				return !b
			}).CollectSlice()
	fmt.Printf("%v", filteredSlice)
	// Output: [false false]
}

func ExampleSeq_Skip() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	skipped := loz.IterSlice(nums).Skip(3).CollectSlice()
	fmt.Printf("%v", skipped)
	// Output: [4 5 6 7 8 9]
}

func ExampleSeq_Take() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	took := loz.IterSlice(nums).Take(3).CollectSlice()
	fmt.Printf("%v", took)
	// Output: [1 2 3]
}

func ExampleSeq_Map() {
	nums := []int{1, 2, 3}
	doubled := loz.IterSlice(nums).Map(func(n int) int { return n * 2 }).CollectSlice()
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
	anyEven := loz.IterSlice(nums).Any(isEven)
	anyBig := loz.IterSlice(nums).Any(isBig)
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
	anyEven := loz.IterSlice(nums).None(isEven)
	anyBig := loz.IterSlice(nums).None(isBig)
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
	anyEven := loz.IterSlice(nums).Every(isOdd)
	anyBig := loz.IterSlice(nums).Every(isBig)
	fmt.Printf("%v, %v", anyEven, anyBig)
	// Output: true, false
}

func ExampleSeq_Expand() {
	expander := func(n int) loz.Seq[int] {
		return func(yield func(int) bool) {
			for i := range n {
				if !yield(i + 1) {
					break
				}
			}
		}
	}

	nums := []int{1, 2, 3, 0, 5}
	expanded := loz.IterSlice(nums).Expand(expander).CollectSlice()
	fmt.Printf("%v", expanded)
	// Output: [1 1 2 1 2 3 1 2 3 4 5]
}

func ExampleSeq_First() {
	first, err := loz.IterSlice([]int{}).First()
	fmt.Printf("%v, %v\n", first, err != nil)
	first, err = loz.IterSlice([]int{1, 2, 3}).First()
	fmt.Printf("%v, %v", first, err)
	// Output: 0, true
	// 1, <nil>
}

func ExampleSeq_Last() {
	first, err := loz.IterSlice([]int{}).Last()
	fmt.Printf("%v, %v\n", first, err != nil)
	first, err = loz.IterSlice([]int{1, 2, 3}).Last()
	fmt.Printf("%v, %v", first, err)
	// Output: 0, true
	// 3, <nil>
}

func ExampleSeq_Fold() {
	mult := func(a, b int) int {
		return a * b
	}
	foldEmpty := loz.IterSlice([]int{}).Fold(100, mult)
	foldVals := loz.IterSlice([]int{2, 4, 2}).Fold(1, mult)
	fmt.Printf("%v, %v", foldEmpty, foldVals)
	// Output: 100, 16
}

func ExampleSeq_Reduce() {
	mult := func(a, b int) int {
		return a * b
	}
	reduced, err := loz.IterSlice([]int{}).Reduce(mult)
	fmt.Printf("%v, %v\n", reduced, err != nil)
	reduced, err = loz.IterSlice([]int{2, 4, 2}).Reduce(mult)
	fmt.Printf("%v, %v\n", reduced, err)
	// Output: 0, true
	// 16, <nil>
}

func ExampleSeq_ForEach() {
	loz.IterSlice([]int{1, 2, 3}).ForEach(func(i int) {
		fmt.Printf("(%v)", i)
	})
	// Output: (1)(2)(3)
}

func TestSkipAll(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	skipped := loz.IterSlice(nums).Skip(100).CollectSlice()
	assert.Empty(t, skipped)
}

func TestSkipAndTake(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	took := loz.IterSlice(nums).Skip(3).Take(3).CollectSlice()
	assert.Equal(t, []int{4, 5, 6}, took)
}

func TestRepeatCalls(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	baseSeq := loz.IterSlice(nums).Skip(3)
	took := baseSeq.Take(3).CollectSlice()
	skipped := baseSeq.Skip(3).CollectSlice()
	assert.Equal(t, []int{4, 5, 6}, took)
	assert.Equal(t, []int{7, 8, 9}, skipped)
}

func TestTMP(t *testing.T) {
}

func ExampleSeq_errorHandling() {
	result, err := loz.Map1[string, int](loz.IterSlice([]string{"1", "foo", "3"})).
		Map(func(s string) int {
			num, err := strconv.Atoi(s)
			if err != nil {
				loz.PanicHaltIteration(err)
			}
			return num
		}).TryCollectSlice()
	fmt.Printf("%v; %v", result, err)
	// Output: []; strconv.Atoi: parsing "foo": invalid syntax
}

func ExampleSeq_incorrectErrorHandling() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("example code panicked: %v", r)
		}
	}()

	result, err := loz.Map1[string, int](loz.IterSlice([]string{"1", "foo", "3"})).
		Map(func(s string) int {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			return num
		}).TryCollectSlice()
	fmt.Printf("%v; %v", result, err)
	// Output: example code panicked: strconv.Atoi: parsing "foo": invalid syntax
}
