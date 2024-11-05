package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func CatchError(err error) error {
	if err != nil {
		panic(err.Error())
	}
	return err
}

func retrieveInfo() {
	//......
}



func PutDataToDb(db *sql.DB, username, email, password string) error {

	sql := "INSERT INTO communicators (username, email, password) VALUES (?, ?, ?)"

	_, err := db.Query(sql, username, email, password)
	CatchError(err)

	fmt.Println("user inserted into database")

	return err
}


type Data struct {
	id int
	username string
}

func main() {

	fmt.Println("lets connect to a db using go!!")

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/users")
	CatchError(err)

	defer db.Close()

	query := "SELECT userId, username FROM communicators"
	info, err := db.Query(query)
	CatchError(err)

	// fmt.Println("value of info variable is: ", info)

	var data []Data

	for info.Next() {
		var i Data
		err = info.Scan(&i.id, &i.username)
		CatchError(err)

		data = append(data, i)
	}
	err = info.Err()
	CatchError(err)


	for _, ident := range data {
		fmt.Println("Id: ", ident.id, "\nusername: ", ident.username)
	}	
}