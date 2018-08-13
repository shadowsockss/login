package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"login/models"
)

type LoginController struct {
	beego.Controller
}

func (index *LoginController) Get() {
	sess := index.StartSession()
	username := sess.Get("username")
	fmt.Println(username)
	if username == nil || username == "" {
		index.TplNames = "index.tpl"
	} else {
		index.TplNames = "success.tpl"
	}

}

func (index *LoginController) Post() {
	sess := index.StartSession()
	var user models.User
	inputs := index.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("username")
	user.Pwd = inputs.Get("pwd")
	err := models.ValidateUser(user)
	if err == nil {
		sess.Set("username", user.Username)
		fmt.Println("username:", sess.Get("username"))
		index.TplNames = "success.tpl"
	} else {
		fmt.Println(err)
		index.TplNames = "error.tpl"
	}
}