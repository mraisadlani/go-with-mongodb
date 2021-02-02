package repository

import (
	"github.com/vanilla/go-with-mongodb/api/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


type ProductRepositoryImpl struct {
	db *mgo.Database
	collection string
}

func NewProductRepositoryImpl(db *mgo.Database, collection string) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db: db,
		collection: collection,
	}
}

func (r *ProductRepositoryImpl) FindAll() ([]entity.Product, error) {
	var products []entity.Product

	err := r.db.C(r.collection).Find(bson.M{}).All(&products)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepositoryImpl) FindById(id string) (entity.Product, error) {
	return entity.Product{}, nil
}

func (r *ProductRepositoryImpl) Save(productDTO entity.Product) (bool, error) {
	err := r.db.C(r.collection).Insert(productDTO)

	if err != nil {
		return false, err
	}
	
	return true, nil
}

func (r *ProductRepositoryImpl) Update(productDTO entity.Product, id string) (bool, error) {
	return true, nil
}

func (r *ProductRepositoryImpl) Delete(id string) (bool, error) {
	return true, nil
}
