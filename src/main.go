package main

import (
	"models"
	"net/http"
	"fmt"
)

type DBEnv struct{
	db *models.DBManager
}

func setDBEnv(fake bool) *DBEnv {
	var db *models.DBManager
	if fake == true {
		db, _ = models.GetFakeDB()
	} else {
		db, _ = models.GetDB("root:777@tcp(localhost:3306)/books")
	}
	return &DBEnv{db}
}

func main() {
	env := setDBEnv(false)

	http.HandleFunc("/books", env.indexPage)
	http.ListenAndServe(":3000", nil)
}

func (env *DBEnv) indexPage (w http.ResponseWriter, r *http.Request) {
	books, _ := env.db.Books.AllBooks()

	for _, book := range books {
		fmt.Fprintf(w, "%s \"%s\" %s %f\n", book.Isbn, book.Title, book.Author, book.Price)
	}
}
