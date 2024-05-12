package internal

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

func bCreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	
}

func CreateCustomerHandler(c echo.Context) error {
	store := c.(*CustomerStoreContext)
	var data CustomerData
	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	rec, err := store.CreateCustomer(data)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusAccepted, map[string]any{
		"data": rec,
	})
}

func GetCustomerHandler(c echo.Context) error {
	store := c.(*CustomerStoreContext)
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.String(http.StatusBadRequest, "id is invalid")
	}
	data, err := store.GetCustomer(int64(id))
	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	}
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, data)
}
