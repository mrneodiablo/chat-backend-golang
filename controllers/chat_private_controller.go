package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"hope-pet-chat-backend/constants"
	"hope-pet-chat-backend/dto"
	"hope-pet-chat-backend/models"
	"hope-pet-chat-backend/utilities"
	"hope-pet-chat-backend/validator"
	"net/http"
	"strconv"
	"time"
)

type ChatPrivateController struct {
	beego.Controller
}

func (this *ChatPrivateController) PrivateChat() {

	userId, errUserId := strconv.ParseInt(this.GetString("userId"), 10, 64)
	session := this.GetString("session")
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)

	if errUserId != nil || session == "" {
		respone := dto.BaseResponeDto{
			Action:       constants.CONNECT,
			Status:       constants.ERROR,
			ErrorCode:    constants.URL_IN_VALID,
			ErrorMessage: "URL_IN_VALID",
			Data:         nil}
		out, _ := json.Marshal(respone)
		ResponeJsonToClient(ws, string(out))
		ws.Close()
	}

	if session != utilities.GenerateSession(userId) {
		respone := dto.BaseResponeDto{
			Action:       constants.CONNECT,
			Status:       constants.ERROR,
			ErrorCode:    constants.CONNECT_NOT_ALLOW,
			ErrorMessage: "CONNECT_NOT_ALLOW",
			Data:         nil}
		out, _ := json.Marshal(respone)
		ResponeJsonToClient(ws, string(out))
		ws.Close()
	}

	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	msgCh := coordinator.Subscribe()
	defer coordinator.Unsubscribe(msgCh)

	chan_receive := make(chan dto.MessageDto)
	go ReceiveMsgFromClient(ws, chan_receive, userId)

	unReadEntity := models.UnReadMessage{}

	for {
		select {

		// nhan message
		case msgDto := <-chan_receive:
			switch msgDto.Action {
			case constants.HEARTBEAT: //HEARTBEAT
				heartbeatEntity := models.ChatHeartBeatModels{}
				heartbeatEntity.UserId = msgDto.UserId
				heartbeatEntity.CreatedDate = time.Now()
				hb, error := models.AddHeartBeat(heartbeatEntity)
				if error != nil {
					beego.Error(error.Error())
					respone := dto.BaseResponeDto{
						Action:       constants.HEARTBEAT,
						Status:       constants.ERROR,
						ErrorCode:    constants.SYSTEM_GENERAL_EXCEPTION,
						ErrorMessage: error.Error(),
						Data:         nil}
					out, _ := json.Marshal(respone)
					ResponeJsonToClient(ws, string(out))
				} else {
					respone := dto.BaseResponeDto{
						Action:       constants.HEARTBEAT,
						Status:       constants.SUCCESS,
						ErrorCode:    constants.SUCCESS,
						ErrorMessage: "",
						Data:         dto.ConverHeartBeatEntityToDto(hb)}
					out, _ := json.Marshal(respone)
					ResponeJsonToClient(ws, string(out))
					beego.Info(string(out))
				}
			case constants.SEND: // SEND
				unReadEntity.MessageId = msgDto.MessageId
				unReadEntity.UserId = msgDto.UserId
				unReadEntity.SenderUserId = msgDto.SenderUserId
				unReadEntity.ReceiverUserId = msgDto.ReceiverUserId
				unReadEntity.Data = msgDto.Data
				unReadEntity.CreatedDate = time.Now()
				output, error := models.AddUnReadMessage(unReadEntity)

				if error != nil {
					respone := dto.BaseResponeDto{
						Action:       constants.SEND,
						Status:       constants.ERROR,
						ErrorCode:    constants.SYSTEM_GENERAL_EXCEPTION,
						ErrorMessage: error.Error(),
						Data:         nil}
					out, _ := json.Marshal(respone)
					ResponeJsonToClient(ws, string(out))
					beego.Error(error.Error())

				} else {
					respone := dto.BaseResponeDto{
						Action:       constants.SEND,
						Status:       constants.SUCCESS,
						ErrorCode:    constants.SUCCESS,
						ErrorMessage: "",
						Data:         output}
					out, _ := json.Marshal(respone)
					ResponeJsonToClient(ws, string(out))
					beego.Info(output)
				}
			case constants.FETCH: //FETCH
				unReadEntitys, error := models.GetUnReadMessage(msgDto.UserId)
				if error != nil {
					respone := dto.BaseResponeDto{
						Action:       constants.FETCH,
						Status:       constants.ERROR,
						ErrorCode:    constants.SYSTEM_GENERAL_EXCEPTION,
						ErrorMessage: error.Error(),
						Data:         nil}
					out, _ := json.Marshal(respone)
					ResponeJsonToClient(ws, string(out))
					beego.Error(error.Error())
				} else {
					respone := dto.BaseResponeDto{
						Action:       constants.FETCH,
						Status:       constants.SUCCESS,
						ErrorCode:    constants.SUCCESS,
						ErrorMessage: "",
						Data:         unReadEntitys}
					out, _ := json.Marshal(respone)
					ResponeJsonToClient(ws, string(out))
				}
				beego.Info(msgDto)
			case constants.CONFIRM: //CONFIRM
				error := models.RemoveUnReadMessage(msgDto.MessageId, msgDto.UserId)
				if error != nil {
					respone := dto.BaseResponeDto{
						Action:       constants.CONFIRM,
						Status:       constants.ERROR,
						ErrorCode:    constants.CONFIRM_INVALID,
						ErrorMessage: error.Error(),
						Data:         nil}
					out, _ := json.Marshal(respone)
					ResponeJsonToClient(ws, string(out))
					beego.Error(error.Error())

				} else {
					respone := dto.BaseResponeDto{
						Action:       constants.CONFIRM,
						Status:       constants.SUCCESS,
						ErrorCode:    constants.SUCCESS,
						ErrorMessage: "",
						Data:         nil}
					out, _ := json.Marshal(respone)
					ResponeJsonToClient(ws, string(out))
					beego.Info(msgDto)
				}

			}
			//respone message
		case data := <-msgCh:
			br := data.(dto.MessageDto)
			if br.ReceiverUserId == userId {
				out, _ := json.Marshal(br)
				ResponeJsonToClient(ws, string(out))
			}
		}
	}

}

