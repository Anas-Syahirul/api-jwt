package repository

import "api-jwt/models"

type ProductRepository interface {
	FindById(id string) *models.Product
	FindAll() *[]models.Product
}