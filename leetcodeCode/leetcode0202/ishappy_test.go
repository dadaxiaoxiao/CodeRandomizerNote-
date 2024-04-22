package leetcode0202

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsHappy(t *testing.T) {
	testCases := []struct {
		name    string
		n       int
		wantRes bool
	}{
		{
			name:    "n=19",
			n:       19,
			wantRes: true,
		},
		{
			name:    "n=2",
			n:       2,
			wantRes: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := isHappy(tc.n)
			assert.Equal(t, tc.wantRes, res)
		})
	}

}
