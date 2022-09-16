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
