package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"moonlord/models"
	"time"
)

type ApiController struct {
	beego.Controller
}

func (this *ApiController) GetOneTras() {
	index, _ := this.GetInt32("index", 0)
	this.Data["json"] = &models.MTras[index]
	this.ServeJSON()
}

func (this *ApiController) GetNearestMTras() {
	var timeStr string
	var pointTime time.Time
	var lat, lon float64
	var err error
	timeStr = this.GetString("time", "")
	pointTime, err = time.Parse("2006-01-02-15-04-05", timeStr)
	lat, err = this.GetFloat("lat")
	lon, err = this.GetFloat("lon")
	limit, _ := this.GetInt("limit", 1)
	if err != nil {
		return
	}
	x, y, zone := models.WGS2UTM(lat, lon)

	mtras := getNearestTrasByPoint(zone, limit, []float64{x, y, float64(pointTime.Unix())})
	this.Data["json"] = &mtras
	this.ServeJSON()
}

func getNearestTrasByPoint(zone, limit int, point []float64) [][]models.Tra {
	fmt.Println(models.RTrees[zone].Dim)
	rects := models.RTrees[zone].NearestNeighbors(limit, point)
	mtras := make([][]models.Tra, len(rects))
	for i, rect := range rects {
		mtras[i] = rect.Bounds().GetTras()
	}
	return mtras
}
