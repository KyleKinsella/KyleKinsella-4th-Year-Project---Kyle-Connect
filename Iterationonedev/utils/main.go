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

func UpdateFriendRequestStatus(db *sql.DB, status string, name string) string {
	sql := "UPDATE friendrequest SET status = ? Where toUserName = ?;"

	_, err := db.Query(sql, status, name)
	CatchError(err)

	fmt.Println(name, "has been updated from pending to accept")
	return name
}

func DeclineFriendRequest(db *sql.DB, name string) string {
	sql := "DELETE FROM friendrequest WHERE toUserName = ?"

	_, err := db.Query(sql, name)
	CatchError(err)

	fmt.Println(name, "has been deleted out of the friend request table")
	return name
}

func PutFriendsToFriendsTable(db *sql.DB, friend1 string, friend2 string) []string {
	sql := "INSERT INTO friends (user1, user2) VALUES (?, ?)"

	_, err := db.Query(sql, friend1,friend2)
	CatchError(err)

	fmt.Println(friend1, "and", friend2, "are in the friends table")
	friends := []string{friend1, friend2}
	return friends
}