package models

import (
	"testing"
)

func TestWGS2UTM(t *testing.T) {
	x, y, zone := WGS2UTM(39.984686, 116.318417)
	min := 0.000001
	if !(x-441807.043047 <= min && y-4426279.938078 <= min && zone == 50) {
		t.Fail()
	}
}
func TestUTM2WGS(t *testing.T) {
	lon, lat := UTM2WGS(441807.043047, 4426279.938078, 50, false)
	min := 0.000001
	if !(lon-116.318417 <= min && lat-39.984686 <= min) {
		t.Fail()
	}
}
