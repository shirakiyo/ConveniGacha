package repository

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/shirakiyo/FamimaGacha/internal/domain/model"
)

type ProductRepository interface {
	ListProducts() ([]*model.Product, error)
}

type productsFile struct {
	filePath string
}

func NewProductRepository(path string) ProductRepository {
	return &productsFile{
		filePath: path,
	}
}

func (pr *productsFile) ListProducts() ([]*model.Product, error) {
	result := make([]*model.Product, 0)

	csvFile, err := os.Open(pr.filePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	var line []string

	for i := 1; ; i++ {
		line, err = reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if len(line) != 3 {
			return nil, fmt.Errorf("CSVファイルの%d行目に誤りがあります", i)
		}

		// "1,000"という数値が来たときの処理
		if strings.Contains(line[1], ",") {
			line[1] = strings.Replace(line[1], `"`, "", -1)
			line[1] = strings.Replace(line[1], ",", "", -1)
		}

		price, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			return nil, fmt.Errorf("CSVファイルの%d行目に誤りがあります: %w", i, err)
		}

		content := &model.Product{
			Name:  line[0],
			Price: price,
			Link:  line[2],
		}

		result = append(result, content)
	}

	return result, nil
}
