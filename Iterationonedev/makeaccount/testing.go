// package main

// import (
// 	"net/http"
// 	"io"
// )

// func index(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "hello world!")
// }

// func main() {

// 	http.HandleFunc("/", index)
// 	http.ListenAndServe(":8080", nil)

// }

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"testing/utils"
)

// Template for the HTML form
var formTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kyle Connect - Make an Account</title>
</head>
<body>
    <h1>Make an Account</h1>
    <form method="POST" action="/form">
        <label for="name">Name:</label>
        <input type="text" id="name" name="name" placeholder="Enter your name" required><br><br>

        <label for="email">Email:</label>
        <input type="email" id="email" name="email" placeholder="Enter your email" required><br><br>

  		<label for="password">Password:</label>
        <input type="password" id="password" name="password" placeholder="Enter your password" required><br><br>

        <input type="submit" value="Create Account">
    </form>

    {{if .Username}}
    <h2>Form Data Received:</h2>
    <p>Name: {{.Username}}</p>
    <p>Email: {{.Email}}</p>
	<p>Password: {{.Password}}</p>
    {{end}}
</body>
</html>
`


type FormData struct {
    Username  string
    Email string
	Password string
}

// Handler function to serve the form and process submissions
func formHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the HTML template
    tmpl, err := template.New("form").Parse(formTemplate)
    if err != nil {
        log.Fatal(err)
    }

    // Initialize form data
    formData := FormData{}

    // Check if the request method is POST (form submission)
    if r.Method == http.MethodPost {
        // Parse form data
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Error parsing form data", http.StatusBadRequest)
            return
        }

        // Retrieve values from the form
        formData.Username = r.FormValue("name")
        formData.Email = r.FormValue("email")
		formData.Password = r.FormValue("password")

		fmt.Println(formData.Username, formData.Email, formData.Password)

        db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/users")
        utils.CatchError(err)

        utils.PutDataToDb(db, formData.Username, formData.Email, formData.Password)
    }

    // Render the HTML template with the form data
    tmpl.Execute(w, formData)
}

func main() {
    // Set up the route and handler for the form
    http.HandleFunc("/", formHandler)

    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
