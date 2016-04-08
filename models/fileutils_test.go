package models

import (
	"fmt"
	"testing"
)

func BenchmarkGetTraData(b *testing.B) {
	traDatas, err := GetTraData("F:/tradata/format", 100)
	if err != nil {
		fmt.Println(err)
	}
	b.StartTimer()
	for j := 0; j < len(traDatas); j++ {
		for i := 0; i < len(traDatas[j]); i++ {
			traDatas[j][i].LatLon2XY()
		}
	}
	b.StopTimer()
}

//func BenchmarkFileFormat(b *testing.B) {
//	b.StartTimer()
//	FileFormat("F:/tradata","F:/tradata/format")
//	b.StopTimer()
//}
