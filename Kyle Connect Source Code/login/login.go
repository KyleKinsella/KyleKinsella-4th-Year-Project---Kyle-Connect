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

    tasks = []string {
        sendFriendRequest,
        sendMessageToFriend,
        createServer,
        sendMessageToServerPlusChannel,
        addFriend,
        removeFriend,
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

// func GetAnswer(db *sql.DB, w http.ResponseWriter, r *http.Request) Ans { 
//     var input = `
//     <html>
//     <head>
//     <style>
//         /* General Body Styling */
//         body {
//             font-family: 'Arial', sans-serif;
//             background-color: #f4f7fa;
//             color: #333;
//             margin: 0;
//             padding: 0;
//         }

//         /* Container for the Form */
//         .kyle {
//             max-width: 500px;
//             margin: 2rem auto;
//             padding: 2rem;
//             background-color: #ffffff;
//             border-radius: 8px;
//             box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
//         }

//         /* Heading Style */
//         .kyle h1 {
//             font-size: 1.75rem;
//             color: #111827;
//             margin-bottom: 1.5rem;
//             text-align: center;
//         }

//         /* Label Styling */
//         label {
//             font-size: 1rem;
//             color: #555;
//             margin-bottom: 0.5rem;
//             display: block;
//         }

//         /* Input Field Styling */
//         input[type="text"] {
//             width: 100%;
//             padding: 0.75rem;
//             margin: 0.5rem 0 1.5rem 0;
//             border: 1px solid #d1d5db;
//             border-radius: 4px;
//             font-size: 1rem;
//             box-sizing: border-box;
//         }

//         /* Submit Button Styling */
//         input[type="submit"] {
//             background-color: #3b82f6;
//             color: white;
//             border: none;
//             padding: 0.75rem 1.5rem;
//             border-radius: 4px;
//             font-size: 1rem;
//             cursor: pointer;
//             transition: background-color 0.3s ease;
//         }

//         /* Submit Button Hover Effect */
//         input[type="submit"]:hover {
//             background-color: #2563eb;
//         }

//         /* For Small Adjustments on Mobile */
//         @media (max-width: 600px) {
//             .kyle {
//                 padding: 1rem;
//                 width: 90%;
//             }

//             .kyle h1 {
//                 font-size: 1.5rem;
//             }

//             input[type="text"] {
//                 font-size: 0.9rem;
//             }

//             input[type="submit"] {
//                 font-size: 0.9rem;
//             }
//         }
//     </style>
// </head>
// <body>

//     <div class="kyle">
//         <h1>Enter accept or decline for your answer</h1>
//         <form method="POST" action="/form">    
//             <label for="kyle">Answer:</label>
//             <input type="text" id="kyle" name="kyle" placeholder="Enter your answer:" required><br><br>

//             <input type="submit" value="Submit Answer">
//         </form>
//     </div>
// </body>
// </html>
// `

func GetAnswer(db *sql.DB, w http.ResponseWriter, r *http.Request) Ans { 
    var input = `
    <html>
    <head>
    <style>
        body {
            font-family: 'Arial', sans-serif;
            background-color: #f4f7fa;
            color: #333;
            margin: 0;
            padding: 0;
        }

        .kyle {
            text-align: center;
            width: 400px;
            padding: 1.5rem;
            border-radius: 1rem;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            margin: 10 auto;
            background: transparent !important;
            box-shadow: 0 0 12px 3px rgba(70, 120, 180, 0.8) !important;
        }

        .kyle h1 {
            font-size: 1.75rem;
            color: #111827;
            margin-bottom: 1.5rem;
        }

        .answer-button {
            background-color: #3b82f6;
            color: white;
            border: none;
            padding: 0.75rem 1.5rem;
            margin: 0.5rem;
            border-radius: 4px;
            font-size: 1rem;
            cursor: pointer;
            transition: background-color 0.3s ease;
        }

        .answer-button:hover {
            background-color: #2563eb;
        }

        @media (max-width: 600px) {
            .kyle {
                padding: 1rem;
                width: 90%;
            }

            .kyle h1 {
                font-size: 1.5rem;
            }

            .answer-button {
                font-size: 0.9rem;
                padding: 0.5rem 1rem;
            }
        }
    </style>
    </head>
    <body>

    <div class="kyle">
        <h1>Please choose your answer</h1>
        <form method="POST" action="/form">
            <button type="submit" name="kyle" value="accept" class="answer-button">Accept</button>
            <button type="submit" name="kyle" value="decline" class="answer-button">Decline</button>
        </form>
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

        userData.Email = r.FormValue("email")
		userData.Password = r.FormValue("password")

        db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect") // this line of code works for localhost but not docker! MAKE SURE TO COMMENT THIS OUT WHEN WORKING WITH DOCKER!!!!!!!!!!!!!!!!
        // db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")

        utils.CatchError(err)
        defer db.Close()

        s := utils.RetrieveEmail(db, userData.Email)
        fmt.Println("value of s is:", s)
        AcceptOrDecline(db, w, r, userData, "Martin") // this string is the logged in user (I am having some issues with it!)

        hashedPassword, err := utils.RetrieveDataFromDb(db, userData.Email)  
        if err != nil {
            if err == sql.ErrNoRows {
                fmt.Println("Email is not in database!")
            } //else {
                // fmt.Println("Error retrieving password from database: / this is the or one of the problems", err)
                // return
            // }
            return
        }

        data := User {
            Email: userData.Email,
            UI: template.HTML(ui.UI),
        }

        if makeAccount.CheckPassword(userData.Password, hashedPassword) {
            fmt.Println("You have logged into your account.")
            
            // Parse and execute the template
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
            redir(w, tasks[1], "Send message to a friend")
            redir(w, tasks[2], "Create a server")

            t, err = template.New("UI").Parse(ui.Line)
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }
            t.Execute(w, t)

            // Prepare servers template
            var servers = `
            <style>

            body {
                background-color: #a7b1c5;
            }

            .servers-container {
                width: 400px;
                padding: 1.5rem;
                border-radius: 1rem;
                font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
                margin: 10px;
                background: transparent !important;
                box-shadow: 0 0 12px 3px rgba(70, 120, 180, 0.8) !important;
            }

            .servers-title {
                font-size: 1.75rem;
                color: #111827;
                margin-bottom: 0.5rem;
            }

            .servers-description {
                font-size: 1rem;
                color: black;
                margin-bottom: 1.25rem;
            }

            .servers-list {
                display: flex;
                list-style: none;
                padding: 0;
                margin: 0;
                gap: 10px;
            }

            .servers-list li {
                margin-bottom: 0.75rem;
            }

            .server-link {
                display: inline-block;
                padding: 0.5rem 1rem;
                background-color: #3b82f6;
                color: #fff;
                border-radius: 0.5rem;
                text-decoration: none;
                transition: background-color 0.2s ease-in-out;
            }

            .server-link:hover {
                background-color: #2563eb;
            }

            </style>

            <div class="servers-container">
                <h3 class="servers-title">Servers</h3>
                <p class="servers-description">Below are all of your servers that you have created:</p>
                <ul class="servers-list">
                    {{range .}}
                        <li><a href="/serverClicked/{{.}}" class="server-link">{{.}}</a></li>                                  
                    {{end}}
                </ul>
            </div>
            `
            
            // Parse the servers template
            t, err = template.New("servers").Parse(servers)
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }

            loggedInId, err := utils.GetLastUserLoggedIn(db)
            // utils.CatchError(err)
            
            converted := convertIntToString(loggedInId)
		
            emailFromId, err := utils.RetrieveEmailFromId(db, converted)
            // utils.CatchError(err)
            
            name := utils.RetrieveEmail(db, emailFromId)

            srs := utils.Servers(db, name)

            if err := t.Execute(w, srs); err != nil {
                http.Error(w, "Template execution error", http.StatusInternalServerError)
                return
            }  

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

                    redir(w, tasks[4], "Add friends to server")
                    redir(w, tasks[5], "Delete friends from server")

                    t.Execute(w, err)
                }
            }

            ////////////////////////////////////////////////////////////////////////////////////////////////////
            // peopleWhoOwnAServer := utils.GetOwnerOfServer(db)
            // fmt.Println("peopleWhoOwnAServer:", peopleWhoOwnAServer)

            // peopleAddedToServer := utils.FriendsAddedToServer(db)
            // fmt.Println("peopleAddedToServer:", peopleAddedToServer)

            // for _, n := range peopleWhoOwnAServer {
            //     if s == n {
            //         utils.ChannelMessagesBeingParsed(db, w)
            //     }
            // }

            // for _, n := range peopleAddedToServer {
            //     if s == n {
            //         utils.ChannelMessagesBeingParsed(db, w)
            //     }
            // }
            ////////////////////////////////////////////////////////////////////////////////////////////////////

            // Prepare serversYouHaveBeenAddedTo template
            var serversYouHaveBeenAddedTo = `
            <style>
            .servers-container {
                max-width: 600px;
                margin: 2rem auto;
                padding: 1.5rem;
                background-color: #f9fafb;
                border-radius: 1rem;
                box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
                font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            }

            .servers-title {
                font-size: 1.75rem;
                color: #111827;
                margin-bottom: 0.5rem;
            }

            .servers-description {
                font-size: 1rem;
                color: #6b7280;
                margin-bottom: 1.25rem;
            }

            .servers-list {
                list-style: none;
                padding: 0;
                margin: 0;
            }

            .servers-list li {
                margin-bottom: 0.75rem;
            }

            .server-link {
                display: inline-block;
                padding: 0.5rem 1rem;
                background-color: #3b82f6;
                color: #fff;
                border-radius: 0.5rem;
                text-decoration: none;
                transition: background-color 0.2s ease-in-out;
            }

            .server-link:hover {
                background-color: #2563eb;
            }

            </style>

            <div class="servers-container">
                <h3 class="servers-title">Servers Linked to Your Profile</h3>
                <p>Below are all of the servers you have been added to:</p>
                <ul>
                    {{range .}}
                        <li>{{.}}</li>                    
                    {{end}}
                </ul>
            </div>
            `

            // Parse the serversYouHaveBeenAddedTo template
            t, err = template.New("serversYouHaveBeenAddedTo").Parse(serversYouHaveBeenAddedTo)
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }

            addedTo := utils.AddedToServer(db, s) // come back to the hard-coded bit later on!!!!!!! (FIXED THIS ON: 01/04/2025)
            if err := t.Execute(w, addedTo); err != nil {
                http.Error(w, "Template execution error", http.StatusInternalServerError)
                return
            }  

            // Prepare messagesRecieved template
            var messagesRecieved = `
            <style>

            .servers-container {
                max-width: 600px;
                margin: 2rem auto;
                padding: 1.5rem;
                background-color: #f9fafb;
                border-radius: 1rem;
                box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
                font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            }

            .servers-title {
                font-size: 1.75rem;
                color: #111827;
                margin-bottom: 0.5rem;
            }

            .servers-description {
                font-size: 1rem;
                color: #6b7280;
                margin-bottom: 1.25rem;
            }

            .servers-list {
                list-style: none;
                padding: 0;
                margin: 0;
            }

            .servers-list li {
                margin-bottom: 0.75rem;
            }

            .server-link {
                display: inline-block;
                padding: 0.5rem 1rem;
                background-color: #3b82f6;
                color: #fff;
                border-radius: 0.5rem;
                text-decoration: none;
                transition: background-color 0.2s ease-in-out;
            }

            .server-link:hover {
                background-color: #2563eb;
            }

            </style>

            <div class="servers-container">
                <h3 class="servers-title">Messages</h3>
                <p>Below are all of your messages:</p>
                <ul>
                    {{range .}}
                        <pre>{{.}}</pre>                    
                    {{end}}
                </ul>
            </div>
            `

            // Parse the messagesRecieved template
            t, err = template.New("message").Parse(messagesRecieved)
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }

            id, err := utils.GetLastUserClicked(db)
            messages := utils.GetMessages(db, id)
            if err := t.Execute(w, messages); err != nil {
                http.Error(w, "Template execution error", http.StatusInternalServerError)
                return
            }  

            // Prepare friendsHTML template
            var friendsHTML = `
            <style>

            .servers-container {
                max-width: 600px;
                margin: 2rem auto;
                padding: 1.5rem;
                background-color: #f9fafb;
                border-radius: 1rem;
                box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
                font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            }

            .servers-title {
                font-size: 1.75rem;
                color: #111827;
                margin-bottom: 0.5rem;
            }

            .servers-description {
                font-size: 1rem;
                color: #6b7280;
                margin-bottom: 1.25rem;
            }

            .servers-list {
                list-style: none;
                padding: 0;
                margin: 0;
            }

            .servers-list li {
                margin-bottom: 0.75rem;
            }

            .server-link {
                display: inline-block;
                padding: 0.5rem 1rem;
                background-color: #3b82f6;
                color: #fff;
                border-radius: 0.5rem;
                text-decoration: none;
                transition: background-color 0.2s ease-in-out;
            }

            .server-link:hover {
                background-color: #2563eb;
            }

            </style>

            <div class="servers-container">
                <h3 class="servers-title">Friends</h3>
                <p class="friendsLi">Below are all of your friends:</p>
                <ul>
                    {{range .}}
                        <li><a href="/friend/{{.}}">{{.}}</a></li>                    
                    {{end}}
                </ul>
            </div>
            `

            // Parse the friendsHTML template
            t, err = template.New("friendsList").Parse(friendsHTML)
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }

            user := utils.RetrieveEmail(db, userData.Email)
            // if er != nil {
            //     log.Fatal(er)
            // }
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

            value, err := utils.GetToUserName(db, user)

            const status = "pending"
            logged, to, stat := utils.GetPendingRequestsForLoggedInUser(db, user, value, status)
            fmt.Println("value of to variable is:", to, "- (this is a print statement!)")

            if value != "pending" || value == "pending" || stat != "pending" || stat == "pending" {
                var showData = `

                <style>
                .servers-container {
                    max-width: 600px;
                    margin: 2rem auto;
                    padding: 1.5rem;
                    background-color: #f9fafb;
                    border-radius: 1rem;
                    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
                    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
                }

                .servers-title {
                    font-size: 1.75rem;
                    color: #111827;
                    margin-bottom: 0.5rem;
                }

                .servers-description {
                    font-size: 1rem;
                    color: #6b7280;
                    margin-bottom: 1.25rem;
                }

                .servers-list {
                    list-style: none;
                    padding: 0;
                    margin: 0;
                }

                .servers-list li {
                    margin-bottom: 0.75rem;
                }

                .server-link {
                    display: inline-block;
                    padding: 0.5rem 1rem;
                    background-color: #3b82f6;
                    color: #fff;
                    border-radius: 0.5rem;
                    text-decoration: none;
                    transition: background-color 0.2s ease-in-out;
                }

                .server-link:hover {
                    background-color: #2563eb;
                }

                </style>

                <div class="servers-container">
                    <h3 class="servers-title">Pending Friend Requests</h3>
                    <p class="friendsLi">Below are the people who sent you friend requests:</p>                    
                    <ul>
                        {{range .}}
                            <li>{{.}}</li> 
                        {{end}}
                    </ul>
                `

                t, err = template.New("showData").Parse(showData)
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
            return // stop the rendering of the login page
        } else {
            t, err := template.New(ui.UIERROR).Parse(ui.UIERROR) 
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
    // getLogged = utils.RetrieveEmail(db, user)
    // fmt.Println("testlogged here is:", testlogged)
    // testlogged = getLogged

    friend := utils.LoggedInPossibleFriend(db, testlogged)
    fmt.Println("friend:", friend)
    
    if result == "accept" {
        status := "accept"
        utils.UpdateFriendRequestStatus(db, status, testlogged) // was testname now -> testlogged

        // here I put user1 and user2 into the friends table
        utils.PutFriendsToFriendsTable(db, testlogged, friend) // was testlogged now -> testname

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
    
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect")
    // db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")
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
        serverName := r.URL.Path[len("/serverClicked/"):]
        // channelsInServer := utils.GetChannelsInServer(db, serverName)
        // fmt.Println("channelsInServer.............................................", channelsInServer)

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
            }
        }
    })

    http.HandleFunc("/channelClicked/", func(w http.ResponseWriter, r *http.Request) {
        channelName := r.URL.Path[len("/channelClicked/"):]
        fmt.Println("channel name:", channelName)

        // this is the server name you clicked on
        fmt.Println("SN:", sn)

        

        // if channelName != sn {
        //     fmt.Println(channelName, "is not in ", sn)
        // } else {
        //     fmt.Println(channelName, "is in ", sn)
        // }
        
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
                fmt.Println("n2:", n2, "\n")
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
