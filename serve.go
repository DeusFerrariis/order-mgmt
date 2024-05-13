package main

import (
	"log"
	"net/http"

	_ "github.com/DeusFerrariis/order-mgmt/customer"
	"github.com/DeusFerrariis/order-mgmt/order"
	clog "github.com/charmbracelet/log"
)

func main() {
	mux := http.NewServeMux()
	if err := order.Api.AttachRoutes("order.db", mux); err != nil {
		clog.Error(err)
		return
	}

	http.Handle("/", mux)
	clog.Info("Listening...", "port", 3001)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
