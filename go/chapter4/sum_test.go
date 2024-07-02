package chapter4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	require.Equal(t, 10, Sum([]int{2, 3, 5}))
}