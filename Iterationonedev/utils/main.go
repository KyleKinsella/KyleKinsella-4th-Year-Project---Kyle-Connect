package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CatchError(err error) error {
	if err != nil {
		panic(err.Error())
	}
	return err
}

func RetrieveDataFromDb(db *sql.DB, email, password string) []string {
	sql := "SELECT email, password FROM communicators WHERE email=? AND password=?"
	_, err := db.Query(sql, email, password)
	CatchError(err)
	
	totalData := []string{email, password}

	for _, i := range totalData {
		fmt.Println("here is some data retrieved from the database\n", i)
	}

	return totalData
}

func PutDataToDb(db *sql.DB, username, email, password string) error {
	sql := "INSERT INTO communicators (username, email, password) VALUES (?, ?, ?)"
	_, err := db.Query(sql, username, email, password)
	CatchError(err)

	fmt.Println(username, "inserted into database.")
	return err
}