package order

import (
	"net/http"

	"github.com/DeusFerrariis/order-mgmt/order/internal"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

var (
	AttachOrderRoutes = func(dbPath string, e *echo.Echo) error {
		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				cc, err := internal.NewLineItemStoreContext(dbPath, c)
				if err != nil {
					log.Error(err)
					return c.NoContent(http.StatusInternalServerError)
				}
				return next(cc)
			}
		})

		e.POST("/orders/:orderId/lineItems", internal.CreateLineItemHandler)
		return nil
	}
)
