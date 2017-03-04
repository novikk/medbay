package controllers

import (
	"github.com/astaxie/beego"
)

type AppController struct {
	beego.Controller
}

type Medicine struct {
	Type       string
	Dose       int
	NumDoses   int
	Prescribed bool
}

var addedMedicines []Medicine

func init() {
	addedMedicines = append(addedMedicines, Medicine{
		Type:       "Amoxicillin",
		Dose:       500,
		NumDoses:   2,
		Prescribed: true,
	})
}

func (c *AppController) Tracking() {
	c.Data["medicines"] = addedMedicines

	c.TplName = "tracking.tpl"
}

func (c *AppController) Prescription() {
	c.TplName = "prescription.tpl"
}

func (c *AppController) PrescriptionManage() {
	c.TplName = "prescriptionManage.tpl"
}

func (c *AppController) TrackingDetail() {
	c.TplName = "trackingDetail.tpl"
}
