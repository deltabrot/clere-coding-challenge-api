package memstore

import (
	"clere-coding-challenge-api/store"
	"clere-coding-challenge-api/store/model"
	"errors"
)

type MemProductStore struct {
	*MemStore
}

func (s *MemStore) Product() store.ProductStore {
	return &MemProductStore{s}
}

func (s *MemProductStore) Get(id int) (model.Product, error) {
	for _, p := range s.products {
		if p.Id == id {
			return p, nil
		}
	}

	return model.Product{}, errors.New("Error: Failed to get product")
}

func (s *MemProductStore) GetAll() ([]model.Product, error) {
	return s.products, nil
}

func (s *MemProductStore) Create(product *model.Product) error {
	largestId := 0
	for _, p := range s.products {
		if p.Id > largestId {
			largestId = p.Id
		}
	}

	product.Id = largestId + 1
	s.products = append(s.products, *product)

	return nil
}

func (s *MemProductStore) Update(product *model.Product) error {
	for i, p := range s.products {
		if p.Id == product.Id {
			s.products[i] = *product
			return nil
		}
	}

	return errors.New("Error: Failed to update product")
}

func (s *MemProductStore) Delete(id int) error {
	for i, p := range s.products {
		if p.Id == id {
			s.products = append(s.products[:i], s.products[i+1:]...)
			return nil
		}
	}

	return errors.New("Error: Failed to delete product")
}
