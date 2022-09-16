package counter

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mock_client "module6/internal/app/services/post/mocks"
	"testing"
)

func TestPostCount(t *testing.T) {
	req := require.New(t)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	newMock := mock_client.NewMockClient(mockCtrl)
	newMock.EXPECT().GetList()
	cl, err := newMock.GetList()
	req.NoError(err)
	fmt.Println(cl)
}
