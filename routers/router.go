package routers

import (
	"github.com/astaxie/beego"
	"github.com/novikk/medbay/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/testqr", &controllers.MainController{}, "get:TestQr")

	beego.Router("/api/events/add", &controllers.EventsController{}, "get:Add")
	beego.Router("/api/events/pending", &controllers.EventsController{}, "get:Pending")

	beego.Router("/app/tracking", &controllers.AppController{}, "get:Tracking")
	beego.Router("/app/tracking/detail", &controllers.AppController{}, "get:TrackingDetail")
	beego.Router("/app/prescription", &controllers.AppController{}, "get:Prescription")
	beego.Router("/app/prescription/manage", &controllers.AppController{}, "get:PrescriptionManage")
}
