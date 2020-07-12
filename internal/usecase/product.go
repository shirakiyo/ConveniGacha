package usecase

import (
	"errors"
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

	LawsonProductsCSV = "lawson_products.csv"
	LawsonFoodsCSV    = "lawson_foods.csv"
	LawsonSweetsCSV   = "lawson_sweets.csv"
	LawsonSnacksCSV   = "lawson_snacks.csv"
)

var (
	ErrInvalidCategory = errors.New("invalid category")
)

type ProductUseCase interface {
	GetFamimaProduct(ProductCategory) (*model.Product, error)
	GetLawsonProduct(ProductCategory) (*model.Product, error)
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
func (pu *productUseCase) GetFamimaProduct(category ProductCategory) (*model.Product, error) {
	var fileName string

	switch category {
	case "":
		fileName = FamimaProductsCSV
	case FoodsPrefix:
		fileName = FamimaFoodsCSV
	case SweetsPrefix:
		fileName = FamimaSweetsCSV
	case SnacksPrefix:
		fileName = FamimaSnacksCSV
	default:
		return nil, ErrInvalidCategory
	}

	products, err := pu.productRepository.ListProducts(fileName)
	if err != nil {
		return nil, err
	}
	index := rand.Int() % len(products)

	return products[index], nil
}

func (pu *productUseCase) GetLawsonProduct(category ProductCategory) (*model.Product, error) {
	var fileName string

	switch category {
	case "":
		fileName = LawsonProductsCSV
	case FoodsPrefix:
		fileName = LawsonFoodsCSV
	case SweetsPrefix:
		fileName = LawsonSweetsCSV
	case SnacksPrefix:
		fileName = LawsonSnacksCSV
	default:
		return nil, ErrInvalidCategory
	}

	products, err := pu.productRepository.ListProducts(fileName)
	if err != nil {
		return nil, err
	}
	index := rand.Int() % len(products)

	return products[index], nil
}
