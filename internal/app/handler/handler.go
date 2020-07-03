package handler

import (
	"log"
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
	product, err := h.productUseCase.GetProduct(c)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, product)
}
