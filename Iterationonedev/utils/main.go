package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

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
		return "", errors.New("\nTHIS HERE IS THE PROBLEM\nno user with that username found")
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

func InsertLoggedInUserToTable(db *sql.DB, name, email string) (string, string) {
	sql := "INSERT into loggedin (name, email) VALUES (?, ?)"

	_, err := db.Query(sql, name, email)
	CatchError(err)

	fmt.Println(name, "and", email, "has been inserted into the logged in table")
	return name, email
}

func RetrieveEmail(db *sql.DB, email string) string {
	findEmail := db.QueryRow("SELECT username FROM communicators WHERE email = ?", email).Scan(&email)
	if findEmail == sql.ErrNoRows {
		return ""
	} else if findEmail != nil {
		return ""
	} 
	return email
}

func RetrieveUsername(db *sql.DB, username string) (string, error) {
	findUser := db.QueryRow("SELECT email FROM loggedin WHERE name = ?", username).Scan(&username)
	if findUser == sql.ErrNoRows {
		return "", errors.New("no username found for that email")
	} else if findUser != nil {
		return "", findUser
	} 
	return username, nil
}

func GetLastUserLoggedIn(db *sql.DB) (int, error) {
	var id int
	err := db.QueryRow("SELECT id FROM loggedin ORDER BY id DESC LIMIT 1").Scan(&id)
	if err == sql.ErrNoRows {
		return 0, errors.New("no users found")
	} else if err != nil {
		return 0, err
	} 
	return id, nil
}

