package usecase

import (
	"math/rand"

	"github.com/shirakiyo/ConveniGacha/internal/domain/model"
	"github.com/shirakiyo/ConveniGacha/internal/domain/repository"
)

type ProductCategory string

func (pc ProductCategory) String() string {
	return string(pc)
}

const (
	FoodsPrefix  ProductCategory = "foods"
	SweetsPrefix ProductCategory = "sweets"
	SnacksPrefix ProductCategory = "snacks"

	FamimaProductsCSV = "famima_products.csv"
	FamimaFoodsCSV    = "famima_foods.csv"
	FamimaSweetsCSV   = "famima_sweets.csv"
	FamimaSnacksCSV   = "famima_snacks.csv"
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
	switch category {
	case "":
		fileName = FamimaProductsCSV
	case FoodsPrefix:
		fileName = FamimaFoodsCSV
	case SweetsPrefix:
		fileName = FamimaSweetsCSV
	case SnacksPrefix:
		fileName = FamimaSnacksCSV
	}

	products, err := pu.productRepository.ListProducts(fileName)
	if err != nil {
		return nil, err
	}
	index := rand.Int() % len(products)

	return products[index], nil
}
