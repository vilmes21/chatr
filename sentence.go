package main

import (
  "database/sql"
  "fmt"
  _"github.com/lib/pq"
  "./getDB"
  "time"
)

func createSentence(content, chatSpeakerId) {
	db := getDB.Db
	sqlStatement := `
	INSERT INTO sentence (chat_speaker_id, content, time)
	VALUES ($1, $2, $3)`

	  err = db.QueryRow(sqlStatement, chatSpeakerId, content, time.now()).Scan(&id)

	  if err != nil {
		panic(err)
	  }

	  fmt.Println("New sentence created:", content)
}