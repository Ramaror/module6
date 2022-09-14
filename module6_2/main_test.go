package module6_2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	mock_client "module6/module6_2/mocks"
	"testing"
)

func TestPostCount(t *testing.T) {
	client := mock_client.MockClient{}

	list, err := client.GetList()
	assert.NoError(t, err)

	fmt.Println(list)

}
