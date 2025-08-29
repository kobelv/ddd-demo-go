//go:build wireinject
// +build wireinject

package interfaces

import (
	"context"

	"github.com/google/wire"

	"ddd-demo-go/application/appservice"
	"ddd-demo-go/application/dto"
	"ddd-demo-go/infrastructure/adapter"
	"ddd-demo-go/infrastructure/common/cache"
	"ddd-demo-go/infrastructure/common/db"
	"ddd-demo-go/infrastructure/common/logit"
	"ddd-demo-go/infrastructure/common/request"
	"ddd-demo-go/infrastructure/common/response"
	"ddd-demo-go/infrastructure/po"
	"ddd-demo-go/infrastructure/repository"
	"ddd-demo-go/interfaces/http"
	"ddd-demo-go/interfaces/http/controller"
)

func NewApp(ctx context.Context) (*app, error) {
	panic(wire.Build(wire.NewSet(
		loadAppConf,
		logit.NewServiceLoggerConf,
		logit.NewServiceLogger,

		db.NewDB,
		cache.NewRedis,
		response.NewHTTPResponseWriter,
		request.NewRequest,

		appservice.NewOrderAS,

		dto.NewOrderConv,

		controller.NewOrder,
		controller.NewHealth,
		http.NewHTTPHandler,
		newHTTPServer,

		po.NewOrderPOConv,

		adapter.NewUserAdapter,
		repository.NewOrderRepository,

		wire.Struct(new(app), "*"),
	)))
}
