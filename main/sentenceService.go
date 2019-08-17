package main

import (
	"fmt"
	"log"
	"../common"
)

func getChatMembersIds(senderUserId int, chatId int) common.ChatMembersIds {
	ids := common.ChatMembersIds {
		SenderUserId: senderUserId,
		SenderSpeakerId: 000,
	}

	rows, err := store.db.Query(
		`SELECT id as speakerId, user_id FROM chat_speaker WHERE chat_id=$1`,
		chatId,
	)

	defer rows.Close()

	if err !=nil {
		fmt.Println(`SQL err in func getChatMembersIds`)
		log.Fatal(err)
	}

	members := [](*[2]int){}
	for rows.Next(){
		speakerAndUserIds := [2]int{}
		if err2 := rows.Scan(&speakerAndUserIds[0], &speakerAndUserIds[1]); err != nil {
			fmt.Println(`err2 SQL err in func getChatMembersIds`)
			log.Fatal(err2)
		}
		members=append(members, &speakerAndUserIds)
	}

	for _, arr := range members {
		if arr[1] == senderUserId{
			ids.SenderSpeakerId = arr[0]
			continue
		}
		
		ids.ListenersUserIds = append(ids.ListenersUserIds, arr[1])
	}

	return ids
}

func (store *dbStore) CreateSentence(s *Sentence) error {
	_, err := store.db.Query("INSERT INTO sentence(chat_speaker_id, content, time) VALUES ($1,$2, $3)", s.ChatSpeakerId, s.Content, s.Time)
	return err
}

func (store *dbStore) GetSentences(chatId int) ([]*Sentence, error) {
	//todo: use chatId to fetch chatSpeakerIds
	chatSpeakerIds := []int{1,2,3}
	rows, err := store.db.Query("SELECT * from sentence where chat_speaker_id in $1", chatSpeakerIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
 
	sentences := []*Sentence{}
	for rows.Next() {
		s := &Sentence{}
		if err := rows.Scan(&s.ChatSpeakerId, &s.Content, &s.Time); err != nil {
			return nil, err
		}
		
		sentences = append(sentences, s)
	}
	return sentences, nil
}