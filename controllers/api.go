package controllers

import (
	"bufio"
	"fmt"
	"moonlord/models"
	"time"

	"github.com/astaxie/beego"
)

type MtrasResult struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Center  models.Tra     `json:"center"`
	Zoom    int            `json:"zoom"`
	Mtras   [][]models.Tra `json:"mtras"`
}

type ApiController struct {
	beego.Controller
}

func (this *ApiController) GetTrasNum() {
	type Num struct {
		Num int
	}
	this.Data["json"] = &Num{
		Num: models.RTree.GetSize(),
	}
	this.ServeJSON()
}

func (this *ApiController) GetNearestMTrasByPoint() {

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
		this.Data["json"] = &MtrasResult{
			Status:  0,
			Message: err.Error(),
		}
		this.ServeJSON()
		return
	}

	market := &models.Tra{
		Lat: lat,
		Lon: lon,
	}
	market.LatLon2XY()
	mtras := getNearestTrasByPoint(limit, []float64{market.X, market.Y, float64(pointTime.Unix() / models.TIME_SCALE)})

	center, zoom := caclCenterAndZoom(GetMTrasMinAndMaxWithMarketNT(mtras, market))

	this.Data["json"] = &MtrasResult{
		Status: 1,
		Center: *center,
		Zoom:   zoom,
		Mtras:  mtras,
	}
	this.ServeJSON()
}

func (this *ApiController) GetSimMTrasByTras() {
	type result struct {
		Status  int            `json:"status"`
		Message string         `json:"message"`
		Tras    [][]models.Tra `json:"tras"`
	}
	file, _, err := this.GetFile("trafile")
	defer file.Close()
	if err != nil {
		this.Data["json"] = &result{
			Status:  0,
			Message: err.Error(),
		}
		this.ServeJSON()
		return
	}

	fileReader := bufio.NewReader(file)
	tras := GetTraDataFromUploadFileReader(fileReader)
	fmt.Println(tras)

	this.Data["json"] = &result{
		Status:  1,
		Message: "sucess",
		Tras:    [][]models.Tra{tras},
	}
	this.ServeJSON()
}
