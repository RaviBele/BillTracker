package main

import (
	"errors"
	"fmt"
	"log"

	"bill_manager/controllers"
	"bill_manager/database"
	"bill_manager/models"
	"bill_manager/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()

	controllers.RegisterValidations()
	// Set up Gin
	router := gin.Default()

	// Register routes
	routes.RegisterAccountRoutes(router)
	routes.RegisterBillRoutes(router)
	routes.RegisterVendorRoutes(router)

	// Start the server
	router.Run(":8080")
}

func loadEnv() {
	if err := godotenv.Load(".env.local"); err != nil {
		handleError(errors.New("error loading .env file"))
	}
}

func loadDatabase() {
	if err := database.Connect(); err != nil {
		handleError(err)
	}

	if err := database.Database.AutoMigrate(&models.Account{}); err != nil {
		log.Printf("error loading database Account: %v", err.Error())
		handleError(err)
	}
	if err := database.Database.AutoMigrate(&models.Bill{}); err != nil {
		log.Printf("error loading database Bill: %v", err.Error())
		handleError(err)
	}
	if err := database.Database.AutoMigrate(&models.Vendor{}); err != nil {
		log.Printf("error loading database Vendor: %v", err.Error())
		handleError(err)
	}
	if err := database.Database.AutoMigrate(&models.VendorProduct{}); err != nil {
		log.Printf("error loading database VendorProduct: %v", err.Error())
		handleError(err)
	}
	if err := database.Database.AutoMigrate(&models.Product{}); err != nil {
		log.Printf("error loading database Product: %v", err.Error())
		handleError(err)
	}
	if err := database.Database.AutoMigrate(&models.BillProduct{}); err != nil {
		log.Printf("error loading database BillProduct: %v", err.Error())
		handleError(err)
	}
	fmt.Println("Migrations executed successfully")
}

func handleError(err error) {
	log.Fatal(err)
}
