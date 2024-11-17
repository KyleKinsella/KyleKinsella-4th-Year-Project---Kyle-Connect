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

        db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect")
        utils.CatchError(err)
        defer db.Close()

        hashedPassword, err := utils.RetrieveDataFromDb(db, userData.Email)  
        if err != nil {
            if err == sql.ErrNoRows {
                fmt.Println("Email is not in database!")
            } else {
                fmt.Println("Error retrieving password from database:", err)
            }
            return
        }

        data := User {
            Email: userData.Email,
            UI: template.HTML(ui.UI),
        }

        if makeAccount.CheckPassword(userData.Password, hashedPassword) {
            fmt.Println("You have logged into your account.")
            // Parse and execute the template
            t, err := template.New(ui.UI).Parse(ui.UI)
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }
            if err := t.Execute(w, data); err != nil {
                http.Error(w, "Template execution error", http.StatusInternalServerError)
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

func main() {
    // Set up the route and handler for the form
    http.HandleFunc("/", formHandler)

    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
