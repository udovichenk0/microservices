package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/udovichenk0/microservices/order/internal/application/core/domain"
)

type DBAdapter struct {
	db *sql.DB
}

func NewDBAdapter(source string) *DBAdapter {
	db, err := sql.Open("mysql", source)
	if err != nil {
		log.Fatal(err)
	}
	return &DBAdapter{db}
}

func (db *DBAdapter) Save(order *domain.Order) error {
	//save to db
	return nil
}

func (db *DBAdapter) Get(id int64) (domain.Order, error) {
	//get from db
	return domain.Order{}, nil
}
