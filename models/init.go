package models

import (
	"fmt"
	"moonlord/filter"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/c4pt0r/ini"
)

//var MTras [][]Tra
var RTree *Rtree
var minChildren int

const (
	SETTING_SECTION = "ser"
	TIME_SCALE      = 60
)

var (
	realpath, _ = exec.LookPath(os.Args[0])
	cfg         = ini.NewConf(filepath.Dir(realpath) + "\\setting.ini")

	initmode     = cfg.Int(SETTING_SECTION, "initmode", 0)
	gobfilepath  = cfg.String(SETTING_SECTION, "gobfilepath", "")
	mtrasnum     = cfg.Int(SETTING_SECTION, "mtrasnum", 0)
	trasfilepath = cfg.String(SETTING_SECTION, "trasfilepath", "")
)

func init() {
	err := cfg.Parse()
	if err != nil {
		fmt.Println(err)
	}

	minChildren = 40
	RTree = NewTree(3, minChildren, minChildren*2)

	if *initmode == 1 {
		fmt.Printf("启动参数：initmode:%d,gobfilepath:%s\n", *initmode, *gobfilepath)
		RTree = &Rtree{}
		b := time.Now().Unix()
		err := RTree.Decode(*gobfilepath)
		if err != nil {
			fmt.Printf("从持久化数据解析到内存失败:%s\n程序退出\n", err.Error())
			os.Exit(-1)
		}
		fmt.Printf("从持久化数据解析到内存成功，用时%ds\n", time.Now().Unix()-b)
	} else {
		fmt.Printf("启动参数：initmode:%d,mtrasnum:%d,trasfilepath:%s,gobfilepath:%s\n", *initmode, *mtrasnum, *trasfilepath, *gobfilepath)
		trasChan := make(chan []Tra, runtime.NumCPU())
		go GetTraDataUseChan(*trasfilepath, *mtrasnum, trasChan)
		go processTrasChan(trasChan)
	}
}

func processTrasChan(trasChan chan []Tra) {
	a := time.Now().Unix()
	k := new(filter.KlamanFilter)
	index := 0
	for tras := range trasChan {
		if index%1000 == 0 {
			fmt.Printf("已经解析轨迹数:%d\n", index)
		}
		FilteOneTras(tras, k)
		InsterTrasToRTrees(RTree, tras)
		index++

	}
	fmt.Printf("数据解析完成，解析到内存的轨迹条数：%d，用时%d s\n正在持久化到硬盘中...\n", index, time.Now().Unix()-a)
	RTree.Encode(*gobfilepath)
	fmt.Printf("持久化完毕\n")
}

var minx, miny, mint, maxx, maxy, maxt float64

func InsterTrasToRTrees(rTree *Rtree, tras []Tra) {
	minx, miny, mint, maxx, maxy, maxt = GetTrasMinAndMax(tras)
	rTree.Insert(NewRectWithTwoPoint([]float64{minx, miny, mint / TIME_SCALE}, []float64{maxx, maxy, maxt / TIME_SCALE}, tras))
}

func GetTrasMinAndMax(tras []Tra) (minx, miny, mint, maxx, maxy, maxt float64) {
	minx, miny, mint = tras[0].FX, tras[0].FY, float64(tras[0].T)
	maxx, maxy, maxt = tras[0].FX, tras[0].FY, float64(tras[0].T)
	for i := 1; i < len(tras); i++ {
		if minx > tras[i].FX {
			minx = tras[i].FX
		}
		if miny > tras[i].FY {
			miny = tras[i].FY
		}
		if mint > float64(tras[i].T) {
			mint = float64(tras[i].T)
		}
		if maxx < tras[i].FX {
			maxx = tras[i].FX
		}
		if maxy < tras[i].FY {
			maxy = tras[i].FY
		}
		if maxt < float64(tras[i].T) {
			maxt = float64(tras[i].T)
		}
	}
	return
}

func FilteMultiTras(mtras [][]Tra) {
	k := new(filter.KlamanFilter)
	for i := 0; i < len(mtras); i++ {
		FilteOneTras(mtras[i], k)
	}
}

func FilteOneTras(tras []Tra, k *filter.KlamanFilter) {
	x := make([]float64, len(tras))
	y := make([]float64, len(tras))
	t := make([]float64, len(tras))
	tras[0].LatLon2XY()
	x[0], y[0], t[0] = 0, 0, 0

	for j := 1; j < len(tras); j++ {
		tras[j].LatLon2XY()
		x[j] = tras[j].X - tras[0].X
		y[j] = tras[j].Y - tras[0].Y
		t[j] = float64(tras[j].T - tras[0].T)
	}

	k.X, k.Y, k.T, k.N = x, y, t, len(x)
	fx, fy := k.Filter()
	tras[0].FX = tras[0].X
	tras[0].FY = tras[0].Y
	tras[0].FxFy2FlatFlon()
	for j := 1; j < len(tras); j++ {
		tras[j].FX = tras[0].X + fx[j]
		tras[j].FY = tras[0].Y + fy[j]
		tras[j].FxFy2FlatFlon()
	}
}
