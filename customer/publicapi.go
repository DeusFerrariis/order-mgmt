package customer

import (
	"database/sql"

	"github.com/DeusFerrariis/order-mgmt/customer/internal"
	"github.com/labstack/echo/v4"
)

var (
	AttachCustomerRoutes = func(dbPath string, e *echo.Echo) error {
		// Database middleware
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			return err
		}
		custStore, err := internal.NewCustomerStore(db)
		if err != nil {
			return err
		}
		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				cc := internal.DatabaseContext{Context: c, Store: *custStore}
				return next(&cc)
			}
		})

		// Routes
		e.GET("/customers/:id", internal.GetCustomerHandler)
		e.POST("/customers", internal.CreateCustomerHandler)
		return nil
	}
)
