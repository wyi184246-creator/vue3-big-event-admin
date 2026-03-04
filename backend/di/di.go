package di

import (
	"backend/api"
	"backend/dao"
	"backend/service"

	"go.uber.org/fx"
)

var controllerModule = fx.Module("controller", fx.Provide(
	fx.Annotate(api.NewUserController, fx.As(new(api.UserControllerInterface))),
))

var serviceModule = fx.Module("service", fx.Provide(
	service.NewUserService,
	fx.Annotate(service.NewUserService, fx.As(new(service.UserServiceInterface))),
))

var daoModule = fx.Module("dao", fx.Provide(
	dao.NewGormDB,
	fx.Annotate(dao.NewUserDao, fx.As(new(dao.UserDaoInterface))),
))

var provideConfig = fx.Options(
	controllerModule,
	serviceModule,
	daoModule,
)

var fxConfig = fx.Options(
	provideConfig,
	fx.Invoke(func() {

	}),
)
