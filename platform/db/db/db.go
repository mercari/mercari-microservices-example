package db

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/mercari/mercari-microservices-example/platform/db/model"
)

var _ DB = (*db)(nil)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

var customerGoldie = &model.Customer{
	ID:   "7c0cde05-4df0-47f4-94c4-978dd9f56e5c",
	Name: "goldie",
}

type DB interface {
	CreateCustomer(ctx context.Context, name string) (*model.Customer, error)
	GetCustomer(ctx context.Context, id string) (*model.Customer, error)
	GetCustomerByName(ctx context.Context, name string) (*model.Customer, error)
	CreateItem(ctx context.Context, item *model.Item) (*model.Item, error)
	GetItem(ctx context.Context, id string) (*model.Item, error)
	GetAllItems(ctx context.Context) ([]*model.Item, error)
}

func New() DB {
	return &db{
		customersByID: map[string]*model.Customer{
			customerGoldie.ID: customerGoldie,
		},
		customersByName: map[string]*model.Customer{
			customerGoldie.Name: customerGoldie,
		},
		items: map[string]*model.Item{
			"e0e58243-4138-48e5-8aba-448a8888e2ff": {
				ID:         "e0e58243-4138-48e5-8aba-448a8888e2ff",
				CustomerID: customerGoldie.ID,
				Title:      "Mobile Phone",
				Price:      10000,
			},
			"0b185d96-d6fa-4eaf-97f6-3f6d2c1649b6": {
				ID:         "0b185d96-d6fa-4eaf-97f6-3f6d2c1649b6",
				CustomerID: customerGoldie.ID,
				Title:      "Laptop",
				Price:      20000,
			},
		},
	}
}

type db struct {
	customersByID   map[string]*model.Customer
	customersByIDMu sync.RWMutex

	customersByName   map[string]*model.Customer
	customersByNameMu sync.RWMutex

	items   map[string]*model.Item
	itemsMu sync.RWMutex
}

func (d *db) CreateCustomer(ctx context.Context, name string) (*model.Customer, error) {
	d.customersByNameMu.Lock()
	defer d.customersByNameMu.Unlock()

	_, ok := d.customersByName[name]
	if ok {
		return nil, ErrAlreadyExists
	}

	c := &model.Customer{
		ID:   uuid.New().String(),
		Name: name,
	}

	d.customersByName[c.Name] = c

	d.customersByIDMu.Lock()
	d.customersByID[c.ID] = c
	d.customersByIDMu.Unlock()

	return c, nil
}

func (d *db) GetCustomer(ctx context.Context, id string) (*model.Customer, error) {
	d.customersByIDMu.RLock()
	defer d.customersByIDMu.RUnlock()

	c, ok := d.customersByID[id]
	if !ok {
		return nil, ErrNotFound
	}

	return c, nil
}

func (d *db) GetCustomerByName(ctx context.Context, name string) (*model.Customer, error) {
	d.customersByNameMu.RLock()
	defer d.customersByNameMu.RUnlock()

	c, ok := d.customersByName[name]
	if !ok {
		return nil, ErrNotFound
	}

	return c, nil
}

func (d *db) CreateItem(ctx context.Context, item *model.Item) (*model.Item, error) {
	id := uuid.New().String()
	d.itemsMu.Lock()
	d.items[id] = item
	d.itemsMu.Unlock()

	item.ID = id

	return item, nil
}

func (d *db) GetItem(ctx context.Context, id string) (*model.Item, error) {
	d.itemsMu.RLock()
	defer d.itemsMu.RUnlock()

	i, ok := d.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	i.ID = id

	return i, nil
}

func (d *db) GetAllItems(ctx context.Context) ([]*model.Item, error) {
	var items []*model.Item

	for _, item := range d.items {
		items = append(items, item)
	}

	return items, nil
}
