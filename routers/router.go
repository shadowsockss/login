package routers



import (

"github.com/astaxie/beego"

"login/controllers"

)



func init() {

beego.Router("/", &controllers.LoginController{}, "GET:GET")

beego.Router("/login", &controllers.LoginController{}, "POST:POST")



beego.Router("/regist", &controllers.RegistController{}, "GET:GET")

beego.Router("/regist", &controllers.RegistController{}, "POST:POST")

}