package main

import (
	//"fmt"
	"github.com/astaxie/beego"
	"login/controllers"
)

type MainController struct {
	beego.Controller
}

func main() {
	beego.SessionOn = true
	beego.RegisterController("/", &controllers.IndexController{})
	beego.RegisterController("/regist", &controllers.RegistController{})
	beego.Run()
}