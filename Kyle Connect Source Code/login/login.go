package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"testing/makeAccount"
	"testing/ui"
	"testing/utils"
)

var (
    sendFriendRequest = "http://localhost:8082/actions/addFriend.go"
    sendMessageToFriend = "http://localhost:8083/sendMessage/sendMessage.go"
    createServer = "http://localhost:8084/createServer/createServer.go"
    sendMessageToServerPlusChannel = "http://localhost:8088/sendMessageToChannelInServer/sendMessageToChannelInServer.go"
    addFriend = "http://localhost:8086/addFriendToServer/addFriendToServer.go"
    removeFriend = "http://localhost:8087/deleteFriendFromServer/deleteFriendFromServer.go"
    createChannel = "http://localhost:8085/createChannel/createChannel.go"

    tasks = []string {
        sendFriendRequest,
        sendMessageToFriend,
        createServer,
        sendMessageToServerPlusChannel,
        addFriend,
        removeFriend,
        createChannel,
    }
)

var account = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kyle Connect - Login</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.1/css/all.min.css">
    <style>
        /* General Styles */
        body {
            font-family: 'Arial', sans-serif;
            background: linear-gradient(to right, #141e30, #243b55);
            color: white;
            text-align: center;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }

        .container {
            background: rgba(255, 255, 255, 0.1);
            padding: 30px;
            border-radius: 12px;
            box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
            width: 100%;
            max-width: 400px;
        }

        h1 {
            margin-bottom: 10px;
            font-size: 28px;
        }

        p {
            margin-bottom: 20px;
            font-size: 16px;
            opacity: 0.8;
        }

        /* Form Styling */
        .input-container {
            display: flex;
            align-items: center;
            background: white;
            padding: 12px;
            border-radius: 8px;
            margin-bottom: 15px;
            transition: 0.3s ease-in-out;
            border: 2px solid transparent;
        }

        .input-container:hover {
            border-color: #007bff;
        }

        .icon {
            margin-right: 10px;
            color: #007bff;
            font-size: 18px;
        }

        input {
            border: none;
            outline: none;
            width: 100%;
            font-size: 16px;
            padding: 5px;
        }

        input:focus {
            border-bottom: 2px solid #007bff;
        }

        /* Submit Button */
        .btn {
            background-color: #007bff;
            color: white;
            padding: 12px;
            border: none;
            border-radius: 8px;
            cursor: pointer;
            width: 100%;
            font-size: 18px;
            font-weight: bold;
            transition: 0.3s;
        }

        .btn:hover {
            background-color: #0056b3;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Login</h1>
        <p>Welcome back to Kyle Connect!</p>
        <form method="POST" action="/form">
            <div class="input-container">
                <i class="fa-solid fa-envelope icon"></i>
                <input type="email" id="email" name="email" placeholder="Enter your email" required>
            </div>
            <div class="input-container">
                <i class="fa-solid fa-lock icon"></i>
                <input type="password" id="password" name="password" placeholder="Enter your password" required>
            </div>
            <button type="submit" class="btn">Login</button>
        </form>
    </div>
</body>
</html>
`

type User struct {
    Email string
	Password string
    UI template.HTML
    ServerName string
    Channels []string
    Channel string
    Name string
}

type Ans struct {
    Answer string
}

func (a Ans) String() string {
    return a.Answer
}

func convertIntToString(id int) string {
	str := strconv.Itoa(id)
    return str
}

func redir(w http.ResponseWriter, newPage string, n string) {   
    data := `
    <html>
    <head>
    <style>
    ul.test {
        list-style-type: none; /* Removes the dots */
        padding: 0;
        margin: 0;
    }

    ul.test li {
        margin: 5px 0;
        padding-left: 15px; /* Moves the link to the right */
    }

    ul.test li a {
        text-decoration: none;
        color: #007BFF;
        font-weight: bold;
    }

    ul.test li a:hover {
        text-decoration: underline;
    }
    </style>
  </head>
  <body>
    <ul class="test">
      <li><a href="%s">%s</a></li>
    </ul>
  </body>
</html>
`

    html := fmt.Sprintf(data, newPage, n)
    w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func KyleConnect(db *sql.DB, w http.ResponseWriter, s string, userData User) {
    data := User {
        Name: s,
        Email: userData.Email,
        UI: template.HTML(ui.UI),
    }

    t, err := template.New("UI").Parse(ui.UI)
    if err != nil {
        http.Error(w, "Template parsing error", http.StatusInternalServerError)
        return
    }

    if err := t.Execute(w, data); err != nil {
        http.Error(w, "Template execution error", http.StatusInternalServerError)
        return
    }

    redir(w, tasks[0], "Add friend (Send Friend Request)")
    redir(w, tasks[2], "Create a server")

    t, err = template.New("UI").Parse(ui.Line)
    if err != nil {
        http.Error(w, "Template parsing error", http.StatusInternalServerError)
        return
    }
    t.Execute(w, t)

    ownerOfServer := utils.GetOwnerOfServer(db)
    for _, n := range ownerOfServer {
        // s is the logged-in user 
        if s == n {
            // Parse the TEST ui/ui.go Admin variable
            t, err = template.New("").Parse(ui.Admin)
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }
            t.Execute(w, nil)

            redir(w, tasks[4], "Add friends to server")
            redir(w, tasks[5], "Delete friends from server")

            t, err = template.New("").Parse(ui.BR)
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }
            t.Execute(w, nil)
        }
    }

var servers = `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <style>
        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }

        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            padding: 20px;
            line-height: 1.6;
            background-color: #a7b1c5;
        }

        .dashboard-container {
            max-width: 1200px;
            margin: 0 auto;
            display: flex;
            flex-direction: column;
            gap: 20px;
        }

        .welcome-banner {
            text-align: center;
            padding: 25px;
            border-radius: 10px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
            background-color: #fff;
            margin-bottom: 20px;
        }

        .content-row {
            display: flex;
            gap: 20px;
            width: 100%;
            flex-wrap: wrap;
        }

        .card {
            flex: 1;
            padding: 1.5rem;
            border-radius: 1rem;
            box-shadow: 0 0 12px 3px rgba(70, 120, 180, 0.8);
            transition: transform 0.3s ease, box-shadow 0.3s ease;
        }

        .card:hover {
            transform: translateY(-10px);
            box-shadow: 0 4px 16px 4px rgba(70, 120, 180, 0.8);
        }

        .card-title {
            font-size: 1.5rem;
            color: #111827;
            margin-bottom: 0.75rem;
            border-bottom: 2px solid #3b82f6;
            padding-bottom: 8px;
        }

        .card-description {
            font-size: 1rem;
            color: #1f2937;
            margin-bottom: 1rem;
        }

        .servers-list {
            display: flex;
            flex-wrap: wrap;
            list-style: none;
            gap: 10px;
            margin-top: 15px;
        }

        .server-link {
            display: inline-block;
            padding: 0.5rem 1rem;
            background-color: #3b82f6;
            color: white;
            border-radius: 0.5rem;
            text-decoration: none;
            transition: background-color 0.2s ease-in-out, transform 0.2s ease-in-out;
        }

        .server-link:hover {
            background-color: #2563eb;
            transform: scale(1.05);
        }

        ul {
            list-style-position: inside;
            margin-top: 10px;
        }

        li {
            margin-bottom: 8px;
        }

        a {
            color: #2563eb;
            text-decoration: none;
        }

        a:hover {
            text-decoration: underline;
        }

        pre {
            padding: 8px;
            border-radius: 4px;
            font-family: inherit;
            white-space: pre-wrap;
        }

        /* Mobile and Tablet Responsive */
        @media (max-width: 900px) {
            .content-row {
                flex-direction: column;
            }

            .card {
                width: 100%;
                margin-bottom: 20px;
            }

            .welcome-banner {
                padding: 15px;
            }

            body {
                padding: 15px;
            }
        }

        @media (max-width: 600px) {
            .card-title {
                font-size: 1.25rem;
            }

            .card-description {
                font-size: 0.9rem;
            }

            .server-link {
                padding: 0.4rem 0.8rem;
            }
        }
        </style>

    </head>
    <body>
    <div class="dashboard-container">
        <div class="content-row">
            <div class="card">
                <h3 class="card-title">Servers</h3>
                <p class="card-description">Below are all of your servers that you have created:</p>
                <ul class="servers-list">
                    {{range .}}
                        <li><a href="/serverClicked/{{.}}" class="server-link">{{.}}</a></li>                                  
                    {{end}}
                </ul>
            </div>
        </div>
