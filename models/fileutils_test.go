package models

//func BenchmarkGetTraData(b *testing.B) {
//	b.StartTimer()
//	mtrasNum := 200000
//	minChildren = int(mtrasNum / 100)
//	RTree = NewTree(3, minChildren, minChildren*2)

//	MTras = make([][]Tra, 0, mtrasNum)

//	trasChan := make(chan []Tra, runtime.NumCPU())
//	go GetTraDataUseChan("F:/tradata/format", mtrasNum, trasChan)
//	processTrasChan(trasChan)
//	b.StopTimer()
//}

//func BenchmarkGetTraData2(b *testing.B) {
//	b.StartTimer()
//	mtrasNum := 200000
//	GetTraData("F:/tradata/format", mtrasNum)
//	b.StopTimer()
//}

//func BenchmarkFileFormat(b *testing.B) {
//	b.StartTimer()
//	FileFormat("F:/tradata", "F:/tradata/format")
//	b.StopTimer()
//}
