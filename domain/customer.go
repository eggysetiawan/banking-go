package domain

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
}
