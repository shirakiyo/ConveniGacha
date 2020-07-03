package usecase

import (
	"fmt"
	"math/rand"

	"github.com/shirakiyo/FamimaGacha/internal/domain/model"
	"github.com/shirakiyo/FamimaGacha/internal/domain/repository"
)

type ProductCategory string

func (pc ProductCategory) String() string {
	return string(pc)
}

const (
	FoodsPrefix  ProductCategory = "foods"
	SweetsPrefix ProductCategory = "sweets"
	SnacksPrefix ProductCategory = "snacks"
)

type ProductUseCase interface {
	GetProduct(ProductCategory) (*model.Product, error)
}

type productUseCase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(pr repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		productRepository: pr,
	}
}

// GetProduct 全商品の中からランダムで商品を選択する
func (pu *productUseCase) GetProduct(category ProductCategory) (*model.Product, error) {
	fileName := category.String()
	if category == "" {
		fileName = "products"
	}

	products, err := pu.productRepository.ListProducts(fmt.Sprintf("%s.csv", fileName))
	if err != nil {
		return nil, err
	}
	index := rand.Int() % len(products)

	return products[index], nil
}
