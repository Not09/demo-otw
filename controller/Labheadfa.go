package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/not09/demo-otw/entity"
)

// POST /labheadfas
func CreateLabheadfa(c *gin.Context) {
	var labheadfa entity.Labheadfa
	if err := c.ShouldBindJSON(&labheadfa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&labheadfa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": labheadfa})
}

// GET /labheadfa/:id
func GetLabheadfas(c *gin.Context) {
	var labheadfa entity.Requirecustomer
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM labheadfas WHERE id = ?", id).Scan(&labheadfa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": labheadfa})
}

// GET /labheadfas
func ListLabheadfas(c *gin.Context) {
	var labheadfas []entity.Requirecustomer
	if err := entity.DB().Raw("SELECT * FROM labheadfas").Scan(&labheadfas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": labheadfas})
}

// DELETE /labheadfas/:id
func DeleteLabheadfa(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM labheadfas WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "labheadfa not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /labheadfas
func UpdateLabheadfa(c *gin.Context) {
	var labheadfa entity.Requirecustomer
	if err := c.ShouldBindJSON(&labheadfa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", labheadfa.ID).First(&labheadfa); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	if err := entity.DB().Save(&labheadfa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": labheadfa})
}