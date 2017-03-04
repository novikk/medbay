package routers

import (
	"github.com/astaxie/beego"
	"github.com/novikk/medbay/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/events/add", &controllers.EventsController{}, "get:Add")
	beego.Router("/api/events/pending", &controllers.EventsController{}, "get:Pending")
}
