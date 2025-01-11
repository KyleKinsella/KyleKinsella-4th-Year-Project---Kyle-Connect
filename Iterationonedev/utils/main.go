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

func RetrieveUsernameFromDb(db *sql.DB, username string) (string, error) {
    err := db.QueryRow("SELECT username FROM communicators WHERE username = ?", username).Scan(&username)
	if err == sql.ErrNoRows {
        return "", errors.New("no user with that username found")
    } else if err != nil {
        return "", err
    }
    return username, nil
}

func PutDataToDb(db *sql.DB, username, email, password string) error {
	sql := "INSERT INTO communicators (username, email, password) VALUES (?, ?, ?)"
	_, err := db.Query(sql, username, email, password)
	CatchError(err)

	fmt.Println(username, "inserted into database.")
	return err
}


func PutDataToFriendRequestTable(db *sql.DB, fromUserId int, fromUserName string, toUserId int, toUserName string, status string) error {
	sql := "INSERT INTO friendrequest (fromUserId, fromUserName, toUserId, toUserName, status) VALUES (?, ?, ?, ?, ?)" 
	_, err := db.Query(sql, fromUserId, fromUserName, toUserId, toUserName, status);	
	CatchError(err);

	fmt.Println(fromUserName, "has sent a friend request to", toUserName)
	return err;
}

func GetUserId(db *sql.DB, username string) (string, error) {
	getUserid := db.QueryRow("SELECT userId FROM communicators WHERE username = ?", username).Scan(&username)
	if getUserid == sql.ErrNoRows {
		return "", errors.New("no userid found for that username")
	} else if getUserid != nil {
		return "", getUserid
	} 
	return username, nil
}