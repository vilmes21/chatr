package common

type MsgReceived struct {
	Msg string `json:"msg"`
	UserNowId int `json:"userNowId"`
	ChatId int `json:"chatId"`
};

type ChatPair struct {
	ChatId int `json:"chatId"`
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
	SpeakerUserId int `json:"speakerUserId"`
	Content string `json:"content"`
	ChatId int `json:"chatId"`
}