package main

import (
	"log"
	// "fmt"
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

func CreateSentenceHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	for {
		var msg common.MessageObj
		err := ws.ReadJSON(&msg)

		//todo: verify in db before setting this
		clients[ws] = msg.ChatId

		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		//save to db
		newSentence := Sentence {Time: time.Now(), Content: msg.Content, ChatSpeakerId: msg.ChatSpeakerId}
		err = store.CreateSentence(&newSentence) 
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		broadcast <- msg
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
