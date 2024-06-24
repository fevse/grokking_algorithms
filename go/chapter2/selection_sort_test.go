package selection_sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		data []int
		exp []int
	}{
		{data: []int{1, 2, 3}, exp: []int{1, 2, 3}},
		{data: []int{2, 1}, exp: []int{1, 2}},
		{data: []int{99, 2, 44, 5}, exp: []int{2, 5, 44, 99}},
		{data: []int{6, 5, 4, 3, 2, 1}, exp: []int{1, 2, 3, 4, 5, 6}},
		{data: []int{}, exp: []int{}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("test", func(t *testing.T) {
			res := SelectionSort(tc.data)
			require.Equal(t, tc.exp, res)
		})
	}
}