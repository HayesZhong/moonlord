package models

import "math"

type Tra struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	T   int64   `json:"time"`
	//	Ht   bool      `json:"ht"` // 是否有时间
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	FLat float64 `json:"flat"`
	FLon float64 `json:"flon"`
	FX   float64 `json:"fx"`
	FY   float64 `json:"fy"`
}

var minFloat64 float64 = 0.00000001

func (t *Tra) Equal(r *Tra) bool {
	if math.Abs(t.X-r.X) > minFloat64 || math.Abs(t.Y-r.Y) > minFloat64 {
		return false
	}
	return true
}

func (t *Tra) LatLon2XY() {
	//	if mode == 0 {
	//		t.X, t.Y, t.Zone = WGS2UTM(t.Lat, t.Lon)
	//	} else if mode == 1 {
	t.X, t.Y = MercatorLonLat2XY(t.Lon, t.Lat)
	//	}

}
func (t *Tra) XY2LatLon() {
	//	if mode == 0 {
	//		t.FLon, t.FLat = UTM2WGS(t.FX, t.FY, t.Zone, false)
	//	} else if mode == 1 {
	t.Lon, t.Lat = MercatorXY2LonLat(t.X, t.Y)
	//	}

}

func (t *Tra) FxFy2FlatFlon() {
	//	if mode == 0 {
	//		t.FLon, t.FLat = UTM2WGS(t.FX, t.FY, t.Zone, false)
	//	} else if mode == 1 {
	t.FLon, t.FLat = MercatorXY2LonLat(t.FX, t.FY)
	//	}

}

func MercatorLonLat2XY(lon, lat float64) (x, y float64) {
	x = lon * 20037508.34 / 180
	y = math.Log(math.Tan((90+lat)*math.Pi/360)) / (math.Pi / 180)
	y = y * 20037508.34 / 180
	return
}
func MercatorXY2LonLat(x, y float64) (lon, lat float64) {
	lon = x / 20037508.34 * 180
	lat = y / 20037508.34 * 180
	lat = 180 / math.Pi * (2*math.Atan(math.Exp(lat*math.Pi/180)) - math.Pi/2)
	return
}

//func (t *Tra) LatLon2XY() {
//	t.X, t.Y, t.Zone = WGS2UTM(t.Lat, t.Lon)
//}

//func (t *Tra) XY2LatLon() {
//	t.FLon, t.FLat = UTM2WGS(t.FX, t.FY, t.Zone, false)
//}

//const (
//	SM_A      = 6378137.0
//	SM_B      = 6356752.314
//	UTM_SCALE = 0.9996
//)

////得到的结果是: x坐标,y坐标
//func WGS2UTM(lat, lon float64) (x, y float64, zone int8) {
//	zone = int8(math.Floor((lon+180.0)/6)) + 1
//	cm := uTMCentralMeridian(int(zone))

//	x, y = mapLatLonToXY(lat/180.0*math.Pi, lon/180*math.Pi, cm)
//	/* Adjust easting and northing for UTM system. */
//	x = x*UTM_SCALE + 500000.0
//	y = y * UTM_SCALE
//	if y < 0.0 {
//		y += 10000000.0
//	}
//	return x, y, zone
//}

//func uTMCentralMeridian(zone int) float64 {
//	deg := float64(-183.0 + (zone * 6.0))
//	cmeridian := deg / 180.0 * math.Pi
//	return cmeridian
//}
//func uTMCentralMeridian2(zone int) float64 {
//	deg := float64(-183.0 + (zone * 6.0))
//	cmeridian := deg / math.Pi * 180.0
//	return cmeridian
//}

//func mapLatLonToXY(phi, lambda, lambda0 float64) (x, y float64) {
//	var N, nu2, ep2, t, t2, l float64
//	var l3coef, l4coef, l5coef, l6coef, l7coef, l8coef float64
//	//var tmp float64

//	/* Precalculate ep2 */
//	ep2 = (math.Pow(SM_A, 2.0) - math.Pow(SM_B, 2.0)) / math.Pow(SM_B, 2.0)
//	/* Precalculate nu2 */
//	nu2 = ep2 * math.Pow(math.Cos(phi), 2.0)
//	/* Precalculate N */
//	N = math.Pow(SM_A, 2.0) / (SM_B * math.Sqrt(1+nu2))
//	/* Precalculate t */
//	t = math.Tan(phi)
//	t2 = t * t
//	//tmp = (t2 * t2 * t2) - math.Pow(t, 6.0)
//	/* Precalculate l */
//	l = lambda - lambda0

