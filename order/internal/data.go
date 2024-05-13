package internal

import (
	"database/sql"
	"sync"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type (
	OrderStore struct {
		mu sync.Mutex
		db *sql.DB
	}
	LineItemStore struct {
		mu    sync.Mutex
		db    *sql.DB
		stmts map[string]*sql.Stmt
	}
)

func (store *LineItemStore) PrepareStatements() error {
	stmt, err := store.db.Prepare(`
		INSERT INTO lineItems VALUES(NULL, ?, ?, ?, ?, ?, ?, ?);
	`)
	store.stmts["createLineItem"] = stmt
	return err
}

func NewOrderStore(path string) (*OrderStore, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	// TODO: init schema for orders table
	return &OrderStore{
		db: db,
	}, nil
}

func NewLineItemStore(path string) (LineItemStore, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return LineItemStore{}, err
	}
	// TODO: init schema for orders table
	return LineItemStore{
		db: db,
	}, nil
}

func (ctx *LineItemStore) CreateLineItemStmt() (*sql.Stmt, error) {
	return ctx.db.Prepare(`
		INSERT INTO lineItems VALUES(NULL, ?, ?, ?, ?, ?, ?, ?);
	`)
}

func (ctx *LineItemStore) CreateLineItem(data LineItemData) (*LineItemRecord, error) {
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
