package store

import "clere-coding-challenge-api/store/model"

var store Store

type Store interface {
	Product() ProductStore
}

type ProductStore interface {
	Get(id int) (model.Product, error)
	GetAll() ([]model.Product, error)
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(id int) error
}
