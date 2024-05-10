package internal

import (
	"database/sql"
	"sync"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

const (
	CustomerTable  = "customers"
	CustomerSchema = `
		CREATE TABLE IF NOT EXISTS customers (
			id INT NOT NULL UNIQUE PRIMARY KEY,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL
		);
	`
	CreateCustomerSql = `
		SELECT * FROM customers WHERE id=?;
	`
)

type CustomerStoreContext struct {
	echo.Context
	mu sync.Mutex
	db *sql.DB
}

func (sc *CustomerStoreContext) GetCustomerStmt() (*sql.Stmt, error) {
	return sc.db.Prepare(CreateCustomerSql)
}

func NewCustomerStoreContext(path string, c echo.Context) (*CustomerStoreContext, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(CustomerSchema)
	if err != nil {
		return nil, err
	}
	return &CustomerStoreContext{
		Context: c,
		db:      db,
	}, nil
}

func (store *CustomerStoreContext) CreateCustomer(data CustomerData) (*int64, error) {
	return nil, nil
}

func (store *CustomerStoreContext) GetCustomer(id int64) (*CustomerRecord, error) {
	var rec CustomerRecord
	stmt, err := store.GetCustomerStmt()
	if err != nil {
		return nil, nil
	}
	err = stmt.
		QueryRow(id).
		Scan(&rec.Id, &rec.FirstName, &rec.LastName)
	return &rec, err
}
