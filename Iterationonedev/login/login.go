package main

import (
	// "database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "testing/utils"
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

    {{if .Username}}
    <p>You have successfully logged into your account {{.Username}}.</p>
    {{end}}
</body>
</html>
`

type User struct {
    Email string
	Password string
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
        fmt.Println("email:", userData.Email, "password:", userData.Password)

        // passwordHashed := mdHashing(userData.Password)
        // mixedUpPassword := jumbleUpHash(passwordHashed, 26)



        // fmt.Println("password hashed is:", passwordHashed)
        // fmt.Println("jumpled up password is:", mixedUpPassword)
        
        
        
        // db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/users")
        // utils.CatchError(err)

        // utils.PutDataToDb(db, userData.Username, userData.Email, mixedUpPassword)
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