package domain

import (
	"github.com/quannguyennn/banking/dto"
	"github.com/quannguyennn/banking/errs"
)

type Customer struct {
	Id 				string			`db:"customer_id"`
	Name 			string
	City 			string
	Zipcode 		string
	DateOfBirth 	string			`db:"date_of_birth"`
	Status			string
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
		Id: 			c.Id,
		Name: 			c.Name,
		City: 			c.City,
		Zipcode: 		c.Zipcode,
		DateOfBirth: 	c.DateOfBirth,
		Status: 		c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindOne(string) (*Customer,  *errs.AppError)
}
