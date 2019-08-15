package main

import (
	"log"
	"fmt"
	"net/http"
	// "time"
	"../common"
	"encoding/json"

	// "github.com/gorilla/websocket"
)

func (store *dbStore) FindChatIdByUserIds(userId int, user2Id int) error {
	rows, err := store.db.Query(
		`select chat_id, COUNT(*) from chat_speaker 
		where chat_id = 
		(select a.chat_id from chat_speaker as a JOIN chat_speaker as b ON a.chat_id=b.chat_id WHERE a.user_id=$1 AND b.user_id=$2)
		GROUP BY chat_id
		HAVING Count(*) =2;`, 
		userId, user2Id)
		if err != nil {
			log.Println("err while querying: ", err)
		}
		defer rows.Close()
		
		log.Println("rows: ", rows)
		fmt.Println("rows: v%", rows)

		chat_id, count, rowNum := 0, 0, 0
		for rows.Next() {
			rowNum++
			err := rows.Scan(&chat_id, &count)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("-----------LOOK")
			log.Println(chat_id, count)
		}

		newChatId := 0
		if rowNum == 0 {
			newChatIdRow, err2 := store.db.Query(`insert into chat (title) values ('') returning id`)

			fmt.Println("newChatIdRow:", newChatIdRow)
			
			if err2 != nil {
				log.Println("err2 while querying: ", err)
			}

			newChatIdRow.Scan(&newChatId)

			if newChatId > 0 {
				chaterIds := [2]int {userId, user2Id}
				for _, chaterId := range chaterIds {
					_, err3 := store.db.Query(
						`INSERT INTO chat_speaker (chat_id, user_id) VALUES ($1, $2)`, 
						newChatId, chaterId)
						if err3 != nil {
							log.Println("err3 while querying: ", err)
						}
				}
			} else {
				fmt.Println("Inserting into Chat Failed!!!!!")
			}
	
		}

		return err
}


func newChatHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
    var p common.ChatPair
    err := decoder.Decode(&p)
    if err != nil {
        panic(err)
    }
    log.Println("p is: ", p)
	
	store.FindChatIdByUserIds(p.UserId, p.User2Id)
	
	
// 	res := common.JsonResp {Success: true, Msg: "You are cool"}
// 	resString, _ := json.Marshal(res)
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
//    w.WriteHeader(200)
//    w.Header().Set("Content-Type", "application/json")
//    w.Write(resString)
}
