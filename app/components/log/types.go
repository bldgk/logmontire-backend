package log

import (
	"logmontire/crypto"
	"strconv"
)

type Log struct {
	Id       string `bson:"_id"json:"_id"`
	Title    string `bson:"title"json:"title"`
	Username string `bson:"username"json:"username"`
	Body     string `bson:"body"json:"body"`
	Channel  string `bson:"channel"json:"channel"`
	//ChannelId string `bson:"channel_id"json:"channel_id"`
}

func (l *Log) New(title string, username string, body string, channel string) {
	l.Id = strconv.Itoa(crypto.GenerateId())
	l.Title = title
	l.Username = username
	l.Body = body
	l.Channel = channel
	//l.ChannelId = channelId
}

func (l *Log) GetId() string {
	return l.Id
}

func (l *Log) GetInstance() interface{} {
	return l
}
