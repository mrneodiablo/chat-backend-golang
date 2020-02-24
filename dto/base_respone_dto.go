package dto

type BaseResponeDto struct {
	Action       int         `json:"action" bson:"action"`             // CONFIRM, FETCH, SEND, HEARTBEAT
	Status       int         `json:"status" bson:"status"`             // 1 sucess, 0 error
	ErrorCode    int         `json:"errorCode" bson:"errorCode"`       // 1 khong loi, 0 error
	ErrorMessage string      `json:"errorMessage" bson:"errorMessage"` // 1 khong loi, 0 error
	Data         interface{} `json:"data" bson:"data"`
}
