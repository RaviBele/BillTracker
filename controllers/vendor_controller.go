package controllers

import (
	"bill_manager/database"
	"bill_manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateVendor(c *gin.Context) {
	var vendor models.Vendor

	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate UUIDs for Vendor and Products
	vendor.ID = uuid.New()
	for i := range vendor.Products {
		vendor.Products[i].ID = uuid.New()
		vendor.Products[i].VendorID = vendor.ID
	}

	if err := database.Database.Create(&vendor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, vendor)
}

func GetVendors(c *gin.Context) {
	var vendors []models.Vendor
	if err := database.Database.Preload("Product").Find(&vendors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vendors)
}
