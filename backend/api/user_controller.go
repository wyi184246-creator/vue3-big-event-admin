package api

import (
	"backend/constants"
	"backend/dto/request"
	"backend/service"
	"errors"

	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	Login(c *gin.Context)
}
type UserController struct {
	userService service.UserServiceInterface
}

func NewUserController(userService service.UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

func (u *UserController) Login(c *gin.Context) {
	req := &request.UserLoginReq{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}
	if req.Password != req.RePassword {
		c.JSON(200, gin.H{
			"code": 2,
			"msg":  "\"repassword\" is required",
		})
		return
	}
	err = u.userService.Login(req)
	ok := errors.Is(err, constants.ErrUserNotFound)
	if ok {
		c.JSON(200, gin.H{
			"code": 3,
			"msg":  "用户不存在",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "登录成功",
	})
}
