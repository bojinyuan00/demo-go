package dao

import (
	"demo-go/app/model"
	"demo-go/common/global"
	"demo-go/common/log"
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
func (u *UserDao) GetUserById(id int) (*model.User, error) {
	defer log.TimeTracker(time.Now()) // 记录执行栈

	var user model.User
	result := global.GormDB.First(&user, id)
	if result.Error != nil {
		log.ErrorLogger(result.Error, "获取用户信息失败-UserDao.GetUserById")
		return nil, result.Error
	}
	return &user, nil
}
