package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	data1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	data2 := []interface{}{"coba1", "coba2", "coba3"}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(data []interface{}, index int) {
			defer wg.Done()
			for j := 0; j < len(data); j++ {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				mu.Lock()
				fmt.Printf("%v %d\n", data, index)
				mu.Unlock()
			}
		}(data1, i+1)
	}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(data []interface{}, index int) {
			defer wg.Done()
			for j := 0; j < len(data); j++ {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				mu.Lock()
				fmt.Printf("%v %d\n", data, index)
				mu.Unlock()
			}
		}(data2, i+1)
	}

	wg.Wait()
}
