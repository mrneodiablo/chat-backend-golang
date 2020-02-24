package validator

import "hope-pet-chat-backend/dto"

//{
//	Action         int    `json:"action" bson:"action"` // CONFIRM, FETCH, SEND, HEARTBEAT
//	UserId         int64  `json:"userId" bson:"userId"`
//	SenderUserId   int64  `json:"senderuserId" bson:"senderuserId"`
//	ReceiverUserId int64  `json:"receiveruserId" bson:"receiveruserId"`
//	Data           string `json:"data" bson:"data"`
//}

func SendValidator(message dto.MessageDto, userId int64) bool {

	flag := true
	if message.Action == 0 {
		flag = false
	}
	if message.UserId == 0 {
		flag = false
	}
	if message.SenderUserId == 0 {
		flag = false
	}
	if message.ReceiverUserId == 0 {
		flag = false
	}
	if message.Data == "" {
		flag = false
	}
	if message.UserId != userId || message.UserId != message.SenderUserId || message.SenderUserId != userId {
		flag = false
	}
	return flag
}
