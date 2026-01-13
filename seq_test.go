package loz_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/jmatth/loz"
	lom "github.com/jmatth/loz/mapping"
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
	result, err := lom.Map1[string, int](loz.IterSlice([]string{"1", "foo", "3"})).
		Map(func(s string) int {
			num, err := strconv.Atoi(s)
			loz.PanicHaltIteration(err)
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

	result, err := lom.Map1[string, int](loz.IterSlice([]string{"1", "foo", "3"})).
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

func ExampleSeq_AppendSlice() {
	s := make([]int, 0, 5)
	s = append(s, 1, 2)
	loz.Generate(3, func(idx int) int {
		return idx + 3
	}).AppendSlice(&s)
	fmt.Print(s)
	// Output: [1 2 3 4 5]
}

func ExampleSeq_FilterMap() {
	matching := loz.IterSlice([]int{0, 10, 2, 3, 44, 55}).
		FilterMap(func(i int) (int, bool) {
			if i < 10 {
				return 0, false
			}
			return i / 10, true
		}).
		CollectSlice()
	fmt.Print(matching)
	// Output: [1 4 5]
}

func TestSeqTryMethods(t *testing.T) {
	seq := loz.Generate(5, func(idx int) int {
		return idx
	})
	haltingErr := errors.New("Testing error")
	haltingSeq := seq.Map(func(i int) int {
		loz.PanicHaltIteration(haltingErr)
		return i
	})
	table := []struct {
		name string
		run  func(loz.Seq[int]) error
	}{
		{
			"TryForEach",
			func(s loz.Seq[int]) error {
				return s.TryForEach(func(i int) {})
			},
		},
		{
			"TryReduce",
			func(s loz.Seq[int]) error {
				_, err := s.TryReduce(func(i1 int, i2 int) int {
					return i1 + i2
				})
				return err
			},
		},
		{
			"TryFold",
			func(s loz.Seq[int]) error {
				_, err := s.TryFold(0, func(i1 int, i2 int) int {
					return i1 + i2
				})
				return err
			},
		},
		{
			"TryFirst",
			func(s loz.Seq[int]) error {
				_, err := s.TryFirst()
				return err
			},
		},
		{
			"TryLast",
			func(s loz.Seq[int]) error {
				_, err := s.TryLast()
				return err
			},
		},
		{
			"TryAny",
			func(s loz.Seq[int]) error {
				_, err := s.TryAny(func(i int) bool {
					return i%2 == 0
				})
				return err
			},
		},
		{
			"TryNone",
			func(s loz.Seq[int]) error {
				_, err := s.TryNone(func(i int) bool {
					return i%2 == 0
				})
				return err
			},
		},
		{
			"TryEvery",
			func(s loz.Seq[int]) error {
				_, err := s.TryEvery(func(i int) bool {
					return i%2 == 0
				})
				return err
			},
		},
	}

	for _, row := range table {
		t.Run(row.name, func(t *testing.T) {
			err := row.run(seq)
			assert.Nil(t, err)
			err = row.run(haltingSeq)
			assert.Equal(t, err, haltingErr)
		})
	}
}
