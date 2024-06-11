package routes

import (
	"gorestapi/controllers/mobilController"
	"gorestapi/controllers/wisataController"

	"github.com/gin-gonic/gin"
)

func RoutesApi(router *gin.Engine) {

	api := router.Group("/api")
	{
		//Unit Mobil
		api.GET("/mobil", mobilController.Index)
		api.GET("/mobil-show/:id", mobilController.Show)
		api.POST("/mobil-create", mobilController.Create)
		api.PUT("/mobil-update/:id", mobilController.Update)
		api.DELETE("/mobil-delete", mobilController.Delete)

		//wisata
		api.GET("/wisata", wisataController.Index)
		api.GET("/wisata-show/:id", wisataController.Show)
		api.POST("/wisata-create", wisataController.Create)
		api.PUT("/wisata-update/:id", wisataController.Update)
		api.DELETE("/wisata-delete", wisataController.Delete)
	}
}
