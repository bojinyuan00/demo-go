package user_service

import (
	"demo-go/app/dao/product_dao"
	"demo-go/app/dao/user_dao"
	"demo-go/app/model"
	"demo-go/common/log"
	"demo-go/common/utils"
	"time"
)

type Service struct {
	UserDao    *user_dao.Dao
	ProductDao *product_dao.Dao
}

// NewService 初始化 UserService
func NewService(userDao *user_dao.Dao, productDao *product_dao.Dao) *Service {
	return &Service{
		UserDao:    userDao,
		ProductDao: productDao,
	}
}

func (u *Service) GetUserInfoById(userId int) (result *model.User, err error, customErr error, message string) {
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
