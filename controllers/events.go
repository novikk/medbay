package controllers

import (
	"github.com/astaxie/beego"
)

type EventsController struct {
	beego.Controller
}

var lastEvent string

func (c *EventsController) Add() {
	lastEvent = c.GetString("event")

	alreadyAdded := false
	foundMed := 0
	for i, med := range addedMedicines {
		if med.Type == lastEvent {
			alreadyAdded = true
			foundMed = i
		}
	}

	if !alreadyAdded {
		addedMedicines = append(addedMedicines, Medicine{
			Type:       lastEvent,
			NumDoses:   1,
			Prescribed: false,
		})
	} else {
		addedMedicines[foundMed].NumDoses++
	}
}

func (c *EventsController) Pending() {
	defer c.ServeJSON()

	if lastEvent == "" {
		c.Data["json"] = "nothing"
		return
	}

	c.Data["json"] = lastEvent
	lastEvent = ""
}
