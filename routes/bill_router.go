package routes

import (
	"bill_manager/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBillRoutes(router *gin.Engine) {
	router.POST("/api/accounts/:account_id/bills", controllers.CreateBill)
	router.GET("/api/accounts/:account_id/bills", controllers.GetBills)
	router.GET("/api/accounts/:account_id/bills/:id", controllers.GetBill)
	router.PUT("/api/accounts/:account_id/bills/:id", controllers.UpdateBill)
	router.DELETE("/api/accounts/:account_id/bills/:id", controllers.DeleteBill)
}
