package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

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

	// 相対パスを絶対パスに変換
	fp, err := filepath.Abs("./data/product.csv")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	repo := repository.NewProductRepository(fp)
	usecase := usecase.NewProductUseCase(repo)
	productHandler := handler.NewHandler(usecase)

	e.GET("/get", productHandler.GetProduct)

	e.Logger.Fatal(e.Start(port))
}
