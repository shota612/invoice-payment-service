package main

import (
	"github.com/shota612/invoice-payment-service/server/api"
	"github.com/shota612/invoice-payment-service/server/controllers"
	"github.com/shota612/invoice-payment-service/server/models"
	"github.com/shota612/invoice-payment-service/server/repository"
	"github.com/shota612/invoice-payment-service/server/usecase"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Company{}, &models.User{}, &models.Client{}, &models.ClientBankAccount{}, &models.Invoice{})
	if err != nil {
		return
	}

	invoiceRepo := repository.NewInvoiceRepository(db)
	invoiceUsecase := usecase.NewInvoiceUsecase(invoiceRepo)
	invoiceController := controllers.NewInvoiceController(invoiceUsecase)

	r := api.SetupRouter(invoiceController)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
