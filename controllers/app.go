package controllers

import (
	"time"

	"github.com/astaxie/beego"
)

type AppController struct {
	beego.Controller
}

type Medicine struct {
	Type       string
	Dose       int
	NumDoses   []int
	Prescribed bool
	Cooldown   int
	LastTake   time.Time
}

var addedMedicines []Medicine

func init() {
	// example med
	addedMedicines = append(addedMedicines, Medicine{
		Type:       "Amoxicillin",
		Dose:       500,
		NumDoses:   []int{0, 0, 0},
		Prescribed: true,
		Cooldown:   8,
		LastTake:   time.Now().Add(-60 * 60 * 24 * time.Hour),
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