//	/* Precalculate coefficients for l**n in the equations below
//	   so a normal human being can read the expressions for easting
//	   and northing
//	   -- l**1 and l**2 have coefficients of 1.0 */
//	l3coef = 1.0 - t2 + nu2
//	l4coef = 5.0 - t2 + 9*nu2 + 4.0*(nu2*nu2)
//	l5coef = 5.0 - 18.0*t2 + (t2 * t2) + 14.0*nu2 - 58.0*t2*nu2
//	l6coef = 61.0 - 58.0*t2 + (t2 * t2) + 270.0*nu2 - 330.0*t2*nu2
//	l7coef = 61.0 - 479.0*t2 + 179.0*(t2*t2) - (t2 * t2 * t2)
//	l8coef = 1385.0 - 3111.0*t2 + 543.0*(t2*t2) - (t2 * t2 * t2)

//	/* Calculate easting (x) */
//	x = N * math.Cos(phi) * l
//	x += (N / 6.0 * math.Pow(math.Cos(phi), 3.0) * l3coef * math.Pow(l, 3.0))
//	x += (N / 120.0 * math.Pow(math.Cos(phi), 5.0) * l5coef * math.Pow(l, 5.0))
//	x += (N / 5040.0 * math.Pow(math.Cos(phi), 7.0) * l7coef * math.Pow(l, 7.0))

//	/* Calculate northing (y) */
//	y = arcLengthOfMeridian(phi)
//	y += (t / 2.0 * N * math.Pow(math.Cos(phi), 2.0) * math.Pow(l, 2.0))
//	y += (t / 24.0 * N * math.Pow(math.Cos(phi), 4.0) * l4coef * math.Pow(l, 4.0))
//	y += (t / 720.0 * N * math.Pow(math.Cos(phi), 6.0) * l6coef * math.Pow(l, 6.0))
//	y += (t / 40320.0 * N * math.Pow(math.Cos(phi), 8.0) * l8coef * math.Pow(l, 8.0))

//	return x, y
//}

//func arcLengthOfMeridian(phi float64) float64 {
//	var alpha, beta, gamma, delta, epsilon, n float64
//	var result float64

//	/* Precalculate n */
//	n = (SM_A - SM_B) / (SM_A + SM_B)
//	/* Precalculate alpha */
//	alpha = ((SM_A + SM_B) / 2.0) * (1.0 + (math.Pow(n, 2.0) / 4.0) + (math.Pow(n, 4.0) / 64.0))
//	/* Precalculate beta */
//	beta = (-3.0 * n / 2.0) + (9.0 * math.Pow(n, 3.0) / 16.0) + (-3.0 * math.Pow(n, 5.0) / 32.0)
//	/* Precalculate gamma */
//	gamma = (15.0 * math.Pow(n, 2.0) / 16.0) + (-15.0 * math.Pow(n, 4.0) / 32.0)
//	/* Precalculate delta */
//	delta = (-35.0 * math.Pow(n, 3.0) / 48.0) + (105.0 * math.Pow(n, 5.0) / 256.0)
//	/* Precalculate epsilon */
//	epsilon = (315.0 * math.Pow(n, 4.0) / 512.0)
//	/* Now calculate the sum of the series and return */
//	result = alpha * (phi + (beta * math.Sin(2.0*phi)) + (gamma * math.Sin(4.0*phi)) + (delta * math.Sin(6.0*phi)) + (epsilon * math.Sin(8.0*phi)))
//	return result
//}

//func footpointLatitude(y float64) float64 {
//	var y_, alpha_, beta_, gamma_, delta_, epsilon_, n float64
//	var result float64

//	/* Precalculate n (Eq. 10.18) */
//	n = (SM_A - SM_B) / (SM_A + SM_B)
//	/* Precalculate alpha_ (Eq. 10.22) */
//	/* (Same as alpha in Eq. 10.17) */
//	alpha_ = ((SM_A + SM_B) / 2.0) * (1 + (math.Pow(n, 2.0) / 4) + (math.Pow(n, 4.0) / 64))
//	/* Precalculate y_ (Eq. 10.23) */
//	y_ = y / alpha_
//	/* Precalculate beta_ (Eq. 10.22) */
//	beta_ = (3.0 * n / 2.0) + (-27.0 * math.Pow(n, 3.0) / 32.0) + (269.0 * math.Pow(n, 5.0) / 512.0)
//	/* Precalculate gamma_ (Eq. 10.22) */
//	gamma_ = (21.0 * math.Pow(n, 2.0) / 16.0) + (-55.0 * math.Pow(n, 4.0) / 32.0)
//	/* Precalculate delta_ (Eq. 10.22) */
//	delta_ = (151.0 * math.Pow(n, 3.0) / 96.0) + (-417.0 * math.Pow(n, 5.0) / 128.0)
//	/* Precalculate epsilon_ (Eq. 10.22) */
//	epsilon_ = (1097.0 * math.Pow(n, 4.0) / 512.0)
//	/* Now calculate the sum of the series (Eq. 10.21) */
//	result = y_ + (beta_ * math.Sin(2.0*y_)) + (gamma_ * math.Sin(4.0*y_)) + (delta_ * math.Sin(6.0*y_)) + (epsilon_ * math.Sin(8.0*y_))
//	return result
//}
//func mapXYToLatLon(x, y, lambda0 float64) (lon, lat float64) {
//	var phif, Nf, Nfpow, nuf2, ep2, tf, tf2, tf4, cf float64
//	var x1frac, x2frac, x3frac, x4frac, x5frac, x6frac, x7frac, x8frac float64
//	var x2poly, x3poly, x4poly, x5poly, x6poly, x7poly, x8poly float64

