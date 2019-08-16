package main

import (
	"log"
	"fmt"
	"net/http"
	// "time"
	"../common"
	"encoding/json"

	// "github.com/davecgh/go-spew/spew"
	// "github.com/gorilla/websocket"
)

func (store *dbStore) CreateChat() int {
	newChatId := 0
	newChatIdRow, err2 := store.db.Query(`insert into chat (title) values ('') returning id`)

	// fmt.Println("newChatIdRow:", newChatIdRow)
	fmt.Printf("----------------")

	
	fmt.Printf("----------------")

	if err2 != nil {
		log.Println("err2 while querying: ", err2)
	}

			for newChatIdRow.Next() {
				err := newChatIdRow.Scan(&newChatId)
				if err != nil {
					log.Fatal(err)
				}
			}
		 
				return newChatId
}

func (store *dbStore) CreateChatSpeakers(p *common.ChatPair) bool {
	done := true
	userIds := [2]int {p.UserId, p.User2Id}
	
	loopSuccess := false
	for _, chaterUserId := range userIds {
		_, err3 := store.db.Query(
			`INSERT INTO chat_speaker (chat_id, user_id) VALUES ($1, $2)`, 
			p.ChatId, chaterUserId)
			if err3 == nil {
				loopSuccess = true
			} else {
				log.Println("err3 while querying: ", err3)
			}

			done = done && loopSuccess
	}

	return done
}

func findChatIdByUserIds(userId int, user2Id int) int {
	existingChatId := 0
	rows, err := store.db.Query(
		`select chat_id, COUNT(*) from chat_speaker 
		where chat_id = 
		(select a.chat_id from chat_speaker as a JOIN chat_speaker as b ON a.chat_id=b.chat_id WHERE a.user_id=$1 AND b.user_id=$2)
		GROUP BY chat_id
		HAVING Count(*) < 3;`, 
		userId, user2Id)
		if err != nil {
			log.Println("err while querying: ", err)
		}
		defer rows.Close()
		
		count := 0
		for rows.Next() {
			err := rows.Scan(&existingChatId, &count)
			if err != nil {
				log.Fatal(err)
			}
		}

		return existingChatId
}


func newChatHandler(w http.ResponseWriter, r *http.Request) {
	res := common.JsonResp {}

	decoder := json.NewDecoder(r.Body)
    var p common.ChatPair
    err := decoder.Decode(&p)
    if err != nil {
        panic(err)
    }
	
	if p.User2Id == p.UserId {
		res.Msg = "Self-talk is bad for the soul"
		resString, _ := json.Marshal(res)
		w.Write(resString)
	}
	
	res.Id = findChatIdByUserIds(p.UserId, p.User2Id)

	if res.Id == 0 {
		res.Id = store.CreateChat()
		if res.Id > 0 {
			p.ChatId = res.Id
			res.Success = store.CreateChatSpeakers(&p)
		}
	}
	
	resString, _ := json.Marshal(res)
	   w.Header().Set("Content-Type", "application/json")
   w.Write(resString)
}
