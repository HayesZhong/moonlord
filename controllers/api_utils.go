package controllers

import (
	"bufio"
	"fmt"
	"moonlord/filter"
	"moonlord/models"
	"time"
)

func getNearestTrasByPoint(limit int, point []float64) [][]models.Tra {
	rects := models.RTree.NearestNeighbors(limit, point)
	mtras := make([][]models.Tra, len(rects))
	for i, rect := range rects {
		mtras[i] = rect.Bounds().GetTras()
	}
	return mtras
}

func getNearestTrasByTras(limit int, tras []models.Tra) [][]models.Tra {
	seeds := make([]models.Point, 0, 8)
	k := new(filter.KlamanFilter)
	models.FilteOneTras(tras, k)
	minx, miny, mint, maxx, maxy, maxt := models.GetTrasMinAndMax(tras)
	//设置种子point
	seeds = append(seeds, []float64{minx, miny, mint})
	seeds = append(seeds, []float64{maxx, maxy, maxt})
	seeds = append(seeds, []float64{minx, maxy, mint})
	seeds = append(seeds, []float64{maxx, miny, mint})
	seeds = append(seeds, []float64{minx, miny, maxt})
	seeds = append(seeds, []float64{maxx, maxy, mint})
	seeds = append(seeds, []float64{minx, maxy, maxt})
	seeds = append(seeds, []float64{maxx, miny, maxt})
	for i := 0; i < len(tras)/3; i++ {
		seeds = append(seeds, []float64{tras[i*3].FX, tras[i*3].FY, float64(tras[i*3].T / models.TIME_SCALE)})
	}
	//获得跟种子point相近的轨迹，并过滤掉重复的
	mtrasMap := make(map[string][]models.Tra, limit)
	for i := 0; i < len(seeds); i++ {
		mtras := getNearestTrasByPoint(6, seeds[i])
		for j := 0; j < len(mtras); j++ {
			tmpKey := fmt.Sprintf("%f%f%f%f%f%f%f%f", mtras[j][0].X, mtras[j][0].Y, mtras[j][0].T, mtras[j][len(mtras)].X, mtras[j][len(mtras)].Y, mtras[j][len(mtras)].T)
			mtrasMap[tmpKey] = mtras[j]
		}
	}
	//过滤掉不可能的(暂时先不做)

	return nil
}

func GetTraDataFromUploadFileReader(dataReader *bufio.Reader) []models.Tra {
	traData := make([]models.Tra, 0, 20)
	var x, y float64
	var dateStr, timeStr string
	var line []byte
	var err error
	var timeTmp time.Time
	for {

		line, _, err = dataReader.ReadLine()
		if err != nil {
			break
		}
		fmt.Sscanf(string(line), "%f,%f,%s %s", &x, &y, &dateStr, &timeStr)
		timeTmp, err = time.Parse("2006/01/02 15:04:05", dateStr+" "+timeStr)
		if err != nil {
			continue
		}
		traData = append(traData, models.Tra{
			Lat: x,
			Lon: y,
			T:   timeTmp.Unix(),
		})
	}
	return traData
}

//获得一组轨迹的最大最小 不包括时间 使用滤波前的xy
func GetMTrasMinAndMaxNT(mtras [][]models.Tra) (minx, miny, maxx, maxy float64) {
	minx, miny = mtras[0][0].X, mtras[0][0].Y
	maxx, maxy = mtras[0][0].X, mtras[0][0].Y
	for i := 1; i < len(mtras); i++ {
		for j := 0; j < len(mtras[i]); j++ {
			if minx > mtras[i][j].X {
				minx = mtras[i][j].X
			}
			if miny > mtras[i][j].Y {
				miny = mtras[i][j].Y
			}

			if maxx < mtras[i][j].X {
				maxx = mtras[i][j].X
			}
			if maxy < mtras[i][j].Y {
				maxy = mtras[i][j].Y
			}
		}
	}
	return
}

//获得一组轨迹和一个market的最大最小 不包括时间 使用滤波前的lon lat
func GetMTrasMinAndMaxWithMarketNT(mtras [][]models.Tra, tra *models.Tra) (minx, miny, maxx, maxy float64) {
	minx, miny, maxx, maxy = GetMTrasMinAndMaxNT(mtras)
	if tra.X < minx {
		minx = tra.X
	}
	if tra.Y < miny {
		miny = tra.Y
	}
	if tra.X > maxx {
		maxx = tra.X
	}
	if tra.Y > maxy {
		maxy = tra.Y
	}
	return
}

func caclCenterAndZoom(minx, miny, maxx, maxy float64) (center *models.Tra, zoom int) {
	center = &models.Tra{
		X: (minx + maxx) / 2,
		Y: (miny + maxy) / 2,
	}
	center.XY2LatLon()

	xzoom := caclZoom((maxx - minx) / 10)
	yzoom := caclZoom((maxy - miny) / 6)
	zoom = xzoom
	if xzoom > yzoom {
		zoom = yzoom
	}
	return
}
func caclZoom(dist float64) int {
	switch {
	case dist <= 20:
		return 19
	case dist <= 50:
		return 18
	case dist <= 100:
		return 17
	case dist <= 200:
		return 16
	case dist <= 500:
		return 15
	case dist <= 1000:
		return 14
	case dist <= 2000:
		return 13
	case dist <= 5000:
		return 12
	case dist <= 10000:
		return 11
	case dist <= 20000:
		return 10
	case dist <= 25000:
		return 9
	case dist <= 50000:
		return 8
	case dist <= 100000:
		return 7
	default:
		return 6
	}

}