`

        // Parse the servers template
        t, err = template.New("").Parse(servers)
        if err != nil {
            http.Error(w, "Template parsing error", http.StatusInternalServerError)
            return
        }
        
        loggedInId, _ := utils.GetLastUserLoggedIn(db)
        converted := convertIntToString(loggedInId)
        emailFromId, _ := utils.RetrieveEmailFromId(db, converted)

        name := utils.RetrieveEmail(db, emailFromId)
        srs := utils.Servers(db, name)

        if err := t.Execute(w, srs); err != nil {
            http.Error(w, "Template execution error", http.StatusInternalServerError)
            return
        }  
                
var serversYouHaveBeenAddedTo = `
        <div class="content-row">
            <div class="card">
                <h3 class="card-title">Servers you have been added to</h3>
                <p class="card-description">Below are all of the servers you have been added to:</p>
                <ul>
                    {{range .}}
                        <li>{{.}}</li>                    
                    {{end}}
                </ul>
            </div>
            `

            // Parse the serversYouHaveBeenAddedTo template
            t, err = template.New("").Parse(serversYouHaveBeenAddedTo)
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }

            addedTo := utils.AddedToServer(db, s)
            if err := t.Execute(w, addedTo); err != nil {
                http.Error(w, "Template execution error", http.StatusInternalServerError)
                return
            }  
            
var Messages = `
        <div class="card">
            <h3 class="card-title">Messages</h3>
            <p class="card-description">Below are all of your messages:</p>
            <ul>
                {{range .}}
                    <pre>{{.}}</pre>                    
                {{end}}
            </ul>
        </div>
