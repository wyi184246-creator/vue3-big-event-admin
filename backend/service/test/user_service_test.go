package test

import (
	"backend/constants"
	"backend/dto/request"
	"backend/mocks"
	"backend/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Login_Success(t *testing.T) {
	req := &request.UserLoginReq{
		UserName: "test-user",
	}

	userDaoMock := mocks.NewUserDaoInterface(t)
	userDaoMock.On("GetUserByName", req).Return(nil).Once()

	userService := service.NewUserService(userDaoMock)
	err := userService.Login(req)
	assert.NoError(t, err)
}

func TestUserService_Login_UserNotFound(t *testing.T) {
	req := &request.UserLoginReq{
		UserName: "not-found-user",
	}

	userDaoMock := mocks.NewUserDaoInterface(t)
	userDaoMock.On("GetUserByName", req).Return(constants.ErrUserNotFound).Once()

	userService := service.NewUserService(userDaoMock)
	err := userService.Login(req)
	assert.ErrorIs(t, err, constants.ErrUserNotFound)
}

func TestUserService_Login_QueryFailed(t *testing.T) {
	req := &request.UserLoginReq{
		UserName: "query-failed-user",
	}
	queryErr := errors.New("query failed")

	userDaoMock := mocks.NewUserDaoInterface(t)
	userDaoMock.On("GetUserByName", req).Return(queryErr).Once()

	userService := service.NewUserService(userDaoMock)
	err := userService.Login(req)
	assert.ErrorIs(t, err, queryErr)
}
