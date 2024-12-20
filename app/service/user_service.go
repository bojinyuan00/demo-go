package service

import (
	"demo-go/app/dao"
	"demo-go/app/model"
)

type UserService struct {
	UserDao *dao.UserDao
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{UserDao: userDao}
}

func (u *UserService) GetUserInfoById(userId int) (*model.User, error) {
	//defer log.TimeTracker(time.Now()) // 记录执行栈

	result, err := u.UserDao.GetUserById(userId)
	//time.Sleep(600 * time.Millisecond) // 模拟慢请求
	if err != nil {
		//log.ErrorLogger(err, "获取用户信息失败-UserService.GetUserInfoById")
		return nil, err
	}
	return result, nil
}
