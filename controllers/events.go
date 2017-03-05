package controllers

import (
	"time"

	"github.com/astaxie/beego"
)

type EventsController struct {
	beego.Controller
}

var lastEvent string

type EventAddResponse struct {
	Status string `json:"status"`
}

func (c *EventsController) Add() {
	defer c.ServeJSON()

	ev := c.GetString("event")
	alreadyAdded := false
	foundMed := 0
	for i, med := range addedMedicines {
		if med.Type == ev {
			alreadyAdded = true
			foundMed = i
		}
	}

	if !alreadyAdded {
		addedMedicines = append(addedMedicines, Medicine{
			Type:       ev,
			NumDoses:   []int{1},
			Prescribed: false,
			Cooldown:   8,
			LastTake:   time.Now(),
			Dose:       600,
		})
	} else {
		if time.Now().Sub(addedMedicines[foundMed].LastTake).Hours() < float64(addedMedicines[foundMed].Cooldown) {
			c.Data["json"] = EventAddResponse{"error_cooldown"}
			return
		}

		addedMedicines[foundMed].LastTake = time.Now()
		addedMedicines[foundMed].NumDoses = append(addedMedicines[foundMed].NumDoses, 1)
	}

	lastEvent = ev
	c.Data["json"] = EventAddResponse{"ok"}
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
