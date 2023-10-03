package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	data1 := []interface{}{"coba1", "coba2", "coba3"}
	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)

		// Goroutine pertama (menampilkan secara acak)
		go func(i int, data []interface{}) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			shuffledData := shuffle(data)
			for _, item := range shuffledData {
				fmt.Printf("[%v] %d\n", item, i)
			}
		}(i, data1)

		// Goroutine kedua (menampilkan secara acak)
		go func(i int, data []interface{}) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			shuffledData := shuffle(data)
			for _, item := range shuffledData {
				fmt.Printf("[%v] %d\n", item, i)
			}
		}(i, data2)
	}

	wg.Wait()
}

// Mengacak data dalam slice
func shuffle(data []interface{}) []interface{} {
	n := len(data)
	shuffledData := make([]interface{}, n)
	copy(shuffledData, data)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(n, func(i, j int) {
		shuffledData[i], shuffledData[j] = shuffledData[j], shuffledData[i]
	})
	return shuffledData
}
