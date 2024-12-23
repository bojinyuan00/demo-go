//go:build wireinject
// +build wireinject

package provider

import (
	"demo-go/app/controller/user_controller"
	"demo-go/app/dao/product_dao"
	"demo-go/app/dao/user_dao"
	"demo-go/app/service/user_service"
	"github.com/google/wire"
)

// InitializeUserController 自动生成构造函数
func InitializeUserController() (*user_controller.Controller, error) {
	wire.Build(
		user_dao.NewDao,
		product_dao.NewDao,
		user_service.NewService,
		user_controller.NewController,
	)
	return nil, nil
}
