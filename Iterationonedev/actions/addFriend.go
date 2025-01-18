package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"testing/utils"
)

var addFriend = `
<!DOCTYPE html>
<html>
<head>
    <title>Kyle Connect - Add a friend</title>
</head>
<body>
    <h1>Add a friend</h1>
	<p>Enter a name of a friend that you would like to add</p>
    <form method="POST" action="/form">
        <label for="username">Username:</label>
        <input type="username" id="username" name="username" placeholder="Enter username to add" required><br><br>

        <input type="submit" value="Send friend request">
    </form>

    {{if .Username}}
    <p>Your friend request has been sent to {{.Username}}!</p>
    {{end}}
</body>
</html>
`

type User struct {
    Username string
	UI template.HTML
}

func question(fromUser string) string {
    var value string
    fmt.Print("Do you accept or decline this friend request from ", fromUser, "?\n")
    fmt.Scan(&value)
    return value
}

func convertStringToInt(number string) int {
    i, err := strconv.Atoi(number)
    if err != nil {
        panic(err)
    }
    return i
}

// Handler function to serve the form and process submissions
func formHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the HTML template
    tmpl, err := template.New("form").Parse(addFriend)
    if err != nil {
        log.Fatal(err)
    }
    // Initialize form data
    userData := User{}

    if r.Method == http.MethodPost {
        // Parse form data
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Error parsing form data", http.StatusBadRequest)
            return
        }

        userData.Username = r.FormValue("username")
		enteredUsername := userData.Username

        db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect") // this line of code works for localhost but not docker!
        // db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")

		username, er := utils.RetrieveUsernameFromDb(db, enteredUsername)
		if err != nil {
			log.Fatal(er)
		}
		
		if er != nil {
			log.Fatal("error retrieving username from database", er)
		}

        // this should get a userid for a given inputted username
        id, err := utils.GetUserId(db, username)

        // convert the id variable above from a string to an int 
        idConverted := convertStringToInt(id)
        if err != nil {
            fmt.Println("error finding userid for", id)
        } 

        kylesId, err := utils.GetUserId(db, "Kyle") // need to change this in the future but keeping it like this for NOW!
        kylesIdConverted := convertStringToInt(kylesId)
        if err != nil {
            fmt.Println("error finding userid for", kylesId)
        } 

        status := "pending"
        // if we have any pending values we can either accept them or not
        // if we accept the friend request then I, add my friend to a new db / table called "friends" --- done
        // if I say decline then I, delete that friend out of the table --- done  

		if strings.TrimSpace(strings.ToLower(username)) == strings.TrimSpace(strings.ToLower(enteredUsername)) {

            // below I put the values for the friend request into a friend request table 
            utils.PutDataToFriendRequestTable(db, kylesIdConverted, "Kyle", idConverted, username, status)

            if status == "pending" {

                answerFromUser := question("Kyle")
                
                if answerFromUser == "accept" {
                    status = "accept"
                    utils.UpdateFriendRequestStatus(db, status, username)

                    // here I put user1 and user2 into the friends table
                    utils.PutFriendsToFriendsTable(db, "Kyle", username) // need to fix the hard-coded Kyle!
                } 

                if answerFromUser == "decline" {
                    status = "decline"
                    utils.DeclineFriendRequest(db, username)
                }
            } else {
                fmt.Println("no status is:", status)
            }
		} else {
			fmt.Println("noooo", er)
		}
        utils.CatchError(err)
        defer db.Close()
    }
    // Render the HTML template with the form data
    tmpl.Execute(w, userData)
}

func main() {
    // Set up the route and handler for the form
    http.HandleFunc("/", formHandler)

    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8082")
    log.Fatal(http.ListenAndServe(":8082", nil))
}
