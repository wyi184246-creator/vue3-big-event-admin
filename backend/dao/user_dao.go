package dao

import (
	"backend/constants"
	"backend/dto/request"
	"backend/model"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserDaoInterface interface {
	GetUserByName(req *request.UserLoginReq) error
}
type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (u *UserDao) GetUserByName(req *request.UserLoginReq) error {
	user := &model.User{}
	err := u.db.Where("username = ?", req.UserName).First(user).Error
	if err != nil {
		ok := errors.Is(err, gorm.ErrRecordNotFound)
		if ok {
			zap.L().Error("用户不存在", zap.String("username", req.UserName))
			return constants.ErrUserNotFound
		}
		zap.L().Error("查询用户失败", zap.String("username", req.UserName), zap.Error(err))
		return err
	}
	return nil
}
