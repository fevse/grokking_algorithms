package chapter4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxn(t *testing.T) {
	require.Equal(t, 9, Maxn([]int{2, 4, 7, 8, 9, 6, 3, 1}))
}