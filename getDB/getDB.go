package getDB

import (
  "database/sql"
  "fmt"
  _"github.com/lib/pq"
  "../keys"
)

func getDB(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
	keys.Host, keys.Port, keys.User, keys.Dbname)
	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
	  panic(err)
	}

	fmt.Println("Exporting db...")
	return db
}

var Db = getDB()
