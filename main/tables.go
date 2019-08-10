package main

import (
	"database/sql"
	"time"
)

type Store interface {
	CreateSentence(s *Sentence) error
	GetSentences(chatId int) ([]*Sentence, error)
}

var store Store

func InitStore(s Store) {
	store = s
}

type dbStore struct {
	db *sql.DB
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
