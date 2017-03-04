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
}

func (c *EventsController) Pending() {
	if lastEvent == "" {
		return
	}

	c.Data["json"] = lastEvent
	c.ServeJSON()
}
