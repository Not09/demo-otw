package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/not09/demo-otw/entity"
)

// POST /engheadfas
func CreateEngheadfa(c *gin.Context) {
	var engheadfa entity.Engheadfa
	if err := c.ShouldBindJSON(&engheadfa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&engheadfa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": engheadfa})
}

// GET /engheadfa/:id
func GetEngheadfas(c *gin.Context) {
	var engheadfa entity.Requirecustomer
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM engheadfas WHERE id = ?", id).Scan(&engheadfa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": engheadfa})
}

// GET /engheadfas
func LisEngheadfas(c *gin.Context) {
	var engheadfas []entity.Requirecustomer
	if err := entity.DB().Raw("SELECT * FROM engheadfas").Scan(&engheadfas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": engheadfas})
}

// DELETE /engheadfas/:id
func DeleteEngheadfa(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM engheadfas WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "engheadfa not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /engheadfas
func UpdateEngheadfa(c *gin.Context) {
	var engheadfa entity.Requirecustomer
	if err := c.ShouldBindJSON(&engheadfa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", engheadfa.ID).First(&engheadfa); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	if err := entity.DB().Save(&engheadfa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": engheadfa})
}