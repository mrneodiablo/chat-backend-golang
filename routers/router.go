package routers

import (
	"hope-pet-chat-backend/controllers"
	"github.com/astaxie/beego"

)

func init() {

	// handler error
	beego.ErrorController(&controllers.ErrorController{})

	//get last login
	beego.Router("/chat/lastlogin", &controllers.ChatStatusController{}, "get:GetLastLogin")
	// private chat
	beego.Router("/chat/private", &controllers.ChatPrivateController{}, "get:PrivateChat")

}

