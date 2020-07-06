package handler

import (
	"net/http"

	"github.com/shirakiyo/ConveniGacha/internal/app/output"

	"github.com/shirakiyo/ConveniGacha/internal/app/input"

	"github.com/labstack/echo/v4"
	"github.com/shirakiyo/ConveniGacha/internal/usecase"
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
	var param input.GetProduct
	if err := c.Bind(&param); err != nil {
		return APIResponseError(c, http.StatusBadRequest, "Bad Request", err)
	}

	product, err := h.productUseCase.GetProduct(usecase.ProductCategory(param.Category))
	if err != nil {
		return APIResponseError(c, http.StatusInternalServerError, "Internal Server Error", err)
	}

	return APIResponseOK(c, output.Product{
		Name:   product.Name,
		Price:  product.Price,
		Link:   product.Link,
		Detail: product.Detail,
	})
}
