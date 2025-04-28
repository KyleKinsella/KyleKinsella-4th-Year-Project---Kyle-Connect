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
    "testing/ui"
)

var addFriend = `
  <style>
    body {
        font-family: 'Arial', sans-serif;
        background: linear-gradient(to right, #141e30, #243b55);
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        margin: 0;
        color: white;
    }

    .container {
        background: rgba(255, 255, 255, 0.1);
        padding: 30px;
        border-radius: 12px;
        box-shadow: 0 4px 20px rgba(0, 0, 0, 0.4);
        text-align: center;
        width: 100%;
        max-width: 500px;

    }
        
    h1 {
        margin-bottom: 20px;
        font-size: 28px;
        color: #00aaff;
        font-weight: 600;
    }

    form {
        display: flex;
        flex-direction: column;
        gap: 20px;
        align-items: center;
    }
   
    .input-group {
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: 100%;
        max-width: 400px;
    }

    label {
        font-size: 16px;
        font-weight: 600;
        margin-right: 10px;
        white-space: nowrap;
    }

    input[type="username"] {
        flex: 1;
        padding: 10px;
        font-size: 16px;
        border-radius: 8px;
        border: 2px solid #444;
        outline: none;
        background-color: #1a2939;
        color: white;
        transition: border-color 0.3s ease-in-out;
    }

    input[type="username"]:focus {
        border-color: #00aaff;
    }

    input[type="submit"] {
        background-color: #00aaff;
        color: white;
        padding: 14px 0;
        border: none;
        border-radius: 8px;
        cursor: pointer;
        font-size: 18px;
        font-weight: 600;
        width: 100%;
        max-width: 400px;
        transition: background-color 0.3s ease;
    }

    input[type="submit"]:hover {
        background-color: #0088cc;
    }
    </style>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kyle Connect - Add a Friend</title>
</head>
<body>
    <div class="container">
        <h1>Add a Friend</h1>
        <p>Enter the username of a friend you’d like to add.</p>
        <form method="POST" action="/form">
            <div class="input-group">
                <label for="username">Username:</label>
                <input type="username" id="username" name="username" placeholder="Enter username to add" required>
            </div>
            <input type="submit" value="Send Friend Request">
        </form>

        {{if .Username}}
        <p class="success-message">Your friend request has been sent to {{.Username}}!</p>
        {{end}}
    </div>
</body>
</html>
`

type User struct {
    Username string
    Answer string
	UI template.HTML
}

func convertStringToInt(number string) int {
    i, err := strconv.Atoi(number)
    if err != nil {
        panic(err)
    }
    return i
}

func convertIntToString(number int) string {
	str := strconv.Itoa(number)
    return str
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

        // db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect") // this line of code works for localhost but not docker!
        db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")

		username, er := utils.RetrieveUsernameFromDb(db, enteredUsername)
		if err != nil {
			log.Fatal(er)
		}

        if username == "" {
            communicators := utils.GetCommunicatorsUsernames(db)
            for _, n := range communicators {
                if n != username {
                    fmt.Println("the name you entered", enteredUsername, "is not within my system, try again with a new name...")
                    
                    w.Header().Set("Content-Type", "text/html")

                    t, err := template.New("").Parse(ui.NoUsernameWithThatName)
                    if err != nil {
                        http.Error(w, "Template parsing error", http.StatusInternalServerError)
                        return
                    }
                    t.Execute(w, nil)
                    return
                }
            }
        } else {
            fmt.Println("username is:", username)
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
        
        lastUser, e := utils.GetLastUserLoggedIn(db)
        if e != nil {
            fmt.Println("error getting last user in logged-in table", e)
        }
        utils.CatchError(e)
        fmt.Println("value variable = ", lastUser)
        
        lastUserFromIntToString := convertIntToString(lastUser)

        emailId, e := utils.RetrieveEmailFromId(db, lastUserFromIntToString)
        if e != nil {
            fmt.Println("error getting email from id", e)
        }
        fmt.Println("emailId is:", emailId)
        
        loggedInUser := utils.RetrieveEmail(db, emailId)
        fmt.Println("loggedInUser is:", loggedInUser)

        loggedInUserId, er := utils.GetUserId(db, loggedInUser)
        convertedValue := convertStringToInt(loggedInUserId)

        friends := utils.GetFriends(db, loggedInUser)
        friends2 := utils.GetFriends2(db, loggedInUser)

        var totalFriends []string
        totalFriends = append(totalFriends, friends...)
        totalFriends = append(totalFriends, friends2...)
        
        status := "pending"
		if strings.TrimSpace(strings.ToLower(username)) == strings.TrimSpace(strings.ToLower(enteredUsername)) {

            if loggedInUser == username {
                fmt.Println("you cannot send a friend reuqest to yourself!")

                t, err := template.New("").Parse(ui.YouCannotSendAFriendRequestToYourself)
                if err != nil {
                    http.Error(w, "Template parsing error", http.StatusInternalServerError)
                    return
                }
                t.Execute(w, nil)
                return
            }

            for _, n := range totalFriends {
                if n == username {
                    fmt.Println("you cannot send a friend request because you are already friends")

                    w.Header().Set("Content-Type", "text/html")

                    t, err := template.New("").Parse(ui.CannotSendFriendRequest)
                    if err != nil {
                        http.Error(w, "Template parsing error", http.StatusInternalServerError)
                        return
                    }
                    t.Execute(w, nil)
                    return
                }
            }
            // below I put the values for the friend request into a friend request table 
            utils.PutDataToFriendRequestTable(db, convertedValue, loggedInUser, idConverted, username, status)
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
