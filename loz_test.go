package loz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhere(t *testing.T) {
	filteredSlice := Values([]bool{true, false, true, false, true}).Where(
		func(b bool) bool {
			return !b
		}).ToSlice()
	assert.Equal(t, []bool{false, false}, filteredSlice)
}

func TestSkip(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	skipped := Values(nums).Skip(3).ToSlice()
	assert.Equal(t, []int{4, 5, 6, 7, 8, 9}, skipped)
}

func TestAll(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	skipped := Values(nums).Skip(100).ToSlice()
	assert.Len(t, skipped, 0)
}

func TestTake(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	took := Values(nums).Take(3).ToSlice()
	assert.Equal(t, []int{1, 2, 3}, took)
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
