package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"testing/makeAccount"
	"testing/ui"
	"testing/utils"
)

var account = `
<!DOCTYPE html>
<html>
<head>
    <title>Kyle Connect - Login</title>
</head>
<body>
    <h1>Login</h1>
    <form method="POST" action="/form">
        <label for="email">Email:</label>
        <input type="email" id="email" name="email" placeholder="Enter your email" required><br><br>

  		<label for="password">Password:</label>
        <input type="password" id="password" name="password" placeholder="Enter your password" required><br><br>

        <input type="submit" value="Login">
    </form>
</body>
</html>
`

type User struct {
    Email string
	Password string
    UI template.HTML
}

type Ans struct {
    Answer string
}

func (a Ans) String() string {
    return a.Answer
}

func GetAnswer(w http.ResponseWriter, r *http.Request) Ans {
    var input = `
    <div class="kyle">
        <h1>Enter accept or decline for your answer</h1>
        <form method="POST" action="/answer">    
            <label for="kyle">Answer:</label>
            <input type="text" id="kyle" name="kyle" placeholder="Enter your answer:" required><br><br>

            <input type="submit" value="Submit Answer">
        </form>
    </div>
    `
    
    // Parse the HTML template
    u, err := template.New("answer").Parse(input)
    if err != nil {
        log.Fatal(err)
    }
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
    u.Execute(w, user)
    return user
}

// Handler function to serve the form and process submissions
func formHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the HTML template
    tmpl, err := template.New("form").Parse(account)
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

        db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect") // this line of code works for localhost but not docker!
        // // db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")

        utils.CatchError(err)
        defer db.Close()


        AcceptOrDecline(db, w, r, userData, "dockerr") // this string is the logged in user (I am having some issues with it!)

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

            // Prepare friendsHTML template
            var friendsHTML = `
            <div class="fri">
                <h3 class="friends">Friends</h3>
                <p>Below are all of your friends.</p>
                <ul>
                    {{range .}}
                        <li>{{.}}</li>
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
            if err := t.Execute(w, friends); err != nil {
                http.Error(w, "Template execution error", http.StatusInternalServerError)
                return
            }

            value, err := utils.GetToUserName(db, user)

            const status = "pending"
            logged, to, stat := utils.GetPendingRequestsForLoggedInUser(db, user, value, status)
            fmt.Println("value of to variable is:", to, "- (this is a print statement!)")

            if value != "pending" || value == "pending" || stat != "pending" || stat == "pending" {
                var showData = `
                <h3 class="showData">Pending Friend Requests</h3>
                    <p>Below are the people who sent you friend requests.</p>                    
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




            // if stat == "pending" || value != "pending" || stat != "pending" {
            //     // show the user who sent the friend request 
            //     // and show some html to the user to either accept or decline it
            //     var showData = `
            //         <h3 class="showData">Pending Friend Requests</h3>
            //         <p>Below are the people who sent you friend requests.</p>                    
            //         <ul>
            //             {{range .}}
            //                 <li>{{.}}</li> 
            //             {{end}}
            //         </ul>
            //     `

            //     t, err = template.New("showData").Parse(showData)
            //     if err != nil {
            //         http.Error(w, "Template parsing error", http.StatusInternalServerError)
            //         return
            //     }
                
            //     many := utils.WhoSentFriendRequest(db, logged)
            //     if err := t.Execute(w, many); err != nil {
            //         http.Error(w, "Template execution error", http.StatusInternalServerError)
            //         return
            //     }     
            //     return
            // } else {
            //     // do something else....
            // }  

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
    answer := GetAnswer(w, r)

    a := Ans{Answer: answer.Answer}
    result := a.String()

    user := f.Email

    testlogged := getLogged
    testlogged = utils.RetrieveEmail(db, user)
    testlogged = getLogged

    friend := utils.LoggedInPossibleFriend(db, testlogged)

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
    }

    if result == "decline" {
        utils.DeclineFriendRequest(db, testlogged)

        no, err := template.New(ui.FriendRequestDeclined).Parse(ui.FriendRequestDeclined)
        if err != nil {
            log.Fatal(err)
        }
        no.Execute(w, user)
    }  
}

func main() {
    // Set up the route and handler for the form
    http.HandleFunc("/", formHandler)

    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
