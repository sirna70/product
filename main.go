package main

import (
	"gorestapi/models"
	"gorestapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()

	routes.RoutesApi(r)

	r.Run()
}
