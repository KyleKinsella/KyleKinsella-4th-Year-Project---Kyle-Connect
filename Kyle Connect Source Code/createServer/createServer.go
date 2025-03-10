package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"testing/utils"
)

var server = `
<!DOCTYPE html>
<html>
<head>
    <title>Kyle Connect - Create a Server</title>
</head>
<body>
    <h1>Create a Server</h1>
    <form method="POST" action="/form">
        <label for="server">Server:</label>
        <input type="text" id="server" name="server" placeholder="Enter your server name" required><br><br>

        <input type="submit" value="Create Server">
    </form>
</body>
</html>
`

type Server struct {
    ServerName string
}

func convertIntToString(id int) string {
	str := strconv.Itoa(id)
    return str
}

// Handler function to serve the form and process submissions
func formHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the HTML template
    tmpl, err := template.New("form").Parse(server) // might need to do it here!?
    if err != nil {
        log.Fatal(err)
    }

    // Initialize form data
    serverData := Server{}

    if r.Method == http.MethodPost {
        // Parse form data
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Error parsing form data", http.StatusBadRequest)
            return
        }

        serverData.ServerName = r.FormValue("server")
		serverName := serverData.ServerName
		
        db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect") // this line of code works for localhost but not docker! MAKE SURE TO COMMENT THIS OUT WHEN WORKING WITH DOCKER!!!!!!!!!!!!!!!!
        // db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")

        utils.CatchError(err)
        defer db.Close()

		loggedInId, err := utils.GetLastUserLoggedIn(db)
		utils.CatchError(err)

		converted := convertIntToString(loggedInId)
		
		emailFromId, err := utils.RetrieveEmailFromId(db, converted)
		utils.CatchError(err)

		name := utils.RetrieveEmail(db, emailFromId)

		utils.CreateServer(db, serverName, name)
	}
	tmpl.Execute(w, serverData)
}

func main() {
    // Set up the route and handler for the form
    http.HandleFunc("/", formHandler)

    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8084")
    log.Fatal(http.ListenAndServe(":8084", nil))
}
