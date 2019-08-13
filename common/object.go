package common

type JsonResp struct {
	Success bool
	Id int
	Msg string
}

type Message struct {
	Username string
	Message  string
}

type MessageObj struct {
	ChatSpeakerId int `json:"chatSpeakerId"`
	Content string `json:"content"`
	ChatId int `json:"chatId"`
}