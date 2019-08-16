package main

import (
	"log"
	"fmt"
	"net/http"
	"time"
	"../common"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]int)
var broadcast = make(chan common.MessageObj)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func getChatSpeakerId(chatId int, userNowId int) int {
	fmt.Println("chatId: ", chatId, " userNowId: ", userNowId)
	
	row, err := store.db.Query(
		`SELECT id FROM chat_speaker WHERE chat_id= $1 AND user_id= $2`, 
		chatId, 
		userNowId,
	)

	if err != nil {
		log.Fatal(err)
	}

	id := 0
	for row.Next() {
		err := row.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	} 

	return id
}

func CreateSentenceHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	for {
		var msg common.MsgReceived
		err := ws.ReadJSON(&msg)

		if msg.ChatId == 0 || msg.UserNowId ==0 {
			log.Printf("1st check error: userNowSpeakerId == 0")
			return
		}
		//todo: verify in db before setting this
		userNowSpeakerId := getChatSpeakerId(msg.ChatId, msg.UserNowId)
		if userNowSpeakerId == 0 {
			log.Printf("error: userNowSpeakerId == 0")
			return
		}
		
		clients[ws] = msg.ChatId

		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		//save to db
		newS := Sentence {Time: time.Now(), Content: msg.Msg, ChatSpeakerId: userNowSpeakerId}
		err = store.CreateSentence(&newS) 
		if err != nil {
			log.Printf("error: %v", err)
			break
		}

		broadcast <- common.MessageObj {
			SpeakerUserId: msg.UserNowId,
			Content: msg.Msg,
			ChatId: msg.ChatId,
		}
	}
}

func pushMsgToClient() {
	for {
		msg := <-broadcast
	
		//todo: for loop can be optimized to using arrary or map of clients of one chatID 
		for client, chatId := range clients {
			if chatId == msg.ChatId {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}
