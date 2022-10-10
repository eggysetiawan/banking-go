package domain

import (
	"github.com/eggysetiawan/banking-go/dto"
	"github.com/eggysetiawan/banking-go/errs"
)

type Customer struct {
	Id          string `json:"customerId"`
	Name        string `json:"name"`
	City        string `json:"city"`
	PostalCode  string `json:"postalCode" db:"zipcode"`
	DateOfBirth string `json:"dateOfBirth" db:"date_of_birth"`
	Status      string `json:"status"`
}

func (c Customer) statusAsText() string {
	statusAsText := "active"

	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		PostalCode:  c.PostalCode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

// customerRepositoryDb.go
// customerRepositoryStub.go

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindAllInactive() ([]Customer, *errs.AppError)
	FindAllActive() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
