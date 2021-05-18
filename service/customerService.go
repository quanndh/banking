package service

import (
	"github.com/quannguyennn/banking/domain"
	"github.com/quannguyennn/banking/dto"
	"github.com/quannguyennn/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *errs.AppError)
	GetCustomerDetail(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	}
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomerDetail(id string) (*dto.CustomerResponse,  *errs.AppError)  {
	c, err := s.repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	resp := c.ToDto()

	return &resp, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{ repo: repository }
}