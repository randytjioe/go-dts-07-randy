package main

import (
	"fmt"
	"strings"
)

func main() {
	kalimat := "selamat malam"

	// memecah kalimat menjadi kata-kata
	kata := strings.Split(kalimat, " ")

	// mencetak setiap kata per karakter
	for _, k := range kata {
		for _, char := range k {
			fmt.Printf("%c\n", char)
		}
		fmt.Println()
	}

	// menghitung kemunculan tiap karakter
	freq := make(map[string]int)
	for _, char := range kalimat {
		if char != ' ' {
			freq[string(char)]++
		}
	}

	// mencetak hasil perhitungan
	fmt.Println(freq)
}
