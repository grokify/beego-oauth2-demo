package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/grokify/beegoutil/demo1/conf"
)

type LogoutController struct {
	beego.Controller
	Logger *logs.BeeLogger
}

func (c *LogoutController) Get() {
	cfg := conf.NewConfig()
	log := cfg.Logger()
	c.Logger = log.Logger

	conf.InitSession(log)

	log.Info("Start Login Controller")

	c.SetSession("user", nil)
	c.SetSession("loggedIn", false)

	c.Redirect("/", 302)
}
