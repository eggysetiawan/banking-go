package dto

type CustomerResponse struct {
	Id          string `json:"customerId"`
	Name        string `json:"fullName"`
	City        string `json:"city"`
	PostalCode  string `json:"postalCode"`
	DateOfBirth string `json:"dateOfBirth"`
	Status      string `json:"status"`
}
