package luckydraw

import "testing"

func BenchmarkGetLotteryNumber(b *testing.B) {
	cases := []struct {
		name   string
		probes []int64
	}{
		{"N-4", []int64{10,20,30,40}},
	}
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var lc = LotteryCaculator{}
				lc.MaxCount = 8
				lc.LoadAwardProb(c.probes)
				lc.Gen()
			}
		})
	}
}

func BenchmarkGetBinaryLotteryNumber(b *testing.B) {
	cases := []struct {
		name   string
		probes []int64
	}{
		{"N-4", []int64{10,20,30,40}},
	}
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var lc = LotteryCaculator{}
				lc.MaxCount = 8
				lc.LoadAwardProb(c.probes)
				lc.BinaryGen()
			}
		})
	}
}

