package service

import (
	"backend/constants"
	"backend/dao"
	"backend/dto/request"
	"errors"

	"go.uber.org/zap"
)

type UserServiceInterface interface {
	Login(req *request.UserLoginReq) error
}

type UserService struct {
	userDao dao.UserDaoInterface
}

func NewUserService(userDao dao.UserDaoInterface) *UserService {
	return &UserService{userDao: userDao}
}

func (u *UserService) Login(req *request.UserLoginReq) error {
	err := u.userDao.GetUserByName(req)
	if err != nil {
		ok := errors.Is(err, constants.ErrUserNotFound)
		if ok {
			zap.L().Error("用户不存在", zap.String("username", req.UserName))
			return constants.ErrUserNotFound
		}
		zap.L().Error("查询用户失败", zap.String("username", req.UserName), zap.Error(err))
		return err
	}
	return nil
}
