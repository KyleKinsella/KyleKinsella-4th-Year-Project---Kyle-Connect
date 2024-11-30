package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"html/template"
	"log"
	// below are libaries to convert from an int to a string
	"strconv"
)

var genRandNum int // ik this is bad but i dont care lol......

var validate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kyle Connect - Authentication System</title>
</head>
<body>
    <h1>Let's verify that it is you</h1>
    <form method="POST" action="/form">
		<h4> Random number is: {{.RandomNumber}}</h4>
        <label for="RandomNumber">Enter the random number displayed on screen: </label>
        <input type="number" id="RandomNumber" name="RandomNumber" placeholder="Enter Random Number" required><br><br>

        <input type="submit" value="Authenticate">
    </form>
</body>
</html>
`

var accessTrue = `    
	<h1>Access Granted</h1>
`

var accessFalse = `    
	<h1>Access Denied</h1>
`

type User struct {
	RandomNumber int // number entered from the form
	RandNum string // number converted to a string
	UI template.HTML
}

func convertIntToString(randomNum int) string {
	return strconv.Itoa(randomNum)
}

func generateRandomNumber() int {
	genRandNum = rand.Intn(500)
	fmt.Println("random number is:", genRandNum)
    return genRandNum
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the HTML template
    tmpl, err := template.New("form").Parse(validate)
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
		userData.RandNum = r.FormValue("RandomNumber")
		userData.RandomNumber = genRandNum
        conversion := convertIntToString(userData.RandomNumber)

		if userData.RandNum == conversion {
			// access given 
			dataTrue := User {
                UI: template.HTML(accessTrue),
            }
            // Parse and execute the template
            t, err := template.New("account").Parse(accessTrue)
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }
            if err := t.Execute(w, dataTrue); err != nil {
                http.Error(w, "Template execution error", http.StatusInternalServerError)
                return 
            }
            return

		} else {
			// access denied 
			dataFalse := User {
                UI: template.HTML(accessFalse),
            }
            // Parse and execute the template
            t, err := template.New("account").Parse(accessFalse)
            if err != nil {
                http.Error(w, "Template parsing error", http.StatusInternalServerError)
                return
            }
            if err := t.Execute(w, dataFalse); err != nil {
                http.Error(w, "Template execution error", http.StatusInternalServerError)
                return
            }
            return
		}
    }
    userData.RandomNumber = genRandNum

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

func main() {
    fmt.Printf("genRandNum updated to: %d\n", genRandNum)
    generateRandomNumber()
    RunServer()
}