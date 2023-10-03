package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	data1 := []interface{}{"coba1", "coba2", "coba3"}
	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)

		// Goroutine pertama (menampilkan secara rapih dengan mutex)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			for j := 1; j <= 3; j++ {
				fmt.Printf("[%v] %d\n", data1[j-1], i)
			}
		}(i)

		// Goroutine kedua (menampilkan secara rapih dengan mutex)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			for j := 1; j <= 3; j++ {
				fmt.Printf("[%v] %d\n", data2[j-1], i)
			}
		}(i)
	}

	wg.Wait()
}
