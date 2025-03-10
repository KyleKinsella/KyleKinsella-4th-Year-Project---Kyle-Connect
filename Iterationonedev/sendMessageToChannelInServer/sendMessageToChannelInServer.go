package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"testing/utils"
	"time"
	"strconv"
)

var channelMessage = `
<!DOCTYPE html>
<html>
<head>
    <title>Kyle Connect - Send your message to a channel</title>
</head>
<body>
    <h1>Send a message to a channel</h1>
    <form method="POST" action="/form">
        <label for="chanmsg">Your Message:</label>
        <input type="text" id="chanmsg" name="chanmsg" placeholder="Enter your message to send" required><br><br>

        <input type="submit" value="Send Message">
    </form>
</body>
</html>
`

type ChannelMessage struct {
    channelMessage string
}

func convertIntToString(number int) string {
	str := strconv.Itoa(number)
    return str
}

// Handler function to serve the form and process submissions
func formHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the HTML template
    tmpl, err := template.New("form").Parse(channelMessage) // might need to do it here!?
    if err != nil {
        log.Fatal(err)
    }

    // Initialize form data
    chanmsgData := ChannelMessage{}

    if r.Method == http.MethodPost {
        // Parse form data
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Error parsing form data", http.StatusBadRequest)
            return
        }

		db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect") // this line of code works for localhost but not docker! MAKE SURE TO COMMENT THIS OUT WHEN WORKING WITH DOCKER!!!!!!!!!!!!!!!!
        // db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")

        utils.CatchError(err)
        defer db.Close()

		id, err := utils.GetLastUserLoggedIn(db)
		utils.CatchError(err)
		converted := convertIntToString(id)

		email, err := utils.RetrieveEmailFromId(db, converted)
		utils.CatchError(err)

		username := utils.RetrieveEmail(db, email)

		serverId, err := utils.GetServerId(db)
		utils.CatchError(err)
		
		chanmsgData.channelMessage = r.FormValue("chanmsg")
		channelMessage := chanmsgData.channelMessage

		currentTime := time.Now()
		timeAndDate := currentTime.Format("2006-01-02 15:04:05")

		utils.SendMessageToChannel(db, username, serverId, channelMessage, timeAndDate)
	}
	tmpl.Execute(w, chanmsgData)
}

func main() {
    // Set up the route and handler for the form
    http.HandleFunc("/", formHandler)

    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8088")
    log.Fatal(http.ListenAndServe(":8088", nil))
}
