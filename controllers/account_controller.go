package controllers

import (
	"bill_manager/database"
	"bill_manager/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	DEFAULT_MIN_LIMIT = 10
	DEFAULT_MAX_LIMIT = 50
	DEFAULT_PAGE      = 1
)

func CreateAccount(c *gin.Context) {
	var newAccount models.Account
	if err := c.ShouldBindJSON(&newAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validate.Struct(newAccount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAccount.ID = uuid.New() // Generate UUID
	result := database.Database.Create(&newAccount)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, newAccount)
}

func GetAccounts(c *gin.Context) {
	limitQuery := c.Query("limit")
	pageQuery := c.Query("page")

	if limitQuery == "" {
		limitQuery = fmt.Sprintf("%d", DEFAULT_MIN_LIMIT)
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		log.Printf("invalid limit query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if limit < 1 {
		limit = DEFAULT_MIN_LIMIT
	}

	if limit > DEFAULT_MAX_LIMIT {
		limit = DEFAULT_MAX_LIMIT
	}

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		log.Printf("invalid page query: %v", err)
		page = DEFAULT_PAGE
	}

	var accounts []models.Account
	offset := (page - 1) * limit

	log.Printf("Fetching accounts with limit %d, offset: %d, page: %d", limit, offset, page)

	if err := database.Database.Table("accounts").Order("created_at desc").Limit(limit).Offset(offset).Find(&accounts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Returning accounts : %d", len(accounts))

	accountResponse := models.AccountListResponse{
		Pagination: models.Pagination{
			Page:  page,
			Count: len(accounts),
		},
		Data: accounts,
	}

	c.JSON(http.StatusOK, accountResponse)
}

func GetAccount(c *gin.Context) {
	id := c.Param("account_id")
	var account models.Account
	if err := database.Database.First(&account, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, account)
}

func UpdateAccount(c *gin.Context) {
	id := c.Param("account_id")
	var account models.Account
	if err := database.Database.First(&account, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.Database.Save(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func DeleteAccount(c *gin.Context) {
	id := c.Param("account_id")
	if err := database.Database.Delete(&models.Account{}, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account deleted"})
}
