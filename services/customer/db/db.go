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
	GetCustomer(ctx context.Context, id string) (*model.Customer, error)
	GetCustomerByName(ctx context.Context, name string) (*model.Customer, error)
}

var customerGopher = &model.Customer{
	ID:   "7c0cde05-4df0-47f4-94c4-978dd9f56e5c",
	Name: "gopher",
}

func New() DB {
	return &db{
		mapID: map[string]*model.Customer{
			customerGopher.ID: customerGopher,
		},
		mapName: map[string]*model.Customer{
			customerGopher.Name: customerGopher,
		},
	}
}

type db struct {
	mapID   map[string]*model.Customer
	mapIDMu sync.RWMutex

	mapName   map[string]*model.Customer
	mapNameMu sync.RWMutex
}

func (d *db) GetCustomer(ctx context.Context, id string) (*model.Customer, error) {
	d.mapIDMu.RLock()
	defer d.mapIDMu.RUnlock()

	c, ok := d.mapID[id]
	if !ok {
		return nil, ErrNotFound
	}

	return c, nil
}

func (d *db) GetCustomerByName(ctx context.Context, name string) (*model.Customer, error) {
	d.mapNameMu.RLock()
	defer d.mapNameMu.RUnlock()

	c, ok := d.mapName[name]
	if !ok {
		return nil, ErrNotFound
	}

	return c, nil
}
