package controllers

import (
	"github.com/astaxie/beego"
	"hope-pet-chat-backend/models"
	"strconv"
)

type ChatStatusController struct {
	beego.Controller
}

func (this *ChatStatusController) GetLastLogin() {
	userId,_ := strconv.ParseInt(this.GetString("userId"), 10, 64)
	heartbeat, _ := models.GetLastTimeHeartBeat(userId)
	this.Data["json"] = heartbeat
	this.ServeJSON()
}