func ReceiveMsgFromClient(ws *websocket.Conn, chan_receive chan dto.MessageDto, userId int64) {
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}
		data := dto.MessageDto{}

		dataDecrypt, err := utilities.Decrypt(string(msg))

		json.Unmarshal(dataDecrypt, &data)

		// basic validator
		if !validator.BaseValidator(data) {
			respone := dto.BaseResponeDto{
				Action:       data.Action,
				Status:       constants.ERROR,
				ErrorCode:    constants.DATA_IN_VALID,
				ErrorMessage: "DATA_IN_VALID",
				Data:         nil}
			out, _ := json.Marshal(respone)
			ResponeJsonToClient(ws, string(out))
			beego.Error(string(out))
		}

		// validate data
		switch data.Action {
		case constants.HEARTBEAT:
			if !validator.HeartBeatValidator(data, userId) {
				respone := dto.BaseResponeDto{
					Action:       data.Action,
					Status:       constants.ERROR,
					ErrorCode:    constants.DATA_IN_VALID,
					ErrorMessage: "DATA_IN_VALID",
					Data:         nil}
				out, _ := json.Marshal(respone)
				ResponeJsonToClient(ws, string(out))
				beego.Error(string(out))
			} else {
				chan_receive <- data
			}
		case constants.SEND:
			if !validator.SendValidator(data, userId) {
				respone := dto.BaseResponeDto{
					Action:       data.Action,
					Status:       constants.ERROR,
					ErrorCode:    constants.DATA_IN_VALID,
					ErrorMessage: "DATA_IN_VALID",
					Data:         nil}
				out, _ := json.Marshal(respone)
				ResponeJsonToClient(ws, string(out))
				beego.Error(string(out))
			} else {

				if data.MessageId == 0 {
					data.MessageId = utilities.GenerateMessageId(data.UserId)
				}

				chan_receive <- data
				// publis all channel with action send
				coordinator.Publish(data)

			}
		case constants.FETCH:
			if !validator.FetchValidator(data, userId) {
				respone := dto.BaseResponeDto{
					Action:       data.Action,
					Status:       constants.ERROR,
					ErrorCode:    constants.DATA_IN_VALID,
					ErrorMessage: "DATA_IN_VALID",
					Data:         nil}
				out, _ := json.Marshal(respone)
				ResponeJsonToClient(ws, string(out))
				beego.Error(string(out))
			}else {
				chan_receive <- data
			}
		case constants.CONFIRM:
			if !validator.ConfirmValidator(data, userId) {
				respone := dto.BaseResponeDto{
					Action:       data.Action,
					Status:       constants.ERROR,
					ErrorCode:    constants.DATA_IN_VALID,
					ErrorMessage: "DATA_IN_VALID",
					Data:         nil}
				out, _ := json.Marshal(respone)
				ResponeJsonToClient(ws, string(out))
				beego.Error(string(out))
			}else {
				chan_receive <- data
			}
		}

	}
}

func ResponeJsonToClient(ws *websocket.Conn, data string) {
	rp, _ := utilities.Encrypt([]byte(data))
	ws.WriteMessage(websocket.TextMessage, rp)
}
