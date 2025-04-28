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
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kyle Connect - Make an Account</title>
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
        <h1>Create an Account</h1>
        <p>Join Kyle Connect today!</p>
        <form method="POST" action="/form">
            <div class="input-container">
                <i class="fa-solid fa-user icon"></i>
                <input type="text" id="name" name="name" placeholder="Enter your name" required>
            </div>
            <div class="input-container">
                <i class="fa-solid fa-envelope icon"></i>
                <input type="email" id="email" name="email" placeholder="Enter your email" required>
            </div>
            <div class="input-container">
                <i class="fa-solid fa-lock icon"></i>
                <input type="password" id="password" name="password" placeholder="Enter your password" required>
            </div>
            <button type="submit" class="btn">Sign Up</button>
        </form>
    </div>
</body>
</html>
`

type User struct {
    Username  string
    Email string
	Password string
}

// Handler function to serve the form and process submissions
func FormHandler(w http.ResponseWriter, r *http.Request) {
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

        db, _ := sql.Open("mysql", "root@tcp(127.0.0.1)/kyleconnect") // this line of code works for localhost but not docker!
        // db, err := sql.Open("mysql", "root@tcp(host.docker.internal:3306)/kyleconnect?parseTime=true")

        usernames := utils.GetCommunicatorsUsernames(db)
        fmt.Println("usernames:", usernames)

        username := User{
            Username: userData.Username,
        }

        if usernames == nil {
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

        for _, n := range usernames {
            if n == userData.Username {
                fmt.Println("this username is already taken, you cannot make anaccount with this name, try again with a new username")
                
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
    http.HandleFunc("/", FormHandler)

    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}