`

        // Parse the messagesRecieved template
        t, err = template.New("").Parse(Messages)
        if err != nil {
            http.Error(w, "Template parsing error", http.StatusInternalServerError)
            return
        }

        id, _ := utils.GetLastUserClicked(db)
        messages := utils.GetMessages(db, id)
        if err := t.Execute(w, messages); err != nil {
            http.Error(w, "Template execution error", http.StatusInternalServerError)
            return
        }     

        var Friends = `
        <div class="content-row">
            <div class="card">
                <h3 class="card-title">Friends</h3>
                <p class="card-description">Below are all of your friends:</p>
                <ul>
                    {{range .}}
                        <li><a href="/friend/{{.}}">{{.}}</a></li>                    
                    {{end}}
                </ul>
            </div>
        </div>
        `

        // Parse the friendsHTML template
        t, err = template.New("").Parse(Friends)
        if err != nil {
            http.Error(w, "Template parsing error", http.StatusInternalServerError)
            return
        }

        user := utils.RetrieveEmail(db, userData.Email)
        utils.InsertLoggedInUserToTable(db, user, userData.Email)

        friends := utils.GetFriends(db, user)
        friends2 := utils.GetFriends2(db, user)

        var totalFriends []string
        totalFriends = append(totalFriends, friends...)
        totalFriends = append(totalFriends, friends2...)

        if err := t.Execute(w, totalFriends); err != nil {
            http.Error(w, "Template execution error", http.StatusInternalServerError)
            return
        }


    value, _ := utils.GetToUserName(db, user)

    const status = "pending"
    logged, to, stat := utils.GetPendingRequestsForLoggedInUser(db, user, value, status)
    fmt.Println("value of to variable is:", to, "- (this is a print statement!)")

    if value != "pending" || value == "pending" || stat != "pending" || stat == "pending" {
        var pendingFriendRequest = `
        <div class="card">
            <h3 class="card-title">Pending Friend Requests</h3>
            <p class="card-description">Below are the people who sent you friend requests:</p>                    
            <ul>
                {{range .}}
                    <li>{{.}}</li>
                {{end}}
            </ul>
        </div>
        <br><br><br>
    </body>
    </html>
