package routes

import (
	"bill_manager/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(router *gin.Engine) {
	router.POST("/api/accounts", controllers.CreateAccount)
	router.GET("/api/accounts", controllers.GetAccounts)
	router.GET("/api/accounts/:account_id", controllers.GetAccount)
	router.PUT("/api/accounts/:account_id", controllers.UpdateAccount)
	router.DELETE("/api/accounts/:account_id", controllers.DeleteAccount)
}
