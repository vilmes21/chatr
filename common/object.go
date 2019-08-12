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