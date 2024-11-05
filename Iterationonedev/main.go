package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func catchError(err error) error {
	if err != nil {
		panic(err.Error())
	}
	return err
}

func retrieveInfo() {
	//......
}


type Data struct {
	id int
	username string
}

func main() {

	fmt.Println("lets connect to a db using go!!")

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/users")
	catchError(err)

	defer db.Close()

	query := "SELECT userId, username FROM communicators"
	info, err := db.Query(query)
	catchError(err)

	// fmt.Println("value of info variable is: ", info)

	var data []Data

	for info.Next() {
		var i Data
		err = info.Scan(&i.id, &i.username)
		catchError(err)

		data = append(data, i)
	}
	err = info.Err()
	catchError(err)


	for _, ident := range data {
		fmt.Println("Id: ", ident.id, "\nusername: ", ident.username)
	}	
}