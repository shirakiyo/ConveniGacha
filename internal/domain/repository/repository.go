package repository

import (
	"errors"

	"github.com/shirakiyo/FamimaGacha/internal/domain/model"
)

type ProductRepository interface {
	ListProducts() ([]*model.Product, error)
}

type productsFile struct {
	filePath string
}

func NewProductRepository(path string) ProductRepository {
	return &productsFile{
		filePath: path,
	}
}

func (pr *productsFile) ListProducts() ([]*model.Product, error) {
	return nil, errors.New("")
}
