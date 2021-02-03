package repository

import (
	"context"

	"github.com/vanilla/go-with-mongodb/api/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var (
	collected string = "products"
	ctx context.Context
)

type ProductRepositoryImpl struct {
	Connection *mongo.Database
}

func NewProductRepositoryImpl(Connection *mongo.Database) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{Connection: Connection}
}

func (r *ProductRepositoryImpl) FindAll(limit  int64, offset int64) ([]entity.Product, error) {
	var (
		products []entity.Product
		product entity.Product
		filterOption = options.Find()
	)

	filterOption.SetLimit(limit)
	filterOption.SetSkip(offset)

	row, err := r.Connection.Collection(collected).Find(ctx, bson.M{}, filterOption)

	if err != nil {
		return nil, err
	}

	for row.Next(ctx) {
		err := row.Decode(&product)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepositoryImpl) CountAllUsers() (int64, error) {
	counter, err := r.Connection.Collection(collected).CountDocuments(ctx, bson.M{}, nil)

	if err != nil {
		return 0, err
	}

	return counter, nil
}

func (r *ProductRepositoryImpl) FindById(id string) (entity.Product, error) {
	var (
		product entity.Product
		productID, _ = primitive.ObjectIDFromHex(id)
	)

	err := r.Connection.Collection(collected).FindOne(ctx, bson.M{"_id": productID}).Decode(&product)

	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (r *ProductRepositoryImpl) Save(productDTO entity.Product) (bool, error) {
	_, err := r.Connection.Collection(collected).InsertOne(ctx, productDTO)

	if err != nil {
		return false, err
	}
	
	return true, nil
}

func (r *ProductRepositoryImpl) Update(productDTO entity.Product, id string) (bool, error) {
	productID, _ := primitive.ObjectIDFromHex(id)

	updateField := bson.M{
		"$set": bson.M{
			"nama_product": productDTO.NamaProduct,
			"quantity": productDTO.Quantity,
			"harga_satuan": productDTO.HargaSatuan,
			"harga_jual": productDTO.HargaJual, 
		},
	}

	_, err := r.Connection.Collection(collected).UpdateOne(ctx, bson.M{"_id": productID}, updateField)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *ProductRepositoryImpl) Delete(id string) (bool, error) {
	productID, _ := primitive.ObjectIDFromHex(id)

	_, err := r.Connection.Collection(collected).DeleteOne(ctx, bson.M{"_id": productID})

	if err != nil {
		return false, err
	}

	return true, nil
}
