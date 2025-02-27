package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shota612/invoice-payment-service/server/controllers"
)

func SetupRouter(invoiceController *controllers.InvoiceController) *gin.Engine {
	router := gin.Default()

	// TODO: Add Authentication Middleware
	//	router.Use(middleware.AuthMiddleware())

	api := router.Group("/api")
	{
		api.POST("/invoices", invoiceController.CreateInvoice)
		api.GET("/invoices", invoiceController.GetInvoicesByDateRange)
	}
	return router
}
