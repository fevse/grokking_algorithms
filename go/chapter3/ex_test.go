package chapter3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCoutndown(t *testing.T) {
	x := countdown(5)
	require.Equal(t, 15, x)
}