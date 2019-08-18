package main

import (
	"log"
	"fmt"
	"net/http"
	"time"
	"../common"

	"github.com/gorilla/websocket"
)

// var clients = make(map[*websocket.Conn]int)
var clients = make(map[int]*websocket.Conn)
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
	
	var msg common.MsgReceived

	//this line is blocking
	err1 := ws.ReadJSON(&msg)
	if err1 != nil {
		log.Fatal(`err1: `)
	}

	if msg.UserNowId == 0 {
		ws.WriteJSON(`{"Provided user id":false}`)
		return
	}

	//todo: auth, check is really logged in 1st
	//note: this way, if same user opens multiple windows, it will only feed the 1st window.
	if _, exist := clients[msg.UserNowId]; !exist {
		clients[msg.UserNowId] = ws
		fmt.Println(`before for loop Just mapped CONN OF msg.UserNowId:`, msg.UserNowId)
	}

	for {
		err := ws.ReadJSON(&msg)
		
		if msg.UserNowId ==0 {
			fmt.Println("GONNA skip. msg.ChatId: ", msg.ChatId, " msg.UserNowId: ", msg.UserNowId)
			continue
		}
		
		memberIds := getChatMembersIds(msg.UserNowId, msg.ChatId)

		if err != nil {
			log.Printf("error: %v", err)
			// delete(clients, ws)
			break
		}

		//save to db
		newS := Sentence {
			Time: time.Now(), 
			Content: msg.Msg, 
			ChatSpeakerId: memberIds.SenderSpeakerId,
		}

		err = store.CreateSentence(&newS) 
		if err != nil {
			log.Printf("error: %v", err)
			break
		}

		//In 2-person chat, only 1 listener. In group chats, more listeners
		//todo: or consider making ReceiverUserId an array?
		for _, listenerId := range memberIds.ListenersUserIds {
			broadcast <- common.MessageObj {
				ReceiverUserId: listenerId, 
				SpeakerUserId: msg.UserNowId,
				Content: msg.Msg,
				ChatId: msg.ChatId,
			}
		}


	}
}

func pushMsgToClient() {
	for {
		msg := <-broadcast
	
		//need to push to all logged in users
		for userId, client := range clients {
			if userId == msg.SpeakerUserId || userId == msg.ReceiverUserId {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					client.Close()
					// delete(clients, client)
				}
			}
		}
	}
}
