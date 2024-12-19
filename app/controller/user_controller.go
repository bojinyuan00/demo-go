package controller

import (
	"demo-go/app/service"
	"demo-go/common/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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
	defer log.TimeTracker(time.Now()) // 记录执行栈

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := u.UserService.GetUserInfoById(id)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	c.JSON(http.StatusOK, user)
}