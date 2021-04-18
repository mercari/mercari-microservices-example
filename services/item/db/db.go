package db

import (
	"context"
	"errors"
	"sync"

	"github.com/mercari/go-conference-2021-spring-office-hour/services/item/model"
)

var _ DB = (*db)(nil)

var ErrNotFound = errors.New("not found")

type DB interface {
	GetItem(ctx context.Context, id string) (*model.Item, error)
}

func New() DB {
	return &db{
		db: map[string]*model.Item{
			"e0e58243-4138-48e5-8aba-448a8888e2ff": {
				CustomerID: "7c0cde05-4df0-47f4-94c4-978dd9f56e5c", // gopher
				Title:      "Mobile Phone",
				Price:      10000,
			},
			"0b185d96-d6fa-4eaf-97f6-3f6d2c1649b6": {
				CustomerID: "7c0cde05-4df0-47f4-94c4-978dd9f56e5c", // gopher
				Title:      "Laptop",
				Price:      20000,
			},
		},
	}
}

type db struct {
	db   map[string]*model.Item
	dbMu sync.RWMutex
}

func (d *db) GetItem(ctx context.Context, id string) (*model.Item, error) {
	d.dbMu.RLock()
	defer d.dbMu.RUnlock()

	i, ok := d.db[id]
	if !ok {
		return nil, ErrNotFound
	}
	i.ID = id

	return i, nil
}
