package internal

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/DeusFerrariis/order-mgmt/handle"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	*LineItemStore
}

func (cr *OrderController) HCreateLineItem(c *handle.HandlerContext) error {
	orderId, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		return err
	}
	var body struct {
		ItemData
		ItemPricing
		Quantity int64 `json:"quantity"`
	}
	if err = c.BodyJson(&body); err != nil {
		return c.LogErr(http.StatusUnprocessableEntity, err)
	}
	lineItem := LineItemData{
		body.ItemData,
		body.ItemPricing,
		0,
		body.Quantity,
		int64(orderId),
	}
	store := cr.LineItemStore
	rec, err := store.CreateLineItem(lineItem)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	// Return record
	return c.JSON(http.StatusAccepted, map[string]any{"data": rec})
}

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
	s := c.(*LineItemStore)
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

func (ctx *OrderStore) CreateOrderStmt() (*sql.Stmt, error) {
	return ctx.db.Prepare("INSERT INTO orders VALUES(NULL, ?);")
}

func (ctx *OrderStore) CreateOrder(custId int64) (*OrderRecord, error) {
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
