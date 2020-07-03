package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shirakiyo/FamimaGacha/internal/usecase"
)

type ProductHandler interface {
	GetProduct(echo.Context) error
}

type productHandler struct {
	productUseCase usecase.ProductUseCase
}

func NewHandler(pu usecase.ProductUseCase) ProductHandler {
	return &productHandler{
		productUseCase: pu,
	}
}

func (h *productHandler) GetProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world!")
}
