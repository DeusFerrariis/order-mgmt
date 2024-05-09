package customer

import (
	"net/http"

	"github.com/DeusFerrariis/order-mgmt/customer/internal"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

var (
	AttachCustomerRoutes = func(dbPath string, e *echo.Echo) error {
		// Database middleware
		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				cc, err := internal.NewCustomerStoreContext("orders.db", c)
				if err != nil {
					log.Error(err)
					return c.NoContent(http.StatusInternalServerError)
				}
				return next(cc)
			}
		})

		// Routes
		e.GET("/customers/:id", internal.GetCustomerHandler)
		e.POST("/customers", internal.CreateCustomerHandler)
		return nil
	}
)