func RetrieveEmailFromId(db *sql.DB, id string) (string, error) {
	findEmail := db.QueryRow("SELECT email FROM loggedin WHERE id = ?", id).Scan(&id)
	if findEmail == sql.ErrNoRows {
		return "", errors.New("no email found for that username")
	} else if findEmail != nil {
		return "", findEmail
	} 
	return id, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func removeDuplicates(remove []string) []string {
	list := []string{}

	for _, item := range remove {
		// fmt.Println(item)
		if contains(list, item) == false {
			list = append(list, item)
		}
	}
	return list
}

func GetFriends(db *sql.DB, name string) []string {
	var friends []string

	rows, err := db.Query("SELECT user2 FROM friends WHERE user1 = ?", name)
	if err != nil {
		fmt.Println("an error has occured when executing query!")	
	}
	defer rows.Close()

	for rows.Next() {
		var friend string 
		if err := rows.Scan(&friend); err != nil {
			fmt.Println("error scanning row:", err)
		}
		friends = append(friends, friend)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("error iterating over rows:", err)
	}
	return removeDuplicates(friends)
}

func GetFriends2(db *sql.DB, name string) []string {
	var friends []string

	rows, err := db.Query("SELECT user1 FROM friends WHERE user2 = ?", name)
	if err != nil {
		fmt.Println("an error has occured when executing query!")	
	}
	defer rows.Close()

	for rows.Next() {
		var friend string 
		if err := rows.Scan(&friend); err != nil {
			fmt.Println("error scanning row:", err)
		}
		friends = append(friends, friend)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("error iterating over rows:", err)
	}
	return removeDuplicates(friends)
}

func GetPendingRequestsForLoggedInUser(db *sql.DB, loggedInUser, toUser, status string) (string, string, string) {
	pending := db.QueryRow("SELECT status FROM friendrequest WHERE status = ?", status).Scan(&status)
	if pending == sql.ErrNoRows {
		// fmt.Println("no status of pending:", status)
		return "", "", ""
	} else if pending != nil {
		return "", "", ""
	} 
	return loggedInUser, toUser, status
}

func GetToUserName(db *sql.DB, username string) (string, error) {
	findtoUserName := db.QueryRow("SELECT toUserName FROM friendrequest WHERE fromUserName = ?", username).Scan(&username)
	if findtoUserName == sql.ErrNoRows {
		return "", errors.New("no username found with a status of pending")
	} else if findtoUserName != nil {
		return "", findtoUserName
	} 
	return username, nil
}

func WhoSentFriendRequest(db *sql.DB, name string) []string {
	var whosent []string

	rows, err := db.Query("SELECT fromUserName FROM friendrequest WHERE toUserName = ?", name)
	if err != nil {
		fmt.Println("an error has occured when executing query!")	
	}
	defer rows.Close()

	for rows.Next() {
		var who string 
		if err := rows.Scan(&who); err != nil {
			fmt.Println("error scanning row:", err)
		}
		whosent = append(whosent, who)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("error iterating over rows:", err)
	}
	return removeDuplicates(whosent)
}

func LoggedInPossibleFriend(db *sql.DB, user string) string {
	var name string
	err := db.QueryRow("SELECT fromUserName FROM friendrequest WHERE toUserName = ?", user).Scan(&name)
	if err == sql.ErrNoRows {
		// No matching rows
		return ""
	} else if err != nil {
		// Log and handle other errors
		log.Println("Error querying database:", err)
		return ""
	}
	return name
}

func InsertIntoClickedTable(db *sql.DB, name string) string {
	sql := "INSERT into clicked (username) VALUES (?)"

	_, err := db.Query(sql, name)
	CatchError(err)

	fmt.Println(name, "has been inserted into the clicked table")
	return name
}

func InsertMessage(db *sql.DB, sender, receiver int, content, timestamp string) {
	sql := "INSERT into messages (sender, receiver, content, timestamp) VALUES (?, ?, ?, ?)"
	
	_, err := db.Query(sql, sender, receiver, content, timestamp)
	CatchError(err)

	fmt.Println(sender, "sent a message to", receiver)
	fmt.Println("Your message has been inserted into the messages table")
}

func GetLastUserClicked(db *sql.DB) (int, error) {
	var id int
	err := db.QueryRow("SELECT clickedId FROM clicked ORDER BY clickedId DESC LIMIT 1").Scan(&id)
	if err == sql.ErrNoRows {
		return 0, errors.New("no users found")
	} else if err != nil {
		return 0, err
	} 
	return id, nil
}

func GetMessages(db *sql.DB, id int) []string {
	var messages []string

	rows, err := db.Query("SELECT content FROM messages WHERE receiver = ?", id)
	if err != nil {
		fmt.Println("an error has occured when executing query!")	
	}
	defer rows.Close()

	for rows.Next() {
		var message string 
		if err := rows.Scan(&message); err != nil {
			fmt.Println("error scanning row:", err)
		}
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("error iterating over rows:", err)
	}
	// return removeDuplicates(friends)
	return messages
}

func CreateServer(db *sql.DB, serverName string, ownerOfServer string) string {
	sql := "INSERT into server (serverName, ownerOfServer) VALUES (?, ?)"

	_, err := db.Query(sql, serverName, ownerOfServer)
	CatchError(err)

	fmt.Println(ownerOfServer, "has created a server called" , serverName, "this has been inserted into the server table")
	return serverName
}


func Servers(db *sql.DB, ownerOfServer string) []string {
	var serverList []string

	rows, err := db.Query("SELECT serverName FROM server WHERE ownerOfServer = ?", ownerOfServer)
	if err != nil {
		fmt.Println("an error has occured when executing query!")	
	}
	defer rows.Close()

	for rows.Next() {
		var server string 
		if err := rows.Scan(&server); err != nil {
			fmt.Println("error scanning row:", err)
		}
		serverList = append(serverList, server)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("error iterating over rows:", err)
	}
	// return removeDuplicates(whosent)
	return serverList
}

func GetServerId(db *sql.DB) (int, error) {
	var id int
	err := db.QueryRow("SELECT serverId FROM server ORDER BY serverId DESC LIMIT 1").Scan(&id)
	if err == sql.ErrNoRows {
		return 0, errors.New("no users found")
	} else if err != nil {
		return 0, err
	} 
	return id, nil
}

func InsertChannelName(db *sql.DB, channelName string, serverId int) string {
	sql := "INSERT into channel (channelName, serverId) VALUES (?, ?)"

	_, err := db.Query(sql, channelName, serverId)
	CatchError(err)

	fmt.Println(channelName, "has been inserted into the channel table")
	return channelName
}