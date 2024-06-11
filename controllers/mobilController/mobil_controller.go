package mobilController

import (
	"encoding/json"
	"gorestapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var mobil []models.Mobil

	models.DB.Find(&mobil)
	c.JSON(http.StatusOK, gin.H{"Mobil": mobil})

}

func Show(c *gin.Context) {

	var mobil models.Mobil
	id := c.Param("id")

	if err := models.DB.First(&mobil, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"mobil": mobil})

}

func Create(c *gin.Context) {

	var mobil models.Mobil

	if err := c.ShouldBindJSON(&mobil); err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}

	models.DB.Create(&mobil)
	c.JSON(http.StatusOK, gin.H{"mobil": mobil})

}

func Update(c *gin.Context) {

	var mobil models.Mobil

	id := c.Param("id")

	if err := c.ShouldBindJSON(&mobil); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}

	if models.DB.Model(&mobil).Where("id = ?", id).Updates(&mobil).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate mobil"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbaharui"})
}

func Delete(c *gin.Context) {

	var mobil models.Mobil

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&mobil, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Mobil yang ingin dihapus tidak ditemukan"})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
