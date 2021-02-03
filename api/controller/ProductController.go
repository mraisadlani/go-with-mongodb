package controller

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vanilla/go-with-mongodb/api/common"
	"github.com/vanilla/go-with-mongodb/api/entity"
	"github.com/vanilla/go-with-mongodb/api/payload"
	"github.com/vanilla/go-with-mongodb/api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetAllProduct(c *gin.Context) {
	db, err := common.Connection()

	if err != nil {
		payload.ResponseError(c, http.StatusInternalServerError, err.Error())
	}
	
	var (
		limit, _ = strconv.ParseInt(c.Param("limit"), 10, 64)
		page, _ = strconv.ParseInt(c.Param("page"), 10, 64)
	)

	// Limit and Paging data
	var pages int64

	// Set paging per page
	if page == 1 {
		pages = page
	} else {
		pages = page * 10
	}

	repo := repository.NewProductRepositoryImpl(db)
	get, err := repo.FindAll(limit, pages)

	if err != nil {
		payload.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	count, err := repo.CountAllUsers()

	if err != nil {
		payload.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	payload.ResponsePage(c, http.StatusOK, "Berhasil mendapatkan data", get, len(get), count, limit, page)
}

func FindProduct(c *gin.Context) {
	db, err := common.Connection()

	if err != nil {
		payload.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	id := c.Param("id")

	repo := repository.NewProductRepositoryImpl(db)
	get, err := repo.FindById(id)

	if err != nil {
		payload.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	payload.Response(c, http.StatusCreated, "Berhasil mendapatkan data", get)
}

func CreateProduct(c *gin.Context) {
	db, err := common.Connection()

	if err != nil {
		payload.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	var product entity.Product

	err = c.ShouldBindJSON(&product)

	if err != nil {
		payload.ResponseError(c, http.StatusUnprocessableEntity, err.Error())
	}

	product.ID = primitive.NewObjectID()
	product.CreateBy = "Dummy"
	product.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	product.UpdateBy = "Dummy"
	product.UpdateDate = time.Now().Format("2006-01-02 15:04:05")

	strings.Replace(product.CreateDate, "\"", "", -1)

	repo := repository.NewProductRepositoryImpl(db)
	get, err := repo.Save(&product)

	if err != nil {
		payload.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	payload.Response(c, http.StatusCreated, "Berhasil membuat product", get)
}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {

}
