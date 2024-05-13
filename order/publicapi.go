package order

import (
	"net/http"

	"github.com/DeusFerrariis/order-mgmt/handle"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

type OrderAPI struct{}

func (*OrderAPI) AttachRoutes(dbPath string, mux *http.ServeMux) error {
	liStore, err := internal.NewLineItemStore(dbPath)
	if err != nil {
		return err
	}
	controller := internal.OrderController{
		LineItemStore: &liStore,
	}
	routes := map[string]handle.HandlerFunc{
		"POST /orders/{orderId}/line-items": controller.HCreateLineItem,
	}
	for route, handler := range routes {
		handle.Attach(mux, route, handler)
	}
	return nil
}

var api = OrderAPI{}

var (
	AttachOrderRoutes = func(dbPath string, e *echo.Echo) error {
		e.POST("/orders/:orderId/lineItems", internal.CreateLineItemHandler)
		return nil
	}
)
