package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vanilla/go-with-mongodb/api/controller"
)


func main() {
	port := os.Getenv("PORT")

	r := gin.Default()

	RestAPI := r.Group("/api")
	{
		RestAPI.GET("/GetProduct", controller.GetAllProduct)
		RestAPI.GET("/FindProduct", controller.FindProduct)
		RestAPI.POST("/CreateProduct", controller.CreateProduct)
		RestAPI.POST("/UpdateProduct/:id", controller.UpdateProduct)
		RestAPI.POST("/DeleteProduct/:id", controller.DeleteProduct)
	}

	r.Run(fmt.Sprintf(":%s", port))
}