package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/not09/demo-otw/entity"
)

// POST /requirecustomers
func CreateRequirecustomer(c *gin.Context) {
	var requirecustomer entity.Requirecustomer
	if err := c.ShouldBindJSON(&requirecustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&requirecustomer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": requirecustomer})
}

// GET /requirecustomer/:id
func GetRequirecustomers(c *gin.Context) {
	var requirecustomer entity.Requirecustomer
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM requirecustomers WHERE id = ?", id).Scan(&requirecustomer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requirecustomer})
}

// GET /requirecustomers
func ListRequirecustomers(c *gin.Context) {
	var requirecustomers []entity.Requirecustomer
	if err := entity.DB().Raw("SELECT * FROM requirecustomers").Scan(&requirecustomers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requirecustomers})
}

// DELETE /requirecustomers/:id
func DeleteRequirecustomer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM requirecustomers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "requirecustomer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /requirecustomers
func UpdateRequirecustomer(c *gin.Context) {
	var requirecustomer entity.Requirecustomer
	if err := c.ShouldBindJSON(&requirecustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", requirecustomer.ID).First(&requirecustomer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "resolution not found"})
		return
	}

	if err := entity.DB().Save(&requirecustomer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requirecustomer})
}