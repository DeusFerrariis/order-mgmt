package main

import (
	"log"

	"github.com/DeusFerrariis/order-mgmt/customer"
	clog "github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	if err := customer.AttachCustomerRoutes("orders.db", e); err != nil {
		clog.Error(err)
	}
	log.Fatal(e.Start(":3001"))
}
