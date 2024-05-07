package internal

import "database/sql"

type CustomerStore struct {
	db *sql.DB
}

type CustomerId = int64

func (store *CustomerStore) CreateCustomer(data CustomerData) (*CustomerId, error) {
	// TODO: implement
	return nil, nil
}

func (store *CustomerStore) GetCustomer(id CustomerId) (*CustomerData, error) {
	// TODO: implement
	return nil, nil
}
