package internal

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/DeusFerrariis/order-mgmt/handle"
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
		*ItemData
		*ItemPricing
		Quantity *int64 `json:"quantity"`
	}
	if err = c.BodyJson(&body); err != nil {
		return c.LogErr(http.StatusUnprocessableEntity, err)
	}
	if body.ItemData == nil ||
		body.ItemPricing == nil ||
		body.Quantity == nil {
		return c.String(http.StatusUnprocessableEntity, "missing required field")
	}
	lineItem := LineItemData{
		*body.ItemData,
		*body.ItemPricing,
		0,
		*body.Quantity,
		int64(orderId),
	}
	store := cr.LineItemStore
	rec, err := store.CreateLineItem(lineItem)
	if err != nil {
		return err
	}
	// Return record
	return c.JSON(http.StatusAccepted, map[string]any{"data": rec})
}

func (cr *OrderController) HGetLineItemById(c *handle.HandlerContext) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Stringf(http.StatusUnprocessableEntity, "Invalid id '%s'", c.Param("id"))
	}

	store := cr.LineItemStore
	rec, err := store.GetLineItemById(int64(id))
	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{"data": rec})
}
