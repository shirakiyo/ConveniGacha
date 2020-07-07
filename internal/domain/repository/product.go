package repository

import (
	"github.com/shirakiyo/ConveniGacha/internal/domain/model"
)

type ProductRepository interface {
	ListProducts(string) ([]*model.Product, error)
}
