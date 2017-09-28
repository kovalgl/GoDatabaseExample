package models

import "database/sql"


/*
Implements IBookManager interface
*/
type BookManager struct {
	connection *sql.DB
}

func (bm BookManager) AllBooks() ([]*Book, error) {
	rows, _ := bm.connection.Query("select * from `books`")

	defer rows.Close()

	bks := make([]*Book, 0)

	for rows.Next() {
		bk := &Book{}

		rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)

		bks = append(bks, bk)
	}

	return bks, nil
}