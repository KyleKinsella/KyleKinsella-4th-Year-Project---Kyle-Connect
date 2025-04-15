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

func convertIntToString(id int) string {
	str := strconv.Itoa(id)
    return str
}

// Handler function to serve the form and process submissions
func formHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {

		t, err := template.New("UI").Parse(ui.Adding)
		if err != nil {
			http.Error(w, "Template parsing error", http.StatusInternalServerError)
			return
		}
		t.Execute(w, err)

        // Parse form data
        err = r.ParseForm()
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
   		<style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }

        .container {
            width: 60%;
            background: white;
            padding: 20px;
            border-radius: 12px;
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
            text-align: center;
            border: 3px solid #0056b3;
        }

        .header {
            font-size: 24px;
            font-weight: bold;
            color: #0056b3;
            border-bottom: 3px solid #0056b3;
            padding-bottom: 10px;
            margin-bottom: 20px;
        }

        .content {
            display: flex;
            justify-content: space-between;
            align-items: flex-start;
            gap: 20px;
			margin: 20px;
        }

        .section {
            flex: 1;
            padding: 15px;
            background: #f8f9fa;
            border-radius: 10px;
            text-align: center;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        }

        .friends-section {
            background: #e3eaf5;
        }

        .servers-section {
            background: #f2e3f5;
        }

        .divider {
            width: 2px;
            background: #0056b3;
            height: auto;
        }

        .section h3 {
            color: #0056b3;
            font-size: 20px;
            font-weight: bold;
        }

		.section ul {
            list-style: none;
            padding: 0;
        }

        .section li {
            background: #d0e1ff;
            margin: 6px 0;
            padding: 10px;
            border-radius: 6px;
            transition: background 0.3s ease, transform 0.2s ease;
        }

        .section a {
            text-decoration: none;
            color: #0056b3;
            font-weight: bold;
        }

        .section li:hover {
            background: #0056b3;
            transform: scale(1.03);
        }

        .section li:hover a {
            color: white;
        }
    	</style>

		<div class="content">
            	<!-- Friends Section -->
            	<div class="section friends-section">    
					<h3 class="friends">Friends</h3>
					<p>Below are all of your friends.</p>
					<ul>
						{{range .}}
							<li><a href="/friend/{{.}}">{{.}}</a></li>                    
						{{end}}
					</ul>
				</div>
		</div>
		`

		// Parse the friendsHTML template
		t, err = template.New("friendsList").Parse(friendsHTML)
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
		friends2 := utils.GetFriends2(db, user)
		
		var totalFriends []string
		totalFriends = append(totalFriends, friends...)
		totalFriends = append(totalFriends, friends2...)

		if err := t.Execute(w, totalFriends); err != nil {
			http.Error(w, "Template execution error", http.StatusInternalServerError)
			return
		}
		 
		// Prepare servers template
		var servers = `
		<div class="content">
			<!-- Friends Section -->
			<div class="section friends-section">    
				<h3 class="servers">Servers</h3>
				<p>Below are all of your servers.</p>
				<ul>
					{{range .}}
						<li><a href="/server/{{.}}">{{.}}</a></li>                    
					{{end}}
				</ul>
			</div>
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
		
		selectedFriend := `
		<div class="selection-container">
    		<h1 class="selection-message">You have selected: <span class="friend-name">%s</span> to add to your server</h1>
		
			<button class="back-button">
        		<a href="http://localhost:8086/addFriendToServer/addFriendToServer.go">Back</a>
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

			w.Header().Set("Content-Type", "text/html")

			t, err := template.New("").Parse(ui.CannotAddThisFriendToThisServer)
			if err != nil {
				http.Error(w, "Template parsing error", http.StatusInternalServerError)
				return
			}
			t.Execute(w, nil)
			return
		} else {
			for _, n := range data {
				// fmt.Println("this is a test, n:", n)
				utils.AddFriendToServer(db, n, serverName)

				if n == data[0] {
					break
				}
				continue
			}
			w.Header().Set("Content-Type", "text/html")

            t, err := template.New("").Parse(ui.FriendsAddedToServer)
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }
			t.Execute(w, err)
		}
    })

    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8086")
    log.Fatal(http.ListenAndServe(":8086", nil))
}
