package product

import (
	"fmt"
	"log"
	"strconv"
)

type Serviced interface {
	Describe(ID uint64) (*Product, error)
	List(cursor uint64, limit uint64) error
	Create(args string) (uint64, error)
	Update(ID uint64, args string) error
	Remove(ID uint64) (bool, error)
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return AllProducts
}

func (s *Service) Describe(idx int) (*Product, error) {
	return &AllProducts[idx-1], nil
}

func (s *Service) Create(args string) (uint64, error) {
	err := "Такой объект уже существует!"
	var product Product
	var newProductId = len(AllProducts) + 1
	for _, elem := range AllProducts {
		if elem.Title == args {
			return 0, fmt.Errorf("%v", err)
		}
	}
	product.ID = newProductId
	product.Title = args
	AllProducts = append(AllProducts, product)
	return uint64(product.ID), nil
}

func (s *Service) Remove(ID uint64) bool {
	object, found := FindById(ID)
	if !found {
		return found
	}
	for idx, elem := range AllProducts {
		if elem == *object {
			AllProducts = append(AllProducts[:idx], AllProducts[idx+1:]...)
		}
	}
	return found
}

func (s *Service) Update(ID uint64, args string) error {
	product, found := FindById(ID)
	if !found {
		return fmt.Errorf("Нет объекта под таким id: %v", found)
	}
	product.Title = args
	for idx, elem := range AllProducts {
		if elem.ID == int(ID) {
			AllProducts[idx] = *product
		}
	}
	return nil
}

func FindById(ID uint64) (*Product, bool) {
	var product Product
	var found bool
	for idx, elem := range AllProducts {
		if idx == int(ID-1) {
			found = true
			product = elem
			break
		}
	}
	return &product, found
}

func Translate(args string) (int, error) {
	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return 0, err
	}
	return arg, nil
}
