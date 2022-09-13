package pkg

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverseInt1(t *testing.T) {
	x, err := ReverseInt(123123)
	req := require.New(t)
	req.NoError(err)
	fmt.Println(x)
}
func TestReverseInt2(t *testing.T) {
	x, err := ReverseInt(-967)
	req := require.New(t)
	req.NoError(err)
	fmt.Println(x)
}
func TestReverseInt3(t *testing.T) {
	x, err := ReverseInt(0)
	req := require.New(t)
	req.NoError(err)

	fmt.Println(x)
}
