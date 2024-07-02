package chapter4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCount(t *testing.T) {
	require.Equal(t, 3, Count([]int{1, 2, 3}))
}