package db

import (
	"context"
	"errors"
	"sync"

	"github.com/mercari/go-conference-2021-spring-office-hour/services/customer/model"
)

var _ DB = (*db)(nil)

var ErrNotFound = errors.New("not found")

type DB interface {
	GetCustomer(ctx context.Context, name string) (*model.Customer, error)
}

func New() DB {
	return &db{
		db: map[string]*model.Customer{
			"gopher": {
				ID:   "7c0cde05-4df0-47f4-94c4-978dd9f56e5c",
				Name: "gopher",
			},
		},
	}
}

type db struct {
	db   map[string]*model.Customer
	dbMu sync.RWMutex
}

func (d *db) GetCustomer(ctx context.Context, name string) (*model.Customer, error) {
	d.dbMu.RLock()
	defer d.dbMu.RUnlock()

	c, ok := d.db[name]
	if !ok {
		return nil, ErrNotFound
	}

	return c, nil
}
