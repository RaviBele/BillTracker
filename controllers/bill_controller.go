package controllers

import (
	"bill_manager/database"
	"bill_manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateBillHandler handles the creation of a models.Bill
func CreateBill(c *gin.Context) {

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

	var bill models.Bill
	bill.ID = uuid.New()
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accountID, _ := uuid.Parse(id)
	bill.AccountID = accountID
	// Decode the request body into the models.Bill struct

	// Check if VendorID exists
	var vendor models.Vendor
	if result := database.Database.First(&vendor, "id = ?", bill.VendorID); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "models.Vendor ID does not exist"})
		return
	}

	// Calculate the amount based on VendorProduct rates, quantities, and taxes
	totalAmount := 0.0
	totalTax := 0.0
	var products []models.Product
	for _, product := range bill.Products {
		var vendorProduct models.VendorProduct
		if result := database.Database.Preload("Vendor").First(&vendorProduct, "id = ? AND vendor_id = ?", product.VendorProductID, bill.VendorID); result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product does not belong to the specified Vendor"})
			return
		}

		// Calculate the amount for this product including tax
		productAmount := vendorProduct.Rate * float64(product.Quantity)
		taxAmount := productAmount * (vendorProduct.Tax / 100)
		totalAmount += productAmount + taxAmount
		totalTax += taxAmount

		productData := models.Product{
			ID:       uuid.New(),
			Product:  vendorProduct,
			Quantity: product.Quantity,
		}
		// Add the product to the list with the correct VendorProduct details
		products = append(products, productData)
	}

	// Set the calculated amount
	bill.Amount = totalAmount
	bill.TotalTax = totalTax

	// Attach the complete product details back to the bill
	bill.Products = products

	// Generate a new UUID for the models.Bill
	if result := database.Database.Create(&bill); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save bill"})
		return
	}

	var createdBill models.Bill
	if result := database.Database.Preload("Products").First(&createdBill, "id = ?", bill.ID); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch created bill with products"})
		return
	}

	// Respond with the created Bill and VendorProduct details
	c.JSON(http.StatusCreated, bill)
}

func GetBills(c *gin.Context) {
	account_id := c.Param("account_id")
	var account models.Account
	if err := database.Database.First(&account, "id = ?", account_id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var bills []models.Bill
	if result := database.Database.Preload("Products.Product").Find(&bills, "account_id = ?", account_id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch created bill with products"})
		return
	}
	c.JSON(http.StatusOK, bills)
}

func GetBill(c *gin.Context) {
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

	bill_id := c.Param("id")
	var bill models.Bill
	if err := database.Database.First(&bill, "id = ?", bill_id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, bill)
}

func UpdateBill(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func DeleteBill(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
