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
    fmt.Print("Do you accept or decline this friend request from ", fromUser, "?")
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
        fmt.Println("the output of id is", id)

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
        // when we send this data to the table it is "sent"!
        // so this means that we can print friend request sent or something like that
        
        // status should be "pending" 
        // query the db / table
        // if we have any pending values we can either accept them or not
        // if we accept the friend request then I, add my friend to a new db / table called "friends"
        // if I say decline then I, delete that friend out of the table  

        // statusArr := []string{"pending", "accepted", "declined"}


        // username = kyle 
        // enteredusename = joe
        // if username != enteredUsername {
        //     fmt.Println("username is not equal to entered username!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
        //     fmt.Println(enteredUsername, "sent a request to ", username)
        // }

		if strings.TrimSpace(strings.ToLower(username)) == strings.TrimSpace(strings.ToLower(enteredUsername)) {

            // below I put the values for the friend request into a friend request table 
            utils.PutDataToFriendRequestTable(db, kylesIdConverted, "Kyle", idConverted, username, status)

            if status == "pending" {
                fmt.Println("yes status is:", status)

                answerFromUser := question("Kyle")
                fmt.Println("answer is:", answerFromUser)

                if answerFromUser == "yes" || answerFromUser == "y" || answerFromUser == "accept" {
                    status = "accept"
                    // i might have to make / change the below query to update instead of making a new insertion
                    utils.PutDataToFriendRequestTable(db, kylesIdConverted, "Kyle", idConverted, username, status)
                } 

                if answerFromUser == "no" || answerFromUser == "n" || answerFromUser == "decline" {
                    status = "decline"
                    // i might have to make / change the below query to update instead of making a new insertion
                    // dont do the below query, I need to delete this request from the friend request table due to me not wanting to 
                    // accept the friend request from x user

                    utils.PutDataToFriendRequestTable(db, kylesIdConverted, "Kyle", idConverted, username, status) //  <= DON'T DO! 
                    fmt.Println("you have declined the friend request, this has been removed from the friend request table!")
                }
            } else {
                fmt.Println("no status is:", status)
            }

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
