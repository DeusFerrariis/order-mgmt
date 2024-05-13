package order

import (
	"net/http"

	"github.com/DeusFerrariis/order-mgmt/handle"
	"github.com/DeusFerrariis/order-mgmt/order/internal"
	clog "github.com/charmbracelet/log"
)

type OrderAPI struct{}

func (*OrderAPI) AttachRoutes(dbPath string, mux *http.ServeMux) error {
	liStore, err := internal.NewLineItemStore(dbPath)
	if err != nil {
		clog.Error("could not make db")
		return err
	}
	controller := internal.OrderController{
		LineItemStore: &liStore,
	}
	routes := map[string]handle.HandlerFunc{
		"POST /orders/{orderId}/line-items":     controller.HCreateLineItem,
		"GET /orders/{orderId}/line-items/{id}": controller.HGetLineItemById,
	}
	for route, handler := range routes {
		handle.Attach(mux, route, handler)
	}
	return nil
}

var Api = OrderAPI{}
