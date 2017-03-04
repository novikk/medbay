package controllers

import (
	"github.com/astaxie/beego"
)

type EventsController struct {
	beego.Controller
}

var lastEvent string

func (c *EventsController) Add() {
	lastEvent = ""
	lastEvent = c.GetString("event")
}

func (c *EventsController) Pending() {
	defer c.ServeJSON()

	if lastEvent == "" {
		c.Data["json"] = "nothing"
		return
	}

	c.Data["json"] = lastEvent
}
