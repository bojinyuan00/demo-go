package controller

import (
	"demo-go/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService *service.UserService
}

// NewUserController 创建 UserController 实例
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

// GetUserInfo 获取用户信息
func (u *UserController) GetUserInfo(c *gin.Context) {
	//defer log.TimeTracker(time.Now(), requestParam, responseParam) // 记录执行栈

	id, _ := strconv.Atoi(c.Param("id"))
	user, err, _, message := u.UserService.GetUserInfoById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "请求失败：" + message})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUserListInfo 获取用户列表信息
func (u *UserController) GetUserListInfo(c *gin.Context) {
	//defer log.TimeTracker(time.Now(), requestParam, responseParam) // 记录执行栈

	id, _ := strconv.Atoi(c.Param("id"))
	user, err, _, message := u.UserService.GetUserInfoById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "请求失败：" + message})
		return
	}
	c.JSON(http.StatusOK, user)
}
