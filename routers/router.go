package routers

import (
	"relation-graph/controllers"
	"github.com/astaxie/beego"
)


func init() {
    beego.Router("/", &controllers.MainController{})
}
