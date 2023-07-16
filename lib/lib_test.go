package lib

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	for _, tc := range []struct {
		a        int
		b        int
		expected int
	}{
		{2, 3, 5},
		{-10, 1, -9},
		{-2147483647, 2147483647, 0},
	} {
		require.Equal(t, tc.expected, add(tc.a, tc.b))
	}
}

func TestMul(t *testing.T) {
	for _, tc := range []struct {
		a        int
		b        int
		expected int
	}{
		{-2, 3, -6},
		{1, 2, 2},
		{1, 100000000, 100000000},
		{198, 35267, 6982866},
	} {
		require.Equal(t, tc.expected, mul(tc.a, tc.b))
	}
}

func TestCalc(t *testing.T) {
	for _, tc := range []struct {
		action   string
		a        int
		b        int
		expected int
		isErr    bool
	}{
		{"add", 1, 10, 11, false},
		{"mul", 15, -15, -225, false},
		{"sub", 1, -2, 0, true},
		{"blabla", 1, 1, 0, true},
	} {
		res, err := Calc(tc.action, tc.a, tc.b)
		require.Equal(t, tc.expected, res)
		if tc.isErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
