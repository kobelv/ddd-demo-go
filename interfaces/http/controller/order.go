package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ddd-demo-go/application/appservice"
	"ddd-demo-go/application/dto"
	"ddd-demo-go/infrastructure/common/request"
	"ddd-demo-go/infrastructure/common/response"
)

type Order struct {
	request.BindInterface
	response.HTTPResponseInterface
	app *appservice.OrderAS
}

// NewOrder 实例化
func NewOrder(req request.BindInterface, res response.HTTPResponseInterface, app *appservice.OrderAS) *Order {
	return &Order{BindInterface: req, HTTPResponseInterface: res, app: app}
}

// PlaceOrder 下单入口
func (order *Order) PlaceOrder(c *gin.Context) {
	dto := dto.OrderDTO{}
	if err := order.Bind(c, &dto); err != nil {
		order.RenderJSONResponse(c, http.StatusOK, nil, err)
		return
	}
	res, err := order.app.PlaceOrder(c, &dto)
	order.RenderJSONResponse(c, http.StatusOK, res, err)
	return
}
