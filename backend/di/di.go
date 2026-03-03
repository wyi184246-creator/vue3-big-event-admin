package di

import "go.uber.org/fx"

var controllerModule = fx.Module("controller", fx.Provide())

var serviceModule = fx.Module("service", fx.Provide())

var daoModule = fx.Module("dao", fx.Provide())

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
