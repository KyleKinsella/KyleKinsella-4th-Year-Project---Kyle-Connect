package utils

import (
	"database/sql"
	"fmt"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

func CatchError(err error) error {
	if err != nil {
		panic(err.Error())
	}
	return err
}

func RetrieveDataFromDb(db *sql.DB, email string) (string, error) {
    var hashedPassword string
    err := db.QueryRow("SELECT password FROM communicators WHERE email = ?", email).Scan(&hashedPassword)
    if err == sql.ErrNoRows {
        return "", errors.New("no user with that email")
    } else if err != nil {
        return "", err
    }
    return hashedPassword, nil
}

func PutDataToDb(db *sql.DB, username, email, password string) error {
	sql := "INSERT INTO communicators (username, email, password) VALUES (?, ?, ?)"
	_, err := db.Query(sql, username, email, password)
	CatchError(err)

	fmt.Println(username, "inserted into database.")
	return err
}