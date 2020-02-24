package models

import (
	"time"
	"hope-pet-chat-backend/connector"
	"gopkg.in/mgo.v2/bson"
)

type ChatHeartBeatModels struct {
	UserId      int64     `json:"userId" bson:"userId"`
	CreatedDate time.Time `json:"createdDate" bson:"createdDate"`
}

func GetLastTimeHeartBeat(userId int64) (ChatHeartBeatModels, error) {
	var (
		err       error
		heartbeat ChatHeartBeatModels
	)
	c := connector.SessionConnectCollection(connector.CHAT_HEARTBEAT_COLLECTION)
	defer c.Close()

	query := bson.M{"userId": userId}

	err = c.Session.Find(query).Sort("-createdDate").One(&heartbeat)
	if err != nil {
		return heartbeat, err
	}
	return heartbeat, err
}

func AddHeartBeat(u ChatHeartBeatModels) (ChatHeartBeatModels, error) {
	var (
		err error
	)
	c := connector.SessionConnectCollection(connector.CHAT_HEARTBEAT_COLLECTION)
	defer c.Close()

	err = c.Session.Insert(u)
	if err != nil {
		panic(err)
	}
	return u, err

}
