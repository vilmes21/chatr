package main

import (
    "fmt"
    // "strconv"
    "log"
    "net/http"
	"database/sql"
    "sync"
    // "time"
    _"github.com/lib/pq"
    "../keys"
    // "../controller"
  )

var counter int
var mutex = &sync.Mutex{}

func signup(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "name to sign up: " + r.FormValue("name"))
}



func main() {
    connectInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
    keys.Host, keys.Port, keys.User, keys.Dbname)
    
    db, err := sql.Open("postgres", connectInfo)
    if err != nil {
        panic(err)
    }
    err = db.Ping()
    if err != nil {
        panic(err)
    }

	InitStore(&dbStore{db: db})

    http.HandleFunc("/signup", signup)
    http.HandleFunc("/sentence/create", CreateSentenceHandler)

    log.Fatal(http.ListenAndServe(":8081", nil))
}
