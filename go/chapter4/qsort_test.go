package chapter4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQsort(t *testing.T) {
	arr := []int{10, 5, 2, 3}
	require.Equal(t, []int{2, 3, 5, 10}, Qsort(arr))
}