package internal

import (
	"database/sql"
	"fmt"
)

const (
	CustomerTable  = "customers"
	CustomerSchema = `
		CREATE IF NOT EXISTS TABLE customers (
			id INT NOT NULL UNIQUE PRIMARY KEY,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL
		);
	`
)

type CustomerStore struct {
	Db    *sql.DB
	stmts map[string]sql.Stmt
}

type CustomerId = int64

func NewCustomerStore(db *sql.DB) (*CustomerStore, error) {
	_, err := db.Exec(CustomerSchema)
	if err != nil {
		return nil, err
	}
	statements := map[string]string{
		"createCustomer": fmt.Sprintf("INSERT INTO %s VALUES(NULL, ?, ?);", CustomerTable),
	}

	store := CustomerStore{Db: db, stmts: make(map[string]sql.Stmt)}

	for k, v := range statements {
		sqlStmt, err := db.Prepare(v)
		if err != nil {
			return nil, err
		}
		store.stmts[k] = *sqlStmt
	}

	return &store, nil
}

func (store *CustomerStore) CreateCustomer(data CustomerData) (*int64, error) {
	query := store.stmts["createCustomer"]
	res, err := query.Exec(data.FirstName, data.LastName)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (store *CustomerStore) GetCustomer(id CustomerId) (*CustomerData, error) {
	query := store.stmts["getCustomer"]
	var (
		skipId int64
		data   CustomerData
	)
	err := query.
		QueryRow(id).
		Scan(&skipId, &data.FirstName, &data.LastName)
	return &data, err
}
