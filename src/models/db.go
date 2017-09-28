package models

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func GetDB(connectionAddress string) (*DBManager, error) {
	connection, _ := sql.Open("mysql", connectionAddress)

	return &DBManager{Books: BookManager{connection:connection}}, nil
}

func GetFakeDB() (*DBManager, error){
	return &DBManager{Books: FakeBookManager{}}, nil
}

type DBManager struct {
	Books IBookManager
}


type IBookManager interface {
	AllBooks() ([]*Book, error)
}

