package main

import (
	"github.com/astaxie/beego"
	_ "hope-pet-chat-backend/routers"
	"runtime"
)

const (
	APP_VER = "0.1.1"
)

func main() {
	runtime.GOMAXPROCS(4)
	beego.SetLogger("file", string(beego.AppConfig.String("logging")))
	beego.Info(beego.BConfig.AppName, APP_VER)
	beego.Run()
}
