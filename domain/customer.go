package domain

import "github.com/eggysetiawan/banking-go/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	PostalCode  string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
