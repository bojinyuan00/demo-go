// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	"demo-go/app/controller"
	"demo-go/app/dao"
	"demo-go/app/service"
)

// Injectors from provider.go:

// InitializeUserController 自动生成构造函数
func InitializeUserController() (*controller.UserController, error) {
	userDao := dao.NewUserDao()
	userService := service.NewUserService(userDao)
	userController := controller.NewUserController(userService)
	return userController, nil
}
