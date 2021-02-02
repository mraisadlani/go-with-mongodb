package entity

type Product struct {
	ID int64 `bson:"_id" json:"_id"`
	NamaProduct string `bson:"nama_product" json:"nama_product"`
	Quantity int64 `bson:"quantity" json:"quantity"`
	HargaSatuan int64 `bson:"harga_satuan" json:"harga_satuan"`
	HargaJual int64 `bson:"harga_jual" json:"harga_jual"`
	CreateBy string `bson:"create_by" json:"create_by"`
	CreateDate string `bson:"create_date" json:"create_date"`
	UpdateBy string `bson:"update_by" json:"update_by"`
	UpdateDate string `bson:"update_date" json:"update_date"`
}