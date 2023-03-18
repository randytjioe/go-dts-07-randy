package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	data1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	data2 := []interface{}{"coba1", "coba2", "coba3"}

	for i := 0; i < 4; i++ {
		go func(data []interface{}, index int) {
			for j := 0; j < len(data); j++ {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				fmt.Printf("%v %d\n", data, index)
			}
		}(data1, i+1)

	}

	for i := 0; i < 4; i++ {
		go func(data []interface{}, index int) {
			for j := 0; j < len(data); j++ {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				fmt.Printf("%v %d\n", data, index)
			}
		}(data2, i+1)
	}

	time.Sleep(1 * time.Second) // tunggu 1 detik agar semua goroutine selesai
}
