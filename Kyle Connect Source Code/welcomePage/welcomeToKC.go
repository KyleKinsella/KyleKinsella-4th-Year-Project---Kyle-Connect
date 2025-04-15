package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var welcomePage = `
		<style>
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
			margin: 10px;
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

        .input-container {
            display: flex;
            flex-direction: column;
            gap: 10px;
            margin-top: 20px;
        }

        .btn {
            background-color: #007bff;
            color: white;
            padding: 12px;
            border: none;
            border-radius: 8px;
            cursor: pointer;
            font-size: 18px;
            font-weight: bold;
            transition: 0.3s;
            width: 100%;
        }

        .btn:hover {
            background-color: #0056b3;
        }

        a {
            text-decoration: none;
        }
		</style>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kyle Connect - Welcome Page</title>
</head>
<body>

	<div class="container">
		<h1>Welcome to Kyle Connect</h1>
		<p>Kyle Connect is the place where you can talk to anyone anywhere in the world! 
		Want to chat with friends, start conversations and do so much more? Join Kyle Connect Today!</p>
		
		<div class="input-container">
			<a href="http://localhost:8080/"><button class="btn">Join Now!</button></a>
			<a href="http://localhost:8081/login/login.go"><button class="btn">Already Have An Account?</button></a>
		</div>
	</div>
</body>
</html>
`

// Handler function to serve the form and process submissions
func formHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the HTML template
    tmpl, err := template.New("").Parse(welcomePage) // might need to do it here!?
    if err != nil {
        log.Fatal(err)
    }
	tmpl.Execute(w, nil)
}

func main() {
    // Set up the route and handler for the form
    http.HandleFunc("/", formHandler)

    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8089")
    log.Fatal(http.ListenAndServe(":8089", nil))
}
