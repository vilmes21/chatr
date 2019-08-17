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
	// r.ParseForm()
	// firstData := r.FormValue("userNowId")
	// fmt.Println(`firstData23: `,firstData)
	
	ws, err := upgrader.Upgrade(w, r, nil)

	// err7 := ws.WriteJSON(`{"test":true,"fun":true,"adsf":3}`)
	// if err7 !=nil {
	// 	fmt.Println("OH err7 !")
	// 	log.Fatal(err7)
	// }

	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	

	// if firstData > 0 {
	// 	//todo: auth, check is really logged in 1st
	// 	if _, exist := clients[msg.UserNowId]; !exist {
	// 		clients[msg.UserNowId] = ws
	// 		fmt.Println(`Just mapped CONN OF msg.UserNowId:`, msg.UserNowId)
	// 	}
	// }

	for {
		var msg common.MsgReceived
		err := ws.ReadJSON(&msg)
		
		if msg.ChatId == 0 || msg.UserNowId ==0 {
			fmt.Println("GONNA break. 1st check error, msg.ChatId: ", msg.ChatId, " msg.UserNowId: ", msg.UserNowId)
			break
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
