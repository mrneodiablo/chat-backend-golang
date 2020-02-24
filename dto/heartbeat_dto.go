package dto

import (
	"hope-pet-chat-backend/models"
	"time"
)

type HeartBeatDto struct {
	UserId      int64     `json:"userId" bson:"userId"`
	CreatedDate time.Time `json:"createdDate" bson:"createdDate"`
}

func ConverHeartBeatEntityToDto(chatHeartBeatModelsEntity models.ChatHeartBeatModels) (HeartBeatDto) {
	heartbeatdto := HeartBeatDto{
		UserId:      chatHeartBeatModelsEntity.UserId,
		CreatedDate: chatHeartBeatModelsEntity.CreatedDate}

	return heartbeatdto
}
