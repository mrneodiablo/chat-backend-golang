package validator

import "hope-pet-chat-backend/dto"

//{
//	Action      int    `json:"action" bson:"action"` // CONFIRM, FETCH, SEND, HEARTBEAT
//	UserId      int64  `json:"userId" bson:"userId"`
//}

func BaseValidator(message dto.MessageDto) (bool) {

	flag := true
	if message.Action == 0 {
		flag = false
	}
	if message.UserId == 0 {
		flag = false
	}
	return flag

}
