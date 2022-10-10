package domain

import (
	"database/sql"
	"time"

	"github.com/eggysetiawan/banking-go/errs"
	"github.com/eggysetiawan/banking-go/errs/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAllActive() ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)

	findAllSql := "select * from customers where status = ?"

	err := d.client.Select(&customers, findAllSql, 1)

	if err != nil {
		logger.Error("Error while querying customers active table" + err.Error())
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindAllInactive() ([]Customer, *errs.AppError) {
	findAllSql := "select * from customers where status = ?"
	rows, err := d.client.Query(findAllSql, 0)

	if err != nil {
		logger.Error("Error while querying customers inactive table" + err.Error())
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.PostalCode, &c.DateOfBirth, &c.Status)

		if err != nil {
			logger.Error("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		customers = append(customers, c)
	}

	return customers, nil

}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select * from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select * from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)

	}

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select * from customers where id = ?"

	row := d.client.QueryRow(customerSql, id)

	var c Customer

	err := row.Scan(&c.Id, &c.Name, &c.City, &c.PostalCode, &c.DateOfBirth, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "rahmat:P@ssw0rd@tcp(localhost:3306)/banking-go")

	if err != nil {
		errs.NewUnexpectedError("Failed connect to database")
		return CustomerRepositoryDb{nil}
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
