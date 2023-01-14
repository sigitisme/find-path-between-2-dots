package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	testcases := []struct {
		name     string
		input    [][]int
		m        int
		n        int
		expected bool
	}{
		{
			name: "5x5 array with no obstacles",
			input: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			m:        5,
			n:        5,
			expected: true,
		},
		{
			name: "5x5 array with little resistances",
			input: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, -1, 0, 0},
				{0, -1, 0, 0, 0},
			},
			m:        5,
			n:        5,
			expected: true,
		},
		{
			name: "5x5 array with high resistances",
			input: [][]int{
				{0, 0, 0, 0, 0},
				{-1, -1, -1, -1, 0},
				{0, 0, 0, 0, 0},
				{0, -1, -1, -1, -1},
				{0, 0, 0, 0, 0},
			},
			m:        5,
			n:        5,
			expected: true,
		},
		{
			name: "5x5 array with high resistances",
			input: [][]int{
				{0, 0, 0, 0, 0},
				{-1, -1, -1, -1, 0},
				{0, 0, 0, 0, 0},
				{0, -1, -1, -1, -1},
				{0, 0, 0, 0, 0},
			},
			m:        5,
			n:        5,
			expected: true,
		},
		{
			name: "5x5 array with case example resistances ",
			input: [][]int{
				{0, 0, 0, -1, 0},
				{-1, 0, 0, -1, -1},
				{0, 0, 0, -1, 0},
				{-1, 0, 0, 0, 0},
				{0, 0, -1, 0, 0},
			},
			m:        5,
			n:        5,
			expected: true,
		},
		{
			name: "5x4 array with case example resistances ",
			input: [][]int{
				{0, 0, 0, -1},
				{-1, 0, 0, -1},
				{0, 0, 0, -1},
				{-1, 0, 0, 0},
				{0, 0, -1, 0},
			},
			m:        5,
			n:        4,
			expected: true,
		},
		{
			name: "5x5 array with no path 1",
			input: [][]int{
				{0, 0, 0, 0, 0},
				{-1, -1, -1, -1, -1},
				{0, 0, 0, 0, 0},
				{0, -1, -1, -1, -1},
				{0, 0, 0, 0, 0},
			},
			m:        5,
			n:        5,
			expected: false,
		},
		{
			name: "5x5 array with no path 2",
			input: [][]int{
				{0, 0, 0, 0, 0},
				{-1, -1, -1, -1, 0},
				{0, 0, 0, 0, 0},
				{0, -1, -1, -1, -1},
				{0, 0, 0, -1, 0},
			},
			m:        5,
			n:        5,
			expected: false,
		},
	}

	for _, tc := range testcases {
		actual := findPath(tc.m, tc.n, tc.input)

		if actual != tc.expected {
			t.Errorf("%s: expected not same with actual. expected: %v - actual: %v", tc.name, tc.expected, actual)
		}
	}

}