`

        t, err = template.New("").Parse(pendingFriendRequest)
        if err != nil {
            http.Error(w, "Template parsing error", http.StatusInternalServerError)
            return
        }
        
        many := utils.WhoSentFriendRequest(db, logged)
        if err := t.Execute(w, many); err != nil {
            http.Error(w, "Template execution error", http.StatusInternalServerError)
            return
        }  
        return
    }
}

func GetAnswer(db *sql.DB, w http.ResponseWriter, r *http.Request) Ans { 
    var input = `
    <html>
    <head>
        <style>
            * {
                box-sizing: border-box;
                margin: 0;
                padding: 0;
            }

            body {
                font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
                padding: 20px;
                color: #333;
            }

            .dashboard-container {
                max-width: 600px;
                margin: 0 auto;
                display: flex;
                flex-direction: column;
                gap: 20px;
            }

            .card {
                padding: 1.5rem;
                border-radius: 1rem;
                box-shadow: 0 0 12px 3px rgba(70, 120, 180, 0.8);
            }

            .card-title {
                font-size: 1.5rem;
                color: #111827;
                margin-bottom: 1rem;
                border-bottom: 2px solid #3b82f6;
                padding-bottom: 8px;
            }

            .answer-button {
                background-color: #3b82f6;
                color: white;
                border: none;
                padding: 0.75rem 1.5rem;
                margin: 0.5rem 0;
                border-radius: 0.5rem;
                font-size: 1rem;
                cursor: pointer;
                transition: background-color 0.3s ease;
            }

            .answer-button:hover {
                background-color: #2563eb;
            }

            @media (max-width: 600px) {
                .dashboard-container {
                    width: 90%;
                }

                .card-title {
                    font-size: 1.3rem;
                }

                .answer-button {
                    font-size: 0.9rem;
                    padding: 0.5rem 1rem;
                }
            }
        </style>
    </head>
    <body>
        <div class="dashboard-container">
            <div class="card">
                <h1 class="card-title">Would you like to accept or decline this request?</h1>
                <form method="POST" action="/form">
                    <button class="answer-button" type="submit" name="kyle" value="accept" class="answer-button">Accept</button>
                    <button class="answer-button" type="submit" name="kyle" value="decline" class="answer-button">Decline</button>
                </form>
            </div>
        </div>
    </body>
    </html>
    `

    // Initialize form data
    user := Ans{}

    if r.Method == http.MethodPost {
        // Parse form data
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Error parsing form data", http.StatusBadRequest)
            return user
        }

        y, e := template.New("kyle").Parse(input)
        if e != nil {
            http.Error(w, "Template parsing error", http.StatusInternalServerError)
            return user
        }
        utils.CatchError(e)

        if err := y.Execute(w, "answer"); err != nil {
            http.Error(w, "Template execution error", http.StatusInternalServerError)
            return user
        } 
                
        user.Answer = r.FormValue("kyle")
        return user
    }
    return user
}

// Handler function to serve the form and process submissions
func FormHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the HTML template
    tmpl, err := template.New("form").Parse(account) // might need to do it here!?
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

        // db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect") // this line of code works for localhost but not docker! MAKE SURE TO COMMENT THIS OUT WHEN WORKING WITH DOCKER!!!!!!!!!!!!!!!!
        db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")

        utils.CatchError(err)
        defer db.Close()

        userData.Email = r.FormValue("email")
		userData.Password = r.FormValue("password")

        // the variable "s" is the logged in user:
        s := utils.RetrieveEmail(db, userData.Email)
        fmt.Println("value of s is:......", s)
        AcceptOrDecline(db, w, r, userData, "test") // this string is the logged in user (I am having some issues with it!)

        hashedPassword, err := utils.RetrieveDataFromDb(db, userData.Email)  
        if err != nil {
            if err == sql.ErrNoRows {
                fmt.Println("Email is not in database!")
            } 
            return
        }

        data := User {
            Email: userData.Email,
            UI: template.HTML(ui.UI),
        }

        if makeAccount.CheckPassword(userData.Password, hashedPassword) {
            fmt.Println("You have logged into your account.")
            KyleConnect(db, w, s, userData)
            return // stop showing the login page
        } else {
            t, err := template.New("").Parse(ui.UIERROR) 
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }

            if err := t.Execute(w, data); err != nil {
                http.Error(w, "Template execution error", http.StatusInternalServerError)
                return
            }
            return 
        }
    }     
    // Render the HTML template with the form data
    tmpl.Execute(w, userData)
}

func AcceptOrDecline(db *sql.DB, w http.ResponseWriter, r *http.Request, f User, getLogged string) {
    answer := GetAnswer(db, w, r)
    fmt.Println("ANSWER:", answer)

    a := Ans{Answer: answer.Answer}
    result := a.String()

    user := f.Email
    fmt.Println("user:", user)

    s := utils.RetrieveEmail(db, user)
    fmt.Println("value of s is:", s)

    testlogged := getLogged
    fmt.Println("value of testlogged is:", testlogged)

    friend := utils.LoggedInPossibleFriend(db, testlogged)
    fmt.Println("friend:", friend)
    
    if result == "accept" {
        status := "accept"
        utils.UpdateFriendRequestStatus(db, status, testlogged)

        // here I put user1 and user2 into the friends table
        utils.PutFriendsToFriendsTable(db, testlogged, friend)

        yes, err := template.New(ui.FriendRequestAccepted).Parse(ui.FriendRequestAccepted)
        if err != nil {
            log.Fatal(err)
        }
        yes.Execute(w, user)
        return
    }

    if result == "decline" {
        utils.DeclineFriendRequest(db, testlogged)

        no, err := template.New(ui.FriendRequestDeclined).Parse(ui.FriendRequestDeclined)
        if err != nil {
            log.Fatal(err)
        }
        no.Execute(w, user)
        return
    }  
}

func main() {
    // Set up the route and handler for the form
    http.HandleFunc("/", FormHandler)
    
    // db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect")
    db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")
    utils.CatchError(err)
    defer db.Close()
    
    http.HandleFunc("/friend/", func(w http.ResponseWriter, r *http.Request) {
        friendName := r.URL.Path[len("/friend/"):]
        utils.InsertIntoClickedTable(db, friendName) // i have what friend you clicked on!!!!
    
        selectedFriend := `
		<div class="selection-container">
    		<h1 class="selection-message">You have chosen: <span class="friend-name">%s</span> to chat with!</h1>
		
			<button class="back-button">
        		<a href="http://localhost:8083/sendMessage/sendMessage.go">Begin Conversation</a>
    		</button>

            <button class="back-button">
            	<a href="">Back</a>
    		</button>
		</div>

		<style>
		body {
			background-color: #f0f2f5;
			font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
			margin: 0;
			padding: 0;
			display: flex;
			justify-content: center;
			align-items: center;
			height: 100vh;
		}

		.selection-container {
			text-align: center;
			padding: 40px;
			background-color: #fff;
			border-radius: 12px;
			box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
			max-width: 600px;
		}

		.selection-message {
			font-size: 28px;
			color: #333;
		}

		.friend-name {
			color: #007bff;
			font-weight: bold;
		}

		.back-button {
			padding: 10px 20px;
			background-color: #007bff   ;
			color: white;
			font-size: 16px;
			border-radius: 6px;
			border: none;
			cursor: pointer;
			transition: background-color 0.3s ease;
    	}

		.back-button:hover {
			background-color: #0056b3;
		}

		.back-button {
			text-align: center;
			margin-top: 20px;
		}

		a {
        	color: black;
        	text-decoration: none;
    	}
		</style>
		`

		fmt.Fprintf(w, selectedFriend, friendName)
    })

    var sn string
    var msgInChan []string

    http.HandleFunc("/serverClicked/", func(w http.ResponseWriter, r *http.Request) {
        // this is the server you clicked on:
        serverName := r.URL.Path[len("/serverClicked/"):]
        sn = serverName

        serverIdFromServerName := utils.GetServerIdFromServerName(db, sn)
        fmt.Println("serverIdFromServerName:", serverIdFromServerName)

        lastServerId, _ := utils.GetServerId(db)
        c := convertIntToString(lastServerId)

        for _, n := range serverIdFromServerName {
            if n != c {
                fmt.Println("server id", n, "is not equal to, server id", c)

                t, err := template.New("").Parse(ui.NoChannelsFound)
                if err != nil {
                    http.Error(w, "Template parsing error", http.StatusInternalServerError)
                    return
                }
                t.Execute(w, nil)
                return // dont show the channels!
            }

            if n == c {
                channelsInServer := utils.GetChannelsInServer(db, serverName)
                fmt.Println("channelsInServer.............................................", channelsInServer, "this code has executed:::::::::")    
                
                user := User{
                    ServerName: serverName,
                    Channels: channelsInServer,
                }
        
                // Prepare channelsInServerTemplate template
                var channelsInServerTemplate = `
                <style>

                body {
                    background-color: #a7b1c5;
                }
                    
                .fri {
                    background-color: #f9f9fb;
                    border: 1px solid #e0e0e0;
                    border-radius: 12px;
                    padding: 24px;
                    max-width: 600px;
                    width: 600px;
                    padding: 1.5rem;
                    border-radius: 1rem;
                    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
                    background: transparent !important;
                    box-shadow: 0 0 12px 3px rgba(70, 120, 180, 0.8) !important;
                    
                    /* Centering styles */
                    position: absolute;
                    top: 50%;
                    left: 50%;
                    transform: translate(-50%, -50%);
                }

                .serversYouHaveBeenAddedTo {
                    font-size: 1.4rem;
                    color: #333;
                    margin-bottom: 16px;
                }

                .fri ul {
                    list-style: none;
                    padding-left: 0;
                }

                .fri li {
                    margin-bottom: 10px;
                }

                .fri a {
                    display: inline-block;
                    padding: 10px 16px;
                    background-color: #4f46e5;
                    color: white;
                    text-decoration: none;
                    border-radius: 8px;
                    transition: background-color 0.3s ease;
                }

                .fri a:hover {
                    background-color: #4338ca;
                }

                @media (max-width: 600px) {
                    .serversYouHaveBeenAddedTo {
                        font-size: 1.2rem;
                    }

                    .fri {
                        padding: 1rem;
                        width: 90%; /* Make the width more fluid */
                        margin: 10px;
                    }

                    .fri a {
                        font-size: 0.9rem;
                        padding: 8px 14px;
                    }
                }
                </style>

                <div class="fri">
                    <h3 class="serversYouHaveBeenAddedTo">Here are all of the Channel's that are in: {{.ServerName}}.</h3>
                    <p></p>
                    <ul>
                        {{range .Channels}}
                            <li><a href="/channelClicked/{{.}}">{{.}}</a></li>                    
                        {{end}}
                    </ul>
                </div>
                `
        
                // Parse the channelsInServerTemplate template
                t, err := template.New("channelsInServerTemplate").Parse(channelsInServerTemplate)
                if err != nil {
                    http.Error(w, "Template parsing error", http.StatusInternalServerError)
                    return
                }
        
                if err := t.Execute(w, user); err != nil {
                    http.Error(w, "Template execution error", http.StatusInternalServerError)
                    return
                }  
                redir(w, tasks[6], "Create a Channel")
            }
        }
    })

    http.HandleFunc("/channelClicked/", func(w http.ResponseWriter, r *http.Request) {
        channelName := r.URL.Path[len("/channelClicked/"):]
        fmt.Println("channel name:", channelName)

        // this is the server name you clicked on
        fmt.Println("SN:", sn)

        // gets a server id depending on what you clicked on for (sn above)
        serverIdFromServerName := utils.GetServerIdFromServerName(db, sn)
        fmt.Println("serverIdFromServerName:", serverIdFromServerName)

        lastServerId, _ := utils.GetServerId(db)
        c := convertIntToString(lastServerId)

        for _, n := range serverIdFromServerName {
            if n != c {
                fmt.Println("server id", n, "is not equal to, server id", c)

                uuu := User{
                    Channel: channelName,
                    Channels: msgInChan,
                }
    
                // Prepare serversYouHaveBeenAddedTo template
                var serversYouHaveBeenAddedTo = `
                <div class="fri">
                    <h3 class="serversYouHaveBeenAddedTo">
                        Messages sent to this channel: {{.ServerName}} {{.Channel}}
                    </h3>
                    <ul class="message-list">
                        {{range .Channels}}
                            <li><pre class="message">{{.}}</pre></li>                    
                        {{end}}
                    </ul>
                </div>
                `

                // Parse the serversYouHaveBeenAddedTo template
                t, err := template.New("serversYouHaveBeenAddedTo").Parse(serversYouHaveBeenAddedTo)
                if err != nil {
                    http.Error(w, "Template parsing error", http.StatusInternalServerError)
                    return
                }

                if err := t.Execute(w, uuu); err != nil {
                    http.Error(w, "Template execution error", http.StatusInternalServerError)
                    return
                }  
            }

            num, err := strconv.Atoi(n)
            utils.CatchError(err)
        
            msgInChan := utils.GetMessagesInChannel(db, num)
            if len(msgInChan) == 0 {
                fmt.Println("no messages where found for:", sn, serverIdFromServerName)
            }

            fmt.Println("msgInChan:", msgInChan)
            for _, n2 := range msgInChan {
                fmt.Println("n2:", n2)
            }

            uuu := User{
                Channel: channelName,
                Channels: msgInChan,
            }

            // Prepare serversYouHaveBeenAddedTo template
            var serversYouHaveBeenAddedTo = `
            <style>
            body {
                background-color: #a7b1c5;
            }
            </style>

            <div class="fri">
                <h3 class="serversYouHaveBeenAddedTo">Messages sent to this channel: {{.ServerName}} {{.Channel}}</h3>
                <p></p>
                <ul>
                    {{range .Channels}}
                        <pre>{{.}}</pre>                    
                    {{end}}
                </ul>
            </div>
            `
 
            // Parse the serversYouHaveBeenAddedTo template
            t, err := template.New("serversYouHaveBeenAddedTo").Parse(serversYouHaveBeenAddedTo)
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }

            if err := t.Execute(w, uuu); err != nil {
                http.Error(w, "Template execution error", http.StatusInternalServerError)
                return
            }  
            redir(w, tasks[3], "Send a message to this Channel")
        }
    })

    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
