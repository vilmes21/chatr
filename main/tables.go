package main

import (
	"database/sql"
	"time"
	"fmt"
	"../common"
	"../keys"
	_"github.com/lib/pq"
)

type Store interface {
	CreateSentence(s *Sentence) error
	GetSentences(chatId int) ([]*Sentence, error)
	// FindChatIdByUserIds(userId int, user2Id int) common.JsonResp
	CreateChat() int
	CreateChatSpeakers(chaterIds *common.ChatPair) bool
}


type dbStore struct {
	db *sql.DB
}

var store dbStore

func InitStore(s *dbStore) {
	store = *s
}

type Chat struct {
	Id int 
	Title string 
}

type User struct {
	Id int 
	Name string 
}

type Friend struct {
	UserId int 
	User2Id int 
}

type ChatSpeaker struct {
	Id int 
	ChatId int 
	UserId int 
}

type Sentence struct {
	ChatSpeakerId int 
	Content string 
	Time time.Time 
}

func initDB(){
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
}