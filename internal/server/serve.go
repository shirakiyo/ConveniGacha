package server

import (
	"net/http"

	"github.com/shirakiyo/FamimaGacha/internal/domain/repository"
	"github.com/shirakiyo/FamimaGacha/internal/usecase"

	"github.com/shirakiyo/FamimaGacha/internal/app/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve(port string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		Skipper: func(c echo.Context) bool {
			return c.Request().Method == echo.OPTIONS
		},
	}))
	e.HidePort = true
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	repo := repository.NewProductRepository("../../data/produce.csv")
	usecase := usecase.NewProductUseCase(repo)
	productHandler := handler.NewHandler(usecase)

	e.GET("/get", productHandler.GetProduct)

	e.Logger.Fatal(e.Start(port))
}
