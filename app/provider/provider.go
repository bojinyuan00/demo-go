//go:build wireinject
// +build wireinject

package provider

import (
	"demo-go/app/controller"
	"demo-go/app/dao"
	"demo-go/app/service"
	"github.com/google/wire"
)

// InitializeUserController 自动生成构造函数
func InitializeUserController() (*controller.UserController, error) {
	wire.Build(
		dao.NewUserDao,
		service.NewUserService,
		controller.NewUserController,
	)
	return nil, nil
}
