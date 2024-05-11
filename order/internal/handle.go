package internal

import (
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

