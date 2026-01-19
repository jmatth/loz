package loz_test

import (
	"fmt"
	"strconv"

	"github.com/jmatth/loz"
	lom "github.com/jmatth/loz/mapping"
)

func Example_foldWithMap() {
	result := lom.Map1[int, string](loz.Generate(5, func(i int) int {
		return i + 1
	})).
		Fold("", func(acc string, n int) string {
			if acc == "" {
				return fmt.Sprintf("%d", n)
			}
			return fmt.Sprintf("%s, %d", acc, n)
		})
	fmt.Printf("%v", result)
	// Output: 1, 2, 3, 4, 5
}

func Example_haltOnErrorWithMap() {
	nums, err := lom.Map1[string, int](loz.IterSlice([]string{"1", "two", "3"})).
		Map(func(str string) int {
			num, err := strconv.Atoi(str)
			loz.PanicHaltIteration(err)
			return num
		}).
		TryCollectSlice()
	fmt.Printf("%v; %v\n", nums, err)
	// Output: []; strconv.Atoi: parsing "two": invalid syntax
}

func Example_skipOnErrorWithMap() {
	nums := lom.Map1[string, int](loz.IterSlice([]string{"1", "two", "3"})).
		FilterMap(func(num string) (int, bool) {
			res, err := strconv.Atoi(num)
			return res, err == nil
		}).
		CollectSlice()
	fmt.Printf("%v\n", nums)
	// Output: [1 3]
}
