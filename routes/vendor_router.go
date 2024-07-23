package routes

import (
	"bill_manager/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterVendorRoutes(router *gin.Engine) {
	router.POST("/api/vendors", controllers.CreateVendor)
	router.GET("/api/vendors", controllers.GetVendors)
	// router.GET("/api/vendors/:vendor_id", controllers.GetVendor)
	// router.PUT("/api/vendors/:vendor_id", controllers.UpdateVendor)
	// router.DELETE("/api/vendors/:vendor_id", controllers.DeleteVendor)
}
