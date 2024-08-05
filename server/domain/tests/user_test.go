package tests

import (
	"github.com/shota612/invoice-payment-service/server/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserModel(t *testing.T) {
	user := domain.User{
		Name:      "Test User",
		Email:     "test@example.com",
		Password:  "securepassword",
		CompanyID: 1,
	}

	assert.Equal(t, "Test User", user.Name)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "securepassword", user.Password)
	assert.Equal(t, uint(1), user.CompanyID)
}
