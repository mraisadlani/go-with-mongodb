package repository

import "github.com/vanilla/go-with-mongodb/api/entity"

type ProductRepository interface {
	FindAll(int64, int64) ([]entity.Product, error)
	CountAllUsers() (int64, error)
	FindById(string) (entity.Product, error)
	Save(*entity.Product) (bool, error)
	// Update(entity.Product, string) (bool, error)
	// Delete(string) (bool, error)
}