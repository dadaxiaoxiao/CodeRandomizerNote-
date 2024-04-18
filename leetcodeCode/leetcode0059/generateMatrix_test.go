package leetcode0059

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestGenerateMatrix(t *testing.T) {
	testCases := []struct {
		name    string
		n       int
		wantRes [][]int
	}{
		{
			name:    "n =1,输出 [[1]]",
			n:       1,
			wantRes: [][]int{{1}},
		},
		{
			name:    "n=3,输出[[1,2,3],[8,9,4],[7,6,5]]",
			n:       3,
			wantRes: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := generateMatrix(tc.n)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
