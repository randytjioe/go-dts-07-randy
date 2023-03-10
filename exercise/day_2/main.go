package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 5, 4, 7, 3, 4, 4, 6, 5, 7, 9, 9}

	// membuat map untuk menghitung jumlah kemunculan setiap value dalam array
	countMap := make(map[int]int)
	for _, num := range arr {
		countMap[num]++
	}

	// mencetak value yang kemunculannya > 1
	fmt.Print("Value dalam array yang berjumlah > 1: ")
	for num, count := range countMap {
		if count > 1 {
			fmt.Printf("%d ", num)
		}
	}
}
