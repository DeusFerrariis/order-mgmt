package internal

import (
	"database/sql"
	"sync"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type (
	OrderStoreContext struct {
		echo.Context
		mu sync.Mutex
		db *sql.DB
	}
	LineItemStoreContext struct {
		echo.Context
		mu sync.Mutex
		db *sql.DB
	}
)

func NewOrderStoreContext(path string, c echo.Context) (*OrderStoreContext, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	// TODO: init schema for orders table
	return &OrderStoreContext{
		Context: c,
		db:      db,
	}, nil
}

func NewLineItemStoreContext(path string, c echo.Context) (*LineItemStoreContext, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	// TODO: init schema for orders table
	return &LineItemStoreContext{
		Context: c,
		db:      db,
	}, nil
}

func (ctx *LineItemStoreContext) CreateLineItemStmt() (*sql.Stmt, error) {
	return ctx.db.Prepare(`
		INSERT INTO lineItems VALUES(NULL, ?, ?, ?, ?, ?, ?, ?);
	`)
}

func (ctx *LineItemStoreContext) CreateLineItem(data LineItemData) (*LineItemRecord, error) {
	stmt, err := ctx.CreateLineItemStmt()
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(data.OrderId, data.LineNumber, data.Sku, data.Description,
		data.Cost, data.Price, data.Quantity)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &LineItemRecord{Id: id, LineItemData: data}, err
}