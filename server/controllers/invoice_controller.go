package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shota612/invoice-payment-service/server/models"
	"github.com/shota612/invoice-payment-service/server/usecase"
	"net/http"
)

type InvoiceController struct {
	invoiceUsecase usecase.InvoiceUsecase
}

func NewInvoiceController(u usecase.InvoiceUsecase) *InvoiceController {
	return &InvoiceController{u}
}

func (ctrl *InvoiceController) CreateInvoice(c *gin.Context) {

	var request struct {
		IssueDate      string               `json:"issue_date" binding:"required"`
		PaymentAmount  float64              `json:"payment_amount" binding:"required"`
		FeeRate        float64              `json:"fee_rate" binding:"required"`
		SalesTaxRate   float64              `json:"sales_tax_rate" binding:"required"`
		PaymentDueDate string               `json:"payment_due_date" binding:"required"`
		Status         models.InvoiceStatus `json:"status" binding:"required"`
		CompanyID      uint                 `json:"company_id" binding:"required"`
		ClientID       uint                 `json:"client_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoice, err := ctrl.invoiceUsecase.CreateInvoice(
		request.IssueDate,
		request.PaymentAmount,
		request.FeeRate,
		request.SalesTaxRate,
		request.PaymentDueDate,
		request.Status,
		request.CompanyID,
		request.ClientID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invoice)
}
