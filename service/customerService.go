package service

import (
	"github.com/eggysetiawan/banking-go/domain"
	"github.com/eggysetiawan/banking-go/dto"
	"github.com/eggysetiawan/banking-go/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetAllCustomerInactive() ([]dto.CustomerResponse, *errs.AppError)
	GetAllCustomerActive() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {

	if status == "inactive" {
		status = "0"
	} else if status == "active" {
		status = "1"
	} else {
		status = ""
	}

	customers, err := s.repo.FindAll(status)

	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}

	return response, nil

}

func (s DefaultCustomerService) GetAllCustomerActive() ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAllActive()

	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, 0)

	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, nil
}

func (s DefaultCustomerService) GetAllCustomerInactive() ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAllInactive()

	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)

	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, nil

}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
