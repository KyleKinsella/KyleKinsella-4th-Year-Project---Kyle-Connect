package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"testing/utils"
)

var channel = `
<!DOCTYPE html>
<html>
<head>
    <title>Kyle Connect - Create a Channel</title>
</head>
<body>
    <h1>Create a Channel</h1>
    <form method="POST" action="/form">
        <label for="channel">Channel:</label>
        <input type="text" id="channel" name="channel" placeholder="Enter your channel name" required><br><br>

        <input type="submit" value="Create Channel">
    </form>
</body>
</html>
`

type Channel struct {
    ChannelName string
}

// Handler function to serve the form and process submissions
func formHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the HTML template
    tmpl, err := template.New("form").Parse(channel) // might need to do it here!?
    if err != nil {
        log.Fatal(err)
    }

    // Initialize form data
    channelData := Channel{}

    if r.Method == http.MethodPost {
        // Parse form data
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Error parsing form data", http.StatusBadRequest)
            return
        }

        channelData.ChannelName = r.FormValue("channel")
		channelName := channelData.ChannelName

        db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect") // this line of code works for localhost but not docker! MAKE SURE TO COMMENT THIS OUT WHEN WORKING WITH DOCKER!!!!!!!!!!!!!!!!
        // db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")

        utils.CatchError(err)
        defer db.Close()

		serverId, err := utils.GetServerId(db)
		utils.CatchError(err)

		utils.InsertChannelName(db, channelName, serverId)
	}
	tmpl.Execute(w, channelData)
}

func main() {
    // Set up the route and handler for the form
    http.HandleFunc("/", formHandler)

    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8085")
    log.Fatal(http.ListenAndServe(":8085", nil))
}
