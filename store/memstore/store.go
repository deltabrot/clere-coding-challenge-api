package memstore

import (
	"clere-coding-challenge-api/store"
	"clere-coding-challenge-api/store/model"
)

var _ store.Store = &MemStore{}

type MemStore struct {
	products []model.Product
}

func Open() (*MemStore, error) {
	var s MemStore
	s.products = []model.Product{}

	return &s, nil
}
