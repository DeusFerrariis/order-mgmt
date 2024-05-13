package internal

import (
	"database/sql"
	"errors"
	"sync"

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

const (
	// id | sku | description | cost | price | line_number | quantity | order_id
	LineItemTableSchema = `
		CREATE TABLE IF NOT EXISTS lineItems (
			id INTEGER UNIQUE NOT NULL PRIMARY KEY,
			sku TEXT NOT NULL,
			description TEXT NOT NULL,
			cost INTEGER NOT NULL,
			price INTEGER NOT NULL,
			line_number INTEGER NOT NULL,
			quantity INTEGER NOT NULL,
			order_id INTEGER NOT NULL
		);
	`
)

func (store *LineItemStore) PrepareStatements() error {
	store.stmts = make(map[string]*sql.Stmt)
	stmt, err := store.db.Prepare(`
		INSERT INTO lineItems VALUES(NULL, ?, ?, ?, ?, ?, ?, ?);
	`)
	store.stmts["createLineItem"] = stmt
	stmt, err = store.db.Prepare(`
		SELECT * FROM lineItems WHERE id=?;
	`)
	store.stmts["getLineItemById"] = stmt
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
	_, err = db.Exec(LineItemSchema)
	if err != nil {
		return LineItemStore{}, err
	}
	store := LineItemStore{
		db: db,
	}
	if err = store.PrepareStatements(); err != nil {
		return LineItemStore{}, errors.Join(errors.New("Could not prepare statements:"), err)
	}
	return store, nil
}

func (ctx *LineItemStore) CreateLineItem(data LineItemData) (*LineItemRecord, error) {
	stmt := ctx.stmts["createLineItem"]
	// id | sku | description | cost | price | line_number | quantity | order_id
	res, err := stmt.Exec(
		data.Sku, data.Description, data.Cost, data.Price, data.LineNumber,
		data.Quantity, data.OrderId,
	)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &LineItemRecord{Id: id, LineItemData: data}, err
}

func (ctx *LineItemStore) GetLineItemById(id int64) (*LineItemRecord, error) {
	var r LineItemRecord
	stmt := ctx.stmts["getLineItemById"].QueryRow(id)
	err := stmt.
		// id | sku | description | cost | price | line_number | quantity | order_id
		Scan(&r.Id, &r.Sku, &r.Description, &r.Cost, &r.Price, &r.LineNumber, &r.Quantity, &r.OrderId)
	return &r, err
}
