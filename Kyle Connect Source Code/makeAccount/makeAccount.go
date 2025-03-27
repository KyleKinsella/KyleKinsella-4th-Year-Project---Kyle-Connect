package makeAccount

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"testing/ui"
	"testing/utils"

	// below are imports for the encryption
	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(input string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hash), nil
}

func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

var account = `
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
</body>
</html>
`

type User struct {
    Username  string
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

        // get the data from the html form
        userData.Username = r.FormValue("name")
        userData.Email = r.FormValue("email")
		userData.Password = r.FormValue("password")    
        // hash the password
        passwordHashed, err := HashingPassword(userData.Password)
        if err != nil {
            fmt.Println("an error has occured!")
            return 
        }

        db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect") // this line of code works for localhost but not docker!
        // db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")
        utils.CatchError(err)

        usernames := utils.GetCommunicatorsUsernames(db)
        for _, n := range usernames {
            if n == userData.Username {
                // fmt.Println("this username is already taken, you cannot make anaccount with this name, try again with a new username")
                
                w.Header().Set("Content-Type", "text/html")

                t, err := template.New("TEST").Parse(ui.ThisUserNameIsAlreadyTaken)
                if err != nil {
                    http.Error(w, "Template parsing error", http.StatusInternalServerError)
                    return
                }
                t.Execute(w, nil)
                return
            } 
        }

        username := User{
            Username: userData.Username,
        }

        for _, n := range usernames {
            if n != userData.Username {
                utils.PutDataToDb(db, userData.Username, userData.Email, passwordHashed)

                // Parse and execute the template
                t, err := template.New("").Parse(ui.YourAccountHasBeenMade)
                if err != nil {
                    http.Error(w, "Template parsing error", http.StatusInternalServerError)
                    return
                }
                
                if err := t.Execute(w, username); err != nil {
                    http.Error(w, "Template execution error", http.StatusInternalServerError)
                    return
                }
                return
            }
        }
    }

    // Render the HTML template with the form data
    tmpl.Execute(w, userData)
}

func RunServer() {
     // Set up the route and handler for the form
    http.HandleFunc("/", formHandler)
    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}