package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

func GetAccounts(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func GetAccount(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func UpdateAccount(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func DeleteAccount(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
