package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/shirakiyo/ConveniGacha/internal/usecase"

	conveniRepo "github.com/shirakiyo/ConveniScraping/pkg/repository"
	conveniService "github.com/shirakiyo/ConveniScraping/pkg/service"
)

var (
	conveni = flag.String("conveni", "", "input convenience store name")
)

func main() {
	flag.Parse()

	switch *conveni {
	case "famima":
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
	case "lawson":
		lawsonRepo := conveniRepo.NewLawsonClient()
		lawsonSvc := conveniService.NewLawsonService(lawsonRepo)
		dataPath, err := filepath.Abs("data")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		lawsonSvc.Sweets(filepath.Join(dataPath, usecase.LawsonSweetsCSV))
		log.Println("sweets ok")
		lawsonSvc.Snacks(filepath.Join(dataPath, usecase.LawsonSnacksCSV))
		log.Println("snacks ok")
		lawsonSvc.Foods(filepath.Join(dataPath, usecase.LawsonFoodsCSV))
		log.Println("foods ok")
		lawsonSvc.Products(filepath.Join(dataPath, usecase.LawsonProductsCSV))
		log.Println("products ok")
	}
}
