package domain

import "github.com/eggysetiawan/banking-go/errs"

type Customer struct {
	Id          string `json:"customerId"`
	Name        string `json:"name"`
	City        string `json:"city"`
	PostalCode  string `json:"postalCode" db:"zipcode"`
	DateOfBirth string `json:"dateOfBirth" db:"date_of_birth"`
	Status      string `json:"status"`
}

// customerRepositoryDb.go
// customerRepositoryStub.go

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindAllInactive() ([]Customer, *errs.AppError)
	FindAllActive() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
