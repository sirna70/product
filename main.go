package main

import (
	"gorestapi/controllers/productController"
	"gorestapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productController.Index)
	r.GET("/api/products/:id", productController.Show)
	r.POST("/api/product", productController.Create)
	r.PUT("/api/product/:id", productController.Update)
	r.DELETE("/api/produc", productController.Delete)

	r.Run()
}