//	/* Get the value of phif, the footpoint latitude. */
//	phif = footpointLatitude(y)
//	/* Precalculate ep2 */
//	ep2 = (math.Pow(SM_A, 2.0) - math.Pow(SM_B, 2.0)) / math.Pow(SM_B, 2.0)

//	/* Precalculate cos (phif) */
//	cf = math.Cos(phif)
//	/* Precalculate nuf2 */
//	nuf2 = ep2 * math.Pow(cf, 2.0)
//	/* Precalculate Nf and initialize Nfpow */
//	Nf = math.Pow(SM_A, 2.0) / (SM_B * math.Sqrt(1+nuf2))
//	Nfpow = Nf
//	/* Precalculate tf */
//	tf = math.Tan(phif)
//	tf2 = tf * tf
//	tf4 = tf2 * tf2

//	/* Precalculate fractional coefficients for x**n in the equations
//	   below to simplify the expressions for latitude and longitude. */
//	x1frac = 1.0 / (Nfpow * cf)

//	Nfpow *= Nf /* now equals Nf**2) */
//	x2frac = tf / (2.0 * Nfpow)

//	Nfpow *= Nf /* now equals Nf**3) */
//	x3frac = 1.0 / (6.0 * Nfpow * cf)

//	Nfpow *= Nf /* now equals Nf**4) */
//	x4frac = tf / (24.0 * Nfpow)

//	Nfpow *= Nf /* now equals Nf**5) */
//	x5frac = 1.0 / (120.0 * Nfpow * cf)

//	Nfpow *= Nf /* now equals Nf**6) */
//	x6frac = tf / (720.0 * Nfpow)

//	Nfpow *= Nf /* now equals Nf**7) */
//	x7frac = 1.0 / (5040.0 * Nfpow * cf)

//	Nfpow *= Nf /* now equals Nf**8) */
//	x8frac = tf / (40320.0 * Nfpow)

//	/* Precalculate polynomial coefficients for x**n.
//	   -- x**1 does not have a polynomial coefficient. */
//	x2poly = -1.0 - nuf2
//	x3poly = -1.0 - 2*tf2 - nuf2
//	x4poly = 5.0 + 3.0*tf2 + 6.0*nuf2 - 6.0*tf2*nuf2 - 3.0*(nuf2*nuf2) - 9.0*tf2*(nuf2*nuf2)
//	x5poly = 5.0 + 28.0*tf2 + 24.0*tf4 + 6.0*nuf2 + 8.0*tf2*nuf2
//	x6poly = -61.0 - 90.0*tf2 - 45.0*tf4 - 107.0*nuf2 + 162.0*tf2*nuf2
//	x7poly = -61.0 - 662.0*tf2 - 1320.0*tf4 - 720.0*(tf4*tf2)
//	x8poly = 1385.0 + 3633.0*tf2 + 4095.0*tf4 + 1575*(tf4*tf2)

//	/* Calculate latitude */
//	lat = phif + x2frac*x2poly*(x*x) + x4frac*x4poly*math.Pow(x, 4.0) + x6frac*x6poly*math.Pow(x, 6.0) + x8frac*x8poly*math.Pow(x, 8.0)

//	/* Calculate longitude */
//	lon = lambda0 + x1frac*x + x3frac*x3poly*math.Pow(x, 3.0) + x5frac*x5poly*math.Pow(x, 5.0) + x7frac*x7poly*math.Pow(x, 7.0)
//	return lon * 57.29577951308409, lat * 57.29577951308409
//}

//func UTM2WGS(x, y float64, zone int8, southhemi bool) (float64, float64) {

//	x -= 500000.0
//	x /= UTM_SCALE

//	/* If in southern hemisphere, adjust y accordingly. */
//	if southhemi {
//		y -= 10000000.0
//	}

//	y /= UTM_SCALE

//	cmeridian := uTMCentralMeridian(int(zone))
//	return mapXYToLatLon(x, y, cmeridian)
//}
