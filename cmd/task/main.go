package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/shirakiyo/ConveniGacha/internal/usecase"

	conveniRepo "github.com/shirakiyo/ConveniScraping/pkg/repository"
	conveniService "github.com/shirakiyo/ConveniScraping/pkg/service"
)

func main() {
	famimaRepo := conveniRepo.NewFamimaClient()
	famimaSvc := conveniService.NewFamimaService(famimaRepo)
	dataPath, err := filepath.Abs("data")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	famimaSvc.Products(filepath.Join(dataPath, usecase.FamimaProductsCSV))
	famimaSvc.Foods(filepath.Join(dataPath, usecase.FamimaFoodsCSV))
	famimaSvc.Sweets(filepath.Join(dataPath, usecase.FamimaSweetsCSV))
	famimaSvc.Snacks(filepath.Join(dataPath, usecase.FamimaSnacksCSV))
}
