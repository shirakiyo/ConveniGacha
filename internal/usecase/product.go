package usecase

import (
	"errors"

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

func (pu *productUseCase) GetProduct(c echo.Context) (*model.Product, error) {
	return nil, errors.New("")
}
