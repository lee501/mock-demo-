package user

import (
	"github.com/golang/mock/gomock"
	"mock-demo/mockd"
	"testing"
)

func TestUser_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id int64 = 1
	mockMale := mockd.NewMockMale(ctl)
	gomock.InOrder(
		mockMale.EXPECT().Get(id).Return(nil),
	)

	user := NewUser(mockMale)
	err := user.GetUserInfo(id)
	if err != nil {
		t.Errorf("user.GetUserInfo err : %v", err)
	}
}
