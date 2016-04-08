package models

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//以后要改造成消费者生产者模式，使用多线程
func GetTraData(filePath string, num int) ([][]Tra, error) {
	dataFile, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer dataFile.Close()

	dataReader := bufio.NewReader(dataFile)

	traDataSize := 0
	if num >= 1 {
		traDataSize = num
	} else {
		traDataSize = 1000
	}
	traDatas := make([][]Tra, 0, num)

	traNowSize := 0
	isend := false
	for {
		traDatas = append(traDatas, make([]Tra, 0, 50))
		for {
			var x, y float64
			var dateStr, timeStr string
			line, _, err := dataReader.ReadLine()
			if err != nil {
				isend = true
				break
			}
			if strings.Compare("######", string(line)) == 0 {
				break
			}
			fmt.Sscanf(string(line), "%f,%f,%s %s", &x, &y, &dateStr, &timeStr)
			timeTmp, err := time.Parse("2006/01/02 15:04:05", dateStr+" "+timeStr)
			if err != nil {
				continue
			}
			traDatas[traNowSize] = append(traDatas[traNowSize], Tra{
				Lat:  x,
				Lon:  y,
				Time: timeTmp,
				Ht:   true,
			})

		}
		traNowSize++
		if isend {
			break
		}
		if num <= 0 {
			continue
		}
		if traDataSize <= traNowSize {
			break
		}
	}
	return traDatas, nil

}

func FileFormat(source string, dest string) {
	formatData, _ := os.Create(dest)
	defer formatData.Close()
	formatDataWriter := bufio.NewWriter(formatData)

	rootPathStr := source + "/Data"
	rootPath, _ := os.Open(rootPathStr)
	defer rootPath.Close()

	traNum := 0
	paths, _ := rootPath.Readdirnames(-1)
	for _, path := range paths {
		tmpPath, _ := os.Open(rootPathStr + "/" + path + "/Trajectory")
		defer tmpPath.Close()

		tmpPaths, _ := tmpPath.Readdirnames(-1)
		for _, tmpFileName := range tmpPaths {
			tmpFile, _ := os.Open(rootPathStr + "/" + path + "/Trajectory/" + tmpFileName)
			tmpFileReader := bufio.NewReader(tmpFile)
			for i := 0; i < 6; i++ {
				tmpFileReader.ReadLine()
			}
			tmpTra := ""
			lineNum := 0

			traPointNum := 0
			if rand.Intn(2) == 0 {
				traPointNum = rand.Intn(40) + 60
			} else {
				traPointNum = rand.Intn(20) + 30
			}
			for {
				if line, _, err := tmpFileReader.ReadLine(); err == nil {
					data := strings.Split(string(line), ",")
					data[5] = strings.Replace(data[5], "-", "/", -1)
					tmpTra += fmt.Sprintf("%s,%s,%s %s\n", data[0], data[1], data[5], data[6])
					lineNum++
				} else {
					break
				}
				if lineNum == traPointNum {
					formatDataWriter.WriteString(tmpTra + "######\n")
					traNum++
					formatDataWriter.Flush()
					tmpTra = ""
					lineNum = 0
				}
			}
			if lineNum >= 20 {
				formatDataWriter.WriteString(tmpTra + "######\n")
				traNum++
				formatDataWriter.Flush()
			}
			tmpFile.Close()
		}
	}
	fmt.Printf("总共处理了%d个轨迹数据\n", traNum)
}
