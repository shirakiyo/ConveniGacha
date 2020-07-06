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

	famimaSvc.Sweets(filepath.Join(dataPath, usecase.FamimaSweetsCSV))
	log.Println("sweets ok")
	famimaSvc.Snacks(filepath.Join(dataPath, usecase.FamimaSnacksCSV))
	log.Println("snacks ok")
	famimaSvc.Foods(filepath.Join(dataPath, usecase.FamimaFoodsCSV))
	log.Println("foods ok")
	famimaSvc.Products(filepath.Join(dataPath, usecase.FamimaProductsCSV))
	log.Println("products ok")
}
