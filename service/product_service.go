package service

import (
	"api-jwt/models"
	"api-jwt/repository"
	"errors"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*models.Product, error) {
	product := service.Repository.FindById(id)

	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (service ProductService) GetAllProducts() (*[]models.Product, error) {
	products := service.Repository.FindAll()

	if products == nil {
		return nil, errors.New("there's no product in db")
	}

	return products, nil
}