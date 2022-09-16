package pkg

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverseInt(t *testing.T) {
	x, err := ReverseInt(123123)
	req := require.New(t)
	req.NoError(err)
	fmt.Println(x)
}
func TestReverseNegativeValue(t *testing.T) {
	x, err := ReverseInt(-960)
	req := require.New(t)
	req.NoError(err)
	fmt.Println(x)
}
func TestReverseZeroValue(t *testing.T) {
	x, err := ReverseInt(0)
	req := require.New(t)
	req.NoError(err)
	fmt.Println(x)
}

// lessons 06.03
// Closure driven
func TestIsPalindrome(t *testing.T) {
	palinCase := func(x int) func(t *testing.T) {
		return func(t *testing.T) {
			res, _ := IsPalindrome(x)
			fmt.Println(res)
		}
	}
	t.Run("first case", palinCase(1))
	t.Run("second case", palinCase(54))
	t.Run("third case", palinCase(43))
}

// Table driven
func TestContainsDuplicate(t *testing.T) {
	tests := map[string][]int{
		"simple":          []int{1, 2, 3},
		"zero values":     []int{0, 0, 0},
		"negative values": []int{-8, -5, -5},
	}
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := ContainsDuplicate(testCase)
			fmt.Println(res)
		})
	}
}
