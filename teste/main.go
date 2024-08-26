package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sum := HighLatency(1234567890)
	end := time.Now()
	fmt.Println(sum)
	fmt.Println("Duração:", end.Sub(start))
}

func LowLatency(n int) int {
	return n * (n + 1) / 2
}

func HighLatency(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}
