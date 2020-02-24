package models

import (
	"gopkg.in/mgo.v2/bson"
	"hope-pet-chat-backend/connector"
	"time"
)

type UnReadMessage struct {
	MessageId      int64     `json:"messageId" bson:"messageId"`
	UserId         int64     `json:"userId" bson:"userId"`
	SenderUserId   int64     `json:"senderuserId" bson:"senderuserId"`
	ReceiverUserId int64     `json:"receiveruserId" bson:"receiveruserId"`
	Data           string    `json:"data" bson:"data"`
	CreatedDate           time.Time `json:"createdDate" bson:"createdDate"`
}

func AddUnReadMessage(this UnReadMessage) (UnReadMessage, error) {
	var (
		err error
	)
	c := connector.SessionConnectCollection(connector.CHAT_UNREADMESSAGE_COLLECTION)
	defer c.Close()

	err = c.Session.Insert(this)
	if err != nil {
		panic(err)
	}
	return this, err

}

func RemoveUnReadMessage(messageId int64, userId int64) (error) {
	var (
		err error
	)
	c := connector.SessionConnectCollection(connector.CHAT_UNREADMESSAGE_COLLECTION)
	defer c.Close()

	err = c.Session.Remove(bson.M{"messageId": messageId, "receiveruserId": userId })
	return err

}

func GetUnReadMessage(userId int64) ([]UnReadMessage, error) {
	var (
		err error
	)
	c := connector.SessionConnectCollection(connector.CHAT_UNREADMESSAGE_COLLECTION)
	defer c.Close()

	var unReadMessage []UnReadMessage
	err = c.Session.Find(bson.M{"receiveruserId": userId}).Sort("-createdDate").All(&unReadMessage)
	if unReadMessage == nil {
		unReadMessage = []UnReadMessage{}
	}
	return unReadMessage, err

}
