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
        width: 100%;
        max-width: 400px;
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

    <!DOCTYPE html>
    <html>
    <head>
        <title>Kyle Connect - Send your message to a channel</title>
    </head>
    <body>
        <div class="container">
            <h1>Send a message to a channel</h1>
            <p>Let's send a message to a channel within a server!</p>
            <form method="POST" action="/form">
                <div class="input-group">
                    <label for="chanmsg">Your Message:</label>
                    <input type="text" id="chanmsg" name="chanmsg" placeholder="Enter your message to send" required><br><br>
                </div>
                <input type="submit" value="Send Message">
            </form>

            {{if .ChannelMessage}}
            <p class="success-message">Your message, {{.ChannelMessage}}, has been successfully sent to your server!</p>
            {{end}}
        </div>
    </body>
    </html>
`

type ChannelMessage struct {
    ChannelMessage string
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

		// db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect") // this line of code works for localhost but not docker! MAKE SURE TO COMMENT THIS OUT WHEN WORKING WITH DOCKER!!!!!!!!!!!!!!!!
        db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")

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
		
		chanmsgData.ChannelMessage = r.FormValue("chanmsg")
		channelMessage := chanmsgData.ChannelMessage

        location, _ := time.LoadLocation("Europe/Dublin")
        currentTime := time.Now().In(location)
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
