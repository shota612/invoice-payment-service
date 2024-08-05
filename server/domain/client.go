package domain

type Client struct {
	ID             uint
	LegalName      string
	Representative string
	PhoneNumber    string
	PostalCode     string
	Address        string
	CompanyID      uint
}
