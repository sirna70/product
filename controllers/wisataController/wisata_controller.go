package wisataController

import (
	"encoding/json"
	"gorestapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var wisata []models.Wisata

	models.DB.Find(&wisata)
	c.JSON(http.StatusOK, gin.H{"wisata": wisata})

}

func Show(c *gin.Context) {

	var wisata models.Wisata
	id := c.Param("id")

	if err := models.DB.First(&wisata, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"wisata": wisata})

}

func Create(c *gin.Context) {

	var wisata models.Wisata

	if err := c.ShouldBindJSON(&wisata); err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}

	models.DB.Create(&wisata)
	c.JSON(http.StatusOK, gin.H{"wisata": wisata})

}

func Update(c *gin.Context) {

	var wisata models.Wisata

	id := c.Param("id")

	if err := c.ShouldBindJSON(&wisata); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}

	if models.DB.Model(&wisata).Where("id = ?", id).Updates(&wisata).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate wisata"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wisata berhasil diperbaharui"})
}

func Delete(c *gin.Context) {

	var wisata models.Wisata

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&wisata, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Wisata yang ingin dihapus tidak ditemukan"})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "Wisata berhasil dihapus"})
}
