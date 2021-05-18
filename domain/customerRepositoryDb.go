package domain

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/quannguyennn/banking/errs"
	"github.com/quannguyennn/banking/logger"
	"os"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error

	customers := make([]Customer, 0)

	if status == "" {
		findAllSql :=  "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql :=  "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		return nil, errs.NewServerError("Error get all customer")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindOne(id string) (*Customer, *errs.AppError) {
	sqlQuery := "select * from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, sqlQuery, id)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Customer not found " + err.Error())
			return nil,	errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error" + err.Error())
		return nil, errs.NewServerError("Something went wrong")
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PW")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/$s", dbUser, dbPass, dbHost, dbPort, dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}