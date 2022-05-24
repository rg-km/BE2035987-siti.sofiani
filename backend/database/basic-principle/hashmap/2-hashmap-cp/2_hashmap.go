package main

import "fmt"

type PrimaryKey int

type InvoiceRow struct {
	ID               PrimaryKey
	SubscriptionCode string
	CustomerName     string
	Address          string
	Phone            string
}

type InvoiceDB struct {
	m map[PrimaryKey]InvoiceRow
}

func NewInvoice() *InvoiceDB {
	return &InvoiceDB{
		m: make(map[PrimaryKey]InvoiceRow),
	}
}

func (db *InvoiceDB) Insert(code string, name string, address string, phone string) {
	db.m[PrimaryKey(len(db.m)+1)] = InvoiceRow{
		ID:               PrimaryKey(len(db.m)) + 1,
		SubscriptionCode: code,
		CustomerName:     name,
		Address:          address,
		Phone:            phone,
	}
}

func (db *InvoiceDB) Where(id PrimaryKey) *InvoiceRow {
<<<<<<< HEAD
	result, ok := db.m[id]
	if !ok {
		return nil
	}
	return &result
}

func (db *InvoiceDB) Update(id PrimaryKey, code string, name string, address string, phone string) (*InvoiceRow, error) {
	invoice, ok := db.m[id]
	if !ok {
		return nil, fmt.Errorf("id %d not found", id)
	}

	invoice.SubscriptionCode = code
	invoice.CustomerName = name
	invoice.Address = address
	invoice.Phone = phone

	return &invoice, nil
}
=======
	return InvoiceRow{} // TODO: replace this
}

func (db *InvoiceDB) Update(id PrimaryKey, code string, name string, address string, phone string) (*InvoiceRow, error) {
	return nil, nil // TODO: replace this
}
>>>>>>> fc12154791502702980a046e3507ab317e48f675
