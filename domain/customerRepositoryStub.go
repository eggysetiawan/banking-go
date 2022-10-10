package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Rahmat Setiawan", "Jakarta", "14250", "1998-02-16", "1"},
		{"2", "Tiara Ramadayanti", "Bekasi", "14430", "1998-01-27", "1"},
	}
	return CustomerRepositoryStub{customers}
}
