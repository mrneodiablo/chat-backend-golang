package validator

import "hope-pet-chat-backend/dto"

//{
//	MessageId      int64  `json:"messageId" bson:"messageId"`
//	Action         int    `json:"action" bson:"action"`
//	UserId         int64  `json:"userId" bson:"userId"`
//	SenderUserId   int64  `json:"senderuserId" bson:"senderuserId"`
//	ReceiverUserId int64  `json:"receiveruserId" bson:"receiveruserId"`
//}


func ConfirmValidator(message dto.MessageDto, userId int64)  bool {

	flag := true
	if message.MessageId == 0 {
		flag = false
	}
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
	if message.UserId != userId || message.UserId != message.ReceiverUserId || message.ReceiverUserId != userId{
		flag = false
	}
	return flag
}