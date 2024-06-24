package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		data []int
		item int
		index int
		step int
	}{
		{data: []int{1, 2, 3}, item: 2, index: 1, step: 1},
		{data: []int{1, 2, 3}, item: 1, index: 0, step: 2},
		{data: []int{1, 2, 3, 4, 5, 6}, item: 4, index: 3, step: 3},
		{data: []int{}, item: 1, index: 0, step: 0},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("test", func(t *testing.T) {
			i, s := BinarySearch(tc.data, tc.item)
			require.Equal(t, tc.index, i)
			require.Equal(t, s, tc.step)
		})
	}
}