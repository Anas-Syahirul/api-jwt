package service

import (
	"api-jwt/models"
	"api-jwt/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductFound(t *testing.T) {
	User := models.User{
		Role: "admin",
		FullName: "Anas Syahirul",
		Email: "try@gmail.com",
		Password: "secret",
	}
	Product := models.Product{
		Title: "Book",
		Description: "It's Book",
		UserID: 2,
		User: &User,
	}
	productRepository.Mock.On("FindById", "3").Return(Product)
	product, err := productService.GetOneProduct("3")
	assert.NotNil(t, product)
	assert.Nil(t, err)
	assert.Equal(t, &Product, product)
}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "3").Return(nil)

	product, err := productService.GetOneProduct("3")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), 
	"error response has to be 'product not found'")
}

func TestProductServiceGetAllProductsNotExisted(t *testing.T) {
	productRepository.Mock.On("FindAll").Return(nil)

	products, err := productService.GetAllProducts()

	assert.Nil(t, products)
	assert.NotNil(t, err)
	assert.Equal(t, "there's no product in db", err.Error(), 
	"error response has to be 'there's no product in db'")
}

func TestProductServiceGetAllProductsExisted(t *testing.T) {
	User := models.User{
		Role: "admin",
		FullName: "Anas Syahirul",
		Email: "try@gmail.com",
		Password: "secret",
	}
	Product1 := models.Product{
		Title: "Book",
		Description: "It's Book",
		UserID: 2,
		User: &User,
	}
	Product2 := models.Product{
		Title: "Chair",
		Description: "It's Chair",
		UserID: 2,
		User: &User,
	}
	prods := &[]models.Product{Product1, Product2}
	productRepository.Mock.On("FindAll").Return(prods)
	products, err := productService.GetAllProducts()

	assert.NotNil(t, products)
	assert.Nil(t, err)
	assert.Equal(t, prods, products)
}