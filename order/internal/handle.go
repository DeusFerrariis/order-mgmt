package internal

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateLineItemHandler(c echo.Context) error {
	// Collect /{id}
	orderIdParam := c.Param("orderId")
	orderId, err := strconv.Atoi(orderIdParam)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid order id")
	}
	// Collect json body
	var data struct {
		ItemData
		ItemPricing
		Quantity int64 `json:"quantity"`
	}
	if err := c.Bind(&data); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	// Attempt create
	s := c.(*LineItemStoreContext)
	rec, err := s.CreateLineItem(LineItemData{
		data.ItemData,
		data.ItemPricing,
		0,
		data.Quantity,
		int64(orderId),
	})
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	// Return record
	return c.JSON(http.StatusAccepted, map[string]any{"data": rec})
}

func (ctx *OrderStoreContext) CreateOrderStmt() (*sql.Stmt, error) {
	return ctx.db.Prepare("INSERT INTO orders VALUES(NULL, ?);")
}

func (ctx *OrderStoreContext) CreateOrder(custId int64) (*OrderRecord, error) {
	stmt, err := ctx.CreateOrderStmt()
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(custId)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
}
