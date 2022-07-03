package domain

import "github.com/eggysetiawan/banking-go/errs"

type Customer struct {
	Id          string `json:"customerId"`
	Name        string `json:"name"`
	City        string `json:"city"`
	PostalCode  string `json:"postalCode"`
	DateOfBirth string `json:"dateOfBirth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
