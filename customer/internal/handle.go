package internal

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DatabaseContext struct {
	echo.Context
	Store CustomerStore
}

func CreateCustomerHandler(c echo.Context) error {
	store := c.(DatabaseContext).Store
	var data CustomerData
	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	id, err := store.CreateCustomer(data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusAccepted, map[string]any{
		"id": id,
	})
}

func GetCustomerHandler(c echo.Context) error {
	store := c.(DatabaseContext).Store
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.String(http.StatusBadRequest, "id is invalid")
	}
	data, err := store.GetCustomer(int64(id))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, data)
}
