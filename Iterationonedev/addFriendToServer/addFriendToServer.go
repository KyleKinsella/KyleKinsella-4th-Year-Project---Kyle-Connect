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

func convertIntToString(id int) string {
	str := strconv.Itoa(id)
    return str
}

// Handler function to serve the form and process submissions
func formHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
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

		// Prepare friendsHTML template
		var friendsHTML = `
		<div class="fri">
			<h3 class="friends">Friends</h3>
			<p>Below are all of your friends.</p>
			<ul>
				{{range .}}
					<li><a href="/friend/{{.}}">{{.}}</a></li>                    
				{{end}}
			</ul>
		</div>
		`

		// Parse the friendsHTML template
		t, err := template.New("friendsList").Parse(friendsHTML)
		if err != nil {
			http.Error(w, "Template parsing error", http.StatusInternalServerError)
			return
		}

		id, err := utils.GetLastUserLoggedIn(db)
		utils.CatchError(err)
		con := convertIntToString(id)

		email, err := utils.RetrieveEmailFromId(db, con)
		utils.CatchError(err)

		user := utils.RetrieveEmail(db, email)

		friends := utils.GetFriends(db, user)
		if err := t.Execute(w, friends); err != nil {
			http.Error(w, "Template execution error", http.StatusInternalServerError)
			return
		}
		
		friends2 := utils.GetFriends2(db, user)
		if err := t.Execute(w, friends2); err != nil {
			http.Error(w, "Template execution error", http.StatusInternalServerError)
			return
		} 
		 
		// Prepare servers template
		var servers = `
		<div class="fri">
			<h3 class="servers">Servers</h3>
			<p>Below are all of your servers.</p>
			<ul>
				{{range .}}
					<li><a href="/server/{{.}}">{{.}}</a></li>                    
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
		converted := convertIntToString(loggedInId)
	
		emailFromId, err := utils.RetrieveEmailFromId(db, converted)
		
		name := utils.RetrieveEmail(db, emailFromId)

		srs := utils.Servers(db, name)

		if err := t.Execute(w, srs); err != nil {
			http.Error(w, "Template execution error", http.StatusInternalServerError)
			return
		}  
	}
}

func main() {
    // Set up the route and handler for the form
    http.HandleFunc("/", formHandler)

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect")
    // db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")
    utils.CatchError(err)
    defer db.Close()

	var data []string
    
	http.HandleFunc("/friend/", func(w http.ResponseWriter, r *http.Request) {
		friendName := r.URL.Path[len("/friend/"):]
		fmt.Println("friendName is:", friendName)
	
		data = append(data, friendName)
		for _, n := range data {
			fmt.Println("n contains:", n)
		}
	})

	http.HandleFunc("/server/", func(w http.ResponseWriter, r *http.Request) {
        serverName := r.URL.Path[len("/server/"):]
		fmt.Println("serverName is:", serverName)

		peopleInServer := utils.NameOfPeopleInServer(db)
		fmt.Println("here are the list of people in your server", peopleInServer)

		var friendsToAdd string
		var peopleInS string

		for _, friendsToAdd = range data {
			fmt.Println("this is what is in friendsToAdd:", friendsToAdd)
		}

		fmt.Println("\n\n")

		for _, peopleInS = range peopleInServer {
			fmt.Println("this is what is in peopleInS:", peopleInS)
		}

		if friendsToAdd == peopleInS {
			fmt.Println("you cannot add someone that is in the server already!")
			return
		} else {
			for _, n := range data {
				// fmt.Println("this is a test, n:", n)
				utils.AddFriendToServer(db, n, serverName)
			}
		}
    })

    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8086")
    log.Fatal(http.ListenAndServe(":8086", nil))
}
