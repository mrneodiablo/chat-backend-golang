package validator

import "hope-pet-chat-backend/dto"

//{
//	Action         int    `json:"action" bson:"action"` // CONFIRM, FETCH, SEND, HEARTBEAT
//	UserId         int64  `json:"userId" bson:"userId"`
//}

func FetchValidator(message dto.MessageDto, userId int64)  bool {

	flag := true
	if message.Action == 0 {
		flag = false
	}
	if message.UserId == 0 {
		flag = false
	}
	if message.UserId != userId {
		flag = false
	}
	return flag
}
