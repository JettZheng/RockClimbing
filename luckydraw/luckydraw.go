package luckydraw

import (
	"math/rand"
	"time"
)

// 假设有八个奖品 每个奖品设置不同的概率 概率总和为100%, 随机一个奖品，该奖品的概率要和设置概率的概率一致，而不是1/8 请写出方法

type LotteryCaculator struct {
	ProbList []int64
	MaxCount int
}

func (lc *LotteryCaculator) random() int64 {
	rand.Seed(time.Now().UnixNano())
	// return rand.Int63n(100)
	return int64(96)
}

func (lc *LotteryCaculator) LoadAwardProb(input []int64) {
	var edge int64
	lc.ProbList = make([]int64, lc.MaxCount)
	for i := 0; i < lc.MaxCount; i++ {
		edge += input[i]
		lc.ProbList[i] = edge
	}
}

func (lc *LotteryCaculator) binary(l, r int, n int64) int {
	mid := (l + r) / 2

	if n < lc.ProbList[mid] {
		return lc.binary(l, mid, n)
	}
	if n > lc.ProbList[mid] {
		return lc.binary(mid, r, n)
	}
	return mid
}
//1,2,3,4,5
func BinarySearch(a []int, x int) int {
	l := 0
	r := len(a)
	for l < r {
		m := l + (r - l)/2
		if a[m] == x {
			return m
		}
		if x > a[m] {
			l = m
		}
		if x < a[m] {
			r = m - 1
		}
	}
	return -1
}

func (lc *LotteryCaculator) Gen() int {
	var fac = lc.random()
	for i := 1; i <= lc.MaxCount; i++ {
		if fac < lc.ProbList[i-1] {
			return i - 1
		}
		if fac >= lc.ProbList[i-1] && fac < lc.ProbList[i] {
			return i
		}
		if fac >= lc.ProbList[lc.MaxCount-1] {
			return i
		}
	}

	return 0
}

func (lc *LotteryCaculator) BinaryGen() int {
	var fac = lc.random()
	return lc.binary(0, lc.MaxCount, fac)
}

// [94,96,98,100,100,100,100,100]

// random a number 19

// for
// a = arr[i]  a = i+1   a>arr[i] a<arr[i+1] a = i +1
