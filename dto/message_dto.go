package dto

type ActionType int

type MessageDto struct {
	MessageId      int64  `json:"messageId" bson:"messageId"`
	Action         int    `json:"action" bson:"action"` // CONFIRM, FETCH, SEND, HEARTBEAT
	UserId         int64  `json:"userId" bson:"userId"`
	SenderUserId   int64  `json:"senderuserId" bson:"senderuserId"`
	ReceiverUserId int64  `json:"receiveruserId" bson:"receiveruserId"`
	Data           string `json:"data" bson:"data"`
}
