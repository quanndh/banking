package service

import (
	"github.com/quannguyennn/banking/domain"
	"github.com/quannguyennn/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *errs.AppError)
	GetCustomerDetail(id string) (*domain.Customer, *errs.AppError)
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

func (s DefaultCustomerService) GetCustomerDetail(id string) (*domain.Customer,  *errs.AppError)  {
	return s.repo.FindOne(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{ repo: repository }
}