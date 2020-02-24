package models

import (
	"hope-pet-chat-backend/connector"
	"time"
)

type ReadMessage struct {
	MessageId      int64     `json:"messageId" bson:"messageId"`
	UserId         int64     `json:"userId" bson:"userId"`
	SenderUserId   int64     `json:"senderuserId" bson:"senderuserId"`
	ReceiverUserId int64     `json:"receiveruserId" bson:"receiveruserId"`
	Data           string    `json:"data" bson:"data"`
	CreatedDate           time.Time `json:"createdDate" bson:"createdDate"`
}

func AddReadMessage(this ReadMessage) (ReadMessage, error) {
	var (
		err error
	)
	c := connector.SessionConnectCollection(connector.CHAT_READMESSAGE_COLLECTION)
	defer c.Close()

	err = c.Session.Insert(this)
	if err != nil {
		panic(err)
	}
	return this, err

}
