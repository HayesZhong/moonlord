package controllers

import (
	"github.com/astaxie/beego"
)

type ViewController struct {
	beego.Controller
}

func (this *ViewController) Index() {
	index, _ := this.GetInt32("index", 0)
	this.Data["index"] = index
	this.TplName = "index.tpl"
}

func (this *ViewController) Nerest() {
	var timeStr string
	var lat, lon float64
	var err error
	timeStr = this.GetString("time", "")
	lat, err = this.GetFloat("lat")
	lon, err = this.GetFloat("lon")
	limit, _ := this.GetInt("limit", 1)
	if err != nil {
		return
	}

	this.Data["time"], this.Data["lat"], this.Data["lon"], this.Data["limit"] = timeStr, lat, lon, limit
	this.TplName = "nerest.tpl"
}
