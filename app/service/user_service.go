package service

import (
	"demo-go/app/dao"
	"demo-go/app/model"
	"demo-go/common/log"
	"demo-go/common/utils"
	"time"
)

type UserService struct {
	UserDao *dao.UserDao
}

// NewUserService 初始化 UserService
func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{UserDao: userDao}
}

func (u *UserService) GetUserInfoById(userId int) (result *model.User, err error, customErr error, message string) {
	start := time.Now()                       //访问时间
	logParams := make(map[string]interface{}) //request参数集合
	logParams["userId"] = userId
	defer func() { log.TimeTracker(start, logParams, result, customErr) }() // defer 捕获日志并执行

	result, err, _, message = u.UserDao.GetUserById(userId)
	//time.Sleep(600 * time.Millisecond) // 模拟慢请求
	if err != nil {
		customErr = utils.CaptureError(err, "获取用户信息失败-UserService.GetUserInfoById")
		return nil, err, customErr, message
	}

	return result, err, customErr, message
}
