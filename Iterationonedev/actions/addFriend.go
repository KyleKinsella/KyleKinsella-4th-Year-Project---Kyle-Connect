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
</body>
</html>
`

type User struct {
    Username string
	UI template.HTML
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

        kylesId, err := utils.GetUserId(db, "Kyle")
        kylesIdConverted := convertStringToInt(kylesId)
        if err != nil {
            fmt.Println("error finding userid for", kylesId)
        } 

        status := "sent"
      
        
		if strings.TrimSpace(strings.ToLower(username)) == strings.TrimSpace(strings.ToLower(enteredUsername)) {
            			
            // below I put the values for the friend request into a friend request table 
            utils.PutDataToFriendRequestTable(db, kylesIdConverted, "Kyle", idConverted, username, status)

			// here i need to make the friend request logic
			// show a message to the user that the friend request has been sent 
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
