package usecase

import (
	"math/rand"

	"github.com/labstack/echo/v4"
	"github.com/shirakiyo/FamimaGacha/internal/domain/model"
	"github.com/shirakiyo/FamimaGacha/internal/domain/repository"
)

type ProductUseCase interface {
	GetProduct(echo.Context) (*model.Product, error)
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
func (pu *productUseCase) GetProduct(c echo.Context) (*model.Product, error) {
	products, err := pu.productRepository.ListProducts()
	if err != nil {
		return nil, err
	}
	index := rand.Int() % len(products)

	return products[index], nil
}
