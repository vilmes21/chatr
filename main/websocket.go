package main

import (
	"log"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Sentence)

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

	clients[ws] = true

	for {
		var msg Sentence
		err := ws.ReadJSON(&msg)

		fmt.Printf("ws.ReadJSON(&msg), so msg is: %v", msg)
		
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		msg.Time= time.Now()
		
		err = store.CreateSentence(&msg) //save into db
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
	
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
