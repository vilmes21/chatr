package main

import (
  "database/sql"
  "fmt"
  _"github.com/lib/pq"
  "./keys"
)

func main(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
	keys.Host, keys.Port, keys.User, keys.Dbname)
	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
	  panic(err)
	}
  
	fmt.Println("Successfully connected!")

	sqlStatement := `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	// sqlStatement := `
	// INSERT INTO chat (title)
	// VALUES ($1)
	// RETURNING id`
	  id := 0
	  err = db.QueryRow(sqlStatement, 30, "a2@b.com", "J4ack", "Smith").Scan(&id)

	//err = db.QueryRow(sqlStatement, "Funny Things deux").Scan(&id)

	  if err != nil {
		panic(err)
	  }
	  fmt.Println("New record ID is:", id)
}