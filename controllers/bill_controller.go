package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBill(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

func GetBills(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func GetBill(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func UpdateBill(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func DeleteBill(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
