package repository

import "github.com/vanilla/go-with-mongodb/api/entity"

type ProductRepository interface {
	FindAll() ([]entity.Product, error)
	FindById(string) (entity.Product, error)
	Save(entity.Product) (bool, error)
	Update(entity.Product, string) (bool, error)
	Delete(string) (bool, error)
}