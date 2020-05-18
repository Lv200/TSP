package util

import (
	"math"
	"math/rand"
	"time"
)

func Distance(x1, y1, x2, y2 float32) float32 {
	return float32(math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2)))
}

func Shuffle(a []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

func IndexOf(a []int, x int) int {
	for i, ele := range a {
		if ele == x {
			return i
		}
	}
	return -1
}

func DeleteByValue(a []int, value int) []int {
	i := IndexOf(a, value)
	return append(a[:i], a[i+1:]...)
}

func NextIndex(a []int, index int) (nextIndex int) {
	return (index + 1) % len(a)
}

func PrevIndex(a []int, index int) (prevIndex int) {
	return (index - 1 + len(a)) % len(a)
}
