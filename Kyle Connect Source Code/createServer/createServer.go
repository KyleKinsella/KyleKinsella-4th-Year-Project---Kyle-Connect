package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"testing/utils"
    "testing/ui"
)

var server = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kyle Connect - Create a Server</title>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600&display=swap" rel="stylesheet">
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

        input[type="text"] {
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

        input[type="text"]:focus {
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
</head>
<body>
    <div class="container">
        <h1>Create a Server</h1>
        <p>Create a server, add your friends, and start communicating seamlessly!</p>
        <form method="POST" action="/form">
            <div class="input-group">
                <label for="server">Server Name:</label>
                <input type="text" id="server" name="server" placeholder="Enter your server name" required>
            </div>
            <input type="submit" value="Create Server">
        </form>
        
        {{if .ServerName}}
        <p class="success-message">Your server, {{.ServerName}}, has been successfully created!</p>
        {{end}}
    </div>
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
		
        // db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect") // this line of code works for localhost but not docker! MAKE SURE TO COMMENT THIS OUT WHEN WORKING WITH DOCKER!!!!!!!!!!!!!!!!
        db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")

        utils.CatchError(err)
        defer db.Close()

		loggedInId, err := utils.GetLastUserLoggedIn(db)
		utils.CatchError(err)
		converted := convertIntToString(loggedInId)

		emailFromId, err := utils.RetrieveEmailFromId(db, converted)
		utils.CatchError(err)
		name := utils.RetrieveEmail(db, emailFromId)

        namesOfServers := utils.Servers(db, name)
        for _, n := range namesOfServers {
            if n == serverData.ServerName {
                fmt.Println("you cannot create another server with this server name:", n)

                t, err := template.New("").Parse(ui.CannotUseThisServerName)
                if err != nil {
                    http.Error(w, "Template parsing error", http.StatusInternalServerError)
                    return
                }
                t.Execute(w, nil)
                return
            }
        }
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
