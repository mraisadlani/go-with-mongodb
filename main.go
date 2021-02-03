package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vanilla/go-with-mongodb/api/common"
	"github.com/vanilla/go-with-mongodb/api/controller"
	"github.com/vanilla/go-with-mongodb/api/payload"
)


func main() {
	port := os.Getenv("PORT")

	r := gin.Default()

	RestAPI := r.Group("/api")
	{
		RestAPI.GET("/GetProduct", controller.GetAllProduct)
		// RestAPI.GET("/FindProduct", controller.FindProduct)
		// RestAPI.POST("/CreateProduct", controller.CreateProduct)
		// RestAPI.POST("/UpdateProduct/:id", controller.UpdateProduct)
		// RestAPI.POST("/DeleteProduct/:id", controller.DeleteProduct)
	}

	r.Run(fmt.Sprintf(":%s", port))
}