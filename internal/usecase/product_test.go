package usecase

import (
	"errors"
	"testing"

	"github.com/shirakiyo/ConveniGacha/internal/domain/model"

	"github.com/golang/mock/gomock"
	mock_repository "github.com/shirakiyo/ConveniGacha/internal/domain/repository/mock"
)

type ProductTestCase struct {
	name  string
	run   func(*testing.T, ProductUseCase) error
	mock  func(repository *mock_repository.MockProductRepository)
	check func(*testing.T, error)
}

type ProductTestCases []ProductTestCase

func (cs ProductTestCases) Run(t *testing.T) {
	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock_repository.NewMockProductRepository(ctrl)
			if c.mock != nil {
				c.mock(mock)
			}

			err := c.run(t, NewProductUseCase(mock))

			c.check(t, err)
		})
	}
}

func TestProductUseCase_GetProduct(t *testing.T) {
	expectedProducts := []*model.Product{
		{
			Name:   "foo",
			Price:  100,
			Link:   "bar",
			Detail: "baz",
		},
	}
	testCases := ProductTestCases{
		{
			name: "category is nil OK",
			run: func(t *testing.T, u ProductUseCase) error {
				product, err := u.GetProduct("")

				if product == nil {
					err = errors.New("product is nil")
				}
				return err
			},
			mock: func(mock *mock_repository.MockProductRepository) {
				mock.EXPECT().
					ListProducts(FamimaProductsCSV).
					Return(expectedProducts, nil)
			},
			check: func(t *testing.T, err error) {
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "category is foods OK",
			run: func(t *testing.T, u ProductUseCase) error {
				product, err := u.GetProduct(FoodsPrefix)

				if product == nil {
					err = errors.New("product is nil")
				}
				return err
			},
			mock: func(mock *mock_repository.MockProductRepository) {
				mock.EXPECT().
					ListProducts(FamimaFoodsCSV).
					Return(expectedProducts, nil)
			},
			check: func(t *testing.T, err error) {
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "category is sweets OK",
			run: func(t *testing.T, u ProductUseCase) error {
				product, err := u.GetProduct(SweetsPrefix)

				if product == nil {
					err = errors.New("product is nil")
				}
				return err
			},
			mock: func(mock *mock_repository.MockProductRepository) {
				mock.EXPECT().
					ListProducts(FamimaSweetsCSV).
					Return(expectedProducts, nil)
			},
			check: func(t *testing.T, err error) {
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "category is foods OK",
			run: func(t *testing.T, u ProductUseCase) error {
				product, err := u.GetProduct(SnacksPrefix)

				if product == nil {
					err = errors.New("product is nil")
				}
				return err
			},
			mock: func(mock *mock_repository.MockProductRepository) {
				mock.EXPECT().
					ListProducts(FamimaSnacksCSV).
					Return(expectedProducts, nil)
			},
			check: func(t *testing.T, err error) {
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "category is foods OK",
			run: func(t *testing.T, u ProductUseCase) error {
				product, err := u.GetProduct(SnacksPrefix)

				if product == nil {
					err = errors.New("product is nil")
				}
				return err
			},
			mock: func(mock *mock_repository.MockProductRepository) {
				mock.EXPECT().
					ListProducts(FamimaSnacksCSV).
					Return(expectedProducts, nil)
			},
			check: func(t *testing.T, err error) {
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "invalid category",
			run: func(t *testing.T, u ProductUseCase) error {
				_, err := u.GetProduct("hoge")
				return err
			},
			mock: func(mock *mock_repository.MockProductRepository) {
			},
			check: func(t *testing.T, err error) {
				if !errors.Is(err, ErrInvalidCategory) {
					t.Errorf("wrong error: %v", err)
				}
			},
		},
	}

	testCases.Run(t)
}
