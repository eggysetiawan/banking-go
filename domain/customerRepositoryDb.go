package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/eggysetiawan/banking-go/errs"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	customers := make([]Customer, 0)
	findAllSql := "select * from customers"

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Println("Error while querying customers table" + err.Error())
		return nil, err
	}

	for rows.Next() {
		var c Customer
		rows.Scan(&c.Id, &c.Name, &c.City, &c.PostalCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers table" + err.Error())
			return nil, err
		}
		customers = append(customers, c)
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
			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "rahmat:P@ssw0rd@tcp(localhost:3306)/banking-go")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
