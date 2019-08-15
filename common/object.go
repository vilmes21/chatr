package common

type ChatPair struct {
	UserId int `json:"userId"`
	User2Id int `json:"user2Id"`
}

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