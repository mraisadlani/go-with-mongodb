package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vanilla/go-with-mongodb/api/common"
	"github.com/vanilla/go-with-mongodb/api/payload"
	"github.com/vanilla/go-with-mongodb/api/repository"
)


func GetAllProduct(c *gin.Context) {
	db, err := common.Connection()

	if err != nil {

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
		payload.ResponseError(http.StatusInternalServerError, err.Error())
	}

	payload.Response(http.StatusOK, "Berhasil mendapatkan data", get)
}

func FindProduct(c *gin.Context) {

}

func CreateProduct(c *gin.Context) {

}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {

}
