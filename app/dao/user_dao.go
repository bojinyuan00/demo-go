package dao

import (
	"demo-go/app/model"
	"demo-go/common/global"
	"demo-go/common/log"
	"demo-go/common/utils"
	"time"
)

// UserDao 数据访问对象
type UserDao struct {
}

// NewUserDao 创建 UserDao 实例
func NewUserDao() *UserDao {
	return &UserDao{}
}

// GetUserById 根据ID获取用户信息
func (u *UserDao) GetUserById(id int) (user *model.User, err error) {
	start := time.Now()

	//初始化日志字段
	logParams := make(map[string]interface{})
	logParams["id"] = id
	defer func() { log.TimeTracker(start, logParams, user, err) }() // defer 捕获日志并执行

	result := global.GormDB.First(&user, id)
	err = result.Error

	if err != nil {
		err = utils.CaptureError(err, "获取用户信息失败-UserDao.GetUserById")
		return nil, err
	}
	return user, nil
}
