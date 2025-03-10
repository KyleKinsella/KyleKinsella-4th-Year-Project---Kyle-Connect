package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"testing/utils"
	"time"
)

var sendMessage = `
<!DOCTYPE html>
<html>
<head>
    <title>Kyle Connect - Send Message</title>
</head>
<body>
    <h1>Send a Message to a friend</h1>
    <form method="POST" action="/form">
        <label for="message">Your Message:</label>
        <input type="text" id="message" name="message" placeholder="Enter your message" required><br><br>

        <input type="submit" value="Send Message">
    </form>
</body>
</html>
`

type Message struct {
	Message string
    UI template.HTML
}

// Handler function to serve the form and process submissions
func formHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the HTML template
    tmpl, err := template.New("form").Parse(sendMessage) // might need to do it here!?
    if err != nil {
        log.Fatal(err)
    }

    // Initialize form data
    msg := Message{}

    if r.Method == http.MethodPost {
        // Parse form data
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Error parsing form data", http.StatusBadRequest)
            return
        }

        msg.Message = r.FormValue("message")

		currentTime := time.Now()
		timeAndDate := currentTime.Format("2006-01-02 15:04:05")

        db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect") // this line of code works for localhost but not docker! MAKE SURE TO COMMENT THIS OUT WHEN WORKING WITH DOCKER!!!!!!!!!!!!!!!!
        // db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")

        utils.CatchError(err)
        defer db.Close()

		loggedInUser, err := utils.GetLastUserLoggedIn(db)
		utils.CatchError(err)

		clickedUser, err := utils.GetLastUserClicked(db)
		utils.CatchError(err)
		
		utils.InsertMessage(db, loggedInUser, clickedUser, msg.Message, timeAndDate)
	}
	tmpl.Execute(w, msg)
}

func main() {
	// Set up the route and handler for the form
	http.HandleFunc("/", formHandler)

	fmt.Println("Server started at http://localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}